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

// 麻将的操作类型
type PlaybackMjActType int32

const (
	PlaybackMjActType_PMAT_MO           PlaybackMjActType = 1
	PlaybackMjActType_PMAT_MO_OUT       PlaybackMjActType = 2
	PlaybackMjActType_PMAT_MO_PENG      PlaybackMjActType = 3
	PlaybackMjActType_PMAT_MO_GANG_BA   PlaybackMjActType = 4
	PlaybackMjActType_PMAT_MO_GANG_AN   PlaybackMjActType = 5
	PlaybackMjActType_PMAT_MO_GANG_DIAN PlaybackMjActType = 6
	PlaybackMjActType_PMAT_MO_GANG_CHI  PlaybackMjActType = 7
	PlaybackMjActType_PMAT_MO_GANG_HU   PlaybackMjActType = 8
)

var PlaybackMjActType_name = map[int32]string{
	1: "PMAT_MO",
	2: "PMAT_MO_OUT",
	3: "PMAT_MO_PENG",
	4: "PMAT_MO_GANG_BA",
	5: "PMAT_MO_GANG_AN",
	6: "PMAT_MO_GANG_DIAN",
	7: "PMAT_MO_GANG_CHI",
	8: "PMAT_MO_GANG_HU",
}
var PlaybackMjActType_value = map[string]int32{
	"PMAT_MO":           1,
	"PMAT_MO_OUT":       2,
	"PMAT_MO_PENG":      3,
	"PMAT_MO_GANG_BA":   4,
	"PMAT_MO_GANG_AN":   5,
	"PMAT_MO_GANG_DIAN": 6,
	"PMAT_MO_GANG_CHI":  7,
	"PMAT_MO_GANG_HU":   8,
}

func (x PlaybackMjActType) Enum() *PlaybackMjActType {
	p := new(PlaybackMjActType)
	*p = x
	return p
}
func (x PlaybackMjActType) String() string {
	return proto.EnumName(PlaybackMjActType_name, int32(x))
}
func (x *PlaybackMjActType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PlaybackMjActType_value, data, "PlaybackMjActType")
	if err != nil {
		return err
	}
	*x = PlaybackMjActType(value)
	return nil
}
func (PlaybackMjActType) EnumDescriptor() ([]byte, []int) { return fileDescriptor24, []int{0} }

// Start： 201704142049 下面是麻将回放的 Proto
// 当前页显示记录：第（length * page）条 到 第（length * (page + 1) - 1）条
type PlaybackReqPage struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Length           *int32       `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	Page             *int32       `protobuf:"varint,3,opt,name=page" json:"page,omitempty"`
	GameNumber       *int32       `protobuf:"varint,4,opt,name=gameNumber" json:"gameNumber,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PlaybackReqPage) Reset()                    { *m = PlaybackReqPage{} }
func (m *PlaybackReqPage) String() string            { return proto.CompactTextString(m) }
func (*PlaybackReqPage) ProtoMessage()               {}
func (*PlaybackReqPage) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{0} }

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

func (m *PlaybackReqPage) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
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
func (*PlaybackAckPage) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{1} }

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
func (*PlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{2} }

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
func (*PlaybackSnapshot) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{3} }

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
func (*PlaybackDeskInfo) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{4} }

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
func (*PlaybackItem) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{5} }

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
	proto.RegisterEnum("ddproto.PlaybackMjActType", PlaybackMjActType_name, PlaybackMjActType_value)
}

