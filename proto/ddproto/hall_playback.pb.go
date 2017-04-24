// Code generated by protoc-gen-go.
// source: hall_playback.proto
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

// Ignoring public import of CardInfo from common_mj.proto

// Ignoring public import of RoomTypeInfo from common_mj.proto

// Ignoring public import of ComposeCard from common_mj.proto

// Ignoring public import of PlayerCard from common_mj.proto

// Ignoring public import of ErrorCode from common_mj.proto

// Ignoring public import of mj_enum_color from common_mj.proto

// Ignoring public import of mj_enum_gangType from common_mj.proto

// Ignoring public import of mj_enum_huType from common_mj.proto

// Ignoring public import of mj_enum_composeCardType from common_mj.proto

// Ignoring public import of mj_enum_paiType from common_mj.proto

// Ignoring public import of mj_enum_userGameStatus from common_mj.proto

// Ignoring public import of mj_enum_deskGameStatus from common_mj.proto

// Ignoring public import of MJRoomType from common_mj.proto

// Start： 201704142049 下面是麻将回放的 Proto
// 当前页显示记录：第（length * page）条 到 第（length * (page + 1) - 1）条
type PlaybackReqPage struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Length           *int32       `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	Page             *int32       `protobuf:"varint,3,opt,name=page" json:"page,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PlaybackReqPage) Reset()                    { *m = PlaybackReqPage{} }
func (m *PlaybackReqPage) String() string            { return proto.CompactTextString(m) }
func (*PlaybackReqPage) ProtoMessage()               {}
func (*PlaybackReqPage) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{0} }

func (m *PlaybackReqPage) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PlaybackReqPage) GetLength() int32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func (m *PlaybackReqPage) GetPage() int32 {
	if m != nil && m.Page != nil {
		return *m.Page
	}
	return 0
}

