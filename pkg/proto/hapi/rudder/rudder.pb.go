// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hapi/rudder/rudder.proto

/*
Package rudder is a generated protocol buffer package.

It is generated from these files:
	hapi/rudder/rudder.proto

It has these top-level messages:
	Result
	VersionReleaseRequest
	VersionReleaseResponse
	InstallReleaseRequest
	InstallReleaseResponse
	DeleteReleaseRequest
	DeleteReleaseResponse
	UpgradeReleaseRequest
	UpgradeReleaseResponse
	RollbackReleaseRequest
	RollbackReleaseResponse
	ReleaseStatusRequest
	ReleaseStatusResponse
*/
package rudder

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import hapi_release3 "github.com/paalka/helm/pkg/proto/hapi/release"
import hapi_release5 "github.com/paalka/helm/pkg/proto/hapi/release"

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

type Result_Status int32

const (
	// No status set
	Result_UNKNOWN Result_Status = 0
	// Operation was successful
	Result_SUCCESS Result_Status = 1
	// Operation had no results (e.g. upgrade identical, rollback to same, delete non-existent)
	Result_UNCHANGED Result_Status = 2
	// Operation failed
	Result_ERROR Result_Status = 3
)

var Result_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "SUCCESS",
	2: "UNCHANGED",
	3: "ERROR",
}
var Result_Status_value = map[string]int32{
	"UNKNOWN":   0,
	"SUCCESS":   1,
	"UNCHANGED": 2,
	"ERROR":     3,
}

