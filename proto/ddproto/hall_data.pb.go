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
	HallEnumProtoId_HALL_PID_HEARTBEAT            HallEnumProtoId = 0
	HallEnumProtoId_HALL_PID_QUICK_CONN           HallEnumProtoId = 1
	HallEnumProtoId_HALL_PID_QUICK_CONN_ACK       HallEnumProtoId = 2
	HallEnumProtoId_HALL_PID_GAME_LOGIN           HallEnumProtoId = 3
	HallEnumProtoId_HALL_PID_GAME_LOGIN_ACK       HallEnumProtoId = 4
	HallEnumProtoId_HALL_PID_USER_DATA            HallEnumProtoId = 9
	HallEnumProtoId_HALL_PID_USER_DATA_ACK        HallEnumProtoId = 10
	HallEnumProtoId_HALL_PID_DRAW_LOTTERY         HallEnumProtoId = 11
	HallEnumProtoId_HALL_PID_DRAW_LOTTERY_ACK     HallEnumProtoId = 12
	HallEnumProtoId_HALL_PID_DS_LOTTERY_ITEMS_ACK HallEnumProtoId = 13
	HallEnumProtoId_HALL_PID_SIGN_LOTTERY         HallEnumProtoId = 14
	HallEnumProtoId_HALL_PID_SIGN_LOTTERY_ACK     HallEnumProtoId = 15
	HallEnumProtoId_HALL_PID_ONLINEWARD_REQ       HallEnumProtoId = 16
	HallEnumProtoId_HALL_PID_ONLINEWARD_ACK       HallEnumProtoId = 17
	HallEnumProtoId_HALL_PID_EVENT_REQ            HallEnumProtoId = 18
	HallEnumProtoId_HALL_PID_EVENT_ACK            HallEnumProtoId = 19
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
	13: "HALL_PID_DS_LOTTERY_ITEMS_ACK",
	14: "HALL_PID_SIGN_LOTTERY",
	15: "HALL_PID_SIGN_LOTTERY_ACK",
	16: "HALL_PID_ONLINEWARD_REQ",
	17: "HALL_PID_ONLINEWARD_ACK",
	18: "HALL_PID_EVENT_REQ",
	19: "HALL_PID_EVENT_ACK",
}
var HallEnumProtoId_value = map[string]int32{
	"HALL_PID_HEARTBEAT":            0,
	"HALL_PID_QUICK_CONN":           1,
	"HALL_PID_QUICK_CONN_ACK":       2,
	"HALL_PID_GAME_LOGIN":           3,
	"HALL_PID_GAME_LOGIN_ACK":       4,
	"HALL_PID_USER_DATA":            9,
	"HALL_PID_USER_DATA_ACK":        10,
	"HALL_PID_DRAW_LOTTERY":         11,
	"HALL_PID_DRAW_LOTTERY_ACK":     12,
	"HALL_PID_DS_LOTTERY_ITEMS_ACK": 13,
	"HALL_PID_SIGN_LOTTERY":         14,
	"HALL_PID_SIGN_LOTTERY_ACK":     15,
	"HALL_PID_ONLINEWARD_REQ":       16,
	"HALL_PID_ONLINEWARD_ACK":       17,
	"HALL_PID_EVENT_REQ":            18,
	"HALL_PID_EVENT_ACK":            19,
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
func (HallEnumProtoId) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

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
func (HallEnumEvent) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

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
func (HallEnum_Reward) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

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
func (HallEnumEmailType) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{3} }

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
func (HallEnumBagItemType) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{4} }

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
func (HallEnumBagItemRank) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{5} }

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
func (HallEnumTaskType) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{6} }

// 奖励物品信息
type HallLotteryItem struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HallLotteryItem) Reset()                    { *m = HallLotteryItem{} }
func (m *HallLotteryItem) String() string            { return proto.CompactTextString(m) }
func (*HallLotteryItem) ProtoMessage()               {}
func (*HallLotteryItem) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *HallLotteryItem) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *HallLotteryItem) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

