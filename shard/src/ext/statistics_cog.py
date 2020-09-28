from datetime import timedelta, datetime
from collections import Counter, defaultdict
from concurrent.futures import ThreadPoolExecutor

from discord.ext import commands
from discord import Forbidden, HTTPException
import discord
import json
import aiohttp
import pytz
from typing import Dict, List, Counter as TCounter
from string import punctuation

import src.generate.wordcount as wordcount_gen
from src.generate import corona, member_growth
from lib.config import DISCORD_EPOCH, logger
from lib.ipc import manager_pb2 as message_type
from src.utils import mention_to_name


class GuildData:

    def __init__(self, bot, guild, dictionary, time_granularity=timedelta(days=1)):
        self.bot = bot
        self._up_to_date_after = pytz.utc.localize(datetime.now())
        self.guild = guild
        self.dictionary, self.stops = dictionary
        self.forbidden = False
        self.time_granularity = time_granularity

        self._join_dates = []

        self._last_activity = {}
        self._message_count = Counter()
        self.correct_word_count = Counter()
        self.words = defaultdict(Counter)
        self.correct_words = defaultdict(Counter)
        self.member_words = defaultdict(Counter)
        self.mentions = defaultdict(Counter)
        self.members = defaultdict(Counter)
        self.channels = Counter()
        self.times = defaultdict(lambda: defaultdict(lambda: defaultdict(int)))

    def count_correct(self, string):
        '''returns the number of correctly spelled words in a string'''
        return len([w for w in string.split() if w in self.dictionary or w.upper() in ('A', 'I')])

    def _filter_words(self, msg: discord.Message, words: List[str]):
        if msg.author == self.bot.user:
            return []
        filtered = []
        for w in words:
            if w in self.stops:
                continue
            elif w in punctuation:
                continue
            else:
                filtered.append(w)
        return filtered

    def _merge_counts(self, count: TCounter[TCounter[int]], member: discord.Member) -> TCounter[int]:
        combined = Counter()
        for ch_id, count in count.items():
            ch = self.guild.get_channel(ch_id)
            if ch is None:
                continue
            perms = ch.permissions_for(member)
            if perms is not None and perms.read_messages and perms.read_message_history:
                combined.update(count)
        return combined

    @property
    def up_to_date(self):
        return self._up_to_date_after == DISCORD_EPOCH

    @property
    def up_to_date_after(self):
        return self._up_to_date_after

    @up_to_date_after.setter
    def up_to_date_after(self, after: datetime):
        if after.tzinfo is None:
            after = pytz.utc.localize(after)
        self._up_to_date_after = max(DISCORD_EPOCH, after)

    @property
    def join_dates(self):
        if self._join_dates is None:
            self._join_dates = tuple(m.created_at for m in self.guild.members)
        return self._join_dates

    async def process_message(self, msg):
        ch = msg.channel
        self._message_count[ch.id] += 1
        self._last_activity[ch.id] = msg.created_at

        words = [w.lower() for w in msg.content.split()]
        self.correct_word_count[ch.id] += self.count_correct(msg.content)
        self.words[ch.id].update(self._filter_words(msg, words))
        self.correct_words[ch.id][msg.author.id] += self.count_correct(msg.content)
        self.member_words[ch.id][msg.author.id] += len(words)

        date = msg.created_at - ((pytz.utc.localize(msg.created_at) - DISCORD_EPOCH) % self.time_granularity)
        self.times[ch.id][date][msg.author.id] += 1

        self.mentions[ch.id].update([m.id for m in msg.mentions])

        self.members[ch.id][msg.author.id] += 1

        self.channels[ch.id] += 1

    def architus_count(self, member: discord.Member):
        return self._merge_counts(self.members, member).get(self.bot.user.id, 0)

    @property
    def member_count(self):
        return self.guild.member_count

    def times_as_strings(self, member: discord.Member):
        combined = {}
        for ch_id, count in self.times.items():
            ch = self.guild.get_channel(ch_id)
            if ch is None:
                continue
            perms = ch.permissions_for(member)
            if perms is not None and perms.read_messages and perms.read_message_history:
                combined.update(count)
        return {k.isoformat(): v for k, v in combined.items()}

    def channel_counts(self, member: discord.Member):
        counts = {}
        for ch_id, count in self.channels.items():
            ch = self.guild.get_channel(ch_id)
            if ch is None:
                continue
            perms = ch.permissions_for(member)
            if perms is not None and perms.read_messages and perms.read_message_history:
                counts[ch_id] = count
        return counts

    def message_count(self, member: discord.Member):
        msgs = 0
        for ch_id, count in self.channels.items():
            ch = self.guild.get_channel(ch_id)
            if ch is None:
                continue
            perms = ch.permissions_for(member)
            if perms is not None and perms.read_messages and perms.read_message_history:
                msgs += count
        return msgs

    def last_activity(self, member: discord.Member):
        last = DISCORD_EPOCH.replace(tzinfo=None)
        for ch_id, time in self._last_activity.items():
            ch = self.guild.get_channel(ch_id)
            if ch is None:
                continue
            perms = ch.permissions_for(member)
            if perms is not None and perms.read_messages and perms.read_message_history:
                if time > last:
                    last = time
        return last

    def member_counts(self, member: discord.Member):
        return dict(self._merge_counts(self.members, member))

    def mention_counts(self, member: discord.Member):
        return dict(self._merge_counts(self.mentions, member))

    def mention_count(self, member: discord.Member):
        return sum(self._merge_counts(self.mentions, member).values())

    def word_count(self, member: discord.Member):
        return sum(self._merge_counts(self.words, member).values())

    def word_counts(self, member: discord.Member):
        return dict(self._merge_counts(self.member_words, member))

    def common_words(self, member: discord.Member):
        words = self._merge_counts(self.words, member).most_common(75)
        for i, pair in enumerate(words):
            try:
                name = mention_to_name(self.guild, pair[0])
                words[i] = (name, pair[1])
            except ValueError:
                continue
        return words


