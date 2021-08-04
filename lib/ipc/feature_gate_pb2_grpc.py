# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import lib.ipc.feature_gate_pb2 as feature__gate__pb2


class FeatureGateStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateFeature = channel.unary_unary(
                '/featuregate.FeatureGate/CreateFeature',
                request_serializer=feature__gate__pb2.Feature.SerializeToString,
                response_deserializer=feature__gate__pb2.CreationResult.FromString,
                )
        self.CheckOpenness = channel.unary_unary(
                '/featuregate.FeatureGate/CheckOpenness',
                request_serializer=feature__gate__pb2.FeatureName.SerializeToString,
                response_deserializer=feature__gate__pb2.OpennessResult.FromString,
                )
        self.AddGuildFeature = channel.unary_unary(
                '/featuregate.FeatureGate/AddGuildFeature',
                request_serializer=feature__gate__pb2.FeatureAddition.SerializeToString,
                response_deserializer=feature__gate__pb2.AddResult.FromString,
                )
        self.RemoveGuildFeature = channel.unary_unary(
                '/featuregate.FeatureGate/RemoveGuildFeature',
                request_serializer=feature__gate__pb2.FeatureRemoval.SerializeToString,
                response_deserializer=feature__gate__pb2.RemoveResult.FromString,
                )
        self.CheckGuildFeature = channel.unary_unary(
                '/featuregate.FeatureGate/CheckGuildFeature',
                request_serializer=feature__gate__pb2.GuildFeature.SerializeToString,
                response_deserializer=feature__gate__pb2.FeatureResult.FromString,
                )
        self.BatchCheckGuildFeatures = channel.unary_unary(
                '/featuregate.FeatureGate/BatchCheckGuildFeatures',
                request_serializer=feature__gate__pb2.BatchCheck.SerializeToString,
                response_deserializer=feature__gate__pb2.BatchCheckResult.FromString,
                )
        self.GetFeatures = channel.unary_stream(
                '/featuregate.FeatureGate/GetFeatures',
                request_serializer=feature__gate__pb2.FeatureList.SerializeToString,
                response_deserializer=feature__gate__pb2.Feature.FromString,
                )
        self.GetGuildFeatures = channel.unary_stream(
                '/featuregate.FeatureGate/GetGuildFeatures',
                request_serializer=feature__gate__pb2.Guild.SerializeToString,
                response_deserializer=feature__gate__pb2.Feature.FromString,
                )


class FeatureGateServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CreateFeature(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CheckOpenness(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddGuildFeature(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RemoveGuildFeature(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CheckGuildFeature(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def BatchCheckGuildFeatures(self, request, context):
        """The number of supported guilds to check at once is at least 256, but may be more.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetFeatures(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetGuildFeatures(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_FeatureGateServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateFeature': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateFeature,
                    request_deserializer=feature__gate__pb2.Feature.FromString,
                    response_serializer=feature__gate__pb2.CreationResult.SerializeToString,
            ),
            'CheckOpenness': grpc.unary_unary_rpc_method_handler(
                    servicer.CheckOpenness,
                    request_deserializer=feature__gate__pb2.FeatureName.FromString,
                    response_serializer=feature__gate__pb2.OpennessResult.SerializeToString,
            ),
            'AddGuildFeature': grpc.unary_unary_rpc_method_handler(
                    servicer.AddGuildFeature,
                    request_deserializer=feature__gate__pb2.FeatureAddition.FromString,
                    response_serializer=feature__gate__pb2.AddResult.SerializeToString,
            ),
            'RemoveGuildFeature': grpc.unary_unary_rpc_method_handler(
                    servicer.RemoveGuildFeature,
                    request_deserializer=feature__gate__pb2.FeatureRemoval.FromString,
                    response_serializer=feature__gate__pb2.RemoveResult.SerializeToString,
            ),
            'CheckGuildFeature': grpc.unary_unary_rpc_method_handler(
                    servicer.CheckGuildFeature,
                    request_deserializer=feature__gate__pb2.GuildFeature.FromString,
                    response_serializer=feature__gate__pb2.FeatureResult.SerializeToString,
            ),
            'BatchCheckGuildFeatures': grpc.unary_unary_rpc_method_handler(
                    servicer.BatchCheckGuildFeatures,
                    request_deserializer=feature__gate__pb2.BatchCheck.FromString,
                    response_serializer=feature__gate__pb2.BatchCheckResult.SerializeToString,
            ),
            'GetFeatures': grpc.unary_stream_rpc_method_handler(
                    servicer.GetFeatures,
                    request_deserializer=feature__gate__pb2.FeatureList.FromString,
                    response_serializer=feature__gate__pb2.Feature.SerializeToString,
            ),
            'GetGuildFeatures': grpc.unary_stream_rpc_method_handler(
                    servicer.GetGuildFeatures,
                    request_deserializer=feature__gate__pb2.Guild.FromString,
                    response_serializer=feature__gate__pb2.Feature.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'featuregate.FeatureGate', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class FeatureGate(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CreateFeature(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/featuregate.FeatureGate/CreateFeature',
            feature__gate__pb2.Feature.SerializeToString,
            feature__gate__pb2.CreationResult.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CheckOpenness(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/featuregate.FeatureGate/CheckOpenness',
            feature__gate__pb2.FeatureName.SerializeToString,
            feature__gate__pb2.OpennessResult.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddGuildFeature(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/featuregate.FeatureGate/AddGuildFeature',
            feature__gate__pb2.FeatureAddition.SerializeToString,
            feature__gate__pb2.AddResult.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def RemoveGuildFeature(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/featuregate.FeatureGate/RemoveGuildFeature',
            feature__gate__pb2.FeatureRemoval.SerializeToString,
            feature__gate__pb2.RemoveResult.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CheckGuildFeature(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/featuregate.FeatureGate/CheckGuildFeature',
            feature__gate__pb2.GuildFeature.SerializeToString,
            feature__gate__pb2.FeatureResult.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def BatchCheckGuildFeatures(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/featuregate.FeatureGate/BatchCheckGuildFeatures',
            feature__gate__pb2.BatchCheck.SerializeToString,
            feature__gate__pb2.BatchCheckResult.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetFeatures(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/featuregate.FeatureGate/GetFeatures',
            feature__gate__pb2.FeatureList.SerializeToString,
            feature__gate__pb2.Feature.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetGuildFeatures(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/featuregate.FeatureGate/GetGuildFeatures',
            feature__gate__pb2.Guild.SerializeToString,
            feature__gate__pb2.Feature.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
