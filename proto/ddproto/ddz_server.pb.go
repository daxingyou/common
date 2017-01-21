// Code generated by protoc-gen-go.
// source: ddz_server.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of common_srv_pokerPai from common_server_poker.proto

// Ignoring public import of ddz_base_roomTypeInfo from ddz_base.proto

// Ignoring public import of ddz_base_playerInfo from ddz_base.proto

// Ignoring public import of ddz_base_playerRateInfo from ddz_base.proto

// Ignoring public import of ddz_base_commonRateInfo from ddz_base.proto

// Ignoring public import of ddz_base_timerInfo from ddz_base.proto

// Ignoring public import of ddz_base_deskInfo from ddz_base.proto

// Ignoring public import of ddz_enum_protoId from ddz_base.proto

// Ignoring public import of ddz_enum_errorCode from ddz_base.proto

// Ignoring public import of ddz_enum_paiType from ddz_base.proto

// Ignoring public import of ddz_enum_actType from ddz_base.proto

// Ignoring public import of ddz_enum_gameStatus from ddz_base.proto

// Ignoring public import of ddz_enum_playerStatus from ddz_base.proto

// Ignoring public import of ddz_enum_roomType from ddz_base.proto

// Ignoring public import of ddz_enum_enterType from ddz_base.proto

// Ignoring public import of ddz_enum_coinRoomLevel from ddz_base.proto

// Ignoring public import of ddz_enum_deskGameStatus from ddz_base.proto

// 打出去的牌
type DdzSrvOutPokerPais struct {
	KeyValue         *int32               `protobuf:"varint,1,opt,name=keyValue" json:"keyValue,omitempty"`
	PokerPais        []*CommonSrvPokerPai `protobuf:"bytes,2,rep,name=pokerPais" json:"pokerPais,omitempty"`
	Type             *int32               `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`
	IsBomb           *bool                `protobuf:"varint,4,opt,name=isBomb" json:"isBomb,omitempty"`
	CountDuizi       *int32               `protobuf:"varint,5,opt,name=countDuizi" json:"countDuizi,omitempty"`
	CountSanzhang    *int32               `protobuf:"varint,6,opt,name=countSanzhang" json:"countSanzhang,omitempty"`
	CountSizhang     *int32               `protobuf:"varint,7,opt,name=countSizhang" json:"countSizhang,omitempty"`
	CountYizhang     *int32               `protobuf:"varint,8,opt,name=countYizhang" json:"countYizhang,omitempty"`
	UserId           *uint32              `protobuf:"varint,9,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *DdzSrvOutPokerPais) Reset()                    { *m = DdzSrvOutPokerPais{} }
func (m *DdzSrvOutPokerPais) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvOutPokerPais) ProtoMessage()               {}
func (*DdzSrvOutPokerPais) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *DdzSrvOutPokerPais) GetKeyValue() int32 {
	if m != nil && m.KeyValue != nil {
		return *m.KeyValue
	}
	return 0
}

func (m *DdzSrvOutPokerPais) GetPokerPais() []*CommonSrvPokerPai {
	if m != nil {
		return m.PokerPais
	}
	return nil
}

func (m *DdzSrvOutPokerPais) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *DdzSrvOutPokerPais) GetIsBomb() bool {
	if m != nil && m.IsBomb != nil {
		return *m.IsBomb
	}
	return false
}

func (m *DdzSrvOutPokerPais) GetCountDuizi() int32 {
	if m != nil && m.CountDuizi != nil {
		return *m.CountDuizi
	}
	return 0
}

func (m *DdzSrvOutPokerPais) GetCountSanzhang() int32 {
	if m != nil && m.CountSanzhang != nil {
		return *m.CountSanzhang
	}
	return 0
}

func (m *DdzSrvOutPokerPais) GetCountSizhang() int32 {
	if m != nil && m.CountSizhang != nil {
		return *m.CountSizhang
	}
	return 0
}

func (m *DdzSrvOutPokerPais) GetCountYizhang() int32 {
	if m != nil && m.CountYizhang != nil {
		return *m.CountYizhang
	}
	return 0
}

func (m *DdzSrvOutPokerPais) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 对 desk的统计
type DdzSrvDeskTongJi struct {
	Bombs            []*DdzSrvOutPokerPais `protobuf:"bytes,1,rep,name=bombs" json:"bombs,omitempty"`
	CountQiangDiZhu  *int32                `protobuf:"varint,2,opt,name=countQiangDiZhu" json:"countQiangDiZhu,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *DdzSrvDeskTongJi) Reset()                    { *m = DdzSrvDeskTongJi{} }
func (m *DdzSrvDeskTongJi) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvDeskTongJi) ProtoMessage()               {}
func (*DdzSrvDeskTongJi) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

func (m *DdzSrvDeskTongJi) GetBombs() []*DdzSrvOutPokerPais {
	if m != nil {
		return m.Bombs
	}
	return nil
}

func (m *DdzSrvDeskTongJi) GetCountQiangDiZhu() int32 {
	if m != nil && m.CountQiangDiZhu != nil {
		return *m.CountQiangDiZhu
	}
	return 0
}