type PlaybackAckPage struct {
	Header *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// 分页信息
	Length *int32 `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	Page   *int32 `protobuf:"varint,3,opt,name=page" json:"page,omitempty"`
	Total  *int32 `protobuf:"varint,4,opt,name=total" json:"total,omitempty"`
	// 快照信息
	PlaybackSnapshots []*PlaybackSnapshot `protobuf:"bytes,5,rep,name=playbackSnapshots" json:"playbackSnapshots,omitempty"`
	XXX_unrecognized  []byte              `json:"-"`
}

func (m *PlaybackAckPage) Reset()                    { *m = PlaybackAckPage{} }
func (m *PlaybackAckPage) String() string            { return proto.CompactTextString(m) }
func (*PlaybackAckPage) ProtoMessage()               {}
func (*PlaybackAckPage) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{1} }

func (m *PlaybackAckPage) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PlaybackAckPage) GetLength() int32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func (m *PlaybackAckPage) GetPage() int32 {
	if m != nil && m.Page != nil {
		return *m.Page
	}
	return 0
}

func (m *PlaybackAckPage) GetTotal() int32 {
	if m != nil && m.Total != nil {
		return *m.Total
	}
	return 0
}

func (m *PlaybackAckPage) GetPlaybackSnapshots() []*PlaybackSnapshot {
	if m != nil {
		return m.PlaybackSnapshots
	}
	return nil
}

// 玩家的信息
type PlayerInfo struct {
	IsBanker         *bool       `protobuf:"varint,1,opt,name=isBanker" json:"isBanker,omitempty"`
	PlayerCard       *PlayerCard `protobuf:"bytes,2,opt,name=playerCard" json:"playerCard,omitempty"`
	Coin             *int64      `protobuf:"varint,3,opt,name=coin" json:"coin,omitempty"`
	NickName         *string     `protobuf:"bytes,4,opt,name=nickName" json:"nickName,omitempty"`
	Sex              *int32      `protobuf:"varint,5,opt,name=sex" json:"sex,omitempty"`
	UserId           *uint32     `protobuf:"varint,6,opt,name=userId" json:"userId,omitempty"`
	IsOwner          *bool       `protobuf:"varint,7,opt,name=isOwner" json:"isOwner,omitempty"`
	BReady           *int32      `protobuf:"varint,8,opt,name=bReady" json:"bReady,omitempty"`
	BDingQue         *int32      `protobuf:"varint,9,opt,name=bDingQue" json:"bDingQue,omitempty"`
	BExchanged       *int32      `protobuf:"varint,10,opt,name=bExchanged" json:"bExchanged,omitempty"`
	NHuPai           *int32      `protobuf:"varint,11,opt,name=nHuPai" json:"nHuPai,omitempty"`
	QuePai           *int32      `protobuf:"varint,12,opt,name=quePai" json:"quePai,omitempty"`
	WxInfo           *WeixinInfo `protobuf:"bytes,13,opt,name=wxInfo" json:"wxInfo,omitempty"`
	GameStatus       *int32      `protobuf:"varint,14,opt,name=GameStatus" json:"GameStatus,omitempty"`
	AgentMode        *bool       `protobuf:"varint,15,opt,name=agentMode" json:"agentMode,omitempty"`
	Ip               *string     `protobuf:"bytes,16,opt,name=ip" json:"ip,omitempty"`
	XiaCount         *int32      `protobuf:"varint,17,opt,name=xiaCount" json:"xiaCount,omitempty"`
	IsBaoTing        *bool       `protobuf:"varint,18,opt,name=isBaoTing" json:"isBaoTing,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *PlayerInfo) Reset()                    { *m = PlayerInfo{} }
func (m *PlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*PlayerInfo) ProtoMessage()               {}
func (*PlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{2} }

func (m *PlayerInfo) GetIsBanker() bool {
	if m != nil && m.IsBanker != nil {
		return *m.IsBanker
	}
	return false
}

func (m *PlayerInfo) GetPlayerCard() *PlayerCard {
	if m != nil {
		return m.PlayerCard
	}
	return nil
}

func (m *PlayerInfo) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *PlayerInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *PlayerInfo) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *PlayerInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PlayerInfo) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *PlayerInfo) GetBReady() int32 {
	if m != nil && m.BReady != nil {
		return *m.BReady
	}
	return 0
}

func (m *PlayerInfo) GetBDingQue() int32 {
	if m != nil && m.BDingQue != nil {
		return *m.BDingQue
	}
	return 0
}

func (m *PlayerInfo) GetBExchanged() int32 {
	if m != nil && m.BExchanged != nil {
		return *m.BExchanged
	}
	return 0
}

func (m *PlayerInfo) GetNHuPai() int32 {
	if m != nil && m.NHuPai != nil {
		return *m.NHuPai
	}
	return 0
}

func (m *PlayerInfo) GetQuePai() int32 {
	if m != nil && m.QuePai != nil {
		return *m.QuePai
	}
	return 0
}

func (m *PlayerInfo) GetWxInfo() *WeixinInfo {
	if m != nil {
		return m.WxInfo
	}
	return nil
}

func (m *PlayerInfo) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *PlayerInfo) GetAgentMode() bool {
	if m != nil && m.AgentMode != nil {
		return *m.AgentMode
	}
	return false
}

func (m *PlayerInfo) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *PlayerInfo) GetXiaCount() int32 {
	if m != nil && m.XiaCount != nil {
		return *m.XiaCount
	}
	return 0
}

func (m *PlayerInfo) GetIsBaoTing() bool {
	if m != nil && m.IsBaoTing != nil {
		return *m.IsBaoTing
	}
	return false
}

