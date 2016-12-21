// Code generated by protoc-gen-go.
// source: hall_data.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HallEnumProtoId int32

const (
	HallEnumProtoId_HALL_PID_HEARTBEAT        HallEnumProtoId = 0
	HallEnumProtoId_HALL_PID_QUICK_CONN       HallEnumProtoId = 1
	HallEnumProtoId_HALL_PID_QUICK_CONN_ACK   HallEnumProtoId = 2
	HallEnumProtoId_HALL_PID_GAME_LOGIN       HallEnumProtoId = 3
	HallEnumProtoId_HALL_PID_GAME_LOGIN_ACK   HallEnumProtoId = 4
	HallEnumProtoId_HALL_PID_USER_DATA        HallEnumProtoId = 9
	HallEnumProtoId_HALL_PID_USER_DATA_ACK    HallEnumProtoId = 10
	HallEnumProtoId_HALL_PID_DRAW_LOTTERY     HallEnumProtoId = 11
	HallEnumProtoId_HALL_PID_DRAW_LOTTERY_ACK HallEnumProtoId = 12
)

var HallEnumProtoId_name = map[int32]string{
	0:  "HALL_PID_HEARTBEAT",
	1:  "HALL_PID_QUICK_CONN",
	2:  "HALL_PID_QUICK_CONN_ACK",
	3:  "HALL_PID_GAME_LOGIN",
	4:  "HALL_PID_GAME_LOGIN_ACK",
	9:  "HALL_PID_USER_DATA",
	10: "HALL_PID_USER_DATA_ACK",
	11: "HALL_PID_DRAW_LOTTERY",
	12: "HALL_PID_DRAW_LOTTERY_ACK",
}
var HallEnumProtoId_value = map[string]int32{
	"HALL_PID_HEARTBEAT":        0,
	"HALL_PID_QUICK_CONN":       1,
	"HALL_PID_QUICK_CONN_ACK":   2,
	"HALL_PID_GAME_LOGIN":       3,
	"HALL_PID_GAME_LOGIN_ACK":   4,
	"HALL_PID_USER_DATA":        9,
	"HALL_PID_USER_DATA_ACK":    10,
	"HALL_PID_DRAW_LOTTERY":     11,
	"HALL_PID_DRAW_LOTTERY_ACK": 12,
}

func (x HallEnumProtoId) Enum() *HallEnumProtoId {
	p := new(HallEnumProtoId)
	*p = x
	return p
}
func (x HallEnumProtoId) String() string {
	return proto.EnumName(HallEnumProtoId_name, int32(x))
}
func (x *HallEnumProtoId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumProtoId_value, data, "HallEnumProtoId")
	if err != nil {
		return err
	}
	*x = HallEnumProtoId(value)
	return nil
}
func (HallEnumProtoId) EnumDescriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

// 活动类型
type HallEnumEvent int32

const (
	HallEnumEvent_TYPE_TIME HallEnumEvent = 1
	HallEnumEvent_TYPE_NEW  HallEnumEvent = 2
	HallEnumEvent_TYPE_NULL HallEnumEvent = 3
)

var HallEnumEvent_name = map[int32]string{
	1: "TYPE_TIME",
	2: "TYPE_NEW",
	3: "TYPE_NULL",
}
var HallEnumEvent_value = map[string]int32{
	"TYPE_TIME": 1,
	"TYPE_NEW":  2,
	"TYPE_NULL": 3,
}

func (x HallEnumEvent) Enum() *HallEnumEvent {
	p := new(HallEnumEvent)
	*p = x
	return p
}
func (x HallEnumEvent) String() string {
	return proto.EnumName(HallEnumEvent_name, int32(x))
}
func (x *HallEnumEvent) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumEvent_value, data, "HallEnumEvent")
	if err != nil {
		return err
	}
	*x = HallEnumEvent(value)
	return nil
}
func (HallEnumEvent) EnumDescriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

// 活动奖品
type HallEnum_Reward int32

const (
	HallEnum_Reward_RE_EXP  HallEnum_Reward = 1
	HallEnum_Reward_RE_GIFT HallEnum_Reward = 2
)

var HallEnum_Reward_name = map[int32]string{
	1: "RE_EXP",
	2: "RE_GIFT",
}
var HallEnum_Reward_value = map[string]int32{
	"RE_EXP":  1,
	"RE_GIFT": 2,
}

func (x HallEnum_Reward) Enum() *HallEnum_Reward {
	p := new(HallEnum_Reward)
	*p = x
	return p
}
func (x HallEnum_Reward) String() string {
	return proto.EnumName(HallEnum_Reward_name, int32(x))
}
func (x *HallEnum_Reward) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnum_Reward_value, data, "HallEnum_Reward")
	if err != nil {
		return err
	}
	*x = HallEnum_Reward(value)
	return nil
}
func (HallEnum_Reward) EnumDescriptor() ([]byte, []int) { return fileDescriptor12, []int{2} }

