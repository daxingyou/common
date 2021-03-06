// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common_server_config.proto

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 对应mgo 中的表单,但是这里主要是做一些配置
type TConfigSys struct {
	Id               *int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	NewUserCoin      *int64 `protobuf:"varint,2,opt,name=newUserCoin" json:"newUserCoin,omitempty"`
	NewUserRoomcard  *int64 `protobuf:"varint,3,opt,name=newUserRoomcard" json:"newUserRoomcard,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TConfigSys) Reset()                    { *m = TConfigSys{} }
func (m *TConfigSys) String() string            { return proto.CompactTextString(m) }
func (*TConfigSys) ProtoMessage()               {}
func (*TConfigSys) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *TConfigSys) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TConfigSys) GetNewUserCoin() int64 {
	if m != nil && m.NewUserCoin != nil {
		return *m.NewUserCoin
	}
	return 0
}

func (m *TConfigSys) GetNewUserRoomcard() int64 {
	if m != nil && m.NewUserRoomcard != nil {
		return *m.NewUserRoomcard
	}
	return 0
}

// 转盘抽奖的配置
type TConfigDrawLottery struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Type             *int32  `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TConfigDrawLottery) Reset()                    { *m = TConfigDrawLottery{} }
func (m *TConfigDrawLottery) String() string            { return proto.CompactTextString(m) }
func (*TConfigDrawLottery) ProtoMessage()               {}
func (*TConfigDrawLottery) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *TConfigDrawLottery) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TConfigDrawLottery) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *TConfigDrawLottery) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func init() {
	proto.RegisterType((*TConfigSys)(nil), "ddproto.TConfigSys")
	proto.RegisterType((*TConfigDrawLottery)(nil), "ddproto.TConfigDrawLottery")
}

func init() { proto.RegisterFile("common_server_config.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x8d, 0x31, 0xcf, 0x82, 0x40,
	0x0c, 0x40, 0x03, 0x7c, 0xe4, 0x8b, 0x35, 0xd1, 0xa4, 0x13, 0x71, 0x22, 0x4c, 0x4c, 0xfe, 0x09,
	0x1c, 0x99, 0x4e, 0x9d, 0x09, 0xe1, 0xaa, 0xde, 0x70, 0x57, 0x52, 0x2e, 0x12, 0xfe, 0xbd, 0xb1,
	0x61, 0x30, 0xba, 0xbd, 0xbe, 0x36, 0x7d, 0x70, 0x18, 0xd8, 0x7b, 0x0e, 0xdd, 0x44, 0xf2, 0x24,
	0xe9, 0x06, 0x0e, 0x37, 0x77, 0x3f, 0x8e, 0xc2, 0x91, 0xf1, 0xdf, 0x5a, 0x85, 0xea, 0x01, 0x70,
	0x69, 0x74, 0x73, 0x5e, 0x26, 0xdc, 0x41, 0xea, 0x6c, 0x91, 0x94, 0x49, 0x9d, 0x9b, 0xd4, 0x59,
	0x2c, 0x61, 0x1b, 0x68, 0xbe, 0x4e, 0x24, 0x0d, 0xbb, 0x50, 0xa4, 0x65, 0x52, 0x67, 0xe6, 0x53,
	0x61, 0x0d, 0xfb, 0x75, 0x34, 0xcc, 0x7e, 0xe8, 0xc5, 0x16, 0x99, 0x5e, 0x7d, 0xeb, 0xaa, 0x05,
	0x5c, 0x4b, 0x27, 0xe9, 0xe7, 0x96, 0x63, 0x24, 0x59, 0x7e, 0x8a, 0x08, 0x7f, 0xa1, 0xf7, 0xa4,
	0xa9, 0x8d, 0x51, 0x7e, 0xbb, 0xb8, 0x8c, 0xa4, 0x8f, 0x73, 0xa3, 0xfc, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0xf4, 0x41, 0x82, 0xf0, 0xdd, 0x00, 0x00, 0x00,
}
