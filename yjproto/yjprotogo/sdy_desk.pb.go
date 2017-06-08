// Code generated by protoc-gen-go.
// source: sdy_desk.proto
// DO NOT EDIT!

package yjprotogo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ProtoHeader from common_client.proto

// Ignoring public import of ServerInfo from common_client.proto

// Ignoring public import of QuickConn from common_client.proto

// Ignoring public import of AckQuickConn from common_client.proto

// Ignoring public import of WeixinInfo from common_client.proto

// Ignoring public import of common_req_reg from common_client.proto

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of cm_offline from common_client.proto

// Ignoring public import of cm_hearbeat from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of sdy_base_userPaiIds from sdy_base.proto

// Ignoring public import of sdy_base_roomTypeInfo from sdy_base.proto

// Ignoring public import of sdy_base_timerInfo from sdy_base.proto

// Ignoring public import of sdy_base_playerInfo from sdy_base.proto

// Ignoring public import of sdy_base_commonRateInfo from sdy_base.proto

// Ignoring public import of sdy_base_deskInfo from sdy_base.proto

// Ignoring public import of sdy_enum_protoId from sdy_base.proto

// Ignoring public import of sdy_enum_errorCode from sdy_base.proto

// Ignoring public import of sdy_enum_actType from sdy_base.proto

// Ignoring public import of sdy_enum_deskStatus from sdy_base.proto

// Ignoring public import of sdy_enum_userStatus from sdy_base.proto

// Ignoring public import of sdy_enum_enterType from sdy_base.proto

// Ignoring public import of sdy_enum_Option from sdy_base.proto

// Ignoring public import of sdy_enum_coinRoomLevel from sdy_base.proto

// Ignoring public import of sdy_enum_flowers from sdy_base.proto

// 准备游戏
type SdyReqReady struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyReqReady) Reset()                    { *m = SdyReqReady{} }
func (m *SdyReqReady) String() string            { return proto.CompactTextString(m) }
func (*SdyReqReady) ProtoMessage()               {}
func (*SdyReqReady) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *SdyReqReady) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyReqReady) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 准备游戏的结果
type SdyAckReady struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserIds          []uint32     `protobuf:"varint,2,rep,name=userIds" json:"userIds,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyAckReady) Reset()                    { *m = SdyAckReady{} }
func (m *SdyAckReady) String() string            { return proto.CompactTextString(m) }
func (*SdyAckReady) ProtoMessage()               {}
func (*SdyAckReady) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

func (m *SdyAckReady) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyAckReady) GetUserIds() []uint32 {
	if m != nil {
		return m.UserIds
	}
	return nil
}

// 赢牌信息：谁赢了多少
type SdyBaseWinCoinInfo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	TotalScore       *int64       `protobuf:"varint,2,opt,name=totalScore" json:"totalScore,omitempty"`
	RoundScore       *int64       `protobuf:"varint,3,opt,name=roundScore" json:"roundScore,omitempty"`
	WinNum           *int32       `protobuf:"varint,4,opt,name=winNum" json:"winNum,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBaseWinCoinInfo) Reset()                    { *m = SdyBaseWinCoinInfo{} }
func (m *SdyBaseWinCoinInfo) String() string            { return proto.CompactTextString(m) }
func (*SdyBaseWinCoinInfo) ProtoMessage()               {}
func (*SdyBaseWinCoinInfo) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

func (m *SdyBaseWinCoinInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBaseWinCoinInfo) GetTotalScore() int64 {
	if m != nil && m.TotalScore != nil {
		return *m.TotalScore
	}
	return 0
}

func (m *SdyBaseWinCoinInfo) GetRoundScore() int64 {
	if m != nil && m.RoundScore != nil {
		return *m.RoundScore
	}
	return 0
}

func (m *SdyBaseWinCoinInfo) GetWinNum() int32 {
	if m != nil && m.WinNum != nil {
		return *m.WinNum
	}
	return 0
}

