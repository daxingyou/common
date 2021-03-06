// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mahjong_hall.proto

package mjproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ProtoHeader from base.proto

// Ignoring public import of Heartbeat from base.proto

// Ignoring public import of WeixinInfo from base.proto

// Ignoring public import of CardInfo from base.proto

// Ignoring public import of PlayOptions from base.proto

// Ignoring public import of ChangShaPlayOptions from base.proto

// Ignoring public import of BaiShanPlayOptions from base.proto

// Ignoring public import of ZhuanZhuanPlayOptions from base.proto

// Ignoring public import of LiuZhouPlayOptions from base.proto

// Ignoring public import of HaiNanPlayOptions from base.proto

// Ignoring public import of ChangBaiPlayOptions from base.proto

// Ignoring public import of RoomTypeInfo from base.proto

// Ignoring public import of ComposeCard from base.proto

// Ignoring public import of PlayerCard from base.proto

// Ignoring public import of GangBean from base.proto

// Ignoring public import of PlayerInfo from base.proto

// Ignoring public import of DeskGameInfo from base.proto

// Ignoring public import of EProtoId from base.proto

// Ignoring public import of ErrorCode from base.proto

// Ignoring public import of MJOption from base.proto

// Ignoring public import of MJRoomType from base.proto

// Ignoring public import of MahjongColor from base.proto

// Ignoring public import of GangType from base.proto

// Ignoring public import of ComposeCardType from base.proto

// Ignoring public import of HuType from base.proto

// Ignoring public import of PaiType from base.proto

// Ignoring public import of COMMON_ENUM_ROOMCARD_BILL_TYPE from base.proto

// Ignoring public import of MJUserGameStatus from base.proto

// Ignoring public import of DeskGameStatus from base.proto

// Ignoring public import of CHANGBAI_PREOPENING_PLAYOPTIONS from base.proto

// 服务器信息
type ServerInfo struct {
	Ip               *string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Port             *int32  `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	Status           *int32  `protobuf:"varint,3,opt,name=status" json:"status,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ServerInfo) Reset()                    { *m = ServerInfo{} }
func (m *ServerInfo) String() string            { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()               {}
func (*ServerInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ServerInfo) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *ServerInfo) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *ServerInfo) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

// 接入服务器
type Game_QuickConn struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ChannelId        *string      `protobuf:"bytes,2,opt,name=channelId" json:"channelId,omitempty"`
	GameId           *int32       `protobuf:"varint,3,opt,name=gameId" json:"gameId,omitempty"`
	CurrVersion      *int32       `protobuf:"varint,4,opt,name=currVersion" json:"currVersion,omitempty"`
	LanguageId       *int32       `protobuf:"varint,5,opt,name=languageId" json:"languageId,omitempty"`
	UserId           *uint32      `protobuf:"varint,6,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_QuickConn) Reset()                    { *m = Game_QuickConn{} }
func (m *Game_QuickConn) String() string            { return proto.CompactTextString(m) }
func (*Game_QuickConn) ProtoMessage()               {}
func (*Game_QuickConn) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *Game_QuickConn) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_QuickConn) GetChannelId() string {
	if m != nil && m.ChannelId != nil {
		return *m.ChannelId
	}
	return ""
}

func (m *Game_QuickConn) GetGameId() int32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *Game_QuickConn) GetCurrVersion() int32 {
	if m != nil && m.CurrVersion != nil {
		return *m.CurrVersion
	}
	return 0
}

func (m *Game_QuickConn) GetLanguageId() int32 {
	if m != nil && m.LanguageId != nil {
		return *m.LanguageId
	}
	return 0
}

func (m *Game_QuickConn) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

type Game_AckQuickConn struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	GameServer       *ServerInfo  `protobuf:"bytes,2,opt,name=gameServer" json:"gameServer,omitempty"`
	ReleaseTag       *int32       `protobuf:"varint,3,opt,name=releaseTag" json:"releaseTag,omitempty"`
	CurrVersion      *int32       `protobuf:"varint,4,opt,name=currVersion" json:"currVersion,omitempty"`
	IsUpdate         *int32       `protobuf:"varint,5,opt,name=isUpdate" json:"isUpdate,omitempty"`
	DownloadUrl      *string      `protobuf:"bytes,6,opt,name=downloadUrl" json:"downloadUrl,omitempty"`
	VersionInfo      *string      `protobuf:"bytes,7,opt,name=versionInfo" json:"versionInfo,omitempty"`
	IsMaintain       *int32       `protobuf:"varint,8,opt,name=isMaintain" json:"isMaintain,omitempty"`
	MaintainMsg      *string      `protobuf:"bytes,9,opt,name=maintainMsg" json:"maintainMsg,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_AckQuickConn) Reset()                    { *m = Game_AckQuickConn{} }
