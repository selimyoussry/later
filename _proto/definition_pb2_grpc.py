# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import definition_pb2 as definition__pb2


class LaterStub(object):

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.CreateInstance = channel.unary_unary(
        '/hippoai.later.Later/CreateInstance',
        request_serializer=definition__pb2.CreateInstanceInput.SerializeToString,
        response_deserializer=definition__pb2.CreateInstanceOutput.FromString,
        )
    self.AbortInstance = channel.unary_unary(
        '/hippoai.later.Later/AbortInstance',
        request_serializer=definition__pb2.AbortInstanceInput.SerializeToString,
        response_deserializer=definition__pb2.AbortInstanceOutput.FromString,
        )
    self.GetInstances = channel.unary_unary(
        '/hippoai.later.Later/GetInstances',
        request_serializer=definition__pb2.GetInstancesInput.SerializeToString,
        response_deserializer=definition__pb2.GetInstancesOutput.FromString,
        )
    self.Stats = channel.unary_unary(
        '/hippoai.later.Later/Stats',
        request_serializer=definition__pb2.StatsInput.SerializeToString,
        response_deserializer=definition__pb2.StatsOutput.FromString,
        )


class LaterServicer(object):

  def CreateInstance(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def AbortInstance(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetInstances(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Stats(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_LaterServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'CreateInstance': grpc.unary_unary_rpc_method_handler(
          servicer.CreateInstance,
          request_deserializer=definition__pb2.CreateInstanceInput.FromString,
          response_serializer=definition__pb2.CreateInstanceOutput.SerializeToString,
      ),
      'AbortInstance': grpc.unary_unary_rpc_method_handler(
          servicer.AbortInstance,
          request_deserializer=definition__pb2.AbortInstanceInput.FromString,
          response_serializer=definition__pb2.AbortInstanceOutput.SerializeToString,
      ),
      'GetInstances': grpc.unary_unary_rpc_method_handler(
          servicer.GetInstances,
          request_deserializer=definition__pb2.GetInstancesInput.FromString,
          response_serializer=definition__pb2.GetInstancesOutput.SerializeToString,
      ),
      'Stats': grpc.unary_unary_rpc_method_handler(
          servicer.Stats,
          request_deserializer=definition__pb2.StatsInput.FromString,
          response_serializer=definition__pb2.StatsOutput.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'hippoai.later.Later', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))