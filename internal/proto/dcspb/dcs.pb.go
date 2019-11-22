// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dcs.proto

package dcspb

import (
	context "context"
	fmt "fmt"
	sourcebackendpb "github.com/Debian/dcs/internal/proto/sourcebackendpb"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Error_ErrorType int32

const (
	Error_CANCELLED           Error_ErrorType = 0
	Error_BACKEND_UNAVAILABLE Error_ErrorType = 1
	Error_FAILED              Error_ErrorType = 2
	Error_INVALID_QUERY       Error_ErrorType = 3
)

var Error_ErrorType_name = map[int32]string{
	0: "CANCELLED",
	1: "BACKEND_UNAVAILABLE",
	2: "FAILED",
	3: "INVALID_QUERY",
}

var Error_ErrorType_value = map[string]int32{
	"CANCELLED":           0,
	"BACKEND_UNAVAILABLE": 1,
	"FAILED":              2,
	"INVALID_QUERY":       3,
}

func (x Error_ErrorType) String() string {
	return proto.EnumName(Error_ErrorType_name, int32(x))
}

func (Error_ErrorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_14f789ee6ef427d2, []int{1, 0}
}

type Event_Type int32

const (
	Event_ERROR      Event_Type = 0
	Event_PROGRESS   Event_Type = 1
	Event_MATCH      Event_Type = 2
	Event_PAGINATION Event_Type = 3
	Event_DONE       Event_Type = 4
)

var Event_Type_name = map[int32]string{
	0: "ERROR",
	1: "PROGRESS",
	2: "MATCH",
	3: "PAGINATION",
	4: "DONE",
}

var Event_Type_value = map[string]int32{
	"ERROR":      0,
	"PROGRESS":   1,
	"MATCH":      2,
	"PAGINATION": 3,
	"DONE":       4,
}

func (x Event_Type) String() string {
	return proto.EnumName(Event_Type_name, int32(x))
}

func (Event_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_14f789ee6ef427d2, []int{4, 0}
}

type SearchRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Literal              bool     `protobuf:"varint,2,opt,name=literal,proto3" json:"literal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_14f789ee6ef427d2, []int{0}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchRequest) GetLiteral() bool {
	if m != nil {
		return m.Literal
	}
	return false
}

type Error struct {
	Type                 Error_ErrorType `protobuf:"varint,1,opt,name=type,proto3,enum=dcspb.Error_ErrorType" json:"type,omitempty"`
	Message              string          `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_14f789ee6ef427d2, []int{1}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetType() Error_ErrorType {
	if m != nil {
		return m.Type
	}
	return Error_CANCELLED
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Progress struct {
	QueryId              string   `protobuf:"bytes,1,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	FilesProcessed       int64    `protobuf:"varint,2,opt,name=files_processed,json=filesProcessed,proto3" json:"files_processed,omitempty"`
	FilesTotal           int64    `protobuf:"varint,3,opt,name=files_total,json=filesTotal,proto3" json:"files_total,omitempty"`
	Results              int64    `protobuf:"varint,4,opt,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Progress) Reset()         { *m = Progress{} }
func (m *Progress) String() string { return proto.CompactTextString(m) }
func (*Progress) ProtoMessage()    {}
func (*Progress) Descriptor() ([]byte, []int) {
	return fileDescriptor_14f789ee6ef427d2, []int{2}
}

func (m *Progress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Progress.Unmarshal(m, b)
}
func (m *Progress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Progress.Marshal(b, m, deterministic)
}
func (m *Progress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Progress.Merge(m, src)
}
func (m *Progress) XXX_Size() int {
	return xxx_messageInfo_Progress.Size(m)
}
func (m *Progress) XXX_DiscardUnknown() {
	xxx_messageInfo_Progress.DiscardUnknown(m)
}

var xxx_messageInfo_Progress proto.InternalMessageInfo

func (m *Progress) GetQueryId() string {
	if m != nil {
		return m.QueryId
	}
	return ""
}

func (m *Progress) GetFilesProcessed() int64 {
	if m != nil {
		return m.FilesProcessed
	}
	return 0
}

func (m *Progress) GetFilesTotal() int64 {
	if m != nil {
		return m.FilesTotal
	}
	return 0
}

func (m *Progress) GetResults() int64 {
	if m != nil {
		return m.Results
	}
	return 0
}

type Pagination struct {
	QueryId              string   `protobuf:"bytes,1,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	ResultPages          int64    `protobuf:"varint,2,opt,name=result_pages,json=resultPages,proto3" json:"result_pages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pagination) Reset()         { *m = Pagination{} }
func (m *Pagination) String() string { return proto.CompactTextString(m) }
func (*Pagination) ProtoMessage()    {}
func (*Pagination) Descriptor() ([]byte, []int) {
	return fileDescriptor_14f789ee6ef427d2, []int{3}
}

func (m *Pagination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pagination.Unmarshal(m, b)
}
func (m *Pagination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pagination.Marshal(b, m, deterministic)
}
func (m *Pagination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pagination.Merge(m, src)
}
func (m *Pagination) XXX_Size() int {
	return xxx_messageInfo_Pagination.Size(m)
}
func (m *Pagination) XXX_DiscardUnknown() {
	xxx_messageInfo_Pagination.DiscardUnknown(m)
}

var xxx_messageInfo_Pagination proto.InternalMessageInfo

func (m *Pagination) GetQueryId() string {
	if m != nil {
		return m.QueryId
	}
	return ""
}

func (m *Pagination) GetResultPages() int64 {
	if m != nil {
		return m.ResultPages
	}
	return 0
}

type Event struct {
	// Types that are valid to be assigned to Data:
	//	*Event_Error
	//	*Event_Progress
	//	*Event_Match
	//	*Event_Pagination
	Data                 isEvent_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_14f789ee6ef427d2, []int{4}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

type isEvent_Data interface {
	isEvent_Data()
}

type Event_Error struct {
	Error *Error `protobuf:"bytes,1,opt,name=error,proto3,oneof"`
}

type Event_Progress struct {
	Progress *Progress `protobuf:"bytes,2,opt,name=progress,proto3,oneof"`
}

type Event_Match struct {
	Match *sourcebackendpb.Match `protobuf:"bytes,3,opt,name=match,proto3,oneof"`
}

type Event_Pagination struct {
	Pagination *Pagination `protobuf:"bytes,4,opt,name=pagination,proto3,oneof"`
}

func (*Event_Error) isEvent_Data() {}

func (*Event_Progress) isEvent_Data() {}

func (*Event_Match) isEvent_Data() {}

func (*Event_Pagination) isEvent_Data() {}

func (m *Event) GetData() isEvent_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Event) GetError() *Error {
	if x, ok := m.GetData().(*Event_Error); ok {
		return x.Error
	}
	return nil
}

func (m *Event) GetProgress() *Progress {
	if x, ok := m.GetData().(*Event_Progress); ok {
		return x.Progress
	}
	return nil
}

func (m *Event) GetMatch() *sourcebackendpb.Match {
	if x, ok := m.GetData().(*Event_Match); ok {
		return x.Match
	}
	return nil
}

func (m *Event) GetPagination() *Pagination {
	if x, ok := m.GetData().(*Event_Pagination); ok {
		return x.Pagination
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Event) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Event_OneofMarshaler, _Event_OneofUnmarshaler, _Event_OneofSizer, []interface{}{
		(*Event_Error)(nil),
		(*Event_Progress)(nil),
		(*Event_Match)(nil),
		(*Event_Pagination)(nil),
	}
}

func _Event_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Event)
	// data
	switch x := m.Data.(type) {
	case *Event_Error:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case *Event_Progress:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Progress); err != nil {
			return err
		}
	case *Event_Match:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Match); err != nil {
			return err
		}
	case *Event_Pagination:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Pagination); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Event.Data has unexpected type %T", x)
	}
	return nil
}

