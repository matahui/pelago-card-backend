// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4
// Source: cardholder.proto

package server

import (
	"context"

	"pelago-card-backend/app/cardholder/internal/logic"
	"pelago-card-backend/app/cardholder/internal/svc"
	"pelago-card-backend/app/pb/cardholder"
)

type CardHolderServiceServer struct {
	svcCtx *svc.ServiceContext
	cardholder.UnimplementedCardHolderServiceServer
}

func NewCardHolderServiceServer(svcCtx *svc.ServiceContext) *CardHolderServiceServer {
	return &CardHolderServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *CardHolderServiceServer) CreateAccount(ctx context.Context, in *cardholder.CreateAccountReq) (*cardholder.CreateAccountResp, error) {
	l := logic.NewCreateAccountLogic(ctx, s.svcCtx)
	return l.CreateAccount(in)
}
