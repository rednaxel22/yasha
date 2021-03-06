// Code generated by protoc-gen-go.
// source: steamdatagram_messages.proto
// DO NOT EDIT!

package dota

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type ESteamDatagramMsgID int32

const (
	ESteamDatagramMsgID_k_ESteamDatagramMsg_RouterPingRequest            ESteamDatagramMsgID = 1
	ESteamDatagramMsgID_k_ESteamDatagramMsg_RouterPingReply              ESteamDatagramMsgID = 2
	ESteamDatagramMsgID_k_ESteamDatagramMsg_GameserverPingRequest        ESteamDatagramMsgID = 3
	ESteamDatagramMsgID_k_ESteamDatagramMsg_GameserverPingReply          ESteamDatagramMsgID = 4
	ESteamDatagramMsgID_k_ESteamDatagramMsg_GameserverSessionRequest     ESteamDatagramMsgID = 5
	ESteamDatagramMsgID_k_ESteamDatagramMsg_GameserverSessionEstablished ESteamDatagramMsgID = 6
	ESteamDatagramMsgID_k_ESteamDatagramMsg_NoSession                    ESteamDatagramMsgID = 7
	ESteamDatagramMsgID_k_ESteamDatagramMsg_Diagnostic                   ESteamDatagramMsgID = 8
	ESteamDatagramMsgID_k_ESteamDatagramMsg_DataClientToRouter           ESteamDatagramMsgID = 9
	ESteamDatagramMsgID_k_ESteamDatagramMsg_DataRouterToServer           ESteamDatagramMsgID = 10
	ESteamDatagramMsgID_k_ESteamDatagramMsg_DataServerToRouter           ESteamDatagramMsgID = 11
	ESteamDatagramMsgID_k_ESteamDatagramMsg_DataRouterToClient           ESteamDatagramMsgID = 12
	ESteamDatagramMsgID_k_ESteamDatagramMsg_Stats                        ESteamDatagramMsgID = 13
	ESteamDatagramMsgID_k_ESteamDatagramMsg_ClientPingSampleRequest      ESteamDatagramMsgID = 14
	ESteamDatagramMsgID_k_ESteamDatagramMsg_ClientPingSampleReply        ESteamDatagramMsgID = 15
)

var ESteamDatagramMsgID_name = map[int32]string{
	1:  "k_ESteamDatagramMsg_RouterPingRequest",
	2:  "k_ESteamDatagramMsg_RouterPingReply",
	3:  "k_ESteamDatagramMsg_GameserverPingRequest",
	4:  "k_ESteamDatagramMsg_GameserverPingReply",
	5:  "k_ESteamDatagramMsg_GameserverSessionRequest",
	6:  "k_ESteamDatagramMsg_GameserverSessionEstablished",
	7:  "k_ESteamDatagramMsg_NoSession",
	8:  "k_ESteamDatagramMsg_Diagnostic",
	9:  "k_ESteamDatagramMsg_DataClientToRouter",
	10: "k_ESteamDatagramMsg_DataRouterToServer",
	11: "k_ESteamDatagramMsg_DataServerToRouter",
	12: "k_ESteamDatagramMsg_DataRouterToClient",
	13: "k_ESteamDatagramMsg_Stats",
	14: "k_ESteamDatagramMsg_ClientPingSampleRequest",
	15: "k_ESteamDatagramMsg_ClientPingSampleReply",
}
var ESteamDatagramMsgID_value = map[string]int32{
	"k_ESteamDatagramMsg_RouterPingRequest":            1,
	"k_ESteamDatagramMsg_RouterPingReply":              2,
	"k_ESteamDatagramMsg_GameserverPingRequest":        3,
	"k_ESteamDatagramMsg_GameserverPingReply":          4,
	"k_ESteamDatagramMsg_GameserverSessionRequest":     5,
	"k_ESteamDatagramMsg_GameserverSessionEstablished": 6,
	"k_ESteamDatagramMsg_NoSession":                    7,
	"k_ESteamDatagramMsg_Diagnostic":                   8,
	"k_ESteamDatagramMsg_DataClientToRouter":           9,
	"k_ESteamDatagramMsg_DataRouterToServer":           10,
	"k_ESteamDatagramMsg_DataServerToRouter":           11,
	"k_ESteamDatagramMsg_DataRouterToClient":           12,
	"k_ESteamDatagramMsg_Stats":                        13,
	"k_ESteamDatagramMsg_ClientPingSampleRequest":      14,
	"k_ESteamDatagramMsg_ClientPingSampleReply":        15,
}

func (x ESteamDatagramMsgID) Enum() *ESteamDatagramMsgID {
	p := new(ESteamDatagramMsgID)
	*p = x
	return p
}
func (x ESteamDatagramMsgID) String() string {
	return proto.EnumName(ESteamDatagramMsgID_name, int32(x))
}
func (x *ESteamDatagramMsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ESteamDatagramMsgID_value, data, "ESteamDatagramMsgID")
	if err != nil {
		return err
	}
	*x = ESteamDatagramMsgID(value)
	return nil
}