func (m *Game_AckQuickConn) String() string            { return proto.CompactTextString(m) }
func (*Game_AckQuickConn) ProtoMessage()               {}
func (*Game_AckQuickConn) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *Game_AckQuickConn) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_AckQuickConn) GetGameServer() *ServerInfo {
	if m != nil {
		return m.GameServer
	}
	return nil
}

func (m *Game_AckQuickConn) GetReleaseTag() int32 {
	if m != nil && m.ReleaseTag != nil {
		return *m.ReleaseTag
	}
	return 0
}

func (m *Game_AckQuickConn) GetCurrVersion() int32 {
	if m != nil && m.CurrVersion != nil {
		return *m.CurrVersion
	}
	return 0
}

func (m *Game_AckQuickConn) GetIsUpdate() int32 {
	if m != nil && m.IsUpdate != nil {
		return *m.IsUpdate
	}
	return 0
}

func (m *Game_AckQuickConn) GetDownloadUrl() string {
	if m != nil && m.DownloadUrl != nil {
		return *m.DownloadUrl
	}
	return ""
}

func (m *Game_AckQuickConn) GetVersionInfo() string {
	if m != nil && m.VersionInfo != nil {
		return *m.VersionInfo
	}
	return ""
}

func (m *Game_AckQuickConn) GetIsMaintain() int32 {
	if m != nil && m.IsMaintain != nil {
		return *m.IsMaintain
	}
	return 0
}

func (m *Game_AckQuickConn) GetMaintainMsg() string {
	if m != nil && m.MaintainMsg != nil {
		return *m.MaintainMsg
	}
	return ""
}

// 游戏登录请求
type Game_Login struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	ProtoVersion     *int32       `protobuf:"varint,3,opt,name=protoVersion" json:"protoVersion,omitempty"`
	WxInfo           *WeixinInfo  `protobuf:"bytes,4,opt,name=wxInfo" json:"wxInfo,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_Login) Reset()                    { *m = Game_Login{} }
func (m *Game_Login) String() string            { return proto.CompactTextString(m) }
func (*Game_Login) ProtoMessage()               {}
func (*Game_Login) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *Game_Login) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_Login) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Game_Login) GetProtoVersion() int32 {
	if m != nil && m.ProtoVersion != nil {
		return *m.ProtoVersion
	}
	return 0
}

func (m *Game_Login) GetWxInfo() *WeixinInfo {
	if m != nil {
		return m.WxInfo
	}
	return nil
}

// 游戏登录回复
type Game_AckLogin struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	NickName         *string      `protobuf:"bytes,3,opt,name=nickName" json:"nickName,omitempty"`
	RoomPassword     *string      `protobuf:"bytes,4,opt,name=roomPassword" json:"roomPassword,omitempty"`
	CostCreateRoom   *int64       `protobuf:"varint,5,opt,name=costCreateRoom" json:"costCreateRoom,omitempty"`
	CostRebuy        *int64       `protobuf:"varint,6,opt,name=costRebuy" json:"costRebuy,omitempty"`
	Championship     *bool        `protobuf:"varint,7,opt,name=championship" json:"championship,omitempty"`
	Chip             *int64       `protobuf:"varint,8,opt,name=chip" json:"chip,omitempty"`
	MailCount        *int32       `protobuf:"varint,9,opt,name=mailCount" json:"mailCount,omitempty"`
	Notice           *string      `protobuf:"bytes,10,opt,name=notice" json:"notice,omitempty"`
	GameStatus       *int32       `protobuf:"varint,11,opt,name=gameStatus" json:"gameStatus,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_AckLogin) Reset()                    { *m = Game_AckLogin{} }
func (m *Game_AckLogin) String() string            { return proto.CompactTextString(m) }
func (*Game_AckLogin) ProtoMessage()               {}
func (*Game_AckLogin) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *Game_AckLogin) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_AckLogin) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Game_AckLogin) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *Game_AckLogin) GetRoomPassword() string {
	if m != nil && m.RoomPassword != nil {
		return *m.RoomPassword
	}
	return ""
}

func (m *Game_AckLogin) GetCostCreateRoom() int64 {
	if m != nil && m.CostCreateRoom != nil {
		return *m.CostCreateRoom
	}
	return 0
}

func (m *Game_AckLogin) GetCostRebuy() int64 {
	if m != nil && m.CostRebuy != nil {
		return *m.CostRebuy
	}
	return 0
}

func (m *Game_AckLogin) GetChampionship() bool {
	if m != nil && m.Championship != nil {
		return *m.Championship
	}
	return false
}

func (m *Game_AckLogin) GetChip() int64 {
	if m != nil && m.Chip != nil {
		return *m.Chip
	}
	return 0
}

