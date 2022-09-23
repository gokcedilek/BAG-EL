// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.13.0
// source: coord.proto

package coord

import (
	any "github.com/golang/protobuf/ptypes/any"
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

type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId  string   `protobuf:"bytes,1,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
	QueryType string   `protobuf:"bytes,2,opt,name=QueryType,proto3" json:"QueryType,omitempty"`
	Nodes     []uint64 `protobuf:"varint,3,rep,packed,name=Nodes,proto3" json:"Nodes,omitempty"`
	Graph     string   `protobuf:"bytes,4,opt,name=Graph,proto3" json:"Graph,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coord_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_coord_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_coord_proto_rawDescGZIP(), []int{0}
}

func (x *Query) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Query) GetQueryType() string {
	if x != nil {
		return x.QueryType
	}
	return ""
}

func (x *Query) GetNodes() []uint64 {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *Query) GetGraph() string {
	if x != nil {
		return x.Graph
	}
	return ""
}

type QueryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query  *Query   `protobuf:"bytes,1,opt,name=Query,proto3" json:"Query,omitempty"`
	Result *any.Any `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error  string   `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *QueryResult) Reset() {
	*x = QueryResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coord_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResult) ProtoMessage() {}

func (x *QueryResult) ProtoReflect() protoreflect.Message {
	mi := &file_coord_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResult.ProtoReflect.Descriptor instead.
func (*QueryResult) Descriptor() ([]byte, []int) {
	return file_coord_proto_rawDescGZIP(), []int{1}
}

func (x *QueryResult) GetQuery() *Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *QueryResult) GetResult() *any.Any {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *QueryResult) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_coord_proto protoreflect.FileDescriptor

var file_coord_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63,
	0x6f, 0x6f, 0x72, 0x64, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6d, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x04, 0x52, 0x05, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x47, 0x72, 0x61, 0x70, 0x68, 0x22, 0x75,
	0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x22, 0x0a,
	0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63,
	0x6f, 0x6f, 0x72, 0x64, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x2c, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x39, 0x0a, 0x05, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x12, 0x30,
	0x0a, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x0c, 0x2e, 0x63,
	0x6f, 0x6f, 0x72, 0x64, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x6f,
	0x72, 0x64, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_coord_proto_rawDescOnce sync.Once
	file_coord_proto_rawDescData = file_coord_proto_rawDesc
)

func file_coord_proto_rawDescGZIP() []byte {
	file_coord_proto_rawDescOnce.Do(func() {
		file_coord_proto_rawDescData = protoimpl.X.CompressGZIP(file_coord_proto_rawDescData)
	})
	return file_coord_proto_rawDescData
}

var file_coord_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_coord_proto_goTypes = []interface{}{
	(*Query)(nil),       // 0: coord.Query
	(*QueryResult)(nil), // 1: coord.QueryResult
	(*any.Any)(nil),     // 2: google.protobuf.Any
}
var file_coord_proto_depIdxs = []int32{
	0, // 0: coord.QueryResult.Query:type_name -> coord.Query
	2, // 1: coord.QueryResult.Result:type_name -> google.protobuf.Any
	0, // 2: coord.Coord.StartQuery:input_type -> coord.Query
	1, // 3: coord.Coord.StartQuery:output_type -> coord.QueryResult
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_coord_proto_init() }
func file_coord_proto_init() {
	if File_coord_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coord_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
		file_coord_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResult); i {
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
			RawDescriptor: file_coord_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coord_proto_goTypes,
		DependencyIndexes: file_coord_proto_depIdxs,
		MessageInfos:      file_coord_proto_msgTypes,
	}.Build()
	File_coord_proto = out.File
	file_coord_proto_rawDesc = nil
	file_coord_proto_goTypes = nil
	file_coord_proto_depIdxs = nil
}