type CMsgSteamDatagramRouterPingReply struct {
	ClientTimestamp      *uint32  `protobuf:"fixed32,1,opt,name=client_timestamp" json:"client_timestamp,omitempty"`
	LatencyDatacenterIds []uint32 `protobuf:"fixed32,2,rep,packed,name=latency_datacenter_ids" json:"latency_datacenter_ids,omitempty"`
	LatencyPingMs        []uint32 `protobuf:"varint,3,rep,packed,name=latency_ping_ms" json:"latency_ping_ms,omitempty"`
	YourPublicIp         *uint32  `protobuf:"fixed32,4,opt,name=your_public_ip" json:"your_public_ip,omitempty"`
	ServerTime           *uint32  `protobuf:"fixed32,5,opt,name=server_time" json:"server_time,omitempty"`
	Challenge            *uint64  `protobuf:"fixed64,6,opt,name=challenge" json:"challenge,omitempty"`
	SecondsUntilShutdown *uint32  `protobuf:"varint,7,opt,name=seconds_until_shutdown" json:"seconds_until_shutdown,omitempty"`
	ClientCookie         *uint32  `protobuf:"fixed32,8,opt,name=client_cookie" json:"client_cookie,omitempty"`
	XXX_unrecognized     []byte   `json:"-"`
}

func (m *CMsgSteamDatagramRouterPingReply) Reset()         { *m = CMsgSteamDatagramRouterPingReply{} }
func (m *CMsgSteamDatagramRouterPingReply) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamDatagramRouterPingReply) ProtoMessage()    {}

func (m *CMsgSteamDatagramRouterPingReply) GetClientTimestamp() uint32 {
	if m != nil && m.ClientTimestamp != nil {
		return *m.ClientTimestamp
	}
	return 0
}

func (m *CMsgSteamDatagramRouterPingReply) GetLatencyDatacenterIds() []uint32 {
	if m != nil {
		return m.LatencyDatacenterIds
	}
	return nil
}

func (m *CMsgSteamDatagramRouterPingReply) GetLatencyPingMs() []uint32 {
	if m != nil {
		return m.LatencyPingMs
	}
	return nil
}

func (m *CMsgSteamDatagramRouterPingReply) GetYourPublicIp() uint32 {
	if m != nil && m.YourPublicIp != nil {
		return *m.YourPublicIp
	}
	return 0
}

func (m *CMsgSteamDatagramRouterPingReply) GetServerTime() uint32 {
	if m != nil && m.ServerTime != nil {
		return *m.ServerTime
	}
	return 0
}

func (m *CMsgSteamDatagramRouterPingReply) GetChallenge() uint64 {
	if m != nil && m.Challenge != nil {
		return *m.Challenge
	}
	return 0
}

func (m *CMsgSteamDatagramRouterPingReply) GetSecondsUntilShutdown() uint32 {
	if m != nil && m.SecondsUntilShutdown != nil {
		return *m.SecondsUntilShutdown
	}
	return 0
}

func (m *CMsgSteamDatagramRouterPingReply) GetClientCookie() uint32 {
	if m != nil && m.ClientCookie != nil {
		return *m.ClientCookie
	}
	return 0
}

type CMsgSteamDatagramGameserverPing struct {
	ClientSession           *uint32 `protobuf:"varint,1,opt,name=client_session" json:"client_session,omitempty"`
	ClientSteamId           *uint64 `protobuf:"fixed64,2,opt,name=client_steam_id" json:"client_steam_id,omitempty"`
	ClientTimestamp         *uint32 `protobuf:"fixed32,3,opt,name=client_timestamp" json:"client_timestamp,omitempty"`
	RouterTimestamp         *uint32 `protobuf:"fixed32,4,opt,name=router_timestamp" json:"router_timestamp,omitempty"`
	RouterGameserverLatency *uint32 `protobuf:"varint,5,opt,name=router_gameserver_latency" json:"router_gameserver_latency,omitempty"`
	SeqNumberRouter         *uint32 `protobuf:"varint,6,opt,name=seq_number_router" json:"seq_number_router,omitempty"`
	SeqNumberE2E            *uint32 `protobuf:"varint,7,opt,name=seq_number_e2e" json:"seq_number_e2e,omitempty"`
	XXX_unrecognized        []byte  `json:"-"`
}

func (m *CMsgSteamDatagramGameserverPing) Reset()         { *m = CMsgSteamDatagramGameserverPing{} }
func (m *CMsgSteamDatagramGameserverPing) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamDatagramGameserverPing) ProtoMessage()    {}

func (m *CMsgSteamDatagramGameserverPing) GetClientSession() uint32 {
	if m != nil && m.ClientSession != nil {
		return *m.ClientSession
	}
	return 0
}

func (m *CMsgSteamDatagramGameserverPing) GetClientSteamId() uint64 {
	if m != nil && m.ClientSteamId != nil {
		return *m.ClientSteamId
	}
	return 0
}

func (m *CMsgSteamDatagramGameserverPing) GetClientTimestamp() uint32 {
	if m != nil && m.ClientTimestamp != nil {
		return *m.ClientTimestamp
	}
	return 0
}

func (m *CMsgSteamDatagramGameserverPing) GetRouterTimestamp() uint32 {
	if m != nil && m.RouterTimestamp != nil {
		return *m.RouterTimestamp
	}
	return 0
}

func (m *CMsgSteamDatagramGameserverPing) GetRouterGameserverLatency() uint32 {
	if m != nil && m.RouterGameserverLatency != nil {
		return *m.RouterGameserverLatency
	}
	return 0
}

func (m *CMsgSteamDatagramGameserverPing) GetSeqNumberRouter() uint32 {
	if m != nil && m.SeqNumberRouter != nil {
		return *m.SeqNumberRouter
	}
	return 0
}

func (m *CMsgSteamDatagramGameserverPing) GetSeqNumberE2E() uint32 {
	if m != nil && m.SeqNumberE2E != nil {
		return *m.SeqNumberE2E
	}
	return 0
}

