// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: server/monit.proto

package server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Monitization_HealthCheck_FullMethodName         = "/monitization.Monitization/HealthCheck"
	Monitization_CreateWallet_FullMethodName        = "/monitization.Monitization/CreateWallet"
	Monitization_GetWallet_FullMethodName           = "/monitization.Monitization/GetWallet"
	Monitization_UpdateWallet_FullMethodName        = "/monitization.Monitization/UpdateWallet"
	Monitization_ParticipationReward_FullMethodName = "/monitization.Monitization/ParticipationReward"
	Monitization_UserRewardHistory_FullMethodName   = "/monitization.Monitization/UserRewardHistory"
	Monitization_VideoReward_FullMethodName         = "/monitization.Monitization/VideoReward"
	Monitization_ExclusiveContent_FullMethodName    = "/monitization.Monitization/ExclusiveContent"
	Monitization_GroupRewardHistory_FullMethodName  = "/monitization.Monitization/GroupRewardHistory"
	Monitization_UserWatchHour_FullMethodName       = "/monitization.Monitization/UserWatchHour"
	Monitization_ConferenceWatchHour_FullMethodName = "/monitization.Monitization/ConferenceWatchHour"
	Monitization_GroupWatchHour_FullMethodName      = "/monitization.Monitization/GroupWatchHour"
)

// MonitizationClient is the client API for Monitization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MonitizationClient interface {
	HealthCheck(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletResponse, error)
	GetWallet(ctx context.Context, in *GetWalletRequest, opts ...grpc.CallOption) (*GetWalletResponse, error)
	UpdateWallet(ctx context.Context, in *UpdateWalletRequest, opts ...grpc.CallOption) (*UpdateWalletResponse, error)
	ParticipationReward(ctx context.Context, in *ParticipationRewardRequest, opts ...grpc.CallOption) (*ParticipationRewardResponse, error)
	UserRewardHistory(ctx context.Context, in *UserRewardHistoryRequest, opts ...grpc.CallOption) (*UserRewardHistoryResponse, error)
	VideoReward(ctx context.Context, in *VideoRewardRequest, opts ...grpc.CallOption) (*VideoRewardResponse, error)
	ExclusiveContent(ctx context.Context, in *ExclusiveContentRequest, opts ...grpc.CallOption) (*ExclusiveContentResponse, error)
	GroupRewardHistory(ctx context.Context, in *GroupRewardHistoryRequest, opts ...grpc.CallOption) (*GroupRewardHistoryResponse, error)
	UserWatchHour(ctx context.Context, in *UserWatchHourRequest, opts ...grpc.CallOption) (*UserWatchHourResponse, error)
	ConferenceWatchHour(ctx context.Context, in *ConferenceWatchHourRequest, opts ...grpc.CallOption) (*ConferenceWatchHourResponse, error)
	GroupWatchHour(ctx context.Context, in *GroupWatchHourRequest, opts ...grpc.CallOption) (*GroupWatchHourResponse, error)
}

type monitizationClient struct {
	cc grpc.ClientConnInterface
}

func NewMonitizationClient(cc grpc.ClientConnInterface) MonitizationClient {
	return &monitizationClient{cc}
}

