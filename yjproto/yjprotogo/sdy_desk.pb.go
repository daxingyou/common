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

// Ignoring public import of sdy_enum_Option from sdy_base.proto

// Ignoring public import of sdy_enum_roomLevel from sdy_base.proto

// Ignoring public import of sdy_enum_flowers from sdy_base.proto

// Ignoring public import of sdy_bc_opening from sdy_play.proto

// Ignoring public import of sdy_dealCards from sdy_play.proto

// Ignoring public import of sdy_bc_playerPokers from sdy_play.proto

// Ignoring public import of sdy_bc_jiaoFen from sdy_play.proto

// Ignoring public import of sdy_req_jiaoFen from sdy_play.proto

// Ignoring public import of sdy_ack_jiaoFen from sdy_play.proto

// Ignoring public import of sdy_bc_jiaoFenResult from sdy_play.proto

// Ignoring public import of sdy_bc_startPlay from sdy_play.proto

// Ignoring public import of sdy_req_outCards from sdy_play.proto

// Ignoring public import of sdy_ack_outCards from sdy_play.proto

// Ignoring public import of sdy_bc_overTurn from sdy_play.proto

// Ignoring public import of sdy_bc_whatPai from sdy_play.proto

// Ignoring public import of sdy_bc_scorePai from sdy_play.proto

// Ignoring public import of sdy_bc_gameInfo from sdy_play.proto

// Ignoring public import of sdy_bc_dingZhu from sdy_play.proto

// Ignoring public import of sdy_req_dingZhu from sdy_play.proto

// Ignoring public import of sdy_bc_dingZhuResult from sdy_play.proto

// Ignoring public import of sdy_ack_dingZhu from sdy_play.proto

// Ignoring public import of sdy_bc_huanDi from sdy_play.proto

// Ignoring public import of sdy_req_huanDi from sdy_play.proto

// Ignoring public import of sdy_ack_huanDi from sdy_play.proto

// Ignoring public import of sdy_req_changeDesk from sdy_play.proto

// Ignoring public import of sdy_ack_changeDesk from sdy_play.proto

// Ignoring public import of sdy_enum_jScore from sdy_play.proto

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
	BaseNum          *int32                `protobuf:"varint,9,opt,name=baseNum" json:"baseNum,omitempty"`
	Ready            *SdyReReady           `protobuf:"bytes,10,opt,name=ready" json:"ready,omitempty"`
	HuanDi           *SdyReHuanDi          `protobuf:"bytes,11,opt,name=huanDi" json:"huanDi,omitempty"`
	OutPai           *SdyRePlay            `protobuf:"bytes,12,opt,name=outPai" json:"outPai,omitempty"`
	CurrentResult    *SdyBcCurrentResult   `protobuf:"bytes,13,opt,name=currentResult" json:"currentResult,omitempty"`
	JiaoFen          []*SdyReJiaoFen       `protobuf:"bytes,14,rep,name=jiaoFen" json:"jiaoFen,omitempty"`
	PaiLen           []*SdyReLenHandPokers `protobuf:"bytes,15,rep,name=paiLen" json:"paiLen,omitempty"`
	CanJScore        []SdyEnumJScore       `protobuf:"varint,16,rep,name=canJScore,enum=yjprotogo.SdyEnumJScore" json:"canJScore,omitempty"`
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