// desk
type DdzSrvDesk struct {
	AllPokerPai      []*CommonSrvPokerPai   `protobuf:"bytes,1,rep,name=allPokerPai" json:"allPokerPai,omitempty"`
	DiPokerPai       []*CommonSrvPokerPai   `protobuf:"bytes,2,rep,name=diPokerPai" json:"diPokerPai,omitempty"`
	OutPai           *DdzSrvOutPokerPais    `protobuf:"bytes,3,opt,name=outPai" json:"outPai,omitempty"`
	DizhuPaiUser     *uint32                `protobuf:"varint,4,opt,name=dizhuPaiUser" json:"dizhuPaiUser,omitempty"`
	DiZhuUserId      *uint32                `protobuf:"varint,5,opt,name=diZhuUserId" json:"diZhuUserId,omitempty"`
	ActiveUserId     *uint32                `protobuf:"varint,6,opt,name=activeUserId" json:"activeUserId,omitempty"`
	Tongji           *DdzSrvDeskTongJi      `protobuf:"bytes,7,opt,name=tongji" json:"tongji,omitempty"`
	DdzType          *int32                 `protobuf:"varint,8,opt,name=ddzType" json:"ddzType,omitempty"`
	RoomType         *int32                 `protobuf:"varint,9,opt,name=RoomType" json:"RoomType,omitempty"`
	BoardsCount      *int32                 `protobuf:"varint,10,opt,name=BoardsCount" json:"BoardsCount,omitempty"`
	CapMax           *int64                 `protobuf:"varint,11,opt,name=CapMax" json:"CapMax,omitempty"`
	IsJiaoFen        *bool                  `protobuf:"varint,12,opt,name=IsJiaoFen" json:"IsJiaoFen,omitempty"`
	CommonRateInfo   *DdzBaseCommonRateInfo `protobuf:"bytes,13,opt,name=commonRateInfo" json:"commonRateInfo,omitempty"`
	PlayRate         *int32                 `protobuf:"varint,14,opt,name=playRate" json:"playRate,omitempty"`
	GameStatus       *int32                 `protobuf:"varint,15,opt,name=GameStatus" json:"GameStatus,omitempty"`
	InitRoomCoin     *int64                 `protobuf:"varint,16,opt,name=initRoomCoin" json:"initRoomCoin,omitempty"`
	CurrPlayCount    *int32                 `protobuf:"varint,17,opt,name=currPlayCount" json:"currPlayCount,omitempty"`
	TotalPlayCount   *int32                 `protobuf:"varint,18,opt,name=totalPlayCount" json:"totalPlayCount,omitempty"`
	PlayerNum        *int32                 `protobuf:"varint,19,opt,name=playerNum" json:"playerNum,omitempty"`
	TimeOutSecond    *int32                 `protobuf:"varint,20,opt,name=timeOutSecond" json:"timeOutSecond,omitempty"`
	UserMinCoin      *int64                 `protobuf:"varint,21,opt,name=userMinCoin" json:"userMinCoin,omitempty"`
	UserMaxCoin      *int64                 `protobuf:"varint,22,opt,name=userMaxCoin" json:"userMaxCoin,omitempty"`
	CoinTicket       *int64                 `protobuf:"varint,23,opt,name=coinTicket" json:"coinTicket,omitempty"`
	CoinRoomLv       *DdzEnumCoinRoomLevel  `protobuf:"varint,24,opt,name=coinRoomLv,enum=ddproto.DdzEnumCoinRoomLevel" json:"coinRoomLv,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *DdzSrvDesk) Reset()                    { *m = DdzSrvDesk{} }
func (m *DdzSrvDesk) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvDesk) ProtoMessage()               {}
func (*DdzSrvDesk) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

func (m *DdzSrvDesk) GetAllPokerPai() []*CommonSrvPokerPai {
	if m != nil {
		return m.AllPokerPai
	}
	return nil
}

func (m *DdzSrvDesk) GetDiPokerPai() []*CommonSrvPokerPai {
	if m != nil {
		return m.DiPokerPai
	}
	return nil
}

func (m *DdzSrvDesk) GetOutPai() *DdzSrvOutPokerPais {
	if m != nil {
		return m.OutPai
	}
	return nil
}

func (m *DdzSrvDesk) GetDizhuPaiUser() uint32 {
	if m != nil && m.DizhuPaiUser != nil {
		return *m.DizhuPaiUser
	}
	return 0
}

func (m *DdzSrvDesk) GetDiZhuUserId() uint32 {
	if m != nil && m.DiZhuUserId != nil {
		return *m.DiZhuUserId
	}
	return 0
}

func (m *DdzSrvDesk) GetActiveUserId() uint32 {
	if m != nil && m.ActiveUserId != nil {
		return *m.ActiveUserId
	}
	return 0
}

func (m *DdzSrvDesk) GetTongji() *DdzSrvDeskTongJi {
	if m != nil {
		return m.Tongji
	}
	return nil
}

func (m *DdzSrvDesk) GetDdzType() int32 {
	if m != nil && m.DdzType != nil {
		return *m.DdzType
	}
	return 0
}

func (m *DdzSrvDesk) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *DdzSrvDesk) GetBoardsCount() int32 {
	if m != nil && m.BoardsCount != nil {
		return *m.BoardsCount
	}
	return 0
}

func (m *DdzSrvDesk) GetCapMax() int64 {
	if m != nil && m.CapMax != nil {
		return *m.CapMax
	}
	return 0
}

func (m *DdzSrvDesk) GetIsJiaoFen() bool {
	if m != nil && m.IsJiaoFen != nil {
		return *m.IsJiaoFen
	}
	return false
}

func (m *DdzSrvDesk) GetCommonRateInfo() *DdzBaseCommonRateInfo {
	if m != nil {
		return m.CommonRateInfo
	}
	return nil
}

func (m *DdzSrvDesk) GetPlayRate() int32 {
	if m != nil && m.PlayRate != nil {
		return *m.PlayRate
	}
	return 0
}

func (m *DdzSrvDesk) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *DdzSrvDesk) GetInitRoomCoin() int64 {
	if m != nil && m.InitRoomCoin != nil {
		return *m.InitRoomCoin
	}
	return 0
}

func (m *DdzSrvDesk) GetCurrPlayCount() int32 {
	if m != nil && m.CurrPlayCount != nil {
		return *m.CurrPlayCount
	}
	return 0
}

func (m *DdzSrvDesk) GetTotalPlayCount() int32 {
	if m != nil && m.TotalPlayCount != nil {
		return *m.TotalPlayCount
	}
	return 0
}

func (m *DdzSrvDesk) GetPlayerNum() int32 {
	if m != nil && m.PlayerNum != nil {
		return *m.PlayerNum
	}
	return 0
}

func (m *DdzSrvDesk) GetTimeOutSecond() int32 {
	if m != nil && m.TimeOutSecond != nil {
		return *m.TimeOutSecond
	}
	return 0
}

func (m *DdzSrvDesk) GetUserMinCoin() int64 {
	if m != nil && m.UserMinCoin != nil {
		return *m.UserMinCoin
	}
	return 0
}

func (m *DdzSrvDesk) GetUserMaxCoin() int64 {
	if m != nil && m.UserMaxCoin != nil {
		return *m.UserMaxCoin
	}
	return 0
}

func (m *DdzSrvDesk) GetCoinTicket() int64 {
	if m != nil && m.CoinTicket != nil {
		return *m.CoinTicket
	}
	return 0
}

func (m *DdzSrvDesk) GetCoinRoomLv() DdzEnumCoinRoomLevel {
	if m != nil && m.CoinRoomLv != nil {
		return *m.CoinRoomLv
	}
	return DdzEnumCoinRoomLevel_LV1
}

// 游戏数据
type DdzSrvGameData struct {
	HandPokers       []*CommonSrvPokerPai  `protobuf:"bytes,1,rep,name=handPokers" json:"handPokers,omitempty"`
	OutPaiList       []*DdzSrvOutPokerPais `protobuf:"bytes,2,rep,name=outPaiList" json:"outPaiList,omitempty"`
	OutPai           *DdzSrvOutPokerPais   `protobuf:"bytes,3,opt,name=outPai" json:"outPai,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *DdzSrvGameData) Reset()                    { *m = DdzSrvGameData{} }
func (m *DdzSrvGameData) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvGameData) ProtoMessage()               {}
func (*DdzSrvGameData) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{3} }

