// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/RangelReale/fproto-wrap-std/time.proto

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

func init() { proto.RegisterFile("github.com/RangelReale/fproto-wrap-std/time.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4c, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x0f, 0x4a, 0xcc, 0x4b, 0x4f, 0xcd, 0x09, 0x4a, 0x4d,
	0xcc, 0x49, 0xd5, 0x4f, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2d, 0x2f, 0x4a, 0x2c, 0xd0, 0x2d,
	0x2e, 0x49, 0xd1, 0x2f, 0xc9, 0xcc, 0x4d, 0xd5, 0x03, 0x8b, 0x09, 0x71, 0x43, 0xe4, 0xe2, 0x41,
	0x72, 0x52, 0xf2, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x60, 0xa1, 0xa4, 0xd2, 0x34, 0xb0,
	0xc2, 0xe2, 0x92, 0xc4, 0xdc, 0x02, 0x88, 0x6a, 0xa5, 0x20, 0x2e, 0x0e, 0xbf, 0xd2, 0x9c, 0x9c,
	0x90, 0xcc, 0xdc, 0x54, 0x21, 0x03, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0x6e, 0x23, 0x29, 0x3d, 0x88, 0x66, 0x3d, 0x98, 0x66, 0xbd, 0x10, 0x98, 0xe6, 0x20,
	0x88, 0x42, 0x21, 0x11, 0xb0, 0x8e, 0xcc, 0x14, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x08,
	0xc7, 0xc9, 0x3c, 0xca, 0x94, 0x48, 0x67, 0xa7, 0xe7, 0x83, 0x98, 0xfa, 0xe9, 0xe5, 0x10, 0x2b,
	0xd8, 0xc0, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x4c, 0x35, 0x35, 0xf6, 0x00, 0x00,
	0x00,
}