func (m *SdyBcReconnectInfo) GetBaseNum() int32 {
	if m != nil && m.BaseNum != nil {
		return *m.BaseNum
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

func (m *SdyBcReconnectInfo) GetCanJScore() []SdyEnumJScore {
	if m != nil {
		return m.CanJScore
	}
	return nil
}

// 用于广播玩家的状态
type SdyBcIsOnLine struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32            `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	UserState        *SdyEnumUserStatus `protobuf:"varint,3,opt,name=userState,enum=yjprotogo.SdyEnumUserStatus" json:"userState,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
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

func (m *SdyBcIsOnLine) GetUserState() SdyEnumUserStatus {
	if m != nil && m.UserState != nil {
		return *m.UserState
	}
	return SdyEnumUserStatus_SDY_U_NORMAL
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
<<<<<<< HEAD
	// 858 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0x5d, 0x4f, 0x33, 0x45,
	0x14, 0xb6, 0x2d, 0xdb, 0x8f, 0x53, 0x5a, 0xde, 0x77, 0x05, 0x5c, 0x09, 0x51, 0xb3, 0x51, 0x42,
	0xfc, 0xa8, 0x49, 0x25, 0x5e, 0x79, 0x23, 0x10, 0x28, 0x8a, 0xd2, 0x48, 0x14, 0x13, 0x2f, 0x9a,
	0x61, 0x77, 0x68, 0x87, 0x6e, 0x67, 0xea, 0xce, 0x2e, 0x58, 0xff, 0x8d, 0x3f, 0xc6, 0xbf, 0xe2,
	0xef, 0xf0, 0xcc, 0x99, 0xd9, 0xb6, 0xd0, 0x5e, 0xc0, 0x7b, 0x41, 0xd8, 0x79, 0xe6, 0x7c, 0x3e,
	0xf3, 0x9c, 0x53, 0x68, 0xeb, 0x78, 0x36, 0x88, 0xb9, 0x1e, 0x77, 0xa6, 0xa9, 0xca, 0x94, 0xdf,
	0x98, 0xdd, 0xd3, 0xc7, 0x50, 0xed, 0xbd, 0x1f, 0xa9, 0xc9, 0x44, 0xc9, 0x41, 0x94, 0x08, 0x2e,
	0x33, 0x7b, 0xbf, 0x47, 0xf6, 0xb7, 0x4c, 0xf3, 0xe5, 0xf3, 0x34, 0x61, 0x33, 0x7b, 0x0e, 0xcf,
	0xa1, 0x65, 0x90, 0x94, 0xff, 0x89, 0x7f, 0x2c, 0x9e, 0xf9, 0x07, 0x50, 0x1d, 0xe1, 0x07, 0x4f,
	0x83, 0xd2, 0x27, 0xa5, 0xc3, 0x66, 0x77, 0xb7, 0x33, 0xcf, 0xd0, 0xe9, 0x9b, 0xff, 0x3d, 0xba,
	0xf5, 0xdb, 0x50, 0xcd, 0x35, 0x4f, 0x2f, 0xe2, 0xa0, 0x8c, 0x76, 0xad, 0xb0, 0x67, 0x03, 0xb1,
	0x68, 0xfc, 0xca, 0x40, 0x5b, 0x50, 0xb3, 0x81, 0x34, 0x46, 0xaa, 0x60, 0xa4, 0x07, 0xd8, 0x2e,
	0x8a, 0x1e, 0x3c, 0x0a, 0x79, 0xa2, 0x84, 0xbc, 0x90, 0x77, 0xea, 0xc5, 0x01, 0x7d, 0x80, 0x4c,
	0x65, 0x2c, 0xb9, 0x8e, 0x54, 0xca, 0xa9, 0xba, 0x8a, 0xc1, 0x52, 0x95, 0xcb, 0xd8, 0x62, 0x15,
	0xc2, 0xb0, 0x03, 0x0c, 0xff, 0x73, 0x3e, 0x09, 0x36, 0xf0, 0xec, 0x85, 0xff, 0x94, 0x5d, 0xe2,
	0x68, 0x10, 0xe5, 0x69, 0x8a, 0x1c, 0xfe, 0xc2, 0x75, 0x9e, 0x64, 0xfe, 0x11, 0x34, 0x97, 0xea,
	0xc0, 0xec, 0x15, 0xcc, 0xfe, 0xf1, 0x52, 0xf6, 0xb5, 0xe5, 0x6e, 0xc3, 0x66, 0x11, 0xc6, 0x64,
	0xa6, 0x42, 0x3c, 0x53, 0xc8, 0xad, 0x62, 0x69, 0xac, 0x4f, 0x54, 0x9e, 0x51, 0x21, 0x84, 0xdd,
	0x29, 0x95, 0xf5, 0xd5, 0x98, 0xa7, 0x1a, 0x8b, 0xa9, 0x20, 0x86, 0xde, 0xa3, 0x9c, 0xc9, 0x53,
	0xe1, 0x50, 0x8f, 0xd0, 0xb7, 0xd0, 0x30, 0x79, 0x7e, 0x63, 0x49, 0xce, 0x83, 0x2a, 0x75, 0x81,
	0x90, 0xd0, 0x37, 0x4c, 0x0e, 0x7f, 0x54, 0x79, 0x50, 0x43, 0xa8, 0x6e, 0x18, 0x15, 0xba, 0xaf,
	0xfa, 0x4c, 0x04, 0xf5, 0x05, 0x80, 0xf7, 0xa7, 0x22, 0x68, 0x2c, 0x80, 0x73, 0x8c, 0x3f, 0x0c,
	0x80, 0x00, 0x2c, 0x41, 0xe8, 0xeb, 0x11, 0x02, 0x27, 0x23, 0x1e, 0x34, 0x17, 0x18, 0x1e, 0xe5,
	0xd0, 0x44, 0xda, 0x34, 0x58, 0xf8, 0x5f, 0xc9, 0x3e, 0x33, 0x75, 0x7b, 0x2b, 0x92, 0xe4, 0xc5,
	0xaf, 0xf2, 0x19, 0xb2, 0xfd, 0x3b, 0xf1, 0x57, 0x26, 0xbb, 0x9d, 0x25, 0xbb, 0x1b, 0x2e, 0xfe,
	0x72, 0xac, 0xbd, 0x81, 0xba, 0xd0, 0xc7, 0x62, 0x78, 0x23, 0x24, 0xb1, 0xe3, 0x6a, 0xbd, 0x7a,
	0x94, 0x98, 0x61, 0xa3, 0x00, 0x1c, 0xcf, 0xc8, 0x8a, 0xa1, 0x00, 0x7d, 0xa2, 0xa2, 0xcc, 0x2a,
	0x31, 0xda, 0x02, 0x6f, 0x4a, 0xfd, 0xd7, 0x8a, 0xe3, 0x98, 0xba, 0xaf, 0xd3, 0x11, 0xed, 0x87,
	0xa6, 0x77, 0x63, 0xd0, 0x28, 0x10, 0x5d, 0x34, 0x0f, 0x24, 0x06, 0x09, 0x3b, 0x4e, 0x0b, 0x5c,
	0xc6, 0x97, 0x2a, 0xcb, 0x78, 0x3a, 0xa3, 0x02, 0xbf, 0x80, 0x86, 0x91, 0xeb, 0x31, 0xf6, 0xae,
	0x9d, 0x14, 0x82, 0x75, 0x52, 0x20, 0x72, 0x5e, 0xac, 0x81, 0xf0, 0x08, 0xde, 0xda, 0x39, 0x1c,
	0x98, 0xe8, 0x57, 0x79, 0x86, 0xc5, 0x2d, 0xcd, 0x98, 0xe1, 0xb6, 0x65, 0x3a, 0x9f, 0x1a, 0x39,
	0xb8, 0xa1, 0xf3, 0xc2, 0x4f, 0x61, 0xd3, 0x79, 0xd9, 0x99, 0x33, 0xaa, 0x61, 0x32, 0xb6, 0x9a,
	0x21, 0x37, 0x54, 0x4d, 0xf8, 0x7d, 0x31, 0xe3, 0x03, 0x2b, 0x29, 0x3f, 0x80, 0x37, 0x66, 0x85,
	0x9c, 0xcd, 0x45, 0x57, 0x98, 0xfa, 0xbb, 0xd0, 0x5e, 0x96, 0x1d, 0x25, 0x32, 0x21, 0x7e, 0x82,
	0xa6, 0x0b, 0x61, 0x76, 0x87, 0xff, 0x25, 0x54, 0x15, 0x95, 0xe8, 0x18, 0xd8, 0x7f, 0xc6, 0xc0,
	0xd3, 0x36, 0x0c, 0xbb, 0x66, 0xee, 0x8c, 0xbd, 0x0d, 0xf7, 0x8d, 0xdd, 0x63, 0x68, 0x76, 0x2f,
	0x98, 0x3a, 0xe3, 0x72, 0xa5, 0x55, 0x94, 0xb5, 0xb9, 0x5a, 0xcc, 0xb0, 0x17, 0x7e, 0x67, 0xc7,
	0x13, 0x9d, 0x12, 0x2e, 0x7b, 0xf3, 0x36, 0x57, 0x5c, 0x91, 0x04, 0x62, 0x49, 0x5f, 0xa2, 0x26,
	0xb2, 0x91, 0xf3, 0xfe, 0x77, 0x63, 0x3e, 0xdd, 0x29, 0x8f, 0x94, 0x94, 0x3c, 0xca, 0x5e, 0xb5,
	0x56, 0xba, 0x00, 0xa6, 0x77, 0x4c, 0x64, 0x45, 0x6c, 0xfa, 0xfe, 0x68, 0xdd, 0xcb, 0x2f, 0xac,
	0xfc, 0x0e, 0xd4, 0x0d, 0xd1, 0xe4, 0x51, 0xa1, 0xe8, 0xfb, 0xeb, 0x3c, 0x0a, 0x1b, 0x7f, 0x07,
	0x5a, 0xf3, 0xe2, 0x7e, 0xd5, 0x4e, 0xf1, 0x44, 0xc6, 0xdf, 0xa3, 0xfc, 0x2c, 0x51, 0x8f, 0x08,
	0x79, 0x85, 0x86, 0x58, 0x94, 0x89, 0x07, 0x4e, 0x66, 0x55, 0x32, 0x43, 0xcc, 0x44, 0xba, 0xce,
	0x58, 0x96, 0x6b, 0x27, 0x7d, 0xda, 0x18, 0x12, 0xd9, 0x40, 0x92, 0x9d, 0xfc, 0x0f, 0xc0, 0x23,
	0xb5, 0x90, 0xf6, 0x9b, 0xdd, 0x0f, 0x56, 0xdf, 0xce, 0x8a, 0xe9, 0x10, 0x89, 0x21, 0x2d, 0xd0,
	0x48, 0xac, 0xca, 0x7c, 0xa1, 0xa7, 0x83, 0xb9, 0x1c, 0x9a, 0x2b, 0x14, 0x2e, 0xcb, 0xe6, 0x5b,
	0x68, 0x3d, 0xd9, 0xac, 0xb4, 0x54, 0xd6, 0xac, 0xd2, 0xe7, 0x0b, 0xf8, 0x73, 0xa8, 0x39, 0x9d,
	0x04, 0x2d, 0xe2, 0xfd, 0xc3, 0xd5, 0x04, 0x85, 0x90, 0xbe, 0x86, 0xea, 0x94, 0x09, 0x7c, 0xfa,
	0xa0, 0xbd, 0x76, 0x4f, 0xaf, 0xc8, 0xe7, 0x2b, 0x68, 0x44, 0x4c, 0xfe, 0x60, 0x95, 0xb6, 0x85,
	0x3e, 0xed, 0xee, 0xde, 0x33, 0x1f, 0x2e, 0xf3, 0xc9, 0xe0, 0x9e, 0x2c, 0xc2, 0x3f, 0x60, 0xcb,
	0xd5, 0x88, 0x5b, 0x49, 0x5e, 0x0a, 0xc9, 0xdf, 0xf5, 0x27, 0xd3, 0xee, 0xba, 0x2b, 0x99, 0x60,
	0x0c, 0xbb, 0xeb, 0x8e, 0xcb, 0xbd, 0x4a, 0xff, 0xbd, 0x7e, 0xa9, 0x5f, 0xfe, 0x3f, 0x00, 0x00,
	0xff, 0xff, 0xf6, 0x3a, 0x8f, 0xb0, 0xe8, 0x07, 0x00, 0x00,
=======
	// 1042 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdf, 0x6e, 0x1b, 0xc5,
	0x17, 0xfe, 0xad, 0xdd, 0x75, 0xec, 0x71, 0x9c, 0xf6, 0xb7, 0xb4, 0x65, 0x89, 0xa2, 0x62, 0xad,
	0x10, 0xb2, 0x90, 0x1a, 0x21, 0x83, 0xa0, 0x17, 0x08, 0x89, 0xa4, 0x35, 0x09, 0x8d, 0x88, 0x35,
	0x11, 0x04, 0xae, 0x56, 0x9b, 0xf5, 0xd4, 0x9e, 0x78, 0x33, 0x63, 0x66, 0x76, 0x31, 0xe6, 0x35,
	0xb8, 0xe0, 0x8e, 0x17, 0xe0, 0x1d, 0x78, 0x32, 0x2e, 0xd0, 0x39, 0x33, 0xb3, 0xbb, 0xfe, 0xd3,
	0x8b, 0x56, 0xbd, 0x4a, 0xce, 0xf7, 0x7d, 0x73, 0xe6, 0x9c, 0x33, 0xe7, 0x9c, 0x35, 0x39, 0xd0,
	0x93, 0x55, 0x3c, 0x61, 0x7a, 0x7e, 0xbc, 0x50, 0x32, 0x97, 0x41, 0x67, 0x75, 0x8b, 0xff, 0x4c,
	0xe5, 0xe1, 0x7b, 0xa9, 0xbc, 0xbb, 0x93, 0x22, 0x4e, 0x33, 0xce, 0x44, 0x6e, 0xf8, 0x43, 0xd4,
	0xdf, 0x24, 0x9a, 0xd5, 0xed, 0x45, 0x96, 0xac, 0x8c, 0x1d, 0x5d, 0x93, 0x1e, 0x20, 0x8a, 0xfd,
	0x12, 0x2b, 0x96, 0x4c, 0x56, 0xc1, 0x31, 0x69, 0xcd, 0x58, 0x32, 0x61, 0x2a, 0xf4, 0xfa, 0xde,
	0xa0, 0x3b, 0x7c, 0x7c, 0x5c, 0xde, 0x70, 0x3c, 0x86, 0xbf, 0x67, 0xc8, 0x52, 0xab, 0x0a, 0x1e,
	0x93, 0x56, 0xa1, 0x99, 0x3a, 0x9f, 0x84, 0x8d, 0xbe, 0x37, 0xe8, 0x51, 0x6b, 0x45, 0x3f, 0x1b,
	0xc7, 0x49, 0x3a, 0x7f, 0x4b, 0xc7, 0x21, 0xd9, 0x33, 0xae, 0x74, 0xd8, 0xe8, 0x37, 0x07, 0x3d,
	0xea, 0xcc, 0xe8, 0x2f, 0x8f, 0x3c, 0x74, 0x69, 0xc5, 0x4b, 0x2e, 0x4e, 0x25, 0x17, 0xe7, 0xe2,
	0x95, 0x7c, 0xe3, 0x2b, 0x9e, 0x10, 0x92, 0xcb, 0x3c, 0xc9, 0xae, 0x52, 0xa9, 0x18, 0xc6, 0xdf,
	0xa4, 0x35, 0x04, 0x78, 0x25, 0x0b, 0x31, 0x31, 0x7c, 0xd3, 0xf0, 0x15, 0x02, 0xb9, 0x2f, 0xb9,
	0xf8, 0xbe, 0xb8, 0x0b, 0xef, 0xf5, 0xbd, 0x81, 0x4f, 0xad, 0x15, 0xfd, 0xdd, 0xb4, 0x01, 0xa6,
	0x71, 0x5a, 0x28, 0xc5, 0x44, 0x4e, 0x99, 0x2e, 0xb2, 0x3c, 0xf8, 0x86, 0x74, 0x6b, 0xf1, 0x86,
	0x5e, 0xbf, 0x39, 0xe8, 0x0e, 0x3f, 0xac, 0x45, 0xb9, 0x2b, 0x2d, 0x5a, 0x3f, 0x13, 0x44, 0x64,
	0xdf, 0xf9, 0x84, 0x40, 0x30, 0x6a, 0x9f, 0xae, 0x61, 0x10, 0xf7, 0x8d, 0x4c, 0xd4, 0x44, 0x9f,
	0xca, 0x22, 0xc7, 0xb8, 0x7d, 0x5a, 0x43, 0x80, 0x7f, 0x25, 0x65, 0x3e, 0x96, 0x73, 0xa6, 0x74,
	0x78, 0xaf, 0xdf, 0x04, 0xbe, 0x42, 0xe0, 0x8e, 0x59, 0x91, 0x88, 0xe7, 0xdc, 0x2a, 0x7c, 0x54,
	0xac, 0x61, 0xc1, 0x11, 0xe9, 0x40, 0xa0, 0x3f, 0x26, 0x59, 0xc1, 0xc2, 0x16, 0x96, 0xa6, 0x02,
	0x80, 0xe5, 0xfa, 0x3a, 0x11, 0xd3, 0x97, 0xb2, 0x08, 0xf7, 0xfa, 0xde, 0xa0, 0x4d, 0x2b, 0x00,
	0x9e, 0x96, 0xeb, 0xb1, 0x1c, 0x27, 0x3c, 0x6c, 0x23, 0xe7, 0x4c, 0xc3, 0xbc, 0x94, 0xc5, 0x73,
	0x1e, 0x76, 0x1c, 0x83, 0xa6, 0x61, 0xbe, 0x2d, 0x12, 0x31, 0x0d, 0x89, 0x63, 0xd0, 0x84, 0x6c,
	0xb8, 0xbe, 0x9a, 0x25, 0x62, 0x7a, 0x3a, 0x63, 0x61, 0x17, 0xc9, 0x1a, 0x62, 0xf8, 0xd3, 0x19,
	0x13, 0x53, 0xb8, 0x70, 0xdf, 0xf1, 0x0e, 0x89, 0xfe, 0x69, 0x98, 0x56, 0xc5, 0xba, 0xdf, 0xf0,
	0x2c, 0x7b, 0xe3, 0x3e, 0x7a, 0x4a, 0x5a, 0xcb, 0x9f, 0xf0, 0x45, 0x1b, 0xa8, 0x7f, 0x54, 0xd3,
	0x5f, 0x33, 0xfe, 0x9b, 0x7d, 0x47, 0x2b, 0x0a, 0x0e, 0x49, 0x9b, 0xeb, 0x13, 0x3e, 0xbd, 0xe6,
	0x02, 0x1f, 0xa7, 0x4d, 0x4b, 0xdb, 0xa4, 0x79, 0xb9, 0x14, 0x4c, 0x61, 0x4f, 0x61, 0x9a, 0x68,
	0x02, 0x63, 0xfb, 0x20, 0xf4, 0xb1, 0xdc, 0xce, 0x04, 0x7f, 0xa9, 0x4b, 0xaf, 0x85, 0x8f, 0x5d,
	0xda, 0xc1, 0x43, 0xe2, 0x2f, 0xb0, 0xd0, 0x7b, 0x48, 0x18, 0x03, 0xd0, 0x39, 0x16, 0xb9, 0x6d,
	0x50, 0x34, 0xc0, 0xcf, 0x14, 0x2a, 0x0a, 0xf2, 0x8e, 0xf1, 0xe3, 0x6c, 0xe0, 0xb4, 0x2b, 0x31,
	0x31, 0x9c, 0xb3, 0xa3, 0x3f, 0x3c, 0xf2, 0xc8, 0xb6, 0x3b, 0x13, 0x93, 0x0b, 0x99, 0xe7, 0x4c,
	0xad, 0x30, 0xd3, 0x2f, 0x48, 0x07, 0x86, 0xf6, 0x84, 0x67, 0x99, 0xb6, 0xdd, 0x1e, 0xee, 0xea,
	0x76, 0xa8, 0x3a, 0xad, 0xa4, 0xef, 0xa2, 0xc9, 0xa3, 0x17, 0xe4, 0xff, 0x66, 0xb3, 0xc5, 0xe0,
	0xf7, 0xb2, 0xc8, 0x21, 0x8d, 0x6a, 0x5b, 0x79, 0xf5, 0x6d, 0x05, 0xc5, 0x5d, 0x40, 0x5f, 0x9f,
	0xbb, 0xbb, 0x9c, 0x19, 0x0d, 0xc9, 0xbe, 0x75, 0x63, 0xd6, 0x18, 0xcc, 0x46, 0x22, 0x26, 0x66,
	0x0a, 0xd0, 0x8f, 0x99, 0x8d, 0x1a, 0x16, 0xa5, 0x6e, 0xa9, 0xc6, 0x66, 0x64, 0x82, 0x4f, 0xc8,
	0x03, 0xd8, 0xd9, 0xa3, 0x72, 0xc4, 0xca, 0x83, 0x5b, 0x78, 0xf0, 0x31, 0x39, 0xa8, 0x0f, 0x1a,
	0x46, 0x04, 0xca, 0x0d, 0x34, 0x8a, 0x49, 0xd7, 0x5e, 0x02, 0xeb, 0x3c, 0xf8, 0x9c, 0xb4, 0x24,
	0xe6, 0x68, 0xeb, 0x7c, 0xb4, 0x51, 0xe7, 0xb5, 0x3a, 0x50, 0xab, 0xc5, 0x67, 0x85, 0x55, 0x06,
	0xe7, 0xcc, 0x35, 0xa5, 0x1d, 0x8d, 0xcc, 0xc7, 0x46, 0xb1, 0xf8, 0x96, 0x27, 0x72, 0xc4, 0xc4,
	0x6b, 0xab, 0x77, 0x44, 0x3a, 0x20, 0xa9, 0xd6, 0xa8, 0x4f, 0x2b, 0x20, 0xa2, 0x66, 0x19, 0x2a,
	0x16, 0x67, 0x4c, 0x9c, 0x95, 0x75, 0x7a, 0xad, 0xb7, 0x88, 0xec, 0x63, 0xf1, 0xf5, 0x05, 0x13,
	0xd3, 0x7c, 0xe6, 0x1e, 0xbf, 0x8e, 0x45, 0xff, 0xfa, 0xe5, 0x86, 0x55, 0x2c, 0x95, 0x42, 0xb0,
	0x34, 0x7f, 0xab, 0x4f, 0xc0, 0xd7, 0x84, 0x40, 0xf9, 0x98, 0xb2, 0xe3, 0x0b, 0xa5, 0x7b, 0xb2,
	0xab, 0x45, 0x2b, 0x15, 0xad, 0x9d, 0x08, 0x9e, 0x91, 0x36, 0xbc, 0x20, 0x9e, 0x6e, 0xe2, 0x8d,
	0x47, 0xbb, 0x4e, 0x3b, 0x0d, 0x2d, 0xd5, 0xc1, 0x47, 0xa4, 0x57, 0x86, 0xfe, 0x83, 0xb6, 0xf3,
	0xde, 0xa3, 0xeb, 0x20, 0x94, 0xf6, 0xf7, 0x59, 0x31, 0xca, 0xe4, 0x92, 0x29, 0x9c, 0x7b, 0x9f,
	0x56, 0x00, 0xcc, 0x40, 0x92, 0xe6, 0xfc, 0x57, 0x86, 0x0e, 0x5a, 0xe8, 0xa0, 0x86, 0x00, 0x0f,
	0xf7, 0x5d, 0xe5, 0x49, 0x5e, 0x68, 0xbb, 0x02, 0x6a, 0x88, 0x59, 0xe2, 0x62, 0xce, 0xd4, 0x88,
	0x09, 0xbb, 0x0b, 0x2a, 0x00, 0x86, 0x02, 0x82, 0x87, 0xef, 0x9b, 0x59, 0x07, 0xce, 0x0c, 0x9e,
	0x12, 0x1f, 0xa7, 0x01, 0x57, 0x41, 0x77, 0xf8, 0xfe, 0x76, 0xaf, 0x21, 0x4d, 0x8d, 0x2a, 0xf8,
	0x94, 0xb4, 0x4c, 0xf3, 0xe2, 0x76, 0xde, 0xde, 0x01, 0xe5, 0xa0, 0x50, 0xab, 0x83, 0x67, 0xb4,
	0xdd, 0xbc, 0xbf, 0xf5, 0x8c, 0xb5, 0xae, 0x2f, 0xfb, 0xf8, 0x05, 0xe9, 0xad, 0x7d, 0x69, 0xc3,
	0x1e, 0x1e, 0xdb, 0xfa, 0xb4, 0x6e, 0x7c, 0x90, 0xe9, 0xfa, 0xa9, 0xe0, 0x33, 0xb2, 0x67, 0x7b,
	0x3d, 0x3c, 0xc0, 0x56, 0xf8, 0x60, 0xfb, 0x5e, 0x2b, 0xa0, 0x4e, 0x19, 0x7c, 0x49, 0x5a, 0x8b,
	0x84, 0x5f, 0x30, 0x11, 0xde, 0xdf, 0xf9, 0x3d, 0xdf, 0x6c, 0x7c, 0x6a, 0xe5, 0xc1, 0x33, 0xd2,
	0x49, 0x13, 0xf1, 0x9d, 0x19, 0x9b, 0x07, 0xfd, 0xe6, 0xe0, 0x60, 0x78, 0xb8, 0x71, 0x96, 0x89,
	0xe2, 0x2e, 0xbe, 0x45, 0x05, 0xad, 0xc4, 0xd1, 0x9f, 0x1e, 0xb9, 0x6f, 0xf3, 0xe1, 0xfa, 0x52,
	0x5c, 0x70, 0xc1, 0xde, 0xd5, 0x0f, 0xb7, 0xe0, 0x2b, 0xb3, 0xb3, 0xa1, 0x43, 0xcc, 0x6f, 0x9e,
	0x83, 0xad, 0x81, 0xc0, 0xa8, 0x9c, 0xa8, 0xd0, 0xb4, 0x3a, 0x70, 0xd2, 0x38, 0x6b, 0x8e, 0xff,
	0x37, 0xf6, 0xc6, 0x8d, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xeb, 0xe2, 0x39, 0xaa, 0x0a,
	0x00, 0x00,
>>>>>>> 1905a8d1dfb58c418a3623ced4808a60bfecf519
}