func (m *DdzSrvGameData) GetHandPokers() []*CommonSrvPokerPai {
	if m != nil {
		return m.HandPokers
	}
	return nil
}

func (m *DdzSrvGameData) GetOutPaiList() []*DdzSrvOutPokerPais {
	if m != nil {
		return m.OutPaiList
	}
	return nil
}

func (m *DdzSrvGameData) GetOutPai() *DdzSrvOutPokerPais {
	if m != nil {
		return m.OutPai
	}
	return nil
}

type DdzSrvBillBean struct {
	// 斗地主的账单
	Coin             *int64  `protobuf:"varint,1,opt,name=coin" json:"coin,omitempty"`
	LoseUser         *uint32 `protobuf:"varint,2,opt,name=loseUser" json:"loseUser,omitempty"`
	WinUser          *uint32 `protobuf:"varint,3,opt,name=winUser" json:"winUser,omitempty"`
	Desc             *string `protobuf:"bytes,4,opt,name=desc" json:"desc,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DdzSrvBillBean) Reset()                    { *m = DdzSrvBillBean{} }
func (m *DdzSrvBillBean) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvBillBean) ProtoMessage()               {}
func (*DdzSrvBillBean) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{4} }

func (m *DdzSrvBillBean) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *DdzSrvBillBean) GetLoseUser() uint32 {
	if m != nil && m.LoseUser != nil {
		return *m.LoseUser
	}
	return 0
}

func (m *DdzSrvBillBean) GetWinUser() uint32 {
	if m != nil && m.WinUser != nil {
		return *m.WinUser
	}
	return 0
}

func (m *DdzSrvBillBean) GetDesc() string {
	if m != nil && m.Desc != nil {
		return *m.Desc
	}
	return ""
}

type DdzSrvBill struct {
	// 斗地主的账单
	WinCoin          *int64            `protobuf:"varint,1,opt,name=winCoin" json:"winCoin,omitempty"`
	BillBean         []*DdzSrvBillBean `protobuf:"bytes,2,rep,name=billBean" json:"billBean,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *DdzSrvBill) Reset()                    { *m = DdzSrvBill{} }
func (m *DdzSrvBill) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvBill) ProtoMessage()               {}
func (*DdzSrvBill) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{5} }

func (m *DdzSrvBill) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *DdzSrvBill) GetBillBean() []*DdzSrvBillBean {
	if m != nil {
		return m.BillBean
	}
	return nil
}