func (m *Game_AckLogin) GetMailCount() int32 {
	if m != nil && m.MailCount != nil {
		return *m.MailCount
	}
	return 0
}

func (m *Game_AckLogin) GetNotice() string {
	if m != nil && m.Notice != nil {
		return *m.Notice
	}
	return ""
}

func (m *Game_AckLogin) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

type Game_Notice struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	NoticeType       *int32       `protobuf:"varint,2,opt,name=noticeType" json:"noticeType,omitempty"`
	ChannelId        *string      `protobuf:"bytes,3,opt,name=channelId" json:"channelId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_Notice) Reset()                    { *m = Game_Notice{} }
func (m *Game_Notice) String() string            { return proto.CompactTextString(m) }
func (*Game_Notice) ProtoMessage()               {}
func (*Game_Notice) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *Game_Notice) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_Notice) GetNoticeType() int32 {
	if m != nil && m.NoticeType != nil {
		return *m.NoticeType
	}
	return 0
}

func (m *Game_Notice) GetChannelId() string {
	if m != nil && m.ChannelId != nil {
		return *m.ChannelId
	}
	return ""
}

// 公告的内容
type Game_AckNotice struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	NoticeType       *int32       `protobuf:"varint,2,opt,name=noticeType" json:"noticeType,omitempty"`
	NoticeTitle      *string      `protobuf:"bytes,3,opt,name=noticeTitle" json:"noticeTitle,omitempty"`
	NoticeContent    *string      `protobuf:"bytes,4,opt,name=noticeContent" json:"noticeContent,omitempty"`
	NoticeMemo       *string      `protobuf:"bytes,5,opt,name=noticeMemo" json:"noticeMemo,omitempty"`
	Id               *int32       `protobuf:"varint,6,opt,name=id" json:"id,omitempty"`
	Fileds           []string     `protobuf:"bytes,7,rep,name=fileds" json:"fileds,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_AckNotice) Reset()                    { *m = Game_AckNotice{} }
func (m *Game_AckNotice) String() string            { return proto.CompactTextString(m) }
func (*Game_AckNotice) ProtoMessage()               {}
func (*Game_AckNotice) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *Game_AckNotice) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_AckNotice) GetNoticeType() int32 {
	if m != nil && m.NoticeType != nil {
		return *m.NoticeType
	}
	return 0
}

func (m *Game_AckNotice) GetNoticeTitle() string {
	if m != nil && m.NoticeTitle != nil {
		return *m.NoticeTitle
	}
	return ""
}

func (m *Game_AckNotice) GetNoticeContent() string {
	if m != nil && m.NoticeContent != nil {
		return *m.NoticeContent
	}
	return ""
}

func (m *Game_AckNotice) GetNoticeMemo() string {
	if m != nil && m.NoticeMemo != nil {
		return *m.NoticeMemo
	}
	return ""
}

func (m *Game_AckNotice) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Game_AckNotice) GetFileds() []string {
	if m != nil {
		return m.Fileds
	}
	return nil
}

// 游戏战绩
type Game_GameRecord struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Id               *int32       `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	GameId           *int32       `protobuf:"varint,3,opt,name=gameId" json:"gameId,omitempty"`
	UserId           *uint32      `protobuf:"varint,4,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_GameRecord) Reset()                    { *m = Game_GameRecord{} }
func (m *Game_GameRecord) String() string            { return proto.CompactTextString(m) }
func (*Game_GameRecord) ProtoMessage()               {}
func (*Game_GameRecord) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *Game_GameRecord) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_GameRecord) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Game_GameRecord) GetGameId() int32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *Game_GameRecord) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

//
type BeanUserRecord struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=UserId" json:"UserId,omitempty"`
	NickName         *string      `protobuf:"bytes,3,opt,name=NickName" json:"NickName,omitempty"`
	WinAmount        *int64       `protobuf:"varint,4,opt,name=WinAmount" json:"WinAmount,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *BeanUserRecord) Reset()                    { *m = BeanUserRecord{} }
func (m *BeanUserRecord) String() string            { return proto.CompactTextString(m) }
func (*BeanUserRecord) ProtoMessage()               {}
func (*BeanUserRecord) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *BeanUserRecord) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *BeanUserRecord) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *BeanUserRecord) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *BeanUserRecord) GetWinAmount() int64 {
	if m != nil && m.WinAmount != nil {
		return *m.WinAmount
	}
	return 0
}

//
type BeanGameRecord struct {
	Header           *ProtoHeader      `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Id               *int32            `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	DeskId           *int32            `protobuf:"varint,3,opt,name=deskId" json:"deskId,omitempty"`
	BeginTime        *string           `protobuf:"bytes,4,opt,name=beginTime" json:"beginTime,omitempty"`
	Users            []*BeanUserRecord `protobuf:"bytes,5,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *BeanGameRecord) Reset()                    { *m = BeanGameRecord{} }
