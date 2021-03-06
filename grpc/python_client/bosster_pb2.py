# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: bosster.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='bosster.proto',
  package='bosster',
  syntax='proto3',
  serialized_pb=_b('\n\rbosster.proto\x12\x07\x62osster\"\x81\x01\n\x07\x43ontact\x12\x0c\n\x04name\x18\x01 \x01(\t\x12#\n\x04type\x18\x02 \x01(\x0e\x32\x15.bosster.Contact.Type\x12\r\n\x05value\x18\x03 \x01(\t\"4\n\x04Type\x12\r\n\tUNIVERSAL\x10\x00\x12\t\n\x05PHONE\x10\x01\x12\x07\n\x03URL\x10\x02\x12\t\n\x05\x45MAIL\x10\x03\"+\n\x04Post\x12\x0f\n\x07message\x18\x01 \x01(\t\x12\x12\n\nimage_urls\x18\x02 \x03(\t\"\x7f\n\x0bPostRequest\x12\x1b\n\x04post\x18\x01 \x01(\x0b\x32\r.bosster.Post\x12\x0c\n\x04sync\x18\x02 \x01(\x08\x12\"\n\x08\x63ontacts\x18\x04 \x03(\x0b\x32\x10.bosster.Contact\x12!\n\x07targets\x18\x05 \x03(\x0b\x32\x10.bosster.PostJob\"\x95\x02\n\x07PostJob\x12\x13\n\x0b\x65xternal_id\x18\x01 \x01(\t\x12!\n\x04type\x18\x02 \x01(\x0e\x32\x13.bosster.SocialType\x12\x11\n\tsocial_id\x18\x03 \x01(\t\x12\x18\n\x10social_wall_type\x18\x0b \x01(\t\x12\x14\n\x0csocial_token\x18\x04 \x01(\t\x12\x14\n\x0csocial_login\x18\x05 \x01(\t\x12\x17\n\x0fsocial_password\x18\x06 \x01(\t\x12\x15\n\rsocial_app_id\x18\x07 \x01(\t\x12\x19\n\x11social_app_secret\x18\x08 \x01(\t\x12\x1f\n\x06status\x18\t \x01(\x0e\x32\x0f.bosster.STATUS\x12\r\n\x05\x65rror\x18\n \x01(\t\"+\n\tPostReply\x12\x1e\n\x04jobs\x18\x01 \x03(\x0b\x32\x10.bosster.PostJob*V\n\nSocialType\x12\x08\n\x04NONE\x10\x00\x12\x06\n\x02VK\x10\x01\x12\x0c\n\x08TELEGRAM\x10\x02\x12\r\n\tINSTAGRAM\x10\x03\x12\x0b\n\x07TWITTER\x10\x04\x12\x0c\n\x08\x46\x41\x43\x45\x42OOK\x10\x05*<\n\x06STATUS\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0c\n\x08\x45NQUEUED\x10\x01\x12\x08\n\x04\x46\x41IL\x10\x02\x12\r\n\tCOMPLETED\x10\x03\x32<\n\x06Poster\x12\x32\n\x04Post\x12\x14.bosster.PostRequest\x1a\x12.bosster.PostReply\"\x00\x62\x06proto3')
)

_SOCIALTYPE = _descriptor.EnumDescriptor(
  name='SocialType',
  full_name='bosster.SocialType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NONE', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='VK', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TELEGRAM', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='INSTAGRAM', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TWITTER', index=4, number=4,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FACEBOOK', index=5, number=5,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=657,
  serialized_end=743,
)
_sym_db.RegisterEnumDescriptor(_SOCIALTYPE)

SocialType = enum_type_wrapper.EnumTypeWrapper(_SOCIALTYPE)
_STATUS = _descriptor.EnumDescriptor(
  name='STATUS',
  full_name='bosster.STATUS',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ENQUEUED', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FAIL', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='COMPLETED', index=3, number=3,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=745,
  serialized_end=805,
)
_sym_db.RegisterEnumDescriptor(_STATUS)

STATUS = enum_type_wrapper.EnumTypeWrapper(_STATUS)
NONE = 0
VK = 1
TELEGRAM = 2
INSTAGRAM = 3
TWITTER = 4
FACEBOOK = 5
UNKNOWN = 0
ENQUEUED = 1
FAIL = 2
COMPLETED = 3