type CMsgSteamDatagramGameServerAuthTicket struct {
	TimeExpiry         *uint32 `protobuf:"fixed32,1,opt,name=time_expiry" json:"time_expiry,omitempty"`
	AuthorizedSteamId  *uint64 `protobuf:"fixed64,2,opt,name=authorized_steam_id" json:"authorized_steam_id,omitempty"`
	AuthorizedPublicIp *uint32 `protobuf:"fixed32,3,opt,name=authorized_public_ip" json:"authorized_public_ip,omitempty"`
	GameserverSteamId  *uint64 `protobuf:"fixed64,4,opt,name=gameserver_steam_id" json:"gameserver_steam_id,omitempty"`
	GameserverNetId    *uint64 `protobuf:"fixed64,5,opt,name=gameserver_net_id" json:"gameserver_net_id,omitempty"`
	Signature          []byte  `protobuf:"bytes,6,opt,name=signature" json:"signature,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *CMsgSteamDatagramGameServerAuthTicket) Reset()         { *m = CMsgSteamDatagramGameServerAuthTicket{} }
func (m *CMsgSteamDatagramGameServerAuthTicket) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamDatagramGameServerAuthTicket) ProtoMessage()    {}

func (m *CMsgSteamDatagramGameServerAuthTicket) GetTimeExpiry() uint32 {
	if m != nil && m.TimeExpiry != nil {
		return *m.TimeExpiry
	}
	return 0
}

func (m *CMsgSteamDatagramGameServerAuthTicket) GetAuthorizedSteamId() uint64 {
	if m != nil && m.AuthorizedSteamId != nil {
		return *m.AuthorizedSteamId
	}
	return 0
}

func (m *CMsgSteamDatagramGameServerAuthTicket) GetAuthorizedPublicIp() uint32 {
	if m != nil && m.AuthorizedPublicIp != nil {
		return *m.AuthorizedPublicIp
	}
	return 0
}

func (m *CMsgSteamDatagramGameServerAuthTicket) GetGameserverSteamId() uint64 {
	if m != nil && m.GameserverSteamId != nil {
		return *m.GameserverSteamId
	}
	return 0
}

func (m *CMsgSteamDatagramGameServerAuthTicket) GetGameserverNetId() uint64 {
	if m != nil && m.GameserverNetId != nil {
		return *m.GameserverNetId
	}
	return 0
}

func (m *CMsgSteamDatagramGameServerAuthTicket) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type CMsgSteamDatagramGameserverSessionRequest struct {
	Ticket           *CMsgSteamDatagramGameServerAuthTicket `protobuf:"bytes,1,opt,name=ticket" json:"ticket,omitempty"`
	ChallengeTime    *uint32                                `protobuf:"fixed32,3,opt,name=challenge_time" json:"challenge_time,omitempty"`
	Challenge        *uint64                                `protobuf:"fixed64,4,opt,name=challenge" json:"challenge,omitempty"`
	ClientCookie     *uint32                                `protobuf:"fixed32,5,opt,name=client_cookie" json:"client_cookie,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *CMsgSteamDatagramGameserverSessionRequest) Reset() {
	*m = CMsgSteamDatagramGameserverSessionRequest{}
}
func (m *CMsgSteamDatagramGameserverSessionRequest) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamDatagramGameserverSessionRequest) ProtoMessage()    {}

func (m *CMsgSteamDatagramGameserverSessionRequest) GetTicket() *CMsgSteamDatagramGameServerAuthTicket {
	if m != nil {
		return m.Ticket
	}
	return nil
}

func (m *CMsgSteamDatagramGameserverSessionRequest) GetChallengeTime() uint32 {
	if m != nil && m.ChallengeTime != nil {
		return *m.ChallengeTime
	}
	return 0
}

func (m *CMsgSteamDatagramGameserverSessionRequest) GetChallenge() uint64 {
	if m != nil && m.Challenge != nil {
		return *m.Challenge
	}
	return 0
}

func (m *CMsgSteamDatagramGameserverSessionRequest) GetClientCookie() uint32 {
	if m != nil && m.ClientCookie != nil {
		return *m.ClientCookie
	}
	return 0
}

type CMsgSteamDatagramGameserverSessionEstablished struct {
	ClientCookie         *uint32 `protobuf:"fixed32,1,opt,name=client_cookie" json:"client_cookie,omitempty"`
	LegacyClientCookie   *uint32 `protobuf:"fixed32,2,opt,name=legacy_client_cookie" json:"legacy_client_cookie,omitempty"`
	GameserverSteamId    *uint64 `protobuf:"fixed64,3,opt,name=gameserver_steam_id" json:"gameserver_steam_id,omitempty"`
	SecondsUntilShutdown *uint32 `protobuf:"varint,4,opt,name=seconds_until_shutdown" json:"seconds_until_shutdown,omitempty"`
	XXX_unrecognized     []byte  `json:"-"`
}

func (m *CMsgSteamDatagramGameserverSessionEstablished) Reset() {
	*m = CMsgSteamDatagramGameserverSessionEstablished{}
}
func (m *CMsgSteamDatagramGameserverSessionEstablished) String() string {
	return proto.CompactTextString(m)
}
func (*CMsgSteamDatagramGameserverSessionEstablished) ProtoMessage() {}

func (m *CMsgSteamDatagramGameserverSessionEstablished) GetClientCookie() uint32 {
	if m != nil && m.ClientCookie != nil {
		return *m.ClientCookie
	}
	return 0
}

func (m *CMsgSteamDatagramGameserverSessionEstablished) GetLegacyClientCookie() uint32 {
	if m != nil && m.LegacyClientCookie != nil {
		return *m.LegacyClientCookie
	}
	return 0
}