// 本局结果(广播)
type SdyBcCurrentResult struct {
	WinCoinInfo      []*SdyBaseWinCoinInfo `protobuf:"bytes,1,rep,name=winCoinInfo" json:"winCoinInfo,omitempty"`
	CurrentRound     *int32                `protobuf:"varint,2,opt,name=currentRound" json:"currentRound,omitempty"`
	BoardsCout       *int32                `protobuf:"varint,3,opt,name=boardsCout" json:"boardsCout,omitempty"`
	FootPokers       []int32               `protobuf:"varint,4,rep,name=footPokers" json:"footPokers,omitempty"`
	HuanDiPokers     []int32               `protobuf:"varint,5,rep,name=huanDiPokers" json:"huanDiPokers,omitempty"`
	BaseValue        *int64                `protobuf:"varint,6,opt,name=baseValue" json:"baseValue,omitempty"`
	IsWangKou        *bool                 `protobuf:"varint,7,opt,name=isWangKou" json:"isWangKou,omitempty"`
	IsPoPai          *bool                 `protobuf:"varint,8,opt,name=isPoPai" json:"isPoPai,omitempty"`
	IsKouDi          *bool                 `protobuf:"varint,9,opt,name=isKouDi" json:"isKouDi,omitempty"`
	IsGuang          *bool                 `protobuf:"varint,10,opt,name=isGuang" json:"isGuang,omitempty"`
	IsShangChe       *bool                 `protobuf:"varint,11,opt,name=isShangChe" json:"isShangChe,omitempty"`
	IsChengPai       *bool                 `protobuf:"varint,12,opt,name=isChengPai" json:"isChengPai,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *SdyBcCurrentResult) Reset()                    { *m = SdyBcCurrentResult{} }
func (m *SdyBcCurrentResult) String() string            { return proto.CompactTextString(m) }
func (*SdyBcCurrentResult) ProtoMessage()               {}
func (*SdyBcCurrentResult) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{3} }

func (m *SdyBcCurrentResult) GetWinCoinInfo() []*SdyBaseWinCoinInfo {
	if m != nil {
		return m.WinCoinInfo
	}
	return nil
}

func (m *SdyBcCurrentResult) GetCurrentRound() int32 {
	if m != nil && m.CurrentRound != nil {
		return *m.CurrentRound
	}
	return 0
}

func (m *SdyBcCurrentResult) GetBoardsCout() int32 {
	if m != nil && m.BoardsCout != nil {
		return *m.BoardsCout
	}
	return 0
}

func (m *SdyBcCurrentResult) GetFootPokers() []int32 {
	if m != nil {
		return m.FootPokers
	}
	return nil
}

func (m *SdyBcCurrentResult) GetHuanDiPokers() []int32 {
	if m != nil {
		return m.HuanDiPokers
	}
	return nil
}

func (m *SdyBcCurrentResult) GetBaseValue() int64 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

func (m *SdyBcCurrentResult) GetIsWangKou() bool {
	if m != nil && m.IsWangKou != nil {
		return *m.IsWangKou
	}
	return false
}

func (m *SdyBcCurrentResult) GetIsPoPai() bool {
	if m != nil && m.IsPoPai != nil {
		return *m.IsPoPai
	}
	return false
}

func (m *SdyBcCurrentResult) GetIsKouDi() bool {
	if m != nil && m.IsKouDi != nil {
		return *m.IsKouDi
	}
	return false
}

func (m *SdyBcCurrentResult) GetIsGuang() bool {
	if m != nil && m.IsGuang != nil {
		return *m.IsGuang
	}
	return false
}

func (m *SdyBcCurrentResult) GetIsShangChe() bool {
	if m != nil && m.IsShangChe != nil {
		return *m.IsShangChe
	}
	return false
}

func (m *SdyBcCurrentResult) GetIsChengPai() bool {
	if m != nil && m.IsChengPai != nil {
		return *m.IsChengPai
	}
	return false
}

type SdyBaseBill struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	WXInfo           *WeixinInfo  `protobuf:"bytes,2,opt,name=wXInfo" json:"wXInfo,omitempty"`
	IsBigWin         *bool        `protobuf:"varint,3,opt,name=isBigWin" json:"isBigWin,omitempty"`
	IsOwner          *bool        `protobuf:"varint,4,opt,name=isOwner" json:"isOwner,omitempty"`
	WinCoin          *int64       `protobuf:"varint,5,opt,name=winCoin" json:"winCoin,omitempty"`
	ChengPai         *int32       `protobuf:"varint,6,opt,name=chengPai" json:"chengPai,omitempty"`
	PoPai            *int32       `protobuf:"varint,7,opt,name=poPai" json:"poPai,omitempty"`
	KouDi            *int32       `protobuf:"varint,8,opt,name=kouDi" json:"kouDi,omitempty"`
	GuangPai         *int32       `protobuf:"varint,9,opt,name=guangPai" json:"guangPai,omitempty"`
	ShangChe         *int32       `protobuf:"varint,10,opt,name=shangChe" json:"shangChe,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBaseBill) Reset()                    { *m = SdyBaseBill{} }
func (m *SdyBaseBill) String() string            { return proto.CompactTextString(m) }
func (*SdyBaseBill) ProtoMessage()               {}
func (*SdyBaseBill) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{4} }