type DdzSrvUserStatisticsRound struct {
	// 玩家每一局的统计信息
	Round            *int32                `protobuf:"varint,1,opt,name=round" json:"round,omitempty"`
	WinCoin          *int64                `protobuf:"varint,2,opt,name=winCoin" json:"winCoin,omitempty"`
	CountBomb        *int32                `protobuf:"varint,3,opt,name=countBomb" json:"countBomb,omitempty"`
	BomBean          []*DdzSrvOutPokerPais `protobuf:"bytes,4,rep,name=bomBean" json:"bomBean,omitempty"`
	PlayRate         *int32                `protobuf:"varint,5,opt,name=playRate" json:"playRate,omitempty"`
	IsSpring         *bool                 `protobuf:"varint,6,opt,name=isSpring" json:"isSpring,omitempty"`
	IsDizhu          *bool                 `protobuf:"varint,7,opt,name=isDizhu" json:"isDizhu,omitempty"`
	IsWin            *bool                 `protobuf:"varint,8,opt,name=isWin" json:"isWin,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *DdzSrvUserStatisticsRound) Reset()                    { *m = DdzSrvUserStatisticsRound{} }
func (m *DdzSrvUserStatisticsRound) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvUserStatisticsRound) ProtoMessage()               {}
func (*DdzSrvUserStatisticsRound) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{6} }

func (m *DdzSrvUserStatisticsRound) GetRound() int32 {
	if m != nil && m.Round != nil {
		return *m.Round
	}
	return 0
}

func (m *DdzSrvUserStatisticsRound) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *DdzSrvUserStatisticsRound) GetCountBomb() int32 {
	if m != nil && m.CountBomb != nil {
		return *m.CountBomb
	}
	return 0
}

func (m *DdzSrvUserStatisticsRound) GetBomBean() []*DdzSrvOutPokerPais {
	if m != nil {
		return m.BomBean
	}
	return nil
}

func (m *DdzSrvUserStatisticsRound) GetPlayRate() int32 {
	if m != nil && m.PlayRate != nil {
		return *m.PlayRate
	}
	return 0
}

func (m *DdzSrvUserStatisticsRound) GetIsSpring() bool {
	if m != nil && m.IsSpring != nil {
		return *m.IsSpring
	}
	return false
}

func (m *DdzSrvUserStatisticsRound) GetIsDizhu() bool {
	if m != nil && m.IsDizhu != nil {
		return *m.IsDizhu
	}
	return false
}

func (m *DdzSrvUserStatisticsRound) GetIsWin() bool {
	if m != nil && m.IsWin != nil {
		return *m.IsWin
	}
	return false
}

type DdzSrvUserStatistics struct {
	// 玩家统计信息
	RoundBean        []*DdzSrvUserStatisticsRound `protobuf:"bytes,1,rep,name=roundBean" json:"roundBean,omitempty"`
	CountWin         *int32                       `protobuf:"varint,2,opt,name=countWin" json:"countWin,omitempty"`
	CountLose        *int32                       `protobuf:"varint,3,opt,name=countLose" json:"countLose,omitempty"`
	CountSpring      *int32                       `protobuf:"varint,4,opt,name=countSpring" json:"countSpring,omitempty"`
	CountDizhu       *int32                       `protobuf:"varint,5,opt,name=countDizhu" json:"countDizhu,omitempty"`
	CountBomb        *int32                       `protobuf:"varint,6,opt,name=countBomb" json:"countBomb,omitempty"`
	MaxWinCoin       *int64                       `protobuf:"varint,7,opt,name=maxWinCoin" json:"maxWinCoin,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *DdzSrvUserStatistics) Reset()                    { *m = DdzSrvUserStatistics{} }
func (m *DdzSrvUserStatistics) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvUserStatistics) ProtoMessage()               {}
func (*DdzSrvUserStatistics) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{7} }

func (m *DdzSrvUserStatistics) GetRoundBean() []*DdzSrvUserStatisticsRound {
	if m != nil {
		return m.RoundBean
	}
	return nil
}

func (m *DdzSrvUserStatistics) GetCountWin() int32 {
	if m != nil && m.CountWin != nil {
		return *m.CountWin
	}
	return 0
}

func (m *DdzSrvUserStatistics) GetCountLose() int32 {
	if m != nil && m.CountLose != nil {
		return *m.CountLose
	}
	return 0
}

func (m *DdzSrvUserStatistics) GetCountSpring() int32 {
	if m != nil && m.CountSpring != nil {
		return *m.CountSpring
	}
	return 0
}

func (m *DdzSrvUserStatistics) GetCountDizhu() int32 {
	if m != nil && m.CountDizhu != nil {
		return *m.CountDizhu
	}
	return 0
}

func (m *DdzSrvUserStatistics) GetCountBomb() int32 {
	if m != nil && m.CountBomb != nil {
		return *m.CountBomb
	}
	return 0
}

func (m *DdzSrvUserStatistics) GetMaxWinCoin() int64 {
	if m != nil && m.MaxWinCoin != nil {
		return *m.MaxWinCoin
	}
	return 0
}