_CONTACT_TYPE = _descriptor.EnumDescriptor(
  name='Type',
  full_name='bosster.Contact.Type',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNIVERSAL', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PHONE', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='URL', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='EMAIL', index=3, number=3,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=104,
  serialized_end=156,
)
_sym_db.RegisterEnumDescriptor(_CONTACT_TYPE)


_CONTACT = _descriptor.Descriptor(
  name='Contact',
  full_name='bosster.Contact',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='bosster.Contact.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='type', full_name='bosster.Contact.type', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='value', full_name='bosster.Contact.value', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _CONTACT_TYPE,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=27,
  serialized_end=156,
)


_POST = _descriptor.Descriptor(
  name='Post',
  full_name='bosster.Post',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='message', full_name='bosster.Post.message', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='image_urls', full_name='bosster.Post.image_urls', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=158,
  serialized_end=201,
)


_POSTREQUEST = _descriptor.Descriptor(
  name='PostRequest',
  full_name='bosster.PostRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='post', full_name='bosster.PostRequest.post', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='sync', full_name='bosster.PostRequest.sync', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='contacts', full_name='bosster.PostRequest.contacts', index=2,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='targets', full_name='bosster.PostRequest.targets', index=3,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=203,
  serialized_end=330,
)


_POSTJOB = _descriptor.Descriptor(
  name='PostJob',
  full_name='bosster.PostJob',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='external_id', full_name='bosster.PostJob.external_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='type', full_name='bosster.PostJob.type', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='social_id', full_name='bosster.PostJob.social_id', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='social_wall_type', full_name='bosster.PostJob.social_wall_type', index=3,
      number=11, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='social_token', full_name='bosster.PostJob.social_token', index=4,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='social_login', full_name='bosster.PostJob.social_login', index=5,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='social_password', full_name='bosster.PostJob.social_password', index=6,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='social_app_id', full_name='bosster.PostJob.social_app_id', index=7,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='social_app_secret', full_name='bosster.PostJob.social_app_secret', index=8,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='status', full_name='bosster.PostJob.status', index=9,
      number=9, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='error', full_name='bosster.PostJob.error', index=10,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=333,
  serialized_end=610,
)


_POSTREPLY = _descriptor.Descriptor(
  name='PostReply',
  full_name='bosster.PostReply',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='jobs', full_name='bosster.PostReply.jobs', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=612,
  serialized_end=655,
)

_CONTACT.fields_by_name['type'].enum_type = _CONTACT_TYPE
_CONTACT_TYPE.containing_type = _CONTACT
_POSTREQUEST.fields_by_name['post'].message_type = _POST
_POSTREQUEST.fields_by_name['contacts'].message_type = _CONTACT
_POSTREQUEST.fields_by_name['targets'].message_type = _POSTJOB
_POSTJOB.fields_by_name['type'].enum_type = _SOCIALTYPE
_POSTJOB.fields_by_name['status'].enum_type = _STATUS
_POSTREPLY.fields_by_name['jobs'].message_type = _POSTJOB
DESCRIPTOR.message_types_by_name['Contact'] = _CONTACT
DESCRIPTOR.message_types_by_name['Post'] = _POST
DESCRIPTOR.message_types_by_name['PostRequest'] = _POSTREQUEST
DESCRIPTOR.message_types_by_name['PostJob'] = _POSTJOB
DESCRIPTOR.message_types_by_name['PostReply'] = _POSTREPLY
DESCRIPTOR.enum_types_by_name['SocialType'] = _SOCIALTYPE
DESCRIPTOR.enum_types_by_name['STATUS'] = _STATUS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Contact = _reflection.GeneratedProtocolMessageType('Contact', (_message.Message,), dict(
  DESCRIPTOR = _CONTACT,
  __module__ = 'bosster_pb2'
  # @@protoc_insertion_point(class_scope:bosster.Contact)
  ))
_sym_db.RegisterMessage(Contact)

Post = _reflection.GeneratedProtocolMessageType('Post', (_message.Message,), dict(
  DESCRIPTOR = _POST,
  __module__ = 'bosster_pb2'
  # @@protoc_insertion_point(class_scope:bosster.Post)
  ))
_sym_db.RegisterMessage(Post)

PostRequest = _reflection.GeneratedProtocolMessageType('PostRequest', (_message.Message,), dict(
  DESCRIPTOR = _POSTREQUEST,
  __module__ = 'bosster_pb2'
  # @@protoc_insertion_point(class_scope:bosster.PostRequest)
  ))
