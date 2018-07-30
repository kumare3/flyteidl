// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/service/admin.proto

package lyft_flyte_flyteadmin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import admin "github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminServiceClient interface {
	CreateTask(ctx context.Context, in *admin.TaskCreateRequest, opts ...grpc.CallOption) (*admin.TaskCreateResponse, error)
	ListTaskIds(ctx context.Context, in *admin.TaskListRequest, opts ...grpc.CallOption) (*admin.IdentifierList, error)
	ListTasks(ctx context.Context, in *admin.TaskListRequest, opts ...grpc.CallOption) (*admin.TaskList, error)
	CreateWorkflow(ctx context.Context, in *admin.WorkflowCreateRequest, opts ...grpc.CallOption) (*admin.WorkflowCreateResponse, error)
	ListWorkflowIds(ctx context.Context, in *admin.WorkflowListRequest, opts ...grpc.CallOption) (*admin.IdentifierList, error)
	ListWorkflows(ctx context.Context, in *admin.WorkflowListRequest, opts ...grpc.CallOption) (*admin.WorkflowList, error)
}

type adminServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdminServiceClient(cc *grpc.ClientConn) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) CreateTask(ctx context.Context, in *admin.TaskCreateRequest, opts ...grpc.CallOption) (*admin.TaskCreateResponse, error) {
	out := new(admin.TaskCreateResponse)
	err := c.cc.Invoke(ctx, "/lyft.flyte.flyteadmin.AdminService/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) ListTaskIds(ctx context.Context, in *admin.TaskListRequest, opts ...grpc.CallOption) (*admin.IdentifierList, error) {
	out := new(admin.IdentifierList)
	err := c.cc.Invoke(ctx, "/lyft.flyte.flyteadmin.AdminService/ListTaskIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) ListTasks(ctx context.Context, in *admin.TaskListRequest, opts ...grpc.CallOption) (*admin.TaskList, error) {
	out := new(admin.TaskList)
	err := c.cc.Invoke(ctx, "/lyft.flyte.flyteadmin.AdminService/ListTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) CreateWorkflow(ctx context.Context, in *admin.WorkflowCreateRequest, opts ...grpc.CallOption) (*admin.WorkflowCreateResponse, error) {
	out := new(admin.WorkflowCreateResponse)
	err := c.cc.Invoke(ctx, "/lyft.flyte.flyteadmin.AdminService/CreateWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) ListWorkflowIds(ctx context.Context, in *admin.WorkflowListRequest, opts ...grpc.CallOption) (*admin.IdentifierList, error) {
	out := new(admin.IdentifierList)
	err := c.cc.Invoke(ctx, "/lyft.flyte.flyteadmin.AdminService/ListWorkflowIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) ListWorkflows(ctx context.Context, in *admin.WorkflowListRequest, opts ...grpc.CallOption) (*admin.WorkflowList, error) {
	out := new(admin.WorkflowList)
	err := c.cc.Invoke(ctx, "/lyft.flyte.flyteadmin.AdminService/ListWorkflows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
type AdminServiceServer interface {
	CreateTask(context.Context, *admin.TaskCreateRequest) (*admin.TaskCreateResponse, error)
	ListTaskIds(context.Context, *admin.TaskListRequest) (*admin.IdentifierList, error)
	ListTasks(context.Context, *admin.TaskListRequest) (*admin.TaskList, error)
	CreateWorkflow(context.Context, *admin.WorkflowCreateRequest) (*admin.WorkflowCreateResponse, error)
	ListWorkflowIds(context.Context, *admin.WorkflowListRequest) (*admin.IdentifierList, error)
	ListWorkflows(context.Context, *admin.WorkflowListRequest) (*admin.WorkflowList, error)
}

func RegisterAdminServiceServer(s *grpc.Server, srv AdminServiceServer) {
	s.RegisterService(&_AdminService_serviceDesc, srv)
}

func _AdminService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.TaskCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lyft.flyte.flyteadmin.AdminService/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CreateTask(ctx, req.(*admin.TaskCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_ListTaskIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.TaskListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ListTaskIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lyft.flyte.flyteadmin.AdminService/ListTaskIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ListTaskIds(ctx, req.(*admin.TaskListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_ListTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.TaskListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ListTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lyft.flyte.flyteadmin.AdminService/ListTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ListTasks(ctx, req.(*admin.TaskListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_CreateWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.WorkflowCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CreateWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lyft.flyte.flyteadmin.AdminService/CreateWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CreateWorkflow(ctx, req.(*admin.WorkflowCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_ListWorkflowIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.WorkflowListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ListWorkflowIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lyft.flyte.flyteadmin.AdminService/ListWorkflowIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ListWorkflowIds(ctx, req.(*admin.WorkflowListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_ListWorkflows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.WorkflowListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ListWorkflows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lyft.flyte.flyteadmin.AdminService/ListWorkflows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ListWorkflows(ctx, req.(*admin.WorkflowListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdminService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lyft.flyte.flyteadmin.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _AdminService_CreateTask_Handler,
		},
		{
			MethodName: "ListTaskIds",
			Handler:    _AdminService_ListTaskIds_Handler,
		},
		{
			MethodName: "ListTasks",
			Handler:    _AdminService_ListTasks_Handler,
		},
		{
			MethodName: "CreateWorkflow",
			Handler:    _AdminService_CreateWorkflow_Handler,
		},
		{
			MethodName: "ListWorkflowIds",
			Handler:    _AdminService_ListWorkflowIds_Handler,
		},
		{
			MethodName: "ListWorkflows",
			Handler:    _AdminService_ListWorkflows_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flyteidl/service/admin.proto",
}

func init() { proto.RegisterFile("flyteidl/service/admin.proto", fileDescriptor_admin_d60d49e580fd0ed2) }

var fileDescriptor_admin_d60d49e580fd0ed2 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xa6, 0x1e, 0x06, 0x46, 0xe7, 0x30, 0xb2, 0xb1, 0xd5, 0xee, 0xd2, 0xa3, 0x87, 0x05, 0x37,
	0x11, 0x19, 0x22, 0x88, 0xa7, 0x81, 0x27, 0x15, 0x84, 0x79, 0x90, 0xb8, 0xa6, 0x23, 0xb6, 0x4d,
	0xba, 0x26, 0x6e, 0x8e, 0xb1, 0x8b, 0xf8, 0x06, 0x3e, 0x84, 0x0f, 0xe1, 0xc9, 0x67, 0xf0, 0x15,
	0x7c, 0x10, 0x69, 0x9a, 0xc8, 0x2c, 0x9b, 0x58, 0x2f, 0x5b, 0xfb, 0x7d, 0xf9, 0xbf, 0xef, 0xfb,
	0xfb, 0xb5, 0xc0, 0xf1, 0xc3, 0xa9, 0x24, 0xd4, 0x0b, 0x91, 0x20, 0xc9, 0x98, 0x0e, 0x08, 0xc2,
	0x5e, 0x44, 0x59, 0x2b, 0x4e, 0xb8, 0xe4, 0xb0, 0x1a, 0x4e, 0x7d, 0xd9, 0x52, 0x47, 0xb2, 0x5f,
	0x45, 0xda, 0xce, 0x90, 0xf3, 0x61, 0x48, 0x10, 0x8e, 0x29, 0xc2, 0x8c, 0x71, 0x89, 0x25, 0xe5,
	0x4c, 0x64, 0x43, 0x76, 0xe3, 0x5b, 0x52, 0x9d, 0x46, 0x12, 0x8b, 0x40, 0x53, 0xcd, 0x1c, 0x35,
	0xe1, 0x49, 0xe0, 0x87, 0x7c, 0xa2, 0xe9, 0xdd, 0x1c, 0x3d, 0xe0, 0x51, 0xc4, 0x75, 0x96, 0xf6,
	0x7b, 0x09, 0x6c, 0x9e, 0xa6, 0xf0, 0x65, 0x16, 0x14, 0xde, 0x00, 0x70, 0x96, 0x10, 0x2c, 0xc9,
	0x15, 0x16, 0x01, 0xac, 0xb7, 0xb2, 0xe0, 0xe9, 0x4d, 0x06, 0x5f, 0x90, 0xd1, 0x03, 0x11, 0xd2,
	0x6e, 0x2c, 0x61, 0x44, 0xcc, 0x99, 0x20, 0x6e, 0xfd, 0xe9, 0xe3, 0xf3, 0x65, 0x0d, 0xba, 0x65,
	0xb5, 0xca, 0x78, 0x5f, 0x65, 0x15, 0x5d, 0x6b, 0x0f, 0x3e, 0x82, 0x8d, 0x73, 0x2a, 0x64, 0x3a,
	0xd3, 0xf3, 0x04, 0xac, 0x2d, 0x68, 0xa4, 0xb8, 0xd1, 0xae, 0x6a, 0xbc, 0xe7, 0x11, 0x26, 0xa9,
	0x4f, 0x49, 0x92, 0xb2, 0x6e, 0x57, 0xe9, 0x1e, 0xc0, 0xb6, 0xd1, 0x8d, 0x13, 0x7e, 0x4f, 0x06,
	0x12, 0xcd, 0xf4, 0xc5, 0x1c, 0x79, 0x3c, 0xc2, 0x94, 0xa1, 0x59, 0xf6, 0x3f, 0x57, 0xce, 0xb7,
	0xd4, 0x13, 0xf0, 0xd5, 0x02, 0xeb, 0xc6, 0x7a, 0xb5, 0x71, 0x25, 0x87, 0xbb, 0x23, 0x65, 0x19,
	0x40, 0x54, 0xcc, 0x52, 0xf4, 0x8f, 0xe0, 0x61, 0xc1, 0x11, 0x34, 0x63, 0x38, 0x22, 0x73, 0x18,
	0x80, 0xad, 0xec, 0x79, 0x5e, 0xeb, 0x1a, 0xa1, 0xa3, 0x53, 0x19, 0xe0, 0x67, 0x11, 0xcd, 0x15,
	0xac, 0x2e, 0xc3, 0x51, 0x1b, 0xd4, 0xdc, 0x6d, 0x13, 0xc7, 0xbc, 0x1d, 0xaa, 0x90, 0x67, 0x0b,
	0x54, 0xd2, 0x45, 0xcd, 0x70, 0xda, 0x8a, 0x9d, 0x13, 0xfc, 0x43, 0x33, 0x27, 0xca, 0xa4, 0xd0,
	0xce, 0x26, 0x86, 0x6a, 0xe7, 0xcd, 0x02, 0xe5, 0xc5, 0x18, 0xbf, 0x87, 0xd8, 0x59, 0xc2, 0xb9,
	0x53, 0x15, 0x41, 0xc0, 0x4e, 0xf1, 0x08, 0xa2, 0x7f, 0x0c, 0xbb, 0xff, 0x18, 0xd3, 0x8d, 0xdd,
	0x95, 0xd4, 0xa7, 0xd4, 0xf9, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x71, 0x24, 0xb7, 0x73, 0xf6, 0x03,
	0x00, 0x00,
}