// user
type DdzSrvUser struct {
	UserId           *uint32               `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	GameData         *DdzSrvGameData       `protobuf:"bytes,2,opt,name=gameData" json:"gameData,omitempty"`
	StatusHLQiang    *int32                `protobuf:"varint,3,opt,name=statusHLQiang" json:"statusHLQiang,omitempty"`
	StatusHLJiao     *int32                `protobuf:"varint,4,opt,name=statusHLJiao" json:"statusHLJiao,omitempty"`
	StatusHLJiaBei   *int32                `protobuf:"varint,5,opt,name=statusHLJiaBei" json:"statusHLJiaBei,omitempty"`
	StatusSCMenZhua  *int32                `protobuf:"varint,6,opt,name=statusSCMenZhua" json:"statusSCMenZhua,omitempty"`
	StatusSCZhuaPai  *int32                `protobuf:"varint,7,opt,name=statusSCZhuaPai" json:"statusSCZhuaPai,omitempty"`
	StatusSCLaDao    *int32                `protobuf:"varint,8,opt,name=statusSCLaDao" json:"statusSCLaDao,omitempty"`
	StatusJDJiao     *int32                `protobuf:"varint,9,opt,name=statusJDJiao" json:"statusJDJiao,omitempty"`
	IsShowPokers     *bool                 `protobuf:"varint,10,opt,name=isShowPokers" json:"isShowPokers,omitempty"`
	Bill             *DdzSrvBill           `protobuf:"bytes,11,opt,name=bill" json:"bill,omitempty"`
	Coin             *int64                `protobuf:"varint,12,opt,name=coin" json:"coin,omitempty"`
	Statistics       *DdzSrvUserStatistics `protobuf:"bytes,13,opt,name=statistics" json:"statistics,omitempty"`
	PlayRate         *int32                `protobuf:"varint,14,opt,name=playRate" json:"playRate,omitempty"`
	JiaoScore        *int32                `protobuf:"varint,15,opt,name=jiaoScore" json:"jiaoScore,omitempty"`
	TimeOutCount     *int32                `protobuf:"varint,16,opt,name=timeOutCount" json:"timeOutCount,omitempty"`
	IsAgent          *bool                 `protobuf:"varint,17,opt,name=isAgent" json:"isAgent,omitempty"`
	Sex              *int32                `protobuf:"varint,18,opt,name=sex" json:"sex,omitempty"`
	RoomCard         *int32                `protobuf:"varint,19,opt,name=roomCard" json:"roomCard,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *DdzSrvUser) Reset()                    { *m = DdzSrvUser{} }
func (m *DdzSrvUser) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvUser) ProtoMessage()               {}
func (*DdzSrvUser) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{8} }

func (m *DdzSrvUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzSrvUser) GetGameData() *DdzSrvGameData {
	if m != nil {
		return m.GameData
	}
	return nil
}

func (m *DdzSrvUser) GetStatusHLQiang() int32 {
	if m != nil && m.StatusHLQiang != nil {
		return *m.StatusHLQiang
	}
	return 0
}

func (m *DdzSrvUser) GetStatusHLJiao() int32 {
	if m != nil && m.StatusHLJiao != nil {
		return *m.StatusHLJiao
	}
	return 0
}

func (m *DdzSrvUser) GetStatusHLJiaBei() int32 {
	if m != nil && m.StatusHLJiaBei != nil {
		return *m.StatusHLJiaBei
	}
	return 0
}

func (m *DdzSrvUser) GetStatusSCMenZhua() int32 {
	if m != nil && m.StatusSCMenZhua != nil {
		return *m.StatusSCMenZhua
	}
	return 0
}

func (m *DdzSrvUser) GetStatusSCZhuaPai() int32 {
	if m != nil && m.StatusSCZhuaPai != nil {
		return *m.StatusSCZhuaPai
	}
	return 0
}

func (m *DdzSrvUser) GetStatusSCLaDao() int32 {
	if m != nil && m.StatusSCLaDao != nil {
		return *m.StatusSCLaDao
	}
	return 0
}

func (m *DdzSrvUser) GetStatusJDJiao() int32 {
	if m != nil && m.StatusJDJiao != nil {
		return *m.StatusJDJiao
	}
	return 0
}

func (m *DdzSrvUser) GetIsShowPokers() bool {
	if m != nil && m.IsShowPokers != nil {
		return *m.IsShowPokers
	}
	return false
}

func (m *DdzSrvUser) GetBill() *DdzSrvBill {
	if m != nil {
		return m.Bill
	}
	return nil
}

func (m *DdzSrvUser) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *DdzSrvUser) GetStatistics() *DdzSrvUserStatistics {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func (m *DdzSrvUser) GetPlayRate() int32 {
	if m != nil && m.PlayRate != nil {
		return *m.PlayRate
	}
	return 0
}

func (m *DdzSrvUser) GetJiaoScore() int32 {
	if m != nil && m.JiaoScore != nil {
		return *m.JiaoScore
	}
	return 0
}

func (m *DdzSrvUser) GetTimeOutCount() int32 {
	if m != nil && m.TimeOutCount != nil {
		return *m.TimeOutCount
	}
	return 0
}

func (m *DdzSrvUser) GetIsAgent() bool {
	if m != nil && m.IsAgent != nil {
		return *m.IsAgent
	}
	return false
}

func (m *DdzSrvUser) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *DdzSrvUser) GetRoomCard() int32 {
	if m != nil && m.RoomCard != nil {
		return *m.RoomCard
	}
	return 0
}

// room
type DdzSrvRoom struct {
	RoomId           *int32 `protobuf:"varint,1,opt,name=roomId" json:"roomId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DdzSrvRoom) Reset()                    { *m = DdzSrvRoom{} }
func (m *DdzSrvRoom) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvRoom) ProtoMessage()               {}
func (*DdzSrvRoom) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{9} }

func (m *DdzSrvRoom) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

// 备份专用...
type DdzSrvBak struct {
	Desk             *DdzSrvDesk   `protobuf:"bytes,1,opt,name=desk" json:"desk,omitempty"`
	Users            []*DdzSrvUser `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *DdzSrvBak) Reset()                    { *m = DdzSrvBak{} }
