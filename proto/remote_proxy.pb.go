// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: proto/remote_proxy.proto

package RemoteProxy

import (
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

//请求参数
type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Class    string `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`       //指定类
	Func     string `protobuf:"bytes,2,opt,name=func,proto3" json:"func,omitempty"`         //指定方法
	Args     string `protobuf:"bytes,3,opt,name=args,proto3" json:"args,omitempty"`         //参数
	Path     string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`         //执行位置
	Hostname string `protobuf:"bytes,5,opt,name=hostname,proto3" json:"hostname,omitempty"` //消费者服务端口
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_remote_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_proto_remote_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_proto_remote_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *Params) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *Params) GetFunc() string {
	if x != nil {
		return x.Func
	}
	return ""
}

func (x *Params) GetArgs() string {
	if x != nil {
		return x.Args
	}
	return ""
}

func (x *Params) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Params) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Rid     int64  `protobuf:"varint,3,opt,name=rid,proto3" json:"rid,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_remote_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_remote_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_proto_remote_proxy_proto_rawDescGZIP(), []int{1}
}

func (x *Reply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Reply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Reply) GetRid() int64 {
	if x != nil {
		return x.Rid
	}
	return 0
}

var File_proto_remote_proxy_proto protoreflect.FileDescriptor

var file_proto_remote_proxy_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x52, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x22, 0x76, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x75, 0x6e, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x75, 0x6e, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x72, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x4d, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x72, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x72, 0x69, 0x64, 0x32, 0x3d,
	0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x04, 0x77, 0x6f,
	0x72, 0x6b, 0x12, 0x13, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x12, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x32, 0x3d, 0x0a,
	0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x04, 0x70, 0x75, 0x73,
	0x68, 0x12, 0x13, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x12, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e,
	0x2e, 0x2f, 0x3b, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_remote_proxy_proto_rawDescOnce sync.Once
	file_proto_remote_proxy_proto_rawDescData = file_proto_remote_proxy_proto_rawDesc
)

func file_proto_remote_proxy_proto_rawDescGZIP() []byte {
	file_proto_remote_proxy_proto_rawDescOnce.Do(func() {
		file_proto_remote_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_remote_proxy_proto_rawDescData)
	})
	return file_proto_remote_proxy_proto_rawDescData
}

var file_proto_remote_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_remote_proxy_proto_goTypes = []interface{}{
	(*Params)(nil), // 0: RemoteProxy.Params
	(*Reply)(nil),  // 1: RemoteProxy.Reply
}
var file_proto_remote_proxy_proto_depIdxs = []int32{
	0, // 0: RemoteProxy.Consumer.work:input_type -> RemoteProxy.Params
	0, // 1: RemoteProxy.Producer.push:input_type -> RemoteProxy.Params
	1, // 2: RemoteProxy.Consumer.work:output_type -> RemoteProxy.Reply
	1, // 3: RemoteProxy.Producer.push:output_type -> RemoteProxy.Reply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_remote_proxy_proto_init() }
func file_proto_remote_proxy_proto_init() {
	if File_proto_remote_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_remote_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
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
		file_proto_remote_proxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
			RawDescriptor: file_proto_remote_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_remote_proxy_proto_goTypes,
		DependencyIndexes: file_proto_remote_proxy_proto_depIdxs,
		MessageInfos:      file_proto_remote_proxy_proto_msgTypes,
	}.Build()
	File_proto_remote_proxy_proto = out.File
	file_proto_remote_proxy_proto_rawDesc = nil
	file_proto_remote_proxy_proto_goTypes = nil
	file_proto_remote_proxy_proto_depIdxs = nil
}
