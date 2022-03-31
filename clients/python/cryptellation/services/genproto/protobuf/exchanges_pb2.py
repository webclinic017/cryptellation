# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protobuf/exchanges.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x18protobuf/exchanges.proto\x12\texchanges\"^\n\x08\x45xchange\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0f\n\x07periods\x18\x02 \x03(\t\x12\r\n\x05pairs\x18\x03 \x03(\t\x12\x0c\n\x04\x66\x65\x65s\x18\x04 \x01(\x02\x12\x16\n\x0elast_sync_time\x18\x05 \x01(\t\"%\n\x14ReadExchangesRequest\x12\r\n\x05names\x18\x01 \x03(\t\"?\n\x15ReadExchangesResponse\x12&\n\texchanges\x18\x01 \x03(\x0b\x32\x13.exchanges.Exchange2h\n\x10\x45xchangesService\x12T\n\rReadExchanges\x12\x1f.exchanges.ReadExchangesRequest\x1a .exchanges.ReadExchangesResponse\"\x00\x42\x19Z\x17/pkg/genproto/exchangesb\x06proto3')



_EXCHANGE = DESCRIPTOR.message_types_by_name['Exchange']
_READEXCHANGESREQUEST = DESCRIPTOR.message_types_by_name['ReadExchangesRequest']
_READEXCHANGESRESPONSE = DESCRIPTOR.message_types_by_name['ReadExchangesResponse']
Exchange = _reflection.GeneratedProtocolMessageType('Exchange', (_message.Message,), {
  'DESCRIPTOR' : _EXCHANGE,
  '__module__' : 'protobuf.exchanges_pb2'
  # @@protoc_insertion_point(class_scope:exchanges.Exchange)
  })
_sym_db.RegisterMessage(Exchange)

ReadExchangesRequest = _reflection.GeneratedProtocolMessageType('ReadExchangesRequest', (_message.Message,), {
  'DESCRIPTOR' : _READEXCHANGESREQUEST,
  '__module__' : 'protobuf.exchanges_pb2'
  # @@protoc_insertion_point(class_scope:exchanges.ReadExchangesRequest)
  })
_sym_db.RegisterMessage(ReadExchangesRequest)

ReadExchangesResponse = _reflection.GeneratedProtocolMessageType('ReadExchangesResponse', (_message.Message,), {
  'DESCRIPTOR' : _READEXCHANGESRESPONSE,
  '__module__' : 'protobuf.exchanges_pb2'
  # @@protoc_insertion_point(class_scope:exchanges.ReadExchangesResponse)
  })
_sym_db.RegisterMessage(ReadExchangesResponse)

_EXCHANGESSERVICE = DESCRIPTOR.services_by_name['ExchangesService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\027/pkg/genproto/exchanges'
  _EXCHANGE._serialized_start=39
  _EXCHANGE._serialized_end=133
  _READEXCHANGESREQUEST._serialized_start=135
  _READEXCHANGESREQUEST._serialized_end=172
  _READEXCHANGESRESPONSE._serialized_start=174
  _READEXCHANGESRESPONSE._serialized_end=237
  _EXCHANGESSERVICE._serialized_start=239
  _EXCHANGESSERVICE._serialized_end=343
# @@protoc_insertion_point(module_scope)
