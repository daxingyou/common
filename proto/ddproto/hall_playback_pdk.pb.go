// Code generated by protoc-gen-go.
// source: hall_playback_pdk.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ProtoHeader from common_client.proto

// Ignoring public import of Heartbeat from common_client.proto

// Ignoring public import of ServerInfo from common_client.proto

// Ignoring public import of QuickConn from common_client.proto

// Ignoring public import of AckQuickConn from common_client.proto

// Ignoring public import of WeixinInfo from common_client.proto

// Ignoring public import of common_req_reg from common_client.proto

// Ignoring public import of common_req_reg_via_input from common_client.proto

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of common_req_gameLogin from common_client.proto

// Ignoring public import of common_req_gameLogin_via_input from common_client.proto

// Ignoring public import of common_ack_gameLogin from common_client.proto

// Ignoring public import of common_req_qrLogin from common_client.proto

// Ignoring public import of common_ack_qrLogin from common_client.proto

// Ignoring public import of common_req_qrWxInfo from common_client.proto

// Ignoring public import of common_ack_qrWxInfo from common_client.proto

// Ignoring public import of common_ack_reconnect from common_client.proto

// Ignoring public import of common_req_reconnect from common_client.proto

// Ignoring public import of common_req_gameState from common_client.proto

// Ignoring public import of common_ack_gameState from common_client.proto

// Ignoring public import of common_req_logout from common_client.proto

// Ignoring public import of common_ack_logout from common_client.proto

// Ignoring public import of common_req_feedback from common_client.proto

// Ignoring public import of client_base_poker from common_client.proto

// Ignoring public import of common_req_message from common_client.proto

// Ignoring public import of common_bc_message from common_client.proto

// Ignoring public import of common_req_notice from common_client.proto

// Ignoring public import of common_ack_notice from common_client.proto

// Ignoring public import of common_req_enterAgentMode from common_client.proto

// Ignoring public import of common_ack_enterAgentMode from common_client.proto

// Ignoring public import of common_req_quitAgentMode from common_client.proto

// Ignoring public import of common_ack_quitAgentMode from common_client.proto

// Ignoring public import of common_req_leaveDesk from common_client.proto

// Ignoring public import of common_ack_leaveDesk from common_client.proto

// Ignoring public import of common_bc_kickout from common_client.proto

// Ignoring public import of common_req_allowance from common_client.proto

// Ignoring public import of common_ack_allowance from common_client.proto

// Ignoring public import of common_req_applyDissolve from common_client.proto

// Ignoring public import of common_bc_applyDissolve from common_client.proto

// Ignoring public import of common_req_applyDissolveBack from common_client.proto

// Ignoring public import of common_ack_applyDissolveBack from common_client.proto

// Ignoring public import of common_ack_timeout from common_client.proto

// Ignoring public import of common_bc_userBreak from common_client.proto

// Ignoring public import of common_req_clickStatistic from common_client.proto

// Ignoring public import of common_req_offline from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// Ignoring public import of pdk_base_roomTypeInfo from common_pdk.proto

// Ignoring public import of pdk_base_timerInfo from common_pdk.proto

// Ignoring public import of pdk_enum_paiType from common_pdk.proto

// Ignoring public import of pdk_enum_playerStatus from common_pdk.proto

// Ignoring public import of pdk_enum_roomType from common_pdk.proto

// Ignoring public import of pdk_enum_enterType from common_pdk.proto

// Ignoring public import of pdk_enum_coinRoomLevel from common_pdk.proto

// Ignoring public import of pdk_enum_deskGameStatus from common_pdk.proto

// 跑得快的操作类型
type PlaybackPdkActType int32

const (
	PlaybackPdkActType_PPAT_GUO   PlaybackPdkActType = 1
	PlaybackPdkActType_PPAT_DAPAI PlaybackPdkActType = 2
)

var PlaybackPdkActType_name = map[int32]string{
	1: "PPAT_GUO",
	2: "PPAT_DAPAI",
}
var PlaybackPdkActType_value = map[string]int32{
	"PPAT_GUO":   1,
	"PPAT_DAPAI": 2,
}

func (x PlaybackPdkActType) Enum() *PlaybackPdkActType {
	p := new(PlaybackPdkActType)
	*p = x
	return p
}
func (x PlaybackPdkActType) String() string {
	return proto.EnumName(PlaybackPdkActType_name, int32(x))
}
func (x *PlaybackPdkActType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PlaybackPdkActType_value, data, "PlaybackPdkActType")
	if err != nil {
		return err
	}
	*x = PlaybackPdkActType(value)
	return nil
}
func (PlaybackPdkActType) EnumDescriptor() ([]byte, []int) { return fileDescriptor26, []int{0} }