func (m *CMsgSteamDatagramGameserverSessionEstablished) GetGameserverSteamId() uint64 {
	if m != nil && m.GameserverSteamId != nil {
		return *m.GameserverSteamId
	}
	return 0
}

func (m *CMsgSteamDatagramGameserverSessionEstablished) GetSecondsUntilShutdown() uint32 {
	if m != nil && m.SecondsUntilShutdown != nil {
		return *m.SecondsUntilShutdown
	}
	return 0
}

type CMsgSteamDatagramNoSession struct {
	ClientCookie         *uint32 `protobuf:"fixed32,7,opt,name=client_cookie" json:"client_cookie,omitempty"`
	LegacyClientCookie_1 *uint32 `protobuf:"varint,1,opt,name=legacy_client_cookie_1" json:"legacy_client_cookie_1,omitempty"`
	LegacyClientCookie_2 *uint32 `protobuf:"varint,6,opt,name=legacy_client_cookie_2" json:"legacy_client_cookie_2,omitempty"`
	YourPublicIp         *uint32 `protobuf:"fixed32,2,opt,name=your_public_ip" json:"your_public_ip,omitempty"`
	ServerTime           *uint32 `protobuf:"fixed32,3,opt,name=server_time" json:"server_time,omitempty"`
	Challenge            *uint64 `protobuf:"fixed64,4,opt,name=challenge" json:"challenge,omitempty"`
	SecondsUntilShutdown *uint32 `protobuf:"varint,5,opt,name=seconds_until_shutdown" json:"seconds_until_shutdown,omitempty"`
	XXX_unrecognized     []byte  `json:"-"`
}

func (m *CMsgSteamDatagramNoSession) Reset()         { *m = CMsgSteamDatagramNoSession{} }
func (m *CMsgSteamDatagramNoSession) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamDatagramNoSession) ProtoMessage()    {}

func (m *CMsgSteamDatagramNoSession) GetClientCookie() uint32 {
	if m != nil && m.ClientCookie != nil {
		return *m.ClientCookie
	}
	return 0
}

func (m *CMsgSteamDatagramNoSession) GetLegacyClientCookie_1() uint32 {
	if m != nil && m.LegacyClientCookie_1 != nil {
		return *m.LegacyClientCookie_1
	}
	return 0
}

func (m *CMsgSteamDatagramNoSession) GetLegacyClientCookie_2() uint32 {
	if m != nil && m.LegacyClientCookie_2 != nil {
		return *m.LegacyClientCookie_2
	}
	return 0
}

func (m *CMsgSteamDatagramNoSession) GetYourPublicIp() uint32 {
	if m != nil && m.YourPublicIp != nil {
		return *m.YourPublicIp
	}
	return 0
}

func (m *CMsgSteamDatagramNoSession) GetServerTime() uint32 {
	if m != nil && m.ServerTime != nil {
		return *m.ServerTime
	}
	return 0
}

func (m *CMsgSteamDatagramNoSession) GetChallenge() uint64 {
	if m != nil && m.Challenge != nil {
		return *m.Challenge
	}
	return 0
}

func (m *CMsgSteamDatagramNoSession) GetSecondsUntilShutdown() uint32 {
	if m != nil && m.SecondsUntilShutdown != nil {
		return *m.SecondsUntilShutdown
	}
	return 0
}

type CMsgSteamDatagramDiagnostic struct {
	Severity         *uint32 `protobuf:"varint,1,opt,name=severity" json:"severity,omitempty"`
	Text             *string `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgSteamDatagramDiagnostic) Reset()         { *m = CMsgSteamDatagramDiagnostic{} }
func (m *CMsgSteamDatagramDiagnostic) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamDatagramDiagnostic) ProtoMessage()    {}

func (m *CMsgSteamDatagramDiagnostic) GetSeverity() uint32 {
	if m != nil && m.Severity != nil {
		return *m.Severity
	}
	return 0
}

func (m *CMsgSteamDatagramDiagnostic) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

type CMsgSteamDatagramDataCenterState struct {
	DataCenters      []*CMsgSteamDatagramDataCenterState_DataCenter `protobuf:"bytes,1,rep,name=data_centers" json:"data_centers,omitempty"`
	XXX_unrecognized []byte                                         `json:"-"`
}

func (m *CMsgSteamDatagramDataCenterState) Reset()         { *m = CMsgSteamDatagramDataCenterState{} }
func (m *CMsgSteamDatagramDataCenterState) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamDatagramDataCenterState) ProtoMessage()    {}

func (m *CMsgSteamDatagramDataCenterState) GetDataCenters() []*CMsgSteamDatagramDataCenterState_DataCenter {
	if m != nil {
		return m.DataCenters
	}
	return nil
}

type CMsgSteamDatagramDataCenterState_Server struct {
	Address          *string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	PingMs           *uint32 `protobuf:"varint,2,opt,name=ping_ms" json:"ping_ms,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgSteamDatagramDataCenterState_Server) Reset() {
	*m = CMsgSteamDatagramDataCenterState_Server{}
}
func (m *CMsgSteamDatagramDataCenterState_Server) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamDatagramDataCenterState_Server) ProtoMessage()    {}

func (m *CMsgSteamDatagramDataCenterState_Server) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *CMsgSteamDatagramDataCenterState_Server) GetPingMs() uint32 {
	if m != nil && m.PingMs != nil {
		return *m.PingMs
	}
	return 0
}

