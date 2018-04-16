// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/RangelReale/fproto-wrap-std/jsontag.proto

package gwproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type JSONTag struct {
	TagDisable   bool   `protobuf:"varint,1,opt,name=tag_disable,json=tagDisable" json:"tag_disable,omitempty"`
	TagFieldname string `protobuf:"bytes,2,opt,name=tag_fieldname,json=tagFieldname" json:"tag_fieldname,omitempty"`
}

func (m *JSONTag) Reset()                    { *m = JSONTag{} }
func (m *JSONTag) String() string            { return proto.CompactTextString(m) }
func (*JSONTag) ProtoMessage()               {}
func (*JSONTag) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *JSONTag) GetTagDisable() bool {
	if m != nil {
		return m.TagDisable
	}
	return false
}

func (m *JSONTag) GetTagFieldname() string {
	if m != nil {
		return m.TagFieldname
	}
	return ""
}

var E_Jsontag = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*JSONTag)(nil),
	Field:         6700,
	Name:          "fproto_wrap.jsontag",
	Tag:           "bytes,6700,opt,name=jsontag",
	Filename:      "github.com/RangelReale/fproto-wrap-std/jsontag.proto",
}

func init() {
	proto.RegisterType((*JSONTag)(nil), "fproto_wrap.JSONTag")
	proto.RegisterExtension(E_Jsontag)
}

func init() {
	proto.RegisterFile("github.com/RangelReale/fproto-wrap-std/jsontag.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4b, 0xc4, 0x40,
	0x10, 0x85, 0x89, 0x85, 0xa7, 0x1b, 0x6d, 0x82, 0x45, 0x10, 0xc4, 0xa0, 0x4d, 0x9a, 0xdb, 0x05,
	0x3d, 0x11, 0x2c, 0x45, 0x2c, 0x2c, 0x2e, 0xb0, 0x5a, 0xd9, 0x1c, 0x93, 0xcb, 0x66, 0x5c, 0xd9,
	0xcb, 0x84, 0xec, 0x1c, 0xf7, 0xa7, 0xfc, 0x91, 0x92, 0xdd, 0x04, 0x2c, 0xad, 0x66, 0xf8, 0xde,
	0x63, 0xde, 0x63, 0xc4, 0x0a, 0x2d, 0x7f, 0xed, 0x6b, 0xb9, 0xa5, 0x9d, 0xd2, 0xd0, 0xa1, 0x71,
	0xda, 0x80, 0x33, 0xaa, 0xed, 0x07, 0x62, 0x5a, 0x1e, 0x06, 0xe8, 0x97, 0x9e, 0x1b, 0xf5, 0xed,
	0xa9, 0x63, 0x40, 0x19, 0x70, 0x96, 0x46, 0x79, 0x33, 0xca, 0x97, 0x05, 0x12, 0xa1, 0x33, 0x2a,
	0xa0, 0x7a, 0xdf, 0xaa, 0xc6, 0xf8, 0xed, 0x60, 0x7b, 0xa6, 0x21, 0xda, 0x6f, 0x2a, 0xb1, 0x78,
	0x7b, 0xaf, 0xd6, 0x1f, 0x80, 0xd9, 0xb5, 0x48, 0x19, 0x70, 0xd3, 0x58, 0x0f, 0xb5, 0x33, 0x79,
	0x52, 0x24, 0xe5, 0x89, 0x16, 0x0c, 0xf8, 0x12, 0x49, 0x76, 0x2b, 0xce, 0x47, 0x43, 0x6b, 0x8d,
	0x6b, 0x3a, 0xd8, 0x99, 0xfc, 0xa8, 0x48, 0xca, 0x53, 0x7d, 0xc6, 0x80, 0xaf, 0x33, 0x7b, 0x5a,
	0x8b, 0xc5, 0x54, 0x28, 0xbb, 0x92, 0x31, 0x5e, 0xce, 0xf1, 0x32, 0xd8, 0xaa, 0x9e, 0x2d, 0x75,
	0x3e, 0xff, 0x59, 0x15, 0x49, 0x99, 0xde, 0x5d, 0xc8, 0x3f, 0x8d, 0xe5, 0x54, 0x46, 0xcf, 0x47,
	0x9e, 0x1f, 0x3f, 0x1f, 0xfe, 0xf9, 0x07, 0xa4, 0x71, 0x55, 0x78, 0x88, 0x71, 0xc7, 0x61, 0xdc,
	0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x62, 0xec, 0x88, 0x2b, 0x47, 0x01, 0x00, 0x00,
}
