package merchant

import (
	"context"

	"pelago-card-backend/api/internal/svc"
	"pelago-card-backend/api/internal/types"
	"pelago-card-backend/rpc/pb/card"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMerchantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新商户信息
func NewUpdateMerchantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMerchantLogic {
	return &UpdateMerchantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMerchantLogic) UpdateMerchant(req *types.UpdateMerchantReq) (resp *types.UpdateMerchantResp, err error) {
	// 调用RPC服务更新商户信息
	updateResp, err := l.svcCtx.CardRpcClient.UpdateMerchant(l.ctx, &card.UpdateMerchantReq{
		MerchantId:          req.MerchantId,
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
		logx.Errorf("更新商户信息失败: %v", err)
		return nil, err
	}

	// 转换响应格式
	return &types.UpdateMerchantResp{
		Merchant: types.MerchantInfo{
			MerchantId:          updateResp.Merchant.MerchantId,
			EntityName:          updateResp.Merchant.EntityName,
			BrandName:           updateResp.Merchant.BrandName,
			WebsiteUrl:          updateResp.Merchant.WebsiteUrl,
			MerchantLogo:        updateResp.Merchant.MerchantLogo,
			MerchantCountry:     updateResp.Merchant.MerchantCountry,
			ContactName:         updateResp.Merchant.ContactName,
			ContactEmail:        updateResp.Merchant.ContactEmail,
			ApiKey:              updateResp.Merchant.ApiKey,
			VerifiedStatus:      updateResp.Merchant.VerifiedStatus,
			MerchantStatus:      updateResp.Merchant.MerchantStatus,
			BusinessDescription: updateResp.Merchant.BusinessDescription,
			CreatedTime:         updateResp.Merchant.CreatedTime,
			UpdatedTime:         updateResp.Merchant.UpdatedTime,
		},
	}, nil
}