func _Event_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Event)
	switch tag {
	case 1: // data.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Error)
		err := b.DecodeMessage(msg)
		m.Data = &Event_Error{msg}
		return true, err
	case 2: // data.progress
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Progress)
		err := b.DecodeMessage(msg)
		m.Data = &Event_Progress{msg}
		return true, err
	case 3: // data.match
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(sourcebackendpb.Match)
		err := b.DecodeMessage(msg)
		m.Data = &Event_Match{msg}
		return true, err
	case 4: // data.pagination
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Pagination)
		err := b.DecodeMessage(msg)
		m.Data = &Event_Pagination{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Event_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Event)
	// data
	switch x := m.Data.(type) {
	case *Event_Error:
		s := proto.Size(x.Error)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_Progress:
		s := proto.Size(x.Progress)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_Match:
		s := proto.Size(x.Match)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_Pagination:
		s := proto.Size(x.Pagination)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterEnum("dcspb.Error_ErrorType", Error_ErrorType_name, Error_ErrorType_value)
	proto.RegisterEnum("dcspb.Event_Type", Event_Type_name, Event_Type_value)
	proto.RegisterType((*SearchRequest)(nil), "dcspb.SearchRequest")
	proto.RegisterType((*Error)(nil), "dcspb.Error")
	proto.RegisterType((*Progress)(nil), "dcspb.Progress")
	proto.RegisterType((*Pagination)(nil), "dcspb.Pagination")
	proto.RegisterType((*Event)(nil), "dcspb.Event")
}