type PlaybackPdkAckPage struct {
	Header *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// 分页信息
	Length *int32 `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	Page   *int32 `protobuf:"varint,3,opt,name=page" json:"page,omitempty"`
	Total  *int32 `protobuf:"varint,4,opt,name=total" json:"total,omitempty"`
	// 快照信息
	PlaybackSnapshots []*PdkPlaybackSnapshot `protobuf:"bytes,5,rep,name=playbackSnapshots" json:"playbackSnapshots,omitempty"`
	XXX_unrecognized  []byte                 `json:"-"`
}

func (m *PlaybackPdkAckPage) Reset()                    { *m = PlaybackPdkAckPage{} }
func (m *PlaybackPdkAckPage) String() string            { return proto.CompactTextString(m) }
func (*PlaybackPdkAckPage) ProtoMessage()               {}
func (*PlaybackPdkAckPage) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{0} }

func (m *PlaybackPdkAckPage) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PlaybackPdkAckPage) GetLength() int32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func (m *PlaybackPdkAckPage) GetPage() int32 {
	if m != nil && m.Page != nil {
		return *m.Page
	}
	return 0
}

func (m *PlaybackPdkAckPage) GetTotal() int32 {
	if m != nil && m.Total != nil {
		return *m.Total
	}
	return 0
}

func (m *PlaybackPdkAckPage) GetPlaybackSnapshots() []*PdkPlaybackSnapshot {
	if m != nil {
		return m.PlaybackSnapshots
	}
	return nil
}

// 玩家的信息
type PdkPlayerInfo struct {
	PlayerPokers     []*ClientBasePoker   `protobuf:"bytes,1,rep,name=playerPokers" json:"playerPokers,omitempty"`
	Coin             *int64               `protobuf:"varint,2,opt,name=coin" json:"coin,omitempty"`
	NickName         *string              `protobuf:"bytes,3,opt,name=nickName" json:"nickName,omitempty"`
	Sex              *int32               `protobuf:"varint,4,opt,name=sex" json:"sex,omitempty"`
	UserId           *uint32              `protobuf:"varint,5,opt,name=userId" json:"userId,omitempty"`
	IsOwner          *bool                `protobuf:"varint,6,opt,name=isOwner" json:"isOwner,omitempty"`
	BReady           *int32               `protobuf:"varint,7,opt,name=bReady" json:"bReady,omitempty"`
	Status           *PdkEnumPlayerStatus `protobuf:"varint,8,opt,name=status,enum=ddproto.PdkEnumPlayerStatus" json:"status,omitempty"`
	WxInfo           *WeixinInfo          `protobuf:"bytes,9,opt,name=wxInfo" json:"wxInfo,omitempty"`
	OnlineStatus     *int32               `protobuf:"varint,10,opt,name=onlineStatus" json:"onlineStatus,omitempty"`
	RemainPaiCount   *int32               `protobuf:"varint,11,opt,name=remainPaiCount" json:"remainPaiCount,omitempty"`
	OutPais          []*ClientBasePoker   `protobuf:"bytes,12,rep,name=outPais" json:"outPais,omitempty"`
	OutPaiType       *PdkEnumPaiType      `protobuf:"varint,13,opt,name=outPaiType,enum=ddproto.PdkEnumPaiType" json:"outPaiType,omitempty"`
	RoomCard         *int64               `protobuf:"varint,14,opt,name=roomCard" json:"roomCard,omitempty"`
	IsBanker         *bool                `protobuf:"varint,15,opt,name=isBanker" json:"isBanker,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PdkPlayerInfo) Reset()                    { *m = PdkPlayerInfo{} }
func (m *PdkPlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*PdkPlayerInfo) ProtoMessage()               {}
func (*PdkPlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{1} }

func (m *PdkPlayerInfo) GetPlayerPokers() []*ClientBasePoker {
	if m != nil {
		return m.PlayerPokers
	}
	return nil
}

func (m *PdkPlayerInfo) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *PdkPlayerInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *PdkPlayerInfo) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *PdkPlayerInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PdkPlayerInfo) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *PdkPlayerInfo) GetBReady() int32 {
	if m != nil && m.BReady != nil {
		return *m.BReady
	}
	return 0
}

