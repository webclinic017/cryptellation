# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: candlesticks.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x12\x63\x61ndlesticks.proto\x12\x0c\x63\x61ndlesticks\"c\n\x0b\x43\x61ndlestick\x12\x0c\n\x04time\x18\x01 \x01(\t\x12\x0c\n\x04open\x18\x02 \x01(\x01\x12\x0c\n\x04high\x18\x03 \x01(\x01\x12\x0b\n\x03low\x18\x04 \x01(\x01\x12\r\n\x05\x63lose\x18\x05 \x01(\x01\x12\x0e\n\x06volume\x18\x06 \x01(\x01\"\x87\x01\n\x17ReadCandlesticksRequest\x12\x15\n\rexchange_name\x18\x01 \x01(\t\x12\x13\n\x0bpair_symbol\x18\x02 \x01(\t\x12\x15\n\rperiod_symbol\x18\x03 \x01(\t\x12\r\n\x05start\x18\x04 \x01(\t\x12\x0b\n\x03\x65nd\x18\x05 \x01(\t\x12\r\n\x05limit\x18\x06 \x01(\x03\"K\n\x18ReadCandlesticksResponse\x12/\n\x0c\x63\x61ndlesticks\x18\x01 \x03(\x0b\x32\x19.candlesticks.Candlestick2z\n\x13\x43\x61ndlesticksService\x12\x63\n\x10ReadCandlesticks\x12%.candlesticks.ReadCandlesticksRequest\x1a&.candlesticks.ReadCandlesticksResponse\"\x00\x42\'Z%/services/candlesticks/pkg/grpc/protob\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'candlesticks_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z%/services/candlesticks/pkg/grpc/proto'
  _CANDLESTICK._serialized_start=36
  _CANDLESTICK._serialized_end=135
  _READCANDLESTICKSREQUEST._serialized_start=138
  _READCANDLESTICKSREQUEST._serialized_end=273
  _READCANDLESTICKSRESPONSE._serialized_start=275
  _READCANDLESTICKSRESPONSE._serialized_end=350
  _CANDLESTICKSSERVICE._serialized_start=352
  _CANDLESTICKSSERVICE._serialized_end=474
# @@protoc_insertion_point(module_scope)