type CMsgSteamDatagramDataCenterState_DataCenter struct {
	Code             *string                                    `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	ServerSample     []*CMsgSteamDatagramDataCenterState_Server `protobuf:"bytes,2,rep,name=server_sample" json:"server_sample,omitempty"`
	XXX_unrecognized []byte                                     `json:"-"`
}

func (m *CMsgSteamDatagramDataCenterState_DataCenter) Reset() {
	*m = CMsgSteamDatagramDataCenterState_DataCenter{}
}
func (m *CMsgSteamDatagramDataCenterState_DataCenter) String() string {
	return proto.CompactTextString(m)
}
func (*CMsgSteamDatagramDataCenterState_DataCenter) ProtoMessage() {}

func (m *CMsgSteamDatagramDataCenterState_DataCenter) GetCode() string {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ""
}

func (m *CMsgSteamDatagramDataCenterState_DataCenter) GetServerSample() []*CMsgSteamDatagramDataCenterState_Server {
	if m != nil {
		return m.ServerSample
	}
	return nil
}

type CMsgSteamDatagramLinkInstantaneousStats struct {
	OutPacketsPerSecX10     *uint32 `protobuf:"varint,1,opt,name=out_packets_per_sec_x10" json:"out_packets_per_sec_x10,omitempty"`
	OutBytesPerSec          *uint32 `protobuf:"varint,2,opt,name=out_bytes_per_sec" json:"out_bytes_per_sec,omitempty"`
	InPacketsPerSecX10      *uint32 `protobuf:"varint,3,opt,name=in_packets_per_sec_x10" json:"in_packets_per_sec_x10,omitempty"`
	InBytesPerSec           *uint32 `protobuf:"varint,4,opt,name=in_bytes_per_sec" json:"in_bytes_per_sec,omitempty"`
	PingMs                  *uint32 `protobuf:"varint,5,opt,name=ping_ms" json:"ping_ms,omitempty"`
	PacketsDroppedPct       *uint32 `protobuf:"varint,6,opt,name=packets_dropped_pct" json:"packets_dropped_pct,omitempty"`
	PacketsWeirdSequencePct *uint32 `protobuf:"varint,7,opt,name=packets_weird_sequence_pct" json:"packets_weird_sequence_pct,omitempty"`
	XXX_unrecognized        []byte  `json:"-"`
}

func (m *CMsgSteamDatagramLinkInstantaneousStats) Reset() {
	*m = CMsgSteamDatagramLinkInstantaneousStats{}
}
func (m *CMsgSteamDatagramLinkInstantaneousStats) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamDatagramLinkInstantaneousStats) ProtoMessage()    {}

func (m *CMsgSteamDatagramLinkInstantaneousStats) GetOutPacketsPerSecX10() uint32 {
	if m != nil && m.OutPacketsPerSecX10 != nil {
		return *m.OutPacketsPerSecX10
	}
	return 0
}

func (m *CMsgSteamDatagramLinkInstantaneousStats) GetOutBytesPerSec() uint32 {
	if m != nil && m.OutBytesPerSec != nil {
		return *m.OutBytesPerSec
	}
	return 0
}

func (m *CMsgSteamDatagramLinkInstantaneousStats) GetInPacketsPerSecX10() uint32 {
	if m != nil && m.InPacketsPerSecX10 != nil {
		return *m.InPacketsPerSecX10
	}
	return 0
}

func (m *CMsgSteamDatagramLinkInstantaneousStats) GetInBytesPerSec() uint32 {
	if m != nil && m.InBytesPerSec != nil {
		return *m.InBytesPerSec
	}
	return 0
}

func (m *CMsgSteamDatagramLinkInstantaneousStats) GetPingMs() uint32 {
	if m != nil && m.PingMs != nil {
		return *m.PingMs
	}
	return 0
}

func (m *CMsgSteamDatagramLinkInstantaneousStats) GetPacketsDroppedPct() uint32 {
	if m != nil && m.PacketsDroppedPct != nil {
		return *m.PacketsDroppedPct
	}
	return 0
}

func (m *CMsgSteamDatagramLinkInstantaneousStats) GetPacketsWeirdSequencePct() uint32 {
	if m != nil && m.PacketsWeirdSequencePct != nil {
		return *m.PacketsWeirdSequencePct
	}
	return 0
}

type CMsgSteamDatagramLinkLifetimeStats struct {
	PacketsSent           *uint64 `protobuf:"varint,3,opt,name=packets_sent" json:"packets_sent,omitempty"`
	KbSent                *uint64 `protobuf:"varint,4,opt,name=kb_sent" json:"kb_sent,omitempty"`
	PacketsRecv           *uint64 `protobuf:"varint,5,opt,name=packets_recv" json:"packets_recv,omitempty"`
	KbRecv                *uint64 `protobuf:"varint,6,opt,name=kb_recv" json:"kb_recv,omitempty"`
	PacketsRecvSequenced  *uint64 `protobuf:"varint,7,opt,name=packets_recv_sequenced" json:"packets_recv_sequenced,omitempty"`
	PacketsRecvDropped    *uint64 `protobuf:"varint,8,opt,name=packets_recv_dropped" json:"packets_recv_dropped,omitempty"`
	PacketsRecvOutOfOrder *uint64 `protobuf:"varint,9,opt,name=packets_recv_out_of_order" json:"packets_recv_out_of_order,omitempty"`
	PacketsRecvDuplicate  *uint64 `protobuf:"varint,10,opt,name=packets_recv_duplicate" json:"packets_recv_duplicate,omitempty"`
	PacketsRecvLurch      *uint64 `protobuf:"varint,11,opt,name=packets_recv_lurch" json:"packets_recv_lurch,omitempty"`
	XXX_unrecognized      []byte  `json:"-"`
}

func (m *CMsgSteamDatagramLinkLifetimeStats) Reset()         { *m = CMsgSteamDatagramLinkLifetimeStats{} }
func (m *CMsgSteamDatagramLinkLifetimeStats) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamDatagramLinkLifetimeStats) ProtoMessage()    {}

func (m *CMsgSteamDatagramLinkLifetimeStats) GetPacketsSent() uint64 {
	if m != nil && m.PacketsSent != nil {
		return *m.PacketsSent
	}
	return 0
}

func (m *CMsgSteamDatagramLinkLifetimeStats) GetKbSent() uint64 {
	if m != nil && m.KbSent != nil {
		return *m.KbSent
	}
	return 0
}

func (m *CMsgSteamDatagramLinkLifetimeStats) GetPacketsRecv() uint64 {
	if m != nil && m.PacketsRecv != nil {
		return *m.PacketsRecv
	}
	return 0
}

func (m *CMsgSteamDatagramLinkLifetimeStats) GetKbRecv() uint64 {
	if m != nil && m.KbRecv != nil {
		return *m.KbRecv
	}
	return 0
}

func (m *CMsgSteamDatagramLinkLifetimeStats) GetPacketsRecvSequenced() uint64 {
	if m != nil && m.PacketsRecvSequenced != nil {
		return *m.PacketsRecvSequenced
	}
	return 0
}

func (m *CMsgSteamDatagramLinkLifetimeStats) GetPacketsRecvDropped() uint64 {
	if m != nil && m.PacketsRecvDropped != nil {
		return *m.PacketsRecvDropped
	}
	return 0
}

func (m *CMsgSteamDatagramLinkLifetimeStats) GetPacketsRecvOutOfOrder() uint64 {
	if m != nil && m.PacketsRecvOutOfOrder != nil {
		return *m.PacketsRecvOutOfOrder
	}
	return 0
}

func (m *CMsgSteamDatagramLinkLifetimeStats) GetPacketsRecvDuplicate() uint64 {
	if m != nil && m.PacketsRecvDuplicate != nil {
		return *m.PacketsRecvDuplicate
	}
	return 0
}

func (m *CMsgSteamDatagramLinkLifetimeStats) GetPacketsRecvLurch() uint64 {
	if m != nil && m.PacketsRecvLurch != nil {
		return *m.PacketsRecvLurch
	}
	return 0
}

type CMsgSteamDatagramConnectionQuality struct {
	Instantaneous    *CMsgSteamDatagramLinkInstantaneousStats `protobuf:"bytes,1,opt,name=instantaneous" json:"instantaneous,omitempty"`
	Lifetime         *CMsgSteamDatagramLinkLifetimeStats      `protobuf:"bytes,2,opt,name=lifetime" json:"lifetime,omitempty"`
	XXX_unrecognized []byte                                   `json:"-"`
}

func (m *CMsgSteamDatagramConnectionQuality) Reset()         { *m = CMsgSteamDatagramConnectionQuality{} }
func (m *CMsgSteamDatagramConnectionQuality) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamDatagramConnectionQuality) ProtoMessage()    {}

func (m *CMsgSteamDatagramConnectionQuality) GetInstantaneous() *CMsgSteamDatagramLinkInstantaneousStats {
	if m != nil {
		return m.Instantaneous
	}
	return nil
}

func (m *CMsgSteamDatagramConnectionQuality) GetLifetime() *CMsgSteamDatagramLinkLifetimeStats {
	if m != nil {
		return m.Lifetime
	}
	return nil
}

type CMsgSteamDatagramConnectionStatsClientToRouter struct {
	C2R              *CMsgSteamDatagramConnectionQuality `protobuf:"bytes,1,opt,name=c2r" json:"c2r,omitempty"`
	C2S              *CMsgSteamDatagramConnectionQuality `protobuf:"bytes,2,opt,name=c2s" json:"c2s,omitempty"`
	ClientTimestamp  *uint32                             `protobuf:"fixed32,3,opt,name=client_timestamp" json:"client_timestamp,omitempty"`
	ClientCookie     *uint32                             `protobuf:"fixed32,8,opt,name=client_cookie" json:"client_cookie,omitempty"`
	SeqNumC2R        *uint32                             `protobuf:"varint,9,opt,name=seq_num_c2r" json:"seq_num_c2r,omitempty"`
	SeqNumC2S        *uint32                             `protobuf:"varint,10,opt,name=seq_num_c2s" json:"seq_num_c2s,omitempty"`
	XXX_unrecognized []byte                              `json:"-"`
}

func (m *CMsgSteamDatagramConnectionStatsClientToRouter) Reset() {
	*m = CMsgSteamDatagramConnectionStatsClientToRouter{}
}
func (m *CMsgSteamDatagramConnectionStatsClientToRouter) String() string {
	return proto.CompactTextString(m)
}
func (*CMsgSteamDatagramConnectionStatsClientToRouter) ProtoMessage() {}

func (m *CMsgSteamDatagramConnectionStatsClientToRouter) GetC2R() *CMsgSteamDatagramConnectionQuality {
	if m != nil {
		return m.C2R
	}
	return nil
}

func (m *CMsgSteamDatagramConnectionStatsClientToRouter) GetC2S() *CMsgSteamDatagramConnectionQuality {
	if m != nil {
		return m.C2S
	}
	return nil
}

func (m *CMsgSteamDatagramConnectionStatsClientToRouter) GetClientTimestamp() uint32 {
	if m != nil && m.ClientTimestamp != nil {
		return *m.ClientTimestamp
	}
	return 0
}

func (m *CMsgSteamDatagramConnectionStatsClientToRouter) GetClientCookie() uint32 {
	if m != nil && m.ClientCookie != nil {
		return *m.ClientCookie
	}
	return 0
}

func (m *CMsgSteamDatagramConnectionStatsClientToRouter) GetSeqNumC2R() uint32 {
	if m != nil && m.SeqNumC2R != nil {
		return *m.SeqNumC2R
	}
	return 0
}

func (m *CMsgSteamDatagramConnectionStatsClientToRouter) GetSeqNumC2S() uint32 {
	if m != nil && m.SeqNumC2S != nil {
		return *m.SeqNumC2S
	}
	return 0
}

type CMsgSteamDatagramConnectionStatsRouterToClient struct {
	R2C                       *CMsgSteamDatagramConnectionQuality `protobuf:"bytes,1,opt,name=r2c" json:"r2c,omitempty"`
	S2C                       *CMsgSteamDatagramConnectionQuality `protobuf:"bytes,2,opt,name=s2c" json:"s2c,omitempty"`
	ClientTimestampFromRouter *uint32                             `protobuf:"fixed32,3,opt,name=client_timestamp_from_router" json:"client_timestamp_from_router,omitempty"`
	ClientTimestampFromServer *uint32                             `protobuf:"fixed32,4,opt,name=client_timestamp_from_server" json:"client_timestamp_from_server,omitempty"`
	RouterGameserverLatency   *uint32                             `protobuf:"varint,5,opt,name=router_gameserver_latency" json:"router_gameserver_latency,omitempty"`
	SecondsUntilShutdown      *uint32                             `protobuf:"varint,6,opt,name=seconds_until_shutdown" json:"seconds_until_shutdown,omitempty"`
	ClientCookie              *uint32                             `protobuf:"fixed32,7,opt,name=client_cookie" json:"client_cookie,omitempty"`
	SeqNumR2C                 *uint32                             `protobuf:"varint,8,opt,name=seq_num_r2c" json:"seq_num_r2c,omitempty"`
	SeqNumS2C                 *uint32                             `protobuf:"varint,9,opt,name=seq_num_s2c" json:"seq_num_s2c,omitempty"`
	XXX_unrecognized          []byte                              `json:"-"`
}

func (m *CMsgSteamDatagramConnectionStatsRouterToClient) Reset() {
	*m = CMsgSteamDatagramConnectionStatsRouterToClient{}
}
func (m *CMsgSteamDatagramConnectionStatsRouterToClient) String() string {
	return proto.CompactTextString(m)
}
func (*CMsgSteamDatagramConnectionStatsRouterToClient) ProtoMessage() {}

func (m *CMsgSteamDatagramConnectionStatsRouterToClient) GetR2C() *CMsgSteamDatagramConnectionQuality {
	if m != nil {
		return m.R2C
	}
	return nil
}

func (m *CMsgSteamDatagramConnectionStatsRouterToClient) GetS2C() *CMsgSteamDatagramConnectionQuality {
	if m != nil {
		return m.S2C
	}
	return nil
}

func (m *CMsgSteamDatagramConnectionStatsRouterToClient) GetClientTimestampFromRouter() uint32 {
	if m != nil && m.ClientTimestampFromRouter != nil {
		return *m.ClientTimestampFromRouter
	}
	return 0
}

func (m *CMsgSteamDatagramConnectionStatsRouterToClient) GetClientTimestampFromServer() uint32 {
	if m != nil && m.ClientTimestampFromServer != nil {
		return *m.ClientTimestampFromServer
	}
	return 0
}

func (m *CMsgSteamDatagramConnectionStatsRouterToClient) GetRouterGameserverLatency() uint32 {
	if m != nil && m.RouterGameserverLatency != nil {
		return *m.RouterGameserverLatency
	}
	return 0
}

func (m *CMsgSteamDatagramConnectionStatsRouterToClient) GetSecondsUntilShutdown() uint32 {
	if m != nil && m.SecondsUntilShutdown != nil {
		return *m.SecondsUntilShutdown
	}
	return 0
}

func (m *CMsgSteamDatagramConnectionStatsRouterToClient) GetClientCookie() uint32 {
	if m != nil && m.ClientCookie != nil {
		return *m.ClientCookie
	}
	return 0
}

func (m *CMsgSteamDatagramConnectionStatsRouterToClient) GetSeqNumR2C() uint32 {
	if m != nil && m.SeqNumR2C != nil {
		return *m.SeqNumR2C
	}
	return 0
}

func (m *CMsgSteamDatagramConnectionStatsRouterToClient) GetSeqNumS2C() uint32 {
	if m != nil && m.SeqNumS2C != nil {
		return *m.SeqNumS2C
	}
	return 0
}

type CMsgSteamDatagramConnectionStatsRouterToServer struct {
	C2S              *CMsgSteamDatagramConnectionQuality `protobuf:"bytes,2,opt,name=c2s" json:"c2s,omitempty"`
	ClientTimestamp  *uint32                             `protobuf:"fixed32,3,opt,name=client_timestamp" json:"client_timestamp,omitempty"`
	RouterTimestamp  *uint32                             `protobuf:"fixed32,4,opt,name=router_timestamp" json:"router_timestamp,omitempty"`
	SeqNumC2S        *uint32                             `protobuf:"varint,6,opt,name=seq_num_c2s" json:"seq_num_c2s,omitempty"`
	ClientSteamId    *uint64                             `protobuf:"fixed64,7,opt,name=client_steam_id" json:"client_steam_id,omitempty"`
	ClientSessionId  *uint32                             `protobuf:"varint,8,opt,name=client_session_id" json:"client_session_id,omitempty"`
	XXX_unrecognized []byte                              `json:"-"`
}

func (m *CMsgSteamDatagramConnectionStatsRouterToServer) Reset() {
	*m = CMsgSteamDatagramConnectionStatsRouterToServer{}
}
func (m *CMsgSteamDatagramConnectionStatsRouterToServer) String() string {
	return proto.CompactTextString(m)
}
func (*CMsgSteamDatagramConnectionStatsRouterToServer) ProtoMessage() {}

func (m *CMsgSteamDatagramConnectionStatsRouterToServer) GetC2S() *CMsgSteamDatagramConnectionQuality {
	if m != nil {
		return m.C2S
	}
	return nil
}

func (m *CMsgSteamDatagramConnectionStatsRouterToServer) GetClientTimestamp() uint32 {
	if m != nil && m.ClientTimestamp != nil {
		return *m.ClientTimestamp
	}
	return 0
}

func (m *CMsgSteamDatagramConnectionStatsRouterToServer) GetRouterTimestamp() uint32 {
	if m != nil && m.RouterTimestamp != nil {
		return *m.RouterTimestamp
	}
	return 0
}

func (m *CMsgSteamDatagramConnectionStatsRouterToServer) GetSeqNumC2S() uint32 {
	if m != nil && m.SeqNumC2S != nil {
		return *m.SeqNumC2S
	}
	return 0
}

func (m *CMsgSteamDatagramConnectionStatsRouterToServer) GetClientSteamId() uint64 {
	if m != nil && m.ClientSteamId != nil {
		return *m.ClientSteamId
	}
	return 0
}

func (m *CMsgSteamDatagramConnectionStatsRouterToServer) GetClientSessionId() uint32 {
	if m != nil && m.ClientSessionId != nil {
		return *m.ClientSessionId
	}
	return 0
}

type CMsgSteamDatagramClientPingSampleRequest struct {
	ClientCookie     *uint32 `protobuf:"fixed32,1,opt,name=client_cookie" json:"client_cookie,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgSteamDatagramClientPingSampleRequest) Reset() {
	*m = CMsgSteamDatagramClientPingSampleRequest{}
}
func (m *CMsgSteamDatagramClientPingSampleRequest) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamDatagramClientPingSampleRequest) ProtoMessage()    {}