// 回放快照信息（包含玩家信息、桌面信息、房间信息、右上角时间信息等），另外包含协议信息（PID 和 二进制数据）
type PlaybackSnapshot struct {
	PlayerInfo       []*PlayerInfo     `protobuf:"bytes,1,rep,name=playerInfo" json:"playerInfo,omitempty"`
	PlayBackDeskInfo *PlaybackDeskInfo `protobuf:"bytes,2,opt,name=playBackDeskInfo" json:"playBackDeskInfo,omitempty"`
	RoomTypeInfo     *RoomTypeInfo     `protobuf:"bytes,3,opt,name=RoomTypeInfo" json:"RoomTypeInfo,omitempty"`
	Time             *string           `protobuf:"bytes,4,opt,name=time" json:"time,omitempty"`
	OverturnType     *int32            `protobuf:"varint,5,opt,name=overturnType" json:"overturnType,omitempty"`
	Userid           *uint32           `protobuf:"varint,6,opt,name=userid" json:"userid,omitempty"`
	UseridOut        *uint32           `protobuf:"varint,7,opt,name=useridOut" json:"useridOut,omitempty"`
	Cardinfo         []*CardInfo       `protobuf:"bytes,8,rep,name=cardinfo" json:"cardinfo,omitempty"`
	Humsg            *string           `protobuf:"bytes,9,opt,name=humsg" json:"humsg,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *PlaybackSnapshot) Reset()                    { *m = PlaybackSnapshot{} }
func (m *PlaybackSnapshot) String() string            { return proto.CompactTextString(m) }
func (*PlaybackSnapshot) ProtoMessage()               {}
func (*PlaybackSnapshot) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{3} }

func (m *PlaybackSnapshot) GetPlayerInfo() []*PlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *PlaybackSnapshot) GetPlayBackDeskInfo() *PlaybackDeskInfo {
	if m != nil {
		return m.PlayBackDeskInfo
	}
	return nil
}

func (m *PlaybackSnapshot) GetRoomTypeInfo() *RoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

func (m *PlaybackSnapshot) GetTime() string {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return ""
}

func (m *PlaybackSnapshot) GetOverturnType() int32 {
	if m != nil && m.OverturnType != nil {
		return *m.OverturnType
	}
	return 0
}

func (m *PlaybackSnapshot) GetUserid() uint32 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

func (m *PlaybackSnapshot) GetUseridOut() uint32 {
	if m != nil && m.UseridOut != nil {
		return *m.UseridOut
	}
	return 0
}

func (m *PlaybackSnapshot) GetCardinfo() []*CardInfo {
	if m != nil {
		return m.Cardinfo
	}
	return nil
}

func (m *PlaybackSnapshot) GetHumsg() string {
	if m != nil && m.Humsg != nil {
		return *m.Humsg
	}
	return ""
}

// 按霄霄需求的【回放桌面信息】
type PlaybackDeskInfo struct {
	GameStatus       *int32  `protobuf:"varint,1,opt,name=GameStatus" json:"GameStatus,omitempty"`
	PlayerNum        *int32  `protobuf:"varint,2,opt,name=playerNum" json:"playerNum,omitempty"`
	ActiveUserId     *uint32 `protobuf:"varint,3,opt,name=activeUserId" json:"activeUserId,omitempty"`
	DeskInfo         *string `protobuf:"bytes,4,opt,name=deskInfo" json:"deskInfo,omitempty"`
	RemainCards      *int32  `protobuf:"varint,5,opt,name=remainCards" json:"remainCards,omitempty"`
	Time             *string `protobuf:"bytes,6,opt,name=time" json:"time,omitempty"`
	RoomType         *int32  `protobuf:"varint,7,opt,name=roomType" json:"roomType,omitempty"`
	RoomNumber       *string `protobuf:"bytes,8,opt,name=roomNumber" json:"roomNumber,omitempty"`
	CurrPlayCount    *int32  `protobuf:"varint,9,opt,name=currPlayCount" json:"currPlayCount,omitempty"`
	TotalPlayCount   *int32  `protobuf:"varint,10,opt,name=totalPlayCount" json:"totalPlayCount,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PlaybackDeskInfo) Reset()                    { *m = PlaybackDeskInfo{} }
func (m *PlaybackDeskInfo) String() string            { return proto.CompactTextString(m) }
func (*PlaybackDeskInfo) ProtoMessage()               {}
func (*PlaybackDeskInfo) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{4} }