func (m *DdzSrvBak) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvBak) ProtoMessage()               {}
func (*DdzSrvBak) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{10} }

func (m *DdzSrvBak) GetDesk() *DdzSrvDesk {
	if m != nil {
		return m.Desk
	}
	return nil
}

func (m *DdzSrvBak) GetUsers() []*DdzSrvUser {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*DdzSrvOutPokerPais)(nil), "ddproto.ddz_srv_outPokerPais")
	proto.RegisterType((*DdzSrvDeskTongJi)(nil), "ddproto.ddz_srv_deskTongJi")
	proto.RegisterType((*DdzSrvDesk)(nil), "ddproto.ddz_srv_desk")
	proto.RegisterType((*DdzSrvGameData)(nil), "ddproto.ddz_srv_gameData")
	proto.RegisterType((*DdzSrvBillBean)(nil), "ddproto.ddz_srv_billBean")
	proto.RegisterType((*DdzSrvBill)(nil), "ddproto.ddz_srv_bill")
	proto.RegisterType((*DdzSrvUserStatisticsRound)(nil), "ddproto.ddz_srv_userStatisticsRound")
	proto.RegisterType((*DdzSrvUserStatistics)(nil), "ddproto.ddz_srv_userStatistics")
	proto.RegisterType((*DdzSrvUser)(nil), "ddproto.ddz_srv_user")
	proto.RegisterType((*DdzSrvRoom)(nil), "ddproto.ddz_srv_room")
	proto.RegisterType((*DdzSrvBak)(nil), "ddproto.ddz_srv_bak")
}

