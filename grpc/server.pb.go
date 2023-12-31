// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: grpc/server.proto

package proto

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

type ErrorCode int32

const (
	ErrorCode_FAIL      ErrorCode = 0
	ErrorCode_SUCCESS   ErrorCode = 1
	ErrorCode_EXCEPTION ErrorCode = 2
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0: "FAIL",
		1: "SUCCESS",
		2: "EXCEPTION",
	}
	ErrorCode_value = map[string]int32{
		"FAIL":      0,
		"SUCCESS":   1,
		"EXCEPTION": 2,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_server_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_grpc_server_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_grpc_server_proto_rawDescGZIP(), []int{0}
}

type OutcomeType int32

const (
	OutcomeType_RESULT      OutcomeType = 0
	OutcomeType_HIGHEST_BID OutcomeType = 1
)

// Enum value maps for OutcomeType.
var (
	OutcomeType_name = map[int32]string{
		0: "RESULT",
		1: "HIGHEST_BID",
	}
	OutcomeType_value = map[string]int32{
		"RESULT":      0,
		"HIGHEST_BID": 1,
	}
)

func (x OutcomeType) Enum() *OutcomeType {
	p := new(OutcomeType)
	*p = x
	return p
}

func (x OutcomeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OutcomeType) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_server_proto_enumTypes[1].Descriptor()
}

func (OutcomeType) Type() protoreflect.EnumType {
	return &file_grpc_server_proto_enumTypes[1]
}

func (x OutcomeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OutcomeType.Descriptor instead.
func (OutcomeType) EnumDescriptor() ([]byte, []int) {
	return file_grpc_server_proto_rawDescGZIP(), []int{1}
}

// Types for bid
type Amount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount int64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Amount) Reset() {
	*x = Amount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Amount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Amount) ProtoMessage() {}

func (x *Amount) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Amount.ProtoReflect.Descriptor instead.
func (*Amount) Descriptor() ([]byte, []int) {
	return file_grpc_server_proto_rawDescGZIP(), []int{0}
}

func (x *Amount) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode ErrorCode `protobuf:"varint,1,opt,name=errorCode,proto3,enum=proto.ErrorCode" json:"errorCode,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_grpc_server_proto_rawDescGZIP(), []int{1}
}

func (x *Ack) GetErrorCode() ErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return ErrorCode_FAIL
}

// Types for result
type EmptyArgument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyArgument) Reset() {
	*x = EmptyArgument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyArgument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyArgument) ProtoMessage() {}

func (x *EmptyArgument) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyArgument.ProtoReflect.Descriptor instead.
func (*EmptyArgument) Descriptor() ([]byte, []int) {
	return file_grpc_server_proto_rawDescGZIP(), []int{2}
}

type Outcome struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutcomeType OutcomeType `protobuf:"varint,1,opt,name=outcomeType,proto3,enum=proto.OutcomeType" json:"outcomeType,omitempty"`
	Outcome     int64       `protobuf:"varint,2,opt,name=outcome,proto3" json:"outcome,omitempty"`
}

func (x *Outcome) Reset() {
	*x = Outcome{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outcome) ProtoMessage() {}

func (x *Outcome) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outcome.ProtoReflect.Descriptor instead.
func (*Outcome) Descriptor() ([]byte, []int) {
	return file_grpc_server_proto_rawDescGZIP(), []int{3}
}

func (x *Outcome) GetOutcomeType() OutcomeType {
	if x != nil {
		return x.OutcomeType
	}
	return OutcomeType_RESULT
}

func (x *Outcome) GetOutcome() int64 {
	if x != nil {
		return x.Outcome
	}
	return 0
}

var File_grpc_server_proto protoreflect.FileDescriptor

var file_grpc_server_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x06, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x35, 0x0a, 0x03,
	0x41, 0x63, 0x6b, 0x12, 0x2e, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x41, 0x72, 0x67, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0x59, 0x0a, 0x07, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12,
	0x34, 0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x75, 0x74,
	0x63, 0x6f, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x2a,
	0x31, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x08, 0x0a, 0x04,
	0x46, 0x41, 0x49, 0x4c, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x58, 0x43, 0x45, 0x50, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x02, 0x2a, 0x2a, 0x0a, 0x0b, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x48, 0x49, 0x47, 0x48, 0x45, 0x53, 0x54, 0x5f, 0x42, 0x49, 0x44, 0x10, 0x01, 0x32, 0x62,
	0x0a, 0x0e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x20, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x63, 0x6b, 0x12, 0x2e, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x75, 0x74, 0x63, 0x6f,
	0x6d, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_server_proto_rawDescOnce sync.Once
	file_grpc_server_proto_rawDescData = file_grpc_server_proto_rawDesc
)

func file_grpc_server_proto_rawDescGZIP() []byte {
	file_grpc_server_proto_rawDescOnce.Do(func() {
		file_grpc_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_server_proto_rawDescData)
	})
	return file_grpc_server_proto_rawDescData
}

var file_grpc_server_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_grpc_server_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_server_proto_goTypes = []interface{}{
	(ErrorCode)(0),        // 0: proto.ErrorCode
	(OutcomeType)(0),      // 1: proto.OutcomeType
	(*Amount)(nil),        // 2: proto.Amount
	(*Ack)(nil),           // 3: proto.Ack
	(*EmptyArgument)(nil), // 4: proto.EmptyArgument
	(*Outcome)(nil),       // 5: proto.Outcome
}
var file_grpc_server_proto_depIdxs = []int32{
	0, // 0: proto.Ack.errorCode:type_name -> proto.ErrorCode
	1, // 1: proto.Outcome.outcomeType:type_name -> proto.OutcomeType
	2, // 2: proto.AuctionService.Bid:input_type -> proto.Amount
	4, // 3: proto.AuctionService.Result:input_type -> proto.EmptyArgument
	3, // 4: proto.AuctionService.Bid:output_type -> proto.Ack
	5, // 5: proto.AuctionService.Result:output_type -> proto.Outcome
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_grpc_server_proto_init() }
func file_grpc_server_proto_init() {
	if File_grpc_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Amount); i {
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
		file_grpc_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
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
		file_grpc_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyArgument); i {
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
		file_grpc_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Outcome); i {
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
			RawDescriptor: file_grpc_server_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_server_proto_goTypes,
		DependencyIndexes: file_grpc_server_proto_depIdxs,
		EnumInfos:         file_grpc_server_proto_enumTypes,
		MessageInfos:      file_grpc_server_proto_msgTypes,
	}.Build()
	File_grpc_server_proto = out.File
	file_grpc_server_proto_rawDesc = nil
	file_grpc_server_proto_goTypes = nil
	file_grpc_server_proto_depIdxs = nil
}
