// Code generated by protoc-gen-go.
// source: common_server_task.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TTask struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *TTask) Reset()                    { *m = TTask{} }
func (m *TTask) String() string            { return proto.CompactTextString(m) }
func (*TTask) ProtoMessage()               {}
func (*TTask) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{0} }

func init() {
	proto.RegisterType((*TTask)(nil), "ddproto.TTask")
}

func init() { proto.RegisterFile("common_server_task.proto", fileDescriptor17) }

var fileDescriptor17 = []byte{
	// 65 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0x8b, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x8a, 0x2f, 0x49, 0x2c, 0xce, 0xd6, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x49, 0x01, 0x33, 0x94, 0xd8, 0xb9, 0x58, 0x43, 0x42,
	0x12, 0x8b, 0xb3, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x45, 0xb3, 0xe1, 0x2c, 0x00, 0x00,
	0x00,
}
