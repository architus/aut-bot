pub const SOMETHING: u8 = 2;

/// Utility module for dealing with discord snowflakes
pub mod snowflake {
    use std::time::{Duration, SystemTime};
    pub const DISCORD_EPOCH_OFFSET: u64 = 1_420_070_400_000;

    pub const TIMESTAMP_BIT_OFFSET: usize = 22;

    pub const fn extract_timestamp(snowflake: u64) -> u64 {
        (snowflake >> TIMESTAMP_BIT_OFFSET) + DISCORD_EPOCH_OFFSET
    }

    /// Gets a system time from the snowflake. This function may panic
    /// if the resulting time cannot be represented by a SystemTime.
    /// That would imply that some snowflake occurred well over
    /// 200_000_000 years into the future so I'm not going to worry
    /// about that and just return a SystemTime and not an
    /// Option<SystemTime>.
    pub fn extract_system_time(snowflake: u64) -> SystemTime {
        let millis = extract_timestamp(snowflake);
        let dur = Duration::from_millis(millis);

        SystemTime::UNIX_EPOCH + dur
    }

    /// Takes a unix timestamp and converts it to a basic snowflake that
    /// can be used as a bound in requests.
    pub fn bound_from_ts(timestamp: u64) -> u64 {
        (timestamp - DISCORD_EPOCH_OFFSET) << TIMESTAMP_BIT_OFFSET
    }
}