func (m *PlaybackDeskInfo) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *PlaybackDeskInfo) GetPlayerNum() int32 {
	if m != nil && m.PlayerNum != nil {
		return *m.PlayerNum
	}
	return 0
}

func (m *PlaybackDeskInfo) GetActiveUserId() uint32 {
	if m != nil && m.ActiveUserId != nil {
		return *m.ActiveUserId
	}
	return 0
}

func (m *PlaybackDeskInfo) GetDeskInfo() string {
	if m != nil && m.DeskInfo != nil {
		return *m.DeskInfo
	}
	return ""
}

func (m *PlaybackDeskInfo) GetRemainCards() int32 {
	if m != nil && m.RemainCards != nil {
		return *m.RemainCards
	}
	return 0
}

func (m *PlaybackDeskInfo) GetTime() string {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return ""
}

func (m *PlaybackDeskInfo) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *PlaybackDeskInfo) GetRoomNumber() string {
	if m != nil && m.RoomNumber != nil {
		return *m.RoomNumber
	}
	return ""
}

func (m *PlaybackDeskInfo) GetCurrPlayCount() int32 {
	if m != nil && m.CurrPlayCount != nil {
		return *m.CurrPlayCount
	}
	return 0
}

func (m *PlaybackDeskInfo) GetTotalPlayCount() int32 {
	if m != nil && m.TotalPlayCount != nil {
		return *m.TotalPlayCount
	}
	return 0
}

