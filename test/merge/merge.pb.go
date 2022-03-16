// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: merge.proto

package merge

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type A struct {
	B                    B        `protobuf:"bytes,1,opt,name=B,proto3" json:"B"`
	C                    []C      `protobuf:"bytes,2,rep,name=c,proto3" json:"c"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *A) Reset()         { *m = A{} }
func (m *A) String() string { return proto.CompactTextString(m) }
func (*A) ProtoMessage()    {}
func (*A) Descriptor() ([]byte, []int) {
	return fileDescriptor_82caea6f5430298b, []int{0}
}
func (m *A) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_A.Unmarshal(m, b)
}
func (m *A) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_A.Marshal(b, m, deterministic)
}
func (m *A) XXX_Merge(src proto.Message) {
	xxx_messageInfo_A.Merge(m, src)
}
func (m *A) XXX_Size() int {
	return xxx_messageInfo_A.Size(m)
}
func (m *A) XXX_DiscardUnknown() {
	xxx_messageInfo_A.DiscardUnknown(m)
}

var xxx_messageInfo_A proto.InternalMessageInfo

func (m *A) GetB() B {
	if m != nil {
		return m.B
	}
	return B{}
}

func (m *A) GetC() []C {
	if m != nil {
		return m.C
	}
	return nil
}

type B struct {
	C                    int64    `protobuf:"varint,1,opt,name=c,proto3" json:"c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *B) Reset()         { *m = B{} }
func (m *B) String() string { return proto.CompactTextString(m) }
func (*B) ProtoMessage()    {}
func (*B) Descriptor() ([]byte, []int) {
	return fileDescriptor_82caea6f5430298b, []int{1}
}
func (m *B) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_B.Unmarshal(m, b)
}
func (m *B) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_B.Marshal(b, m, deterministic)
}
func (m *B) XXX_Merge(src proto.Message) {
	xxx_messageInfo_B.Merge(m, src)
}
func (m *B) XXX_Size() int {
	return xxx_messageInfo_B.Size(m)
}
func (m *B) XXX_DiscardUnknown() {
	xxx_messageInfo_B.DiscardUnknown(m)
}

var xxx_messageInfo_B proto.InternalMessageInfo

func (m *B) GetC() int64 {
	if m != nil {
		return m.C
	}
	return 0
}

type C struct {
	D                    int64    `protobuf:"varint,1,opt,name=d,proto3" json:"d,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C) Reset()         { *m = C{} }
func (m *C) String() string { return proto.CompactTextString(m) }
func (*C) ProtoMessage()    {}
func (*C) Descriptor() ([]byte, []int) {
	return fileDescriptor_82caea6f5430298b, []int{2}
}
func (m *C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C.Unmarshal(m, b)
}
func (m *C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C.Marshal(b, m, deterministic)
}
func (m *C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C.Merge(m, src)
}
func (m *C) XXX_Size() int {
	return xxx_messageInfo_C.Size(m)
}
func (m *C) XXX_DiscardUnknown() {
	xxx_messageInfo_C.DiscardUnknown(m)
}

var xxx_messageInfo_C proto.InternalMessageInfo

func (m *C) GetD() int64 {
	if m != nil {
		return m.D
	}
	return 0
}

func init() {
	proto.RegisterType((*A)(nil), "merge.A")
	proto.RegisterType((*B)(nil), "merge.B")
	proto.RegisterType((*C)(nil), "merge.C")
}

func init() { proto.RegisterFile("merge.proto", fileDescriptor_82caea6f5430298b) }

var fileDescriptor_82caea6f5430298b = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x4d, 0x2d, 0x4a,
	0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0xa4, 0x0c, 0xd2, 0x33, 0x4b,
	0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x93, 0xf3, 0x8b, 0x73, 0xf3, 0x8b, 0xf5, 0xd3,
	0xf3, 0xd3, 0xf3, 0xc1, 0x6a, 0xd0, 0x58, 0x10, 0x8d, 0x4a, 0xf6, 0x5c, 0x8c, 0x8e, 0x42, 0x32,
	0x5c, 0x8c, 0x4e, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x1c, 0x7a, 0x10, 0x63, 0x9d, 0x9c,
	0x58, 0x4e, 0xdc, 0x93, 0x67, 0x08, 0x62, 0x74, 0x02, 0xc9, 0x26, 0x4b, 0x30, 0x29, 0x30, 0x23,
	0xc9, 0x3a, 0xc3, 0x64, 0x93, 0x95, 0x04, 0xb9, 0x18, 0x9d, 0x84, 0x78, 0x40, 0x4a, 0x40, 0x06,
	0x30, 0x43, 0x85, 0x9c, 0x41, 0x42, 0x29, 0x30, 0xa1, 0x94, 0x24, 0x36, 0xb0, 0x6d, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x0a, 0x1a, 0xa9, 0xb5, 0x00, 0x00, 0x00,
}
