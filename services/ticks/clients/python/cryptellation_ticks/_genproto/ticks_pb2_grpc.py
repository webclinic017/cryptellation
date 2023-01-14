# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import ticks_pb2 as ticks__pb2


class TicksServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Register = channel.unary_unary(
                '/ticks.TicksService/Register',
                request_serializer=ticks__pb2.RegisterRequest.SerializeToString,
                response_deserializer=ticks__pb2.RegisterResponse.FromString,
                )
        self.Unregister = channel.unary_unary(
                '/ticks.TicksService/Unregister',
                request_serializer=ticks__pb2.UnregisterRequest.SerializeToString,
                response_deserializer=ticks__pb2.UnregisterResponse.FromString,
                )


class TicksServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Register(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Unregister(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_TicksServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Register': grpc.unary_unary_rpc_method_handler(
                    servicer.Register,
                    request_deserializer=ticks__pb2.RegisterRequest.FromString,
                    response_serializer=ticks__pb2.RegisterResponse.SerializeToString,
            ),
            'Unregister': grpc.unary_unary_rpc_method_handler(
                    servicer.Unregister,
                    request_deserializer=ticks__pb2.UnregisterRequest.FromString,
                    response_serializer=ticks__pb2.UnregisterResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'ticks.TicksService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class TicksService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Register(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ticks.TicksService/Register',
            ticks__pb2.RegisterRequest.SerializeToString,
            ticks__pb2.RegisterResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Unregister(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ticks.TicksService/Unregister',
            ticks__pb2.UnregisterRequest.SerializeToString,
            ticks__pb2.UnregisterResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
