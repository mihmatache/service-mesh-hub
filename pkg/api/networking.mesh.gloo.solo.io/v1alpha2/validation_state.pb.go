// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/networking/v1alpha2/validation_state.proto

package v1alpha2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// State of a Policy resource reflected in the status by Gloo Mesh while processing a resource.
type ApprovalState int32

const (
	// Resources are in a Pending state before they have been processed by Gloo Mesh.
	ApprovalState_PENDING ApprovalState = 0
	// Resources are in a Accepted state when they are valid and have been applied successfully to
	// the Gloo Mesh configuration.
	ApprovalState_ACCEPTED ApprovalState = 1
	// Resources are in an Invalid state when they contain incorrect configuration parameters,
	// such as missing required values or invalid resource references.
	// An invalid state can also result when a resource's configuration is valid
	// but conflicts with another resource which was accepted in an earlier point in time.
	ApprovalState_INVALID ApprovalState = 2
	// Resources are in a Failed state when they contain correct configuration parameters,
	// but the server encountered an error trying to synchronize the system to
	// the desired state.
	ApprovalState_FAILED ApprovalState = 3
)

var ApprovalState_name = map[int32]string{
	0: "PENDING",
	1: "ACCEPTED",
	2: "INVALID",
	3: "FAILED",
}

var ApprovalState_value = map[string]int32{
	"PENDING":  0,
	"ACCEPTED": 1,
	"INVALID":  2,
	"FAILED":   3,
}

func (x ApprovalState) String() string {
	return proto.EnumName(ApprovalState_name, int32(x))
}

func (ApprovalState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a12a8e2f43a030bc, []int{0}
}