// 邮件类型
type HallEnumEmailType int32

const (
	HallEnumEmailType_EMAIL_SYSTEM HallEnumEmailType = 1
	HallEnumEmailType_TYPE_FRIEND  HallEnumEmailType = 2
)

var HallEnumEmailType_name = map[int32]string{
	1: "EMAIL_SYSTEM",
	2: "TYPE_FRIEND",
}
var HallEnumEmailType_value = map[string]int32{
	"EMAIL_SYSTEM": 1,
	"TYPE_FRIEND":  2,
}

func (x HallEnumEmailType) Enum() *HallEnumEmailType {
	p := new(HallEnumEmailType)
	*p = x
	return p
}
func (x HallEnumEmailType) String() string {
	return proto.EnumName(HallEnumEmailType_name, int32(x))
}
func (x *HallEnumEmailType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumEmailType_value, data, "HallEnumEmailType")
	if err != nil {
		return err
	}
	*x = HallEnumEmailType(value)
	return nil
}
func (HallEnumEmailType) EnumDescriptor() ([]byte, []int) { return fileDescriptor12, []int{3} }

// 道具类型
type HallEnumBagItemType int32

const (
	HallEnumBagItemType_TYPE_MEAT   HallEnumBagItemType = 1
	HallEnumBagItemType_TYPE_EGG    HallEnumBagItemType = 2
	HallEnumBagItemType_TYPE_ICE    HallEnumBagItemType = 3
	HallEnumBagItemType_TYPE_BUN    HallEnumBagItemType = 4
	HallEnumBagItemType_TYPE_SHRIMP HallEnumBagItemType = 5
)

var HallEnumBagItemType_name = map[int32]string{
	1: "TYPE_MEAT",
	2: "TYPE_EGG",
	3: "TYPE_ICE",
	4: "TYPE_BUN",
	5: "TYPE_SHRIMP",
}
var HallEnumBagItemType_value = map[string]int32{
	"TYPE_MEAT":   1,
	"TYPE_EGG":    2,
	"TYPE_ICE":    3,
	"TYPE_BUN":    4,
	"TYPE_SHRIMP": 5,
}

func (x HallEnumBagItemType) Enum() *HallEnumBagItemType {
	p := new(HallEnumBagItemType)
	*p = x
	return p
}
func (x HallEnumBagItemType) String() string {
	return proto.EnumName(HallEnumBagItemType_name, int32(x))
}
func (x *HallEnumBagItemType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumBagItemType_value, data, "HallEnumBagItemType")
	if err != nil {
		return err
	}
	*x = HallEnumBagItemType(value)
	return nil
}
func (HallEnumBagItemType) EnumDescriptor() ([]byte, []int) { return fileDescriptor12, []int{4} }

// 道具阶层
type HallEnumBagItemRank int32

const (
	HallEnumBagItemRank_RANK_ONE   HallEnumBagItemRank = 1
	HallEnumBagItemRank_RANK_TWO   HallEnumBagItemRank = 2
	HallEnumBagItemRank_RANK_THREE HallEnumBagItemRank = 3
)

var HallEnumBagItemRank_name = map[int32]string{
	1: "RANK_ONE",
	2: "RANK_TWO",
	3: "RANK_THREE",
}
var HallEnumBagItemRank_value = map[string]int32{
	"RANK_ONE":   1,
	"RANK_TWO":   2,
	"RANK_THREE": 3,
}

func (x HallEnumBagItemRank) Enum() *HallEnumBagItemRank {
	p := new(HallEnumBagItemRank)
	*p = x
	return p
}
func (x HallEnumBagItemRank) String() string {
	return proto.EnumName(HallEnumBagItemRank_name, int32(x))
}
func (x *HallEnumBagItemRank) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumBagItemRank_value, data, "HallEnumBagItemRank")
	if err != nil {
		return err
	}
	*x = HallEnumBagItemRank(value)
	return nil
}
func (HallEnumBagItemRank) EnumDescriptor() ([]byte, []int) { return fileDescriptor12, []int{5} }

// 任务类型
type HallEnumTaskType int32

const (
	HallEnumTaskType_TYPE_MJ  HallEnumTaskType = 1
	HallEnumTaskType_TYPE_DDZ HallEnumTaskType = 2
	HallEnumTaskType_TYPE_ZJH HallEnumTaskType = 3
)

var HallEnumTaskType_name = map[int32]string{
	1: "TYPE_MJ",
	2: "TYPE_DDZ",
	3: "TYPE_ZJH",
}
var HallEnumTaskType_value = map[string]int32{
	"TYPE_MJ":  1,
	"TYPE_DDZ": 2,
	"TYPE_ZJH": 3,
}