class MessageStats(commands.Cog, name="Server Statistics"):

    def __init__(self, bot):
        self.bot = bot
        self.cache = {}  # type: Dict[int, GuildData]
        with open('res/words/words.json') as f:
            self.dictionary = json.loads(f.read())
        with open('res/words/stops.json') as f:
            self.stops = json.loads(f.read())

    async def cache_guilds_history(self):
        while not all(d.up_to_date for d in self.cache.values()):
            for guild_d in self.cache.values():
                if guild_d.up_to_date:
                    continue
                before = guild_d.up_to_date_after
                after = max(before - timedelta(days=30), DISCORD_EPOCH)
                msgs = []
                for ch in guild_d.guild.text_channels:
                    try:
                        async for msg in ch.history(
                                before=before.replace(tzinfo=None), after=after.replace(tzinfo=None)):
                            msgs.append(msg)
                    except Forbidden:
                        guild_d.forbidden = True
                    except HTTPException:
                        logger.exception("error while downloading '{guild.name}.{channel.name}'")
                        break
                else:
                    for msg in msgs:
                        await guild_d.process_message(msg)
                    guild_d.up_to_date_after = after

    def append_warning(self, data: GuildData, em: discord.Embed):
        if not data.up_to_date:
            em.set_footer(
                text=f"Still indexing data before {data.up_to_date_after}",
                icon_url="https://emojipedia-us.s3.dualstack.us-west-1"
                         ".amazonaws.com/thumbs/120/twitter/259/warning_26a0.png")

    @commands.Cog.listener()
    async def on_ready(self):
        logger.debug(f"Caching messages for {len(self.bot.guilds)} guilds...")
        self.cache = {g.id: GuildData(self.bot, g, (self.dictionary, self.stops)) for g in self.bot.guilds}
        await self.cache_guilds_history()
        logger.debug(f"Message cache up-to-date for {len(self.bot.guilds)} guilds...")

    @commands.Cog.listener()
    async def on_message(self, msg):
        if not msg.channel.guild:
            return
        await self.cache[msg.channel.guild.id].process_message(msg)

    @commands.Cog.listener()
    async def on_member_update(self, before, after):
        if before != self.bot.user:
            return
        if before.guild_permissions != after.guild_permissions:
            self.cache[before.guild.id] = GuildData(self.bot, before.guild, (self.dictionary, self.stops))
            await self.cache_guilds_history()

    @commands.Cog.listener()
    async def on_guild_join(self, guild):
        await self.cache_guilds_history()

    @commands.command(aliases=['exclude'])
    async def optout(self, ctx):
        """Prevents Architus from displaying statistics about you
        run again to reallow collection
        """
        settings = self.bot.settings[ctx.guild]
        excludes = settings.stats_exclude
        author = ctx.author
        if author.id in excludes:
            excludes.remove(author.id)
            await ctx.send(f"{author.display_name}'s message data is now available")
        else:
            excludes.append(author.id)
            await ctx.send(f"{author.display_name}'s message data is now hidden")
        settings.stats_exclude = excludes

    @commands.command(aliases=['growth'])
    async def joins(self, ctx):
        img = member_growth.generate(ctx.guild.members)
        data = await self.bot.manager_client.publish_file(iter([message_type.File(file=img)]))
        em = discord.Embed(title="Server Growth", description=ctx.guild.name)
        em.set_image(url=data.url)
        em.color = 0x35a125
        em.set_footer(text=f"{ctx.guild.name} has a total of {ctx.guild.member_count} members")
        await ctx.channel.send(embed=em)

    @commands.command()
    async def spellcheck(self, ctx, victim: discord.Member):
        '''Checks the spelling of a user'''
        if victim.id in self.bot.settings[ctx.guild].stats_exclude:
            await ctx.send(f"Sorry, {victim.display_name} has requested that their stats not be recorded :confused:")
            return
        data = self.cache[ctx.guild.id]
        words = data.member_words[victim.id]
        ratio = data.correct_words[victim.id] / (words or 1) * 100
        msg = f"{ratio:.1f}% of the {words:,} words sent by {victim.display_name} are spelled correctly."
        em = discord.Embed(title="Spellcheck", description=msg, color=0x03fc8c)
        em.set_author(name=victim.display_name, icon_url=victim.avatar_url)
        self.append_warning(data, em)
        await ctx.send(embed=em)

    @commands.command()
    async def messagecount(self, ctx, victim: discord.Member = None):
        data = self.cache[ctx.guild.id]

        with ThreadPoolExecutor() as pool:
            img = await self.bot.loop.run_in_executor(
                pool, wordcount_gen.generate, ctx.guild, data.members, data.member_words, victim)
        resp = await self.bot.manager_client.publish_file(
            iter([message_type.File(file=img)]))

        em = discord.Embed(title="Top 5 Message Senders", description=ctx.guild.name, color=0x7b8fb7)
        em.set_image(url=resp.url)
        if victim:
            if victim.id in self.bot.settings[ctx.guild].stats_exclude:
                em.set_footer(text=f"{victim.display_name} has hidden their stats")
            else:
                em.set_footer(text="{0} has sent {1:,} words across {2:,} messages".format(
                    victim.display_name, data.member_words[victim], data.members[victim]), icon_url=victim.avatar_url)
        self.append_warning(data, em)

        await ctx.channel.send(embed=em)


class CoronaStats(commands.Cog, name="Coronavirus Data"):
    def __init__(self, bot):
        self.bot = bot

    @commands.command(aliases=['covid', 'covid-19', 'covid19', 'coronavirus'])
    async def corona(self, ctx, deaths_only: bool = False):
        async with aiohttp.ClientSession() as session:
            url = 'http://coronavirusapi.com/time_series.csv'
            async with session.get(url) as resp:
                text = await resp.text()
                parsed = (e.split(',') for e in text.split('\n')[1:-1])

                img = corona.generate(parsed, deaths_only)

                data = await self.bot.manager_client.publish_file(
                    iter([message_type.File(file=img)]))

                em = discord.Embed(title="Coronavirus in the US", description="More Information",
                                   url="https://www.cdc.gov/coronavirus/2019-ncov/index.html")
                em.set_image(url=data.url)
                em.color = 0x7b8fb7
                em.set_footer(text="data collected from state websites by http://coronavirusapi.com")

                await ctx.channel.send(embed=em)


def setup(bot):
    bot.add_cog(MessageStats(bot))
    bot.add_cog(CoronaStats(bot))
