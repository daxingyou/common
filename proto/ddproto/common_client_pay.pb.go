// Code generated by protoc-gen-go.
// source: common_client_pay.proto
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

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of common_req_gameLogin from common_client.proto

// Ignoring public import of common_ack_gameLogin from common_client.proto

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

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// 订单的状态
type PayEnumTradeStatus int32

const (
	PayEnumTradeStatus_PAY_S_UNIFIEDORDER PayEnumTradeStatus = 1
	PayEnumTradeStatus_PAY_S_WATINGPAY    PayEnumTradeStatus = 2
	PayEnumTradeStatus_PAY_S_SUCC         PayEnumTradeStatus = 3
)

var PayEnumTradeStatus_name = map[int32]string{
	1: "PAY_S_UNIFIEDORDER",
	2: "PAY_S_WATINGPAY",
	3: "PAY_S_SUCC",
}
var PayEnumTradeStatus_value = map[string]int32{
	"PAY_S_UNIFIEDORDER": 1,
	"PAY_S_WATINGPAY":    2,
	"PAY_S_SUCC":         3,
}

func (x PayEnumTradeStatus) Enum() *PayEnumTradeStatus {
	p := new(PayEnumTradeStatus)
	*p = x
	return p
}
func (x PayEnumTradeStatus) String() string {
	return proto.EnumName(PayEnumTradeStatus_name, int32(x))
}
func (x *PayEnumTradeStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PayEnumTradeStatus_value, data, "PayEnumTradeStatus")
	if err != nil {
		return err
	}
	*x = PayEnumTradeStatus(value)
	return nil
}
func (PayEnumTradeStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// app收到同步回调之后，请求服务器检测接收异步回调，得到充值之后的信息
type PayReturn int32

const (
	PayReturn_ERR_ONLY_APPLE PayReturn = -250
)

var PayReturn_name = map[int32]string{
	-250: "ERR_ONLY_APPLE",
}
var PayReturn_value = map[string]int32{
	"ERR_ONLY_APPLE": -250,
}

func (x PayReturn) Enum() *PayReturn {
	p := new(PayReturn)
	*p = x
	return p
}
func (x PayReturn) String() string {
	return proto.EnumName(PayReturn_name, int32(x))
}
func (x *PayReturn) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PayReturn_value, data, "PayReturn")
	if err != nil {
		return err
	}
	*x = PayReturn(value)
	return nil
}
func (PayReturn) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

// 充值套餐
type PayBaseProduct struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Money            *int64  `protobuf:"varint,2,opt,name=money" json:"money,omitempty"`
	Diamond          *int64  `protobuf:"varint,3,opt,name=diamond" json:"diamond,omitempty"`
	Memo             *string `protobuf:"bytes,4,opt,name=memo" json:"memo,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PayBaseProduct) Reset()                    { *m = PayBaseProduct{} }
func (m *PayBaseProduct) String() string            { return proto.CompactTextString(m) }
func (*PayBaseProduct) ProtoMessage()               {}
func (*PayBaseProduct) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *PayBaseProduct) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *PayBaseProduct) GetMoney() int64 {
	if m != nil && m.Money != nil {
		return *m.Money
	}
	return 0
}

func (m *PayBaseProduct) GetDiamond() int64 {
	if m != nil && m.Diamond != nil {
		return *m.Diamond
	}
	return 0
}

func (m *PayBaseProduct) GetMemo() string {
	if m != nil && m.Memo != nil {
		return *m.Memo
	}
	return ""
}

// 付款的模式,制定对应的商户号，appid,appkey
type PayBasePaymodel struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	AppId            *string `protobuf:"bytes,2,opt,name=appId" json:"appId,omitempty"`
	MchId            *string `protobuf:"bytes,3,opt,name=mchId" json:"mchId,omitempty"`
	AppKey           *string `protobuf:"bytes,4,opt,name=appKey" json:"appKey,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PayBasePaymodel) Reset()                    { *m = PayBasePaymodel{} }