func (x HallEnumTaskType) Enum() *HallEnumTaskType {
	p := new(HallEnumTaskType)
	*p = x
	return p
}
func (x HallEnumTaskType) String() string {
	return proto.EnumName(HallEnumTaskType_name, int32(x))
}
func (x *HallEnumTaskType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumTaskType_value, data, "HallEnumTaskType")
	if err != nil {
		return err
	}
	*x = HallEnumTaskType(value)
	return nil
}
func (HallEnumTaskType) EnumDescriptor() ([]byte, []int) { return fileDescriptor12, []int{6} }

// 交易类型
type HallEnumShopType int32

const (
	HallEnumShopType_TYPE_COIN     HallEnumShopType = 1
	HallEnumShopType_TYPE_DIAMOND  HallEnumShopType = 2
	HallEnumShopType_TYPE_EXCHANGE HallEnumShopType = 3
	HallEnumShopType_TYPE_BUY      HallEnumShopType = 4
	HallEnumShopType_TYPE_VIP      HallEnumShopType = 5
)

var HallEnumShopType_name = map[int32]string{
	1: "TYPE_COIN",
	2: "TYPE_DIAMOND",
	3: "TYPE_EXCHANGE",
	4: "TYPE_BUY",
	5: "TYPE_VIP",
}
var HallEnumShopType_value = map[string]int32{
	"TYPE_COIN":     1,
	"TYPE_DIAMOND":  2,
	"TYPE_EXCHANGE": 3,
	"TYPE_BUY":      4,
	"TYPE_VIP":      5,
}

func (x HallEnumShopType) Enum() *HallEnumShopType {
	p := new(HallEnumShopType)
	*p = x
	return p
}
func (x HallEnumShopType) String() string {
	return proto.EnumName(HallEnumShopType_name, int32(x))
}
func (x *HallEnumShopType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumShopType_value, data, "HallEnumShopType")
	if err != nil {
		return err
	}
	*x = HallEnumShopType(value)
	return nil
}
func (HallEnumShopType) EnumDescriptor() ([]byte, []int) { return fileDescriptor12, []int{7} }

// 单个活动
type HallItemEvent struct {
	Id               *int32           `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Type             *HallEnumEvent   `protobuf:"varint,2,opt,name=type,enum=ddproto.HallEnumEvent" json:"type,omitempty"`
	Reward           *HallEnum_Reward `protobuf:"varint,3,opt,name=reward,enum=ddproto.HallEnum_Reward" json:"reward,omitempty"`
	RichText         []string         `protobuf:"bytes,5,rep,name=richText" json:"richText,omitempty"`
	Title            *string          `protobuf:"bytes,6,opt,name=title" json:"title,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *HallItemEvent) Reset()                    { *m = HallItemEvent{} }
func (m *HallItemEvent) String() string            { return proto.CompactTextString(m) }
func (*HallItemEvent) ProtoMessage()               {}
func (*HallItemEvent) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *HallItemEvent) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *HallItemEvent) GetType() HallEnumEvent {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return HallEnumEvent_TYPE_TIME
}

func (m *HallItemEvent) GetReward() HallEnum_Reward {
	if m != nil && m.Reward != nil {
		return *m.Reward
	}
	return HallEnum_Reward_RE_EXP
}

func (m *HallItemEvent) GetRichText() []string {
	if m != nil {
		return m.RichText
	}
	return nil
}

func (m *HallItemEvent) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

// 单个邮件
type HallItemEmail struct {
	Id               *int32             `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Type             *HallEnumEmailType `protobuf:"varint,2,opt,name=type,enum=ddproto.HallEnumEmailType" json:"type,omitempty"`
	EmailTitle       *string            `protobuf:"bytes,3,opt,name=emailTitle" json:"emailTitle,omitempty"`
	EmailContent     *string            `protobuf:"bytes,4,opt,name=emailContent" json:"emailContent,omitempty"`
	BtnCheck         *bool              `protobuf:"varint,5,opt,name=btnCheck" json:"btnCheck,omitempty"`
	Date             *int32             `protobuf:"varint,6,opt,name=date" json:"date,omitempty"`
	Reward           *string            `protobuf:"bytes,7,opt,name=reward" json:"reward,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *HallItemEmail) Reset()                    { *m = HallItemEmail{} }
func (m *HallItemEmail) String() string            { return proto.CompactTextString(m) }
func (*HallItemEmail) ProtoMessage()               {}
func (*HallItemEmail) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

func (m *HallItemEmail) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *HallItemEmail) GetType() HallEnumEmailType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return HallEnumEmailType_EMAIL_SYSTEM
}

func (m *HallItemEmail) GetEmailTitle() string {
	if m != nil && m.EmailTitle != nil {
		return *m.EmailTitle
	}
	return ""
}

func (m *HallItemEmail) GetEmailContent() string {
	if m != nil && m.EmailContent != nil {
		return *m.EmailContent
	}
	return ""
}