func (m *BeanGameRecord) String() string            { return proto.CompactTextString(m) }
func (*BeanGameRecord) ProtoMessage()               {}
func (*BeanGameRecord) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *BeanGameRecord) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *BeanGameRecord) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *BeanGameRecord) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *BeanGameRecord) GetBeginTime() string {
	if m != nil && m.BeginTime != nil {
		return *m.BeginTime
	}
	return ""
}

func (m *BeanGameRecord) GetUsers() []*BeanUserRecord {
	if m != nil {
		return m.Users
	}
	return nil
}

//
type Game_AckGameRecord struct {
	Header           *ProtoHeader      `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32           `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	GameId           *int32            `protobuf:"varint,3,opt,name=gameId" json:"gameId,omitempty"`
	Records          []*BeanGameRecord `protobuf:"bytes,4,rep,name=records" json:"records,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Game_AckGameRecord) Reset()                    { *m = Game_AckGameRecord{} }
func (m *Game_AckGameRecord) String() string            { return proto.CompactTextString(m) }
func (*Game_AckGameRecord) ProtoMessage()               {}
func (*Game_AckGameRecord) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *Game_AckGameRecord) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_AckGameRecord) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Game_AckGameRecord) GetGameId() int32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *Game_AckGameRecord) GetRecords() []*BeanGameRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

// 反馈信息的协议
type Game_Feedback struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Message          *string      `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_Feedback) Reset()                    { *m = Game_Feedback{} }
func (m *Game_Feedback) String() string            { return proto.CompactTextString(m) }
func (*Game_Feedback) ProtoMessage()               {}
func (*Game_Feedback) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

func (m *Game_Feedback) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_Feedback) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

// 创建房间
type Game_CreateRoom struct {
	Header           *ProtoHeader  `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	RoomTypeInfo     *RoomTypeInfo `protobuf:"bytes,2,opt,name=roomTypeInfo" json:"roomTypeInfo,omitempty"`
	IsDaiKai         *bool         `protobuf:"varint,3,opt,name=isDaiKai" json:"isDaiKai,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Game_CreateRoom) Reset()                    { *m = Game_CreateRoom{} }
func (m *Game_CreateRoom) String() string            { return proto.CompactTextString(m) }
func (*Game_CreateRoom) ProtoMessage()               {}
func (*Game_CreateRoom) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{12} }

func (m *Game_CreateRoom) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_CreateRoom) GetRoomTypeInfo() *RoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

func (m *Game_CreateRoom) GetIsDaiKai() bool {
	if m != nil && m.IsDaiKai != nil {
		return *m.IsDaiKai
	}
	return false
}