// 签到信息与奖励物品
type HallSignLotteryInfo struct {
	SignItems        []*HallLotteryItem `protobuf:"bytes,1,rep,name=signItems" json:"signItems,omitempty"`
	TotalDays        *int32             `protobuf:"varint,2,opt,name=totalDays" json:"totalDays,omitempty"`
	ContinuousDays   *int32             `protobuf:"varint,3,opt,name=continuousDays" json:"continuousDays,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *HallSignLotteryInfo) Reset()                    { *m = HallSignLotteryInfo{} }
func (m *HallSignLotteryInfo) String() string            { return proto.CompactTextString(m) }
func (*HallSignLotteryInfo) ProtoMessage()               {}
func (*HallSignLotteryInfo) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

func (m *HallSignLotteryInfo) GetSignItems() []*HallLotteryItem {
	if m != nil {
		return m.SignItems
	}
	return nil
}

func (m *HallSignLotteryInfo) GetTotalDays() int32 {
	if m != nil && m.TotalDays != nil {
		return *m.TotalDays
	}
	return 0
}

func (m *HallSignLotteryInfo) GetContinuousDays() int32 {
	if m != nil && m.ContinuousDays != nil {
		return *m.ContinuousDays
	}
	return 0
}

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
func (*HallItemEvent) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

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
func (*HallItemEmail) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{3} }

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
func (*HallItemBag) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{4} }

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
func (*HallItemTask) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{5} }

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
func (*HallUserData) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{6} }

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
func (*HallItemRank) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{7} }

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

// 金币专区
type CoinZone struct {
	Pay              []*GoodsItem `protobuf:"bytes,1,rep,name=pay" json:"pay,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *CoinZone) Reset()                    { *m = CoinZone{} }
func (m *CoinZone) String() string            { return proto.CompactTextString(m) }
func (*CoinZone) ProtoMessage()               {}
func (*CoinZone) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{8} }

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
func (*DiamondZone) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{9} }

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
func (*ExchangeZone) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{10} }

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
func (*BuyZone) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{11} }

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
func (*GoodsItem) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{12} }

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
	proto.RegisterType((*HallLotteryItem)(nil), "ddproto.hall_lottery_item")
	proto.RegisterType((*HallSignLotteryInfo)(nil), "ddproto.hall_sign_lottery_info")
	proto.RegisterType((*HallItemEvent)(nil), "ddproto.hall_item_event")
	proto.RegisterType((*HallItemEmail)(nil), "ddproto.hall_item_email")
	proto.RegisterType((*HallItemBag)(nil), "ddproto.hall_item_bag")
	proto.RegisterType((*HallItemTask)(nil), "ddproto.hall_item_task")
	proto.RegisterType((*HallUserData)(nil), "ddproto.hall_userData")
	proto.RegisterType((*HallItemRank)(nil), "ddproto.hall_item_rank")
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
}