func (m *SdyBaseBill) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBaseBill) GetWXInfo() *WeixinInfo {
	if m != nil {
		return m.WXInfo
	}
	return nil
}

func (m *SdyBaseBill) GetIsBigWin() bool {
	if m != nil && m.IsBigWin != nil {
		return *m.IsBigWin
	}
	return false
}

func (m *SdyBaseBill) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *SdyBaseBill) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *SdyBaseBill) GetChengPai() int32 {
	if m != nil && m.ChengPai != nil {
		return *m.ChengPai
	}
	return 0
}

func (m *SdyBaseBill) GetPoPai() int32 {
	if m != nil && m.PoPai != nil {
		return *m.PoPai
	}
	return 0
}

func (m *SdyBaseBill) GetKouDi() int32 {
	if m != nil && m.KouDi != nil {
		return *m.KouDi
	}
	return 0
}

func (m *SdyBaseBill) GetGuangPai() int32 {
	if m != nil && m.GuangPai != nil {
		return *m.GuangPai
	}
	return 0
}

func (m *SdyBaseBill) GetShangChe() int32 {
	if m != nil && m.ShangChe != nil {
		return *m.ShangChe
	}
	return 0
}

// 全局结果
type SdyBcEndLotteryInfo struct {
	UserBills        []*SdyBaseBill `protobuf:"bytes,1,rep,name=userBills" json:"userBills,omitempty"`
	CurrentRound     *int32         `protobuf:"varint,2,opt,name=currentRound" json:"currentRound,omitempty"`
	BoardsCout       *int32         `protobuf:"varint,3,opt,name=boardsCout" json:"boardsCout,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *SdyBcEndLotteryInfo) Reset()                    { *m = SdyBcEndLotteryInfo{} }
func (m *SdyBcEndLotteryInfo) String() string            { return proto.CompactTextString(m) }
func (*SdyBcEndLotteryInfo) ProtoMessage()               {}
func (*SdyBcEndLotteryInfo) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{5} }

func (m *SdyBcEndLotteryInfo) GetUserBills() []*SdyBaseBill {
	if m != nil {
		return m.UserBills
	}
	return nil
}

func (m *SdyBcEndLotteryInfo) GetCurrentRound() int32 {
	if m != nil && m.CurrentRound != nil {
		return *m.CurrentRound
	}
	return 0
}

func (m *SdyBcEndLotteryInfo) GetBoardsCout() int32 {
	if m != nil && m.BoardsCout != nil {
		return *m.BoardsCout
	}
	return 0
}

// 出的牌,出牌的玩家id
type SdyReUserOutPai struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	PokerId          *int32  `protobuf:"varint,2,opt,name=pokerId" json:"pokerId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SdyReUserOutPai) Reset()                    { *m = SdyReUserOutPai{} }
func (m *SdyReUserOutPai) String() string            { return proto.CompactTextString(m) }
func (*SdyReUserOutPai) ProtoMessage()               {}
func (*SdyReUserOutPai) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{6} }

func (m *SdyReUserOutPai) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyReUserOutPai) GetPokerId() int32 {
	if m != nil && m.PokerId != nil {
		return *m.PokerId
	}
	return 0
}

// 准备阶段需要发的消息
type SdyReReady struct {
	HandPokersId     []int32 `protobuf:"varint,1,rep,name=handPokersId" json:"handPokersId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SdyReReady) Reset()                    { *m = SdyReReady{} }
func (m *SdyReReady) String() string            { return proto.CompactTextString(m) }
func (*SdyReReady) ProtoMessage()               {}
func (*SdyReReady) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{7} }

func (m *SdyReReady) GetHandPokersId() []int32 {
	if m != nil {
		return m.HandPokersId
	}
	return nil
}

type SdyReHuanDi struct {
	DeskFootPokersId []int32 `protobuf:"varint,1,rep,name=deskFootPokersId" json:"deskFootPokersId,omitempty"`
	HuanDiPokersId   []int32 `protobuf:"varint,2,rep,name=huanDiPokersId" json:"huanDiPokersId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SdyReHuanDi) Reset()                    { *m = SdyReHuanDi{} }
func (m *SdyReHuanDi) String() string            { return proto.CompactTextString(m) }
func (*SdyReHuanDi) ProtoMessage()               {}
func (*SdyReHuanDi) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{8} }

func (m *SdyReHuanDi) GetDeskFootPokersId() []int32 {
	if m != nil {
		return m.DeskFootPokersId
	}
	return nil
}