// 创建房间回复的信息
type Game_AckCreateRoom struct {
	Header           *ProtoHeader  `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	DeskId           *int32        `protobuf:"varint,2,opt,name=deskId" json:"deskId,omitempty"`
	Password         *string       `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	UserBalance      *int64        `protobuf:"varint,4,opt,name=userBalance" json:"userBalance,omitempty"`
	CreateFee        *int64        `protobuf:"varint,5,opt,name=createFee" json:"createFee,omitempty"`
	RoomTypeInfo     *RoomTypeInfo `protobuf:"bytes,6,opt,name=roomTypeInfo" json:"roomTypeInfo,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Game_AckCreateRoom) Reset()                    { *m = Game_AckCreateRoom{} }
func (m *Game_AckCreateRoom) String() string            { return proto.CompactTextString(m) }
func (*Game_AckCreateRoom) ProtoMessage()               {}
func (*Game_AckCreateRoom) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{13} }

func (m *Game_AckCreateRoom) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_AckCreateRoom) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *Game_AckCreateRoom) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *Game_AckCreateRoom) GetUserBalance() int64 {
	if m != nil && m.UserBalance != nil {
		return *m.UserBalance
	}
	return 0
}

func (m *Game_AckCreateRoom) GetCreateFee() int64 {
	if m != nil && m.CreateFee != nil {
		return *m.CreateFee
	}
	return 0
}

func (m *Game_AckCreateRoom) GetRoomTypeInfo() *RoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

// 创建房间v2 用作版本迭代 update for 6位随机房间号
type Game_CreateRoomV2 struct {
	Header           *ProtoHeader  `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	RoomTypeInfo     *RoomTypeInfo `protobuf:"bytes,2,opt,name=roomTypeInfo" json:"roomTypeInfo,omitempty"`
	IsDaiKai         *bool         `protobuf:"varint,3,opt,name=isDaiKai" json:"isDaiKai,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Game_CreateRoomV2) Reset()                    { *m = Game_CreateRoomV2{} }
func (m *Game_CreateRoomV2) String() string            { return proto.CompactTextString(m) }
func (*Game_CreateRoomV2) ProtoMessage()               {}
func (*Game_CreateRoomV2) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{14} }

func (m *Game_CreateRoomV2) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_CreateRoomV2) GetRoomTypeInfo() *RoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

func (m *Game_CreateRoomV2) GetIsDaiKai() bool {
	if m != nil && m.IsDaiKai != nil {
		return *m.IsDaiKai
	}
	return false
}

// 客户端请求进入 room, 服务器返回game_SendGameInfo
// 请求进入金币场低中高倍级别时 需要 roomType roomLevel
// 请求进入朋友桌 需要 password
// 请求进入金币场私密房 需要password roomType
// 请求进入金币场公共房 需要deskid
type Game_EnterRoom struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	MatchId          *int32       `protobuf:"varint,2,opt,name=matchId" json:"matchId,omitempty"`
	TableId          *int32       `protobuf:"varint,3,opt,name=tableId" json:"tableId,omitempty"`
	PassWord         *string      `protobuf:"bytes,4,opt,name=PassWord" json:"PassWord,omitempty"`
	UserId           *uint32      `protobuf:"varint,5,opt,name=userId" json:"userId,omitempty"`
	RoomType         *int32       `protobuf:"varint,6,opt,name=roomType" json:"roomType,omitempty"`
	RoomLevel        *int32       `protobuf:"varint,7,opt,name=roomLevel" json:"roomLevel,omitempty"`
	EnterType        *int32       `protobuf:"varint,8,opt,name=enterType" json:"enterType,omitempty"`
	DeskId           *int32       `protobuf:"varint,9,opt,name=deskId" json:"deskId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_EnterRoom) Reset()                    { *m = Game_EnterRoom{} }
func (m *Game_EnterRoom) String() string            { return proto.CompactTextString(m) }
func (*Game_EnterRoom) ProtoMessage()               {}
func (*Game_EnterRoom) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{15} }

func (m *Game_EnterRoom) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_EnterRoom) GetMatchId() int32 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

func (m *Game_EnterRoom) GetTableId() int32 {
	if m != nil && m.TableId != nil {
		return *m.TableId
	}
	return 0
}

func (m *Game_EnterRoom) GetPassWord() string {
	if m != nil && m.PassWord != nil {
		return *m.PassWord
	}
	return ""
}

func (m *Game_EnterRoom) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Game_EnterRoom) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *Game_EnterRoom) GetRoomLevel() int32 {
	if m != nil && m.RoomLevel != nil {
		return *m.RoomLevel
	}
	return 0
}

func (m *Game_EnterRoom) GetEnterType() int32 {
	if m != nil && m.EnterType != nil {
		return *m.EnterType
	}
	return 0
}

func (m *Game_EnterRoom) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

type Game_AckEnterRoom struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_AckEnterRoom) Reset()                    { *m = Game_AckEnterRoom{} }
func (m *Game_AckEnterRoom) String() string            { return proto.CompactTextString(m) }
func (*Game_AckEnterRoom) ProtoMessage()               {}
func (*Game_AckEnterRoom) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{16} }

func (m *Game_AckEnterRoom) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*ServerInfo)(nil), "mjproto.ServerInfo")
	proto.RegisterType((*Game_QuickConn)(nil), "mjproto.game_QuickConn")
	proto.RegisterType((*Game_AckQuickConn)(nil), "mjproto.game_AckQuickConn")
	proto.RegisterType((*Game_Login)(nil), "mjproto.game_Login")
	proto.RegisterType((*Game_AckLogin)(nil), "mjproto.game_AckLogin")
	proto.RegisterType((*Game_Notice)(nil), "mjproto.game_Notice")
	proto.RegisterType((*Game_AckNotice)(nil), "mjproto.game_AckNotice")
	proto.RegisterType((*Game_GameRecord)(nil), "mjproto.game_GameRecord")
	proto.RegisterType((*BeanUserRecord)(nil), "mjproto.BeanUserRecord")
	proto.RegisterType((*BeanGameRecord)(nil), "mjproto.BeanGameRecord")
	proto.RegisterType((*Game_AckGameRecord)(nil), "mjproto.game_AckGameRecord")
	proto.RegisterType((*Game_Feedback)(nil), "mjproto.game_Feedback")
	proto.RegisterType((*Game_CreateRoom)(nil), "mjproto.game_CreateRoom")
	proto.RegisterType((*Game_AckCreateRoom)(nil), "mjproto.game_AckCreateRoom")
	proto.RegisterType((*Game_CreateRoomV2)(nil), "mjproto.game_CreateRoom_v2")
	proto.RegisterType((*Game_EnterRoom)(nil), "mjproto.game_EnterRoom")
	proto.RegisterType((*Game_AckEnterRoom)(nil), "mjproto.game_AckEnterRoom")
}