func (m *CMsgSteamDatagramClientPingSampleRequest) GetClientCookie() uint32 {
	if m != nil && m.ClientCookie != nil {
		return *m.ClientCookie
	}
	return 0
}

type CMsgSteamDatagramClientPingSampleReply struct {
	ClientCookie     *uint32                                                  `protobuf:"fixed32,1,opt,name=client_cookie" json:"client_cookie,omitempty"`
	RoutingClusters  []*CMsgSteamDatagramClientPingSampleReply_RoutingCluster `protobuf:"bytes,2,rep,name=routing_clusters" json:"routing_clusters,omitempty"`
	XXX_unrecognized []byte                                                   `json:"-"`
}

func (m *CMsgSteamDatagramClientPingSampleReply) Reset() {
	*m = CMsgSteamDatagramClientPingSampleReply{}
}
func (m *CMsgSteamDatagramClientPingSampleReply) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamDatagramClientPingSampleReply) ProtoMessage()    {}

func (m *CMsgSteamDatagramClientPingSampleReply) GetClientCookie() uint32 {
	if m != nil && m.ClientCookie != nil {
		return *m.ClientCookie
	}
	return 0
}

func (m *CMsgSteamDatagramClientPingSampleReply) GetRoutingClusters() []*CMsgSteamDatagramClientPingSampleReply_RoutingCluster {
	if m != nil {
		return m.RoutingClusters
	}
	return nil
}