var fileDescriptor14 = []byte{
	// 1045 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x56, 0xdb, 0x72, 0x1b, 0x45,
	0x10, 0x45, 0x91, 0x56, 0x5e, 0x8d, 0x2e, 0x76, 0x36, 0xb6, 0x33, 0x49, 0x80, 0xb8, 0x96, 0x3c,
	0xa4, 0x2a, 0x60, 0x88, 0x79, 0x80, 0x57, 0x2c, 0x15, 0x60, 0x97, 0x0c, 0x8a, 0xe5, 0x60, 0x2e,
	0x0f, 0xaa, 0xb1, 0x76, 0xb0, 0x27, 0xd2, 0xee, 0xa8, 0x76, 0x76, 0xe5, 0xcb, 0x07, 0xf1, 0xc6,
	0x5f, 0x50, 0x7c, 0x03, 0x4f, 0x7c, 0x0b, 0xdd, 0xbd, 0x33, 0xba, 0x94, 0x55, 0xd8, 0x79, 0xb2,
	0xf6, 0xcc, 0xa5, 0x4f, 0x77, 0x9f, 0x3e, 0x63, 0xb6, 0x11, 0x45, 0x37, 0x03, 0x23, 0xd3, 0xa9,
	0x4c, 0x77, 0x27, 0xa9, 0xce, 0x74, 0xb0, 0x16, 0x45, 0xf4, 0xe3, 0xe9, 0x93, 0xa1, 0x8e, 0x63,
	0x9d, 0xd8, 0xd5, 0xc1, 0x44, 0x8f, 0xdc, 0x9e, 0xa7, 0x2d, 0x3c, 0x75, 0x26, 0x8c, 0x2c, 0xbe,
	0xc3, 0x7f, 0x4b, 0x6c, 0x93, 0x2e, 0x4a, 0xa7, 0x03, 0x9d, 0x67, 0x3d, 0xdc, 0xda, 0x13, 0xca,
	0x04, 0x1b, 0xcc, 0x1f, 0xc9, 0xeb, 0x9f, 0xc4, 0x38, 0x97, 0xbc, 0xb4, 0x53, 0x7a, 0xe9, 0x05,
	0x9f, 0xb3, 0xda, 0xc4, 0x2d, 0xf3, 0x07, 0x3b, 0xe5, 0x97, 0xf5, 0xbd, 0x0f, 0x77, 0x6d, 0xc8,
	0x5d, 0x17, 0x11, 0xae, 0x71, 0x9b, 0x82, 0x06, 0xab, 0x64, 0xd7, 0x13, 0xc9, 0xcb, 0x74, 0xbc,
	0xc5, 0xaa, 0xca, 0xec, 0xeb, 0xf8, 0x8c, 0x57, 0xe0, 0xdb, 0x0f, 0x02, 0xc6, 0x86, 0x3a, 0x4f,
	0xb2, 0x4e, 0xae, 0x6e, 0x14, 0xf7, 0x68, 0xcf, 0x16, 0x6b, 0x12, 0xd6, 0x17, 0xc9, 0xcd, 0x85,
	0x48, 0xce, 0x79, 0x95, 0xe0, 0x4d, 0xd6, 0x28, 0x60, 0x55, 0xa0, 0x6b, 0x4b, 0xe8, 0x2f, 0x16,
	0xf5, 0x5d, 0x98, 0x1c, 0xf2, 0x3e, 0x88, 0x78, 0x0d, 0xbe, 0x9b, 0xe1, 0x6f, 0x2c, 0x70, 0xf9,
	0x45, 0xd2, 0x8c, 0x4e, 0x74, 0x72, 0x7e, 0xa8, 0x82, 0x4f, 0x99, 0x77, 0x06, 0x54, 0x0c, 0xa4,
	0x86, 0x79, 0x7c, 0x34, 0xcb, 0x63, 0x65, 0x2d, 0x1e, 0xb3, 0x75, 0x8a, 0xf4, 0x46, 0x41, 0x9c,
	0x8e, 0xfa, 0xf5, 0x22, 0x87, 0xfc, 0x21, 0x58, 0xf8, 0xa7, 0xc7, 0x1a, 0x8b, 0xb7, 0x07, 0xaf,
	0x59, 0x5d, 0x8c, 0xc7, 0xee, 0xa4, 0xbd, 0xfd, 0xff, 0xab, 0xf4, 0x05, 0x63, 0x91, 0x9a, 0x9d,
	0xb8, 0x4f, 0x5d, 0x3f, 0x63, 0x55, 0xa4, 0x07, 0xbb, 0xb1, 0xb2, 0x77, 0xb2, 0x87, 0x3a, 0x45,
	0x50, 0xa2, 0x1c, 0x3e, 0xde, 0x42, 0x65, 0xa8, 0xfc, 0xcd, 0xe0, 0x11, 0xab, 0x47, 0x98, 0xc9,
	0xdb, 0xa2, 0x58, 0x1e, 0x81, 0xb0, 0x55, 0x0c, 0x33, 0x35, 0x95, 0x16, 0xad, 0x12, 0xfa, 0x8a,
	0x55, 0x33, 0x28, 0xdb, 0x3b, 0x45, 0x85, 0xaf, 0xef, 0x3d, 0xbb, 0x15, 0x6f, 0xa1, 0xb2, 0xeb,
	0x0c, 0x64, 0x78, 0x73, 0x82, 0x7d, 0x2f, 0x1a, 0x02, 0x42, 0x3a, 0xd6, 0x3a, 0x26, 0xa4, 0x46,
	0x08, 0x84, 0xde, 0xd7, 0x22, 0x8d, 0x4c, 0x1b, 0x8b, 0xca, 0x99, 0xeb, 0x5b, 0x5b, 0x4c, 0x8e,
	0xc4, 0x15, 0xaf, 0xc3, 0x77, 0x39, 0x78, 0xc8, 0x6a, 0x07, 0xe6, 0x50, 0x09, 0xfd, 0xad, 0x4c,
	0x78, 0x83, 0x14, 0xf3, 0x35, 0x6b, 0x15, 0xe5, 0x38, 0x16, 0x99, 0x3c, 0x48, 0x7e, 0xd7, 0xbc,
	0x49, 0x7c, 0x76, 0x96, 0xf8, 0xa0, 0xb8, 0x07, 0xcb, 0xfb, 0x90, 0xc3, 0x64, 0x2c, 0xae, 0xf1,
	0x9b, 0xb7, 0x28, 0x1c, 0xa8, 0xef, 0x3b, 0x11, 0xcb, 0x7e, 0x26, 0xb2, 0xdc, 0xf0, 0x75, 0x27,
	0x28, 0x95, 0xa8, 0x0c, 0xd9, 0xb6, 0xb5, 0x4a, 0xf8, 0x06, 0x11, 0x41, 0x4d, 0xe6, 0x69, 0xda,
	0x83, 0xf3, 0x05, 0xdf, 0x87, 0xb4, 0x79, 0x9b, 0xb5, 0x32, 0x9d, 0x89, 0xf1, 0x1c, 0x0f, 0x08,
	0x07, 0xde, 0x18, 0x4a, 0xa6, 0x3f, 0xe4, 0x31, 0x7f, 0xe4, 0x54, 0x9d, 0xa9, 0x58, 0xfe, 0x98,
	0x67, 0x7d, 0x39, 0xd4, 0x49, 0xc4, 0x37, 0x5d, 0x19, 0x50, 0xa9, 0x47, 0x2a, 0xa1, 0x68, 0x5b,
	0x14, 0xcd, 0x81, 0xe2, 0x8a, 0xc0, 0x6d, 0x02, 0x69, 0x54, 0x54, 0x72, 0xa2, 0x86, 0x23, 0x99,
	0xf1, 0xc7, 0x84, 0x7d, 0x59, 0x60, 0x48, 0xb6, 0x3b, 0xe5, 0x1c, 0xb0, 0xd6, 0xde, 0xf3, 0xa5,
	0x42, 0xc8, 0x24, 0x8f, 0x07, 0xb3, 0x3d, 0x72, 0x2a, 0xc7, 0xe1, 0x1f, 0x25, 0x6b, 0x1b, 0xd0,
	0xb3, 0x73, 0x48, 0xbf, 0x23, 0x32, 0x81, 0x02, 0x84, 0xf9, 0x89, 0x48, 0x30, 0xe6, 0x5e, 0x92,
	0x7d, 0xcd, 0x58, 0x21, 0xc0, 0xae, 0x32, 0x99, 0x95, 0xec, 0x1d, 0x22, 0x7c, 0x3f, 0xcd, 0x86,
	0x6f, 0xe6, 0x3c, 0xcf, 0xd4, 0x78, 0xbc, 0x2f, 0x45, 0x82, 0x76, 0x82, 0xd9, 0x90, 0x1b, 0x95,
	0xb1, 0xa5, 0x63, 0x6d, 0x48, 0xa8, 0x34, 0x8c, 0x4d, 0x54, 0xde, 0xa5, 0x4a, 0x08, 0x28, 0x13,
	0x00, 0x07, 0x40, 0x98, 0x43, 0x12, 0x7c, 0x2d, 0xec, 0xce, 0x47, 0x15, 0xaf, 0xb4, 0xdb, 0xdb,
	0xf3, 0x1b, 0x5f, 0x31, 0xdf, 0xc5, 0xb2, 0x39, 0x3d, 0xb9, 0x45, 0xd2, 0x6d, 0x08, 0xff, 0x2a,
	0xb1, 0x67, 0x0e, 0xc4, 0x86, 0xa1, 0x90, 0xa0, 0x18, 0x6a, 0x68, 0x8e, 0x41, 0x0c, 0x51, 0xd0,
	0x64, 0x5e, 0x8a, 0x3f, 0xac, 0x77, 0x2e, 0x04, 0x7b, 0xe0, 0xe4, 0x4d, 0x96, 0x42, 0x86, 0x58,
	0x18, 0xe4, 0x2e, 0x5b, 0x03, 0x4f, 0xa2, 0xf0, 0x95, 0xfb, 0x94, 0x74, 0x51, 0xd4, 0x9e, 0x1b,
	0x35, 0x65, 0xfa, 0x93, 0x54, 0x59, 0xe7, 0xf4, 0x31, 0xae, 0x32, 0x1d, 0x9c, 0x7e, 0x9a, 0x5d,
	0x1f, 0x79, 0x29, 0x73, 0x0a, 0x34, 0x70, 0x38, 0xfd, 0xf0, 0xef, 0x12, 0xdb, 0x5e, 0x9d, 0x46,
	0xf0, 0x15, 0xab, 0x51, 0x06, 0x44, 0xa8, 0x50, 0xc5, 0x8b, 0x5b, 0x84, 0x56, 0xa5, 0x0e, 0x2c,
	0x28, 0xb5, 0x53, 0x9b, 0xac, 0x37, 0x4b, 0xb6, 0x0b, 0x0d, 0xb3, 0xc9, 0x82, 0xce, 0x0b, 0x4b,
	0x2f, 0xd8, 0x56, 0xdc, 0x50, 0x16, 0x4f, 0x02, 0x11, 0xf6, 0x96, 0xce, 0x52, 0xa1, 0xaa, 0x6e,
	0x5b, 0x2c, 0xae, 0x4e, 0x6d, 0x3d, 0x31, 0xaf, 0x72, 0xf8, 0x4f, 0x79, 0xde, 0x5e, 0x24, 0xb5,
	0xf0, 0x0e, 0x94, 0xac, 0x89, 0xf9, 0x4e, 0xf1, 0xc4, 0x6a, 0x55, 0x77, 0x67, 0x23, 0x01, 0x13,
	0x6b, 0xc8, 0x19, 0xbe, 0xef, 0x92, 0xe7, 0x5b, 0xd2, 0x60, 0x10, 0x0e, 0x46, 0x67, 0xb2, 0xac,
	0xc1, 0x09, 0x16, 0xd0, 0x7d, 0xe9, 0x1e, 0x33, 0x78, 0x35, 0x0a, 0xbc, 0xdf, 0x3e, 0x92, 0x09,
	0x58, 0xad, 0xb0, 0xfc, 0x17, 0x16, 0x10, 0xc5, 0xa1, 0x58, 0x73, 0x46, 0xe1, 0x16, 0xba, 0xa2,
	0x03, 0x01, 0xfc, 0xe5, 0xb0, 0x87, 0x1d, 0x0a, 0x5b, 0x9b, 0xb9, 0x95, 0xe9, 0x5f, 0xe8, 0x4b,
	0x3b, 0xb8, 0x8c, 0xfa, 0xfb, 0x09, 0xab, 0xa0, 0x46, 0xc9, 0x44, 0xeb, 0x7b, 0x5b, 0x2b, 0x05,
	0x3c, 0x9b, 0xa4, 0x86, 0x73, 0x12, 0x33, 0x6b, 0xa1, 0xb5, 0xd4, 0xe7, 0x77, 0x74, 0x7a, 0x85,
	0xa3, 0x42, 0xa3, 0xde, 0x01, 0xbb, 0xfe, 0x50, 0xa7, 0x72, 0x6e, 0xa8, 0xd6, 0xf8, 0x0a, 0x87,
	0xdc, 0x70, 0xb3, 0xa0, 0xcc, 0x37, 0xe7, 0xd2, 0x5a, 0xa9, 0x1f, 0xd4, 0x59, 0xd9, 0xc8, 0x2b,
	0xeb, 0x9f, 0x70, 0x71, 0x8a, 0x06, 0x0c, 0xef, 0x43, 0x61, 0x9f, 0xe1, 0xc7, 0xf3, 0xce, 0xe2,
	0x0a, 0x76, 0x16, 0xff, 0xda, 0xce, 0x7a, 0xe1, 0xcf, 0xf0, 0x92, 0xb9, 0xec, 0xc4, 0x08, 0x2b,
	0x80, 0xcf, 0x11, 0x2d, 0xae, 0xaa, 0x00, 0xbd, 0xd3, 0x2f, 0x98, 0x87, 0x09, 0xb9, 0xff, 0x63,
	0xb6, 0x56, 0xa6, 0xdb, 0xfb, 0xa0, 0x57, 0xfa, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x46, 0xc3,
	0x7d, 0x67, 0x09, 0x00, 0x00,
}
