package logic

import (
	"context"

	"pelago-card-backend/app/integration/internal/svc"
	"pelago-card-backend/app/pb/integration"

	"github.com/zeromicro/go-zero/core/logx"
)

type RechargeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRechargeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RechargeLogic {
	return &RechargeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RechargeLogic) Recharge(in *integration.RechargeReq) (*integration.RechargeResp, error) {
	// todo: add your logic here and delete this line

	return &integration.RechargeResp{}, nil
}