func (x Result_Status) String() string {
	return proto.EnumName(Result_Status_name, int32(x))
}
func (Result_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Result struct {
	Info string   `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
	Log  []string `protobuf:"bytes,2,rep,name=log" json:"log,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Result) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *Result) GetLog() []string {
	if m != nil {
		return m.Log
	}
	return nil
}

type VersionReleaseRequest struct {
}

func (m *VersionReleaseRequest) Reset()                    { *m = VersionReleaseRequest{} }
func (m *VersionReleaseRequest) String() string            { return proto.CompactTextString(m) }
func (*VersionReleaseRequest) ProtoMessage()               {}
func (*VersionReleaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type VersionReleaseResponse struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *VersionReleaseResponse) Reset()                    { *m = VersionReleaseResponse{} }
func (m *VersionReleaseResponse) String() string            { return proto.CompactTextString(m) }
func (*VersionReleaseResponse) ProtoMessage()               {}
func (*VersionReleaseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *VersionReleaseResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VersionReleaseResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type InstallReleaseRequest struct {
	Release *hapi_release5.Release `protobuf:"bytes,1,opt,name=release" json:"release,omitempty"`
}

func (m *InstallReleaseRequest) Reset()                    { *m = InstallReleaseRequest{} }
func (m *InstallReleaseRequest) String() string            { return proto.CompactTextString(m) }
func (*InstallReleaseRequest) ProtoMessage()               {}
func (*InstallReleaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *InstallReleaseRequest) GetRelease() *hapi_release5.Release {
	if m != nil {
		return m.Release
	}
	return nil
}

type InstallReleaseResponse struct {
	Release *hapi_release5.Release `protobuf:"bytes,1,opt,name=release" json:"release,omitempty"`
	Result  *Result                `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
}

func (m *InstallReleaseResponse) Reset()                    { *m = InstallReleaseResponse{} }
func (m *InstallReleaseResponse) String() string            { return proto.CompactTextString(m) }
func (*InstallReleaseResponse) ProtoMessage()               {}
func (*InstallReleaseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *InstallReleaseResponse) GetRelease() *hapi_release5.Release {
	if m != nil {
		return m.Release
	}
	return nil
}

func (m *InstallReleaseResponse) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

type DeleteReleaseRequest struct {
	Release *hapi_release5.Release `protobuf:"bytes,1,opt,name=release" json:"release,omitempty"`
}

func (m *DeleteReleaseRequest) Reset()                    { *m = DeleteReleaseRequest{} }
func (m *DeleteReleaseRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteReleaseRequest) ProtoMessage()               {}
func (*DeleteReleaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeleteReleaseRequest) GetRelease() *hapi_release5.Release {
	if m != nil {
		return m.Release
	}
	return nil
}

type DeleteReleaseResponse struct {
	Release *hapi_release5.Release `protobuf:"bytes,1,opt,name=release" json:"release,omitempty"`
	Result  *Result                `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
}

func (m *DeleteReleaseResponse) Reset()                    { *m = DeleteReleaseResponse{} }
func (m *DeleteReleaseResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteReleaseResponse) ProtoMessage()               {}
func (*DeleteReleaseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteReleaseResponse) GetRelease() *hapi_release5.Release {
	if m != nil {
		return m.Release
	}
	return nil
}

func (m *DeleteReleaseResponse) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

type UpgradeReleaseRequest struct {
	Current  *hapi_release5.Release `protobuf:"bytes,1,opt,name=current" json:"current,omitempty"`
	Target   *hapi_release5.Release `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
	Timeout  int64                  `protobuf:"varint,3,opt,name=Timeout" json:"Timeout,omitempty"`
	Wait     bool                   `protobuf:"varint,4,opt,name=Wait" json:"Wait,omitempty"`
	Recreate bool                   `protobuf:"varint,5,opt,name=Recreate" json:"Recreate,omitempty"`
	Force    bool                   `protobuf:"varint,6,opt,name=Force" json:"Force,omitempty"`
}

func (m *UpgradeReleaseRequest) Reset()                    { *m = UpgradeReleaseRequest{} }
func (m *UpgradeReleaseRequest) String() string            { return proto.CompactTextString(m) }
func (*UpgradeReleaseRequest) ProtoMessage()               {}
func (*UpgradeReleaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UpgradeReleaseRequest) GetCurrent() *hapi_release5.Release {
	if m != nil {
		return m.Current
	}
	return nil
}

func (m *UpgradeReleaseRequest) GetTarget() *hapi_release5.Release {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *UpgradeReleaseRequest) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *UpgradeReleaseRequest) GetWait() bool {
	if m != nil {
		return m.Wait
	}
	return false
}

func (m *UpgradeReleaseRequest) GetRecreate() bool {
	if m != nil {
		return m.Recreate
	}
	return false
}

func (m *UpgradeReleaseRequest) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

type UpgradeReleaseResponse struct {
	Release *hapi_release5.Release `protobuf:"bytes,1,opt,name=release" json:"release,omitempty"`
	Result  *Result                `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
}

func (m *UpgradeReleaseResponse) Reset()                    { *m = UpgradeReleaseResponse{} }
func (m *UpgradeReleaseResponse) String() string            { return proto.CompactTextString(m) }
func (*UpgradeReleaseResponse) ProtoMessage()               {}
func (*UpgradeReleaseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UpgradeReleaseResponse) GetRelease() *hapi_release5.Release {
	if m != nil {
		return m.Release
	}
	return nil
}

func (m *UpgradeReleaseResponse) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

type RollbackReleaseRequest struct {
	Current  *hapi_release5.Release `protobuf:"bytes,1,opt,name=current" json:"current,omitempty"`
	Target   *hapi_release5.Release `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
	Timeout  int64                  `protobuf:"varint,3,opt,name=Timeout" json:"Timeout,omitempty"`
	Wait     bool                   `protobuf:"varint,4,opt,name=Wait" json:"Wait,omitempty"`
	Recreate bool                   `protobuf:"varint,5,opt,name=Recreate" json:"Recreate,omitempty"`
	Force    bool                   `protobuf:"varint,6,opt,name=Force" json:"Force,omitempty"`
}

func (m *RollbackReleaseRequest) Reset()                    { *m = RollbackReleaseRequest{} }
func (m *RollbackReleaseRequest) String() string            { return proto.CompactTextString(m) }
func (*RollbackReleaseRequest) ProtoMessage()               {}
func (*RollbackReleaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RollbackReleaseRequest) GetCurrent() *hapi_release5.Release {
	if m != nil {
		return m.Current
	}
	return nil
}

func (m *RollbackReleaseRequest) GetTarget() *hapi_release5.Release {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *RollbackReleaseRequest) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *RollbackReleaseRequest) GetWait() bool {
	if m != nil {
		return m.Wait
	}
	return false
}

func (m *RollbackReleaseRequest) GetRecreate() bool {
	if m != nil {
		return m.Recreate
	}
	return false
}

func (m *RollbackReleaseRequest) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

type RollbackReleaseResponse struct {
	Release *hapi_release5.Release `protobuf:"bytes,1,opt,name=release" json:"release,omitempty"`
	Result  *Result                `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
}

func (m *RollbackReleaseResponse) Reset()                    { *m = RollbackReleaseResponse{} }
func (m *RollbackReleaseResponse) String() string            { return proto.CompactTextString(m) }
func (*RollbackReleaseResponse) ProtoMessage()               {}
func (*RollbackReleaseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *RollbackReleaseResponse) GetRelease() *hapi_release5.Release {
	if m != nil {
		return m.Release
	}
	return nil
}

func (m *RollbackReleaseResponse) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

type ReleaseStatusRequest struct {
	Release *hapi_release5.Release `protobuf:"bytes,1,opt,name=release" json:"release,omitempty"`
}

func (m *ReleaseStatusRequest) Reset()                    { *m = ReleaseStatusRequest{} }
func (m *ReleaseStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*ReleaseStatusRequest) ProtoMessage()               {}
func (*ReleaseStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ReleaseStatusRequest) GetRelease() *hapi_release5.Release {
	if m != nil {
		return m.Release
	}
	return nil
}

type ReleaseStatusResponse struct {
	Release *hapi_release5.Release `protobuf:"bytes,1,opt,name=release" json:"release,omitempty"`
	Info    *hapi_release3.Info    `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
}

func (m *ReleaseStatusResponse) Reset()                    { *m = ReleaseStatusResponse{} }
func (m *ReleaseStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*ReleaseStatusResponse) ProtoMessage()               {}
func (*ReleaseStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *ReleaseStatusResponse) GetRelease() *hapi_release5.Release {
	if m != nil {
		return m.Release
	}
	return nil
}

func (m *ReleaseStatusResponse) GetInfo() *hapi_release3.Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Result)(nil), "hapi.services.rudder.Result")
	proto.RegisterType((*VersionReleaseRequest)(nil), "hapi.services.rudder.VersionReleaseRequest")
	proto.RegisterType((*VersionReleaseResponse)(nil), "hapi.services.rudder.VersionReleaseResponse")
	proto.RegisterType((*InstallReleaseRequest)(nil), "hapi.services.rudder.InstallReleaseRequest")
	proto.RegisterType((*InstallReleaseResponse)(nil), "hapi.services.rudder.InstallReleaseResponse")
	proto.RegisterType((*DeleteReleaseRequest)(nil), "hapi.services.rudder.DeleteReleaseRequest")
	proto.RegisterType((*DeleteReleaseResponse)(nil), "hapi.services.rudder.DeleteReleaseResponse")
	proto.RegisterType((*UpgradeReleaseRequest)(nil), "hapi.services.rudder.UpgradeReleaseRequest")
	proto.RegisterType((*UpgradeReleaseResponse)(nil), "hapi.services.rudder.UpgradeReleaseResponse")
	proto.RegisterType((*RollbackReleaseRequest)(nil), "hapi.services.rudder.RollbackReleaseRequest")
	proto.RegisterType((*RollbackReleaseResponse)(nil), "hapi.services.rudder.RollbackReleaseResponse")
	proto.RegisterType((*ReleaseStatusRequest)(nil), "hapi.services.rudder.ReleaseStatusRequest")
	proto.RegisterType((*ReleaseStatusResponse)(nil), "hapi.services.rudder.ReleaseStatusResponse")
	proto.RegisterEnum("hapi.services.rudder.Result_Status", Result_Status_name, Result_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ReleaseModuleService service

type ReleaseModuleServiceClient interface {
	Version(ctx context.Context, in *VersionReleaseRequest, opts ...grpc.CallOption) (*VersionReleaseResponse, error)
	// InstallRelease requests installation of a chart as a new release.
	InstallRelease(ctx context.Context, in *InstallReleaseRequest, opts ...grpc.CallOption) (*InstallReleaseResponse, error)
	// DeleteRelease requests deletion of a named release.
	DeleteRelease(ctx context.Context, in *DeleteReleaseRequest, opts ...grpc.CallOption) (*DeleteReleaseResponse, error)
	// RollbackRelease rolls back a release to a previous version.
	RollbackRelease(ctx context.Context, in *RollbackReleaseRequest, opts ...grpc.CallOption) (*RollbackReleaseResponse, error)
	// UpgradeRelease updates release content.
	UpgradeRelease(ctx context.Context, in *UpgradeReleaseRequest, opts ...grpc.CallOption) (*UpgradeReleaseResponse, error)
	// ReleaseStatus retrieves release status.
	ReleaseStatus(ctx context.Context, in *ReleaseStatusRequest, opts ...grpc.CallOption) (*ReleaseStatusResponse, error)
}

type releaseModuleServiceClient struct {
	cc *grpc.ClientConn
}

func NewReleaseModuleServiceClient(cc *grpc.ClientConn) ReleaseModuleServiceClient {
	return &releaseModuleServiceClient{cc}
}

func (c *releaseModuleServiceClient) Version(ctx context.Context, in *VersionReleaseRequest, opts ...grpc.CallOption) (*VersionReleaseResponse, error) {
	out := new(VersionReleaseResponse)
	err := grpc.Invoke(ctx, "/hapi.services.rudder.ReleaseModuleService/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseModuleServiceClient) InstallRelease(ctx context.Context, in *InstallReleaseRequest, opts ...grpc.CallOption) (*InstallReleaseResponse, error) {
	out := new(InstallReleaseResponse)
	err := grpc.Invoke(ctx, "/hapi.services.rudder.ReleaseModuleService/InstallRelease", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseModuleServiceClient) DeleteRelease(ctx context.Context, in *DeleteReleaseRequest, opts ...grpc.CallOption) (*DeleteReleaseResponse, error) {
	out := new(DeleteReleaseResponse)
	err := grpc.Invoke(ctx, "/hapi.services.rudder.ReleaseModuleService/DeleteRelease", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseModuleServiceClient) RollbackRelease(ctx context.Context, in *RollbackReleaseRequest, opts ...grpc.CallOption) (*RollbackReleaseResponse, error) {
	out := new(RollbackReleaseResponse)
	err := grpc.Invoke(ctx, "/hapi.services.rudder.ReleaseModuleService/RollbackRelease", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseModuleServiceClient) UpgradeRelease(ctx context.Context, in *UpgradeReleaseRequest, opts ...grpc.CallOption) (*UpgradeReleaseResponse, error) {
	out := new(UpgradeReleaseResponse)
	err := grpc.Invoke(ctx, "/hapi.services.rudder.ReleaseModuleService/UpgradeRelease", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseModuleServiceClient) ReleaseStatus(ctx context.Context, in *ReleaseStatusRequest, opts ...grpc.CallOption) (*ReleaseStatusResponse, error) {
	out := new(ReleaseStatusResponse)
	err := grpc.Invoke(ctx, "/hapi.services.rudder.ReleaseModuleService/ReleaseStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ReleaseModuleService service

type ReleaseModuleServiceServer interface {
	Version(context.Context, *VersionReleaseRequest) (*VersionReleaseResponse, error)
	// InstallRelease requests installation of a chart as a new release.
	InstallRelease(context.Context, *InstallReleaseRequest) (*InstallReleaseResponse, error)
	// DeleteRelease requests deletion of a named release.
	DeleteRelease(context.Context, *DeleteReleaseRequest) (*DeleteReleaseResponse, error)
	// RollbackRelease rolls back a release to a previous version.
	RollbackRelease(context.Context, *RollbackReleaseRequest) (*RollbackReleaseResponse, error)
	// UpgradeRelease updates release content.
	UpgradeRelease(context.Context, *UpgradeReleaseRequest) (*UpgradeReleaseResponse, error)
	// ReleaseStatus retrieves release status.
	ReleaseStatus(context.Context, *ReleaseStatusRequest) (*ReleaseStatusResponse, error)
}

func RegisterReleaseModuleServiceServer(s *grpc.Server, srv ReleaseModuleServiceServer) {
	s.RegisterService(&_ReleaseModuleService_serviceDesc, srv)
}

func _ReleaseModuleService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseModuleServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hapi.services.rudder.ReleaseModuleService/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseModuleServiceServer).Version(ctx, req.(*VersionReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseModuleService_InstallRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstallReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseModuleServiceServer).InstallRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hapi.services.rudder.ReleaseModuleService/InstallRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseModuleServiceServer).InstallRelease(ctx, req.(*InstallReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseModuleService_DeleteRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseModuleServiceServer).DeleteRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hapi.services.rudder.ReleaseModuleService/DeleteRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseModuleServiceServer).DeleteRelease(ctx, req.(*DeleteReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseModuleService_RollbackRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseModuleServiceServer).RollbackRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hapi.services.rudder.ReleaseModuleService/RollbackRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseModuleServiceServer).RollbackRelease(ctx, req.(*RollbackReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseModuleService_UpgradeRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpgradeReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseModuleServiceServer).UpgradeRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hapi.services.rudder.ReleaseModuleService/UpgradeRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseModuleServiceServer).UpgradeRelease(ctx, req.(*UpgradeReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseModuleService_ReleaseStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseModuleServiceServer).ReleaseStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hapi.services.rudder.ReleaseModuleService/ReleaseStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseModuleServiceServer).ReleaseStatus(ctx, req.(*ReleaseStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReleaseModuleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hapi.services.rudder.ReleaseModuleService",
	HandlerType: (*ReleaseModuleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _ReleaseModuleService_Version_Handler,
		},
		{
			MethodName: "InstallRelease",
			Handler:    _ReleaseModuleService_InstallRelease_Handler,
		},
		{
			MethodName: "DeleteRelease",
			Handler:    _ReleaseModuleService_DeleteRelease_Handler,
		},
		{
			MethodName: "RollbackRelease",
			Handler:    _ReleaseModuleService_RollbackRelease_Handler,
		},
		{
			MethodName: "UpgradeRelease",
			Handler:    _ReleaseModuleService_UpgradeRelease_Handler,
		},
		{
			MethodName: "ReleaseStatus",
			Handler:    _ReleaseModuleService_ReleaseStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hapi/rudder/rudder.proto",
}

func init() { proto.RegisterFile("hapi/rudder/rudder.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 597 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x56, 0x5f, 0x8f, 0xd2, 0x4e,
	0x14, 0xa5, 0xb0, 0x14, 0xb8, 0x64, 0x7f, 0x3f, 0x32, 0xa1, 0xd0, 0x34, 0x3e, 0x90, 0x3e, 0x18,
	0xe2, 0xba, 0x25, 0x41, 0x1f, 0x7d, 0x51, 0x96, 0xfd, 0x13, 0x23, 0x9b, 0x0c, 0xe2, 0x26, 0xbe,
	0x75, 0xe1, 0x82, 0xd5, 0xd2, 0xd6, 0xe9, 0x74, 0x1f, 0xd5, 0x4f, 0xe3, 0x57, 0xd2, 0x8f, 0x63,
	0xda, 0x69, 0x89, 0xad, 0xd3, 0x88, 0x6b, 0xc2, 0x83, 0x4f, 0x9d, 0xe9, 0x3d, 0xdc, 0x39, 0xe7,
	0xf4, 0xce, 0x09, 0xa0, 0xbf, 0xb3, 0x03, 0x67, 0xc4, 0xa2, 0xd5, 0x0a, 0x59, 0xfa, 0xb0, 0x02,
	0xe6, 0x73, 0x9f, 0x74, 0xe3, 0x8a, 0x15, 0x22, 0xbb, 0x73, 0x96, 0x18, 0x5a, 0xa2, 0x66, 0xf4,
	0x05, 0x1e, 0x5d, 0xb4, 0x43, 0x1c, 0x39, 0xde, 0xda, 0x17, 0x70, 0xc3, 0xc8, 0x15, 0xd2, 0xa7,
	0xa8, 0x99, 0x2e, 0xa8, 0x14, 0xc3, 0xc8, 0xe5, 0x84, 0xc0, 0x51, 0xfc, 0x1b, 0x5d, 0x19, 0x28,
	0xc3, 0x16, 0x4d, 0xd6, 0xa4, 0x03, 0x35, 0xd7, 0xdf, 0xe8, 0xd5, 0x41, 0x6d, 0xd8, 0xa2, 0xf1,
	0xd2, 0x7c, 0x06, 0xea, 0x9c, 0xdb, 0x3c, 0x0a, 0x49, 0x1b, 0x1a, 0x8b, 0xd9, 0xcb, 0xd9, 0xf5,
	0xcd, 0xac, 0x53, 0x89, 0x37, 0xf3, 0xc5, 0x64, 0x32, 0x9d, 0xcf, 0x3b, 0x0a, 0x39, 0x86, 0xd6,
	0x62, 0x36, 0xb9, 0x7c, 0x3e, 0xbb, 0x98, 0x9e, 0x75, 0xaa, 0xa4, 0x05, 0xf5, 0x29, 0xa5, 0xd7,
	0xb4, 0x53, 0x33, 0xfb, 0xa0, 0xbd, 0x41, 0x16, 0x3a, 0xbe, 0x47, 0x05, 0x0b, 0x8a, 0x1f, 0x23,
	0x0c, 0xb9, 0x79, 0x0e, 0xbd, 0x62, 0x21, 0x0c, 0x7c, 0x2f, 0xc4, 0x98, 0x96, 0x67, 0x6f, 0x31,
	0xa3, 0x15, 0xaf, 0x89, 0x0e, 0x8d, 0x3b, 0x81, 0xd6, 0xab, 0xc9, 0xeb, 0x6c, 0x6b, 0x5e, 0x82,
	0x76, 0xe5, 0x85, 0xdc, 0x76, 0xdd, 0xfc, 0x01, 0x64, 0x04, 0x8d, 0x54, 0x78, 0xd2, 0xa9, 0x3d,
	0xd6, 0xac, 0xc4, 0xc4, 0xcc, 0x8d, 0x0c, 0x9e, 0xa1, 0xcc, 0xcf, 0xd0, 0x2b, 0x76, 0x4a, 0x19,
	0xfd, 0x69, 0x2b, 0xf2, 0x14, 0x54, 0x96, 0x78, 0x9c, 0xb0, 0x6d, 0x8f, 0x1f, 0x58, 0xb2, 0xef,
	0x67, 0x89, 0xef, 0x40, 0x53, 0xac, 0x79, 0x01, 0xdd, 0x33, 0x74, 0x91, 0xe3, 0xdf, 0x2a, 0xf9,
	0x04, 0x5a, 0xa1, 0xd1, 0x61, 0x85, 0x7c, 0x53, 0x40, 0x5b, 0x04, 0x1b, 0x66, 0xaf, 0x24, 0x52,
	0x96, 0x11, 0x63, 0xe8, 0xf1, 0xdf, 0x10, 0x48, 0x51, 0xe4, 0x14, 0x54, 0x6e, 0xb3, 0x0d, 0x66,
	0x04, 0x4a, 0xf0, 0x29, 0x28, 0x9e, 0x93, 0xd7, 0xce, 0x16, 0xfd, 0x88, 0xeb, 0xb5, 0x81, 0x32,
	0xac, 0xd1, 0x6c, 0x1b, 0x4f, 0xd5, 0x8d, 0xed, 0x70, 0xfd, 0x68, 0xa0, 0x0c, 0x9b, 0x34, 0x59,
	0x13, 0x03, 0x9a, 0x14, 0x97, 0x0c, 0x6d, 0x8e, 0x7a, 0x3d, 0x79, 0xbf, 0xdb, 0x93, 0x2e, 0xd4,
	0xcf, 0x7d, 0xb6, 0x44, 0x5d, 0x4d, 0x0a, 0x62, 0x13, 0xcf, 0x48, 0x51, 0xd8, 0x61, 0xad, 0xfd,
	0xae, 0x40, 0x8f, 0xfa, 0xae, 0x7b, 0x6b, 0x2f, 0x3f, 0xfc, 0x63, 0xde, 0x7e, 0x51, 0xa0, 0xff,
	0x8b, 0xb4, 0x83, 0xdf, 0xc0, 0xb4, 0x93, 0x88, 0xbc, 0x7b, 0xdf, 0xc0, 0x00, 0xb4, 0x42, 0xa3,
	0xfb, 0x0a, 0x79, 0x98, 0x86, 0xb4, 0x90, 0x41, 0xf2, 0xe8, 0x2b, 0x6f, 0xed, 0x8b, 0xe0, 0x1e,
	0x7f, 0xad, 0xef, 0xb8, 0xbf, 0xf2, 0x57, 0x91, 0x8b, 0x73, 0x21, 0x95, 0xac, 0xa1, 0x91, 0x06,
	0x2d, 0x39, 0x91, 0x9b, 0x20, 0x0d, 0x68, 0xe3, 0xf1, 0x7e, 0x60, 0xa1, 0xcb, 0xac, 0x90, 0x2d,
	0xfc, 0x97, 0x8f, 0xcf, 0xb2, 0xe3, 0xa4, 0x71, 0x5d, 0x76, 0x9c, 0x3c, 0x91, 0xcd, 0x0a, 0x79,
	0x0f, 0xc7, 0xb9, 0x8c, 0x23, 0x8f, 0xe4, 0x0d, 0x64, 0x89, 0x6a, 0x9c, 0xec, 0x85, 0xdd, 0x9d,
	0x15, 0xc0, 0xff, 0x85, 0xc1, 0x24, 0x25, 0x74, 0xe5, 0x57, 0xd3, 0x38, 0xdd, 0x13, 0xfd, 0xb3,
	0x99, 0xf9, 0x9c, 0x29, 0x33, 0x53, 0x1a, 0xb3, 0x65, 0x66, 0xca, 0xa3, 0x4b, 0x98, 0x99, 0x1b,
	0xd7, 0x32, 0x33, 0x65, 0x97, 0xa3, 0xcc, 0x4c, 0xe9, 0xfc, 0x9b, 0x95, 0x17, 0xcd, 0xb7, 0xaa,
	0x40, 0xdc, 0xaa, 0xc9, 0x1f, 0x92, 0x27, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xa5, 0x37,
	0x75, 0xf7, 0x08, 0x00, 0x00,
}
