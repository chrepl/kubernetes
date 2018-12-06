// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_lb_v1/messages/messages.proto

package messages

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LoadBalanceRequest struct {
	// Types that are valid to be assigned to LoadBalanceRequestType:
	//	*LoadBalanceRequest_InitialRequest
	//	*LoadBalanceRequest_ClientStats
	LoadBalanceRequestType isLoadBalanceRequest_LoadBalanceRequestType `protobuf_oneof:"load_balance_request_type"`
	XXX_NoUnkeyedLiteral   struct{}                                    `json:"-"`
	XXX_unrecognized       []byte                                      `json:"-"`
	XXX_sizecache          int32                                       `json:"-"`
}

func (m *LoadBalanceRequest) Reset()         { *m = LoadBalanceRequest{} }
func (m *LoadBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*LoadBalanceRequest) ProtoMessage()    {}
func (*LoadBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_b3d89fcb5aa158f8, []int{0}
}
func (m *LoadBalanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalanceRequest.Unmarshal(m, b)
}
func (m *LoadBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalanceRequest.Marshal(b, m, deterministic)
}
func (dst *LoadBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalanceRequest.Merge(dst, src)
}
func (m *LoadBalanceRequest) XXX_Size() int {
	return xxx_messageInfo_LoadBalanceRequest.Size(m)
}
func (m *LoadBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalanceRequest proto.InternalMessageInfo

type isLoadBalanceRequest_LoadBalanceRequestType interface {
	isLoadBalanceRequest_LoadBalanceRequestType()
}

type LoadBalanceRequest_InitialRequest struct {
	InitialRequest *InitialLoadBalanceRequest `protobuf:"bytes,1,opt,name=initial_request,json=initialRequest,proto3,oneof"`
}
type LoadBalanceRequest_ClientStats struct {
	ClientStats *ClientStats `protobuf:"bytes,2,opt,name=client_stats,json=clientStats,proto3,oneof"`
}

func (*LoadBalanceRequest_InitialRequest) isLoadBalanceRequest_LoadBalanceRequestType() {}
func (*LoadBalanceRequest_ClientStats) isLoadBalanceRequest_LoadBalanceRequestType()    {}

func (m *LoadBalanceRequest) GetLoadBalanceRequestType() isLoadBalanceRequest_LoadBalanceRequestType {
	if m != nil {
		return m.LoadBalanceRequestType
	}
	return nil
}

func (m *LoadBalanceRequest) GetInitialRequest() *InitialLoadBalanceRequest {
	if x, ok := m.GetLoadBalanceRequestType().(*LoadBalanceRequest_InitialRequest); ok {
		return x.InitialRequest
	}
	return nil
}

func (m *LoadBalanceRequest) GetClientStats() *ClientStats {
	if x, ok := m.GetLoadBalanceRequestType().(*LoadBalanceRequest_ClientStats); ok {
		return x.ClientStats
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LoadBalanceRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LoadBalanceRequest_OneofMarshaler, _LoadBalanceRequest_OneofUnmarshaler, _LoadBalanceRequest_OneofSizer, []interface{}{
		(*LoadBalanceRequest_InitialRequest)(nil),
		(*LoadBalanceRequest_ClientStats)(nil),
	}
}

func _LoadBalanceRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LoadBalanceRequest)
	// load_balance_request_type
	switch x := m.LoadBalanceRequestType.(type) {
	case *LoadBalanceRequest_InitialRequest:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.InitialRequest); err != nil {
			return err
		}
	case *LoadBalanceRequest_ClientStats:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClientStats); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LoadBalanceRequest.LoadBalanceRequestType has unexpected type %T", x)
	}
	return nil
}

