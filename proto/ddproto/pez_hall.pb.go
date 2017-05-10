// Code generated by protoc-gen-go.
// source: pez_hall.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of pez_base_PaiInfo from pez_base.proto

// Ignoring public import of pez_base_PlayConf from pez_base.proto

// Ignoring public import of pez_base_RoomTypeInfo from pez_base.proto

// Ignoring public import of pez_base_timerInfo from pez_base.proto

// Ignoring public import of pez_base_PaiValue from pez_base.proto

// Ignoring public import of pez_base_PlayerCard from pez_base.proto

// Ignoring public import of pez_lastScore from pez_base.proto

// Ignoring public import of pez_base_PlayerInfo from pez_base.proto

// Ignoring public import of pez_base_DeskGameInfo from pez_base.proto

// Ignoring public import of pez_tuozi from pez_base.proto

// Ignoring public import of pez_enum_protoId from pez_base.proto

// Ignoring public import of pez_enum_PEZTYPE from pez_base.proto

// Ignoring public import of pez_enum_ErrorCode from pez_base.proto

// Ignoring public import of pez_enum_Option from pez_base.proto

// Ignoring public import of pez_RoomType from pez_base.proto

// Ignoring public import of pez_enum_mjColor from pez_base.proto

// Ignoring public import of pez_enum_PaiType from pez_base.proto

// Ignoring public import of pez_enum_Base from pez_base.proto

// Ignoring public import of pez_enum_Bet from pez_base.proto

// Ignoring public import of pez_enum_UserGameStatus from pez_base.proto

// Ignoring public import of pez_enum_DeskGameStatus from pez_base.proto

// Ignoring public import of pez_enum_TuoziType from pez_base.proto

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

// 拼二张创建房间
type PezReq_CreateRoom struct {
	Header           *ProtoHeader          `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	RoomTypeInfo     *PezBase_RoomTypeInfo `protobuf:"bytes,2,opt,name=roomTypeInfo" json:"roomTypeInfo,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *PezReq_CreateRoom) Reset()                    { *m = PezReq_CreateRoom{} }
func (m *PezReq_CreateRoom) String() string            { return proto.CompactTextString(m) }
func (*PezReq_CreateRoom) ProtoMessage()               {}
func (*PezReq_CreateRoom) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{0} }

func (m *PezReq_CreateRoom) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PezReq_CreateRoom) GetRoomTypeInfo() *PezBase_RoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

// 拼二张创建房间回复的信息
type PezReq_AckCreateRoom struct {
	Header           *ProtoHeader          `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	DeskId           *int32                `protobuf:"varint,2,opt,name=deskId" json:"deskId,omitempty"`
	Password         *string               `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	UserBalance      *int64                `protobuf:"varint,4,opt,name=userBalance" json:"userBalance,omitempty"`
	CreateFee        *int64                `protobuf:"varint,5,opt,name=createFee" json:"createFee,omitempty"`
	RoomTypeInfo     *PezBase_RoomTypeInfo `protobuf:"bytes,6,opt,name=roomTypeInfo" json:"roomTypeInfo,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *PezReq_AckCreateRoom) Reset()                    { *m = PezReq_AckCreateRoom{} }
func (m *PezReq_AckCreateRoom) String() string            { return proto.CompactTextString(m) }
func (*PezReq_AckCreateRoom) ProtoMessage()               {}
func (*PezReq_AckCreateRoom) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{1} }

func (m *PezReq_AckCreateRoom) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PezReq_AckCreateRoom) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *PezReq_AckCreateRoom) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *PezReq_AckCreateRoom) GetUserBalance() int64 {
	if m != nil && m.UserBalance != nil {
		return *m.UserBalance
	}
	return 0
}

func (m *PezReq_AckCreateRoom) GetCreateFee() int64 {
	if m != nil && m.CreateFee != nil {
		return *m.CreateFee
	}
	return 0
}

func (m *PezReq_AckCreateRoom) GetRoomTypeInfo() *PezBase_RoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

// 客户端请求进入 room, 服务器返回game_SendGameInfo
type PezReq_EnterRoom struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	MatchId          *int32       `protobuf:"varint,2,opt,name=matchId" json:"matchId,omitempty"`
	TableId          *int32       `protobuf:"varint,3,opt,name=tableId" json:"tableId,omitempty"`
	PassWord         *string      `protobuf:"bytes,4,opt,name=PassWord" json:"PassWord,omitempty"`
	UserId           *uint32      `protobuf:"varint,5,opt,name=userId" json:"userId,omitempty"`
	RoomType         *int32       `protobuf:"varint,6,opt,name=roomType" json:"roomType,omitempty"`
	EnterType        *int32       `protobuf:"varint,7,opt,name=enterType" json:"enterType,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PezReq_EnterRoom) Reset()                    { *m = PezReq_EnterRoom{} }
func (m *PezReq_EnterRoom) String() string            { return proto.CompactTextString(m) }
func (*PezReq_EnterRoom) ProtoMessage()               {}
func (*PezReq_EnterRoom) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{2} }