func (m *PayBasePaymodel) String() string            { return proto.CompactTextString(m) }
func (*PayBasePaymodel) ProtoMessage()               {}
func (*PayBasePaymodel) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *PayBasePaymodel) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *PayBasePaymodel) GetAppId() string {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return ""
}

func (m *PayBasePaymodel) GetMchId() string {
	if m != nil && m.MchId != nil {
		return *m.MchId
	}
	return ""
}

func (m *PayBasePaymodel) GetAppKey() string {
	if m != nil && m.AppKey != nil {
		return *m.AppKey
	}
	return ""
}

// 支付明细
type PayBaseDetails struct {
	Id               *int32              `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	UserId           *uint32             `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	PayModelId       *int32              `protobuf:"varint,3,opt,name=payModelId" json:"payModelId,omitempty"`
	ProductId        *int32              `protobuf:"varint,4,opt,name=productId" json:"productId,omitempty"`
	TradeNo          *string             `protobuf:"bytes,5,opt,name=tradeNo" json:"tradeNo,omitempty"`
	Status           *PayEnumTradeStatus `protobuf:"varint,6,opt,name=status,enum=ddproto.PayEnumTradeStatus" json:"status,omitempty"`
	Diamond          *int64              `protobuf:"varint,7,opt,name=Diamond" json:"Diamond,omitempty"`
	ChangeDiamond    *int64              `protobuf:"varint,8,opt,name=ChangeDiamond" json:"ChangeDiamond,omitempty"`
	Coin             *int64              `protobuf:"varint,9,opt,name=Coin" json:"Coin,omitempty"`
	ChangeCoin       *int64              `protobuf:"varint,10,opt,name=ChangeCoin" json:"ChangeCoin,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *PayBaseDetails) Reset()                    { *m = PayBaseDetails{} }
func (m *PayBaseDetails) String() string            { return proto.CompactTextString(m) }
func (*PayBaseDetails) ProtoMessage()               {}
func (*PayBaseDetails) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *PayBaseDetails) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *PayBaseDetails) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PayBaseDetails) GetPayModelId() int32 {
	if m != nil && m.PayModelId != nil {
		return *m.PayModelId
	}
	return 0
}

func (m *PayBaseDetails) GetProductId() int32 {
	if m != nil && m.ProductId != nil {
		return *m.ProductId
	}
	return 0
}

func (m *PayBaseDetails) GetTradeNo() string {
	if m != nil && m.TradeNo != nil {
		return *m.TradeNo
	}
	return ""
}

func (m *PayBaseDetails) GetStatus() PayEnumTradeStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return PayEnumTradeStatus_PAY_S_UNIFIEDORDER
}

func (m *PayBaseDetails) GetDiamond() int64 {
	if m != nil && m.Diamond != nil {
		return *m.Diamond
	}
	return 0
}

func (m *PayBaseDetails) GetChangeDiamond() int64 {
	if m != nil && m.ChangeDiamond != nil {
		return *m.ChangeDiamond
	}
	return 0
}

func (m *PayBaseDetails) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *PayBaseDetails) GetChangeCoin() int64 {
	if m != nil && m.ChangeCoin != nil {
		return *m.ChangeCoin
	}
	return 0
}

// 统一下单 proto
//    总金额	total_fee	是	Int	888	订单总金额，单位为分，详见支付金额
type WxpayReqUnifiedorder struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ProductId        *int32       `protobuf:"varint,2,opt,name=productId" json:"productId,omitempty"`
	PayModelId       *int32       `protobuf:"varint,3,opt,name=payModelId" json:"payModelId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *WxpayReqUnifiedorder) Reset()                    { *m = WxpayReqUnifiedorder{} }
func (m *WxpayReqUnifiedorder) String() string            { return proto.CompactTextString(m) }
func (*WxpayReqUnifiedorder) ProtoMessage()               {}
func (*WxpayReqUnifiedorder) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *WxpayReqUnifiedorder) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *WxpayReqUnifiedorder) GetProductId() int32 {
	if m != nil && m.ProductId != nil {
		return *m.ProductId
	}
	return 0
}

