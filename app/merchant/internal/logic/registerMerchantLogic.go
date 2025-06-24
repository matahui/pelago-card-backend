package logic

import (
	"context"

	"pelago-card-backend/app/merchant/internal/svc"
	"pelago-card-backend/app/pb/merchant"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterMerchantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterMerchantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterMerchantLogic {
	return &RegisterMerchantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterMerchantLogic) RegisterMerchant(in *merchant.RegisterMerchantReq) (*merchant.RegisterMerchantResp, error) {
	// todo: add your logic here and delete this line

	return &merchant.RegisterMerchantResp{}, nil
}