func (m *SdyReHuanDi) GetHuanDiPokersId() []int32 {
	if m != nil {
		return m.HuanDiPokersId
	}
	return nil
}

type SdyRePlay struct {
	OutPai           []*SdyReUserOutPai `protobuf:"bytes,1,rep,name=outPai" json:"outPai,omitempty"`
	ScorePai         []int32            `protobuf:"varint,2,rep,name=scorePai" json:"scorePai,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *SdyRePlay) Reset()                    { *m = SdyRePlay{} }
func (m *SdyRePlay) String() string            { return proto.CompactTextString(m) }
func (*SdyRePlay) ProtoMessage()               {}
func (*SdyRePlay) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{9} }

func (m *SdyRePlay) GetOutPai() []*SdyReUserOutPai {
	if m != nil {
		return m.OutPai
	}
	return nil
}

func (m *SdyRePlay) GetScorePai() []int32 {
	if m != nil {
		return m.ScorePai
	}
	return nil
}

type SdyReJiaoFen struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	JiaoScore        *int32  `protobuf:"varint,2,opt,name=jiaoScore" json:"jiaoScore,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SdyReJiaoFen) Reset()                    { *m = SdyReJiaoFen{} }
func (m *SdyReJiaoFen) String() string            { return proto.CompactTextString(m) }
func (*SdyReJiaoFen) ProtoMessage()               {}
func (*SdyReJiaoFen) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{10} }

func (m *SdyReJiaoFen) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyReJiaoFen) GetJiaoScore() int32 {
	if m != nil && m.JiaoScore != nil {
		return *m.JiaoScore
	}
	return 0
}

type SdyReLenHandPokers struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	PokersLength     *int32  `protobuf:"varint,2,opt,name=pokersLength" json:"pokersLength,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SdyReLenHandPokers) Reset()                    { *m = SdyReLenHandPokers{} }
func (m *SdyReLenHandPokers) String() string            { return proto.CompactTextString(m) }
func (*SdyReLenHandPokers) ProtoMessage()               {}
func (*SdyReLenHandPokers) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{11} }

func (m *SdyReLenHandPokers) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyReLenHandPokers) GetPokersLength() int32 {
	if m != nil && m.PokersLength != nil {
		return *m.PokersLength
	}
	return 0
}

// 用于断线重连回来后发送的消息
type SdyBcReconnectInfo struct {
	Header           *ProtoHeader          `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PlayerInfo       []*SdyBasePlayerInfo  `protobuf:"bytes,2,rep,name=playerInfo" json:"playerInfo,omitempty"`
	DeskInfo         *SdyBaseDeskInfo      `protobuf:"bytes,3,opt,name=deskInfo" json:"deskInfo,omitempty"`
	ReconnectUser    *uint32               `protobuf:"varint,4,opt,name=reconnectUser" json:"reconnectUser,omitempty"`
	ZhuFlower        *int32                `protobuf:"varint,5,opt,name=zhuFlower" json:"zhuFlower,omitempty"`
	ActiveUser       *uint32               `protobuf:"varint,6,opt,name=activeUser" json:"activeUser,omitempty"`
	DeskStatus       *int32                `protobuf:"varint,7,opt,name=deskStatus" json:"deskStatus,omitempty"`
	BankerFen        *int32                `protobuf:"varint,8,opt,name=bankerFen" json:"bankerFen,omitempty"`
	Ready            *SdyReReady           `protobuf:"bytes,9,opt,name=ready" json:"ready,omitempty"`
	HuanDi           *SdyReHuanDi          `protobuf:"bytes,10,opt,name=huanDi" json:"huanDi,omitempty"`
	OutPai           *SdyRePlay            `protobuf:"bytes,11,opt,name=outPai" json:"outPai,omitempty"`
	CurrentResult    *SdyBcCurrentResult   `protobuf:"bytes,12,opt,name=currentResult" json:"currentResult,omitempty"`
	JiaoFen          []*SdyReJiaoFen       `protobuf:"bytes,13,rep,name=jiaoFen" json:"jiaoFen,omitempty"`
	PaiLen           []*SdyReLenHandPokers `protobuf:"bytes,14,rep,name=paiLen" json:"paiLen,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *SdyBcReconnectInfo) Reset()                    { *m = SdyBcReconnectInfo{} }
func (m *SdyBcReconnectInfo) String() string            { return proto.CompactTextString(m) }
func (*SdyBcReconnectInfo) ProtoMessage()               {}
func (*SdyBcReconnectInfo) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{12} }

func (m *SdyBcReconnectInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcReconnectInfo) GetPlayerInfo() []*SdyBasePlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *SdyBcReconnectInfo) GetDeskInfo() *SdyBaseDeskInfo {
	if m != nil {
		return m.DeskInfo
	}
	return nil
}

func (m *SdyBcReconnectInfo) GetReconnectUser() uint32 {
	if m != nil && m.ReconnectUser != nil {
		return *m.ReconnectUser
	}
	return 0
}

func (m *SdyBcReconnectInfo) GetZhuFlower() int32 {
	if m != nil && m.ZhuFlower != nil {
		return *m.ZhuFlower
	}
	return 0
}

func (m *SdyBcReconnectInfo) GetActiveUser() uint32 {
	if m != nil && m.ActiveUser != nil {
		return *m.ActiveUser
	}
	return 0
}

func (m *SdyBcReconnectInfo) GetDeskStatus() int32 {
	if m != nil && m.DeskStatus != nil {
		return *m.DeskStatus
	}
	return 0
}

func (m *SdyBcReconnectInfo) GetBankerFen() int32 {
	if m != nil && m.BankerFen != nil {
		return *m.BankerFen
	}
	return 0
}

func (m *SdyBcReconnectInfo) GetReady() *SdyReReady {
	if m != nil {
		return m.Ready
	}
	return nil
}

func (m *SdyBcReconnectInfo) GetHuanDi() *SdyReHuanDi {
	if m != nil {
		return m.HuanDi
	}
	return nil
}

func (m *SdyBcReconnectInfo) GetOutPai() *SdyRePlay {
	if m != nil {
		return m.OutPai
	}
	return nil
}

func (m *SdyBcReconnectInfo) GetCurrentResult() *SdyBcCurrentResult {
	if m != nil {
		return m.CurrentResult
	}
	return nil
}

func (m *SdyBcReconnectInfo) GetJiaoFen() []*SdyReJiaoFen {
	if m != nil {
		return m.JiaoFen
	}
	return nil
}

func (m *SdyBcReconnectInfo) GetPaiLen() []*SdyReLenHandPokers {
	if m != nil {
		return m.PaiLen
	}
	return nil
}

// 用于玩家断线重连回来后广播给其他玩家，通知该玩家是否在线
type SdyBcIsOnLine struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	IsOnline         *bool        `protobuf:"varint,3,opt,name=isOnline" json:"isOnline,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBcIsOnLine) Reset()                    { *m = SdyBcIsOnLine{} }
func (m *SdyBcIsOnLine) String() string            { return proto.CompactTextString(m) }
func (*SdyBcIsOnLine) ProtoMessage()               {}
func (*SdyBcIsOnLine) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{13} }