func _LoadBalanceRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LoadBalanceRequest)
	switch tag {
	case 1: // load_balance_request_type.initial_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(InitialLoadBalanceRequest)
		err := b.DecodeMessage(msg)
		m.LoadBalanceRequestType = &LoadBalanceRequest_InitialRequest{msg}
		return true, err
	case 2: // load_balance_request_type.client_stats
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ClientStats)
		err := b.DecodeMessage(msg)
		m.LoadBalanceRequestType = &LoadBalanceRequest_ClientStats{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LoadBalanceRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LoadBalanceRequest)
	// load_balance_request_type
	switch x := m.LoadBalanceRequestType.(type) {
	case *LoadBalanceRequest_InitialRequest:
		s := proto.Size(x.InitialRequest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LoadBalanceRequest_ClientStats:
		s := proto.Size(x.ClientStats)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type InitialLoadBalanceRequest struct {
	// Name of load balanced service (IE, balancer.service.com). Its
	// length should be less than 256 bytes.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitialLoadBalanceRequest) Reset()         { *m = InitialLoadBalanceRequest{} }
func (m *InitialLoadBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*InitialLoadBalanceRequest) ProtoMessage()    {}
func (*InitialLoadBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_b3d89fcb5aa158f8, []int{1}
}
func (m *InitialLoadBalanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitialLoadBalanceRequest.Unmarshal(m, b)
}
func (m *InitialLoadBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitialLoadBalanceRequest.Marshal(b, m, deterministic)
}
func (dst *InitialLoadBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitialLoadBalanceRequest.Merge(dst, src)
}
func (m *InitialLoadBalanceRequest) XXX_Size() int {
	return xxx_messageInfo_InitialLoadBalanceRequest.Size(m)
}
func (m *InitialLoadBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitialLoadBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitialLoadBalanceRequest proto.InternalMessageInfo

func (m *InitialLoadBalanceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Contains the number of calls finished for a particular load balance token.
type ClientStatsPerToken struct {
	// See Server.load_balance_token.
	LoadBalanceToken string `protobuf:"bytes,1,opt,name=load_balance_token,json=loadBalanceToken,proto3" json:"load_balance_token,omitempty"`
	// The total number of RPCs that finished associated with the token.
	NumCalls             int64    `protobuf:"varint,2,opt,name=num_calls,json=numCalls,proto3" json:"num_calls,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientStatsPerToken) Reset()         { *m = ClientStatsPerToken{} }
func (m *ClientStatsPerToken) String() string { return proto.CompactTextString(m) }
func (*ClientStatsPerToken) ProtoMessage()    {}
func (*ClientStatsPerToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_b3d89fcb5aa158f8, []int{2}
}
func (m *ClientStatsPerToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStatsPerToken.Unmarshal(m, b)
}
func (m *ClientStatsPerToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStatsPerToken.Marshal(b, m, deterministic)
}
func (dst *ClientStatsPerToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStatsPerToken.Merge(dst, src)
}
func (m *ClientStatsPerToken) XXX_Size() int {
	return xxx_messageInfo_ClientStatsPerToken.Size(m)
}
func (m *ClientStatsPerToken) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStatsPerToken.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStatsPerToken proto.InternalMessageInfo

func (m *ClientStatsPerToken) GetLoadBalanceToken() string {
	if m != nil {
		return m.LoadBalanceToken
	}
	return ""
}

func (m *ClientStatsPerToken) GetNumCalls() int64 {
	if m != nil {
		return m.NumCalls
	}
	return 0
}

// Contains client level statistics that are useful to load balancing. Each
// count except the timestamp should be reset to zero after reporting the stats.
type ClientStats struct {
	// The timestamp of generating the report.
	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The total number of RPCs that started.
	NumCallsStarted int64 `protobuf:"varint,2,opt,name=num_calls_started,json=numCallsStarted,proto3" json:"num_calls_started,omitempty"`
	// The total number of RPCs that finished.
	NumCallsFinished int64 `protobuf:"varint,3,opt,name=num_calls_finished,json=numCallsFinished,proto3" json:"num_calls_finished,omitempty"`
	// The total number of RPCs that failed to reach a server except dropped RPCs.
	NumCallsFinishedWithClientFailedToSend int64 `protobuf:"varint,6,opt,name=num_calls_finished_with_client_failed_to_send,json=numCallsFinishedWithClientFailedToSend,proto3" json:"num_calls_finished_with_client_failed_to_send,omitempty"`
	// The total number of RPCs that finished and are known to have been received
	// by a server.
	NumCallsFinishedKnownReceived int64 `protobuf:"varint,7,opt,name=num_calls_finished_known_received,json=numCallsFinishedKnownReceived,proto3" json:"num_calls_finished_known_received,omitempty"`
	// The list of dropped calls.
	CallsFinishedWithDrop []*ClientStatsPerToken `protobuf:"bytes,8,rep,name=calls_finished_with_drop,json=callsFinishedWithDrop,proto3" json:"calls_finished_with_drop,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *ClientStats) Reset()         { *m = ClientStats{} }
func (m *ClientStats) String() string { return proto.CompactTextString(m) }
func (*ClientStats) ProtoMessage()    {}
func (*ClientStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_b3d89fcb5aa158f8, []int{3}
}
func (m *ClientStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStats.Unmarshal(m, b)
}
func (m *ClientStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStats.Marshal(b, m, deterministic)
}
func (dst *ClientStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStats.Merge(dst, src)
}
func (m *ClientStats) XXX_Size() int {
	return xxx_messageInfo_ClientStats.Size(m)
}
func (m *ClientStats) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStats.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStats proto.InternalMessageInfo

func (m *ClientStats) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ClientStats) GetNumCallsStarted() int64 {
	if m != nil {
		return m.NumCallsStarted
	}
	return 0
}

func (m *ClientStats) GetNumCallsFinished() int64 {
	if m != nil {
		return m.NumCallsFinished
	}
	return 0
}

func (m *ClientStats) GetNumCallsFinishedWithClientFailedToSend() int64 {
	if m != nil {
		return m.NumCallsFinishedWithClientFailedToSend
	}
	return 0
}

func (m *ClientStats) GetNumCallsFinishedKnownReceived() int64 {
	if m != nil {
		return m.NumCallsFinishedKnownReceived
	}
	return 0
}

func (m *ClientStats) GetCallsFinishedWithDrop() []*ClientStatsPerToken {
	if m != nil {
		return m.CallsFinishedWithDrop
	}
	return nil
}

type LoadBalanceResponse struct {
	// Types that are valid to be assigned to LoadBalanceResponseType:
	//	*LoadBalanceResponse_InitialResponse
	//	*LoadBalanceResponse_ServerList
	LoadBalanceResponseType isLoadBalanceResponse_LoadBalanceResponseType `protobuf_oneof:"load_balance_response_type"`
	XXX_NoUnkeyedLiteral    struct{}                                      `json:"-"`
	XXX_unrecognized        []byte                                        `json:"-"`
	XXX_sizecache           int32                                         `json:"-"`
}

func (m *LoadBalanceResponse) Reset()         { *m = LoadBalanceResponse{} }
func (m *LoadBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*LoadBalanceResponse) ProtoMessage()    {}
func (*LoadBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_b3d89fcb5aa158f8, []int{4}
}
func (m *LoadBalanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalanceResponse.Unmarshal(m, b)
}
func (m *LoadBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalanceResponse.Marshal(b, m, deterministic)
}
func (dst *LoadBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalanceResponse.Merge(dst, src)
}
func (m *LoadBalanceResponse) XXX_Size() int {
	return xxx_messageInfo_LoadBalanceResponse.Size(m)
}
func (m *LoadBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalanceResponse proto.InternalMessageInfo

type isLoadBalanceResponse_LoadBalanceResponseType interface {
	isLoadBalanceResponse_LoadBalanceResponseType()
}

type LoadBalanceResponse_InitialResponse struct {
	InitialResponse *InitialLoadBalanceResponse `protobuf:"bytes,1,opt,name=initial_response,json=initialResponse,proto3,oneof"`
}
type LoadBalanceResponse_ServerList struct {
	ServerList *ServerList `protobuf:"bytes,2,opt,name=server_list,json=serverList,proto3,oneof"`
}

func (*LoadBalanceResponse_InitialResponse) isLoadBalanceResponse_LoadBalanceResponseType() {}
func (*LoadBalanceResponse_ServerList) isLoadBalanceResponse_LoadBalanceResponseType()      {}

func (m *LoadBalanceResponse) GetLoadBalanceResponseType() isLoadBalanceResponse_LoadBalanceResponseType {
	if m != nil {
		return m.LoadBalanceResponseType
	}
	return nil
}

func (m *LoadBalanceResponse) GetInitialResponse() *InitialLoadBalanceResponse {
	if x, ok := m.GetLoadBalanceResponseType().(*LoadBalanceResponse_InitialResponse); ok {
		return x.InitialResponse
	}
	return nil
}

func (m *LoadBalanceResponse) GetServerList() *ServerList {
	if x, ok := m.GetLoadBalanceResponseType().(*LoadBalanceResponse_ServerList); ok {
		return x.ServerList
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LoadBalanceResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LoadBalanceResponse_OneofMarshaler, _LoadBalanceResponse_OneofUnmarshaler, _LoadBalanceResponse_OneofSizer, []interface{}{
		(*LoadBalanceResponse_InitialResponse)(nil),
		(*LoadBalanceResponse_ServerList)(nil),
	}
}

func _LoadBalanceResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LoadBalanceResponse)
	// load_balance_response_type
	switch x := m.LoadBalanceResponseType.(type) {
	case *LoadBalanceResponse_InitialResponse:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.InitialResponse); err != nil {
			return err
		}
	case *LoadBalanceResponse_ServerList:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ServerList); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LoadBalanceResponse.LoadBalanceResponseType has unexpected type %T", x)
	}
	return nil
}

func _LoadBalanceResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LoadBalanceResponse)
	switch tag {
	case 1: // load_balance_response_type.initial_response
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(InitialLoadBalanceResponse)
		err := b.DecodeMessage(msg)
		m.LoadBalanceResponseType = &LoadBalanceResponse_InitialResponse{msg}
		return true, err
	case 2: // load_balance_response_type.server_list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ServerList)
		err := b.DecodeMessage(msg)
		m.LoadBalanceResponseType = &LoadBalanceResponse_ServerList{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LoadBalanceResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LoadBalanceResponse)
	// load_balance_response_type
	switch x := m.LoadBalanceResponseType.(type) {
	case *LoadBalanceResponse_InitialResponse:
		s := proto.Size(x.InitialResponse)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LoadBalanceResponse_ServerList:
		s := proto.Size(x.ServerList)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type InitialLoadBalanceResponse struct {
	// This is an application layer redirect that indicates the client should use
	// the specified server for load balancing. When this field is non-empty in
	// the response, the client should open a separate connection to the
	// load_balancer_delegate and call the BalanceLoad method. Its length should
	// be less than 64 bytes.
	LoadBalancerDelegate string `protobuf:"bytes,1,opt,name=load_balancer_delegate,json=loadBalancerDelegate,proto3" json:"load_balancer_delegate,omitempty"`
	// This interval defines how often the client should send the client stats
	// to the load balancer. Stats should only be reported when the duration is
	// positive.
	ClientStatsReportInterval *duration.Duration `protobuf:"bytes,2,opt,name=client_stats_report_interval,json=clientStatsReportInterval,proto3" json:"client_stats_report_interval,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}           `json:"-"`
	XXX_unrecognized          []byte             `json:"-"`
	XXX_sizecache             int32              `json:"-"`
}

func (m *InitialLoadBalanceResponse) Reset()         { *m = InitialLoadBalanceResponse{} }
func (m *InitialLoadBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*InitialLoadBalanceResponse) ProtoMessage()    {}
func (*InitialLoadBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_b3d89fcb5aa158f8, []int{5}
}
func (m *InitialLoadBalanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitialLoadBalanceResponse.Unmarshal(m, b)
}
func (m *InitialLoadBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitialLoadBalanceResponse.Marshal(b, m, deterministic)
}
func (dst *InitialLoadBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitialLoadBalanceResponse.Merge(dst, src)
}
func (m *InitialLoadBalanceResponse) XXX_Size() int {
	return xxx_messageInfo_InitialLoadBalanceResponse.Size(m)
}
func (m *InitialLoadBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InitialLoadBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InitialLoadBalanceResponse proto.InternalMessageInfo

func (m *InitialLoadBalanceResponse) GetLoadBalancerDelegate() string {
	if m != nil {
		return m.LoadBalancerDelegate
	}
	return ""
}

func (m *InitialLoadBalanceResponse) GetClientStatsReportInterval() *duration.Duration {
	if m != nil {
		return m.ClientStatsReportInterval
	}
	return nil
}

type ServerList struct {
	// Contains a list of servers selected by the load balancer. The list will
	// be updated when server resolutions change or as needed to balance load
	// across more servers. The client should consume the server list in order
	// unless instructed otherwise via the client_config.
	Servers              []*Server `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ServerList) Reset()         { *m = ServerList{} }
func (m *ServerList) String() string { return proto.CompactTextString(m) }
func (*ServerList) ProtoMessage()    {}
func (*ServerList) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_b3d89fcb5aa158f8, []int{6}
}
func (m *ServerList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerList.Unmarshal(m, b)
}
func (m *ServerList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerList.Marshal(b, m, deterministic)
}
func (dst *ServerList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerList.Merge(dst, src)
}
func (m *ServerList) XXX_Size() int {
	return xxx_messageInfo_ServerList.Size(m)
}
func (m *ServerList) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerList.DiscardUnknown(m)
}

var xxx_messageInfo_ServerList proto.InternalMessageInfo

func (m *ServerList) GetServers() []*Server {
	if m != nil {
		return m.Servers
	}
	return nil
}

// Contains server information. When the drop field is not true, use the other
// fields.
type Server struct {
	// A resolved address for the server, serialized in network-byte-order. It may
	// either be an IPv4 or IPv6 address.
	IpAddress []byte `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	// A resolved port number for the server.
	Port int32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// An opaque but printable token given to the frontend for each pick. All
	// frontend requests for that pick must include the token in its initial
	// metadata. The token is used by the backend to verify the request and to
	// allow the backend to report load to the gRPC LB system. The token is also
	// used in client stats for reporting dropped calls.
	LoadBalanceToken string `protobuf:"bytes,3,opt,name=load_balance_token,json=loadBalanceToken,proto3" json:"load_balance_token,omitempty"`
	// Indicates whether this particular request should be dropped by the client.
	// If the request is dropped, there will be a corresponding entry in
	// ClientStats.calls_finished_with_drop.
	Drop                 bool     `protobuf:"varint,4,opt,name=drop,proto3" json:"drop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Server) Reset()         { *m = Server{} }
func (m *Server) String() string { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()    {}
func (*Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_b3d89fcb5aa158f8, []int{7}
}
func (m *Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Server.Unmarshal(m, b)
}
func (m *Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Server.Marshal(b, m, deterministic)
}
func (dst *Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server.Merge(dst, src)
}
func (m *Server) XXX_Size() int {
	return xxx_messageInfo_Server.Size(m)
}
func (m *Server) XXX_DiscardUnknown() {
	xxx_messageInfo_Server.DiscardUnknown(m)
}

var xxx_messageInfo_Server proto.InternalMessageInfo

func (m *Server) GetIpAddress() []byte {
	if m != nil {
		return m.IpAddress
	}
	return nil
}

func (m *Server) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Server) GetLoadBalanceToken() string {
	if m != nil {
		return m.LoadBalanceToken
	}
	return ""
}

func (m *Server) GetDrop() bool {
	if m != nil {
		return m.Drop
	}
	return false
}

func init() {
	proto.RegisterType((*LoadBalanceRequest)(nil), "grpc.lb.v1.LoadBalanceRequest")
	proto.RegisterType((*InitialLoadBalanceRequest)(nil), "grpc.lb.v1.InitialLoadBalanceRequest")
	proto.RegisterType((*ClientStatsPerToken)(nil), "grpc.lb.v1.ClientStatsPerToken")
	proto.RegisterType((*ClientStats)(nil), "grpc.lb.v1.ClientStats")
	proto.RegisterType((*LoadBalanceResponse)(nil), "grpc.lb.v1.LoadBalanceResponse")
	proto.RegisterType((*InitialLoadBalanceResponse)(nil), "grpc.lb.v1.InitialLoadBalanceResponse")
	proto.RegisterType((*ServerList)(nil), "grpc.lb.v1.ServerList")
	proto.RegisterType((*Server)(nil), "grpc.lb.v1.Server")
}

func init() {
	proto.RegisterFile("grpc_lb_v1/messages/messages.proto", fileDescriptor_messages_b3d89fcb5aa158f8)
}

var fileDescriptor_messages_b3d89fcb5aa158f8 = []byte{
	// 708 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x61, 0x4f, 0xf3, 0x36,
	0x10, 0x26, 0x6b, 0xe8, 0xdb, 0x5e, 0x5f, 0x8d, 0xce, 0x6c, 0x2c, 0x2d, 0x30, 0x58, 0xa4, 0x21,
	0x34, 0xb1, 0x54, 0xc0, 0x3e, 0x6c, 0xd2, 0x3e, 0x6c, 0x05, 0xa1, 0xc2, 0xf8, 0x80, 0x52, 0xa4,
	0x4d, 0x48, 0x93, 0xe7, 0x36, 0x26, 0x58, 0xb8, 0x76, 0x66, 0xbb, 0x45, 0xfb, 0xbc, 0xff, 0x33,
	0xed, 0x2f, 0x4c, 0xfb, 0x63, 0x53, 0xec, 0xa4, 0x0d, 0x2d, 0xd5, 0xfb, 0x25, 0x72, 0xee, 0x9e,
	0x7b, 0xee, 0xce, 0x77, 0x8f, 0x21, 0x4c, 0x55, 0x36, 0xc6, 0x7c, 0x84, 0x67, 0xa7, 0xbd, 0x09,
	0xd5, 0x9a, 0xa4, 0x54, 0xcf, 0x0f, 0x51, 0xa6, 0xa4, 0x91, 0x08, 0x72, 0x4c, 0xc4, 0x47, 0xd1,
	0xec, 0xb4, 0xfb, 0x45, 0x2a, 0x65, 0xca, 0x69, 0xcf, 0x7a, 0x46, 0xd3, 0xc7, 0x5e, 0x32, 0x55,
	0xc4, 0x30, 0x29, 0x1c, 0xb6, 0x7b, 0xb0, 0xec, 0x37, 0x6c, 0x42, 0xb5, 0x21, 0x93, 0xcc, 0x01,
	0xc2, 0x7f, 0x3d, 0x40, 0xb7, 0x92, 0x24, 0x7d, 0xc2, 0x89, 0x18, 0xd3, 0x98, 0xfe, 0x31, 0xa5,
	0xda, 0xa0, 0x3b, 0xd8, 0x62, 0x82, 0x19, 0x46, 0x38, 0x56, 0xce, 0x14, 0x78, 0x87, 0xde, 0x71,
	0xeb, 0xec, 0xab, 0x68, 0x91, 0x3d, 0xba, 0x76, 0x90, 0xd5, 0xf8, 0xc1, 0x46, 0xfc, 0x71, 0x11,
	0x5f, 0x32, 0xfe, 0x00, 0xef, 0xc7, 0x9c, 0x51, 0x61, 0xb0, 0x36, 0xc4, 0xe8, 0xe0, 0x23, 0x4b,
	0xf7, 0x79, 0x95, 0xee, 0xc2, 0xfa, 0x87, 0xb9, 0x7b, 0xb0, 0x11, 0xb7, 0xc6, 0x8b, 0xdf, 0xfe,
	0x2e, 0x74, 0xb8, 0x24, 0x09, 0x1e, 0xb9, 0x34, 0x65, 0x51, 0xd8, 0xfc, 0x99, 0xd1, 0xb0, 0x07,
	0x9d, 0xb5, 0x95, 0x20, 0x04, 0xbe, 0x20, 0x13, 0x6a, 0xcb, 0x6f, 0xc6, 0xf6, 0x1c, 0xfe, 0x0e,
	0xdb, 0x95, 0x5c, 0x77, 0x54, 0xdd, 0xcb, 0x67, 0x2a, 0xd0, 0x09, 0xa0, 0x57, 0x49, 0x4c, 0x6e,
	0x2d, 0x02, 0xdb, 0x7c, 0x41, 0xed, 0xd0, 0xbb, 0xd0, 0x14, 0xd3, 0x09, 0x1e, 0x13, 0xce, 0x5d,
	0x37, 0xb5, 0xb8, 0x21, 0xa6, 0x93, 0x8b, 0xfc, 0x3f, 0xfc, 0xa7, 0x06, 0xad, 0x4a, 0x0a, 0xf4,
	0x1d, 0x34, 0xe7, 0x37, 0x5f, 0xdc, 0x64, 0x37, 0x72, 0xb3, 0x89, 0xca, 0xd9, 0x44, 0xf7, 0x25,
	0x22, 0x5e, 0x80, 0xd1, 0xd7, 0xf0, 0xc9, 0x3c, 0x4d, 0x7e, 0x75, 0xca, 0xd0, 0xa4, 0x48, 0xb7,
	0x55, 0xa6, 0x1b, 0x3a, 0x73, 0xde, 0xc0, 0x02, 0xfb, 0xc8, 0x04, 0xd3, 0x4f, 0x34, 0x09, 0x6a,
	0x16, 0xdc, 0x2e, 0xc1, 0x57, 0x85, 0x1d, 0xfd, 0x06, 0xdf, 0xac, 0xa2, 0xf1, 0x0b, 0x33, 0x4f,
	0xb8, 0x98, 0xd4, 0x23, 0x61, 0x9c, 0x26, 0xd8, 0x48, 0xac, 0xa9, 0x48, 0x82, 0xba, 0x25, 0x3a,
	0x5a, 0x26, 0xfa, 0x85, 0x99, 0x27, 0xd7, 0xeb, 0x95, 0xc5, 0xdf, 0xcb, 0x21, 0x15, 0x09, 0x1a,
	0xc0, 0x97, 0x6f, 0xd0, 0x3f, 0x0b, 0xf9, 0x22, 0xb0, 0xa2, 0x63, 0xca, 0x66, 0x34, 0x09, 0xde,
	0x59, 0xca, 0xfd, 0x65, 0xca, 0x9f, 0x73, 0x54, 0x5c, 0x80, 0xd0, 0xaf, 0x10, 0xbc, 0x55, 0x64,
	0xa2, 0x64, 0x16, 0x34, 0x0e, 0x6b, 0xc7, 0xad, 0xb3, 0x83, 0x35, 0x6b, 0x54, 0x8e, 0x36, 0xfe,
	0x6c, 0xbc, 0x5c, 0xf1, 0xa5, 0x92, 0xd9, 0x8d, 0xdf, 0xf0, 0xdb, 0x9b, 0x37, 0x7e, 0x63, 0xb3,
	0x5d, 0x0f, 0xff, 0xf3, 0x60, 0xfb, 0xd5, 0xfe, 0xe8, 0x4c, 0x0a, 0x4d, 0xd1, 0x10, 0xda, 0x0b,
	0x29, 0x38, 0x5b, 0x31, 0xc1, 0xa3, 0x0f, 0x69, 0xc1, 0xa1, 0x07, 0x1b, 0xf1, 0xd6, 0x5c, 0x0c,
	0x05, 0xe9, 0xf7, 0xd0, 0xd2, 0x54, 0xcd, 0xa8, 0xc2, 0x9c, 0x69, 0x53, 0x88, 0x61, 0xa7, 0xca,
	0x37, 0xb4, 0xee, 0x5b, 0x66, 0xc5, 0x04, 0x7a, 0xfe, 0xd7, 0xdf, 0x83, 0xee, 0x92, 0x14, 0x1c,
	0xa7, 0xd3, 0xc2, 0xdf, 0x1e, 0x74, 0xd7, 0x97, 0x82, 0xbe, 0x85, 0x9d, 0x6a, 0xb0, 0xc2, 0x09,
	0xe5, 0x34, 0x25, 0xa6, 0xd4, 0xc7, 0xa7, 0x95, 0x35, 0x57, 0x97, 0x85, 0x0f, 0x3d, 0xc0, 0x5e,
	0x55, 0xbb, 0x58, 0xd1, 0x4c, 0x2a, 0x83, 0x99, 0x30, 0x54, 0xcd, 0x08, 0x2f, 0xca, 0xef, 0xac,
	0x2c, 0xf4, 0x65, 0xf1, 0x18, 0xc5, 0x9d, 0x8a, 0x96, 0x63, 0x1b, 0x7c, 0x5d, 0xc4, 0x86, 0x3f,
	0x02, 0x2c, 0x5a, 0x45, 0x27, 0xf0, 0xce, 0xb5, 0xaa, 0x03, 0xcf, 0x4e, 0x16, 0xad, 0xde, 0x49,
	0x5c, 0x42, 0x6e, 0xfc, 0x46, 0xad, 0xed, 0x87, 0x7f, 0x79, 0x50, 0x77, 0x1e, 0xb4, 0x0f, 0xc0,
	0x32, 0x4c, 0x92, 0x44, 0x51, 0xad, 0x6d, 0x4b, 0xef, 0xe3, 0x26, 0xcb, 0x7e, 0x72, 0x86, 0xfc,
	0x2d, 0xc8, 0x73, 0xdb, 0x7a, 0x37, 0x63, 0x7b, 0x5e, 0x23, 0xfa, 0xda, 0x1a, 0xd1, 0x23, 0xf0,
	0xed, 0xda, 0xf9, 0x87, 0xde, 0x71, 0x23, 0xb6, 0x67, 0xb7, 0x3e, 0xfd, 0xf3, 0x87, 0xd3, 0xa2,
	0xfd, 0x54, 0x72, 0x22, 0xd2, 0x48, 0xaa, 0xb4, 0x97, 0xd7, 0x6e, 0x3f, 0x7c, 0xd4, 0x7b, 0xe3,
	0x65, 0x1f, 0xd5, 0xed, 0x55, 0x9d, 0xff, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x88, 0xe6, 0xf4,
	0xf7, 0x05, 0x00, 0x00,
}