func init() { proto.RegisterFile("mahjong_hall.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 1011 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x56, 0xdd, 0x6e, 0x23, 0x35,
	0x14, 0x26, 0x93, 0x9f, 0x26, 0x0e, 0x2d, 0x62, 0x80, 0x65, 0x54, 0xa1, 0x55, 0x34, 0x42, 0xa8,
	0x12, 0x50, 0x89, 0x72, 0xc5, 0x65, 0xb7, 0xb0, 0x6c, 0xc5, 0xb6, 0x2a, 0xa6, 0xa1, 0x97, 0x95,
	0x33, 0x73, 0x36, 0xf1, 0x66, 0xc6, 0x1e, 0x8d, 0x27, 0xed, 0xf6, 0x8a, 0x57, 0xe0, 0x6a, 0x25,
	0xae, 0x11, 0x4f, 0xc0, 0x7b, 0xf0, 0x00, 0xbc, 0x01, 0x37, 0x48, 0xbc, 0x01, 0xf2, 0xb1, 0x67,
	0xec, 0x09, 0xbb, 0xd2, 0x26, 0x5a, 0xa4, 0xbd, 0x89, 0xe6, 0x7c, 0x3e, 0xf6, 0xf9, 0xf9, 0x8e,
	0x3f, 0x87, 0x84, 0x39, 0x5b, 0x3c, 0x95, 0x62, 0x7e, 0xbd, 0x60, 0x59, 0x76, 0x58, 0x94, 0xb2,
	0x92, 0xe1, 0x4e, 0xfe, 0x14, 0x3f, 0xf6, 0xc9, 0x8c, 0x29, 0x30, 0x60, 0xfc, 0x88, 0x90, 0x1f,
	0xa0, 0xbc, 0x81, 0xf2, 0x54, 0x3c, 0x91, 0xe1, 0x1e, 0x09, 0x78, 0x11, 0x75, 0x26, 0x9d, 0x83,
	0x11, 0x0d, 0x78, 0x11, 0x86, 0xa4, 0x57, 0xc8, 0xb2, 0x8a, 0x82, 0x49, 0xe7, 0xa0, 0x4f, 0xf1,
	0x3b, 0xbc, 0x47, 0x06, 0xaa, 0x62, 0xd5, 0x4a, 0x45, 0x5d, 0x44, 0xad, 0x15, 0xff, 0xd1, 0x21,
	0x7b, 0x73, 0x96, 0xc3, 0xf5, 0xf7, 0x2b, 0x9e, 0x2c, 0x4f, 0xa4, 0x10, 0xe1, 0x67, 0x64, 0xb0,
	0x00, 0x96, 0x42, 0x89, 0x47, 0x8e, 0x8f, 0xde, 0x3f, 0xb4, 0x29, 0x1c, 0x5e, 0xe8, 0xdf, 0x47,
	0xb8, 0x46, 0xad, 0x4f, 0xf8, 0x11, 0x19, 0x25, 0x0b, 0x26, 0x04, 0x64, 0xa7, 0x29, 0x46, 0x1c,
	0x51, 0x07, 0xe8, 0xb0, 0xfa, 0xf4, 0xd3, 0xb4, 0x0e, 0x6b, 0xac, 0x70, 0x42, 0xc6, 0xc9, 0xaa,
	0x2c, 0x7f, 0x84, 0x52, 0x71, 0x29, 0xa2, 0x1e, 0x2e, 0xfa, 0x50, 0x78, 0x9f, 0x90, 0x8c, 0x89,
	0xf9, 0x8a, 0xcd, 0xf5, 0xee, 0x3e, 0x3a, 0x78, 0x88, 0x3e, 0x79, 0xa5, 0xa0, 0x3c, 0x4d, 0xa3,
	0xc1, 0xa4, 0x73, 0xb0, 0x4b, 0xad, 0x15, 0xff, 0x19, 0x90, 0x77, 0xb1, 0xa0, 0xe3, 0x64, 0xb9,
	0x6d, 0x4d, 0x5f, 0x12, 0xa2, 0x8f, 0x30, 0x2d, 0xc6, 0xa2, 0xc6, 0x47, 0xef, 0x35, 0x3b, 0x5c,
	0xe7, 0xa9, 0xe7, 0xa6, 0x13, 0x2e, 0x21, 0x03, 0xa6, 0xe0, 0x92, 0xcd, 0x6d, 0xb9, 0x1e, 0xf2,
	0x0a, 0x25, 0xef, 0x93, 0x21, 0x57, 0xd3, 0x22, 0x65, 0x15, 0xd8, 0x82, 0x1b, 0x5b, 0xef, 0x4e,
	0xe5, 0xad, 0xc8, 0x24, 0x4b, 0xa7, 0x65, 0x86, 0x35, 0x8f, 0xa8, 0x0f, 0x69, 0x8f, 0x1b, 0x73,
	0x90, 0x4e, 0x2d, 0xda, 0x31, 0x1e, 0x1e, 0xa4, 0x33, 0xe4, 0xea, 0x8c, 0x71, 0x51, 0x31, 0x2e,
	0xa2, 0xa1, 0xc9, 0xd0, 0x21, 0xfa, 0x84, 0xdc, 0x7e, 0x9f, 0xa9, 0x79, 0x34, 0x32, 0x27, 0x78,
	0x50, 0xfc, 0x6b, 0xc7, 0x74, 0xe6, 0xfa, 0xb1, 0x9c, 0xf3, 0x4d, 0xbb, 0xea, 0x18, 0x0b, 0x7c,
	0xc6, 0xc2, 0x98, 0xbc, 0x8d, 0x9b, 0xea, 0xce, 0x98, 0xd6, 0xb5, 0xb0, 0xf0, 0x53, 0x32, 0xb8,
	0x7d, 0x86, 0x75, 0xf5, 0xd6, 0xd8, 0xb8, 0x02, 0xfe, 0x8c, 0x63, 0x7d, 0xd4, 0xba, 0xc4, 0x7f,
	0x05, 0x64, 0xb7, 0x1e, 0x81, 0xd7, 0x99, 0xe8, 0x3e, 0x19, 0x0a, 0x9e, 0x2c, 0xcf, 0x59, 0x0e,
	0x98, 0xe4, 0x88, 0x36, 0xb6, 0x2e, 0xa2, 0x94, 0x32, 0xbf, 0x60, 0x4a, 0xdd, 0xca, 0x32, 0xc5,
	0x34, 0x47, 0xb4, 0x85, 0x85, 0x9f, 0x90, 0xbd, 0x44, 0xaa, 0xea, 0xa4, 0x04, 0x56, 0x01, 0x95,
	0x32, 0x47, 0x96, 0xbb, 0x74, 0x0d, 0xc5, 0x2b, 0x25, 0x55, 0x45, 0x61, 0xb6, 0xba, 0x43, 0xa6,
	0xbb, 0xd4, 0x01, 0x3a, 0x52, 0xb2, 0x60, 0x79, 0xc1, 0xa5, 0x50, 0x0b, 0x5e, 0x20, 0xd1, 0x43,
	0xda, 0xc2, 0xb4, 0x02, 0x24, 0x7a, 0x6d, 0x88, 0x9b, 0xf1, 0x5b, 0x9f, 0x9a, 0x33, 0x9e, 0x9d,
	0xc8, 0x95, 0xa8, 0x90, 0xdb, 0x3e, 0x75, 0x80, 0xae, 0x59, 0xc8, 0x8a, 0x27, 0x10, 0x11, 0xcc,
	0xdc, 0x5a, 0x7a, 0x66, 0x70, 0xc6, 0x8d, 0x76, 0x8c, 0xcd, 0xcc, 0x38, 0x24, 0xbe, 0x23, 0x63,
	0x6c, 0xf5, 0xb9, 0x71, 0xdf, 0xac, 0xd1, 0xf7, 0x09, 0x31, 0x61, 0x2e, 0xef, 0x0a, 0xb0, 0x72,
	0xe5, 0x21, 0x6d, 0x6d, 0xe9, 0xae, 0x69, 0x4b, 0xfc, 0x77, 0x2d, 0x5d, 0xc7, 0xc9, 0xf2, 0x7f,
	0x09, 0x3f, 0x21, 0x63, 0x6b, 0xf1, 0x2a, 0xab, 0x29, 0xf7, 0xa1, 0xf0, 0x63, 0xb2, 0x6b, 0xcc,
	0x13, 0x29, 0x2a, 0x10, 0x95, 0xa5, 0xbd, 0x0d, 0xba, 0x38, 0x67, 0x90, 0x4b, 0xe4, 0x7c, 0x44,
	0x3d, 0x04, 0xf5, 0xdb, 0xc8, 0x58, 0x9f, 0x06, 0x1c, 0xa5, 0xed, 0x09, 0xcf, 0x20, 0x55, 0xd1,
	0xce, 0xa4, 0xab, 0xb9, 0x30, 0x56, 0xfc, 0x13, 0x79, 0x07, 0xeb, 0xfd, 0x96, 0xe5, 0x40, 0x21,
	0xd1, 0x23, 0xb5, 0x59, 0xc1, 0x26, 0x50, 0xe0, 0x07, 0x7a, 0xa1, 0x3a, 0xbb, 0x0b, 0xd0, 0x6b,
	0x69, 0xeb, 0xcf, 0x1d, 0xb2, 0xf7, 0x00, 0x98, 0x98, 0x2a, 0x28, 0xb7, 0x4a, 0xe0, 0x1e, 0x19,
	0x4c, 0x5b, 0x37, 0x6b, 0xda, 0xdc, 0xac, 0xf3, 0xb5, 0x9b, 0x55, 0xdb, 0x7a, 0x08, 0xae, 0xb8,
	0x38, 0xce, 0x71, 0x6e, 0x7b, 0xe6, 0x36, 0x34, 0x40, 0xfc, 0xbb, 0x4d, 0xe9, 0x75, 0xf6, 0x24,
	0x05, 0xb5, 0x74, 0x3d, 0x31, 0x96, 0x4e, 0x63, 0x06, 0x73, 0x2e, 0x2e, 0x79, 0x0e, 0x96, 0x66,
	0x07, 0x84, 0x9f, 0x93, 0xbe, 0xee, 0x91, 0x8a, 0xfa, 0x93, 0xee, 0xc1, 0xf8, 0xe8, 0xc3, 0x26,
	0x64, 0xbb, 0x5d, 0xd4, 0x78, 0xc5, 0xbf, 0x75, 0x48, 0x58, 0x8f, 0xee, 0xd6, 0x99, 0xbf, 0x4c,
	0xa6, 0x5e, 0xc6, 0xea, 0x17, 0x64, 0xa7, 0xc4, 0x38, 0x2a, 0xea, 0xbd, 0x20, 0x4b, 0x97, 0x07,
	0xad, 0xfd, 0xe2, 0x2b, 0x2b, 0xa4, 0x0f, 0x01, 0xd2, 0x19, 0x4b, 0x96, 0x1b, 0x66, 0x18, 0x91,
	0x9d, 0x1c, 0x94, 0x62, 0x73, 0xb0, 0xff, 0x0c, 0x6a, 0x33, 0x7e, 0xde, 0xb1, 0xb3, 0xec, 0xc9,
	0xde, 0x66, 0x67, 0x7f, 0x65, 0x04, 0x57, 0x5f, 0x54, 0x7c, 0x17, 0xcc, 0x2b, 0xfd, 0x41, 0xb3,
	0x87, 0x7a, 0x8b, 0xb4, 0xe5, 0x6a, 0xde, 0xd9, 0xaf, 0x19, 0xff, 0x8e, 0x71, 0x6c, 0xd1, 0x90,
	0x36, 0x76, 0xfc, 0x8f, 0xc7, 0xcc, 0xd6, 0xb9, 0xb9, 0x19, 0x0a, 0x5a, 0x33, 0xb4, 0x4f, 0x86,
	0x45, 0xfd, 0x40, 0xd8, 0x31, 0xaf, 0x6d, 0x2d, 0x36, 0x9a, 0xbf, 0x07, 0x2c, 0x63, 0x22, 0x01,
	0x3b, 0xe8, 0x3e, 0x84, 0x6a, 0x88, 0x19, 0x3d, 0x04, 0xb0, 0x2f, 0x87, 0x03, 0xfe, 0xd3, 0x8f,
	0xc1, 0x2b, 0xf7, 0x23, 0xfe, 0xa5, 0xae, 0xd9, 0x15, 0x7c, 0x7d, 0x73, 0xf4, 0x66, 0xf0, 0xf1,
	0x3c, 0xb0, 0x22, 0xff, 0x8d, 0xa8, 0xa0, 0xdc, 0x82, 0x0b, 0x3d, 0x83, 0xac, 0x4a, 0x16, 0x0d,
	0x19, 0xb5, 0xa9, 0x57, 0x2a, 0x36, 0xcb, 0xdc, 0x45, 0xa9, 0x4d, 0x9d, 0x90, 0x7e, 0xb4, 0xaf,
	0xdc, 0x43, 0xde, 0xd8, 0xde, 0xad, 0xeb, 0xaf, 0xff, 0x39, 0xa8, 0x8b, 0xb2, 0x52, 0xde, 0xd8,
	0x9a, 0x39, 0xfd, 0xfd, 0x18, 0x6e, 0x20, 0xc3, 0xf7, 0xba, 0x4f, 0x1d, 0xa0, 0x57, 0x41, 0x17,
	0x87, 0x5b, 0xcd, 0xbf, 0x32, 0x07, 0x78, 0xb3, 0x34, 0xf2, 0x67, 0x29, 0x3e, 0x76, 0x7f, 0x73,
	0xb7, 0x6c, 0xcd, 0xc5, 0x5b, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x07, 0xec, 0xae, 0x71,
	0x0c, 0x00, 0x00,
}
