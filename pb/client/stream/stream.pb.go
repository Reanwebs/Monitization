// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: client/stream/stream.proto

package stream

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_stream_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_client_stream_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_client_stream_stream_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_stream_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_client_stream_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_client_stream_stream_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type VideoDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoID string `protobuf:"bytes,1,opt,name=VideoID,proto3" json:"VideoID,omitempty"`
}

func (x *VideoDetailsRequest) Reset() {
	*x = VideoDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_stream_stream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoDetailsRequest) ProtoMessage() {}

func (x *VideoDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_stream_stream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoDetailsRequest.ProtoReflect.Descriptor instead.
func (*VideoDetailsRequest) Descriptor() ([]byte, []int) {
	return file_client_stream_stream_proto_rawDescGZIP(), []int{2}
}

func (x *VideoDetailsRequest) GetVideoID() string {
	if x != nil {
		return x.VideoID
	}
	return ""
}

type VideoDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerID string `protobuf:"bytes,1,opt,name=OwnerID,proto3" json:"OwnerID,omitempty"`
	Coins   int32  `protobuf:"varint,2,opt,name=Coins,proto3" json:"Coins,omitempty"`
	Title   string `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
}

func (x *VideoDetailsResponse) Reset() {
	*x = VideoDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_stream_stream_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoDetailsResponse) ProtoMessage() {}

func (x *VideoDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_stream_stream_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoDetailsResponse.ProtoReflect.Descriptor instead.
func (*VideoDetailsResponse) Descriptor() ([]byte, []int) {
	return file_client_stream_stream_proto_rawDescGZIP(), []int{3}
}

func (x *VideoDetailsResponse) GetOwnerID() string {
	if x != nil {
		return x.OwnerID
	}
	return ""
}

func (x *VideoDetailsResponse) GetCoins() int32 {
	if x != nil {
		return x.Coins
	}
	return 0
}

func (x *VideoDetailsResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_client_stream_stream_proto protoreflect.FileDescriptor

var file_client_stream_stream_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x22, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x2f, 0x0a, 0x13, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x49, 0x44, 0x22, 0x5c, 0x0a, 0x14, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x32, 0x7d, 0x0a, 0x0c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x28, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0c,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x17, 0x2e, 0x70,
	0x62, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x12, 0x5a, 0x10, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_stream_stream_proto_rawDescOnce sync.Once
	file_client_stream_stream_proto_rawDescData = file_client_stream_stream_proto_rawDesc
)

func file_client_stream_stream_proto_rawDescGZIP() []byte {
	file_client_stream_stream_proto_rawDescOnce.Do(func() {
		file_client_stream_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_stream_stream_proto_rawDescData)
	})
	return file_client_stream_stream_proto_rawDescData
}

var file_client_stream_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_client_stream_stream_proto_goTypes = []interface{}{
	(*Request)(nil),              // 0: pb.Request
	(*Response)(nil),             // 1: pb.Response
	(*VideoDetailsRequest)(nil),  // 2: pb.VideoDetailsRequest
	(*VideoDetailsResponse)(nil), // 3: pb.VideoDetailsResponse
}
var file_client_stream_stream_proto_depIdxs = []int32{
	0, // 0: pb.VideoService.HealthCheck:input_type -> pb.Request
	2, // 1: pb.VideoService.VideoDetails:input_type -> pb.VideoDetailsRequest
	1, // 2: pb.VideoService.HealthCheck:output_type -> pb.Response
	3, // 3: pb.VideoService.VideoDetails:output_type -> pb.VideoDetailsResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_client_stream_stream_proto_init() }
func file_client_stream_stream_proto_init() {
	if File_client_stream_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_stream_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_client_stream_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_client_stream_stream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoDetailsRequest); i {
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
		file_client_stream_stream_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoDetailsResponse); i {
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
			RawDescriptor: file_client_stream_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_stream_stream_proto_goTypes,
		DependencyIndexes: file_client_stream_stream_proto_depIdxs,
		MessageInfos:      file_client_stream_stream_proto_msgTypes,
	}.Build()
	File_client_stream_stream_proto = out.File
	file_client_stream_stream_proto_rawDesc = nil
	file_client_stream_stream_proto_goTypes = nil
	file_client_stream_stream_proto_depIdxs = nil
}