func (m *SdyBcIsOnLine) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcIsOnLine) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyBcIsOnLine) GetIsOnline() bool {
	if m != nil && m.IsOnline != nil {
		return *m.IsOnline
	}
	return false
}

func init() {
	proto.RegisterType((*SdyReqReady)(nil), "yjprotogo.sdy_req_ready")
	proto.RegisterType((*SdyAckReady)(nil), "yjprotogo.sdy_ack_ready")
	proto.RegisterType((*SdyBaseWinCoinInfo)(nil), "yjprotogo.sdy_base_winCoinInfo")
	proto.RegisterType((*SdyBcCurrentResult)(nil), "yjprotogo.sdy_bc_currentResult")
	proto.RegisterType((*SdyBaseBill)(nil), "yjprotogo.sdy_base_bill")
	proto.RegisterType((*SdyBcEndLotteryInfo)(nil), "yjprotogo.sdy_bc_endLotteryInfo")
	proto.RegisterType((*SdyReUserOutPai)(nil), "yjprotogo.sdy_re_userOutPai")
	proto.RegisterType((*SdyReReady)(nil), "yjprotogo.sdy_re_ready")
	proto.RegisterType((*SdyReHuanDi)(nil), "yjprotogo.sdy_re_huanDi")
	proto.RegisterType((*SdyRePlay)(nil), "yjprotogo.sdy_re_play")
	proto.RegisterType((*SdyReJiaoFen)(nil), "yjprotogo.sdy_re_jiaoFen")
	proto.RegisterType((*SdyReLenHandPokers)(nil), "yjprotogo.sdy_re_lenHandPokers")
	proto.RegisterType((*SdyBcReconnectInfo)(nil), "yjprotogo.sdy_bc_reconnectInfo")
	proto.RegisterType((*SdyBcIsOnLine)(nil), "yjprotogo.sdy_bc_isOnLine")
}

