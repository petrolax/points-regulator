// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.6.1
// source: server/proto/regulator/regulator.proto

package regulator

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

type PointsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Points []*Point `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
}

func (x *PointsRequest) Reset() {
	*x = PointsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_regulator_regulator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PointsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PointsRequest) ProtoMessage() {}

func (x *PointsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_regulator_regulator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PointsRequest.ProtoReflect.Descriptor instead.
func (*PointsRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_regulator_regulator_proto_rawDescGZIP(), []int{0}
}

func (x *PointsRequest) GetPoints() []*Point {
	if x != nil {
		return x.Points
	}
	return nil
}

type PointsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Points []*Point `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
}

func (x *PointsResponse) Reset() {
	*x = PointsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_regulator_regulator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PointsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PointsResponse) ProtoMessage() {}

func (x *PointsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_regulator_regulator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PointsResponse.ProtoReflect.Descriptor instead.
func (*PointsResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_regulator_regulator_proto_rawDescGZIP(), []int{1}
}

func (x *PointsResponse) GetPoints() []*Point {
	if x != nil {
		return x.Points
	}
	return nil
}

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int64 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int64 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_regulator_regulator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_regulator_regulator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_server_proto_regulator_regulator_proto_rawDescGZIP(), []int{2}
}

func (x *Point) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Point) GetY() int64 {
	if x != nil {
		return x.Y
	}
	return 0
}

var File_server_proto_regulator_regulator_proto protoreflect.FileDescriptor

var file_server_proto_regulator_regulator_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72,
	0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x0d, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x06, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x30, 0x0a, 0x0e, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x23, 0x0a, 0x05, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x79,
	0x32, 0x50, 0x0a, 0x0f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x3d, 0x0a, 0x18, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x74, 0x73, 0x4f, 0x66, 0x46, 0x6f, 0x75, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12,
	0x0e, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0f, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_proto_regulator_regulator_proto_rawDescOnce sync.Once
	file_server_proto_regulator_regulator_proto_rawDescData = file_server_proto_regulator_regulator_proto_rawDesc
)

func file_server_proto_regulator_regulator_proto_rawDescGZIP() []byte {
	file_server_proto_regulator_regulator_proto_rawDescOnce.Do(func() {
		file_server_proto_regulator_regulator_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_proto_regulator_regulator_proto_rawDescData)
	})
	return file_server_proto_regulator_regulator_proto_rawDescData
}

var file_server_proto_regulator_regulator_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_server_proto_regulator_regulator_proto_goTypes = []interface{}{
	(*PointsRequest)(nil),  // 0: PointsRequest
	(*PointsResponse)(nil), // 1: PointsResponse
	(*Point)(nil),          // 2: Point
}
var file_server_proto_regulator_regulator_proto_depIdxs = []int32{
	2, // 0: PointsRequest.points:type_name -> Point
	2, // 1: PointsResponse.points:type_name -> Point
	0, // 2: PointsRegulator.RegulateSetsOfFourPoints:input_type -> PointsRequest
	1, // 3: PointsRegulator.RegulateSetsOfFourPoints:output_type -> PointsResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_server_proto_regulator_regulator_proto_init() }
func file_server_proto_regulator_regulator_proto_init() {
	if File_server_proto_regulator_regulator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_proto_regulator_regulator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PointsRequest); i {
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
		file_server_proto_regulator_regulator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PointsResponse); i {
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
		file_server_proto_regulator_regulator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
			RawDescriptor: file_server_proto_regulator_regulator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_proto_regulator_regulator_proto_goTypes,
		DependencyIndexes: file_server_proto_regulator_regulator_proto_depIdxs,
		MessageInfos:      file_server_proto_regulator_regulator_proto_msgTypes,
	}.Build()
	File_server_proto_regulator_regulator_proto = out.File
	file_server_proto_regulator_regulator_proto_rawDesc = nil
	file_server_proto_regulator_regulator_proto_goTypes = nil
	file_server_proto_regulator_regulator_proto_depIdxs = nil
}
