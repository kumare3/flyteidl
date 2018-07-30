// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/task.proto

package admin // import "github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TaskCreateRequest struct {
	Id                   *Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version              string      `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Spec                 *TaskSpec   `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TaskCreateRequest) Reset()         { *m = TaskCreateRequest{} }
func (m *TaskCreateRequest) String() string { return proto.CompactTextString(m) }
func (*TaskCreateRequest) ProtoMessage()    {}
func (*TaskCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_ec178c90a216dba4, []int{0}
}
func (m *TaskCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskCreateRequest.Unmarshal(m, b)
}
func (m *TaskCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskCreateRequest.Marshal(b, m, deterministic)
}
func (dst *TaskCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskCreateRequest.Merge(dst, src)
}
func (m *TaskCreateRequest) XXX_Size() int {
	return xxx_messageInfo_TaskCreateRequest.Size(m)
}
func (m *TaskCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskCreateRequest proto.InternalMessageInfo

func (m *TaskCreateRequest) GetId() *Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *TaskCreateRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *TaskCreateRequest) GetSpec() *TaskSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type TaskListRequest struct {
	Project              string   `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Domain               string   `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Limit                uint32   `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               uint32   `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	Filters              string   `protobuf:"bytes,6,opt,name=filters,proto3" json:"filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskListRequest) Reset()         { *m = TaskListRequest{} }
func (m *TaskListRequest) String() string { return proto.CompactTextString(m) }
func (*TaskListRequest) ProtoMessage()    {}
func (*TaskListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_ec178c90a216dba4, []int{1}
}
func (m *TaskListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskListRequest.Unmarshal(m, b)
}
func (m *TaskListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskListRequest.Marshal(b, m, deterministic)
}
func (dst *TaskListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskListRequest.Merge(dst, src)
}
func (m *TaskListRequest) XXX_Size() int {
	return xxx_messageInfo_TaskListRequest.Size(m)
}
func (m *TaskListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskListRequest proto.InternalMessageInfo

func (m *TaskListRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *TaskListRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *TaskListRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskListRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *TaskListRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *TaskListRequest) GetFilters() string {
	if m != nil {
		return m.Filters
	}
	return ""
}

type TaskCreateResponse struct {
	Urn                  string   `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskCreateResponse) Reset()         { *m = TaskCreateResponse{} }
