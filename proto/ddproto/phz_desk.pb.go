// Code generated by protoc-gen-go. DO NOT EDIT.
// source: phz_desk.proto

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

// Ignoring public import of common_req_kickout from common_client.proto

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

// Ignoring public import of common_req_upload_location from common_client.proto

// Ignoring public import of common_bc_leaveTimeout from common_client.proto

// Ignoring public import of common_desk_by_agent from common_client.proto

// Ignoring public import of common_req_list_coin_desk from common_client.proto

// Ignoring public import of common_ack_list_coin_desk from common_client.proto

// Ignoring public import of CommonCoinDeskInfo from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// Ignoring public import of phz_play_paiIds from phz_base.proto

// Ignoring public import of phz_base_createOption from phz_base.proto

// Ignoring public import of phz_base_roomInfo from phz_base.proto

// Ignoring public import of phz_base_deskInfo from phz_base.proto

// Ignoring public import of phz_base_playerInfo from phz_base.proto

// Ignoring public import of phz_enum_protoId from phz_base.proto

// Ignoring public import of phz_enum_roomType from phz_base.proto

// Ignoring public import of phz_enum_tiType from phz_base.proto

// Ignoring public import of phz_enum_pengType from phz_base.proto

// Ignoring public import of phz_enum_paoType from phz_base.proto

// Ignoring public import of phz_enum_deskStatus from phz_base.proto

// Ignoring public import of phz_enum_userStatus from phz_base.proto

// Ignoring public import of phz_req_Ready from phz_play.proto

// Ignoring public import of phz_ack_Ready from phz_play.proto

// Ignoring public import of phz_play_ReadyStatus from phz_play.proto

// Ignoring public import of phz_play_Opening from phz_play.proto

// Ignoring public import of phz_play_SendCards from phz_play.proto

// Ignoring public import of phz_req_OutCards from phz_play.proto

// Ignoring public import of phz_ack_OutCards from phz_play.proto

// Ignoring public import of phz_play_MoPai from phz_play.proto

// Ignoring public import of phz_play_Overturn from phz_play.proto

// Ignoring public import of phz_play_CanPeng from phz_play.proto

// Ignoring public import of phz_req_Peng from phz_play.proto

// Ignoring public import of phz_ack_Peng from phz_play.proto

// Ignoring public import of phz_play_ChiPai from phz_play.proto

// Ignoring public import of phz_play_canChi from phz_play.proto

// Ignoring public import of phz_req_chiPai from phz_play.proto

// Ignoring public import of phz_req_biPai from phz_play.proto

// Ignoring public import of phz_ack_biPai from phz_play.proto

// Ignoring public import of phz_ack_chiPai from phz_play.proto

// Ignoring public import of phz_play_tiPai from phz_play.proto

// Ignoring public import of phz_play_canTiPai from phz_play.proto

// Ignoring public import of phz_play_weiPai from phz_play.proto

// Ignoring public import of phz_play_paoPai from phz_play.proto

// Ignoring public import of phz_play_canHuPai from phz_play.proto

// Ignoring public import of phz_req_huPai from phz_play.proto

// Ignoring public import of phz_ack_huPai from phz_play.proto

// Ignoring public import of phz_req_pass from phz_play.proto

// Ignoring public import of phz_ack_pass from phz_play.proto

// Ignoring public import of phz_play_handPokers from phz_play.proto

// Ignoring public import of phz_play_roundBillBean from phz_play.proto

// Ignoring public import of phz_play_currentResult from phz_play.proto

// Ignoring public import of phz_play_userLotteryInfo from phz_play.proto

// Ignoring public import of phz_play_userEndResult from phz_play.proto

// Ignoring public import of phz_play_endLottery from phz_play.proto

// Ignoring public import of phz_enum_actionType from phz_play.proto

// Ignoring public import of phz_enum_huType from phz_play.proto

// 创建房间
type PhzReqCreateDesk struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	RoomInfo         *PhzBaseCreateOption `protobuf:"bytes,2,opt,name=roomInfo" json:"roomInfo,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PhzReqCreateDesk) Reset()                    { *m = PhzReqCreateDesk{} }
func (m *PhzReqCreateDesk) String() string            { return proto.CompactTextString(m) }
func (*PhzReqCreateDesk) ProtoMessage()               {}
func (*PhzReqCreateDesk) Descriptor() ([]byte, []int) { return fileDescriptor46, []int{0} }

func (m *PhzReqCreateDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PhzReqCreateDesk) GetRoomInfo() *PhzBaseCreateOption {
	if m != nil {
		return m.RoomInfo
	}
	return nil
}

// 进入房间
type PhzReqEnterDesk struct {
	Header           *ProtoHeader     `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	RoomInfo         *PhzBaseRoomInfo `protobuf:"bytes,2,opt,name=roomInfo" json:"roomInfo,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *PhzReqEnterDesk) Reset()                    { *m = PhzReqEnterDesk{} }
func (m *PhzReqEnterDesk) String() string            { return proto.CompactTextString(m) }
func (*PhzReqEnterDesk) ProtoMessage()               {}
func (*PhzReqEnterDesk) Descriptor() ([]byte, []int) { return fileDescriptor46, []int{1} }

func (m *PhzReqEnterDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PhzReqEnterDesk) GetRoomInfo() *PhzBaseRoomInfo {
	if m != nil {
		return m.RoomInfo
	}
	return nil
}

// 进入、创建房间回复的ack
type PhzAckDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	IsCreateDeskOK   *bool        `protobuf:"varint,2,opt,name=isCreateDeskOK" json:"isCreateDeskOK,omitempty"`
	IsEnterDeskOK    *bool        `protobuf:"varint,3,opt,name=isEnterDeskOK" json:"isEnterDeskOK,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PhzAckDesk) Reset()                    { *m = PhzAckDesk{} }