func (m *WxpayReqUnifiedorder) GetPayModelId() int32 {
	if m != nil && m.PayModelId != nil {
		return *m.PayModelId
	}
	return 0
}

// app请求统一下单之后，返回的加密字符串
type WxpayAckUnifiedorder struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PartnerId        *string      `protobuf:"bytes,2,opt,name=partnerId" json:"partnerId,omitempty"`
	PrepayId         *string      `protobuf:"bytes,3,opt,name=prepay_id" json:"prepay_id,omitempty"`
	NonceStr         *string      `protobuf:"bytes,4,opt,name=nonce_str" json:"nonce_str,omitempty"`
	TimeStamp        *string      `protobuf:"bytes,5,opt,name=timeStamp" json:"timeStamp,omitempty"`
	Package          *string      `protobuf:"bytes,6,opt,name=package" json:"package,omitempty"`
	Sign             *string      `protobuf:"bytes,7,opt,name=sign" json:"sign,omitempty"`
	TradeNo          *string      `protobuf:"bytes,8,opt,name=tradeNo" json:"tradeNo,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *WxpayAckUnifiedorder) Reset()                    { *m = WxpayAckUnifiedorder{} }
func (m *WxpayAckUnifiedorder) String() string            { return proto.CompactTextString(m) }
func (*WxpayAckUnifiedorder) ProtoMessage()               {}
func (*WxpayAckUnifiedorder) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *WxpayAckUnifiedorder) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *WxpayAckUnifiedorder) GetPartnerId() string {
	if m != nil && m.PartnerId != nil {
		return *m.PartnerId
	}
	return ""
}

func (m *WxpayAckUnifiedorder) GetPrepayId() string {
	if m != nil && m.PrepayId != nil {
		return *m.PrepayId
	}
	return ""
}

func (m *WxpayAckUnifiedorder) GetNonceStr() string {
	if m != nil && m.NonceStr != nil {
		return *m.NonceStr
	}
	return ""
}

func (m *WxpayAckUnifiedorder) GetTimeStamp() string {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return ""
}

func (m *WxpayAckUnifiedorder) GetPackage() string {
	if m != nil && m.Package != nil {
		return *m.Package
	}
	return ""
}

func (m *WxpayAckUnifiedorder) GetSign() string {
	if m != nil && m.Sign != nil {
		return *m.Sign
	}
	return ""
}

func (m *WxpayAckUnifiedorder) GetTradeNo() string {
	if m != nil && m.TradeNo != nil {
		return *m.TradeNo
	}
	return ""
}

type WxpayReqSyncChecker struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	TradeNo          *string      `protobuf:"bytes,2,opt,name=tradeNo" json:"tradeNo,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *WxpayReqSyncChecker) Reset()                    { *m = WxpayReqSyncChecker{} }
func (m *WxpayReqSyncChecker) String() string            { return proto.CompactTextString(m) }
func (*WxpayReqSyncChecker) ProtoMessage()               {}
func (*WxpayReqSyncChecker) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *WxpayReqSyncChecker) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *WxpayReqSyncChecker) GetTradeNo() string {
	if m != nil && m.TradeNo != nil {
		return *m.TradeNo
	}
	return ""
}

