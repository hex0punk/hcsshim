// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/Microsoft/hcsshim/internal/ncproxyttrpc/networkconfigproxy.proto

package ncproxyttrpc

import (
	context "context"
	fmt "fmt"
	github_com_containerd_ttrpc "github.com/containerd/ttrpc"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RequestTypeInternal int32

const (
	RequestTypeInternal_Setup    RequestTypeInternal = 0
	RequestTypeInternal_Teardown RequestTypeInternal = 1
)

var RequestTypeInternal_name = map[int32]string{
	0: "Setup",
	1: "Teardown",
}

var RequestTypeInternal_value = map[string]int32{
	"Setup":    0,
	"Teardown": 1,
}

func (x RequestTypeInternal) String() string {
	return proto.EnumName(RequestTypeInternal_name, int32(x))
}

func (RequestTypeInternal) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_11f9efc6dfbf9b45, []int{0}
}

type RegisterComputeAgentRequest struct {
	AgentAddress         string   `protobuf:"bytes,1,opt,name=agent_address,json=agentAddress,proto3" json:"agent_address,omitempty"`
	ContainerID          string   `protobuf:"bytes,2,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterComputeAgentRequest) Reset()      { *m = RegisterComputeAgentRequest{} }
func (*RegisterComputeAgentRequest) ProtoMessage() {}
func (*RegisterComputeAgentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_11f9efc6dfbf9b45, []int{0}
}
func (m *RegisterComputeAgentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterComputeAgentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterComputeAgentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterComputeAgentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterComputeAgentRequest.Merge(m, src)
}
func (m *RegisterComputeAgentRequest) XXX_Size() int {
	return m.Size()
}
func (m *RegisterComputeAgentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterComputeAgentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterComputeAgentRequest proto.InternalMessageInfo

type RegisterComputeAgentResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterComputeAgentResponse) Reset()      { *m = RegisterComputeAgentResponse{} }
func (*RegisterComputeAgentResponse) ProtoMessage() {}
func (*RegisterComputeAgentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_11f9efc6dfbf9b45, []int{1}
}
func (m *RegisterComputeAgentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterComputeAgentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterComputeAgentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterComputeAgentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterComputeAgentResponse.Merge(m, src)
}
func (m *RegisterComputeAgentResponse) XXX_Size() int {
	return m.Size()
}
func (m *RegisterComputeAgentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterComputeAgentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterComputeAgentResponse proto.InternalMessageInfo

type ConfigureNetworkingInternalRequest struct {
	ContainerID          string              `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	RequestType          RequestTypeInternal `protobuf:"varint,2,opt,name=request_type,json=requestType,proto3,enum=RequestTypeInternal" json:"request_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ConfigureNetworkingInternalRequest) Reset()      { *m = ConfigureNetworkingInternalRequest{} }
func (*ConfigureNetworkingInternalRequest) ProtoMessage() {}
func (*ConfigureNetworkingInternalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_11f9efc6dfbf9b45, []int{2}
}
func (m *ConfigureNetworkingInternalRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigureNetworkingInternalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigureNetworkingInternalRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigureNetworkingInternalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigureNetworkingInternalRequest.Merge(m, src)
}
func (m *ConfigureNetworkingInternalRequest) XXX_Size() int {
	return m.Size()
}
func (m *ConfigureNetworkingInternalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigureNetworkingInternalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigureNetworkingInternalRequest proto.InternalMessageInfo

type ConfigureNetworkingInternalResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigureNetworkingInternalResponse) Reset()      { *m = ConfigureNetworkingInternalResponse{} }
func (*ConfigureNetworkingInternalResponse) ProtoMessage() {}
func (*ConfigureNetworkingInternalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_11f9efc6dfbf9b45, []int{3}
}
func (m *ConfigureNetworkingInternalResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigureNetworkingInternalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigureNetworkingInternalResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigureNetworkingInternalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigureNetworkingInternalResponse.Merge(m, src)
}
func (m *ConfigureNetworkingInternalResponse) XXX_Size() int {
	return m.Size()
}
func (m *ConfigureNetworkingInternalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigureNetworkingInternalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigureNetworkingInternalResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("RequestTypeInternal", RequestTypeInternal_name, RequestTypeInternal_value)
	proto.RegisterType((*RegisterComputeAgentRequest)(nil), "RegisterComputeAgentRequest")
	proto.RegisterType((*RegisterComputeAgentResponse)(nil), "RegisterComputeAgentResponse")
	proto.RegisterType((*ConfigureNetworkingInternalRequest)(nil), "ConfigureNetworkingInternalRequest")
	proto.RegisterType((*ConfigureNetworkingInternalResponse)(nil), "ConfigureNetworkingInternalResponse")
}

func init() {
	proto.RegisterFile("github.com/Microsoft/hcsshim/internal/ncproxyttrpc/networkconfigproxy.proto", fileDescriptor_11f9efc6dfbf9b45)
}

var fileDescriptor_11f9efc6dfbf9b45 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4f, 0x8f, 0xd2, 0x40,
	0x14, 0xef, 0x98, 0x68, 0x64, 0xa8, 0x4a, 0x06, 0x0e, 0x04, 0xb1, 0x9a, 0xa2, 0x89, 0xf1, 0x30,
	0x4d, 0xf0, 0xe0, 0xc1, 0x8b, 0x80, 0x17, 0x62, 0x34, 0xa6, 0x62, 0x62, 0xb8, 0x90, 0xd2, 0x3e,
	0xca, 0x44, 0x99, 0x19, 0x67, 0xa6, 0x22, 0x37, 0xbf, 0x82, 0xdf, 0x8a, 0xdb, 0xee, 0x71, 0x4f,
	0x9b, 0xa5, 0x9f, 0x64, 0xd3, 0x16, 0xd8, 0xcd, 0xa6, 0x81, 0xcd, 0xde, 0xe6, 0xfd, 0xf2, 0xe6,
	0xf7, 0xe7, 0xbd, 0x87, 0x3f, 0xc5, 0xcc, 0xcc, 0x93, 0x29, 0x0d, 0xc5, 0xc2, 0xfb, 0xcc, 0x42,
	0x25, 0xb4, 0x98, 0x19, 0x6f, 0x1e, 0x6a, 0x3d, 0x67, 0x0b, 0x8f, 0x71, 0x03, 0x8a, 0x07, 0xbf,
	0x3c, 0x1e, 0x4a, 0x25, 0xfe, 0xae, 0x8c, 0x51, 0x32, 0xf4, 0x38, 0x98, 0xa5, 0x50, 0x3f, 0x43,
	0xc1, 0x67, 0x2c, 0xce, 0x71, 0x2a, 0x95, 0x30, 0xa2, 0xd5, 0x88, 0x45, 0x2c, 0xf2, 0xa7, 0x97,
	0xbd, 0x0a, 0xd4, 0xfd, 0x83, 0x9f, 0xfa, 0x10, 0x33, 0x6d, 0x40, 0x0d, 0xc4, 0x42, 0x26, 0x06,
	0x7a, 0x31, 0x70, 0xe3, 0xc3, 0xef, 0x04, 0xb4, 0x21, 0x1d, 0xfc, 0x28, 0xc8, 0xea, 0x49, 0x10,
	0x45, 0x0a, 0xb4, 0x6e, 0xa2, 0x17, 0xe8, 0x75, 0xc5, 0xb7, 0x73, 0xb0, 0x57, 0x60, 0xa4, 0x8b,
	0xed, 0x50, 0x70, 0x13, 0x30, 0x0e, 0x6a, 0xc2, 0xa2, 0xe6, 0xbd, 0xac, 0xa7, 0xff, 0x24, 0x3d,
	0x7f, 0x5e, 0x1d, 0xec, 0xf0, 0xe1, 0x47, 0xbf, 0xba, 0x6f, 0x1a, 0x46, 0xae, 0x83, 0xdb, 0xe5,
	0xba, 0x5a, 0x0a, 0xae, 0xc1, 0xfd, 0x8f, 0xb0, 0x3b, 0xc8, 0x33, 0x24, 0x0a, 0xbe, 0x14, 0x99,
	0x18, 0x8f, 0x87, 0xdb, 0xcc, 0x3b, 0x7f, 0x37, 0xa5, 0xd1, 0x71, 0x69, 0xf2, 0x0e, 0xdb, 0xaa,
	0xf8, 0x3e, 0x31, 0x2b, 0x09, 0xb9, 0xdd, 0xc7, 0xdd, 0x06, 0xdd, 0x72, 0x8e, 0x56, 0x12, 0xf6,
	0x32, 0x55, 0x75, 0x05, 0xba, 0xaf, 0x70, 0xe7, 0xa0, 0xa5, 0xc2, 0xfa, 0x1b, 0x8a, 0xeb, 0x25,
	0x54, 0xa4, 0x82, 0xef, 0x7f, 0x03, 0x93, 0xc8, 0x9a, 0x45, 0x6c, 0xfc, 0x70, 0x04, 0x81, 0x8a,
	0xc4, 0x92, 0xd7, 0x50, 0xf7, 0x04, 0x61, 0xb2, 0xa5, 0x2b, 0xe8, 0xbf, 0x66, 0x5b, 0x23, 0xdf,
	0x71, 0xa3, 0x6c, 0x42, 0xa4, 0x4d, 0x0f, 0x2c, 0xac, 0xf5, 0x8c, 0x1e, 0x1c, 0xab, 0x45, 0xa6,
	0xb8, 0x5e, 0x12, 0x82, 0x74, 0xe8, 0xf1, 0x69, 0xb7, 0x5e, 0xd2, 0x5b, 0xe4, 0x77, 0xad, 0xfe,
	0x78, 0xbd, 0x71, 0xac, 0xb3, 0x8d, 0x63, 0xfd, 0x4b, 0x1d, 0xb4, 0x4e, 0x1d, 0x74, 0x9a, 0x3a,
	0xe8, 0x22, 0x75, 0xd0, 0xf8, 0xc3, 0x1d, 0x2e, 0xfa, 0xfd, 0xf5, 0xea, 0x87, 0x35, 0x7d, 0x90,
	0x5f, 0xee, 0xdb, 0xcb, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x7f, 0xad, 0x6c, 0x1e, 0x03, 0x00,
	0x00,
}

func (m *RegisterComputeAgentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterComputeAgentRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.AgentAddress) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNetworkconfigproxy(dAtA, i, uint64(len(m.AgentAddress)))
		i += copy(dAtA[i:], m.AgentAddress)
	}
	if len(m.ContainerID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNetworkconfigproxy(dAtA, i, uint64(len(m.ContainerID)))
		i += copy(dAtA[i:], m.ContainerID)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *RegisterComputeAgentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterComputeAgentResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ConfigureNetworkingInternalRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigureNetworkingInternalRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ContainerID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNetworkconfigproxy(dAtA, i, uint64(len(m.ContainerID)))
		i += copy(dAtA[i:], m.ContainerID)
	}
	if m.RequestType != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintNetworkconfigproxy(dAtA, i, uint64(m.RequestType))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ConfigureNetworkingInternalResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigureNetworkingInternalResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintNetworkconfigproxy(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RegisterComputeAgentRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AgentAddress)
	if l > 0 {
		n += 1 + l + sovNetworkconfigproxy(uint64(l))
	}
	l = len(m.ContainerID)
	if l > 0 {
		n += 1 + l + sovNetworkconfigproxy(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RegisterComputeAgentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ConfigureNetworkingInternalRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContainerID)
	if l > 0 {
		n += 1 + l + sovNetworkconfigproxy(uint64(l))
	}
	if m.RequestType != 0 {
		n += 1 + sovNetworkconfigproxy(uint64(m.RequestType))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ConfigureNetworkingInternalResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovNetworkconfigproxy(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNetworkconfigproxy(x uint64) (n int) {
	return sovNetworkconfigproxy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *RegisterComputeAgentRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RegisterComputeAgentRequest{`,
		`AgentAddress:` + fmt.Sprintf("%v", this.AgentAddress) + `,`,
		`ContainerID:` + fmt.Sprintf("%v", this.ContainerID) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RegisterComputeAgentResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RegisterComputeAgentResponse{`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ConfigureNetworkingInternalRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ConfigureNetworkingInternalRequest{`,
		`ContainerID:` + fmt.Sprintf("%v", this.ContainerID) + `,`,
		`RequestType:` + fmt.Sprintf("%v", this.RequestType) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ConfigureNetworkingInternalResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ConfigureNetworkingInternalResponse{`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringNetworkconfigproxy(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}

type NetworkConfigProxyService interface {
	RegisterComputeAgent(ctx context.Context, req *RegisterComputeAgentRequest) (*RegisterComputeAgentResponse, error)
	ConfigureNetworking(ctx context.Context, req *ConfigureNetworkingInternalRequest) (*ConfigureNetworkingInternalResponse, error)
}

func RegisterNetworkConfigProxyService(srv *github_com_containerd_ttrpc.Server, svc NetworkConfigProxyService) {
	srv.Register("NetworkConfigProxy", map[string]github_com_containerd_ttrpc.Method{
		"RegisterComputeAgent": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req RegisterComputeAgentRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.RegisterComputeAgent(ctx, &req)
		},
		"ConfigureNetworking": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req ConfigureNetworkingInternalRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.ConfigureNetworking(ctx, &req)
		},
	})
}

type networkConfigProxyClient struct {
	client *github_com_containerd_ttrpc.Client
}

func NewNetworkConfigProxyClient(client *github_com_containerd_ttrpc.Client) NetworkConfigProxyService {
	return &networkConfigProxyClient{
		client: client,
	}
}

func (c *networkConfigProxyClient) RegisterComputeAgent(ctx context.Context, req *RegisterComputeAgentRequest) (*RegisterComputeAgentResponse, error) {
	var resp RegisterComputeAgentResponse
	if err := c.client.Call(ctx, "NetworkConfigProxy", "RegisterComputeAgent", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *networkConfigProxyClient) ConfigureNetworking(ctx context.Context, req *ConfigureNetworkingInternalRequest) (*ConfigureNetworkingInternalResponse, error) {
	var resp ConfigureNetworkingInternalResponse
	if err := c.client.Call(ctx, "NetworkConfigProxy", "ConfigureNetworking", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
func (m *RegisterComputeAgentRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkconfigproxy
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisterComputeAgentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterComputeAgentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AgentAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkconfigproxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AgentAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContainerID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkconfigproxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContainerID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkconfigproxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RegisterComputeAgentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkconfigproxy
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisterComputeAgentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterComputeAgentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkconfigproxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConfigureNetworkingInternalRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkconfigproxy
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConfigureNetworkingInternalRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigureNetworkingInternalRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContainerID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkconfigproxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContainerID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestType", wireType)
			}
			m.RequestType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkconfigproxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestType |= RequestTypeInternal(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkconfigproxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConfigureNetworkingInternalResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkconfigproxy
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConfigureNetworkingInternalResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigureNetworkingInternalResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkconfigproxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipNetworkconfigproxy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNetworkconfigproxy
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNetworkconfigproxy
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNetworkconfigproxy
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthNetworkconfigproxy
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthNetworkconfigproxy
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNetworkconfigproxy
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipNetworkconfigproxy(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthNetworkconfigproxy
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthNetworkconfigproxy = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNetworkconfigproxy   = fmt.Errorf("proto: integer overflow")
)