func (m *HallItemEmail) GetBtnCheck() bool {
	if m != nil && m.BtnCheck != nil {
		return *m.BtnCheck
	}
	return false
}

func (m *HallItemEmail) GetDate() int32 {
	if m != nil && m.Date != nil {
		return *m.Date
	}
	return 0
}

func (m *HallItemEmail) GetReward() string {
	if m != nil && m.Reward != nil {
		return *m.Reward
	}
	return ""
}

// 道具
type HallItemBag struct {
	Id               *int32               `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Type             *HallEnumBagItemType `protobuf:"varint,2,opt,name=type,enum=ddproto.HallEnumBagItemType" json:"type,omitempty"`
	Rank             *HallEnumBagItemRank `protobuf:"varint,3,opt,name=rank,enum=ddproto.HallEnumBagItemRank" json:"rank,omitempty"`
	Exp              *int32               `protobuf:"varint,4,opt,name=exp" json:"exp,omitempty"`
	Amount           *int32               `protobuf:"varint,5,opt,name=amount" json:"amount,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *HallItemBag) Reset()                    { *m = HallItemBag{} }
func (m *HallItemBag) String() string            { return proto.CompactTextString(m) }
func (*HallItemBag) ProtoMessage()               {}
func (*HallItemBag) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{2} }

func (m *HallItemBag) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *HallItemBag) GetType() HallEnumBagItemType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return HallEnumBagItemType_TYPE_MEAT
}

func (m *HallItemBag) GetRank() HallEnumBagItemRank {
	if m != nil && m.Rank != nil {
		return *m.Rank
	}
	return HallEnumBagItemRank_RANK_ONE
}

func (m *HallItemBag) GetExp() int32 {
	if m != nil && m.Exp != nil {
		return *m.Exp
	}
	return 0
}

func (m *HallItemBag) GetAmount() int32 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

// 单个任务
type HallItemTask struct {
	Id               *int32            `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Type             *HallEnumTaskType `protobuf:"varint,2,opt,name=type,enum=ddproto.HallEnumTaskType" json:"type,omitempty"`
	Title            *string           `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Reward           *string           `protobuf:"bytes,4,opt,name=reward" json:"reward,omitempty"`
	Num              *int32            `protobuf:"varint,5,opt,name=Num" json:"Num,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *HallItemTask) Reset()                    { *m = HallItemTask{} }
func (m *HallItemTask) String() string            { return proto.CompactTextString(m) }
func (*HallItemTask) ProtoMessage()               {}
func (*HallItemTask) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{3} }

func (m *HallItemTask) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *HallItemTask) GetType() HallEnumTaskType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return HallEnumTaskType_TYPE_MJ
}

func (m *HallItemTask) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *HallItemTask) GetReward() string {
	if m != nil && m.Reward != nil {
		return *m.Reward
	}
	return ""
}

func (m *HallItemTask) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

// 个人信息
type HallUserData struct {
	UserName         *string `protobuf:"bytes,1,opt,name=userName" json:"userName,omitempty"`
	UserId           *int32  `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	NiceValue        *int32  `protobuf:"varint,3,opt,name=niceValue" json:"niceValue,omitempty"`
	EvilValue        *int32  `protobuf:"varint,4,opt,name=evilValue" json:"evilValue,omitempty"`
	UserLevel        *int32  `protobuf:"varint,5,opt,name=userLevel" json:"userLevel,omitempty"`
	UserVIP          *bool   `protobuf:"varint,6,opt,name=userVIP" json:"userVIP,omitempty"`
	UserMoney        *int32  `protobuf:"varint,7,opt,name=userMoney" json:"userMoney,omitempty"`
	UserDiamond      *int32  `protobuf:"varint,8,opt,name=userDiamond" json:"userDiamond,omitempty"`
	UserRedBag       *int32  `protobuf:"varint,9,opt,name=userRedBag" json:"userRedBag,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HallUserData) Reset()                    { *m = HallUserData{} }
func (m *HallUserData) String() string            { return proto.CompactTextString(m) }
func (*HallUserData) ProtoMessage()               {}
func (*HallUserData) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{4} }

func (m *HallUserData) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

func (m *HallUserData) GetUserId() int32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *HallUserData) GetNiceValue() int32 {
	if m != nil && m.NiceValue != nil {
		return *m.NiceValue
	}
	return 0
}

func (m *HallUserData) GetEvilValue() int32 {
	if m != nil && m.EvilValue != nil {
		return *m.EvilValue
	}
	return 0
}

func (m *HallUserData) GetUserLevel() int32 {
	if m != nil && m.UserLevel != nil {
		return *m.UserLevel
	}
	return 0
}

func (m *HallUserData) GetUserVIP() bool {
	if m != nil && m.UserVIP != nil {
		return *m.UserVIP
	}
	return false
}

func (m *HallUserData) GetUserMoney() int32 {
	if m != nil && m.UserMoney != nil {
		return *m.UserMoney
	}
	return 0
}

func (m *HallUserData) GetUserDiamond() int32 {
	if m != nil && m.UserDiamond != nil {
		return *m.UserDiamond
	}
	return 0
}

func (m *HallUserData) GetUserRedBag() int32 {
	if m != nil && m.UserRedBag != nil {
		return *m.UserRedBag
	}
	return 0
}

// 排行信息
type HallItemRank struct {
	Placing          *int32  `protobuf:"varint,1,opt,name=placing" json:"placing,omitempty"`
	UserId           *int32  `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	RankInfo         *string `protobuf:"bytes,3,opt,name=rankInfo" json:"rankInfo,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HallItemRank) Reset()                    { *m = HallItemRank{} }
func (m *HallItemRank) String() string            { return proto.CompactTextString(m) }
func (*HallItemRank) ProtoMessage()               {}
func (*HallItemRank) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{5} }

func (m *HallItemRank) GetPlacing() int32 {
	if m != nil && m.Placing != nil {
		return *m.Placing
	}
	return 0
}

func (m *HallItemRank) GetUserId() int32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *HallItemRank) GetRankInfo() string {
	if m != nil && m.RankInfo != nil {
		return *m.RankInfo
	}
	return ""
}

// 商城
type HallShop struct {
	Type             *HallEnumShopType `protobuf:"varint,1,opt,name=type,enum=ddproto.HallEnumShopType" json:"type,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *HallShop) Reset()                    { *m = HallShop{} }
func (m *HallShop) String() string            { return proto.CompactTextString(m) }
func (*HallShop) ProtoMessage()               {}
func (*HallShop) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{6} }