func (m *PezReq_EnterRoom) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PezReq_EnterRoom) GetMatchId() int32 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

func (m *PezReq_EnterRoom) GetTableId() int32 {
	if m != nil && m.TableId != nil {
		return *m.TableId
	}
	return 0
}

func (m *PezReq_EnterRoom) GetPassWord() string {
	if m != nil && m.PassWord != nil {
		return *m.PassWord
	}
	return ""
}

func (m *PezReq_EnterRoom) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PezReq_EnterRoom) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *PezReq_EnterRoom) GetEnterType() int32 {
	if m != nil && m.EnterType != nil {
		return *m.EnterType
	}
	return 0
}

type PezReq_AckEnterRoom struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PezReq_AckEnterRoom) Reset()                    { *m = PezReq_AckEnterRoom{} }
func (m *PezReq_AckEnterRoom) String() string            { return proto.CompactTextString(m) }
func (*PezReq_AckEnterRoom) ProtoMessage()               {}
func (*PezReq_AckEnterRoom) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{3} }

func (m *PezReq_AckEnterRoom) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 公告消息
type Pez_Notice struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	NoticeType       *int32       `protobuf:"varint,2,opt,name=noticeType" json:"noticeType,omitempty"`
	ChannelId        *string      `protobuf:"bytes,3,opt,name=channelId" json:"channelId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Pez_Notice) Reset()                    { *m = Pez_Notice{} }
func (m *Pez_Notice) String() string            { return proto.CompactTextString(m) }
func (*Pez_Notice) ProtoMessage()               {}
func (*Pez_Notice) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{4} }

func (m *Pez_Notice) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_Notice) GetNoticeType() int32 {
	if m != nil && m.NoticeType != nil {
		return *m.NoticeType
	}
	return 0
}

func (m *Pez_Notice) GetChannelId() string {
	if m != nil && m.ChannelId != nil {
		return *m.ChannelId
	}
	return ""
}

// 公告的内容
type Pez_AckNotice struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	NoticeType       *int32       `protobuf:"varint,2,opt,name=noticeType" json:"noticeType,omitempty"`
	NoticeTitle      *string      `protobuf:"bytes,3,opt,name=noticeTitle" json:"noticeTitle,omitempty"`
	NoticeContent    *string      `protobuf:"bytes,4,opt,name=noticeContent" json:"noticeContent,omitempty"`
	NoticeMemo       *string      `protobuf:"bytes,5,opt,name=noticeMemo" json:"noticeMemo,omitempty"`
	Id               *int32       `protobuf:"varint,6,opt,name=id" json:"id,omitempty"`
	Fileds           []string     `protobuf:"bytes,7,rep,name=fileds" json:"fileds,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Pez_AckNotice) Reset()                    { *m = Pez_AckNotice{} }
func (m *Pez_AckNotice) String() string            { return proto.CompactTextString(m) }
func (*Pez_AckNotice) ProtoMessage()               {}
func (*Pez_AckNotice) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{5} }

func (m *Pez_AckNotice) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_AckNotice) GetNoticeType() int32 {
	if m != nil && m.NoticeType != nil {
		return *m.NoticeType
	}
	return 0
}

func (m *Pez_AckNotice) GetNoticeTitle() string {
	if m != nil && m.NoticeTitle != nil {
		return *m.NoticeTitle
	}
	return ""
}

func (m *Pez_AckNotice) GetNoticeContent() string {
	if m != nil && m.NoticeContent != nil {
		return *m.NoticeContent
	}
	return ""
}

