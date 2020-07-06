// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/quality_of_service.proto

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

type QualityOfService_Tier int32

const (
	// Default: no quality of service specified.
	QualityOfService_UNDEFINED QualityOfService_Tier = 0
	QualityOfService_HIGH      QualityOfService_Tier = 1
	QualityOfService_MEDIUM    QualityOfService_Tier = 2
	QualityOfService_LOW       QualityOfService_Tier = 3
)

var QualityOfService_Tier_name = map[int32]string{
	0: "UNDEFINED",
	1: "HIGH",
	2: "MEDIUM",
	3: "LOW",
}

var QualityOfService_Tier_value = map[string]int32{
	"UNDEFINED": 0,
	"HIGH":      1,
	"MEDIUM":    2,
	"LOW":       3,
}

func (x QualityOfService_Tier) String() string {
	return proto.EnumName(QualityOfService_Tier_name, int32(x))
}

func (QualityOfService_Tier) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_681e33256d2f8033, []int{1, 0}
}

type QualityOfServiceSpec struct {
	// Indicates how much queueing delay an execution can tolerate.
	QueueingBudgetMins   uint32   `protobuf:"varint,1,opt,name=queueing_budget_mins,json=queueingBudgetMins,proto3" json:"queueing_budget_mins,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QualityOfServiceSpec) Reset()         { *m = QualityOfServiceSpec{} }
func (m *QualityOfServiceSpec) String() string { return proto.CompactTextString(m) }
func (*QualityOfServiceSpec) ProtoMessage()    {}
func (*QualityOfServiceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_681e33256d2f8033, []int{0}
}

func (m *QualityOfServiceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QualityOfServiceSpec.Unmarshal(m, b)
}
func (m *QualityOfServiceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QualityOfServiceSpec.Marshal(b, m, deterministic)
}
func (m *QualityOfServiceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QualityOfServiceSpec.Merge(m, src)
}
func (m *QualityOfServiceSpec) XXX_Size() int {
	return xxx_messageInfo_QualityOfServiceSpec.Size(m)
}
func (m *QualityOfServiceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_QualityOfServiceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_QualityOfServiceSpec proto.InternalMessageInfo

func (m *QualityOfServiceSpec) GetQueueingBudgetMins() uint32 {
	if m != nil {
		return m.QueueingBudgetMins
	}
	return 0
}

type QualityOfService struct {
	// Types that are valid to be assigned to Designation:
	//	*QualityOfService_Tier_
	//	*QualityOfService_Spec
	Designation          isQualityOfService_Designation `protobuf_oneof:"designation"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *QualityOfService) Reset()         { *m = QualityOfService{} }
func (m *QualityOfService) String() string { return proto.CompactTextString(m) }
func (*QualityOfService) ProtoMessage()    {}
func (*QualityOfService) Descriptor() ([]byte, []int) {
	return fileDescriptor_681e33256d2f8033, []int{1}
}

func (m *QualityOfService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QualityOfService.Unmarshal(m, b)
}
func (m *QualityOfService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QualityOfService.Marshal(b, m, deterministic)
}
func (m *QualityOfService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QualityOfService.Merge(m, src)
}
func (m *QualityOfService) XXX_Size() int {
	return xxx_messageInfo_QualityOfService.Size(m)
}
func (m *QualityOfService) XXX_DiscardUnknown() {
	xxx_messageInfo_QualityOfService.DiscardUnknown(m)
}

var xxx_messageInfo_QualityOfService proto.InternalMessageInfo

type isQualityOfService_Designation interface {
	isQualityOfService_Designation()
}

type QualityOfService_Tier_ struct {
	Tier QualityOfService_Tier `protobuf:"varint,1,opt,name=tier,proto3,enum=flyteidl.admin.QualityOfService_Tier,oneof"`
}

type QualityOfService_Spec struct {
	Spec *QualityOfServiceSpec `protobuf:"bytes,2,opt,name=spec,proto3,oneof"`
}

func (*QualityOfService_Tier_) isQualityOfService_Designation() {}

func (*QualityOfService_Spec) isQualityOfService_Designation() {}

func (m *QualityOfService) GetDesignation() isQualityOfService_Designation {
	if m != nil {
		return m.Designation
	}
	return nil
}

func (m *QualityOfService) GetTier() QualityOfService_Tier {
	if x, ok := m.GetDesignation().(*QualityOfService_Tier_); ok {
		return x.Tier
	}
	return QualityOfService_UNDEFINED
}

func (m *QualityOfService) GetSpec() *QualityOfServiceSpec {
	if x, ok := m.GetDesignation().(*QualityOfService_Spec); ok {
		return x.Spec
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*QualityOfService) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*QualityOfService_Tier_)(nil),
		(*QualityOfService_Spec)(nil),
	}
}