func (m *HallShop) GetType() HallEnumShopType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return HallEnumShopType_TYPE_COIN
}

// 金币专区
type CoinZone struct {
	Pay              []*GoodsItem `protobuf:"bytes,1,rep,name=pay" json:"pay,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *CoinZone) Reset()                    { *m = CoinZone{} }
func (m *CoinZone) String() string            { return proto.CompactTextString(m) }
func (*CoinZone) ProtoMessage()               {}
func (*CoinZone) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{7} }

func (m *CoinZone) GetPay() []*GoodsItem {
	if m != nil {
		return m.Pay
	}
	return nil
}

// 钻石专区
type DiamondZone struct {
	Pay              []*GoodsItem `protobuf:"bytes,1,rep,name=pay" json:"pay,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DiamondZone) Reset()                    { *m = DiamondZone{} }
func (m *DiamondZone) String() string            { return proto.CompactTextString(m) }
func (*DiamondZone) ProtoMessage()               {}
func (*DiamondZone) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{8} }

func (m *DiamondZone) GetPay() []*GoodsItem {
	if m != nil {
		return m.Pay
	}
	return nil
}

// 兑换专区
type ExchangeZone struct {
	Money            []*GoodsItem `protobuf:"bytes,1,rep,name=money" json:"money,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ExchangeZone) Reset()                    { *m = ExchangeZone{} }
func (m *ExchangeZone) String() string            { return proto.CompactTextString(m) }
func (*ExchangeZone) ProtoMessage()               {}
func (*ExchangeZone) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{9} }

func (m *ExchangeZone) GetMoney() []*GoodsItem {
	if m != nil {
		return m.Money
	}
	return nil
}

// 购买专区
type BuyZone struct {
	Pay              []*GoodsItem `protobuf:"bytes,1,rep,name=pay" json:"pay,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *BuyZone) Reset()                    { *m = BuyZone{} }
func (m *BuyZone) String() string            { return proto.CompactTextString(m) }
func (*BuyZone) ProtoMessage()               {}
func (*BuyZone) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{10} }

func (m *BuyZone) GetPay() []*GoodsItem {
	if m != nil {
		return m.Pay
	}
	return nil
}

