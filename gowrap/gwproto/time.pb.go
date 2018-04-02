// Code generated by protoc-gen-go. DO NOT EDIT.
// source: time.proto

package gwproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NullTime struct {
	Value *google_protobuf1.Timestamp `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	Valid bool                        `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *NullTime) Reset()                    { *m = NullTime{} }
func (m *NullTime) String() string            { return proto.CompactTextString(m) }
func (*NullTime) ProtoMessage()               {}
func (*NullTime) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *NullTime) GetValue() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *NullTime) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func init() {
	proto.RegisterType((*NullTime)(nil), "fproto_wrap.NullTime")
}

func init() { proto.RegisterFile("time.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xc9, 0xcc, 0x4d,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0x03, 0xd3, 0xf1, 0xe5, 0x45, 0x89, 0x05,
	0x52, 0xf2, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x60, 0xa1, 0xa4, 0xd2, 0x34, 0x7d, 0x90,
	0xc2, 0xe2, 0x92, 0xc4, 0xdc, 0x02, 0x88, 0x6a, 0xa5, 0x20, 0x2e, 0x0e, 0xbf, 0xd2, 0x9c, 0x9c,
	0x90, 0xcc, 0xdc, 0x54, 0x21, 0x03, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0x6e, 0x23, 0x29, 0x3d, 0x88, 0x66, 0x3d, 0x98, 0x66, 0xbd, 0x10, 0x98, 0xe6, 0x20,
	0x88, 0x42, 0x21, 0x11, 0xb0, 0x8e, 0xcc, 0x14, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x08,
	0xc7, 0xc9, 0x3c, 0xca, 0x34, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f,
	0x28, 0x31, 0x2f, 0x3d, 0x35, 0x27, 0x28, 0x35, 0x31, 0x27, 0x55, 0x1f, 0xe2, 0x34, 0x5d, 0x90,
	0xd3, 0x74, 0x8b, 0x4b, 0x52, 0xf4, 0xd3, 0xf3, 0x41, 0x4c, 0xfd, 0xf4, 0x72, 0x88, 0x15, 0x6c,
	0x60, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x21, 0x58, 0xe0, 0xee, 0xcf, 0x00, 0x00, 0x00,
}
