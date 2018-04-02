// Code generated by protoc-gen-go. DO NOT EDIT.
// source: jsonobject.proto

/*
Package gwproto is a generated protocol buffer package.

It is generated from these files:
	jsonobject.proto
	jsontag.proto
	options.proto
	sqltag.proto
	time.proto
	uuid.proto

It has these top-level messages:
	JSONObject
	JSONTag
	FProtoWrapOptions
	SQLTag
	NullTime
	UUID
	NullUUID
*/
package gwproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type JSONObject struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *JSONObject) Reset()                    { *m = JSONObject{} }
func (m *JSONObject) String() string            { return proto.CompactTextString(m) }
func (*JSONObject) ProtoMessage()               {}
func (*JSONObject) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *JSONObject) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*JSONObject)(nil), "fproto_wrap.JSONObject")
}

func init() { proto.RegisterFile("jsonobject.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x2a, 0xce, 0xcf,
	0xcb, 0x4f, 0xca, 0x4a, 0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0x03,
	0xd3, 0xf1, 0xe5, 0x45, 0x89, 0x05, 0x4a, 0x4a, 0x5c, 0x5c, 0x5e, 0xc1, 0xfe, 0x7e, 0xfe, 0x60,
	0x05, 0x42, 0x22, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0x10, 0x8e, 0x93, 0x79, 0x94, 0x69, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e,
	0xae, 0x7e, 0x50, 0x62, 0x5e, 0x7a, 0x6a, 0x4e, 0x50, 0x6a, 0x62, 0x4e, 0xaa, 0x3e, 0xc4, 0x24,
	0x5d, 0x90, 0x49, 0xba, 0xc5, 0x25, 0x29, 0xfa, 0xe9, 0xf9, 0x20, 0xa6, 0x7e, 0x7a, 0x39, 0x58,
	0x3c, 0x89, 0x0d, 0x4c, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc3, 0xd4, 0x8d, 0xc9, 0x84,
	0x00, 0x00, 0x00,
}