// 商品类型
type GoodsItem struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Money            *int32  `protobuf:"varint,2,opt,name=money" json:"money,omitempty"`
	Img              *string `protobuf:"bytes,3,opt,name=img" json:"img,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GoodsItem) Reset()                    { *m = GoodsItem{} }
func (m *GoodsItem) String() string            { return proto.CompactTextString(m) }
func (*GoodsItem) ProtoMessage()               {}
func (*GoodsItem) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{11} }

func (m *GoodsItem) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *GoodsItem) GetMoney() int32 {
	if m != nil && m.Money != nil {
		return *m.Money
	}
	return 0
}

func (m *GoodsItem) GetImg() string {
	if m != nil && m.Img != nil {
		return *m.Img
	}
	return ""
}

func init() {
	proto.RegisterType((*HallItemEvent)(nil), "ddproto.hall_item_event")
	proto.RegisterType((*HallItemEmail)(nil), "ddproto.hall_item_email")
	proto.RegisterType((*HallItemBag)(nil), "ddproto.hall_item_bag")
	proto.RegisterType((*HallItemTask)(nil), "ddproto.hall_item_task")
	proto.RegisterType((*HallUserData)(nil), "ddproto.hall_userData")
	proto.RegisterType((*HallItemRank)(nil), "ddproto.hall_item_rank")
	proto.RegisterType((*HallShop)(nil), "ddproto.hall_shop")
	proto.RegisterType((*CoinZone)(nil), "ddproto.CoinZone")
	proto.RegisterType((*DiamondZone)(nil), "ddproto.DiamondZone")
	proto.RegisterType((*ExchangeZone)(nil), "ddproto.ExchangeZone")
	proto.RegisterType((*BuyZone)(nil), "ddproto.BuyZone")
	proto.RegisterType((*GoodsItem)(nil), "ddproto.GoodsItem")
	proto.RegisterEnum("ddproto.HallEnumProtoId", HallEnumProtoId_name, HallEnumProtoId_value)
	proto.RegisterEnum("ddproto.HallEnumEvent", HallEnumEvent_name, HallEnumEvent_value)
	proto.RegisterEnum("ddproto.HallEnum_Reward", HallEnum_Reward_name, HallEnum_Reward_value)
	proto.RegisterEnum("ddproto.HallEnumEmailType", HallEnumEmailType_name, HallEnumEmailType_value)
	proto.RegisterEnum("ddproto.HallEnumBagItemType", HallEnumBagItemType_name, HallEnumBagItemType_value)
	proto.RegisterEnum("ddproto.HallEnumBagItemRank", HallEnumBagItemRank_name, HallEnumBagItemRank_value)
	proto.RegisterEnum("ddproto.HallEnumTaskType", HallEnumTaskType_name, HallEnumTaskType_value)
	proto.RegisterEnum("ddproto.HallEnumShopType", HallEnumShopType_name, HallEnumShopType_value)
}

var fileDescriptor12 = []byte{
	// 896 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x94, 0xdd, 0x72, 0xda, 0x46,
	0x14, 0xc7, 0xcb, 0x87, 0x0c, 0x3a, 0x80, 0xad, 0xac, 0x9b, 0x44, 0x6e, 0xfa, 0x91, 0x72, 0xd1,
	0x71, 0x48, 0xc7, 0x33, 0x4d, 0x67, 0x3a, 0xbd, 0xe9, 0x85, 0x2c, 0x14, 0x50, 0x0c, 0x82, 0xca,
	0x72, 0x1c, 0xfb, 0x46, 0xa3, 0xa0, 0x2d, 0x68, 0x0c, 0x12, 0x03, 0x22, 0x35, 0x0f, 0xd0, 0x07,
	0xe8, 0x4c, 0xdf, 0xa0, 0x0f, 0xd1, 0xd7, 0xeb, 0xd9, 0x23, 0x04, 0xa8, 0x51, 0x3a, 0xbe, 0xd3,
	0x9e, 0x3d, 0xff, 0x3d, 0xbf, 0xf3, 0xdf, 0xd5, 0x81, 0xa3, 0x89, 0x37, 0x9d, 0xba, 0xbe, 0x17,
	0x7b, 0x67, 0xf3, 0x45, 0x14, 0x47, 0xac, 0xe2, 0xfb, 0xf4, 0xd1, 0xfc, 0xb3, 0xb0, 0xd9, 0x0c,
	0x62, 0x3e, 0x73, 0xf9, 0x07, 0x1e, 0xc6, 0x0c, 0xa0, 0x18, 0xf8, 0x6a, 0xe1, 0x79, 0xe1, 0x54,
	0x62, 0xdf, 0x41, 0x39, 0x5e, 0xcf, 0xb9, 0x5a, 0xc4, 0xd5, 0xe1, 0x2b, 0xf5, 0x6c, 0xa3, 0x3b,
	0x23, 0x0d, 0x0f, 0x57, 0xa9, 0xe6, 0x05, 0x1c, 0x2c, 0xf8, 0xef, 0xde, 0xc2, 0x57, 0x4b, 0x94,
	0x79, 0x92, 0x93, 0x69, 0x53, 0x02, 0x53, 0xa0, 0xba, 0x08, 0x46, 0x13, 0x87, 0xdf, 0xc7, 0xaa,
	0xf4, 0xbc, 0x74, 0x2a, 0xb3, 0x06, 0x48, 0x71, 0x10, 0x4f, 0xb9, 0x7a, 0x80, 0x5a, 0xb9, 0xf9,
	0x77, 0x96, 0x69, 0xe6, 0x05, 0xd3, 0x0c, 0x53, 0x2b, 0xc3, 0xf4, 0x65, 0x1e, 0x93, 0xd0, 0x38,
	0x98, 0xc3, 0x50, 0x98, 0x2c, 0xe8, 0x7c, 0xc1, 0x26, 0xb3, 0xcf, 0xa1, 0x4e, 0x31, 0x3d, 0x0a,
	0x63, 0x64, 0x57, 0xcb, 0x14, 0x45, 0xac, 0xf7, 0x71, 0xa8, 0x4f, 0xf8, 0xe8, 0x0e, 0xb1, 0x0a,
	0xa7, 0x55, 0x56, 0x87, 0x32, 0x5a, 0x96, 0x50, 0x49, 0xec, 0x70, 0xdb, 0x61, 0x85, 0x28, 0xff,
	0x2a, 0x40, 0x63, 0x47, 0xf9, 0xde, 0x1b, 0x67, 0x18, 0xbf, 0xcf, 0x30, 0x7e, 0x9d, 0xc3, 0x88,
	0x0a, 0x13, 0x95, 0x44, 0x89, 0xd9, 0x0b, 0x2f, 0xbc, 0xdb, 0x78, 0xf7, 0x3f, 0xd9, 0x36, 0x66,
	0xb1, 0x1a, 0x94, 0xf8, 0xfd, 0x9c, 0xb0, 0x09, 0xcb, 0x9b, 0x45, 0xab, 0x30, 0x26, 0x68, 0xa9,
	0xb9, 0x84, 0xc3, 0x1d, 0x55, 0xec, 0x2d, 0xef, 0x32, 0x58, 0x2f, 0x32, 0x58, 0xcf, 0x72, 0x0a,
	0x09, 0x09, 0x31, 0x6d, 0x2f, 0x25, 0x31, 0x6d, 0xd7, 0x7e, 0x62, 0x17, 0x42, 0x58, 0xab, 0xd9,
	0xa6, 0xe8, 0x3f, 0xa9, 0x17, 0xab, 0x25, 0x5f, 0xb4, 0xf1, 0x99, 0x09, 0x37, 0xc5, 0xb7, 0xe5,
	0xcd, 0x38, 0x95, 0xa6, 0x03, 0x44, 0xc4, 0xf4, 0xa9, 0xb8, 0xc4, 0x1e, 0x81, 0x1c, 0x06, 0x23,
	0xfe, 0xd6, 0x9b, 0xae, 0x92, 0x1a, 0x14, 0xe2, 0x1f, 0x82, 0x69, 0x12, 0x2a, 0xa7, 0x21, 0xa1,
	0xea, 0xe1, 0x2b, 0x9b, 0x26, 0xc5, 0xd8, 0x11, 0x54, 0x44, 0xe8, 0xad, 0x39, 0xa4, 0x9b, 0xa9,
	0xa6, 0x39, 0xfd, 0x28, 0xe4, 0x6b, 0xba, 0x1c, 0x89, 0x1d, 0x43, 0x8d, 0x50, 0x02, 0xf4, 0x26,
	0xf4, 0xd5, 0x2a, 0x05, 0xd1, 0x09, 0x11, 0xb4, 0xb9, 0x7f, 0xee, 0x8d, 0x55, 0x99, 0xc8, 0xf5,
	0x7d, 0xbb, 0xc4, 0x1d, 0x88, 0xe3, 0xe7, 0x53, 0x6f, 0x14, 0x84, 0xe3, 0x8d, 0x67, 0xff, 0x05,
	0x17, 0xef, 0x17, 0x13, 0xcd, 0xf0, 0xb7, 0x28, 0xf1, 0xa6, 0xf9, 0x13, 0xc8, 0x74, 0xc8, 0x72,
	0x12, 0xcd, 0xb7, 0x16, 0x17, 0x3e, 0x69, 0xb1, 0x48, 0x13, 0x16, 0x37, 0x5f, 0x42, 0x55, 0x8f,
	0x82, 0xf0, 0x16, 0xc1, 0xd9, 0x37, 0x50, 0x9a, 0x7b, 0x6b, 0x54, 0x95, 0x4e, 0x6b, 0xaf, 0xd8,
	0x56, 0xd5, 0x89, 0x22, 0x7f, 0x29, 0x6e, 0xbe, 0x79, 0x06, 0xb5, 0x4d, 0x3b, 0x0f, 0xcb, 0xff,
	0x01, 0xea, 0xc6, 0xfd, 0x68, 0xe2, 0x85, 0x63, 0x4e, 0x82, 0x6f, 0x41, 0x9a, 0x91, 0x43, 0x9f,
	0x96, 0xb4, 0xa0, 0x72, 0xbe, 0x5a, 0x3f, 0xec, 0xf8, 0x1f, 0x41, 0xde, 0x2e, 0x32, 0x4f, 0xac,
	0x91, 0xd6, 0x49, 0xdc, 0xc2, 0x77, 0x12, 0xcc, 0xc6, 0x89, 0x51, 0xad, 0x3f, 0x8a, 0xf0, 0x68,
	0xe7, 0x03, 0x1d, 0x69, 0xfa, 0xec, 0x09, 0xb0, 0xae, 0xd6, 0xeb, 0xb9, 0x43, 0xb3, 0xed, 0x76,
	0x0d, 0xcd, 0x76, 0xce, 0x0d, 0xcd, 0x51, 0x3e, 0x63, 0x4f, 0xe1, 0x78, 0x1b, 0xff, 0xf5, 0xca,
	0xd4, 0x2f, 0x5c, 0x7d, 0x60, 0x59, 0x4a, 0x81, 0x3d, 0x83, 0xa7, 0x39, 0x1b, 0xae, 0xa6, 0x5f,
	0x28, 0xc5, 0x8c, 0xaa, 0xa3, 0xf5, 0x0d, 0xb7, 0x37, 0xe8, 0x98, 0x96, 0x52, 0xca, 0xa8, 0x76,
	0x1b, 0xa4, 0x2a, 0x67, 0x18, 0xae, 0x2e, 0x0d, 0xdb, 0x6d, 0x6b, 0x8e, 0xa6, 0xc8, 0xec, 0x0b,
	0x78, 0xf2, 0x71, 0x9c, 0x34, 0xc0, 0x4e, 0xe0, 0xf1, 0x76, 0xaf, 0x6d, 0x6b, 0xd7, 0x78, 0xa0,
	0xe3, 0x18, 0xf6, 0x8d, 0x52, 0x63, 0x5f, 0xc1, 0x49, 0xee, 0x16, 0x29, 0xeb, 0xad, 0x5f, 0x36,
	0x03, 0x6e, 0x6f, 0x80, 0x36, 0x40, 0x76, 0x6e, 0x86, 0x86, 0xeb, 0x98, 0x7d, 0x03, 0x5b, 0xac,
	0x43, 0x95, 0x96, 0x96, 0x71, 0x8d, 0x3d, 0xa5, 0x9b, 0xd6, 0x55, 0xaf, 0xa7, 0x94, 0x5a, 0x2f,
	0x41, 0xf9, 0x68, 0xaa, 0x02, 0x1c, 0xd8, 0x86, 0x6b, 0xbc, 0x1b, 0xa2, 0xb8, 0x06, 0x15, 0xfc,
	0xee, 0x98, 0xaf, 0x1d, 0xa5, 0xd8, 0xfa, 0x19, 0x8e, 0xf3, 0x06, 0xa3, 0x82, 0xcf, 0xa3, 0xaf,
	0x99, 0x3d, 0xf7, 0xf2, 0xe6, 0xd2, 0x31, 0xfa, 0xa8, 0x3a, 0x82, 0x1a, 0x15, 0x79, 0x6d, 0x9b,
	0x86, 0xd5, 0x46, 0xa5, 0x07, 0x8f, 0xf3, 0xc7, 0x55, 0x8a, 0xd3, 0x17, 0xf7, 0xb4, 0x63, 0x35,
	0x3a, 0x1d, 0x64, 0x4d, 0x57, 0xa6, 0x6e, 0xa0, 0xe9, 0xe9, 0xea, 0xfc, 0xca, 0x42, 0x97, 0xd3,
	0x12, 0x97, 0x5d, 0xdb, 0xec, 0x0f, 0x15, 0xa9, 0xa5, 0xe7, 0x94, 0xa0, 0x19, 0x87, 0x3a, 0x5b,
	0xb3, 0x2e, 0xdc, 0x81, 0xb5, 0x71, 0x83, 0x56, 0xce, 0xf5, 0x00, 0x2b, 0x1c, 0x02, 0x24, 0xab,
	0xae, 0x6d, 0x60, 0x0d, 0x74, 0x93, 0xe5, 0xcc, 0x2f, 0x34, 0x21, 0x81, 0x7c, 0xb3, 0x87, 0xd8,
	0x6e, 0xdf, 0xee, 0x21, 0xde, 0xbe, 0xe9, 0xa2, 0xdc, 0xdf, 0x97, 0xa7, 0xff, 0xe6, 0xb6, 0x47,
	0x7d, 0x60, 0x8a, 0x27, 0x87, 0x76, 0x25, 0x07, 0x98, 0x5a, 0x7f, 0x20, 0xdc, 0xc1, 0xa9, 0xd3,
	0x48, 0xba, 0x7e, 0xa7, 0x77, 0x35, 0xab, 0x93, 0x6d, 0xf6, 0x06, 0x9b, 0x4d, 0x57, 0x38, 0xa8,
	0x14, 0xe9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x15, 0xbe, 0xb5, 0x83, 0x07, 0x00, 0x00,
}