// 一项回放记录
type PlaybackItem struct {
	Id *uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// optional common_enum_game gameId = 2; // 回放记录 游戏类型，显示用
	GameId           *uint32 `protobuf:"varint,2,opt,name=gameId" json:"gameId,omitempty"`
	Time             *string `protobuf:"bytes,3,opt,name=time" json:"time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PlaybackItem) Reset()                    { *m = PlaybackItem{} }
func (m *PlaybackItem) String() string            { return proto.CompactTextString(m) }
func (*PlaybackItem) ProtoMessage()               {}
func (*PlaybackItem) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{5} }

func (m *PlaybackItem) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *PlaybackItem) GetGameId() uint32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *PlaybackItem) GetTime() string {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return ""
}

func init() {
	proto.RegisterType((*PlaybackReqPage)(nil), "ddproto.playback_req_page")
	proto.RegisterType((*PlaybackAckPage)(nil), "ddproto.playback_ack_page")
	proto.RegisterType((*PlayerInfo)(nil), "ddproto.PlayerInfo")
	proto.RegisterType((*PlaybackSnapshot)(nil), "ddproto.PlaybackSnapshot")
	proto.RegisterType((*PlaybackDeskInfo)(nil), "ddproto.PlaybackDeskInfo")
	proto.RegisterType((*PlaybackItem)(nil), "ddproto.PlaybackItem")
}

var fileDescriptor23 = []byte{
	// 642 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x53, 0xcd, 0x4e, 0xdb, 0x4e,
	0x10, 0xff, 0x3b, 0x81, 0x90, 0x4c, 0x62, 0x48, 0x36, 0xf0, 0xd7, 0x96, 0x13, 0x4a, 0x2b, 0xb5,
	0x52, 0x25, 0x0e, 0xb4, 0x87, 0x9e, 0x81, 0xaa, 0x70, 0x28, 0xa4, 0x40, 0xc5, 0x31, 0xda, 0xd8,
	0xdb, 0x78, 0x1b, 0x7b, 0xd7, 0xd8, 0x6b, 0x08, 0x6f, 0xd2, 0x43, 0x5f, 0xac, 0x6f, 0xd0, 0xc7,
	0xe8, 0xec, 0xd8, 0x4e, 0x20, 0xbd, 0xf6, 0x00, 0xca, 0xcc, 0xce, 0xc7, 0xef, 0x63, 0x0c, 0xc3,
	0x48, 0xc4, 0xf1, 0x24, 0x8d, 0xc5, 0xe3, 0x54, 0x04, 0xf3, 0xc3, 0x34, 0x33, 0xd6, 0xb0, 0xad,
	0x30, 0xa4, 0x1f, 0xfb, 0xc3, 0xc0, 0x24, 0x89, 0xd1, 0x93, 0x20, 0x56, 0x52, 0xdb, 0xf2, 0x75,
	0x7f, 0xa7, 0x4a, 0x26, 0xdf, 0xcb, 0xc4, 0xe8, 0x16, 0x06, 0xf5, 0x80, 0x49, 0x26, 0xef, 0x26,
	0xa9, 0x98, 0x49, 0xf6, 0x0a, 0x5a, 0x91, 0x14, 0xa1, 0xcc, 0xb8, 0x77, 0xe0, 0xbd, 0xe9, 0x1e,
	0xed, 0x1e, 0x56, 0x43, 0x0f, 0xc7, 0xee, 0xff, 0x19, 0xbd, 0xb1, 0x6d, 0x68, 0xc5, 0x52, 0xcf,
	0x6c, 0xc4, 0x1b, 0x58, 0xb5, 0xc9, 0x7a, 0xb0, 0xe1, 0xba, 0x79, 0xd3, 0x45, 0xa3, 0x9f, 0xde,
	0x93, 0xc9, 0xee, 0xef, 0x5f, 0x4d, 0x66, 0x3e, 0x6c, 0x5a, 0x63, 0x45, 0xcc, 0x37, 0x28, 0x7c,
	0xbf, 0xda, 0x73, 0xad, 0x45, 0x9a, 0x47, 0xc6, 0xe6, 0x7c, 0xf3, 0xa0, 0x89, 0xd3, 0x5f, 0xac,
	0xa6, 0xaf, 0x55, 0x8c, 0x7e, 0x37, 0x00, 0x5c, 0x52, 0x66, 0xe7, 0xfa, 0x9b, 0x61, 0x7d, 0x68,
	0xab, 0xfc, 0x58, 0xe8, 0x79, 0x85, 0xac, 0xcd, 0x5e, 0x03, 0xa4, 0xf4, 0x7e, 0x22, 0xb2, 0x90,
	0x70, 0x74, 0x8f, 0x86, 0xcf, 0xe6, 0x95, 0x4f, 0x0e, 0x5c, 0x60, 0x94, 0x26, 0x70, 0x4d, 0x37,
	0x48, 0xab, 0x60, 0x7e, 0x21, 0x12, 0x49, 0xf8, 0x3a, 0xac, 0x0b, 0xcd, 0x5c, 0x2e, 0x10, 0x91,
	0x03, 0x8b, 0xcc, 0x8a, 0x1c, 0x77, 0x86, 0xbc, 0x85, 0xb1, 0xcf, 0x76, 0x60, 0x4b, 0xe5, 0x97,
	0x0f, 0x1a, 0xd7, 0x6e, 0xd1, 0x5a, 0x2c, 0x98, 0x5e, 0xa1, 0x0a, 0x8f, 0xbc, 0x4d, 0x0d, 0x38,
	0x6f, 0x7a, 0xaa, 0xf4, 0xec, 0x4b, 0x21, 0x79, 0x87, 0x32, 0x0c, 0x60, 0xfa, 0x71, 0x11, 0x44,
	0x42, 0xcf, 0x64, 0xc8, 0xa1, 0x1e, 0xab, 0xcf, 0x8a, 0xb1, 0x50, 0xbc, 0x5b, 0xc7, 0x77, 0x85,
	0x74, 0x71, 0x8f, 0xe2, 0x97, 0xd0, 0x7a, 0x58, 0x38, 0xa2, 0xdc, 0x5f, 0x23, 0x72, 0x2b, 0xd5,
	0x42, 0x69, 0xd2, 0x00, 0x07, 0x7f, 0x42, 0xd8, 0xd7, 0x56, 0xd8, 0x22, 0xe7, 0xdb, 0xd4, 0x38,
	0x80, 0x0e, 0x0a, 0xaf, 0xed, 0x67, 0x13, 0x4a, 0xbe, 0x43, 0x08, 0x01, 0x1a, 0x2a, 0xe5, 0x7d,
	0xe2, 0x86, 0xe8, 0x16, 0x4a, 0x9c, 0x98, 0x42, 0x5b, 0x3e, 0xa8, 0x1b, 0x9c, 0x90, 0xe6, 0x06,
	0x31, 0x73, 0xe6, 0x1a, 0x46, 0x3f, 0x1a, 0xd0, 0x5f, 0xd7, 0x7f, 0x25, 0x2f, 0xa1, 0xf2, 0xc8,
	0xae, 0x75, 0x79, 0x09, 0xd5, 0x3b, 0xe8, 0xbb, 0xc2, 0x63, 0x6c, 0x3e, 0x95, 0xf9, 0x9c, 0xca,
	0x4b, 0x37, 0xfe, 0x76, 0xb7, 0x2e, 0x60, 0x6f, 0xa1, 0x77, 0x65, 0x4c, 0x72, 0xf3, 0x98, 0x4a,
	0x6a, 0x68, 0x52, 0xc3, 0xde, 0xb2, 0xe1, 0xe9, 0xa3, 0x33, 0xd0, 0xaa, 0xa5, 0x5d, 0xbb, 0xd0,
	0x33, 0xf7, 0x32, 0xb3, 0x45, 0xa6, 0x5d, 0xc5, 0x73, 0xdf, 0x54, 0xed, 0x1b, 0xd2, 0x2c, 0xe3,
	0xcb, 0xc2, 0x92, 0x73, 0x3e, 0x6a, 0xdc, 0x0e, 0xf0, 0x1e, 0x94, 0xdb, 0xd7, 0x26, 0x3e, 0x83,
	0xe5, 0x3e, 0x77, 0x28, 0xb4, 0x0b, 0x6f, 0x37, 0x2a, 0x92, 0x7c, 0x46, 0x5e, 0x76, 0x46, 0xbf,
	0xbc, 0x95, 0x34, 0x4b, 0xf0, 0xcf, 0x7d, 0xf0, 0x6a, 0x59, 0x4b, 0xb9, 0x2e, 0x8a, 0xa4, 0xfa,
	0x28, 0x10, 0xa8, 0x08, 0xac, 0xba, 0x97, 0x5f, 0xcb, 0x83, 0x6a, 0x12, 0x0a, 0x74, 0x24, 0xac,
	0x65, 0x2a, 0x09, 0x0d, 0xa1, 0x9b, 0xc9, 0x44, 0x28, 0xed, 0x40, 0xe4, 0x15, 0x9f, 0x9a, 0x73,
	0xab, 0xb6, 0x31, 0xab, 0x14, 0x21, 0x32, 0x74, 0x64, 0x2e, 0x83, 0xdb, 0xa6, 0x78, 0x9a, 0x6d,
	0xaa, 0xda, 0x03, 0x3f, 0x28, 0xb2, 0xcc, 0xe1, 0x2d, 0x1d, 0x2f, 0xef, 0xf1, 0x7f, 0xd8, 0xa6,
	0xcf, 0x71, 0x95, 0xa7, 0x9b, 0x1c, 0x7d, 0x80, 0x5e, 0x4d, 0xed, 0xdc, 0xca, 0x84, 0xee, 0x26,
	0x24, 0x3a, 0xbe, 0x93, 0x73, 0x86, 0x14, 0xcf, 0xcb, 0x0f, 0xcb, 0x5f, 0xc2, 0x71, 0x1c, 0x3a,
	0xe3, 0xff, 0xc6, 0xde, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x35, 0x86, 0x7e, 0x0c, 0xdc, 0x04,
	0x00, 0x00,
}
