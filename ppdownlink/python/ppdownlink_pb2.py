# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ppdownlink.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='ppdownlink.proto',
  package='ppdownlink',
  syntax='proto3',
  serialized_options=_b('Z)github.com/sukhajata/ppmessage/ppdownlink'),
  serialized_pb=_b('\n\x10ppdownlink.proto\x12\nppdownlink\"|\n\x15\x43onfigDownlinkMessage\x12\x11\n\tdeviceeui\x18\x01 \x01(\t\x12\x0c\n\x04slot\x18\x02 \x01(\r\x12\r\n\x05index\x18\x03 \x01(\r\x12\x10\n\x08\x66irmware\x18\x04 \x01(\t\x12\r\n\x05value\x18\x05 \x01(\x0c\x12\x12\n\nnumretries\x18\x06 \x01(\r\"\x9b\x01\n\x17\x46unctionDownlinkMessage\x12\x11\n\tdeviceeui\x18\x01 \x01(\t\x12\x0c\n\x04slot\x18\x02 \x01(\r\x12\r\n\x05index\x18\x03 \x01(\r\x12\x10\n\x08\x66irmware\x18\x04 \x01(\t\x12\x0e\n\x06param1\x18\x05 \x01(\x0c\x12\x0e\n\x06param2\x18\x06 \x01(\x0c\x12\x0e\n\x06param3\x18\x07 \x01(\x0c\x12\x0e\n\x06param4\x18\x08 \x01(\x0c\"\x81\x01\n\x14ResendRequestMessage\x12\x11\n\tdeviceeui\x18\x01 \x01(\t\x12\x10\n\x08timesent\x18\x02 \x01(\x04\x12\x11\n\tmessageid\x18\x03 \x01(\r\x12\x13\n\x0bmessagetype\x18\x04 \x01(\r\x12\r\n\x05spare\x18\x05 \x01(\x05\x12\r\n\x05state\x18\x06 \x01(\x05\"Z\n\x13RelayCommandMessage\x12\x11\n\tdeviceeui\x18\x01 \x01(\t\x12\x10\n\x08timesent\x18\x02 \x01(\x04\x12\r\n\x05index\x18\x03 \x01(\r\x12\x0f\n\x07\x63ommand\x18\x04 \x01(\rB+Z)github.com/sukhajata/ppmessage/ppdownlinkb\x06proto3')
)




_CONFIGDOWNLINKMESSAGE = _descriptor.Descriptor(
  name='ConfigDownlinkMessage',
  full_name='ppdownlink.ConfigDownlinkMessage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='deviceeui', full_name='ppdownlink.ConfigDownlinkMessage.deviceeui', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='slot', full_name='ppdownlink.ConfigDownlinkMessage.slot', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='index', full_name='ppdownlink.ConfigDownlinkMessage.index', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='firmware', full_name='ppdownlink.ConfigDownlinkMessage.firmware', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='ppdownlink.ConfigDownlinkMessage.value', index=4,
      number=5, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='numretries', full_name='ppdownlink.ConfigDownlinkMessage.numretries', index=5,
      number=6, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=32,
  serialized_end=156,
)


_FUNCTIONDOWNLINKMESSAGE = _descriptor.Descriptor(
  name='FunctionDownlinkMessage',
  full_name='ppdownlink.FunctionDownlinkMessage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='deviceeui', full_name='ppdownlink.FunctionDownlinkMessage.deviceeui', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='slot', full_name='ppdownlink.FunctionDownlinkMessage.slot', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='index', full_name='ppdownlink.FunctionDownlinkMessage.index', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='firmware', full_name='ppdownlink.FunctionDownlinkMessage.firmware', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='param1', full_name='ppdownlink.FunctionDownlinkMessage.param1', index=4,
      number=5, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='param2', full_name='ppdownlink.FunctionDownlinkMessage.param2', index=5,
      number=6, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='param3', full_name='ppdownlink.FunctionDownlinkMessage.param3', index=6,
      number=7, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='param4', full_name='ppdownlink.FunctionDownlinkMessage.param4', index=7,
      number=8, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=159,
  serialized_end=314,
)


_RESENDREQUESTMESSAGE = _descriptor.Descriptor(
  name='ResendRequestMessage',
  full_name='ppdownlink.ResendRequestMessage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='deviceeui', full_name='ppdownlink.ResendRequestMessage.deviceeui', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='timesent', full_name='ppdownlink.ResendRequestMessage.timesent', index=1,
      number=2, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='messageid', full_name='ppdownlink.ResendRequestMessage.messageid', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='messagetype', full_name='ppdownlink.ResendRequestMessage.messagetype', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='spare', full_name='ppdownlink.ResendRequestMessage.spare', index=4,
      number=5, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='state', full_name='ppdownlink.ResendRequestMessage.state', index=5,
      number=6, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=317,
  serialized_end=446,
)


_RELAYCOMMANDMESSAGE = _descriptor.Descriptor(
  name='RelayCommandMessage',
  full_name='ppdownlink.RelayCommandMessage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='deviceeui', full_name='ppdownlink.RelayCommandMessage.deviceeui', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='timesent', full_name='ppdownlink.RelayCommandMessage.timesent', index=1,
      number=2, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='index', full_name='ppdownlink.RelayCommandMessage.index', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='command', full_name='ppdownlink.RelayCommandMessage.command', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=448,
  serialized_end=538,
)

DESCRIPTOR.message_types_by_name['ConfigDownlinkMessage'] = _CONFIGDOWNLINKMESSAGE
DESCRIPTOR.message_types_by_name['FunctionDownlinkMessage'] = _FUNCTIONDOWNLINKMESSAGE
DESCRIPTOR.message_types_by_name['ResendRequestMessage'] = _RESENDREQUESTMESSAGE
DESCRIPTOR.message_types_by_name['RelayCommandMessage'] = _RELAYCOMMANDMESSAGE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ConfigDownlinkMessage = _reflection.GeneratedProtocolMessageType('ConfigDownlinkMessage', (_message.Message,), {
  'DESCRIPTOR' : _CONFIGDOWNLINKMESSAGE,
  '__module__' : 'ppdownlink_pb2'
  # @@protoc_insertion_point(class_scope:ppdownlink.ConfigDownlinkMessage)
  })
_sym_db.RegisterMessage(ConfigDownlinkMessage)

FunctionDownlinkMessage = _reflection.GeneratedProtocolMessageType('FunctionDownlinkMessage', (_message.Message,), {
  'DESCRIPTOR' : _FUNCTIONDOWNLINKMESSAGE,
  '__module__' : 'ppdownlink_pb2'
  # @@protoc_insertion_point(class_scope:ppdownlink.FunctionDownlinkMessage)
  })
_sym_db.RegisterMessage(FunctionDownlinkMessage)

ResendRequestMessage = _reflection.GeneratedProtocolMessageType('ResendRequestMessage', (_message.Message,), {
  'DESCRIPTOR' : _RESENDREQUESTMESSAGE,
  '__module__' : 'ppdownlink_pb2'
  # @@protoc_insertion_point(class_scope:ppdownlink.ResendRequestMessage)
  })
_sym_db.RegisterMessage(ResendRequestMessage)

RelayCommandMessage = _reflection.GeneratedProtocolMessageType('RelayCommandMessage', (_message.Message,), {
  'DESCRIPTOR' : _RELAYCOMMANDMESSAGE,
  '__module__' : 'ppdownlink_pb2'
  # @@protoc_insertion_point(class_scope:ppdownlink.RelayCommandMessage)
  })
_sym_db.RegisterMessage(RelayCommandMessage)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
