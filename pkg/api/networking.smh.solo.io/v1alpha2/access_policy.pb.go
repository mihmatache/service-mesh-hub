// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/networking/v1alpha2/access_policy.proto

package v1alpha2

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha2/types"
	math "math"
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

//
//access control policies apply ALLOW policies to communication in a mesh
//access control policies specify the following:
//ALLOW those requests:
//- originating from from **source pods**
//- sent to **destination pods**
//- matching the indicated request criteria (allowed_paths, allowed_methods, allowed_ports)
//if no access control policies are present, all traffic in the mesh will be set to ALLOW
type AccessPolicySpec struct {
	//
	//requests originating from these pods will have the rule applied
	//leave empty to have all pods in the mesh apply these policies
	//
	//note that access control policies are mapped to source pods by their
	//service account. if other pods share the same service account,
	//this access control rule will apply to those pods as well.
	//
	//for fine-grained access control policies, ensure that your
	//service accounts properly reflect the desired
	//boundary for your access control policies
	SourceSelector []*IdentitySelector `protobuf:"bytes,2,rep,name=source_selector,json=sourceSelector,proto3" json:"source_selector,omitempty"`
	//
	//requests destined for these pods will have the rule applied
	//leave empty to apply to all destination pods in the mesh
	DestinationSelector []*ServiceSelector `protobuf:"bytes,3,rep,name=destination_selector,json=destinationSelector,proto3" json:"destination_selector,omitempty"`
	//
	//Optional. A list of HTTP paths or gRPC methods to allow.
	//gRPC methods must be presented as fully-qualified name in the form of
	//"/packageName.serviceName/methodName" and are case sensitive.
	//Exact match, prefix match, and suffix match are supported for paths.
	//For example, the path "/books/review" matches
	//"/books/review" (exact match), "*books/" (suffix match), or "/books*" (prefix match),
	//
	//If not specified, it allows to any path.
	AllowedPaths []string `protobuf:"bytes,4,rep,name=allowed_paths,json=allowedPaths,proto3" json:"allowed_paths,omitempty"`
	//
	//Optional. A list of HTTP methods to allow (e.g., "GET", "POST").
	//It is ignored in gRPC case because the value is always "POST".
	//If not specified, allows any method.
	AllowedMethods []types.HttpMethodValue `protobuf:"varint,5,rep,packed,name=allowed_methods,json=allowedMethods,proto3,enum=networking.smh.solo.io.HttpMethodValue" json:"allowed_methods,omitempty"`
	//
	//Optional. A list of ports which to allow
	//if not set any port is allowed
	AllowedPorts         []uint32 `protobuf:"varint,6,rep,packed,name=allowed_ports,json=allowedPorts,proto3" json:"allowed_ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessPolicySpec) Reset()         { *m = AccessPolicySpec{} }
func (m *AccessPolicySpec) String() string { return proto.CompactTextString(m) }
func (*AccessPolicySpec) ProtoMessage()    {}
func (*AccessPolicySpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_a607a654bf8f02aa, []int{0}
}
func (m *AccessPolicySpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessPolicySpec.Unmarshal(m, b)
}
func (m *AccessPolicySpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessPolicySpec.Marshal(b, m, deterministic)
}
func (m *AccessPolicySpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessPolicySpec.Merge(m, src)
}
func (m *AccessPolicySpec) XXX_Size() int {
	return xxx_messageInfo_AccessPolicySpec.Size(m)
}
func (m *AccessPolicySpec) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessPolicySpec.DiscardUnknown(m)
}

var xxx_messageInfo_AccessPolicySpec proto.InternalMessageInfo

func (m *AccessPolicySpec) GetSourceSelector() []*IdentitySelector {
	if m != nil {
		return m.SourceSelector
	}
	return nil
}

func (m *AccessPolicySpec) GetDestinationSelector() []*ServiceSelector {
	if m != nil {
		return m.DestinationSelector
	}
	return nil
}

func (m *AccessPolicySpec) GetAllowedPaths() []string {
	if m != nil {
		return m.AllowedPaths
	}
	return nil
}

func (m *AccessPolicySpec) GetAllowedMethods() []types.HttpMethodValue {
	if m != nil {
		return m.AllowedMethods
	}
	return nil
}

func (m *AccessPolicySpec) GetAllowedPorts() []uint32 {
	if m != nil {
		return m.AllowedPorts
	}
	return nil
}

type AccessPolicyStatus struct {
	// More detailed errors than the base status provided by `translation_status`.
	TranslatorErrors     []*AccessPolicyStatus_TranslatorError `protobuf:"bytes,2,rep,name=translator_errors,json=translatorErrors,proto3" json:"translator_errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *AccessPolicyStatus) Reset()         { *m = AccessPolicyStatus{} }
func (m *AccessPolicyStatus) String() string { return proto.CompactTextString(m) }
func (*AccessPolicyStatus) ProtoMessage()    {}
func (*AccessPolicyStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_a607a654bf8f02aa, []int{1}
}
func (m *AccessPolicyStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessPolicyStatus.Unmarshal(m, b)
}
func (m *AccessPolicyStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessPolicyStatus.Marshal(b, m, deterministic)
}
func (m *AccessPolicyStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessPolicyStatus.Merge(m, src)
}
func (m *AccessPolicyStatus) XXX_Size() int {
	return xxx_messageInfo_AccessPolicyStatus.Size(m)
}
func (m *AccessPolicyStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessPolicyStatus.DiscardUnknown(m)
}

var xxx_messageInfo_AccessPolicyStatus proto.InternalMessageInfo

func (m *AccessPolicyStatus) GetTranslatorErrors() []*AccessPolicyStatus_TranslatorError {
	if m != nil {
		return m.TranslatorErrors
	}
	return nil
}

// TODO use a shared Status message with TrafficPolicy once autopilot allows for it
type AccessPolicyStatus_TranslatorError struct {
	// ID representing a translator that translates TrafficPolicy to Mesh-specific config
	TranslatorId         string   `protobuf:"bytes,1,opt,name=translator_id,json=translatorId,proto3" json:"translator_id,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessPolicyStatus_TranslatorError) Reset()         { *m = AccessPolicyStatus_TranslatorError{} }
func (m *AccessPolicyStatus_TranslatorError) String() string { return proto.CompactTextString(m) }
func (*AccessPolicyStatus_TranslatorError) ProtoMessage()    {}
func (*AccessPolicyStatus_TranslatorError) Descriptor() ([]byte, []int) {
	return fileDescriptor_a607a654bf8f02aa, []int{1, 0}
}
func (m *AccessPolicyStatus_TranslatorError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessPolicyStatus_TranslatorError.Unmarshal(m, b)
}
func (m *AccessPolicyStatus_TranslatorError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessPolicyStatus_TranslatorError.Marshal(b, m, deterministic)
}
func (m *AccessPolicyStatus_TranslatorError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessPolicyStatus_TranslatorError.Merge(m, src)
}
func (m *AccessPolicyStatus_TranslatorError) XXX_Size() int {
	return xxx_messageInfo_AccessPolicyStatus_TranslatorError.Size(m)
}
func (m *AccessPolicyStatus_TranslatorError) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessPolicyStatus_TranslatorError.DiscardUnknown(m)
}

var xxx_messageInfo_AccessPolicyStatus_TranslatorError proto.InternalMessageInfo

func (m *AccessPolicyStatus_TranslatorError) GetTranslatorId() string {
	if m != nil {
		return m.TranslatorId
	}
	return ""
}

func (m *AccessPolicyStatus_TranslatorError) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*AccessPolicySpec)(nil), "networking.smh.solo.io.AccessPolicySpec")
	proto.RegisterType((*AccessPolicyStatus)(nil), "networking.smh.solo.io.AccessPolicyStatus")
	proto.RegisterType((*AccessPolicyStatus_TranslatorError)(nil), "networking.smh.solo.io.AccessPolicyStatus.TranslatorError")
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/networking/v1alpha2/access_policy.proto", fileDescriptor_a607a654bf8f02aa)
}

var fileDescriptor_a607a654bf8f02aa = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x8b, 0xd3, 0x50,
	0x10, 0xc7, 0x69, 0xa3, 0x0b, 0x3e, 0xdd, 0x76, 0x8d, 0x8b, 0x84, 0x1e, 0x24, 0xac, 0x07, 0x73,
	0x69, 0x82, 0xbb, 0x17, 0xf1, 0xa6, 0x20, 0xb8, 0xc8, 0x62, 0xb7, 0x15, 0x0f, 0xeb, 0x21, 0xbc,
	0x26, 0x43, 0xf2, 0xd8, 0x24, 0xf3, 0x78, 0x33, 0xd9, 0x65, 0xbf, 0x91, 0x9f, 0x49, 0xf0, 0xe2,
	0x27, 0x91, 0xbc, 0xa4, 0x4d, 0x5b, 0x2d, 0xf4, 0xf6, 0xf2, 0xe7, 0x3f, 0xbf, 0xf9, 0xcf, 0x30,
	0x11, 0x5f, 0x33, 0xc5, 0x79, 0xbd, 0x0c, 0x13, 0x2c, 0x23, 0xc2, 0x02, 0xa7, 0x0a, 0x23, 0x02,
	0x73, 0xa7, 0x12, 0x98, 0x96, 0x40, 0xf9, 0x34, 0xaf, 0x97, 0x91, 0xd4, 0x2a, 0xaa, 0x80, 0xef,
	0xd1, 0xdc, 0xaa, 0x2a, 0x8b, 0xee, 0xde, 0xca, 0x42, 0xe7, 0xf2, 0x3c, 0x92, 0x49, 0x02, 0x44,
	0xb1, 0xc6, 0x42, 0x25, 0x0f, 0xa1, 0x36, 0xc8, 0xe8, 0xbe, 0xec, 0x8d, 0x21, 0x95, 0x79, 0xd8,
	0x40, 0x43, 0x85, 0x93, 0x8b, 0x83, 0xa9, 0x39, 0xb3, 0x6e, 0x61, 0x93, 0x77, 0x07, 0x17, 0x11,
	0x14, 0x90, 0x30, 0x1a, 0xea, 0x2a, 0x4f, 0x33, 0xcc, 0xd0, 0x3e, 0xa3, 0xe6, 0xd5, 0xaa, 0x67,
	0xbf, 0x87, 0xe2, 0xe4, 0x83, 0x0d, 0x3d, 0xb3, 0x99, 0x17, 0x1a, 0x12, 0xf7, 0x5a, 0x8c, 0x09,
	0x6b, 0x93, 0x40, 0xbc, 0x82, 0x78, 0x43, 0xdf, 0x09, 0x9e, 0x9e, 0x07, 0xe1, 0xff, 0x67, 0x09,
	0x2f, 0x53, 0xa8, 0x58, 0xf1, 0xc3, 0xa2, 0xf3, 0xcf, 0x47, 0x2d, 0x60, 0xf5, 0xed, 0xde, 0x88,
	0xd3, 0x14, 0x88, 0x55, 0x25, 0x59, 0x61, 0xd5, 0x73, 0x1d, 0xcb, 0x7d, 0xb3, 0x8f, 0xbb, 0x68,
	0xa7, 0x5d, 0x63, 0x5f, 0x6c, 0x40, 0xd6, 0xec, 0xd7, 0xe2, 0x58, 0x16, 0x05, 0xde, 0x43, 0x1a,
	0x6b, 0xc9, 0x39, 0x79, 0x8f, 0x7c, 0x27, 0x78, 0x32, 0x7f, 0xd6, 0x89, 0xb3, 0x46, 0x73, 0x67,
	0x62, 0xbc, 0x32, 0x95, 0xc0, 0x39, 0xa6, 0xe4, 0x3d, 0xf6, 0x9d, 0x60, 0xb4, 0xbf, 0xf7, 0x67,
	0x66, 0x7d, 0x65, 0xad, 0xdf, 0x65, 0x51, 0xc3, 0x7c, 0xd4, 0xd5, 0xb7, 0x1a, 0x6d, 0xb5, 0x45,
	0xc3, 0xe4, 0x1d, 0xf9, 0x4e, 0x70, 0xdc, 0xb7, 0x6d, 0xb4, 0xb3, 0x5f, 0x03, 0xe1, 0x6e, 0xed,
	0x97, 0x25, 0xd7, 0xe4, 0x66, 0xe2, 0x39, 0x1b, 0x59, 0x51, 0x21, 0x19, 0x4d, 0x0c, 0xc6, 0xa0,
	0xa1, 0x6e, 0xc7, 0xef, 0xf7, 0xe5, 0xf9, 0x17, 0x13, 0x7e, 0x5b, 0x33, 0x3e, 0x35, 0x88, 0xf9,
	0x09, 0x6f, 0x0b, 0x34, 0xf9, 0x21, 0xc6, 0x3b, 0xa6, 0x26, 0xf7, 0x46, 0x6f, 0x95, 0x7a, 0x03,
	0x7f, 0xd0, 0xac, 0xab, 0x17, 0x2f, 0xd3, 0xc6, 0x64, 0x53, 0xc5, 0x25, 0x10, 0xc9, 0x0c, 0xbc,
	0x61, 0x6b, 0xb2, 0xe2, 0x55, 0xab, 0x7d, 0xbc, 0xfe, 0xf9, 0xe7, 0xd5, 0xe0, 0xe6, 0xcb, 0x21,
	0x3f, 0x8c, 0xbe, 0xcd, 0x76, 0x2e, 0x75, 0x73, 0xb6, 0xf5, 0xd5, 0x2e, 0x8f, 0xec, 0x59, 0x5e,
	0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x86, 0xe6, 0xb6, 0x29, 0x86, 0x03, 0x00, 0x00,
}

func (this *AccessPolicySpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccessPolicySpec)
	if !ok {
		that2, ok := that.(AccessPolicySpec)
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
	if len(this.SourceSelector) != len(that1.SourceSelector) {
		return false
	}
	for i := range this.SourceSelector {
		if !this.SourceSelector[i].Equal(that1.SourceSelector[i]) {
			return false
		}
	}
	if len(this.DestinationSelector) != len(that1.DestinationSelector) {
		return false
	}
	for i := range this.DestinationSelector {
		if !this.DestinationSelector[i].Equal(that1.DestinationSelector[i]) {
			return false
		}
	}
	if len(this.AllowedPaths) != len(that1.AllowedPaths) {
		return false
	}
	for i := range this.AllowedPaths {
		if this.AllowedPaths[i] != that1.AllowedPaths[i] {
			return false
		}
	}
	if len(this.AllowedMethods) != len(that1.AllowedMethods) {
		return false
	}
	for i := range this.AllowedMethods {
		if this.AllowedMethods[i] != that1.AllowedMethods[i] {
			return false
		}
	}
	if len(this.AllowedPorts) != len(that1.AllowedPorts) {
		return false
	}
	for i := range this.AllowedPorts {
		if this.AllowedPorts[i] != that1.AllowedPorts[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AccessPolicyStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccessPolicyStatus)
	if !ok {
		that2, ok := that.(AccessPolicyStatus)
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
	if len(this.TranslatorErrors) != len(that1.TranslatorErrors) {
		return false
	}
	for i := range this.TranslatorErrors {
		if !this.TranslatorErrors[i].Equal(that1.TranslatorErrors[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AccessPolicyStatus_TranslatorError) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccessPolicyStatus_TranslatorError)
	if !ok {
		that2, ok := that.(AccessPolicyStatus_TranslatorError)
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
	if this.TranslatorId != that1.TranslatorId {
		return false
	}
	if this.ErrorMessage != that1.ErrorMessage {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}