var fileDescriptor24 = []byte{
	// 739 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xda, 0x40,
	0x10, 0xae, 0x21, 0xfc, 0x0d, 0x10, 0x8c, 0x49, 0x2a, 0x37, 0xa7, 0x88, 0x56, 0x6a, 0xd5, 0x4a,
	0x39, 0xa4, 0x3d, 0xf4, 0x4a, 0x7e, 0x04, 0x1c, 0x02, 0x34, 0x21, 0xea, 0xd1, 0x5a, 0xec, 0x2d,
	0x6c, 0xc0, 0x6b, 0x62, 0xaf, 0x13, 0xf2, 0x26, 0x3d, 0xf4, 0x29, 0xfa, 0x36, 0x7d, 0x83, 0x3e,
	0x46, 0x67, 0xc7, 0x36, 0x04, 0x72, 0xed, 0x21, 0x11, 0x33, 0x3b, 0x3f, 0xdf, 0xf7, 0xcd, 0x8c,
	0xa1, 0x35, 0x63, 0x8b, 0x85, 0xb3, 0x5c, 0xb0, 0xa7, 0x09, 0x73, 0xe7, 0x27, 0xcb, 0x30, 0x50,
	0x81, 0x55, 0xf2, 0x3c, 0xfa, 0x71, 0xd4, 0x72, 0x03, 0xdf, 0x0f, 0xa4, 0xe3, 0x2e, 0x04, 0x97,
	0x2a, 0x79, 0x3d, 0x6a, 0xa4, 0x4e, 0xff, 0x2e, 0x71, 0xb4, 0xe7, 0xd0, 0xcc, 0x0a, 0x38, 0x21,
	0xbf, 0x77, 0x96, 0x6c, 0xca, 0xad, 0x77, 0x50, 0x9c, 0x71, 0xe6, 0xf1, 0xd0, 0x36, 0x8e, 0x8d,
	0x0f, 0xd5, 0xd3, 0x83, 0x93, 0xb4, 0xe8, 0xc9, 0x48, 0xff, 0xef, 0xd1, 0x9b, 0xb5, 0x0f, 0xc5,
	0x05, 0x97, 0x53, 0x35, 0xb3, 0x73, 0x18, 0x55, 0xb0, 0x6a, 0xb0, 0xa7, 0xb3, 0xed, 0x3c, 0x59,
	0x16, 0xc0, 0x94, 0xf9, 0x7c, 0x10, 0xfb, 0x13, 0xac, 0xb3, 0xa7, 0x7d, 0xed, 0x5f, 0xc6, 0xb3,
	0x6e, 0xfa, 0xef, 0xbf, 0x75, 0xab, 0x43, 0x41, 0x05, 0x8a, 0x2d, 0x92, 0x46, 0xd6, 0x97, 0x4d,
	0x9f, 0x1b, 0xc9, 0x96, 0xd1, 0x2c, 0x50, 0x91, 0x5d, 0x38, 0xce, 0x63, 0xf5, 0x37, 0x9b, 0xea,
	0x3b, 0x11, 0xed, 0xbf, 0x39, 0x00, 0xed, 0xe4, 0x61, 0x5f, 0xfe, 0x08, 0x2c, 0x13, 0xca, 0x22,
	0x3a, 0x63, 0x72, 0x9e, 0x22, 0x2b, 0x5b, 0xef, 0x01, 0x96, 0xf4, 0x7e, 0xce, 0x42, 0x8f, 0x70,
	0x54, 0x4f, 0x5b, 0x5b, 0xf5, 0x92, 0x27, 0x0d, 0xce, 0x0d, 0x84, 0x24, 0x70, 0x79, 0x5d, 0x48,
	0x0a, 0x77, 0x3e, 0x40, 0x39, 0x08, 0x5f, 0xc5, 0xaa, 0x42, 0x3e, 0xe2, 0x2b, 0x44, 0xa4, 0xc1,
	0x22, 0xb3, 0x38, 0xc2, 0x9e, 0x9e, 0x5d, 0x44, 0xbb, 0x6e, 0x35, 0xa0, 0x24, 0xa2, 0xe1, 0xa3,
	0xc4, 0xb6, 0x25, 0x6a, 0x8b, 0x01, 0x93, 0x6b, 0x54, 0xe1, 0xc9, 0x2e, 0x53, 0x02, 0xd6, 0x9b,
	0x5c, 0x08, 0x39, 0xfd, 0x16, 0x73, 0xbb, 0x92, 0x89, 0x3d, 0xb9, 0x5c, 0xb9, 0x33, 0x26, 0xa7,
	0xdc, 0xb3, 0x21, 0x2b, 0x2b, 0x7b, 0xf1, 0x88, 0x09, 0xbb, 0x9a, 0xd9, 0xf7, 0x31, 0xd7, 0x76,
	0x8d, 0xec, 0xb7, 0x50, 0x7c, 0x5c, 0x69, 0xa2, 0x76, 0x7d, 0x87, 0xc8, 0x77, 0x2e, 0x56, 0x42,
	0x92, 0x06, 0x58, 0xb8, 0x8b, 0xb0, 0x6f, 0x14, 0x53, 0x71, 0x64, 0xef, 0x53, 0x62, 0x13, 0x2a,
	0x28, 0xbc, 0x54, 0x57, 0x81, 0xc7, 0xed, 0x06, 0x21, 0x04, 0xc8, 0x89, 0xa5, 0x6d, 0x12, 0x37,
	0x44, 0xb7, 0x12, 0xec, 0x3c, 0x88, 0xa5, 0xb2, 0x9b, 0x59, 0x82, 0x16, 0x32, 0x18, 0x23, 0x66,
	0xdb, 0xd2, 0x09, 0xed, 0x9f, 0x39, 0x30, 0x77, 0xf5, 0xdf, 0xc8, 0x4b, 0xa8, 0x0c, 0x1a, 0xd7,
	0xae, 0xbc, 0x84, 0xea, 0x33, 0x98, 0x3a, 0xf0, 0x0c, 0x93, 0x2f, 0x78, 0x34, 0xa7, 0xf0, 0x64,
	0x1a, 0x2f, 0xa7, 0x9b, 0x05, 0x58, 0x9f, 0xa0, 0x76, 0x1d, 0x04, 0xfe, 0xf8, 0x69, 0xc9, 0x29,
	0x21, 0x4f, 0x09, 0x87, 0xeb, 0x84, 0xe7, 0x8f, 0x7a, 0x80, 0x4a, 0xac, 0xc7, 0x75, 0x00, 0xb5,
	0xe0, 0x81, 0x87, 0x2a, 0x0e, 0xa5, 0x8e, 0xd8, 0x9e, 0x9b, 0xc8, 0xe6, 0x86, 0x34, 0x13, 0x7b,
	0x18, 0x2b, 0x9a, 0x5c, 0x1d, 0x35, 0x2e, 0xbb, 0xb8, 0x0f, 0x42, 0xf7, 0x2b, 0x13, 0x9f, 0xe6,
	0xba, 0x9f, 0x5e, 0x14, 0xea, 0x85, 0xbb, 0x3b, 0x8b, 0xfd, 0x68, 0x4a, 0xb3, 0xac, 0xb4, 0xff,
	0x18, 0x1b, 0x69, 0xd6, 0xe0, 0xb7, 0xe7, 0x60, 0x64, 0xb2, 0x26, 0x72, 0xe1, 0x8d, 0xa5, 0x47,
	0x81, 0x40, 0x99, 0xab, 0xc4, 0x03, 0xbf, 0x4d, 0x16, 0x2a, 0x4f, 0x28, 0x70, 0x22, 0x5e, 0x26,
	0x53, 0x42, 0xa8, 0x05, 0xd5, 0x90, 0xfb, 0x4c, 0x48, 0x0d, 0x22, 0x4a, 0xf9, 0x64, 0x9c, 0x8b,
	0xd9, 0x18, 0xc3, 0x54, 0x11, 0x22, 0x43, 0x4b, 0xa6, 0x3d, 0xe9, 0x45, 0x97, 0x29, 0xea, 0x10,
	0xea, 0x6e, 0x1c, 0x86, 0x1a, 0x6f, 0x32, 0xf1, 0x64, 0x1f, 0x5f, 0xc3, 0x3e, 0x9d, 0xe3, 0xc6,
	0x4f, 0x3b, 0xd9, 0xfe, 0x0a, 0xb5, 0x8c, 0x5a, 0x5f, 0x71, 0x9f, 0xf6, 0xc6, 0x23, 0x3a, 0x75,
	0x2d, 0xa7, 0xfe, 0x60, 0xf4, 0x93, 0xc3, 0xaa, 0xaf, 0xe1, 0x68, 0x0e, 0x95, 0x8f, 0xbf, 0x0d,
	0x38, 0x58, 0x7f, 0x3a, 0xfc, 0x3b, 0xfc, 0x7a, 0x28, 0x47, 0x21, 0x36, 0x3c, 0xa5, 0xd2, 0xe8,
	0xaa, 0x33, 0x76, 0xae, 0x86, 0xa6, 0x81, 0xa7, 0x53, 0x4d, 0x0d, 0x67, 0x78, 0x3b, 0x36, 0x73,
	0xc8, 0xa2, 0x96, 0x39, 0x46, 0x97, 0x83, 0xae, 0x99, 0x47, 0xea, 0x8d, 0xcc, 0xd3, 0xed, 0x0c,
	0xba, 0xce, 0x59, 0xc7, 0xdc, 0x7b, 0xe1, 0xec, 0x0c, 0xcc, 0x02, 0x72, 0x6b, 0x6e, 0x39, 0x2f,
	0xfa, 0xe8, 0x2e, 0xa2, 0xc6, 0xe6, 0x96, 0xfb, 0xbc, 0xd7, 0x37, 0x4b, 0x2f, 0x2a, 0xf4, 0x6e,
	0xcd, 0xf2, 0xe8, 0xd5, 0xc8, 0xf8, 0x17, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x21, 0xf7, 0xbb, 0xa5,
	0x05, 0x00, 0x00,
}