func (m *Pez_AckNotice) GetNoticeMemo() string {
	if m != nil && m.NoticeMemo != nil {
		return *m.NoticeMemo
	}
	return ""
}

func (m *Pez_AckNotice) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Pez_AckNotice) GetFileds() []string {
	if m != nil {
		return m.Fileds
	}
	return nil
}

// 游戏战绩
type Pez_GameRecord struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Id               *int32       `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	GameId           *int32       `protobuf:"varint,3,opt,name=gameId" json:"gameId,omitempty"`
	UserId           *uint32      `protobuf:"varint,4,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Pez_GameRecord) Reset()                    { *m = Pez_GameRecord{} }
func (m *Pez_GameRecord) String() string            { return proto.CompactTextString(m) }
func (*Pez_GameRecord) ProtoMessage()               {}
func (*Pez_GameRecord) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{6} }

func (m *Pez_GameRecord) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_GameRecord) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Pez_GameRecord) GetGameId() int32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *Pez_GameRecord) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

//
type PezBeanUserRecord struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=UserId" json:"UserId,omitempty"`
	NickName         *string      `protobuf:"bytes,3,opt,name=NickName" json:"NickName,omitempty"`
	WinAmount        *int64       `protobuf:"varint,4,opt,name=WinAmount" json:"WinAmount,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PezBeanUserRecord) Reset()                    { *m = PezBeanUserRecord{} }
func (m *PezBeanUserRecord) String() string            { return proto.CompactTextString(m) }
func (*PezBeanUserRecord) ProtoMessage()               {}
func (*PezBeanUserRecord) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{7} }

func (m *PezBeanUserRecord) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PezBeanUserRecord) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PezBeanUserRecord) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *PezBeanUserRecord) GetWinAmount() int64 {
	if m != nil && m.WinAmount != nil {
		return *m.WinAmount
	}
	return 0
}

//
type PezBeanGameRecord struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Id               *int32               `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	DeskId           *int32               `protobuf:"varint,3,opt,name=deskId" json:"deskId,omitempty"`
	BeginTime        *string              `protobuf:"bytes,4,opt,name=beginTime" json:"beginTime,omitempty"`
	Users            []*PezBeanUserRecord `protobuf:"bytes,5,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PezBeanGameRecord) Reset()                    { *m = PezBeanGameRecord{} }
func (m *PezBeanGameRecord) String() string            { return proto.CompactTextString(m) }
func (*PezBeanGameRecord) ProtoMessage()               {}
func (*PezBeanGameRecord) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{8} }

func (m *PezBeanGameRecord) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PezBeanGameRecord) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *PezBeanGameRecord) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *PezBeanGameRecord) GetBeginTime() string {
	if m != nil && m.BeginTime != nil {
		return *m.BeginTime
	}
	return ""
}

func (m *PezBeanGameRecord) GetUsers() []*PezBeanUserRecord {
	if m != nil {
		return m.Users
	}
	return nil
}

//
type Pez_AckGameRecord struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32              `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	GameId           *int32               `protobuf:"varint,3,opt,name=gameId" json:"gameId,omitempty"`
	Records          []*PezBeanGameRecord `protobuf:"bytes,4,rep,name=records" json:"records,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *Pez_AckGameRecord) Reset()                    { *m = Pez_AckGameRecord{} }
func (m *Pez_AckGameRecord) String() string            { return proto.CompactTextString(m) }
func (*Pez_AckGameRecord) ProtoMessage()               {}
func (*Pez_AckGameRecord) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{9} }

func (m *Pez_AckGameRecord) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_AckGameRecord) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Pez_AckGameRecord) GetGameId() int32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *Pez_AckGameRecord) GetRecords() []*PezBeanGameRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

// 反馈信息的协议
type Pez_Feedback struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Message          *string      `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Pez_Feedback) Reset()                    { *m = Pez_Feedback{} }
func (m *Pez_Feedback) String() string            { return proto.CompactTextString(m) }
func (*Pez_Feedback) ProtoMessage()               {}
func (*Pez_Feedback) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{10} }