func (m *PdkPlayerInfo) GetStatus() PdkEnumPlayerStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return PdkEnumPlayerStatus_PDK_NORMAL_DEFAULT
}

func (m *PdkPlayerInfo) GetWxInfo() *WeixinInfo {
	if m != nil {
		return m.WxInfo
	}
	return nil
}

func (m *PdkPlayerInfo) GetOnlineStatus() int32 {
	if m != nil && m.OnlineStatus != nil {
		return *m.OnlineStatus
	}
	return 0
}

func (m *PdkPlayerInfo) GetRemainPaiCount() int32 {
	if m != nil && m.RemainPaiCount != nil {
		return *m.RemainPaiCount
	}
	return 0
}

func (m *PdkPlayerInfo) GetOutPais() []*ClientBasePoker {
	if m != nil {
		return m.OutPais
	}
	return nil
}

func (m *PdkPlayerInfo) GetOutPaiType() PdkEnumPaiType {
	if m != nil && m.OutPaiType != nil {
		return *m.OutPaiType
	}
	return PdkEnumPaiType_PDK_ERRORCARD
}

func (m *PdkPlayerInfo) GetRoomCard() int64 {
	if m != nil && m.RoomCard != nil {
		return *m.RoomCard
	}
	return 0
}

func (m *PdkPlayerInfo) GetIsBanker() bool {
	if m != nil && m.IsBanker != nil {
		return *m.IsBanker
	}
	return false
}

// 回放快照信息（包含玩家信息、桌面信息、房间信息、右上角时间信息等），另外包含协议信息（PID 和 二进制数据）
type PdkPlaybackSnapshot struct {
	PlayerInfo       []*PdkPlayerInfo     `protobuf:"bytes,1,rep,name=PlayerInfo" json:"PlayerInfo,omitempty"`
	PlayBackDeskInfo *PdkPlaybackDeskInfo `protobuf:"bytes,2,opt,name=PlayBackDeskInfo" json:"PlayBackDeskInfo,omitempty"`
	RoomTypeInfo     *PdkBaseRoomTypeInfo `protobuf:"bytes,3,opt,name=RoomTypeInfo" json:"RoomTypeInfo,omitempty"`
	Time             *string              `protobuf:"bytes,4,opt,name=time" json:"time,omitempty"`
	ActType          *PlaybackPdkActType  `protobuf:"varint,5,opt,name=actType,enum=ddproto.PlaybackPdkActType" json:"actType,omitempty"`
	Userid           *uint32              `protobuf:"varint,6,opt,name=userid" json:"userid,omitempty"`
	OutPais          []*ClientBasePoker   `protobuf:"bytes,7,rep,name=outPais" json:"outPais,omitempty"`
	OutPaiType       *PdkEnumPaiType      `protobuf:"varint,8,opt,name=outPaiType,enum=ddproto.PdkEnumPaiType" json:"outPaiType,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PdkPlaybackSnapshot) Reset()                    { *m = PdkPlaybackSnapshot{} }
func (m *PdkPlaybackSnapshot) String() string            { return proto.CompactTextString(m) }
func (*PdkPlaybackSnapshot) ProtoMessage()               {}
func (*PdkPlaybackSnapshot) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{2} }

func (m *PdkPlaybackSnapshot) GetPlayerInfo() []*PdkPlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *PdkPlaybackSnapshot) GetPlayBackDeskInfo() *PdkPlaybackDeskInfo {
	if m != nil {
		return m.PlayBackDeskInfo
	}
	return nil
}

func (m *PdkPlaybackSnapshot) GetRoomTypeInfo() *PdkBaseRoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

func (m *PdkPlaybackSnapshot) GetTime() string {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return ""
}

func (m *PdkPlaybackSnapshot) GetActType() PlaybackPdkActType {
	if m != nil && m.ActType != nil {
		return *m.ActType
	}
	return PlaybackPdkActType_PPAT_GUO
}

func (m *PdkPlaybackSnapshot) GetUserid() uint32 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

func (m *PdkPlaybackSnapshot) GetOutPais() []*ClientBasePoker {
	if m != nil {
		return m.OutPais
	}
	return nil
}

func (m *PdkPlaybackSnapshot) GetOutPaiType() PdkEnumPaiType {
	if m != nil && m.OutPaiType != nil {
		return *m.OutPaiType
	}
	return PdkEnumPaiType_PDK_ERRORCARD
}

// 回放桌面信息
type PdkPlaybackDeskInfo struct {
	GameStatus       *int32            `protobuf:"varint,1,opt,name=GameStatus" json:"GameStatus,omitempty"`
	PlayerNum        *int32            `protobuf:"varint,3,opt,name=playerNum" json:"playerNum,omitempty"`
	ActiveUserId     *uint32           `protobuf:"varint,4,opt,name=activeUserId" json:"activeUserId,omitempty"`
	CurrPlayCount    *int32            `protobuf:"varint,10,opt,name=currPlayCount" json:"currPlayCount,omitempty"`
	TotalPlayCount   *int32            `protobuf:"varint,11,opt,name=totalPlayCount" json:"totalPlayCount,omitempty"`
	RoomNumber       *string           `protobuf:"bytes,12,opt,name=roomNumber" json:"roomNumber,omitempty"`
	BankerId         *uint32           `protobuf:"varint,14,opt,name=bankerId" json:"bankerId,omitempty"`
	EnterType        *PdkEnumEnterType `protobuf:"varint,18,opt,name=enterType,enum=ddproto.PdkEnumEnterType" json:"enterType,omitempty"`
	CoinTicket       *int64            `protobuf:"varint,19,opt,name=coinTicket" json:"coinTicket,omitempty"`
	TimerInfo        *PdkBaseTimerInfo `protobuf:"bytes,20,opt,name=timerInfo" json:"timerInfo,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *PdkPlaybackDeskInfo) Reset()                    { *m = PdkPlaybackDeskInfo{} }
func (m *PdkPlaybackDeskInfo) String() string            { return proto.CompactTextString(m) }
func (*PdkPlaybackDeskInfo) ProtoMessage()               {}
func (*PdkPlaybackDeskInfo) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{3} }