type CMsgSteamDatagramClientPingSampleReply_RoutingCluster struct {
	Id               *uint32 `protobuf:"fixed32,1,opt,name=id" json:"id,omitempty"`
	FrontPingMs      *uint32 `protobuf:"varint,2,opt,name=front_ping_ms" json:"front_ping_ms,omitempty"`
	E2EPingMs        *uint32 `protobuf:"varint,3,opt,name=e2e_ping_ms" json:"e2e_ping_ms,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgSteamDatagramClientPingSampleReply_RoutingCluster) Reset() {
	*m = CMsgSteamDatagramClientPingSampleReply_RoutingCluster{}
}
func (m *CMsgSteamDatagramClientPingSampleReply_RoutingCluster) String() string {
	return proto.CompactTextString(m)
}
func (*CMsgSteamDatagramClientPingSampleReply_RoutingCluster) ProtoMessage() {}

func (m *CMsgSteamDatagramClientPingSampleReply_RoutingCluster) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *CMsgSteamDatagramClientPingSampleReply_RoutingCluster) GetFrontPingMs() uint32 {
	if m != nil && m.FrontPingMs != nil {
		return *m.FrontPingMs
	}
	return 0
}

func (m *CMsgSteamDatagramClientPingSampleReply_RoutingCluster) GetE2EPingMs() uint32 {
	if m != nil && m.E2EPingMs != nil {
		return *m.E2EPingMs
	}
	return 0
}

func init() {
	proto.RegisterEnum("dota.ESteamDatagramMsgID", ESteamDatagramMsgID_name, ESteamDatagramMsgID_value)
}