var fileDescriptor14 = []byte{
	// 1074 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x72, 0xda, 0xc6,
	0x17, 0xff, 0x0b, 0x41, 0x80, 0x83, 0x43, 0x36, 0x9b, 0xc4, 0x51, 0x9c, 0x8f, 0xe1, 0xaf, 0xe9,
	0x74, 0x18, 0xda, 0x71, 0x52, 0xf7, 0xa2, 0xbe, 0xc9, 0x05, 0x86, 0x0d, 0x56, 0x0c, 0xc2, 0x59,
	0x64, 0xbb, 0xf6, 0x45, 0x35, 0x6b, 0xb4, 0xc1, 0x1a, 0x83, 0xc4, 0x80, 0x70, 0xcd, 0x63, 0xf4,
	0xaa, 0x7d, 0x8d, 0xce, 0xf4, 0x31, 0xfa, 0x0e, 0x7d, 0x95, 0xce, 0x2e, 0x6b, 0x7d, 0x04, 0xdc,
	0x99, 0xdc, 0xed, 0xef, 0x9c, 0xf3, 0xdb, 0xf3, 0x3b, 0x1f, 0xd2, 0xc2, 0xa3, 0x2b, 0x36, 0x1e,
	0xbb, 0x1e, 0x8b, 0xd8, 0xee, 0x74, 0x16, 0x46, 0x21, 0x2e, 0x7a, 0x9e, 0x3c, 0x98, 0x3f, 0xc1,
	0x63, 0xe9, 0x1b, 0x87, 0x51, 0xc4, 0x67, 0x4b, 0xd7, 0x8f, 0xf8, 0x04, 0x57, 0x21, 0xe7, 0x7b,
	0x86, 0x56, 0xd3, 0xea, 0x05, 0x9a, 0xf3, 0x3d, 0x8c, 0x21, 0x1f, 0xb0, 0x09, 0x37, 0x72, 0x35,
	0xad, 0x5e, 0xa6, 0xf2, 0x6c, 0xfe, 0xa1, 0xc1, 0xb6, 0x64, 0xce, 0xfd, 0x51, 0x90, 0xd0, 0x83,
	0xcf, 0x21, 0xde, 0x87, 0xb2, 0x30, 0x5a, 0x11, 0x9f, 0xcc, 0x0d, 0xad, 0xa6, 0xd7, 0x2b, 0x7b,
	0x3b, 0xbb, 0x2a, 0xe1, 0xee, 0x5a, 0x36, 0x9a, 0x04, 0xe3, 0x57, 0x50, 0x8e, 0xc2, 0x88, 0x8d,
	0xdb, 0x6c, 0x39, 0x97, 0xd9, 0x0a, 0x34, 0x31, 0xe0, 0x6f, 0xa1, 0x3a, 0x0c, 0x83, 0xc8, 0x0f,
	0x16, 0xe1, 0x62, 0x2e, 0x43, 0x74, 0x19, 0xf2, 0x85, 0xd5, 0xfc, 0x4b, 0x53, 0x05, 0x8b, 0xeb,
	0x5d, 0x7e, 0xc3, 0x83, 0x68, 0xad, 0xa4, 0xef, 0x21, 0x1f, 0x2d, 0xa7, 0xab, 0x92, 0xaa, 0x7b,
	0x46, 0x56, 0x1e, 0x0f, 0x16, 0x8a, 0x47, 0x65, 0x14, 0xfe, 0x01, 0x1e, 0xcc, 0xf8, 0xaf, 0x6c,
	0xe6, 0xc9, 0x8c, 0xd5, 0xbd, 0x17, 0x1b, 0xe2, 0xa9, 0x0c, 0xa0, 0x2a, 0x10, 0xef, 0x40, 0x69,
	0xe6, 0x0f, 0xaf, 0x1c, 0x7e, 0x1b, 0x19, 0x85, 0x9a, 0x5e, 0x2f, 0xd3, 0x18, 0xe3, 0xa7, 0x50,
	0x88, 0xfc, 0x68, 0xcc, 0x8d, 0x07, 0xb2, 0xa1, 0x2b, 0x60, 0xfe, 0x93, 0x95, 0x3d, 0x61, 0xfe,
	0x78, 0x4d, 0xf6, 0xbb, 0x8c, 0xec, 0x57, 0x9b, 0x64, 0x0b, 0x9e, 0xb3, 0x9c, 0x72, 0x25, 0xfd,
	0x0d, 0xc0, 0xca, 0x24, 0x13, 0xea, 0x32, 0x61, 0xca, 0x82, 0x4d, 0xd8, 0x92, 0xa8, 0x15, 0x06,
	0x11, 0x0f, 0x22, 0x23, 0x2f, 0x23, 0x32, 0x36, 0x51, 0xcb, 0x65, 0x14, 0xb4, 0xae, 0xf8, 0xf0,
	0xda, 0x28, 0xd4, 0xb4, 0x7a, 0x89, 0xc6, 0x58, 0xec, 0x86, 0xc7, 0xa2, 0x55, 0x29, 0x05, 0x2a,
	0xcf, 0x78, 0x3b, 0x6e, 0x57, 0x51, 0xde, 0xa6, 0x90, 0xf9, 0xa7, 0x06, 0x0f, 0x93, 0x0a, 0x2f,
	0xd9, 0x68, 0xad, 0xbe, 0xbd, 0x4c, 0x7d, 0x6f, 0x36, 0xd4, 0x77, 0xc9, 0x46, 0x62, 0x57, 0x52,
	0x15, 0xee, 0x41, 0x7e, 0xc6, 0x82, 0x6b, 0x35, 0x9a, 0xff, 0xe0, 0x50, 0x16, 0x5c, 0x53, 0x19,
	0x8b, 0x11, 0xe8, 0xfc, 0x76, 0x2a, 0x8b, 0x2d, 0x50, 0x71, 0x14, 0x9a, 0xd9, 0x24, 0x5c, 0x04,
	0x91, 0xac, 0xb0, 0x40, 0x15, 0x32, 0x7f, 0xd3, 0xa0, 0x9a, 0x68, 0x8e, 0xd8, 0xfc, 0x7a, 0x4d,
	0xf4, 0xdb, 0x8c, 0xe8, 0x97, 0x1b, 0x04, 0x08, 0x5a, 0x4a, 0x71, 0x3c, 0x7f, 0x3d, 0x35, 0xff,
	0x54, 0xd7, 0xf2, 0xe9, 0xae, 0x09, 0xad, 0xf6, 0x62, 0xa2, 0x64, 0x89, 0xa3, 0xf9, 0x7b, 0x4e,
	0xf5, 0x71, 0x31, 0xe7, 0xb3, 0x36, 0x8b, 0x98, 0x98, 0x90, 0x38, 0xdb, 0xe2, 0x2b, 0xd5, 0x24,
	0x3b, 0xc6, 0xe2, 0x5e, 0x71, 0xb6, 0x3c, 0xf5, 0x45, 0x29, 0x24, 0x3e, 0xb6, 0xc0, 0x1f, 0xf2,
	0x53, 0x36, 0x5e, 0x70, 0xf5, 0x25, 0x25, 0x06, 0xe1, 0xe5, 0x37, 0xfe, 0x78, 0xe5, 0x5d, 0xf5,
	0x29, 0x31, 0x08, 0xaf, 0xb8, 0xa5, 0xcb, 0x6f, 0xf8, 0x58, 0x29, 0x4b, 0x0c, 0xd8, 0x80, 0xa2,
	0x00, 0xa7, 0xd6, 0xb1, 0x5c, 0x8b, 0x12, 0xbd, 0x83, 0x77, 0xbc, 0x5e, 0x18, 0xf0, 0xa5, 0x5c,
	0x0e, 0xc5, 0x93, 0x06, 0x5c, 0x83, 0x8a, 0xac, 0xc8, 0x67, 0x93, 0x30, 0xf0, 0x8c, 0x92, 0xf4,
	0xa7, 0x4d, 0x62, 0x9b, 0x05, 0xa4, 0xdc, 0x3b, 0x60, 0x23, 0xa3, 0x2c, 0x03, 0x52, 0x16, 0xf3,
	0x97, 0xf4, 0xb0, 0xe4, 0xa4, 0x0d, 0x28, 0x4e, 0xc7, 0x6c, 0xe8, 0x07, 0x23, 0x35, 0xb1, 0x3b,
	0x78, 0x6f, 0x5f, 0xc4, 0x97, 0xcb, 0x82, 0x6b, 0x2b, 0xf8, 0x1c, 0xaa, 0x01, 0xc5, 0xd8, 0x7c,
	0x07, 0xa5, 0x56, 0xe8, 0x07, 0x17, 0x61, 0xc0, 0xf1, 0x37, 0xa0, 0x4f, 0xd9, 0x52, 0xfd, 0xe0,
	0x70, 0x3c, 0xf5, 0x4e, 0x18, 0x7a, 0x73, 0xb9, 0x6e, 0xc2, 0x6d, 0xfe, 0x08, 0x15, 0x25, 0xfe,
	0x2b, 0x48, 0xfb, 0xb0, 0x45, 0x6e, 0x87, 0x57, 0x2c, 0x18, 0x71, 0xc9, 0xaa, 0x43, 0x61, 0x22,
	0x5b, 0x76, 0x3f, 0x6f, 0x15, 0x60, 0xbe, 0x85, 0xe2, 0xc1, 0x62, 0xf9, 0x15, 0xa9, 0x5a, 0x50,
	0x8e, 0x2d, 0x6b, 0x9b, 0xfd, 0xf4, 0x2e, 0xef, 0xaa, 0x43, 0x2b, 0x20, 0x16, 0xd2, 0x9f, 0x8c,
	0x54, 0x6f, 0xc4, 0xb1, 0xf1, 0xb7, 0xae, 0x9e, 0x11, 0xb9, 0xed, 0x32, 0x8f, 0xe5, 0xe1, 0x6d,
	0xc0, 0x87, 0xcd, 0x6e, 0xd7, 0x3d, 0xb6, 0xda, 0xee, 0x21, 0x69, 0x52, 0xe7, 0x80, 0x34, 0x1d,
	0xf4, 0x3f, 0xfc, 0x1c, 0x9e, 0xc4, 0xf6, 0x4f, 0x27, 0x56, 0xeb, 0xc8, 0x6d, 0xf5, 0x6d, 0x1b,
	0x69, 0xf8, 0x25, 0x3c, 0xdf, 0xe0, 0x70, 0x9b, 0xad, 0x23, 0x94, 0xcb, 0xb0, 0x3a, 0xcd, 0x1e,
	0x71, 0xbb, 0xfd, 0x8e, 0x65, 0x23, 0x3d, 0xc3, 0x4a, 0x1c, 0x92, 0x95, 0xcf, 0x68, 0x38, 0x19,
	0x10, 0xea, 0xb6, 0x9b, 0x4e, 0x13, 0x95, 0xf1, 0x0e, 0x6c, 0xaf, 0xdb, 0x25, 0x07, 0xf0, 0x0b,
	0x78, 0x16, 0xfb, 0xda, 0xb4, 0x79, 0xe6, 0x76, 0xfb, 0x8e, 0x43, 0xe8, 0x39, 0xaa, 0xe0, 0xd7,
	0xf0, 0x62, 0xa3, 0x4b, 0x32, 0xb7, 0xf0, 0xff, 0xe1, 0x75, 0xe2, 0x1e, 0xc4, 0x4e, 0xcb, 0x21,
	0xbd, 0x81, 0x0c, 0x79, 0x98, 0xb9, 0x7c, 0x60, 0x75, 0xec, 0xf8, 0xf2, 0x6a, 0xe6, 0xf2, 0xb4,
	0x4b, 0x32, 0x1f, 0x65, 0xea, 0xec, 0xdb, 0x5d, 0xcb, 0x26, 0x67, 0x4d, 0xda, 0x76, 0x29, 0xf9,
	0x84, 0xd0, 0x7d, 0x4e, 0xc1, 0x7c, 0x9c, 0x69, 0x02, 0x39, 0x25, 0xb6, 0x23, 0x49, 0x78, 0x83,
	0x5d, 0xc4, 0x3f, 0x69, 0xbc, 0x57, 0x0f, 0x51, 0xf2, 0x0e, 0xe2, 0x87, 0x50, 0x76, 0xce, 0x8f,
	0x89, 0xeb, 0x58, 0x3d, 0x82, 0x34, 0xbc, 0x05, 0x25, 0x09, 0x6d, 0x72, 0x86, 0x72, 0xb1, 0xd3,
	0x3e, 0xe9, 0x76, 0x91, 0xde, 0xf8, 0x0e, 0xd0, 0x97, 0xcf, 0x22, 0x06, 0x78, 0x40, 0x89, 0x4b,
	0x7e, 0x3e, 0x46, 0x1a, 0xae, 0x40, 0x91, 0x12, 0xb7, 0x63, 0x7d, 0x70, 0x50, 0xae, 0xb1, 0x0f,
	0x4f, 0x36, 0x3c, 0x5e, 0x18, 0xc1, 0x16, 0xe9, 0x35, 0xad, 0xae, 0x3b, 0x38, 0x1f, 0x38, 0xa4,
	0x87, 0x34, 0xfc, 0x08, 0x2a, 0x32, 0xc9, 0x07, 0x6a, 0x11, 0xbb, 0x8d, 0x72, 0x0d, 0x06, 0xcf,
	0x36, 0x3e, 0x0b, 0xb1, 0x9c, 0x9e, 0x58, 0xb7, 0x44, 0x2b, 0xe9, 0x74, 0x50, 0x2e, 0x46, 0x56,
	0x8b, 0x20, 0x3d, 0x46, 0x07, 0x27, 0x36, 0xca, 0xc7, 0x29, 0x06, 0x87, 0xd4, 0xea, 0x1d, 0xa3,
	0x42, 0xa3, 0xb5, 0x21, 0x85, 0x78, 0x45, 0x04, 0x8f, 0x36, 0xed, 0x23, 0xb7, 0x6f, 0xab, 0x6e,
	0x48, 0xe4, 0x9c, 0xf5, 0x51, 0x0e, 0x57, 0x01, 0x56, 0xe8, 0x90, 0x12, 0x82, 0xf4, 0xc6, 0x7b,
	0xc0, 0xeb, 0x2f, 0x81, 0x68, 0xc2, 0x4a, 0xe4, 0xc7, 0x94, 0xc4, 0x76, 0xfb, 0x22, 0x25, 0xf1,
	0xe2, 0xe3, 0x21, 0xd2, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x39, 0xe4, 0x0e, 0xf2, 0xbc, 0x09,
	0x00, 0x00,
}