func init() {
	proto.RegisterEnum("flyteidl.admin.QualityOfService_Tier", QualityOfService_Tier_name, QualityOfService_Tier_value)
	proto.RegisterType((*QualityOfServiceSpec)(nil), "flyteidl.admin.QualityOfServiceSpec")
	proto.RegisterType((*QualityOfService)(nil), "flyteidl.admin.QualityOfService")
}

func init() {
	proto.RegisterFile("flyteidl/admin/quality_of_service.proto", fileDescriptor_681e33256d2f8033)
}

var fileDescriptor_681e33256d2f8033 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x6b, 0xf2, 0x40,
	0x10, 0x86, 0x8d, 0x06, 0xbf, 0xaf, 0x23, 0x4a, 0x58, 0x3c, 0x78, 0x14, 0x69, 0xa9, 0x97, 0xee,
	0xb6, 0xda, 0x53, 0x7b, 0x13, 0x6d, 0x23, 0x54, 0xa5, 0x5a, 0x29, 0xf4, 0x12, 0x4c, 0x32, 0xd9,
	0x0e, 0xc4, 0xdd, 0x98, 0x6c, 0x0a, 0xfe, 0xcb, 0xfe, 0xa4, 0x92, 0x88, 0x05, 0xbd, 0xf4, 0xfc,
	0x3e, 0xcf, 0x3b, 0xc3, 0x0b, 0xd7, 0x51, 0xbc, 0x37, 0x48, 0x61, 0x2c, 0x36, 0xe1, 0x96, 0x94,
	0xd8, 0xe5, 0x9b, 0x98, 0xcc, 0xde, 0xd3, 0x91, 0x97, 0x61, 0xfa, 0x45, 0x01, 0xf2, 0x24, 0xd5,
	0x46, 0xb3, 0xd6, 0x11, 0xe4, 0x25, 0xd8, 0x73, 0xa1, 0xfd, 0x7a, 0x60, 0x17, 0xd1, 0xea, 0x40,
	0xae, 0x12, 0x0c, 0xd8, 0x2d, 0xb4, 0x77, 0x39, 0xe6, 0x48, 0x4a, 0x7a, 0x7e, 0x1e, 0x4a, 0x34,
	0xde, 0x96, 0x54, 0xd6, 0xb1, 0xba, 0x56, 0xbf, 0xb9, 0x64, 0xc7, 0x6c, 0x54, 0x46, 0x33, 0x52,
	0x59, 0xef, 0xdb, 0x02, 0xe7, 0xbc, 0x8a, 0x3d, 0x82, 0x6d, 0x08, 0xd3, 0x52, 0x6b, 0x0d, 0xae,
	0xf8, 0xe9, 0x75, 0x7e, 0xce, 0xf3, 0x37, 0xc2, 0xd4, 0xad, 0x2c, 0x4b, 0x89, 0x3d, 0x80, 0x9d,
	0x25, 0x18, 0x74, 0xaa, 0x5d, 0xab, 0xdf, 0x18, 0x5c, 0xfe, 0x25, 0x17, 0x7f, 0x17, 0x6e, 0xe1,
	0xf4, 0xee, 0xc1, 0x2e, 0xba, 0x58, 0x13, 0x2e, 0xd6, 0xf3, 0xf1, 0xe4, 0x69, 0x3a, 0x9f, 0x8c,
	0x9d, 0x0a, 0xfb, 0x0f, 0xb6, 0x3b, 0x7d, 0x76, 0x1d, 0x8b, 0x01, 0xd4, 0x67, 0x93, 0xf1, 0x74,
	0x3d, 0x73, 0xaa, 0xec, 0x1f, 0xd4, 0x5e, 0x16, 0xef, 0x4e, 0x6d, 0xd4, 0x84, 0x46, 0x88, 0x19,
	0x49, 0xb5, 0x31, 0xa4, 0xd5, 0x68, 0xf8, 0x71, 0x27, 0xc9, 0x7c, 0xe6, 0x3e, 0x0f, 0xf4, 0x56,
	0xc4, 0xfb, 0xc8, 0x88, 0xdf, 0x9d, 0x25, 0x2a, 0x91, 0xf8, 0x37, 0x52, 0x8b, 0xd3, 0xe9, 0xfd,
	0x7a, 0x39, 0xf4, 0xf0, 0x27, 0x00, 0x00, 0xff, 0xff, 0x57, 0x30, 0x24, 0xa3, 0x93, 0x01, 0x00,
	0x00,
}