// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.2
// source: message-service.proto

package notifications

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Later add SMS, Push, etc.
type NotificationType int32

const (
	NotificationType_EMAIL NotificationType = 0
)

// Enum value maps for NotificationType.
var (
	NotificationType_name = map[int32]string{
		0: "EMAIL",
	}
	NotificationType_value = map[string]int32{
		"EMAIL": 0,
	}
)

func (x NotificationType) Enum() *NotificationType {
	p := new(NotificationType)
	*p = x
	return p
}

func (x NotificationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationType) Descriptor() protoreflect.EnumDescriptor {
	return file_message_service_proto_enumTypes[0].Descriptor()
}

func (NotificationType) Type() protoreflect.EnumType {
	return &file_message_service_proto_enumTypes[0]
}

func (x NotificationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationType.Descriptor instead.
func (NotificationType) EnumDescriptor() ([]byte, []int) {
	return file_message_service_proto_rawDescGZIP(), []int{0}
}

type MessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    NotificationType `protobuf:"varint,1,opt,name=type,proto3,enum=NotificationType" json:"type,omitempty"`
	To      string           `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Msg     string           `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	Subject string           `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *MessageRequest) Reset() {
	*x = MessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageRequest) ProtoMessage() {}

func (x *MessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageRequest.ProtoReflect.Descriptor instead.
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return file_message_service_proto_rawDescGZIP(), []int{0}
}

func (x *MessageRequest) GetType() NotificationType {
	if x != nil {
		return x.Type
	}
	return NotificationType_EMAIL
}

func (x *MessageRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *MessageRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *MessageRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

type MessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *MessageResponse) Reset() {
	*x = MessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageResponse) ProtoMessage() {}

func (x *MessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageResponse.ProtoReflect.Descriptor instead.
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return file_message_service_proto_rawDescGZIP(), []int{1}
}

func (x *MessageResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_message_service_proto protoreflect.FileDescriptor

var file_message_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2b, 0x0a, 0x0f, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2a, 0x1d, 0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d,
	0x41, 0x49, 0x4c, 0x10, 0x00, 0x32, 0xb2, 0x01, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a,
	0x0a, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x0f, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a,
	0x0a, 0x0f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_service_proto_rawDescOnce sync.Once
	file_message_service_proto_rawDescData = file_message_service_proto_rawDesc
)

func file_message_service_proto_rawDescGZIP() []byte {
	file_message_service_proto_rawDescOnce.Do(func() {
		file_message_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_service_proto_rawDescData)
	})
	return file_message_service_proto_rawDescData
}

var file_message_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_message_service_proto_goTypes = []interface{}{
	(NotificationType)(0),   // 0: NotificationType
	(*MessageRequest)(nil),  // 1: MessageRequest
	(*MessageResponse)(nil), // 2: MessageResponse
	(*empty.Empty)(nil),     // 3: google.protobuf.Empty
}
var file_message_service_proto_depIdxs = []int32{
	0, // 0: MessageRequest.type:type_name -> NotificationType
	1, // 1: Notification.SendNotification:input_type -> MessageRequest
	1, // 2: Notification.AddToQueue:input_type -> MessageRequest
	3, // 3: Notification.RemoveFromQueue:input_type -> google.protobuf.Empty
	2, // 4: Notification.SendNotification:output_type -> MessageResponse
	2, // 5: Notification.AddToQueue:output_type -> MessageResponse
	1, // 6: Notification.RemoveFromQueue:output_type -> MessageRequest
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_message_service_proto_init() }
func file_message_service_proto_init() {
	if File_message_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_message_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_service_proto_goTypes,
		DependencyIndexes: file_message_service_proto_depIdxs,
		EnumInfos:         file_message_service_proto_enumTypes,
		MessageInfos:      file_message_service_proto_msgTypes,
	}.Build()
	File_message_service_proto = out.File
	file_message_service_proto_rawDesc = nil
	file_message_service_proto_goTypes = nil
	file_message_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NotificationClient is the client API for Notification service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationClient interface {
	// rpc AddToQueue(MessageRequest) returns (MessageResponse);
	SendNotification(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error)
	AddToQueue(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error)
	RemoveFromQueue(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*MessageRequest, error)
}

type notificationClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationClient(cc grpc.ClientConnInterface) NotificationClient {
	return &notificationClient{cc}
}

func (c *notificationClient) SendNotification(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/Notification/SendNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) AddToQueue(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/Notification/AddToQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) RemoveFromQueue(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*MessageRequest, error) {
	out := new(MessageRequest)
	err := c.cc.Invoke(ctx, "/Notification/RemoveFromQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServer is the server API for Notification service.
type NotificationServer interface {
	// rpc AddToQueue(MessageRequest) returns (MessageResponse);
	SendNotification(context.Context, *MessageRequest) (*MessageResponse, error)
	AddToQueue(context.Context, *MessageRequest) (*MessageResponse, error)
	RemoveFromQueue(context.Context, *empty.Empty) (*MessageRequest, error)
}

// UnimplementedNotificationServer can be embedded to have forward compatible implementations.
type UnimplementedNotificationServer struct {
}

func (*UnimplementedNotificationServer) SendNotification(context.Context, *MessageRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNotification not implemented")
}
func (*UnimplementedNotificationServer) AddToQueue(context.Context, *MessageRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToQueue not implemented")
}
func (*UnimplementedNotificationServer) RemoveFromQueue(context.Context, *empty.Empty) (*MessageRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromQueue not implemented")
}

func RegisterNotificationServer(s *grpc.Server, srv NotificationServer) {
	s.RegisterService(&_Notification_serviceDesc, srv)
}

func _Notification_SendNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).SendNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Notification/SendNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).SendNotification(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_AddToQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).AddToQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Notification/AddToQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).AddToQueue(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_RemoveFromQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).RemoveFromQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Notification/RemoveFromQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).RemoveFromQueue(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notification_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Notification",
	HandlerType: (*NotificationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendNotification",
			Handler:    _Notification_SendNotification_Handler,
		},
		{
			MethodName: "AddToQueue",
			Handler:    _Notification_AddToQueue_Handler,
		},
		{
			MethodName: "RemoveFromQueue",
			Handler:    _Notification_RemoveFromQueue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message-service.proto",
}
