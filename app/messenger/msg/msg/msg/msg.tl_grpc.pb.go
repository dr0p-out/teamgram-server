//
// WARNING! All changes made in this file will be lost!
// Created from 'scheme.tl' by 'mtprotoc'
//
// Copyright (c) 2024-present,  Teamgram Authors.
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: msg.tl.proto

package msg

import (
	context "context"
	mtproto "github.com/teamgram/proto/mtproto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RPCMsg_MsgPushUserMessage_FullMethodName        = "/msg.RPCMsg/msg_pushUserMessage"
	RPCMsg_MsgReadMessageContents_FullMethodName    = "/msg.RPCMsg/msg_readMessageContents"
	RPCMsg_MsgSendMessageV2_FullMethodName          = "/msg.RPCMsg/msg_sendMessageV2"
	RPCMsg_MsgEditMessageV2_FullMethodName          = "/msg.RPCMsg/msg_editMessageV2"
	RPCMsg_MsgDeleteMessages_FullMethodName         = "/msg.RPCMsg/msg_deleteMessages"
	RPCMsg_MsgDeleteHistory_FullMethodName          = "/msg.RPCMsg/msg_deleteHistory"
	RPCMsg_MsgDeletePhoneCallHistory_FullMethodName = "/msg.RPCMsg/msg_deletePhoneCallHistory"
	RPCMsg_MsgDeleteChatHistory_FullMethodName      = "/msg.RPCMsg/msg_deleteChatHistory"
	RPCMsg_MsgReadHistory_FullMethodName            = "/msg.RPCMsg/msg_readHistory"
	RPCMsg_MsgReadHistoryV2_FullMethodName          = "/msg.RPCMsg/msg_readHistoryV2"
	RPCMsg_MsgUpdatePinnedMessage_FullMethodName    = "/msg.RPCMsg/msg_updatePinnedMessage"
	RPCMsg_MsgUnpinAllMessages_FullMethodName       = "/msg.RPCMsg/msg_unpinAllMessages"
)

// RPCMsgClient is the client API for RPCMsg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCMsgClient interface {
	MsgPushUserMessage(ctx context.Context, in *TLMsgPushUserMessage, opts ...grpc.CallOption) (*mtproto.Bool, error)
	MsgReadMessageContents(ctx context.Context, in *TLMsgReadMessageContents, opts ...grpc.CallOption) (*mtproto.Messages_AffectedMessages, error)
	MsgSendMessageV2(ctx context.Context, in *TLMsgSendMessageV2, opts ...grpc.CallOption) (*mtproto.Updates, error)
	MsgEditMessageV2(ctx context.Context, in *TLMsgEditMessageV2, opts ...grpc.CallOption) (*mtproto.Updates, error)
	MsgDeleteMessages(ctx context.Context, in *TLMsgDeleteMessages, opts ...grpc.CallOption) (*mtproto.Messages_AffectedMessages, error)
	MsgDeleteHistory(ctx context.Context, in *TLMsgDeleteHistory, opts ...grpc.CallOption) (*mtproto.Messages_AffectedHistory, error)
	MsgDeletePhoneCallHistory(ctx context.Context, in *TLMsgDeletePhoneCallHistory, opts ...grpc.CallOption) (*mtproto.Messages_AffectedFoundMessages, error)
	MsgDeleteChatHistory(ctx context.Context, in *TLMsgDeleteChatHistory, opts ...grpc.CallOption) (*mtproto.Bool, error)
	MsgReadHistory(ctx context.Context, in *TLMsgReadHistory, opts ...grpc.CallOption) (*mtproto.Messages_AffectedMessages, error)
	MsgReadHistoryV2(ctx context.Context, in *TLMsgReadHistoryV2, opts ...grpc.CallOption) (*mtproto.Messages_AffectedMessages, error)
	MsgUpdatePinnedMessage(ctx context.Context, in *TLMsgUpdatePinnedMessage, opts ...grpc.CallOption) (*mtproto.Updates, error)
	MsgUnpinAllMessages(ctx context.Context, in *TLMsgUnpinAllMessages, opts ...grpc.CallOption) (*mtproto.Messages_AffectedHistory, error)
}

type rPCMsgClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCMsgClient(cc grpc.ClientConnInterface) RPCMsgClient {
	return &rPCMsgClient{cc}
}

