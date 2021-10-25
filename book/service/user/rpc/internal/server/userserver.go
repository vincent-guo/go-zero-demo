// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"go-zero-demo/book/service/user/rpc/internal/logic"
	"go-zero-demo/book/service/user/rpc/internal/svc"
	"go-zero-demo/book/service/user/rpc/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetUser(ctx context.Context, in *user.IdReq) (*user.UserInfoReply, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}