func (m *PhzAckDesk) String() string            { return proto.CompactTextString(m) }
func (*PhzAckDesk) ProtoMessage()               {}
func (*PhzAckDesk) Descriptor() ([]byte, []int) { return fileDescriptor46, []int{2} }

func (m *PhzAckDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PhzAckDesk) GetIsCreateDeskOK() bool {
	if m != nil && m.IsCreateDeskOK != nil {
		return *m.IsCreateDeskOK
	}
	return false
}

func (m *PhzAckDesk) GetIsEnterDeskOK() bool {
	if m != nil && m.IsEnterDeskOK != nil {
		return *m.IsEnterDeskOK
	}
	return false
}

type PhzReqGameInfo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PhzReqGameInfo) Reset()                    { *m = PhzReqGameInfo{} }
func (m *PhzReqGameInfo) String() string            { return proto.CompactTextString(m) }
func (*PhzReqGameInfo) ProtoMessage()               {}
func (*PhzReqGameInfo) Descriptor() ([]byte, []int) { return fileDescriptor46, []int{3} }

func (m *PhzReqGameInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type PhzDeskGameInfo struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	DeskInfo         *PhzBaseDeskInfo     `protobuf:"bytes,2,opt,name=deskInfo" json:"deskInfo,omitempty"`
	PlayerInfo       []*PhzBasePlayerInfo `protobuf:"bytes,3,rep,name=playerInfo" json:"playerInfo,omitempty"`
	IsReconnect      *bool                `protobuf:"varint,4,opt,name=isReconnect" json:"isReconnect,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PhzDeskGameInfo) Reset()                    { *m = PhzDeskGameInfo{} }
func (m *PhzDeskGameInfo) String() string            { return proto.CompactTextString(m) }
func (*PhzDeskGameInfo) ProtoMessage()               {}
func (*PhzDeskGameInfo) Descriptor() ([]byte, []int) { return fileDescriptor46, []int{4} }

func (m *PhzDeskGameInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PhzDeskGameInfo) GetDeskInfo() *PhzBaseDeskInfo {
	if m != nil {
		return m.DeskInfo
	}
	return nil
}

func (m *PhzDeskGameInfo) GetPlayerInfo() []*PhzBasePlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *PhzDeskGameInfo) GetIsReconnect() bool {
	if m != nil && m.IsReconnect != nil {
		return *m.IsReconnect
	}
	return false
}

// 房主直接解散房间
type PhzReqDissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PhzReqDissolveDesk) Reset()                    { *m = PhzReqDissolveDesk{} }
func (m *PhzReqDissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*PhzReqDissolveDesk) ProtoMessage()               {}
func (*PhzReqDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor46, []int{5} }

func (m *PhzReqDissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PhzReqDissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

type PhzAckDissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	DeskId           *int32       `protobuf:"varint,2,opt,name=deskId" json:"deskId,omitempty"`
	Password         *string      `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	UserId           *uint32      `protobuf:"varint,4,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PhzAckDissolveDesk) Reset()                    { *m = PhzAckDissolveDesk{} }
func (m *PhzAckDissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*PhzAckDissolveDesk) ProtoMessage()               {}
func (*PhzAckDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor46, []int{6} }

func (m *PhzAckDissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PhzAckDissolveDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *PhzAckDissolveDesk) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *PhzAckDissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func init() {
	proto.RegisterType((*PhzReqCreateDesk)(nil), "ddproto.phz_req_createDesk")
	proto.RegisterType((*PhzReqEnterDesk)(nil), "ddproto.phz_req_enterDesk")
	proto.RegisterType((*PhzAckDesk)(nil), "ddproto.phz_ack_desk")
	proto.RegisterType((*PhzReqGameInfo)(nil), "ddproto.phz_req_gameInfo")
	proto.RegisterType((*PhzDeskGameInfo)(nil), "ddproto.phz_desk_gameInfo")
	proto.RegisterType((*PhzReqDissolveDesk)(nil), "ddproto.phz_req_dissolveDesk")
	proto.RegisterType((*PhzAckDissolveDesk)(nil), "ddproto.phz_ack_dissolveDesk")
}

func init() { proto.RegisterFile("phz_desk.proto", fileDescriptor46) }

var fileDescriptor46 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x4a, 0xf3, 0x40,
	0x10, 0xc7, 0xbf, 0xb4, 0xfd, 0x6a, 0x9d, 0xda, 0xa2, 0x6b, 0x91, 0x10, 0x44, 0x4a, 0x10, 0xe9,
	0x41, 0x7a, 0xe8, 0xc1, 0x83, 0x78, 0x10, 0x54, 0xb0, 0xf4, 0xd0, 0xb0, 0x67, 0xa1, 0xc4, 0x64,
	0xb4, 0xa1, 0xcd, 0x6e, 0xdc, 0x8d, 0x4a, 0x3d, 0xf8, 0x00, 0x9e, 0x7c, 0x3d, 0xdf, 0x46, 0x76,
	0x93, 0x8d, 0xd1, 0x82, 0xd0, 0x7a, 0x29, 0x9d, 0xd9, 0xdf, 0xcc, 0x7f, 0xe6, 0x9f, 0x81, 0x76,
	0x32, 0x7d, 0x99, 0x84, 0x28, 0x67, 0xfd, 0x44, 0xf0, 0x94, 0x93, 0x8d, 0x30, 0xd4, 0x7f, 0x9c,
	0xdd, 0x80, 0xc7, 0x31, 0x67, 0x93, 0x60, 0x1e, 0x21, 0x4b, 0xb3, 0x57, 0x47, 0xd3, 0xb7, 0xbe,
	0xc4, 0x72, 0x9c, 0xcc, 0xfd, 0x45, 0x16, 0xbb, 0xaf, 0x40, 0x54, 0x46, 0xe0, 0xc3, 0x24, 0x10,
	0xe8, 0xa7, 0x78, 0x89, 0x72, 0x46, 0x8e, 0xa1, 0x3e, 0x45, 0x3f, 0x44, 0x61, 0x5b, 0x5d, 0xab,
	0xd7, 0x1c, 0x74, 0xfa, 0xb9, 0x48, 0xdf, 0x53, 0xbf, 0xd7, 0xfa, 0x8d, 0xe6, 0x0c, 0x39, 0x85,
	0x86, 0xe0, 0x3c, 0x1e, 0xb2, 0x3b, 0x6e, 0x57, 0x34, 0x7f, 0x50, 0xf0, 0x46, 0x3e, 0xef, 0x3e,
	0x4e, 0xd2, 0x88, 0x33, 0x5a, 0xf0, 0xee, 0x02, 0x76, 0x8c, 0x3e, 0xb2, 0x14, 0xc5, 0x1a, 0xf2,
	0x27, 0x4b, 0xf2, 0xce, 0xb2, 0xbc, 0x21, 0x4a, 0xd2, 0x6f, 0x16, 0x6c, 0xa9, 0x77, 0x3f, 0x98,
	0x69, 0x3f, 0x57, 0x94, 0x3d, 0x82, 0x76, 0x24, 0x2f, 0x0a, 0xcf, 0xc6, 0x23, 0x2d, 0xde, 0xa0,
	0x3f, 0xb2, 0xe4, 0x10, 0x5a, 0x91, 0xbc, 0x32, 0xbb, 0x8d, 0x47, 0x76, 0x55, 0x63, 0xdf, 0x93,
	0xee, 0x39, 0x6c, 0x1b, 0x1f, 0xee, 0xfd, 0x18, 0xd5, 0x80, 0xab, 0xcd, 0xe3, 0x7e, 0x58, 0x99,
	0x95, 0x6a, 0x95, 0x35, 0x7b, 0x28, 0x2b, 0x55, 0xf9, 0xef, 0x56, 0x1a, 0x82, 0x16, 0x2c, 0x39,
	0x03, 0x50, 0x37, 0x85, 0x42, 0x57, 0x56, 0xbb, 0xd5, 0x5e, 0x73, 0xb0, 0xbf, 0x5c, 0xf9, 0xc5,
	0xd0, 0x12, 0x4f, 0xba, 0xd0, 0x8c, 0x24, 0xc5, 0x80, 0x33, 0x86, 0x41, 0x6a, 0xd7, 0xb4, 0x3f,
	0xe5, 0x94, 0x7b, 0x03, 0x1d, 0xe3, 0x4e, 0x18, 0x49, 0xc9, 0xe7, 0x4f, 0xeb, 0xdc, 0xe9, 0x1e,
	0xd4, 0x1f, 0x25, 0x8a, 0x61, 0xa8, 0x77, 0x6b, 0xd1, 0x3c, 0x72, 0xdf, 0xad, 0xac, 0xbd, 0x3e,
	0x84, 0x3f, 0xb5, 0xd7, 0x86, 0x64, 0xed, 0xff, 0xd3, 0x3c, 0x22, 0x0e, 0x34, 0x12, 0x5f, 0xca,
	0x67, 0x2e, 0x42, 0xfd, 0xed, 0x37, 0x69, 0x11, 0x97, 0x46, 0xaa, 0x95, 0x47, 0xf2, 0xfe, 0x79,
	0x96, 0x57, 0xf9, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xb8, 0x84, 0x87, 0xeb, 0x03, 0x00, 0x00,
}