_sym_db.RegisterMessage(PostRequest)

PostJob = _reflection.GeneratedProtocolMessageType('PostJob', (_message.Message,), dict(
  DESCRIPTOR = _POSTJOB,
  __module__ = 'bosster_pb2'
  # @@protoc_insertion_point(class_scope:bosster.PostJob)
  ))
_sym_db.RegisterMessage(PostJob)

PostReply = _reflection.GeneratedProtocolMessageType('PostReply', (_message.Message,), dict(
  DESCRIPTOR = _POSTREPLY,
  __module__ = 'bosster_pb2'
  # @@protoc_insertion_point(class_scope:bosster.PostReply)
  ))
_sym_db.RegisterMessage(PostReply)



_POSTER = _descriptor.ServiceDescriptor(
  name='Poster',
  full_name='bosster.Poster',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=807,
  serialized_end=867,
  methods=[
  _descriptor.MethodDescriptor(
    name='Post',
    full_name='bosster.Poster.Post',
    index=0,
    containing_service=None,
    input_type=_POSTREQUEST,
    output_type=_POSTREPLY,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_POSTER)

DESCRIPTOR.services_by_name['Poster'] = _POSTER

try:
  # THESE ELEMENTS WILL BE DEPRECATED.
  # Please use the generated *_pb2_grpc.py files instead.
  import grpc
  from grpc.beta import implementations as beta_implementations
  from grpc.beta import interfaces as beta_interfaces
  from grpc.framework.common import cardinality
  from grpc.framework.interfaces.face import utilities as face_utilities


  class PosterStub(object):
    """The greeting service definition.
    """

    def __init__(self, channel):
      """Constructor.

      Args:
        channel: A grpc.Channel.
      """
      self.Post = channel.unary_unary(
          '/bosster.Poster/Post',
          request_serializer=PostRequest.SerializeToString,
          response_deserializer=PostReply.FromString,
          )


  class PosterServicer(object):
    """The greeting service definition.
    """

    def Post(self, request, context):
      """Sends a greeting
      """
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')


  def add_PosterServicer_to_server(servicer, server):
    rpc_method_handlers = {
        'Post': grpc.unary_unary_rpc_method_handler(
            servicer.Post,
            request_deserializer=PostRequest.FromString,
            response_serializer=PostReply.SerializeToString,
        ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
        'bosster.Poster', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


  class BetaPosterServicer(object):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This class was generated
    only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0."""
    """The greeting service definition.
    """
    def Post(self, request, context):
      """Sends a greeting
      """
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)


  class BetaPosterStub(object):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This class was generated
    only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0."""
    """The greeting service definition.
    """
    def Post(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      """Sends a greeting
      """
      raise NotImplementedError()
    Post.future = None


  def beta_create_Poster_server(servicer, pool=None, pool_size=None, default_timeout=None, maximum_timeout=None):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This function was
    generated only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0"""
    request_deserializers = {
      ('bosster.Poster', 'Post'): PostRequest.FromString,
    }
    response_serializers = {
      ('bosster.Poster', 'Post'): PostReply.SerializeToString,
    }
    method_implementations = {
      ('bosster.Poster', 'Post'): face_utilities.unary_unary_inline(servicer.Post),
    }
    server_options = beta_implementations.server_options(request_deserializers=request_deserializers, response_serializers=response_serializers, thread_pool=pool, thread_pool_size=pool_size, default_timeout=default_timeout, maximum_timeout=maximum_timeout)
    return beta_implementations.server(method_implementations, options=server_options)


  def beta_create_Poster_stub(channel, host=None, metadata_transformer=None, pool=None, pool_size=None):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This function was
    generated only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0"""
    request_serializers = {
      ('bosster.Poster', 'Post'): PostRequest.SerializeToString,
    }
    response_deserializers = {
      ('bosster.Poster', 'Post'): PostReply.FromString,
    }
    cardinalities = {
      'Post': cardinality.Cardinality.UNARY_UNARY,
    }
    stub_options = beta_implementations.stub_options(host=host, metadata_transformer=metadata_transformer, request_serializers=request_serializers, response_deserializers=response_deserializers, thread_pool=pool, thread_pool_size=pool_size)
    return beta_implementations.dynamic_stub(channel, 'bosster.Poster', cardinalities, options=stub_options)
except ImportError:
  pass
# @@protoc_insertion_point(module_scope)