// The approval status of a policy that has been applied to a discovery resource.
type ApprovalStatus struct {
	// AcceptanceOrder represents the order in which the Policy
	// was accepted and applied to a discovery resource. The first accepted policy
	// will have an acceptance_order of 0, the second 1, etc.
	// When conflicts are detected in the system,
	// the Policy with the lowest acceptance_order
	// will be chosen (and all other conflicting policies will be rejected).
	AcceptanceOrder uint32 `protobuf:"varint,1,opt,name=acceptance_order,json=acceptanceOrder,proto3" json:"acceptance_order,omitempty"`
	// The result of attempting to apply the policy to the discovery resource,
	// reported by the Policy controller (mesh-networking).
	State ApprovalState `protobuf:"varint,2,opt,name=state,proto3,enum=networking.mesh.gloo.solo.io.ApprovalState" json:"state,omitempty"`
	// Any errors observed which prevented the resource from being Accepted.
	Errors               []string `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApprovalStatus) Reset()         { *m = ApprovalStatus{} }
func (m *ApprovalStatus) String() string { return proto.CompactTextString(m) }
func (*ApprovalStatus) ProtoMessage()    {}
func (*ApprovalStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_a12a8e2f43a030bc, []int{0}
}
func (m *ApprovalStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApprovalStatus.Unmarshal(m, b)
}
func (m *ApprovalStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApprovalStatus.Marshal(b, m, deterministic)
}
func (m *ApprovalStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApprovalStatus.Merge(m, src)
}
func (m *ApprovalStatus) XXX_Size() int {
	return xxx_messageInfo_ApprovalStatus.Size(m)
}
func (m *ApprovalStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ApprovalStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ApprovalStatus proto.InternalMessageInfo

func (m *ApprovalStatus) GetAcceptanceOrder() uint32 {
	if m != nil {
		return m.AcceptanceOrder
	}
	return 0
}

func (m *ApprovalStatus) GetState() ApprovalState {
	if m != nil {
		return m.State
	}
	return ApprovalState_PENDING
}

func (m *ApprovalStatus) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

func init() {
	proto.RegisterEnum("networking.mesh.gloo.solo.io.ApprovalState", ApprovalState_name, ApprovalState_value)
	proto.RegisterType((*ApprovalStatus)(nil), "networking.mesh.gloo.solo.io.ApprovalStatus")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo-mesh/api/networking/v1alpha2/validation_state.proto", fileDescriptor_a12a8e2f43a030bc)
}

var fileDescriptor_a12a8e2f43a030bc = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xdd, 0x4e, 0x32, 0x31,
	0x14, 0xfc, 0x16, 0xf2, 0xa1, 0x56, 0xc1, 0x4d, 0x63, 0x0c, 0x41, 0x63, 0x88, 0x57, 0xa8, 0xa1,
	0x0d, 0xf8, 0x04, 0x2b, 0xbb, 0x9a, 0x55, 0x82, 0x04, 0x8d, 0x17, 0xde, 0x90, 0xb2, 0x34, 0xa5,
	0x61, 0xd9, 0xd3, 0xb4, 0x65, 0x7d, 0x13, 0x9f, 0xc1, 0xe7, 0xf2, 0x49, 0x4c, 0x17, 0x23, 0xfe,
	0x24, 0x5c, 0xf5, 0x74, 0x3a, 0x33, 0x9d, 0x9c, 0x41, 0x77, 0x42, 0xda, 0xd9, 0x72, 0x42, 0x12,
	0x58, 0x50, 0x03, 0x29, 0xb4, 0x25, 0x50, 0x91, 0x02, 0xb4, 0x17, 0xdc, 0xcc, 0x28, 0x53, 0x92,
	0x66, 0xdc, 0xbe, 0x80, 0x9e, 0xcb, 0x4c, 0xd0, 0xbc, 0xc3, 0x52, 0x35, 0x63, 0x5d, 0x9a, 0xb3,
	0x54, 0x4e, 0x99, 0x95, 0x90, 0x8d, 0x8d, 0x65, 0x96, 0x13, 0xa5, 0xc1, 0x02, 0x3e, 0x5e, 0x73,
	0x89, 0xd3, 0x13, 0xe7, 0x44, 0x9c, 0x2d, 0x91, 0xd0, 0x38, 0x32, 0xf3, 0xbc, 0x5b, 0x58, 0x26,
	0xa0, 0x39, 0xcd, 0x3b, 0xc5, 0xb9, 0x92, 0x36, 0x0e, 0x04, 0x08, 0x28, 0x46, 0xea, 0xa6, 0x15,
	0x7a, 0xfa, 0xea, 0xa1, 0x5a, 0xa0, 0x94, 0x86, 0x9c, 0xa5, 0x0f, 0x96, 0xd9, 0xa5, 0xc1, 0x67,
	0xc8, 0x67, 0x49, 0xc2, 0x95, 0x65, 0x59, 0xc2, 0xc7, 0xa0, 0xa7, 0x5c, 0xd7, 0xbd, 0xa6, 0xd7,
	0xaa, 0x8e, 0xf6, 0xd7, 0xf8, 0xbd, 0x83, 0x71, 0x80, 0xfe, 0x17, 0xe9, 0xea, 0xa5, 0xa6, 0xd7,
	0xaa, 0x75, 0x2f, 0xc8, 0xa6, 0x78, 0xe4, 0xfb, 0x3f, 0x7c, 0xb4, 0x52, 0xe2, 0x43, 0x54, 0xe1,
	0x5a, 0x83, 0x36, 0xf5, 0x72, 0xb3, 0xdc, 0xda, 0x19, 0x7d, 0xde, 0xce, 0x7b, 0xa8, 0xfa, 0x83,
	0x8f, 0x77, 0xd1, 0xd6, 0x30, 0x1a, 0x84, 0xf1, 0xe0, 0xc6, 0xff, 0x87, 0xf7, 0xd0, 0x76, 0xd0,
	0xeb, 0x45, 0xc3, 0xc7, 0x28, 0xf4, 0x3d, 0xf7, 0x14, 0x0f, 0x9e, 0x82, 0x7e, 0x1c, 0xfa, 0x25,
	0x8c, 0x50, 0xe5, 0x3a, 0x88, 0xfb, 0x51, 0xe8, 0x97, 0xaf, 0x86, 0x6f, 0xef, 0x27, 0xde, 0xf3,
	0xed, 0xc6, 0x06, 0xd4, 0x5c, 0xfc, 0x6a, 0xe1, 0x6f, 0xf4, 0xaf, 0x5e, 0x26, 0x95, 0x62, 0x6d,
	0x97, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x10, 0xf4, 0x3f, 0x5a, 0xd6, 0x01, 0x00, 0x00,
}

func (this *ApprovalStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ApprovalStatus)
	if !ok {
		that2, ok := that.(ApprovalStatus)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.AcceptanceOrder != that1.AcceptanceOrder {
		return false
	}
	if this.State != that1.State {
		return false
	}
	if len(this.Errors) != len(that1.Errors) {
		return false
	}
	for i := range this.Errors {
		if this.Errors[i] != that1.Errors[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}