var fileDescriptor9 = []byte{
	// 984 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdf, 0x6e, 0xe3, 0xc4,
	0x17, 0xfe, 0x39, 0x59, 0xa7, 0xc9, 0x49, 0xd3, 0x1f, 0x98, 0xdd, 0xc5, 0x54, 0xd5, 0x12, 0x59,
	0x08, 0x45, 0x48, 0x5b, 0xa1, 0x82, 0x80, 0x2b, 0x24, 0xda, 0xdd, 0xd0, 0x6a, 0x2b, 0x1a, 0x4d,
	0x05, 0x85, 0x2b, 0xcb, 0x75, 0x66, 0x93, 0xd9, 0xb8, 0x33, 0x61, 0x6c, 0x13, 0xc2, 0x6b, 0x70,
	0xcf, 0x35, 0x12, 0xef, 0xc0, 0xb3, 0xa1, 0x73, 0x66, 0xc6, 0x76, 0x9a, 0xec, 0xc5, 0xae, 0xf6,
	0xaa, 0x3d, 0xdf, 0xf7, 0xf9, 0xcc, 0x99, 0xf3, 0x6f, 0x02, 0x07, 0xf9, 0x74, 0x1d, 0x4f, 0x79,
	0xbe, 0x38, 0x5e, 0x6a, 0x55, 0xa8, 0xa0, 0xb7, 0x7e, 0x45, 0xff, 0xcc, 0xd4, 0xe1, 0x07, 0xa9,
	0xba, 0xbb, 0x53, 0x32, 0x4e, 0x33, 0xc1, 0x65, 0x61, 0xf8, 0x43, 0xd2, 0xdf, 0x26, 0x39, 0x37,
	0x76, 0x74, 0x03, 0x03, 0x44, 0x34, 0xff, 0x35, 0xd6, 0x3c, 0x99, 0xae, 0x83, 0x63, 0xe8, 0xcc,
	0x79, 0x32, 0xe5, 0x3a, 0xf4, 0x86, 0xde, 0xa8, 0x7f, 0xf2, 0xf8, 0xb8, 0xf2, 0x78, 0x3c, 0xc1,
	0xbf, 0xe7, 0xc4, 0x32, 0xab, 0x0a, 0x1e, 0x43, 0xa7, 0xcc, 0xb9, 0xbe, 0x98, 0x86, 0xad, 0xa1,
	0x37, 0x1a, 0x30, 0x6b, 0x45, 0xbf, 0x18, 0xc7, 0x49, 0xba, 0x78, 0x4b, 0xc7, 0x21, 0xec, 0x19,
	0x57, 0x79, 0xd8, 0x1a, 0xb6, 0x47, 0x03, 0xe6, 0xcc, 0xe8, 0x2f, 0x0f, 0x1e, 0xba, 0x6b, 0xc4,
	0x2b, 0x21, 0xcf, 0x94, 0x90, 0x17, 0xf2, 0xa5, 0x7a, 0xe3, 0x23, 0x9e, 0x00, 0x14, 0xaa, 0x48,
	0xb2, 0xeb, 0x54, 0x69, 0x4e, 0xf1, 0xb7, 0x59, 0x03, 0x41, 0x5e, 0xab, 0x52, 0x4e, 0x0d, 0xdf,
	0x36, 0x7c, 0x8d, 0xe0, 0xdd, 0x57, 0x42, 0xfe, 0x50, 0xde, 0x85, 0x0f, 0x86, 0xde, 0xc8, 0x67,
	0xd6, 0x8a, 0xfe, 0x69, 0xdb, 0x00, 0xd3, 0x38, 0x2d, 0xb5, 0xe6, 0xb2, 0x60, 0x3c, 0x2f, 0xb3,
	0x22, 0xf8, 0x0e, 0xfa, 0x8d, 0x78, 0x43, 0x6f, 0xd8, 0x1e, 0xf5, 0x4f, 0x3e, 0x6e, 0x44, 0xb9,
	0xeb, 0x5a, 0xac, 0xf9, 0x4d, 0x10, 0xc1, 0xbe, 0xf3, 0x89, 0x81, 0x50, 0xd4, 0x3e, 0xdb, 0xc0,
	0x30, 0xee, 0x5b, 0x95, 0xe8, 0x69, 0x7e, 0xa6, 0xca, 0x82, 0xe2, 0xf6, 0x59, 0x03, 0x41, 0xfe,
	0xa5, 0x52, 0xc5, 0x44, 0x2d, 0xb8, 0xce, 0xc3, 0x07, 0xc3, 0x36, 0xf2, 0x35, 0x82, 0x67, 0xcc,
	0xcb, 0x44, 0x3e, 0x13, 0x56, 0xe1, 0x93, 0x62, 0x03, 0x0b, 0x8e, 0xa0, 0x87, 0x81, 0xfe, 0x94,
	0x64, 0x25, 0x0f, 0x3b, 0x94, 0x9a, 0x1a, 0x40, 0x56, 0xe4, 0x37, 0x89, 0x9c, 0xbd, 0x50, 0x65,
	0xb8, 0x37, 0xf4, 0x46, 0x5d, 0x56, 0x03, 0x58, 0x5a, 0x91, 0x4f, 0xd4, 0x24, 0x11, 0x61, 0x97,
	0x38, 0x67, 0x1a, 0xe6, 0x85, 0x2a, 0x9f, 0x89, 0xb0, 0xe7, 0x18, 0x32, 0x0d, 0xf3, 0x7d, 0x99,
	0xc8, 0x59, 0x08, 0x8e, 0x21, 0x13, 0x6f, 0x23, 0xf2, 0xeb, 0x79, 0x22, 0x67, 0x67, 0x73, 0x1e,
	0xf6, 0x89, 0x6c, 0x20, 0x86, 0x3f, 0x9b, 0x73, 0x39, 0xc3, 0x03, 0xf7, 0x1d, 0xef, 0x90, 0xe8,
	0xdf, 0x96, 0x69, 0x55, 0xca, 0xfb, 0xad, 0xc8, 0xb2, 0x37, 0xee, 0xa3, 0xa7, 0xd0, 0x59, 0xfd,
	0x4c, 0x15, 0x6d, 0x91, 0xfe, 0x51, 0x43, 0x7f, 0xc3, 0xc5, 0xef, 0xb6, 0x8e, 0x56, 0x14, 0x1c,
	0x42, 0x57, 0xe4, 0xa7, 0x62, 0x76, 0x23, 0x24, 0x15, 0xa7, 0xcb, 0x2a, 0xdb, 0x5c, 0xf3, 0x6a,
	0x25, 0xb9, 0xa6, 0x9e, 0xa2, 0x6b, 0x92, 0x89, 0x8c, 0xed, 0x83, 0xd0, 0xa7, 0x74, 0x3b, 0x13,
	0xfd, 0xa5, 0xee, 0x7a, 0x1d, 0x2a, 0x76, 0x65, 0x07, 0x0f, 0xc1, 0x5f, 0x52, 0xa2, 0xf7, 0x88,
	0x30, 0x06, 0xa2, 0x0b, 0x4a, 0x72, 0xd7, 0xa0, 0x64, 0xa0, 0x9f, 0x19, 0x66, 0x14, 0xe5, 0x3d,
	0xe3, 0xc7, 0xd9, 0xc8, 0xe5, 0x2e, 0xc5, 0x60, 0x38, 0x67, 0x47, 0x7f, 0x7a, 0xf0, 0xc8, 0xb6,
	0x3b, 0x97, 0xd3, 0x4b, 0x55, 0x14, 0x5c, 0xaf, 0xe9, 0xa6, 0x5f, 0x41, 0x0f, 0x87, 0xf6, 0x54,
	0x64, 0x59, 0x6e, 0xbb, 0x3d, 0xdc, 0xd5, 0xed, 0x98, 0x75, 0x56, 0x4b, 0xdf, 0x45, 0x93, 0x47,
	0xcf, 0xe1, 0x7d, 0xb3, 0xd9, 0x62, 0xf4, 0x7b, 0x55, 0x16, 0x78, 0x8d, 0x7a, 0x5b, 0x79, 0xcd,
	0x6d, 0x85, 0xc9, 0x5d, 0x62, 0x5f, 0x5f, 0xb8, 0xb3, 0x9c, 0x19, 0x9d, 0xc0, 0xbe, 0x75, 0x63,
	0xd6, 0x18, 0xce, 0x46, 0x22, 0xa7, 0x66, 0x0a, 0xc8, 0x8f, 0x99, 0x8d, 0x06, 0x16, 0xa5, 0x6e,
	0xa9, 0xc6, 0x66, 0x64, 0x82, 0xcf, 0xe0, 0x3d, 0xdc, 0xd1, 0xe3, 0x6a, 0xc4, 0xaa, 0x0f, 0xb7,
	0xf0, 0xe0, 0x53, 0x38, 0x68, 0x0e, 0x1a, 0x45, 0x84, 0xca, 0x7b, 0x68, 0x14, 0x43, 0xdf, 0x1e,
	0xb2, 0xcc, 0x92, 0x75, 0xf0, 0x25, 0x74, 0x14, 0xdd, 0xd1, 0xe6, 0xf9, 0xe8, 0x5e, 0x9e, 0x37,
	0xf2, 0xc0, 0xac, 0x96, 0xca, 0x8a, 0xab, 0x0c, 0xbf, 0x33, 0xc7, 0x54, 0x76, 0x34, 0x36, 0x8f,
	0x8b, 0xe6, 0xf1, 0x2b, 0x91, 0xa8, 0x31, 0x97, 0xaf, 0xcd, 0xde, 0x11, 0xf4, 0x50, 0x52, 0xaf,
	0x51, 0x9f, 0xd5, 0x40, 0xc4, 0xcc, 0x32, 0xd4, 0x3c, 0xce, 0xb8, 0x3c, 0xaf, 0xf2, 0xf4, 0x5a,
	0x6f, 0x11, 0xec, 0x53, 0xf2, 0xf3, 0x4b, 0x2e, 0x67, 0xc5, 0xdc, 0x15, 0xbf, 0x89, 0x45, 0x7f,
	0xfb, 0xd5, 0x86, 0xd5, 0x3c, 0x55, 0x52, 0xf2, 0xb4, 0x78, 0xab, 0x27, 0xe0, 0x5b, 0x00, 0x4c,
	0x1f, 0xd7, 0x76, 0x7c, 0x31, 0x75, 0x4f, 0x76, 0xb5, 0x68, 0xad, 0x62, 0x8d, 0x2f, 0x82, 0x6f,
	0xa0, 0x8b, 0x15, 0xa4, 0xaf, 0xdb, 0x74, 0xe2, 0xd1, 0xae, 0xaf, 0x9d, 0x86, 0x55, 0xea, 0xe0,
	0x13, 0x18, 0x54, 0xa1, 0xff, 0x98, 0xdb, 0x79, 0x1f, 0xb0, 0x4d, 0x10, 0x53, 0xfb, 0xc7, 0xbc,
	0x1c, 0x67, 0x6a, 0xc5, 0x35, 0xcd, 0xbd, 0xcf, 0x6a, 0x00, 0x67, 0x20, 0x49, 0x0b, 0xf1, 0x1b,
	0x27, 0x07, 0x1d, 0x72, 0xd0, 0x40, 0x90, 0xc7, 0xf3, 0xae, 0x8b, 0xa4, 0x28, 0x73, 0xbb, 0x02,
	0x1a, 0x88, 0x59, 0xe2, 0x72, 0xc1, 0xf5, 0x98, 0x4b, 0xbb, 0x0b, 0x6a, 0x20, 0x78, 0x0a, 0x3e,
	0xf5, 0x3c, 0x2d, 0x83, 0xfe, 0xc9, 0x87, 0xdb, 0x1d, 0x45, 0x34, 0x33, 0xaa, 0xe0, 0x73, 0xe8,
	0x98, 0x16, 0xa5, 0x05, 0xb1, 0x3d, 0xe9, 0xd5, 0x38, 0x30, 0xab, 0xc3, 0x62, 0xd9, 0x9e, 0xed,
	0x6f, 0x15, 0xab, 0xd1, 0xdb, 0x55, 0xb7, 0x3e, 0x87, 0xc1, 0xc6, 0x7b, 0x4a, 0xcb, 0x7c, 0xc7,
	0x03, 0x7a, 0xef, 0xd9, 0x65, 0x9b, 0x5f, 0x05, 0x5f, 0xc0, 0x9e, 0xed, 0xe8, 0x70, 0x40, 0x05,
	0xff, 0x68, 0xfb, 0x5c, 0x2b, 0x60, 0x4e, 0x19, 0x7c, 0x0d, 0x9d, 0x65, 0x22, 0x2e, 0xb9, 0x0c,
	0x0f, 0x76, 0xbe, 0xda, 0xf7, 0xdb, 0x9b, 0x59, 0x79, 0x54, 0xc2, 0xff, 0x6d, 0x50, 0x22, 0xbf,
	0x92, 0x97, 0x42, 0xf2, 0x77, 0xf5, 0x1b, 0xcb, 0x3c, 0x24, 0x57, 0x32, 0x13, 0x92, 0xd7, 0x0f,
	0x89, 0xb1, 0x4f, 0x5b, 0xe7, 0xed, 0xc9, 0xff, 0x26, 0xde, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x4e, 0xb2, 0x58, 0x5b, 0x21, 0x0a, 0x00, 0x00,
}