func (c *rPCMsgClient) MsgPushUserMessage(ctx context.Context, in *TLMsgPushUserMessage, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, RPCMsg_MsgPushUserMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) MsgReadMessageContents(ctx context.Context, in *TLMsgReadMessageContents, opts ...grpc.CallOption) (*mtproto.Messages_AffectedMessages, error) {
	out := new(mtproto.Messages_AffectedMessages)
	err := c.cc.Invoke(ctx, RPCMsg_MsgReadMessageContents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) MsgSendMessageV2(ctx context.Context, in *TLMsgSendMessageV2, opts ...grpc.CallOption) (*mtproto.Updates, error) {
	out := new(mtproto.Updates)
	err := c.cc.Invoke(ctx, RPCMsg_MsgSendMessageV2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) MsgEditMessageV2(ctx context.Context, in *TLMsgEditMessageV2, opts ...grpc.CallOption) (*mtproto.Updates, error) {
	out := new(mtproto.Updates)
	err := c.cc.Invoke(ctx, RPCMsg_MsgEditMessageV2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) MsgDeleteMessages(ctx context.Context, in *TLMsgDeleteMessages, opts ...grpc.CallOption) (*mtproto.Messages_AffectedMessages, error) {
	out := new(mtproto.Messages_AffectedMessages)
	err := c.cc.Invoke(ctx, RPCMsg_MsgDeleteMessages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) MsgDeleteHistory(ctx context.Context, in *TLMsgDeleteHistory, opts ...grpc.CallOption) (*mtproto.Messages_AffectedHistory, error) {
	out := new(mtproto.Messages_AffectedHistory)
	err := c.cc.Invoke(ctx, RPCMsg_MsgDeleteHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) MsgDeletePhoneCallHistory(ctx context.Context, in *TLMsgDeletePhoneCallHistory, opts ...grpc.CallOption) (*mtproto.Messages_AffectedFoundMessages, error) {
	out := new(mtproto.Messages_AffectedFoundMessages)
	err := c.cc.Invoke(ctx, RPCMsg_MsgDeletePhoneCallHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) MsgDeleteChatHistory(ctx context.Context, in *TLMsgDeleteChatHistory, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, RPCMsg_MsgDeleteChatHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) MsgReadHistory(ctx context.Context, in *TLMsgReadHistory, opts ...grpc.CallOption) (*mtproto.Messages_AffectedMessages, error) {
	out := new(mtproto.Messages_AffectedMessages)
	err := c.cc.Invoke(ctx, RPCMsg_MsgReadHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) MsgReadHistoryV2(ctx context.Context, in *TLMsgReadHistoryV2, opts ...grpc.CallOption) (*mtproto.Messages_AffectedMessages, error) {
	out := new(mtproto.Messages_AffectedMessages)
	err := c.cc.Invoke(ctx, RPCMsg_MsgReadHistoryV2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) MsgUpdatePinnedMessage(ctx context.Context, in *TLMsgUpdatePinnedMessage, opts ...grpc.CallOption) (*mtproto.Updates, error) {
	out := new(mtproto.Updates)
	err := c.cc.Invoke(ctx, RPCMsg_MsgUpdatePinnedMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) MsgUnpinAllMessages(ctx context.Context, in *TLMsgUnpinAllMessages, opts ...grpc.CallOption) (*mtproto.Messages_AffectedHistory, error) {
	out := new(mtproto.Messages_AffectedHistory)
	err := c.cc.Invoke(ctx, RPCMsg_MsgUnpinAllMessages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCMsgServer is the server API for RPCMsg service.
// All implementations should embed UnimplementedRPCMsgServer
// for forward compatibility
type RPCMsgServer interface {
	MsgPushUserMessage(context.Context, *TLMsgPushUserMessage) (*mtproto.Bool, error)
	MsgReadMessageContents(context.Context, *TLMsgReadMessageContents) (*mtproto.Messages_AffectedMessages, error)
	MsgSendMessageV2(context.Context, *TLMsgSendMessageV2) (*mtproto.Updates, error)
	MsgEditMessageV2(context.Context, *TLMsgEditMessageV2) (*mtproto.Updates, error)
	MsgDeleteMessages(context.Context, *TLMsgDeleteMessages) (*mtproto.Messages_AffectedMessages, error)
	MsgDeleteHistory(context.Context, *TLMsgDeleteHistory) (*mtproto.Messages_AffectedHistory, error)
	MsgDeletePhoneCallHistory(context.Context, *TLMsgDeletePhoneCallHistory) (*mtproto.Messages_AffectedFoundMessages, error)
	MsgDeleteChatHistory(context.Context, *TLMsgDeleteChatHistory) (*mtproto.Bool, error)
	MsgReadHistory(context.Context, *TLMsgReadHistory) (*mtproto.Messages_AffectedMessages, error)
	MsgReadHistoryV2(context.Context, *TLMsgReadHistoryV2) (*mtproto.Messages_AffectedMessages, error)
	MsgUpdatePinnedMessage(context.Context, *TLMsgUpdatePinnedMessage) (*mtproto.Updates, error)
	MsgUnpinAllMessages(context.Context, *TLMsgUnpinAllMessages) (*mtproto.Messages_AffectedHistory, error)
}

// UnimplementedRPCMsgServer should be embedded to have forward compatible implementations.
type UnimplementedRPCMsgServer struct {
}

func (UnimplementedRPCMsgServer) MsgPushUserMessage(context.Context, *TLMsgPushUserMessage) (*mtproto.Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgPushUserMessage not implemented")
}
func (UnimplementedRPCMsgServer) MsgReadMessageContents(context.Context, *TLMsgReadMessageContents) (*mtproto.Messages_AffectedMessages, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgReadMessageContents not implemented")
}
func (UnimplementedRPCMsgServer) MsgSendMessageV2(context.Context, *TLMsgSendMessageV2) (*mtproto.Updates, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgSendMessageV2 not implemented")
}
func (UnimplementedRPCMsgServer) MsgEditMessageV2(context.Context, *TLMsgEditMessageV2) (*mtproto.Updates, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgEditMessageV2 not implemented")
}
func (UnimplementedRPCMsgServer) MsgDeleteMessages(context.Context, *TLMsgDeleteMessages) (*mtproto.Messages_AffectedMessages, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgDeleteMessages not implemented")
}
func (UnimplementedRPCMsgServer) MsgDeleteHistory(context.Context, *TLMsgDeleteHistory) (*mtproto.Messages_AffectedHistory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgDeleteHistory not implemented")
}
func (UnimplementedRPCMsgServer) MsgDeletePhoneCallHistory(context.Context, *TLMsgDeletePhoneCallHistory) (*mtproto.Messages_AffectedFoundMessages, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgDeletePhoneCallHistory not implemented")
}
func (UnimplementedRPCMsgServer) MsgDeleteChatHistory(context.Context, *TLMsgDeleteChatHistory) (*mtproto.Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgDeleteChatHistory not implemented")
}
func (UnimplementedRPCMsgServer) MsgReadHistory(context.Context, *TLMsgReadHistory) (*mtproto.Messages_AffectedMessages, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgReadHistory not implemented")
}
func (UnimplementedRPCMsgServer) MsgReadHistoryV2(context.Context, *TLMsgReadHistoryV2) (*mtproto.Messages_AffectedMessages, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgReadHistoryV2 not implemented")
}
func (UnimplementedRPCMsgServer) MsgUpdatePinnedMessage(context.Context, *TLMsgUpdatePinnedMessage) (*mtproto.Updates, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgUpdatePinnedMessage not implemented")
}
func (UnimplementedRPCMsgServer) MsgUnpinAllMessages(context.Context, *TLMsgUnpinAllMessages) (*mtproto.Messages_AffectedHistory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgUnpinAllMessages not implemented")
}

// UnsafeRPCMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCMsgServer will
// result in compilation errors.
type UnsafeRPCMsgServer interface {
	mustEmbedUnimplementedRPCMsgServer()
}

func RegisterRPCMsgServer(s grpc.ServiceRegistrar, srv RPCMsgServer) {
	s.RegisterService(&RPCMsg_ServiceDesc, srv)
}

func _RPCMsg_MsgPushUserMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLMsgPushUserMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).MsgPushUserMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCMsg_MsgPushUserMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).MsgPushUserMessage(ctx, req.(*TLMsgPushUserMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_MsgReadMessageContents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLMsgReadMessageContents)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).MsgReadMessageContents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCMsg_MsgReadMessageContents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).MsgReadMessageContents(ctx, req.(*TLMsgReadMessageContents))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_MsgSendMessageV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLMsgSendMessageV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).MsgSendMessageV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCMsg_MsgSendMessageV2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).MsgSendMessageV2(ctx, req.(*TLMsgSendMessageV2))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_MsgEditMessageV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLMsgEditMessageV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).MsgEditMessageV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCMsg_MsgEditMessageV2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).MsgEditMessageV2(ctx, req.(*TLMsgEditMessageV2))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_MsgDeleteMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLMsgDeleteMessages)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).MsgDeleteMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCMsg_MsgDeleteMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).MsgDeleteMessages(ctx, req.(*TLMsgDeleteMessages))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_MsgDeleteHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLMsgDeleteHistory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).MsgDeleteHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCMsg_MsgDeleteHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).MsgDeleteHistory(ctx, req.(*TLMsgDeleteHistory))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_MsgDeletePhoneCallHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLMsgDeletePhoneCallHistory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).MsgDeletePhoneCallHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCMsg_MsgDeletePhoneCallHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).MsgDeletePhoneCallHistory(ctx, req.(*TLMsgDeletePhoneCallHistory))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_MsgDeleteChatHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLMsgDeleteChatHistory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).MsgDeleteChatHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCMsg_MsgDeleteChatHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).MsgDeleteChatHistory(ctx, req.(*TLMsgDeleteChatHistory))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_MsgReadHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLMsgReadHistory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).MsgReadHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCMsg_MsgReadHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).MsgReadHistory(ctx, req.(*TLMsgReadHistory))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_MsgReadHistoryV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLMsgReadHistoryV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).MsgReadHistoryV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCMsg_MsgReadHistoryV2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).MsgReadHistoryV2(ctx, req.(*TLMsgReadHistoryV2))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_MsgUpdatePinnedMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLMsgUpdatePinnedMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).MsgUpdatePinnedMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCMsg_MsgUpdatePinnedMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).MsgUpdatePinnedMessage(ctx, req.(*TLMsgUpdatePinnedMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_MsgUnpinAllMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLMsgUnpinAllMessages)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).MsgUnpinAllMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCMsg_MsgUnpinAllMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).MsgUnpinAllMessages(ctx, req.(*TLMsgUnpinAllMessages))
	}
	return interceptor(ctx, in, info, handler)
}

