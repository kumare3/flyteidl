// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/workflow_attributes.proto

package admin

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type WorkflowAttributes struct {
	// Unique project id for which this set of attributes will be applied.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Unique domain id for which this set of attributes will be applied.
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// Workflow name for which this set of attributes will be applied.
	Workflow             string              `protobuf:"bytes,3,opt,name=workflow,proto3" json:"workflow,omitempty"`
	MatchingAttributes   *MatchingAttributes `protobuf:"bytes,4,opt,name=matching_attributes,json=matchingAttributes,proto3" json:"matching_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *WorkflowAttributes) Reset()         { *m = WorkflowAttributes{} }
func (m *WorkflowAttributes) String() string { return proto.CompactTextString(m) }
func (*WorkflowAttributes) ProtoMessage()    {}
func (*WorkflowAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ba8a51ab86bc38c, []int{0}
}

func (m *WorkflowAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowAttributes.Unmarshal(m, b)
}
func (m *WorkflowAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowAttributes.Marshal(b, m, deterministic)
}
func (m *WorkflowAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowAttributes.Merge(m, src)
}
func (m *WorkflowAttributes) XXX_Size() int {
	return xxx_messageInfo_WorkflowAttributes.Size(m)
}
func (m *WorkflowAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowAttributes proto.InternalMessageInfo

func (m *WorkflowAttributes) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *WorkflowAttributes) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *WorkflowAttributes) GetWorkflow() string {
	if m != nil {
		return m.Workflow
	}
	return ""
}

func (m *WorkflowAttributes) GetMatchingAttributes() *MatchingAttributes {
	if m != nil {
		return m.MatchingAttributes
	}
	return nil
}

// Sets custom attributes for a project, domain and workflow combination.
type WorkflowAttributesUpdateRequest struct {
	Attributes           *WorkflowAttributes `protobuf:"bytes,1,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *WorkflowAttributesUpdateRequest) Reset()         { *m = WorkflowAttributesUpdateRequest{} }
func (m *WorkflowAttributesUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*WorkflowAttributesUpdateRequest) ProtoMessage()    {}
func (*WorkflowAttributesUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ba8a51ab86bc38c, []int{1}
}

func (m *WorkflowAttributesUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowAttributesUpdateRequest.Unmarshal(m, b)
}
func (m *WorkflowAttributesUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowAttributesUpdateRequest.Marshal(b, m, deterministic)
}
func (m *WorkflowAttributesUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowAttributesUpdateRequest.Merge(m, src)
}
func (m *WorkflowAttributesUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_WorkflowAttributesUpdateRequest.Size(m)
}
func (m *WorkflowAttributesUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowAttributesUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowAttributesUpdateRequest proto.InternalMessageInfo

func (m *WorkflowAttributesUpdateRequest) GetAttributes() *WorkflowAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// Purposefully empty, may be populated in the future.
type WorkflowAttributesUpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowAttributesUpdateResponse) Reset()         { *m = WorkflowAttributesUpdateResponse{} }
func (m *WorkflowAttributesUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*WorkflowAttributesUpdateResponse) ProtoMessage()    {}
func (*WorkflowAttributesUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ba8a51ab86bc38c, []int{2}
}

func (m *WorkflowAttributesUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowAttributesUpdateResponse.Unmarshal(m, b)
}
func (m *WorkflowAttributesUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowAttributesUpdateResponse.Marshal(b, m, deterministic)
}
func (m *WorkflowAttributesUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowAttributesUpdateResponse.Merge(m, src)
}
func (m *WorkflowAttributesUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_WorkflowAttributesUpdateResponse.Size(m)
}
func (m *WorkflowAttributesUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowAttributesUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowAttributesUpdateResponse proto.InternalMessageInfo

type WorkflowAttributesGetRequest struct {
	// Unique project id which this set of attributes references.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Unique domain id which this set of attributes references.
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// Workflow name which this set of attributes references.
	Workflow             string   `protobuf:"bytes,3,opt,name=workflow,proto3" json:"workflow,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowAttributesGetRequest) Reset()         { *m = WorkflowAttributesGetRequest{} }
func (m *WorkflowAttributesGetRequest) String() string { return proto.CompactTextString(m) }
func (*WorkflowAttributesGetRequest) ProtoMessage()    {}
func (*WorkflowAttributesGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ba8a51ab86bc38c, []int{3}
}

func (m *WorkflowAttributesGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowAttributesGetRequest.Unmarshal(m, b)
}
func (m *WorkflowAttributesGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowAttributesGetRequest.Marshal(b, m, deterministic)
}
func (m *WorkflowAttributesGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowAttributesGetRequest.Merge(m, src)
}
func (m *WorkflowAttributesGetRequest) XXX_Size() int {
	return xxx_messageInfo_WorkflowAttributesGetRequest.Size(m)
}
func (m *WorkflowAttributesGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowAttributesGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowAttributesGetRequest proto.InternalMessageInfo

func (m *WorkflowAttributesGetRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *WorkflowAttributesGetRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *WorkflowAttributesGetRequest) GetWorkflow() string {
	if m != nil {
		return m.Workflow
	}
	return ""
}

type WorkflowAttributesGetResponse struct {
	Attributes           *WorkflowAttributes `protobuf:"bytes,1,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *WorkflowAttributesGetResponse) Reset()         { *m = WorkflowAttributesGetResponse{} }
func (m *WorkflowAttributesGetResponse) String() string { return proto.CompactTextString(m) }
func (*WorkflowAttributesGetResponse) ProtoMessage()    {}
func (*WorkflowAttributesGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ba8a51ab86bc38c, []int{4}
}

func (m *WorkflowAttributesGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowAttributesGetResponse.Unmarshal(m, b)
}
func (m *WorkflowAttributesGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowAttributesGetResponse.Marshal(b, m, deterministic)
}
func (m *WorkflowAttributesGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowAttributesGetResponse.Merge(m, src)
}
func (m *WorkflowAttributesGetResponse) XXX_Size() int {
	return xxx_messageInfo_WorkflowAttributesGetResponse.Size(m)
}
func (m *WorkflowAttributesGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowAttributesGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowAttributesGetResponse proto.InternalMessageInfo

func (m *WorkflowAttributesGetResponse) GetAttributes() *WorkflowAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type WorkflowAttributesDeleteRequest struct {
	// Unique project id which this set of attributes references.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Unique domain id which this set of attributes references.
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// Workflow name which this set of attributes references.
	Workflow             string   `protobuf:"bytes,3,opt,name=workflow,proto3" json:"workflow,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowAttributesDeleteRequest) Reset()         { *m = WorkflowAttributesDeleteRequest{} }
func (m *WorkflowAttributesDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*WorkflowAttributesDeleteRequest) ProtoMessage()    {}
func (*WorkflowAttributesDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ba8a51ab86bc38c, []int{5}
}

func (m *WorkflowAttributesDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowAttributesDeleteRequest.Unmarshal(m, b)
}
func (m *WorkflowAttributesDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowAttributesDeleteRequest.Marshal(b, m, deterministic)
}
func (m *WorkflowAttributesDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowAttributesDeleteRequest.Merge(m, src)
}
func (m *WorkflowAttributesDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_WorkflowAttributesDeleteRequest.Size(m)
}
func (m *WorkflowAttributesDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowAttributesDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowAttributesDeleteRequest proto.InternalMessageInfo

func (m *WorkflowAttributesDeleteRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *WorkflowAttributesDeleteRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *WorkflowAttributesDeleteRequest) GetWorkflow() string {
	if m != nil {
		return m.Workflow
	}
	return ""
}

// Purposefully empty, may be populated in the future.
type WorkflowAttributesDeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowAttributesDeleteResponse) Reset()         { *m = WorkflowAttributesDeleteResponse{} }
func (m *WorkflowAttributesDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*WorkflowAttributesDeleteResponse) ProtoMessage()    {}
func (*WorkflowAttributesDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ba8a51ab86bc38c, []int{6}
}

func (m *WorkflowAttributesDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowAttributesDeleteResponse.Unmarshal(m, b)
}
func (m *WorkflowAttributesDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowAttributesDeleteResponse.Marshal(b, m, deterministic)
}
func (m *WorkflowAttributesDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowAttributesDeleteResponse.Merge(m, src)
}
func (m *WorkflowAttributesDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_WorkflowAttributesDeleteResponse.Size(m)
}
func (m *WorkflowAttributesDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowAttributesDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowAttributesDeleteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*WorkflowAttributes)(nil), "flyteidl.admin.WorkflowAttributes")
	proto.RegisterType((*WorkflowAttributesUpdateRequest)(nil), "flyteidl.admin.WorkflowAttributesUpdateRequest")
	proto.RegisterType((*WorkflowAttributesUpdateResponse)(nil), "flyteidl.admin.WorkflowAttributesUpdateResponse")
	proto.RegisterType((*WorkflowAttributesGetRequest)(nil), "flyteidl.admin.WorkflowAttributesGetRequest")
	proto.RegisterType((*WorkflowAttributesGetResponse)(nil), "flyteidl.admin.WorkflowAttributesGetResponse")
	proto.RegisterType((*WorkflowAttributesDeleteRequest)(nil), "flyteidl.admin.WorkflowAttributesDeleteRequest")
	proto.RegisterType((*WorkflowAttributesDeleteResponse)(nil), "flyteidl.admin.WorkflowAttributesDeleteResponse")
}

func init() {
	proto.RegisterFile("flyteidl/admin/workflow_attributes.proto", fileDescriptor_8ba8a51ab86bc38c)
}

var fileDescriptor_8ba8a51ab86bc38c = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xc9, 0xff, 0x2f, 0x55, 0x47, 0xf0, 0xb0, 0x82, 0x84, 0xa2, 0x58, 0xf6, 0x62, 0x2f,
	0x66, 0xd1, 0x3e, 0x81, 0x45, 0xf0, 0xe4, 0x25, 0x22, 0x82, 0x97, 0xb2, 0x49, 0xa6, 0xe9, 0xea,
	0x26, 0x1b, 0x37, 0x13, 0x4a, 0x9f, 0xcc, 0xd7, 0x13, 0xd2, 0x24, 0x4d, 0x4d, 0x7a, 0xcb, 0x71,
	0x76, 0xbf, 0x99, 0xef, 0x37, 0x1f, 0x03, 0xd3, 0xa5, 0xde, 0x10, 0xaa, 0x48, 0x0b, 0x19, 0x25,
	0x2a, 0x15, 0x6b, 0x63, 0xbf, 0x96, 0xda, 0xac, 0x17, 0x92, 0xc8, 0xaa, 0xa0, 0x20, 0xcc, 0xbd,
	0xcc, 0x1a, 0x32, 0xec, 0xbc, 0x56, 0x7a, 0xa5, 0x72, 0x7c, 0xfb, 0xa7, 0x33, 0x91, 0x14, 0xae,
	0x64, 0xa0, 0x71, 0x61, 0x31, 0x37, 0x85, 0x0d, 0x71, 0xdb, 0xc8, 0x7f, 0x1c, 0x60, 0xef, 0xd5,
	0xd8, 0xc7, 0x66, 0x2a, 0x73, 0xe1, 0x38, 0xb3, 0xe6, 0x13, 0x43, 0x72, 0x9d, 0x89, 0x33, 0x3d,
	0xf5, 0xeb, 0x92, 0x5d, 0xc2, 0x28, 0x32, 0x89, 0x54, 0xa9, 0xfb, 0xaf, 0xfc, 0xa8, 0x2a, 0x36,
	0x86, 0x93, 0x1a, 0xcf, 0xfd, 0x5f, 0xfe, 0x34, 0x35, 0x7b, 0x85, 0x8b, 0x12, 0x40, 0xa5, 0x71,
	0x0b, 0xdd, 0x3d, 0x9a, 0x38, 0xd3, 0xb3, 0x07, 0xee, 0xed, 0xb3, 0x7b, 0x2f, 0x95, 0x74, 0x87,
	0xe3, 0xb3, 0xa4, 0xf3, 0xc6, 0x11, 0x6e, 0xba, 0xe0, 0x6f, 0x59, 0x24, 0x09, 0x7d, 0xfc, 0x2e,
	0x30, 0x27, 0x36, 0x07, 0x68, 0xd9, 0x39, 0xfd, 0x76, 0xdd, 0x21, 0x7e, 0xab, 0x8b, 0x73, 0x98,
	0x1c, 0xb6, 0xc9, 0x33, 0x93, 0xe6, 0xc8, 0x35, 0x5c, 0x75, 0x35, 0xcf, 0x48, 0x35, 0xc7, 0xa0,
	0x69, 0xf2, 0x10, 0xae, 0x0f, 0xb8, 0x6d, 0x71, 0x06, 0x59, 0xdb, 0xf4, 0xa5, 0xfb, 0x84, 0x1a,
	0x77, 0xe9, 0x0e, 0xbb, 0x55, 0x6f, 0xce, 0xb5, 0xe1, 0x76, 0xb1, 0xf9, 0xec, 0xe3, 0x3e, 0x56,
	0xb4, 0x2a, 0x02, 0x2f, 0x34, 0x89, 0xd0, 0x9b, 0x25, 0x89, 0xe6, 0xce, 0x63, 0x4c, 0x45, 0x16,
	0xdc, 0xc5, 0x46, 0xec, 0x9f, 0x7e, 0x30, 0x2a, 0x0f, 0x7d, 0xf6, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0x9e, 0xef, 0x81, 0x5e, 0x4d, 0x03, 0x00, 0x00,
}