// 服务器收到微信充值的异步回调之后，返回给app当前的余额信息
type WxpayAckSyncChecker struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	WxpayId          *int32       `protobuf:"varint,2,opt,name=wxpayId" json:"wxpayId,omitempty"`
	Diamond          *int64       `protobuf:"varint,3,opt,name=diamond" json:"diamond,omitempty"`
	ChangeDiamond    *int64       `protobuf:"varint,4,opt,name=changeDiamond" json:"changeDiamond,omitempty"`
	Coin             *int64       `protobuf:"varint,5,opt,name=coin" json:"coin,omitempty"`
	ChangeCoin       *int64       `protobuf:"varint,6,opt,name=changeCoin" json:"changeCoin,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *WxpayAckSyncChecker) Reset()                    { *m = WxpayAckSyncChecker{} }
func (m *WxpayAckSyncChecker) String() string            { return proto.CompactTextString(m) }
func (*WxpayAckSyncChecker) ProtoMessage()               {}
func (*WxpayAckSyncChecker) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *WxpayAckSyncChecker) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *WxpayAckSyncChecker) GetWxpayId() int32 {
	if m != nil && m.WxpayId != nil {
		return *m.WxpayId
	}
	return 0
}

func (m *WxpayAckSyncChecker) GetDiamond() int64 {
	if m != nil && m.Diamond != nil {
		return *m.Diamond
	}
	return 0
}

func (m *WxpayAckSyncChecker) GetChangeDiamond() int64 {
	if m != nil && m.ChangeDiamond != nil {
		return *m.ChangeDiamond
	}
	return 0
}

func (m *WxpayAckSyncChecker) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *WxpayAckSyncChecker) GetChangeCoin() int64 {
	if m != nil && m.ChangeCoin != nil {
		return *m.ChangeCoin
	}
	return 0
}

// 苹果支付充值回调
type ApplepayReqRechargecb struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ProductId        *int32       `protobuf:"varint,2,opt,name=productId" json:"productId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ApplepayReqRechargecb) Reset()                    { *m = ApplepayReqRechargecb{} }
func (m *ApplepayReqRechargecb) String() string            { return proto.CompactTextString(m) }
func (*ApplepayReqRechargecb) ProtoMessage()               {}
func (*ApplepayReqRechargecb) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *ApplepayReqRechargecb) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ApplepayReqRechargecb) GetProductId() int32 {
	if m != nil && m.ProductId != nil {
		return *m.ProductId
	}
	return 0
}

