// Code generated by protoc-gen-go.
// source: common_server_award.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// todo 等待增加
type Award int32

const (
	Award_taskType_time Award = 1
	Award_taskType_2    Award = 2
)

var Award_name = map[int32]string{
	1: "taskType_time",
	2: "taskType_2",
}
var Award_value = map[string]int32{
	"taskType_time": 1,
	"taskType_2":    2,
}

func (x Award) Enum() *Award {
	p := new(Award)
	*p = x
	return p
}
func (x Award) String() string {
	return proto.EnumName(Award_name, int32(x))
}
func (x *Award) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Award_value, data, "Award")
	if err != nil {
		return err
	}
	*x = Award(value)
	return nil
}
func (Award) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

// 用户奖励，保存到数据库的结构
type Taward struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	TaskType         *int32  `protobuf:"varint,2,opt,name=taskType" json:"taskType,omitempty"`
	AwardType        *int32  `protobuf:"varint,3,opt,name=awardType" json:"awardType,omitempty"`
	UserId           *uint32 `protobuf:"varint,4,opt,name=userId" json:"userId,omitempty"`
	Memo             *string `protobuf:"bytes,5,opt,name=memo" json:"memo,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Taward) Reset()                    { *m = Taward{} }
func (m *Taward) String() string            { return proto.CompactTextString(m) }
func (*Taward) ProtoMessage()               {}
func (*Taward) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Taward) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Taward) GetTaskType() int32 {
	if m != nil && m.TaskType != nil {
		return *m.TaskType
	}
	return 0
}

func (m *Taward) GetAwardType() int32 {
	if m != nil && m.AwardType != nil {
		return *m.AwardType
	}
	return 0
}

func (m *Taward) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Taward) GetMemo() string {
	if m != nil && m.Memo != nil {
		return *m.Memo
	}
	return ""
}

func init() {
	proto.RegisterType((*Taward)(nil), "ddproto.Taward")
	proto.RegisterEnum("ddproto.Award", Award_name, Award_value)
}

var fileDescriptor3 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0x8b, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x8a, 0x4f, 0x2c, 0x4f, 0x2c, 0x4a, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x49, 0x01, 0x33, 0x94, 0xc2, 0xb9, 0xd8, 0x42,
	0xc0, 0x12, 0x42, 0x5c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x42, 0x02,
	0x5c, 0x1c, 0x25, 0x89, 0xc5, 0xd9, 0x21, 0x95, 0x05, 0xa9, 0x12, 0x4c, 0x60, 0x11, 0x41, 0x2e,
	0x4e, 0xb0, 0x32, 0xb0, 0x10, 0x33, 0x58, 0x88, 0x8f, 0x8b, 0xad, 0x14, 0x68, 0xb4, 0x67, 0x8a,
	0x04, 0x0b, 0x90, 0xcf, 0x2b, 0xc4, 0xc3, 0xc5, 0x92, 0x9b, 0x9a, 0x9b, 0x2f, 0xc1, 0x0a, 0xe4,
	0x71, 0x6a, 0x69, 0x71, 0xb1, 0x42, 0xcc, 0x15, 0xe4, 0xe2, 0x85, 0x99, 0x15, 0x5f, 0x92, 0x99,
	0x9b, 0x2a, 0xc0, 0x08, 0xd4, 0xc9, 0x05, 0x17, 0x32, 0x12, 0x60, 0x02, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xe8, 0x51, 0x0d, 0x88, 0xa9, 0x00, 0x00, 0x00,
}