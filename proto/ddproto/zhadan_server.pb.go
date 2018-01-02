// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zhadan_server.proto

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of common_srv_pokerPai from common_server_poker.proto

// Ignoring public import of zhadan_client_poker from zhadan_base.proto

// Ignoring public import of zhadan_user_bill from zhadan_base.proto

// Ignoring public import of zhadan_desk_option from zhadan_base.proto

// Ignoring public import of zhadan_srv_room from zhadan_base.proto

// Ignoring public import of zhadan_enum_protoid from zhadan_base.proto

// Ignoring public import of zhadan_enum_pokerType from zhadan_base.proto

// Ignoring public import of zhadan_enum_desk_status from zhadan_base.proto

// Ignoring public import of zhadan_gan_opt from zhadan_base.proto

// 打出去的牌
type ZhadanSrvPoker struct {
	Pais             []*CommonSrvPokerPai `protobuf:"bytes,1,rep,name=pais" json:"pais,omitempty"`
	Type             *ZhadanEnumPokerType `protobuf:"varint,2,opt,name=Type,enum=ddproto.ZhadanEnumPokerType" json:"Type,omitempty"`
	KeyValue         *int32               `protobuf:"varint,3,opt,name=keyValue" json:"keyValue,omitempty"`
	KeyLength        *int32               `protobuf:"varint,4,opt,name=keyLength" json:"keyLength,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *ZhadanSrvPoker) Reset()                    { *m = ZhadanSrvPoker{} }
func (m *ZhadanSrvPoker) String() string            { return proto.CompactTextString(m) }
func (*ZhadanSrvPoker) ProtoMessage()               {}
func (*ZhadanSrvPoker) Descriptor() ([]byte, []int) { return fileDescriptor58, []int{0} }

func (m *ZhadanSrvPoker) GetPais() []*CommonSrvPokerPai {
	if m != nil {
		return m.Pais
	}
	return nil
}

func (m *ZhadanSrvPoker) GetType() ZhadanEnumPokerType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ZhadanEnumPokerType_ZHADAN_POKER_TYPE_OTHER
}

func (m *ZhadanSrvPoker) GetKeyValue() int32 {
	if m != nil && m.KeyValue != nil {
		return *m.KeyValue
	}
	return 0
}

func (m *ZhadanSrvPoker) GetKeyLength() int32 {
	if m != nil && m.KeyLength != nil {
		return *m.KeyLength
	}
	return 0
}

// desk 的信息
type ZhadanSrvDesk struct {
	DeskId         *int32                `protobuf:"varint,1,opt,name=deskId" json:"deskId,omitempty"`
	Pwd            *string               `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
	GameNumber     *int32                `protobuf:"varint,3,opt,name=gameNumber" json:"gameNumber,omitempty"`
	RoomId         *int32                `protobuf:"varint,4,opt,name=roomId" json:"roomId,omitempty"`
	Status         *ZhadanEnumDeskStatus `protobuf:"varint,5,opt,name=status,enum=ddproto.ZhadanEnumDeskStatus" json:"status,omitempty"`
	CircleNo       *int32                `protobuf:"varint,7,opt,name=circleNo" json:"circleNo,omitempty"`
	CurrDeskScore  *int32                `protobuf:"varint,8,opt,name=currDeskScore" json:"currDeskScore,omitempty"`
	DeskScorePais  *ZhadanSrvPoker       `protobuf:"bytes,29,opt,name=deskScorePais" json:"deskScorePais,omitempty"`
	ShowPoker      *int32                `protobuf:"varint,30,opt,name=showPoker" json:"showPoker,omitempty"`
	Owner          *uint32               `protobuf:"varint,11,opt,name=owner" json:"owner,omitempty"`
	IsDaikai       *bool                 `protobuf:"varint,12,opt,name=isDaikai" json:"isDaikai,omitempty"`
	DaikaiUser     *uint32               `protobuf:"varint,13,opt,name=daikaiUser" json:"daikaiUser,omitempty"`
	DeskOption     *ZhadanDeskOption     `protobuf:"bytes,14,opt,name=deskOption" json:"deskOption,omitempty"`
	LastActUser    *uint32               `protobuf:"varint,15,opt,name=lastActUser" json:"lastActUser,omitempty"`
	LastChupaiUser *uint32               `protobuf:"varint,16,opt,name=lastChupaiUser" json:"lastChupaiUser,omitempty"`
	CurrActUser    *uint32               `protobuf:"varint,26,opt,name=currActUser" json:"currActUser,omitempty"`
	CanGuo         *bool                 `protobuf:"varint,27,opt,name=canGuo" json:"canGuo,omitempty"`
	BaoZhuangUser  *uint32               `protobuf:"varint,28,opt,name=baoZhuangUser" json:"baoZhuangUser,omitempty"`
	IsOnDissolve   *bool                 `protobuf:"varint,18,opt,name=isOnDissolve" json:"isOnDissolve,omitempty"`
	DissolveTime   *int64                `protobuf:"varint,19,opt,name=dissolveTime" json:"dissolveTime,omitempty"`
	IsStart        *bool                 `protobuf:"varint,22,opt,name=isStart" json:"isStart,omitempty"`
	OneStartTime   *int64                `protobuf:"varint,23,opt,name=oneStartTime" json:"oneStartTime,omitempty"`
	AllStartTime   *int64                `protobuf:"varint,24,opt,name=allStartTime" json:"allStartTime,omitempty"`
	DissolveUser   *uint32               `protobuf:"varint,25,opt,name=dissolveUser" json:"dissolveUser,omitempty"`
	// 金币场
	IsCoinRoom       *bool  `protobuf:"varint,20,opt,name=isCoinRoom" json:"isCoinRoom,omitempty"`
	SurplusTime      *int32 `protobuf:"varint,21,opt,name=surplusTime" json:"surplusTime,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ZhadanSrvDesk) Reset()                    { *m = ZhadanSrvDesk{} }
func (m *ZhadanSrvDesk) String() string            { return proto.CompactTextString(m) }
func (*ZhadanSrvDesk) ProtoMessage()               {}
func (*ZhadanSrvDesk) Descriptor() ([]byte, []int) { return fileDescriptor58, []int{1} }

func (m *ZhadanSrvDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *ZhadanSrvDesk) GetPwd() string {
	if m != nil && m.Pwd != nil {
		return *m.Pwd
	}
	return ""
}

func (m *ZhadanSrvDesk) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *ZhadanSrvDesk) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *ZhadanSrvDesk) GetStatus() ZhadanEnumDeskStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ZhadanEnumDeskStatus_ZHADAN_DESK_STATUS_WAIT_READY
}

func (m *ZhadanSrvDesk) GetCircleNo() int32 {
	if m != nil && m.CircleNo != nil {
		return *m.CircleNo
	}
	return 0
}

func (m *ZhadanSrvDesk) GetCurrDeskScore() int32 {
	if m != nil && m.CurrDeskScore != nil {
		return *m.CurrDeskScore
	}
	return 0
}

func (m *ZhadanSrvDesk) GetDeskScorePais() *ZhadanSrvPoker {
	if m != nil {
		return m.DeskScorePais
	}
	return nil
}

func (m *ZhadanSrvDesk) GetShowPoker() int32 {
	if m != nil && m.ShowPoker != nil {
		return *m.ShowPoker
	}
	return 0
}

func (m *ZhadanSrvDesk) GetOwner() uint32 {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return 0
}

func (m *ZhadanSrvDesk) GetIsDaikai() bool {
	if m != nil && m.IsDaikai != nil {
		return *m.IsDaikai
	}
	return false
}

func (m *ZhadanSrvDesk) GetDaikaiUser() uint32 {
	if m != nil && m.DaikaiUser != nil {
		return *m.DaikaiUser
	}
	return 0
}

func (m *ZhadanSrvDesk) GetDeskOption() *ZhadanDeskOption {
	if m != nil {
		return m.DeskOption
	}
	return nil
}

func (m *ZhadanSrvDesk) GetLastActUser() uint32 {
	if m != nil && m.LastActUser != nil {
		return *m.LastActUser
	}
	return 0
}

func (m *ZhadanSrvDesk) GetLastChupaiUser() uint32 {
	if m != nil && m.LastChupaiUser != nil {
		return *m.LastChupaiUser
	}
	return 0
}

func (m *ZhadanSrvDesk) GetCurrActUser() uint32 {
	if m != nil && m.CurrActUser != nil {
		return *m.CurrActUser
	}
	return 0
}

func (m *ZhadanSrvDesk) GetCanGuo() bool {
	if m != nil && m.CanGuo != nil {
		return *m.CanGuo
	}
	return false
}

func (m *ZhadanSrvDesk) GetBaoZhuangUser() uint32 {
	if m != nil && m.BaoZhuangUser != nil {
		return *m.BaoZhuangUser
	}
	return 0
}

func (m *ZhadanSrvDesk) GetIsOnDissolve() bool {
	if m != nil && m.IsOnDissolve != nil {
		return *m.IsOnDissolve
	}
	return false
}

func (m *ZhadanSrvDesk) GetDissolveTime() int64 {
	if m != nil && m.DissolveTime != nil {
		return *m.DissolveTime
	}
	return 0
}

func (m *ZhadanSrvDesk) GetIsStart() bool {
	if m != nil && m.IsStart != nil {
		return *m.IsStart
	}
	return false
}

func (m *ZhadanSrvDesk) GetOneStartTime() int64 {
	if m != nil && m.OneStartTime != nil {
		return *m.OneStartTime
	}
	return 0
}

func (m *ZhadanSrvDesk) GetAllStartTime() int64 {
	if m != nil && m.AllStartTime != nil {
		return *m.AllStartTime
	}
	return 0
}

func (m *ZhadanSrvDesk) GetDissolveUser() uint32 {
	if m != nil && m.DissolveUser != nil {
		return *m.DissolveUser
	}
	return 0
}

func (m *ZhadanSrvDesk) GetIsCoinRoom() bool {
	if m != nil && m.IsCoinRoom != nil {
		return *m.IsCoinRoom
	}
	return false
}

func (m *ZhadanSrvDesk) GetSurplusTime() int32 {
	if m != nil && m.SurplusTime != nil {
		return *m.SurplusTime
	}
	return 0
}

// 用户属性
type ZhadanSrvUser struct {
	UserId           *uint32           `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Bill             *ZhadanUserBill   `protobuf:"bytes,2,opt,name=bill" json:"bill,omitempty"`
	IsOnline         *bool             `protobuf:"varint,3,opt,name=isOnline" json:"isOnline,omitempty"`
	Index            *int32            `protobuf:"varint,4,opt,name=index" json:"index,omitempty"`
	Pokers           *ZhadanSrvPoker   `protobuf:"bytes,5,opt,name=pokers" json:"pokers,omitempty"`
	OutPai           *ZhadanSrvPoker   `protobuf:"bytes,6,opt,name=out_pai,json=outPai" json:"out_pai,omitempty"`
	IsPass           *bool             `protobuf:"varint,7,opt,name=isPass" json:"isPass,omitempty"`
	DeskScore        *int32            `protobuf:"varint,16,opt,name=deskScore" json:"deskScore,omitempty"`
	IsBanker         *bool             `protobuf:"varint,17,opt,name=isBanker" json:"isBanker,omitempty"`
	BaozhuangStatus  *int32            `protobuf:"varint,18,opt,name=baozhuangStatus" json:"baozhuangStatus,omitempty"`
	ScorePais        *ZhadanSrvPoker   `protobuf:"bytes,22,opt,name=score_pais,json=scorePais" json:"score_pais,omitempty"`
	ExtScorePoker    []*ZhadanSrvPoker `protobuf:"bytes,23,rep,name=ext_score_poker,json=extScorePoker" json:"ext_score_poker,omitempty"`
	ExtScore         *int32            `protobuf:"varint,24,opt,name=ext_score,json=extScore" json:"ext_score,omitempty"`
	PaoScore         *int32            `protobuf:"varint,25,opt,name=paoScore" json:"paoScore,omitempty"`
	IsReady          *bool             `protobuf:"varint,8,opt,name=isReady" json:"isReady,omitempty"`
	WxInfo           *WeixinInfo       `protobuf:"bytes,10,opt,name=wxInfo" json:"wxInfo,omitempty"`
	DissolveState    *int32            `protobuf:"varint,11,opt,name=dissolveState" json:"dissolveState,omitempty"`
	IsRobot          *bool             `protobuf:"varint,12,opt,name=isRobot" json:"isRobot,omitempty"`
	IsOnWhiteList    *bool             `protobuf:"varint,13,opt,name=isOnWhiteList" json:"isOnWhiteList,omitempty"`
	WhiteWinRate     *int32            `protobuf:"varint,14,opt,name=whiteWinRate" json:"whiteWinRate,omitempty"`
	IsLeave          *bool             `protobuf:"varint,15,opt,name=isLeave" json:"isLeave,omitempty"`
	PaoNo            *int32            `protobuf:"varint,19,opt,name=paoNo" json:"paoNo,omitempty"`
	IsOnGaming       *bool             `protobuf:"varint,20,opt,name=isOnGaming" json:"isOnGaming,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *ZhadanSrvUser) Reset()                    { *m = ZhadanSrvUser{} }
func (m *ZhadanSrvUser) String() string            { return proto.CompactTextString(m) }
func (*ZhadanSrvUser) ProtoMessage()               {}
func (*ZhadanSrvUser) Descriptor() ([]byte, []int) { return fileDescriptor58, []int{2} }

func (m *ZhadanSrvUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ZhadanSrvUser) GetBill() *ZhadanUserBill {
	if m != nil {
		return m.Bill
	}
	return nil
}

func (m *ZhadanSrvUser) GetIsOnline() bool {
	if m != nil && m.IsOnline != nil {
		return *m.IsOnline
	}
	return false
}

func (m *ZhadanSrvUser) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *ZhadanSrvUser) GetPokers() *ZhadanSrvPoker {
	if m != nil {
		return m.Pokers
	}
	return nil
}

func (m *ZhadanSrvUser) GetOutPai() *ZhadanSrvPoker {
	if m != nil {
		return m.OutPai
	}
	return nil
}

func (m *ZhadanSrvUser) GetIsPass() bool {
	if m != nil && m.IsPass != nil {
		return *m.IsPass
	}
	return false
}

func (m *ZhadanSrvUser) GetDeskScore() int32 {
	if m != nil && m.DeskScore != nil {
		return *m.DeskScore
	}
	return 0
}

func (m *ZhadanSrvUser) GetIsBanker() bool {
	if m != nil && m.IsBanker != nil {
		return *m.IsBanker
	}
	return false
}

func (m *ZhadanSrvUser) GetBaozhuangStatus() int32 {
	if m != nil && m.BaozhuangStatus != nil {
		return *m.BaozhuangStatus
	}
	return 0
}

func (m *ZhadanSrvUser) GetScorePais() *ZhadanSrvPoker {
	if m != nil {
		return m.ScorePais
	}
	return nil
}

func (m *ZhadanSrvUser) GetExtScorePoker() []*ZhadanSrvPoker {
	if m != nil {
		return m.ExtScorePoker
	}
	return nil
}

func (m *ZhadanSrvUser) GetExtScore() int32 {
	if m != nil && m.ExtScore != nil {
		return *m.ExtScore
	}
	return 0
}

func (m *ZhadanSrvUser) GetPaoScore() int32 {
	if m != nil && m.PaoScore != nil {
		return *m.PaoScore
	}
	return 0
}

func (m *ZhadanSrvUser) GetIsReady() bool {
	if m != nil && m.IsReady != nil {
		return *m.IsReady
	}
	return false
}

func (m *ZhadanSrvUser) GetWxInfo() *WeixinInfo {
	if m != nil {
		return m.WxInfo
	}
	return nil
}

func (m *ZhadanSrvUser) GetDissolveState() int32 {
	if m != nil && m.DissolveState != nil {
		return *m.DissolveState
	}
	return 0
}

func (m *ZhadanSrvUser) GetIsRobot() bool {
	if m != nil && m.IsRobot != nil {
		return *m.IsRobot
	}
	return false
}

func (m *ZhadanSrvUser) GetIsOnWhiteList() bool {
	if m != nil && m.IsOnWhiteList != nil {
		return *m.IsOnWhiteList
	}
	return false
}

func (m *ZhadanSrvUser) GetWhiteWinRate() int32 {
	if m != nil && m.WhiteWinRate != nil {
		return *m.WhiteWinRate
	}
	return 0
}

func (m *ZhadanSrvUser) GetIsLeave() bool {
	if m != nil && m.IsLeave != nil {
		return *m.IsLeave
	}
	return false
}

func (m *ZhadanSrvUser) GetPaoNo() int32 {
	if m != nil && m.PaoNo != nil {
		return *m.PaoNo
	}
	return 0
}

func (m *ZhadanSrvUser) GetIsOnGaming() bool {
	if m != nil && m.IsOnGaming != nil {
		return *m.IsOnGaming
	}
	return false
}

// desk快照索引列表
type ZhadanSrvDeskSnapshotIdIndex struct {
	DeskId           []int32 `protobuf:"varint,1,rep,name=deskId" json:"deskId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ZhadanSrvDeskSnapshotIdIndex) Reset()                    { *m = ZhadanSrvDeskSnapshotIdIndex{} }
func (m *ZhadanSrvDeskSnapshotIdIndex) String() string            { return proto.CompactTextString(m) }
func (*ZhadanSrvDeskSnapshotIdIndex) ProtoMessage()               {}
func (*ZhadanSrvDeskSnapshotIdIndex) Descriptor() ([]byte, []int) { return fileDescriptor58, []int{3} }

func (m *ZhadanSrvDeskSnapshotIdIndex) GetDeskId() []int32 {
	if m != nil {
		return m.DeskId
	}
	return nil
}

// 牌桌快照
type ZhadanSrvDeskSnapshot struct {
	DeskState        *ZhadanSrvDesk   `protobuf:"bytes,1,opt,name=deskState" json:"deskState,omitempty"`
	Users            []*ZhadanSrvUser `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *ZhadanSrvDeskSnapshot) Reset()                    { *m = ZhadanSrvDeskSnapshot{} }
func (m *ZhadanSrvDeskSnapshot) String() string            { return proto.CompactTextString(m) }
func (*ZhadanSrvDeskSnapshot) ProtoMessage()               {}
func (*ZhadanSrvDeskSnapshot) Descriptor() ([]byte, []int) { return fileDescriptor58, []int{4} }

func (m *ZhadanSrvDeskSnapshot) GetDeskState() *ZhadanSrvDesk {
	if m != nil {
		return m.DeskState
	}
	return nil
}

func (m *ZhadanSrvDeskSnapshot) GetUsers() []*ZhadanSrvUser {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*ZhadanSrvPoker)(nil), "ddproto.zhadan_srv_poker")
	proto.RegisterType((*ZhadanSrvDesk)(nil), "ddproto.zhadan_srv_desk")
	proto.RegisterType((*ZhadanSrvUser)(nil), "ddproto.zhadan_srv_user")
	proto.RegisterType((*ZhadanSrvDeskSnapshotIdIndex)(nil), "ddproto.zhadan_srv_desk_snapshot_id_index")
	proto.RegisterType((*ZhadanSrvDeskSnapshot)(nil), "ddproto.zhadan_srv_desk_snapshot")
}

func init() { proto.RegisterFile("zhadan_server.proto", fileDescriptor58) }

var fileDescriptor58 = []byte{
	// 963 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x55, 0x5f, 0x6f, 0xdb, 0x36,
	0x10, 0x9f, 0xea, 0xf8, 0x4f, 0x98, 0x3a, 0x49, 0x99, 0x2e, 0x65, 0xfe, 0x2c, 0xf0, 0x8c, 0x61,
	0x30, 0x30, 0x2c, 0xd8, 0xfc, 0x30, 0x14, 0xe8, 0xc3, 0xd0, 0x35, 0x40, 0x11, 0x20, 0x48, 0x0c,
	0xa6, 0x5b, 0x80, 0xbd, 0x08, 0xb4, 0xc5, 0xc5, 0x84, 0x25, 0x52, 0x10, 0xa5, 0xd8, 0xe9, 0xe3,
	0xbe, 0xc4, 0xbe, 0xc7, 0x3e, 0xc3, 0x3e, 0xd8, 0x70, 0x47, 0x4a, 0x96, 0x8c, 0xb4, 0x7d, 0xb2,
	0x7e, 0x3f, 0xfe, 0xee, 0x78, 0x77, 0xbc, 0x3b, 0x93, 0x83, 0x8f, 0x73, 0x11, 0x09, 0x1d, 0x5a,
	0x99, 0x3d, 0xc8, 0xec, 0x3c, 0xcd, 0x4c, 0x6e, 0x68, 0x37, 0x8a, 0xf0, 0xe3, 0xf8, 0x68, 0x66,
	0x92, 0xc4, 0x94, 0xa7, 0x61, 0x6a, 0x16, 0xa5, 0xe6, 0xf8, 0x85, 0x37, 0x9c, 0x0a, 0x2b, 0x1d,
	0x35, 0xfc, 0x37, 0x20, 0xfb, 0xa5, 0xbb, 0xec, 0xc1, 0xa9, 0xe9, 0x4f, 0x64, 0x2b, 0x15, 0xca,
	0xb2, 0x60, 0xd0, 0x1a, 0xed, 0x8c, 0x4f, 0xcf, 0xbd, 0xeb, 0xf3, 0xd2, 0x73, 0x29, 0x9c, 0x08,
	0xc5, 0x51, 0x49, 0xc7, 0x64, 0xeb, 0xc3, 0x63, 0x2a, 0xd9, 0xb3, 0x41, 0x30, 0xda, 0x1d, 0x9f,
	0x55, 0x16, 0xde, 0xb5, 0xd4, 0x45, 0xe2, 0x4c, 0x40, 0xc5, 0x51, 0x4b, 0x8f, 0x49, 0x6f, 0x21,
	0x1f, 0xff, 0x10, 0x71, 0x21, 0x59, 0x6b, 0x10, 0x8c, 0xda, 0xbc, 0xc2, 0xf4, 0x94, 0x6c, 0x2f,
	0xe4, 0xe3, 0x95, 0xd4, 0xf7, 0xf9, 0x9c, 0x6d, 0xe1, 0xe1, 0x9a, 0x18, 0xfe, 0xd3, 0x25, 0x7b,
	0xb5, 0xa0, 0x23, 0x69, 0x17, 0xf4, 0x90, 0x74, 0xe0, 0xf7, 0x32, 0x62, 0x01, 0xca, 0x3d, 0xa2,
	0xfb, 0xa4, 0x95, 0x2e, 0x23, 0x0c, 0x6c, 0x9b, 0xc3, 0x27, 0x3d, 0x23, 0xe4, 0x5e, 0x24, 0xf2,
	0xba, 0x48, 0xa6, 0x32, 0xf3, 0x37, 0xd7, 0x18, 0xf0, 0x94, 0x19, 0x93, 0x5c, 0x46, 0xfe, 0x62,
	0x8f, 0xe8, 0x6b, 0xd2, 0xb1, 0xb9, 0xc8, 0x0b, 0xcb, 0xda, 0x98, 0xe5, 0xe0, 0xc9, 0x2c, 0xe1,
	0xda, 0xd0, 0xe9, 0xb8, 0xd7, 0x43, 0xa6, 0x33, 0x95, 0xcd, 0x62, 0x79, 0x6d, 0x58, 0xd7, 0x65,
	0x5a, 0x62, 0xfa, 0x1d, 0xe9, 0xcf, 0x8a, 0x2c, 0xbb, 0x90, 0x76, 0x71, 0x3b, 0x33, 0x99, 0x64,
	0x3d, 0x14, 0x34, 0x49, 0xfa, 0x2b, 0xe9, 0x47, 0x25, 0x98, 0xc0, 0xd3, 0x7c, 0x33, 0x08, 0x46,
	0x3b, 0xe3, 0xa3, 0xcd, 0x10, 0xaa, 0xa7, 0xe1, 0x4d, 0x3d, 0x14, 0xd4, 0xce, 0xcd, 0x72, 0x02,
	0x67, 0xec, 0xcc, 0x15, 0xb4, 0x22, 0xe8, 0x4b, 0xd2, 0x36, 0x4b, 0x2d, 0x33, 0xb6, 0x33, 0x08,
	0x46, 0x7d, 0xee, 0x00, 0x84, 0xad, 0xec, 0x85, 0x50, 0x0b, 0xa1, 0xd8, 0xf3, 0x41, 0x30, 0xea,
	0xf1, 0x0a, 0x43, 0x11, 0x23, 0xfc, 0xfa, 0xdd, 0xca, 0x8c, 0xf5, 0xd1, 0xac, 0xc6, 0xd0, 0x37,
	0x84, 0x40, 0x00, 0x37, 0x69, 0xae, 0x8c, 0x66, 0xbb, 0x18, 0xed, 0xc9, 0x66, 0xb4, 0x58, 0x2b,
	0x83, 0x12, 0x5e, 0x93, 0xd3, 0x01, 0xd9, 0x89, 0x85, 0xcd, 0xdf, 0xce, 0x72, 0xf4, 0xbe, 0x87,
	0xde, 0xeb, 0x14, 0xfd, 0x9e, 0xec, 0x02, 0x7c, 0x37, 0x2f, 0x52, 0x1f, 0xc2, 0x3e, 0x8a, 0x36,
	0x58, 0xf0, 0x04, 0x85, 0x2c, 0x3d, 0x1d, 0x3b, 0x4f, 0x35, 0x0a, 0x5e, 0x7b, 0x26, 0xf4, 0xfb,
	0xc2, 0xb0, 0x13, 0x4c, 0xd1, 0x23, 0x78, 0x97, 0xa9, 0x30, 0x7f, 0xce, 0x0b, 0xa1, 0xef, 0xd1,
	0xf6, 0x14, 0x6d, 0x9b, 0x24, 0x1d, 0x92, 0xe7, 0xca, 0xde, 0xe8, 0x0b, 0x65, 0xad, 0x89, 0x1f,
	0x24, 0xa3, 0xe8, 0xa3, 0xc1, 0x81, 0x26, 0xf2, 0xdf, 0x1f, 0x54, 0x22, 0xd9, 0xc1, 0x20, 0x18,
	0xb5, 0x78, 0x83, 0xa3, 0x8c, 0x74, 0x95, 0xbd, 0xcd, 0x45, 0x96, 0xb3, 0x43, 0x74, 0x51, 0x42,
	0xb0, 0x36, 0x5a, 0xe2, 0x37, 0x5a, 0xbf, 0x72, 0xd6, 0x75, 0x0e, 0x34, 0x22, 0x8e, 0xd7, 0x1a,
	0xe6, 0x34, 0x75, 0xae, 0x1e, 0x05, 0xa6, 0x73, 0x84, 0xe9, 0x34, 0x38, 0x78, 0x54, 0x65, 0xdf,
	0x19, 0xa5, 0xb9, 0x31, 0x09, 0x7b, 0x89, 0x81, 0xd4, 0x18, 0xa8, 0xa6, 0x2d, 0xb2, 0x34, 0x2e,
	0x2c, 0x5e, 0xf3, 0x35, 0xb6, 0x51, 0x9d, 0x1a, 0xfe, 0xd7, 0x69, 0x4c, 0x66, 0xe1, 0x2b, 0x0c,
	0xbf, 0x7e, 0x32, 0xfb, 0xdc, 0x23, 0xfa, 0x23, 0xd9, 0x9a, 0xaa, 0x38, 0xc6, 0xd1, 0x7c, 0xa2,
	0x95, 0x41, 0x15, 0x82, 0x80, 0xa3, 0xcc, 0x75, 0xe3, 0x8d, 0x8e, 0x95, 0x76, 0xeb, 0x02, 0xbb,
	0xd1, 0x61, 0xe8, 0x5f, 0xa5, 0x23, 0xb9, 0xf2, 0x13, 0xeb, 0x00, 0xfd, 0x99, 0x74, 0x70, 0x16,
	0xdc, 0xc0, 0x7e, 0x76, 0x5a, 0xbc, 0x90, 0x8e, 0x49, 0xd7, 0x14, 0x79, 0x98, 0x0a, 0xc5, 0x3a,
	0x5f, 0xb4, 0x31, 0x45, 0x3e, 0x11, 0x0a, 0xf2, 0x53, 0x76, 0x22, 0xac, 0xc5, 0xd9, 0xee, 0x71,
	0x8f, 0x60, 0xe4, 0xaa, 0x19, 0xc4, 0xf6, 0x6c, 0xf3, 0x35, 0xe1, 0xd2, 0xf9, 0x4d, 0x68, 0x98,
	0xc7, 0x17, 0x65, 0x3a, 0x0e, 0xd3, 0x11, 0xd9, 0x9b, 0x0a, 0xf3, 0x11, 0xdb, 0xec, 0xd6, 0xad,
	0x1c, 0x8a, 0xf6, 0x9b, 0x34, 0x7d, 0x4d, 0x88, 0x05, 0x77, 0x21, 0xee, 0xeb, 0xc3, 0x2f, 0x85,
	0xbc, 0x6d, 0xab, 0x85, 0xf0, 0x96, 0xec, 0xc9, 0x55, 0x1e, 0x7a, 0x6b, 0x5c, 0x0b, 0xaf, 0x70,
	0xdd, 0x7f, 0x6e, 0xa7, 0xc8, 0x55, 0xee, 0x56, 0x0a, 0x6e, 0x8d, 0x13, 0xb2, 0x5d, 0xb9, 0xc0,
	0x9e, 0x6b, 0xf3, 0x5e, 0xa9, 0x80, 0xfc, 0x52, 0x61, 0x5c, 0xf2, 0x47, 0xee, 0xac, 0xc4, 0xae,
	0xdb, 0xb9, 0x14, 0xd1, 0x23, 0x6e, 0x3b, 0xec, 0x76, 0x84, 0xf4, 0x07, 0xd2, 0x59, 0xae, 0x2e,
	0xf5, 0x5f, 0x86, 0x11, 0xcc, 0xe5, 0xa0, 0x0a, 0xe6, 0x4e, 0xaa, 0x95, 0xd2, 0x70, 0xc4, 0xbd,
	0x04, 0x46, 0xb4, 0x6c, 0x5f, 0x28, 0x87, 0xc4, 0xed, 0xd5, 0xe6, 0x4d, 0xd2, 0x5f, 0x66, 0xa6,
	0x26, 0xf7, 0x4b, 0xac, 0x84, 0x60, 0x0f, 0x1d, 0x74, 0x37, 0x57, 0xb9, 0xbc, 0x52, 0x36, 0xc7,
	0x35, 0xd6, 0xe3, 0x4d, 0x12, 0x06, 0x67, 0x09, 0xe0, 0x4e, 0x69, 0x0e, 0x97, 0xec, 0xe2, 0x25,
	0x0d, 0xce, 0xdd, 0x71, 0x25, 0xc5, 0x83, 0xc4, 0x65, 0x85, 0x77, 0x20, 0x84, 0xce, 0x4c, 0x85,
	0xb9, 0x36, 0x38, 0xf5, 0x6d, 0xee, 0x80, 0x1b, 0xb4, 0x1b, 0xfd, 0x5e, 0x24, 0x4a, 0xdf, 0xaf,
	0x07, 0xad, 0x64, 0x86, 0x6f, 0xc8, 0xb7, 0x1b, 0xff, 0x6f, 0xa1, 0xd5, 0x22, 0xb5, 0x73, 0x93,
	0x87, 0x2a, 0x0a, 0x5d, 0x7b, 0xd7, 0xff, 0xf1, 0x5a, 0xeb, 0x7f, 0xbc, 0xe1, 0xdf, 0x01, 0x61,
	0x9f, 0xb2, 0xa6, 0xbf, 0xf8, 0xa6, 0xc4, 0x7a, 0x05, 0x58, 0x63, 0xf6, 0xd4, 0x83, 0x83, 0x88,
	0xaf, 0xa5, 0xf4, 0x9c, 0xb4, 0x61, 0x20, 0x2d, 0x7b, 0x86, 0x4d, 0xf2, 0xa4, 0x0d, 0x08, 0xb8,
	0x93, 0x4d, 0xbe, 0x9a, 0x04, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xcd, 0x42, 0x99, 0xa8,
	0x08, 0x00, 0x00,
}