type ApplepayAcksRechargecb struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	RechargeDiamond  *int64       `protobuf:"varint,2,opt,name=rechargeDiamond" json:"rechargeDiamond,omitempty"`
	Diamond          *int64       `protobuf:"varint,3,opt,name=diamond" json:"diamond,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ApplepayAcksRechargecb) Reset()                    { *m = ApplepayAcksRechargecb{} }
func (m *ApplepayAcksRechargecb) String() string            { return proto.CompactTextString(m) }
func (*ApplepayAcksRechargecb) ProtoMessage()               {}
func (*ApplepayAcksRechargecb) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *ApplepayAcksRechargecb) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ApplepayAcksRechargecb) GetRechargeDiamond() int64 {
	if m != nil && m.RechargeDiamond != nil {
		return *m.RechargeDiamond
	}
	return 0
}

func (m *ApplepayAcksRechargecb) GetDiamond() int64 {
	if m != nil && m.Diamond != nil {
		return *m.Diamond
	}
	return 0
}

func init() {
	proto.RegisterType((*PayBaseProduct)(nil), "ddproto.pay_base_product")
	proto.RegisterType((*PayBasePaymodel)(nil), "ddproto.pay_base_paymodel")
	proto.RegisterType((*PayBaseDetails)(nil), "ddproto.pay_base_details")
	proto.RegisterType((*WxpayReqUnifiedorder)(nil), "ddproto.wxpay_req_unifiedorder")
	proto.RegisterType((*WxpayAckUnifiedorder)(nil), "ddproto.wxpay_ack_unifiedorder")
	proto.RegisterType((*WxpayReqSyncChecker)(nil), "ddproto.wxpay_req_syncChecker")
	proto.RegisterType((*WxpayAckSyncChecker)(nil), "ddproto.wxpay_ack_syncChecker")
	proto.RegisterType((*ApplepayReqRechargecb)(nil), "ddproto.applepay_req_rechargecb")
	proto.RegisterType((*ApplepayAcksRechargecb)(nil), "ddproto.applepay_acks_rechargecb")
	proto.RegisterEnum("ddproto.PayEnumTradeStatus", PayEnumTradeStatus_name, PayEnumTradeStatus_value)
	proto.RegisterEnum("ddproto.PayReturn", PayReturn_name, PayReturn_value)
}

var fileDescriptor2 = []byte{
	// 592 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x8e, 0xd3, 0x3c,
	0x10, 0xfd, 0xd2, 0x6d, 0xbb, 0x9b, 0xf9, 0xe8, 0xcf, 0x66, 0x7f, 0x1a, 0x81, 0x90, 0x50, 0xc4,
	0x05, 0xac, 0x44, 0x2f, 0xf6, 0x0d, 0xaa, 0xb6, 0x40, 0x45, 0xb7, 0x5b, 0xb5, 0xac, 0xd0, 0x5e,
	0x45, 0x5e, 0xdb, 0xb4, 0x61, 0x1b, 0x27, 0x38, 0x8e, 0xa0, 0x2f, 0xc0, 0x33, 0xf0, 0x2c, 0xbc,
	0x0c, 0x6f, 0x02, 0x8c, 0x9d, 0xa4, 0x3f, 0xd0, 0x0b, 0x68, 0x2f, 0x2a, 0xfb, 0x78, 0xe6, 0xcc,
	0xcc, 0xf1, 0x71, 0xa0, 0x45, 0xa3, 0x30, 0x8c, 0x84, 0x4f, 0x17, 0x01, 0x17, 0xca, 0x8f, 0xc9,
	0xb2, 0x1d, 0xcb, 0x48, 0x45, 0xce, 0x21, 0x63, 0x66, 0xf1, 0xf0, 0x64, 0x2b, 0x22, 0x3b, 0xf5,
	0x86, 0xd0, 0xc4, 0x50, 0xff, 0x8e, 0x24, 0xdc, 0x47, 0x84, 0xa5, 0x54, 0x39, 0x00, 0xa5, 0x80,
	0xb9, 0xd6, 0x13, 0xeb, 0x59, 0xc5, 0xa9, 0x41, 0x05, 0x73, 0xf8, 0xd2, 0x2d, 0xe1, 0xf6, 0xc0,
	0x69, 0xc0, 0x21, 0x0b, 0x08, 0x22, 0xcc, 0x3d, 0x30, 0xc0, 0x03, 0x28, 0x87, 0x3c, 0x8c, 0xdc,
	0x32, 0xee, 0x6c, 0xef, 0x0a, 0x8e, 0xd7, 0x6c, 0x64, 0x19, 0x46, 0x8c, 0x2f, 0x7e, 0xa7, 0x23,
	0x71, 0x3c, 0x60, 0x86, 0xce, 0x36, 0xec, 0x74, 0x3e, 0xc8, 0xc8, 0x6c, 0xa7, 0x0e, 0x55, 0x3c,
	0x7d, 0x83, 0xd5, 0x32, 0xba, 0xef, 0xd6, 0x46, 0x77, 0x8c, 0x2b, 0x12, 0x2c, 0x92, 0x2d, 0x3a,
	0x4c, 0x48, 0x13, 0x2e, 0x73, 0xbe, 0x9a, 0x83, 0x87, 0x18, 0x7f, 0xa5, 0xcb, 0xe6, 0xa4, 0x15,
	0xe7, 0x18, 0xec, 0x7c, 0x30, 0x84, 0xca, 0x06, 0xc2, 0x29, 0x94, 0x24, 0x8c, 0x8f, 0x22, 0xb7,
	0x62, 0x0a, 0xbf, 0x80, 0x6a, 0xa2, 0x88, 0x4a, 0x13, 0xb7, 0x8a, 0xfb, 0xfa, 0xe5, 0xe3, 0x76,
	0x2e, 0x5a, 0x5b, 0x97, 0xe7, 0x22, 0x0d, 0x7d, 0x93, 0x30, 0x35, 0x41, 0x3a, 0xbf, 0x97, 0xab,
	0x70, 0x68, 0x54, 0x38, 0x83, 0x5a, 0x77, 0x4e, 0xc4, 0x8c, 0x17, 0xf0, 0x51, 0x21, 0x4e, 0x37,
	0x0a, 0x84, 0x6b, 0x9b, 0x1d, 0x36, 0x97, 0x05, 0x19, 0x0c, 0x34, 0xe6, 0x71, 0x38, 0xff, 0xf4,
	0x59, 0xd7, 0x90, 0xfc, 0xa3, 0x9f, 0x8a, 0xe0, 0x7d, 0xc0, 0x59, 0x24, 0x19, 0x97, 0xce, 0x53,
	0xa8, 0xce, 0x39, 0x96, 0x94, 0x66, 0xd4, 0xff, 0x2f, 0x4f, 0x57, 0x2d, 0x8d, 0xf5, 0xff, 0x6b,
	0x73, 0xb6, 0x3d, 0x5c, 0xc9, 0x0c, 0xb7, 0x43, 0x03, 0xef, 0x9b, 0x55, 0xd4, 0x21, 0xf4, 0x7e,
	0xdf, 0x3a, 0x44, 0x2a, 0xb1, 0xd2, 0xda, 0xce, 0x4a, 0x73, 0xcd, 0x19, 0x14, 0xf7, 0x87, 0x90,
	0x88, 0x04, 0xe5, 0x7e, 0xa2, 0x64, 0x76, 0x85, 0x1a, 0x52, 0x41, 0xa8, 0x85, 0x0b, 0xe3, 0x5c,
	0x6c, 0x54, 0x2f, 0xc6, 0x36, 0xc8, 0x8c, 0x1b, 0xb5, 0x6d, 0x2d, 0x53, 0x12, 0xcc, 0x84, 0xd1,
	0xd2, 0xde, 0xbc, 0x9c, 0x23, 0xe3, 0x82, 0x11, 0x9c, 0xad, 0x35, 0x4a, 0x96, 0x82, 0x76, 0xe7,
	0x9c, 0xde, 0xff, 0x75, 0xeb, 0x1b, 0x7c, 0xa6, 0x71, 0xef, 0xab, 0x55, 0x10, 0x6a, 0x31, 0xf6,
	0x22, 0x34, 0xe9, 0x2b, 0xc5, 0xff, 0x78, 0x14, 0x68, 0x07, 0xba, 0x65, 0x87, 0x72, 0x61, 0x07,
	0xaa, 0xaf, 0xbe, 0x52, 0xd8, 0x81, 0xae, 0xed, 0x50, 0x35, 0x76, 0x98, 0x40, 0x0b, 0x1f, 0xc0,
	0x82, 0x17, 0xd3, 0x4a, 0x8e, 0x21, 0x72, 0xc6, 0xe9, 0xdd, 0xde, 0x7e, 0xf0, 0x3e, 0x80, 0xbb,
	0xe2, 0xc4, 0x81, 0x93, 0x7f, 0x27, 0x6d, 0x41, 0xa3, 0xc8, 0x29, 0x06, 0xda, 0xfd, 0x35, 0xb8,
	0x98, 0xc2, 0xe9, 0xce, 0x07, 0x73, 0x0e, 0xce, 0xb8, 0x73, 0xeb, 0x4f, 0xfd, 0x9b, 0xd1, 0xe0,
	0xe5, 0xa0, 0xdf, 0xbb, 0x9e, 0xf4, 0xfa, 0x93, 0xa6, 0xe5, 0x9c, 0x40, 0x23, 0xc3, 0xdf, 0x75,
	0xde, 0x0e, 0x46, 0xaf, 0x70, 0xdd, 0x2c, 0xe1, 0xa3, 0x86, 0x0c, 0x9c, 0xde, 0x74, 0xbb, 0xcd,
	0x83, 0x8b, 0xe7, 0xc6, 0xd0, 0xd8, 0xb6, 0x4a, 0xa5, 0x70, 0x1e, 0x41, 0xbd, 0x3f, 0x99, 0xf8,
	0xd7, 0xa3, 0xe1, 0xad, 0xdf, 0x19, 0x8f, 0x87, 0xfd, 0xe6, 0x97, 0x1f, 0x3f, 0xb3, 0x9f, 0x35,
	0xfe, 0xef, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x7c, 0x07, 0xff, 0x07, 0x05, 0x00, 0x00,
}