func (m *PdkPlaybackDeskInfo) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *PdkPlaybackDeskInfo) GetPlayerNum() int32 {
	if m != nil && m.PlayerNum != nil {
		return *m.PlayerNum
	}
	return 0
}

func (m *PdkPlaybackDeskInfo) GetActiveUserId() uint32 {
	if m != nil && m.ActiveUserId != nil {
		return *m.ActiveUserId
	}
	return 0
}

func (m *PdkPlaybackDeskInfo) GetCurrPlayCount() int32 {
	if m != nil && m.CurrPlayCount != nil {
		return *m.CurrPlayCount
	}
	return 0
}

func (m *PdkPlaybackDeskInfo) GetTotalPlayCount() int32 {
	if m != nil && m.TotalPlayCount != nil {
		return *m.TotalPlayCount
	}
	return 0
}

func (m *PdkPlaybackDeskInfo) GetRoomNumber() string {
	if m != nil && m.RoomNumber != nil {
		return *m.RoomNumber
	}
	return ""
}

func (m *PdkPlaybackDeskInfo) GetBankerId() uint32 {
	if m != nil && m.BankerId != nil {
		return *m.BankerId
	}
	return 0
}

func (m *PdkPlaybackDeskInfo) GetEnterType() PdkEnumEnterType {
	if m != nil && m.EnterType != nil {
		return *m.EnterType
	}
	return PdkEnumEnterType_PDK_FRIEND
}

func (m *PdkPlaybackDeskInfo) GetCoinTicket() int64 {
	if m != nil && m.CoinTicket != nil {
		return *m.CoinTicket
	}
	return 0
}

func (m *PdkPlaybackDeskInfo) GetTimerInfo() *PdkBaseTimerInfo {
	if m != nil {
		return m.TimerInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*PlaybackPdkAckPage)(nil), "ddproto.playback_pdk_ack_page")
	proto.RegisterType((*PdkPlayerInfo)(nil), "ddproto.PdkPlayerInfo")
	proto.RegisterType((*PdkPlaybackSnapshot)(nil), "ddproto.PdkPlaybackSnapshot")
	proto.RegisterType((*PdkPlaybackDeskInfo)(nil), "ddproto.PdkPlaybackDeskInfo")
	proto.RegisterEnum("ddproto.PlaybackPdkActType", PlaybackPdkActType_name, PlaybackPdkActType_value)
}

