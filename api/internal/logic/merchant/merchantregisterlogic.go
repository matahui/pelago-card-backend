package merchant

import (
	"context"

	"pelago-card-backend/api/internal/svc"
	"pelago-card-backend/api/internal/types"
	"pelago-card-backend/rpc/pb/card"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 商户入驻注册
func NewMerchantRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantRegisterLogic {
	return &MerchantRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantRegisterLogic) MerchantRegister(req *types.MerchantRegisterReq) (resp *types.MerchantRegisterResp, err error) {
	// 调用RPC服务注册商户
	registerResp, err := l.svcCtx.CardRpcClient.RegisterMerchant(l.ctx, &card.RegisterMerchantReq{
		EntityName:          req.EntityName,
		BrandName:           req.BrandName,
		WebsiteUrl:          req.WebsiteUrl,
		MerchantLogo:        req.MerchantLogo,
		MerchantCountry:     req.MerchantCountry,
		ContactName:         req.ContactName,
		ContactEmail:        req.ContactEmail,
		BusinessDescription: req.BusinessDescription,
	})
	if err != nil {
		logx.Errorf("商户注册失败: %v", err)
		return nil, err
	}

	// 转换响应格式
	return &types.MerchantRegisterResp{
		MerchantId:     registerResp.MerchantId,
		ApiKey:         registerResp.ApiKey,
		VerifiedStatus: registerResp.VerifiedStatus,
		MerchantStatus: registerResp.MerchantStatus,
		CreatedTime:    registerResp.CreatedTime,
	}, nil
}