// RPCMsg_ServiceDesc is the grpc.ServiceDesc for RPCMsg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPCMsg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "msg.RPCMsg",
	HandlerType: (*RPCMsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "msg_pushUserMessage",
			Handler:    _RPCMsg_MsgPushUserMessage_Handler,
		},
		{
			MethodName: "msg_readMessageContents",
			Handler:    _RPCMsg_MsgReadMessageContents_Handler,
		},
		{
			MethodName: "msg_sendMessageV2",
			Handler:    _RPCMsg_MsgSendMessageV2_Handler,
		},
		{
			MethodName: "msg_editMessageV2",
			Handler:    _RPCMsg_MsgEditMessageV2_Handler,
		},
		{
			MethodName: "msg_deleteMessages",
			Handler:    _RPCMsg_MsgDeleteMessages_Handler,
		},
		{
			MethodName: "msg_deleteHistory",
			Handler:    _RPCMsg_MsgDeleteHistory_Handler,
		},
		{
			MethodName: "msg_deletePhoneCallHistory",
			Handler:    _RPCMsg_MsgDeletePhoneCallHistory_Handler,
		},
		{
			MethodName: "msg_deleteChatHistory",
			Handler:    _RPCMsg_MsgDeleteChatHistory_Handler,
		},
		{
			MethodName: "msg_readHistory",
			Handler:    _RPCMsg_MsgReadHistory_Handler,
		},
		{
			MethodName: "msg_readHistoryV2",
			Handler:    _RPCMsg_MsgReadHistoryV2_Handler,
		},
		{
			MethodName: "msg_updatePinnedMessage",
			Handler:    _RPCMsg_MsgUpdatePinnedMessage_Handler,
		},
		{
			MethodName: "msg_unpinAllMessages",
			Handler:    _RPCMsg_MsgUnpinAllMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "msg.tl.proto",
}