func init() { proto.RegisterFile("dcs.proto", fileDescriptor_14f789ee6ef427d2) }

var fileDescriptor_14f789ee6ef427d2 = []byte{
	// 567 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x8d, 0x93, 0x38, 0x4d, 0x6e, 0xfa, 0xe3, 0xce, 0x57, 0xf5, 0x0b, 0xdd, 0x50, 0x0c, 0x12,
	0x55, 0x05, 0x76, 0xe5, 0x2e, 0x58, 0x22, 0x27, 0x36, 0x8d, 0x21, 0x75, 0xcc, 0x24, 0xad, 0x04,
	0x9b, 0x68, 0x6c, 0x0f, 0xa9, 0x85, 0x6b, 0xbb, 0x33, 0x13, 0xa4, 0x3c, 0x02, 0x2b, 0x9e, 0x81,
	0x37, 0x45, 0x1e, 0xc7, 0x51, 0xcb, 0x82, 0x4d, 0xa4, 0xf3, 0x33, 0x77, 0xee, 0x39, 0x19, 0x43,
	0x2f, 0x8e, 0xb8, 0x51, 0xb0, 0x5c, 0xe4, 0x48, 0x8d, 0x23, 0x5e, 0x84, 0x27, 0x2f, 0x79, 0xbe,
	0x62, 0x11, 0x0d, 0x49, 0xf4, 0x9d, 0x66, 0x71, 0x11, 0x9a, 0x4f, 0x70, 0xe5, 0xd5, 0xdf, 0xc3,
	0xde, 0x8c, 0x12, 0x16, 0xdd, 0x61, 0xfa, 0xb0, 0xa2, 0x5c, 0xa0, 0x23, 0x50, 0x1f, 0x56, 0x94,
	0xad, 0x07, 0xca, 0xa9, 0x72, 0xd6, 0xc3, 0x15, 0x40, 0x03, 0xd8, 0x49, 0x13, 0x41, 0x19, 0x49,
	0x07, 0xcd, 0x53, 0xe5, 0xac, 0x8b, 0x6b, 0xa8, 0xff, 0x56, 0x40, 0x75, 0x19, 0xcb, 0x19, 0x3a,
	0x87, 0xb6, 0x58, 0x17, 0x54, 0x1e, 0xdc, 0xb7, 0x8e, 0x0d, 0xb9, 0x85, 0x21, 0xb5, 0xea, 0x77,
	0xbe, 0x2e, 0x28, 0x96, 0x9e, 0x72, 0xde, 0x3d, 0xe5, 0x9c, 0x2c, 0xa9, 0x9c, 0xd7, 0xc3, 0x35,
	0xd4, 0x31, 0xf4, 0xb6, 0x66, 0xb4, 0x07, 0xbd, 0x91, 0xed, 0x8f, 0xdc, 0xc9, 0xc4, 0x75, 0xb4,
	0x06, 0xfa, 0x1f, 0xfe, 0x1b, 0xda, 0xa3, 0x4f, 0xae, 0xef, 0x2c, 0x6e, 0x7c, 0xfb, 0xd6, 0xf6,
	0x26, 0xf6, 0x70, 0xe2, 0x6a, 0x0a, 0x02, 0xe8, 0x7c, 0xb0, 0xbd, 0xd2, 0xd4, 0x44, 0x87, 0xb0,
	0xe7, 0xf9, 0xb7, 0xf6, 0xc4, 0x73, 0x16, 0x9f, 0x6f, 0x5c, 0xfc, 0x45, 0x6b, 0xe9, 0x3f, 0x15,
	0xe8, 0x06, 0x2c, 0x5f, 0x32, 0xca, 0x39, 0x7a, 0x06, 0x5d, 0x99, 0x69, 0x91, 0xc4, 0x9b, 0x8c,
	0x3b, 0x12, 0x7b, 0x31, 0x7a, 0x0d, 0x07, 0xdf, 0x92, 0x94, 0xf2, 0x45, 0xc1, 0xf2, 0x88, 0x72,
	0x4e, 0x63, 0xb9, 0x5d, 0x0b, 0xef, 0x4b, 0x3a, 0xa8, 0x59, 0xf4, 0x1c, 0xfa, 0x95, 0x51, 0xe4,
	0x82, 0xa4, 0x83, 0x96, 0x34, 0x81, 0xa4, 0xe6, 0x25, 0x53, 0xe6, 0x63, 0x94, 0xaf, 0x52, 0xc1,
	0x07, 0x6d, 0x29, 0xd6, 0x50, 0xff, 0x08, 0x10, 0x90, 0x65, 0x92, 0x11, 0x91, 0xe4, 0xd9, 0xbf,
	0x96, 0x79, 0x01, 0xbb, 0xd5, 0x99, 0x45, 0x41, 0x96, 0x94, 0x6f, 0x36, 0xe9, 0x57, 0x5c, 0x50,
	0x52, 0xfa, 0xaf, 0x26, 0xa8, 0xee, 0x0f, 0x9a, 0x09, 0xf4, 0x0a, 0x54, 0x5a, 0xb6, 0x26, 0x87,
	0xf4, 0xad, 0xdd, 0xc7, 0xe5, 0x8f, 0x1b, 0xb8, 0x12, 0xd1, 0x5b, 0xe8, 0x16, 0x9b, 0x1a, 0xe4,
	0xb8, 0xbe, 0x75, 0xb0, 0x31, 0xd6, 0xed, 0x8c, 0x1b, 0x78, 0x6b, 0x41, 0x06, 0xa8, 0xf7, 0x44,
	0x44, 0x77, 0x32, 0x5f, 0xdf, 0x3a, 0x36, 0xfe, 0x7a, 0x50, 0xc6, 0x75, 0xa9, 0x96, 0xe3, 0xa5,
	0x0d, 0x5d, 0x02, 0x14, 0xdb, 0x68, 0x32, 0x77, 0xdf, 0x3a, 0xac, 0x2f, 0xd8, 0x0a, 0xe3, 0x06,
	0x7e, 0x64, 0xd3, 0x1d, 0x68, 0xcb, 0xbf, 0xba, 0x07, 0xaa, 0x8b, 0xf1, 0x14, 0x6b, 0x0d, 0xb4,
	0x0b, 0xdd, 0x00, 0x4f, 0xaf, 0xb0, 0x3b, 0x9b, 0x69, 0x4a, 0x29, 0x5c, 0xdb, 0xf3, 0xd1, 0x58,
	0x6b, 0xa2, 0x7d, 0x80, 0xc0, 0xbe, 0xf2, 0x7c, 0x7b, 0xee, 0x4d, 0x7d, 0xad, 0x85, 0xba, 0xd0,
	0x76, 0xa6, 0xbe, 0xab, 0xb5, 0x87, 0x1d, 0x68, 0xc7, 0x44, 0x10, 0xeb, 0x1d, 0xb4, 0x9c, 0xd1,
	0x0c, 0x5d, 0x40, 0xa7, 0x7a, 0xd5, 0xe8, 0x68, 0x73, 0xff, 0x93, 0x47, 0x7e, 0xb2, 0xed, 0xa7,
	0x2c, 0x4f, 0x6f, 0x5c, 0x28, 0xc3, 0x37, 0x5f, 0xcf, 0x97, 0x89, 0xb8, 0x5b, 0x85, 0x46, 0x94,
	0xdf, 0x9b, 0x0e, 0x0d, 0x13, 0x92, 0x99, 0x71, 0xc4, 0xcd, 0x24, 0x13, 0x94, 0x65, 0x24, 0x35,
	0xe5, 0xf7, 0x62, 0xca, 0x73, 0x61, 0x47, 0x82, 0xcb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa0,
	0x29, 0x80, 0x03, 0x75, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DCSClient is the client API for DCS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DCSClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (DCS_SearchClient, error)
}

type dCSClient struct {
	cc *grpc.ClientConn
}

func NewDCSClient(cc *grpc.ClientConn) DCSClient {
	return &dCSClient{cc}
}

func (c *dCSClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (DCS_SearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DCS_serviceDesc.Streams[0], "/dcspb.DCS/Search", opts...)
	if err != nil {
		return nil, err
	}
	x := &dCSSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DCS_SearchClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type dCSSearchClient struct {
	grpc.ClientStream
}

func (x *dCSSearchClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DCSServer is the server API for DCS service.
type DCSServer interface {
	Search(*SearchRequest, DCS_SearchServer) error
}

func RegisterDCSServer(s *grpc.Server, srv DCSServer) {
	s.RegisterService(&_DCS_serviceDesc, srv)
}

func _DCS_Search_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DCSServer).Search(m, &dCSSearchServer{stream})
}

type DCS_SearchServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type dCSSearchServer struct {
	grpc.ServerStream
}

func (x *dCSSearchServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

var _DCS_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dcspb.DCS",
	HandlerType: (*DCSServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Search",
			Handler:       _DCS_Search_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dcs.proto",
}
