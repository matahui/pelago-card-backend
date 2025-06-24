package logic

import (
	"context"

	"pelago-card-backend/app/cardholder/internal/svc"
	"pelago-card-backend/app/pb/cardholder"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAccountLogic {
	return &CreateAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateAccountLogic) CreateAccount(in *cardholder.CreateAccountReq) (*cardholder.CreateAccountResp, error) {
	// todo: add your logic here and delete this line

	return &cardholder.CreateAccountResp{}, nil
}