func (m *Pez_Feedback) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_Feedback) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*PezReq_CreateRoom)(nil), "ddproto.pez_req_CreateRoom")
	proto.RegisterType((*PezReq_AckCreateRoom)(nil), "ddproto.pez_req_AckCreateRoom")
	proto.RegisterType((*PezReq_EnterRoom)(nil), "ddproto.pez_req_EnterRoom")
	proto.RegisterType((*PezReq_AckEnterRoom)(nil), "ddproto.pez_req_AckEnterRoom")
	proto.RegisterType((*Pez_Notice)(nil), "ddproto.pez_Notice")
	proto.RegisterType((*Pez_AckNotice)(nil), "ddproto.pez_AckNotice")
	proto.RegisterType((*Pez_GameRecord)(nil), "ddproto.pez_GameRecord")
	proto.RegisterType((*PezBeanUserRecord)(nil), "ddproto.pezBeanUserRecord")
	proto.RegisterType((*PezBeanGameRecord)(nil), "ddproto.pezBeanGameRecord")
	proto.RegisterType((*Pez_AckGameRecord)(nil), "ddproto.pez_AckGameRecord")
	proto.RegisterType((*Pez_Feedback)(nil), "ddproto.pez_Feedback")
}

var fileDescriptor39 = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x94, 0xc1, 0x6e, 0xd3, 0x4e,
	0x10, 0xc6, 0xff, 0xb6, 0x9b, 0xe4, 0xef, 0x69, 0x53, 0xa8, 0xdb, 0x4a, 0x56, 0x0e, 0x08, 0x59,
	0x1c, 0x40, 0x48, 0x39, 0x20, 0x8e, 0x5c, 0x9a, 0xaa, 0xd0, 0x1c, 0xa8, 0xa2, 0xa8, 0x55, 0xc5,
	0x29, 0x5a, 0xaf, 0xa7, 0x89, 0x15, 0x7b, 0xd7, 0x78, 0x5d, 0x21, 0x78, 0x00, 0x5e, 0x80, 0x37,
	0x40, 0xe2, 0x31, 0x78, 0x37, 0x66, 0xd7, 0x6b, 0x27, 0x54, 0x1c, 0x62, 0xf5, 0x52, 0x55, 0x9f,
	0x67, 0x67, 0x7e, 0xdf, 0x37, 0xbb, 0x81, 0xc3, 0x02, 0xbf, 0x2d, 0x56, 0x2c, 0xcb, 0xc6, 0x45,
	0x29, 0x2b, 0x19, 0x0c, 0x92, 0xc4, 0xfc, 0x33, 0x32, 0x1f, 0x62, 0xa6, 0xb0, 0xfe, 0x30, 0x3a,
	0xe6, 0x32, 0xcf, 0xa5, 0x58, 0xf0, 0x2c, 0x45, 0x51, 0xd5, 0x62, 0x54, 0x40, 0xa0, 0xcb, 0x4a,
	0xfc, 0xbc, 0x38, 0x2f, 0x91, 0x55, 0x38, 0x97, 0x32, 0x0f, 0x5e, 0x40, 0x7f, 0x85, 0x2c, 0xc1,
	0x32, 0x74, 0x9e, 0x3b, 0x2f, 0xf7, 0xdf, 0x9c, 0x8c, 0x6d, 0xd3, 0xf1, 0x4c, 0xff, 0xbd, 0x34,
	0xdf, 0x82, 0xb7, 0x70, 0x50, 0x52, 0xf5, 0xf5, 0xd7, 0x02, 0xa7, 0xe2, 0x4e, 0x86, 0xae, 0xa9,
	0x7d, 0xd6, 0xd6, 0x36, 0xf3, 0x17, 0xf3, 0xad, 0xaa, 0xe8, 0xb7, 0x03, 0xa7, 0xcd, 0xc8, 0x33,
	0xbe, 0xee, 0x3c, 0xf5, 0x10, 0xfa, 0x09, 0xaa, 0xf5, 0x34, 0x31, 0xf3, 0x7a, 0xc1, 0x53, 0xf8,
	0xbf, 0x60, 0x4a, 0x7d, 0x91, 0x65, 0x12, 0x7a, 0xa4, 0xf8, 0xc1, 0x31, 0xec, 0xdf, 0x2b, 0x2c,
	0x27, 0x2c, 0x63, 0x82, 0x63, 0xb8, 0x47, 0xa2, 0x17, 0x1c, 0x81, 0xcf, 0xcd, 0xa8, 0xf7, 0x88,
	0x61, 0xcf, 0x48, 0x0f, 0xf9, 0xfb, 0x3b, 0xf1, 0xff, 0x74, 0xe0, 0xa8, 0xe1, 0xbf, 0x10, 0x15,
	0x96, 0x1d, 0xd8, 0x9f, 0xc0, 0x20, 0x67, 0x15, 0x5f, 0xb5, 0xf0, 0x24, 0x54, 0x2c, 0xce, 0x70,
	0x5a, 0xb3, 0x1b, 0x37, 0x33, 0x72, 0x73, 0xab, 0xdd, 0xec, 0x19, 0x37, 0xe4, 0x57, 0xbb, 0xa1,
	0x0a, 0x4d, 0x3d, 0xd4, 0x15, 0x0d, 0xb5, 0x21, 0xee, 0x69, 0x6b, 0xa8, 0x41, 0x8c, 0x34, 0xd0,
	0x52, 0xf4, 0x0e, 0x4e, 0xb6, 0x32, 0xee, 0x88, 0x19, 0x7d, 0x02, 0xd0, 0xa7, 0xaf, 0x64, 0x95,
	0x72, 0xdc, 0xd1, 0x5a, 0x00, 0x20, 0x4c, 0xbd, 0xa1, 0x70, 0x1b, 0x30, 0xbe, 0x62, 0x42, 0x60,
	0x66, 0xfd, 0xf9, 0xd1, 0x2f, 0x07, 0x86, 0xba, 0x37, 0x51, 0x3d, 0xba, 0x3d, 0xed, 0xd9, 0x6a,
	0x69, 0x95, 0xa1, 0x5d, 0xfe, 0x29, 0x0c, 0x6b, 0xf1, 0x5c, 0x92, 0x6b, 0x51, 0xd9, 0x14, 0xdb,
	0xf3, 0x1f, 0x31, 0x97, 0x26, 0x49, 0x9f, 0x24, 0x37, 0x4d, 0x6c, 0x86, 0x94, 0xf2, 0x5d, 0x9a,
	0x61, 0xa2, 0x28, 0x40, 0x8f, 0x38, 0xe3, 0xfa, 0x5d, 0x7d, 0x60, 0x39, 0xce, 0x91, 0xd3, 0x36,
	0x76, 0xe4, 0xac, 0x7b, 0xba, 0x4d, 0xcf, 0x25, 0x9d, 0x6f, 0x77, 0xbb, 0xd9, 0xa4, 0x66, 0x1a,
	0x46, 0xc2, 0x5c, 0xa4, 0x09, 0x32, 0x71, 0x43, 0x72, 0xa7, 0x31, 0xd4, 0xea, 0xa6, 0x6e, 0xe5,
	0x36, 0x97, 0xe2, 0x2a, 0xa5, 0x48, 0x69, 0x9c, 0xcd, 0x81, 0xb2, 0xbf, 0x4d, 0xc5, 0x59, 0x2e,
	0xef, 0x6d, 0x06, 0x5e, 0xf4, 0xc3, 0x69, 0x07, 0x3e, 0xd6, 0x97, 0x7d, 0x81, 0x5e, 0xb3, 0xe6,
	0x18, 0x97, 0xa9, 0xb8, 0x4e, 0x73, 0xb4, 0x71, 0xbf, 0x82, 0x9e, 0xb6, 0xaa, 0x28, 0x69, 0x8f,
	0x7a, 0x8e, 0xb6, 0xdf, 0xd4, 0xdf, 0x86, 0xa3, 0xef, 0xf6, 0x3d, 0xd1, 0x8d, 0xe8, 0x4c, 0xb5,
	0x49, 0xb4, 0x8e, 0xe1, 0x61, 0xe2, 0xaf, 0x61, 0x50, 0x9a, 0x7e, 0x8a, 0xb8, 0xfe, 0x09, 0xb2,
	0x19, 0x19, 0x5d, 0xc0, 0x81, 0xe6, 0xa0, 0xdf, 0x87, 0x24, 0x66, 0x7c, 0xdd, 0xe1, 0x49, 0xa3,
	0x52, 0x6c, 0x59, 0xdf, 0x4a, 0x7f, 0xe2, 0x5e, 0x7a, 0xb3, 0xff, 0x66, 0xce, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x37, 0x26, 0x77, 0x89, 0x98, 0x05, 0x00, 0x00,
}