var fileDescriptor26 = []byte{
	// 667 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0x49, 0x4f, 0xdc, 0x4c,
	0x10, 0xfd, 0x3c, 0xfb, 0x14, 0xf6, 0x30, 0x78, 0x80, 0xaf, 0x03, 0x51, 0x84, 0x26, 0x39, 0x20,
	0xa2, 0x4c, 0x22, 0x94, 0xe5, 0xcc, 0x22, 0x11, 0x2e, 0x60, 0xb1, 0x28, 0x47, 0xab, 0xc7, 0xee,
	0x30, 0xd6, 0xd8, 0xdd, 0x96, 0xdd, 0x0e, 0xf0, 0x5b, 0x72, 0x4b, 0xfe, 0x55, 0x7e, 0x4d, 0xaa,
	0xcb, 0xc6, 0xac, 0x12, 0xca, 0x05, 0xe8, 0xf6, 0xab, 0xea, 0x57, 0xef, 0xbd, 0x02, 0xfe, 0x9f,
	0xf1, 0x38, 0xf6, 0xd3, 0x98, 0x5f, 0x4f, 0x79, 0x30, 0xf7, 0xd3, 0x70, 0x3e, 0x49, 0x33, 0xa5,
	0x95, 0xdb, 0x0d, 0x43, 0xfa, 0x63, 0x6d, 0x14, 0xa8, 0x24, 0x51, 0xd2, 0x0f, 0xe2, 0x48, 0x48,
	0x5d, 0x7e, 0x5d, 0x1b, 0x56, 0x97, 0x35, 0x7e, 0xfc, 0xdb, 0x82, 0x95, 0xbb, 0x6d, 0x7c, 0xfa,
	0xcd, 0x2f, 0x84, 0xfb, 0x06, 0x3a, 0x33, 0xc1, 0x43, 0x91, 0x31, 0x6b, 0xc3, 0xda, 0x5c, 0xd8,
	0x5e, 0x9e, 0x54, 0xad, 0x27, 0x9e, 0xf9, 0xf9, 0x95, 0xbe, 0xb9, 0x03, 0xe8, 0xc4, 0x42, 0x5e,
	0xe8, 0x19, 0x6b, 0x20, 0xaa, 0xed, 0xda, 0xd0, 0x32, 0xd5, 0xac, 0x49, 0x27, 0x07, 0xda, 0x5a,
	0x69, 0x1e, 0xb3, 0x16, 0x1d, 0xbf, 0xc0, 0xd2, 0xcd, 0x5b, 0xa7, 0x92, 0xa7, 0xf9, 0x4c, 0xe9,
	0x9c, 0xb5, 0x37, 0x9a, 0xd8, 0xfd, 0xe5, 0x6d, 0xf7, 0x70, 0xee, 0x3d, 0x00, 0x8d, 0x7f, 0x35,
	0xc1, 0xa9, 0xee, 0x45, 0x76, 0x28, 0xbf, 0x2b, 0xf7, 0x03, 0xd8, 0x29, 0x9d, 0x3c, 0x35, 0x17,
	0x59, 0x8e, 0x1c, 0x4d, 0x97, 0xb5, 0xba, 0x4b, 0x39, 0xb6, 0x3f, 0xe5, 0xb9, 0xf0, 0x53, 0x03,
	0x31, 0xcc, 0x02, 0x15, 0x49, 0xe2, 0xd9, 0x74, 0x87, 0xd0, 0x93, 0x51, 0x30, 0x3f, 0xe2, 0x49,
	0xc9, 0xb5, 0xef, 0x2e, 0x40, 0x33, 0x17, 0x57, 0x15, 0x53, 0x1c, 0xab, 0xc8, 0xf1, 0xa9, 0x10,
	0xe9, 0x59, 0x9b, 0x8e, 0xbb, 0x08, 0xdd, 0x28, 0x3f, 0xbe, 0x94, 0xa8, 0x46, 0x07, 0x2f, 0x7a,
	0x06, 0x30, 0x3d, 0x41, 0x09, 0xae, 0x59, 0x97, 0x0a, 0x26, 0xd0, 0xc9, 0x35, 0xd7, 0x45, 0xce,
	0x7a, 0x78, 0x1e, 0x6c, 0xbf, 0xaa, 0x99, 0x18, 0x51, 0x85, 0x2c, 0x12, 0xbf, 0xe4, 0x7b, 0x4a,
	0x28, 0xf7, 0x35, 0x74, 0x2e, 0xaf, 0xcc, 0x24, 0xac, 0x4f, 0xea, 0x8e, 0x6a, 0xfc, 0x37, 0x11,
	0x5d, 0x45, 0x92, 0x86, 0x5c, 0x06, 0x5b, 0xc9, 0x38, 0x92, 0xa2, 0x2c, 0x62, 0x40, 0x4f, 0xad,
	0xc2, 0x20, 0x13, 0x09, 0x8f, 0xa4, 0xc7, 0xa3, 0x3d, 0x55, 0x48, 0xcd, 0x16, 0xe8, 0xfe, 0x2d,
	0x74, 0x55, 0xa1, 0xf1, 0x32, 0x67, 0xf6, 0xb3, 0x6a, 0xbc, 0x03, 0x28, 0xc1, 0x67, 0xd7, 0xa9,
	0x60, 0x0e, 0x71, 0x7e, 0xf1, 0x04, 0xe7, 0x12, 0x60, 0xe4, 0xca, 0x94, 0x4a, 0xf6, 0x78, 0x16,
	0xb2, 0xc1, 0x8d, 0x80, 0x51, 0xbe, 0xcb, 0x25, 0x36, 0x63, 0x8b, 0x46, 0x92, 0xf1, 0x9f, 0x06,
	0x8c, 0x9e, 0x30, 0xcf, 0xdd, 0x02, 0xb8, 0x35, 0xae, 0x32, 0x6a, 0xf5, 0xa1, 0xdd, 0x95, 0xad,
	0x9f, 0x61, 0x68, 0x4e, 0xbb, 0x58, 0xbf, 0x2f, 0xf2, 0x39, 0x55, 0x34, 0x48, 0xa0, 0x27, 0x03,
	0x72, 0x83, 0x71, 0x3f, 0x82, 0x7d, 0x82, 0xfc, 0x0c, 0x57, 0xaa, 0x69, 0x52, 0xcd, 0x7d, 0x13,
	0x68, 0xfa, 0xec, 0x0e, 0xca, 0x44, 0x42, 0x47, 0x18, 0x80, 0x16, 0x05, 0xe0, 0x3d, 0x74, 0x79,
	0xa0, 0x49, 0x8f, 0xf6, 0x43, 0x0f, 0xef, 0x6f, 0x88, 0xf6, 0xb5, 0x11, 0xa5, 0x0a, 0x49, 0x14,
	0x52, 0x26, 0x9c, 0xbb, 0x06, 0x74, 0xff, 0xd1, 0x80, 0xde, 0x33, 0x06, 0x8c, 0x7f, 0xde, 0x17,
	0xb7, 0x1e, 0xdc, 0x05, 0x38, 0xc0, 0x0c, 0x57, 0x01, 0xb1, 0x28, 0x08, 0x4b, 0xd0, 0x2f, 0xb3,
	0x76, 0x54, 0x24, 0xd5, 0x22, 0x62, 0x92, 0x90, 0x76, 0xf4, 0x43, 0x9c, 0x97, 0xa9, 0x6e, 0x11,
	0xe1, 0x15, 0x70, 0x82, 0x22, 0xcb, 0x4c, 0xd3, 0x32, 0x48, 0x75, 0xc0, 0x68, 0x6b, 0x6f, 0xef,
	0xcb, 0x80, 0xe1, 0x5b, 0x46, 0x3e, 0xec, 0x3a, 0x45, 0xd3, 0x6d, 0x12, 0x0d, 0x63, 0x30, 0xa5,
	0x10, 0x1c, 0x96, 0xc1, 0x70, 0x70, 0x13, 0xfa, 0x38, 0xaa, 0xc8, 0x68, 0x2e, 0x97, 0xe6, 0x5a,
	0x7f, 0x3c, 0x57, 0x0d, 0x31, 0x5d, 0xcd, 0x5e, 0x9e, 0xe1, 0x36, 0x0a, 0xcd, 0x46, 0x14, 0x2e,
	0xec, 0x61, 0x8c, 0x29, 0x13, 0xb3, 0x4c, 0x5e, 0xae, 0x3f, 0xf6, 0xb2, 0x86, 0x6c, 0x7d, 0x7a,
	0xf4, 0x4f, 0xac, 0xb2, 0xc8, 0x86, 0x9e, 0xe7, 0xed, 0x9c, 0xf9, 0x07, 0xe7, 0xc7, 0x43, 0x0b,
	0x0d, 0x03, 0x3a, 0xed, 0xef, 0x78, 0x3b, 0x87, 0xc3, 0x86, 0xf7, 0x9f, 0x67, 0xfd, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0xa2, 0x61, 0xb7, 0x8a, 0x4a, 0x05, 0x00, 0x00,
}