func (c *monitizationClient) HealthCheck(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Monitization_HealthCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitizationClient) CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletResponse, error) {
	out := new(CreateWalletResponse)
	err := c.cc.Invoke(ctx, Monitization_CreateWallet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitizationClient) GetWallet(ctx context.Context, in *GetWalletRequest, opts ...grpc.CallOption) (*GetWalletResponse, error) {
	out := new(GetWalletResponse)
	err := c.cc.Invoke(ctx, Monitization_GetWallet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitizationClient) UpdateWallet(ctx context.Context, in *UpdateWalletRequest, opts ...grpc.CallOption) (*UpdateWalletResponse, error) {
	out := new(UpdateWalletResponse)
	err := c.cc.Invoke(ctx, Monitization_UpdateWallet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitizationClient) ParticipationReward(ctx context.Context, in *ParticipationRewardRequest, opts ...grpc.CallOption) (*ParticipationRewardResponse, error) {
	out := new(ParticipationRewardResponse)
	err := c.cc.Invoke(ctx, Monitization_ParticipationReward_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitizationClient) UserRewardHistory(ctx context.Context, in *UserRewardHistoryRequest, opts ...grpc.CallOption) (*UserRewardHistoryResponse, error) {
	out := new(UserRewardHistoryResponse)
	err := c.cc.Invoke(ctx, Monitization_UserRewardHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitizationClient) VideoReward(ctx context.Context, in *VideoRewardRequest, opts ...grpc.CallOption) (*VideoRewardResponse, error) {
	out := new(VideoRewardResponse)
	err := c.cc.Invoke(ctx, Monitization_VideoReward_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitizationClient) ExclusiveContent(ctx context.Context, in *ExclusiveContentRequest, opts ...grpc.CallOption) (*ExclusiveContentResponse, error) {
	out := new(ExclusiveContentResponse)
	err := c.cc.Invoke(ctx, Monitization_ExclusiveContent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitizationClient) GroupRewardHistory(ctx context.Context, in *GroupRewardHistoryRequest, opts ...grpc.CallOption) (*GroupRewardHistoryResponse, error) {
	out := new(GroupRewardHistoryResponse)
	err := c.cc.Invoke(ctx, Monitization_GroupRewardHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitizationClient) UserWatchHour(ctx context.Context, in *UserWatchHourRequest, opts ...grpc.CallOption) (*UserWatchHourResponse, error) {
	out := new(UserWatchHourResponse)
	err := c.cc.Invoke(ctx, Monitization_UserWatchHour_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitizationClient) ConferenceWatchHour(ctx context.Context, in *ConferenceWatchHourRequest, opts ...grpc.CallOption) (*ConferenceWatchHourResponse, error) {
	out := new(ConferenceWatchHourResponse)
	err := c.cc.Invoke(ctx, Monitization_ConferenceWatchHour_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitizationClient) GroupWatchHour(ctx context.Context, in *GroupWatchHourRequest, opts ...grpc.CallOption) (*GroupWatchHourResponse, error) {
	out := new(GroupWatchHourResponse)
	err := c.cc.Invoke(ctx, Monitization_GroupWatchHour_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitizationServer is the server API for Monitization service.
// All implementations must embed UnimplementedMonitizationServer
// for forward compatibility
type MonitizationServer interface {
	HealthCheck(context.Context, *Request) (*Response, error)
	CreateWallet(context.Context, *CreateWalletRequest) (*CreateWalletResponse, error)
	GetWallet(context.Context, *GetWalletRequest) (*GetWalletResponse, error)
	UpdateWallet(context.Context, *UpdateWalletRequest) (*UpdateWalletResponse, error)
	ParticipationReward(context.Context, *ParticipationRewardRequest) (*ParticipationRewardResponse, error)
	UserRewardHistory(context.Context, *UserRewardHistoryRequest) (*UserRewardHistoryResponse, error)
	VideoReward(context.Context, *VideoRewardRequest) (*VideoRewardResponse, error)
	ExclusiveContent(context.Context, *ExclusiveContentRequest) (*ExclusiveContentResponse, error)
	GroupRewardHistory(context.Context, *GroupRewardHistoryRequest) (*GroupRewardHistoryResponse, error)
	UserWatchHour(context.Context, *UserWatchHourRequest) (*UserWatchHourResponse, error)
	ConferenceWatchHour(context.Context, *ConferenceWatchHourRequest) (*ConferenceWatchHourResponse, error)
	GroupWatchHour(context.Context, *GroupWatchHourRequest) (*GroupWatchHourResponse, error)
	mustEmbedUnimplementedMonitizationServer()
}

// UnimplementedMonitizationServer must be embedded to have forward compatible implementations.
type UnimplementedMonitizationServer struct {
}

func (UnimplementedMonitizationServer) HealthCheck(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedMonitizationServer) CreateWallet(context.Context, *CreateWalletRequest) (*CreateWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWallet not implemented")
}
func (UnimplementedMonitizationServer) GetWallet(context.Context, *GetWalletRequest) (*GetWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWallet not implemented")
}
func (UnimplementedMonitizationServer) UpdateWallet(context.Context, *UpdateWalletRequest) (*UpdateWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWallet not implemented")
}
func (UnimplementedMonitizationServer) ParticipationReward(context.Context, *ParticipationRewardRequest) (*ParticipationRewardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParticipationReward not implemented")
}
func (UnimplementedMonitizationServer) UserRewardHistory(context.Context, *UserRewardHistoryRequest) (*UserRewardHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRewardHistory not implemented")
}
func (UnimplementedMonitizationServer) VideoReward(context.Context, *VideoRewardRequest) (*VideoRewardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VideoReward not implemented")
}
func (UnimplementedMonitizationServer) ExclusiveContent(context.Context, *ExclusiveContentRequest) (*ExclusiveContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExclusiveContent not implemented")
}
func (UnimplementedMonitizationServer) GroupRewardHistory(context.Context, *GroupRewardHistoryRequest) (*GroupRewardHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupRewardHistory not implemented")
}
func (UnimplementedMonitizationServer) UserWatchHour(context.Context, *UserWatchHourRequest) (*UserWatchHourResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserWatchHour not implemented")
}
func (UnimplementedMonitizationServer) ConferenceWatchHour(context.Context, *ConferenceWatchHourRequest) (*ConferenceWatchHourResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConferenceWatchHour not implemented")
}
func (UnimplementedMonitizationServer) GroupWatchHour(context.Context, *GroupWatchHourRequest) (*GroupWatchHourResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupWatchHour not implemented")
}
func (UnimplementedMonitizationServer) mustEmbedUnimplementedMonitizationServer() {}

// UnsafeMonitizationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MonitizationServer will
// result in compilation errors.
type UnsafeMonitizationServer interface {
	mustEmbedUnimplementedMonitizationServer()
}

func RegisterMonitizationServer(s grpc.ServiceRegistrar, srv MonitizationServer) {
	s.RegisterService(&Monitization_ServiceDesc, srv)
}

func _Monitization_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitizationServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Monitization_HealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitizationServer).HealthCheck(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitization_CreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitizationServer).CreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Monitization_CreateWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitizationServer).CreateWallet(ctx, req.(*CreateWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitization_GetWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitizationServer).GetWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Monitization_GetWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitizationServer).GetWallet(ctx, req.(*GetWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitization_UpdateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitizationServer).UpdateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Monitization_UpdateWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitizationServer).UpdateWallet(ctx, req.(*UpdateWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitization_ParticipationReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParticipationRewardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitizationServer).ParticipationReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Monitization_ParticipationReward_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitizationServer).ParticipationReward(ctx, req.(*ParticipationRewardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitization_UserRewardHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRewardHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitizationServer).UserRewardHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Monitization_UserRewardHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitizationServer).UserRewardHistory(ctx, req.(*UserRewardHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitization_VideoReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoRewardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitizationServer).VideoReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Monitization_VideoReward_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitizationServer).VideoReward(ctx, req.(*VideoRewardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitization_ExclusiveContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExclusiveContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitizationServer).ExclusiveContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Monitization_ExclusiveContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitizationServer).ExclusiveContent(ctx, req.(*ExclusiveContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitization_GroupRewardHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupRewardHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitizationServer).GroupRewardHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Monitization_GroupRewardHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitizationServer).GroupRewardHistory(ctx, req.(*GroupRewardHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitization_UserWatchHour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserWatchHourRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitizationServer).UserWatchHour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Monitization_UserWatchHour_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitizationServer).UserWatchHour(ctx, req.(*UserWatchHourRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitization_ConferenceWatchHour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConferenceWatchHourRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitizationServer).ConferenceWatchHour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Monitization_ConferenceWatchHour_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitizationServer).ConferenceWatchHour(ctx, req.(*ConferenceWatchHourRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitization_GroupWatchHour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupWatchHourRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitizationServer).GroupWatchHour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Monitization_GroupWatchHour_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitizationServer).GroupWatchHour(ctx, req.(*GroupWatchHourRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Monitization_ServiceDesc is the grpc.ServiceDesc for Monitization service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Monitization_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "monitization.Monitization",
	HandlerType: (*MonitizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _Monitization_HealthCheck_Handler,
		},
		{
			MethodName: "CreateWallet",
			Handler:    _Monitization_CreateWallet_Handler,
		},
		{
			MethodName: "GetWallet",
			Handler:    _Monitization_GetWallet_Handler,
		},
		{
			MethodName: "UpdateWallet",
			Handler:    _Monitization_UpdateWallet_Handler,
		},
		{
			MethodName: "ParticipationReward",
			Handler:    _Monitization_ParticipationReward_Handler,
		},
		{
			MethodName: "UserRewardHistory",
			Handler:    _Monitization_UserRewardHistory_Handler,
		},
		{
			MethodName: "VideoReward",
			Handler:    _Monitization_VideoReward_Handler,
		},
		{
			MethodName: "ExclusiveContent",
			Handler:    _Monitization_ExclusiveContent_Handler,
		},
		{
			MethodName: "GroupRewardHistory",
			Handler:    _Monitization_GroupRewardHistory_Handler,
		},
		{
			MethodName: "UserWatchHour",
			Handler:    _Monitization_UserWatchHour_Handler,
		},
		{
			MethodName: "ConferenceWatchHour",
			Handler:    _Monitization_ConferenceWatchHour_Handler,
		},
		{
			MethodName: "GroupWatchHour",
			Handler:    _Monitization_GroupWatchHour_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/monit.proto",
}