func (m *TaskCreateResponse) String() string { return proto.CompactTextString(m) }
func (*TaskCreateResponse) ProtoMessage()    {}
func (*TaskCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_ec178c90a216dba4, []int{2}
}
func (m *TaskCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskCreateResponse.Unmarshal(m, b)
}
func (m *TaskCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskCreateResponse.Marshal(b, m, deterministic)
}
func (dst *TaskCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskCreateResponse.Merge(dst, src)
}
func (m *TaskCreateResponse) XXX_Size() int {
	return xxx_messageInfo_TaskCreateResponse.Size(m)
}
func (m *TaskCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskCreateResponse proto.InternalMessageInfo

func (m *TaskCreateResponse) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

type Task struct {
	Id                   *Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version              string      `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Urn                  string      `protobuf:"bytes,3,opt,name=urn,proto3" json:"urn,omitempty"`
	Spec                 *TaskSpec   `protobuf:"bytes,6,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_ec178c90a216dba4, []int{3}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (dst *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(dst, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetId() *Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Task) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Task) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *Task) GetSpec() *TaskSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type TaskList struct {
	Tasks                []*Task  `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskList) Reset()         { *m = TaskList{} }
func (m *TaskList) String() string { return proto.CompactTextString(m) }
func (*TaskList) ProtoMessage()    {}
func (*TaskList) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_ec178c90a216dba4, []int{4}
}
func (m *TaskList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskList.Unmarshal(m, b)
}
func (m *TaskList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskList.Marshal(b, m, deterministic)
}
func (dst *TaskList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskList.Merge(dst, src)
}
func (m *TaskList) XXX_Size() int {
	return xxx_messageInfo_TaskList.Size(m)
}
func (m *TaskList) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskList.DiscardUnknown(m)
}

var xxx_messageInfo_TaskList proto.InternalMessageInfo

func (m *TaskList) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type TaskSpec struct {
	Task                 *core.TaskTemplate `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TaskSpec) Reset()         { *m = TaskSpec{} }
func (m *TaskSpec) String() string { return proto.CompactTextString(m) }
func (*TaskSpec) ProtoMessage()    {}
func (*TaskSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_ec178c90a216dba4, []int{5}
}
func (m *TaskSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskSpec.Unmarshal(m, b)
}
func (m *TaskSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskSpec.Marshal(b, m, deterministic)
}
func (dst *TaskSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskSpec.Merge(dst, src)
}
func (m *TaskSpec) XXX_Size() int {
	return xxx_messageInfo_TaskSpec.Size(m)
}
func (m *TaskSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskSpec.DiscardUnknown(m)
}

var xxx_messageInfo_TaskSpec proto.InternalMessageInfo

func (m *TaskSpec) GetTask() *core.TaskTemplate {
	if m != nil {
		return m.Task
	}
	return nil
}

func init() {
	proto.RegisterType((*TaskCreateRequest)(nil), "admin.TaskCreateRequest")
	proto.RegisterType((*TaskListRequest)(nil), "admin.TaskListRequest")
	proto.RegisterType((*TaskCreateResponse)(nil), "admin.TaskCreateResponse")
	proto.RegisterType((*Task)(nil), "admin.Task")
	proto.RegisterType((*TaskList)(nil), "admin.TaskList")
	proto.RegisterType((*TaskSpec)(nil), "admin.TaskSpec")
}

func init() { proto.RegisterFile("flyteidl/admin/task.proto", fileDescriptor_task_ec178c90a216dba4) }

var fileDescriptor_task_ec178c90a216dba4 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x55, 0x9a, 0x8f, 0xa5, 0x53, 0xa1, 0x65, 0x2d, 0x84, 0xcc, 0x72, 0x49, 0x83, 0x54, 0xe5,
	0xd2, 0x44, 0xb4, 0xff, 0x00, 0x4e, 0x48, 0x9c, 0x4c, 0x4f, 0xdc, 0xd2, 0x64, 0x52, 0x4c, 0x63,
	0x3b, 0xd8, 0x0e, 0x52, 0xf9, 0x33, 0xfc, 0x55, 0x64, 0x27, 0x29, 0x85, 0x03, 0x97, 0xbd, 0xf9,
	0xcd, 0xbc, 0x37, 0x6f, 0xf2, 0x32, 0xf0, 0xba, 0xed, 0x2e, 0x16, 0x79, 0xd3, 0x95, 0x55, 0x23,
	0xb8, 0x2c, 0x6d, 0x65, 0xce, 0x45, 0xaf, 0x95, 0x55, 0x24, 0xf6, 0x95, 0xc7, 0x3f, 0x8c, 0x5a,
	0x69, 0xf4, 0x04, 0x33, 0x32, 0x1e, 0xdf, 0xfc, 0x23, 0xae, 0x95, 0x10, 0x4a, 0x8e, 0xcd, 0x6c,
	0x80, 0x87, 0x43, 0x65, 0xce, 0x1f, 0x34, 0x56, 0x16, 0x19, 0x7e, 0x1f, 0xd0, 0x58, 0xb2, 0x86,
	0x05, 0x6f, 0x68, 0x90, 0x06, 0xf9, 0x6a, 0xf7, 0x50, 0x78, 0x55, 0xf1, 0xb1, 0x41, 0x69, 0x79,
	0xcb, 0x51, 0xb3, 0x05, 0x6f, 0x08, 0x85, 0xbb, 0x1f, 0xa8, 0x0d, 0x57, 0x92, 0x2e, 0xd2, 0x20,
	0x5f, 0xb2, 0x19, 0x92, 0xb7, 0x10, 0x99, 0x1e, 0x6b, 0x1a, 0x7a, 0xf9, 0xfd, 0x24, 0x77, 0x26,
	0x9f, 0x7b, 0xac, 0x99, 0x6f, 0x66, 0xbf, 0x02, 0xb8, 0x77, 0xa5, 0x4f, 0xdc, 0xd8, 0xd9, 0x95,
	0xc2, 0x5d, 0xaf, 0xd5, 0x37, 0xac, 0xad, 0xb7, 0x5e, 0xb2, 0x19, 0x92, 0x57, 0x90, 0x34, 0x4a,
	0x54, 0x7c, 0xf6, 0x9a, 0x10, 0x21, 0x10, 0xc9, 0x4a, 0xa0, 0xb7, 0x5a, 0x32, 0xff, 0x26, 0x2f,
	0x21, 0xee, 0xb8, 0xe0, 0x96, 0x46, 0x69, 0x90, 0x3f, 0x67, 0x23, 0x70, 0x13, 0x54, 0xdb, 0x1a,
	0xb4, 0x34, 0xf6, 0xe5, 0x09, 0x39, 0xcf, 0x96, 0x77, 0x16, 0xb5, 0xa1, 0xc9, 0xe8, 0x39, 0xc1,
	0x6c, 0x03, 0xe4, 0x36, 0x18, 0xd3, 0x2b, 0x69, 0x90, 0xbc, 0x80, 0x70, 0xd0, 0x72, 0xda, 0xcf,
	0x3d, 0xb3, 0x9f, 0x10, 0x39, 0xde, 0xd3, 0x32, 0x9b, 0xc6, 0x86, 0xd7, 0xb1, 0xd7, 0x14, 0x93,
	0xff, 0xa5, 0xb8, 0x85, 0x67, 0x73, 0x88, 0x64, 0x0d, 0xb1, 0xff, 0xe9, 0x34, 0x48, 0xc3, 0x7c,
	0xb5, 0x5b, 0xdd, 0x28, 0xd8, 0xd8, 0xc9, 0x76, 0x23, 0xdd, 0x0d, 0x20, 0x1b, 0x88, 0x5c, 0x71,
	0x5a, 0x98, 0x14, 0xee, 0x6a, 0x3c, 0xf9, 0x80, 0xa2, 0xef, 0xdc, 0x27, 0xfb, 0xfe, 0xfb, 0xfd,
	0x97, 0x77, 0x27, 0x6e, 0xbf, 0x0e, 0xc7, 0xa2, 0x56, 0xa2, 0xec, 0x2e, 0xad, 0x2d, 0xaf, 0xe7,
	0x74, 0x42, 0x59, 0xf6, 0xc7, 0xed, 0x49, 0x95, 0x7f, 0x5f, 0xd8, 0x31, 0xf1, 0xb7, 0xb5, 0xff,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x9f, 0x1a, 0x8f, 0xb7, 0x02, 0x00, 0x00,
}