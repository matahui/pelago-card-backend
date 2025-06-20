package merchant

import (
	"context"

	"pelago-card-backend/api/internal/svc"
	"pelago-card-backend/api/internal/types"
	"pelago-card-backend/rpc/pb/card"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMerchantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取商户信息
func NewGetMerchantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMerchantLogic {
	return &GetMerchantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMerchantLogic) GetMerchant(req *types.GetMerchantReq) (resp *types.GetMerchantResp, err error) {
	// 调用RPC服务获取商户信息
	getMerchantResp, err := l.svcCtx.CardRpcClient.GetMerchant(l.ctx, &card.GetMerchantReq{
		MerchantId: req.MerchantId,
	})
	if err != nil {
		logx.Errorf("获取商户信息失败: %v", err)
		return nil, err
	}

	// 转换响应格式
	return &types.GetMerchantResp{
		Merchant: types.MerchantInfo{
			MerchantId:          getMerchantResp.Merchant.MerchantId,
			EntityName:          getMerchantResp.Merchant.EntityName,
			BrandName:           getMerchantResp.Merchant.BrandName,
			WebsiteUrl:          getMerchantResp.Merchant.WebsiteUrl,
			MerchantLogo:        getMerchantResp.Merchant.MerchantLogo,
			MerchantCountry:     getMerchantResp.Merchant.MerchantCountry,
			ContactName:         getMerchantResp.Merchant.ContactName,
			ContactEmail:        getMerchantResp.Merchant.ContactEmail,
			ApiKey:              getMerchantResp.Merchant.ApiKey,
			VerifiedStatus:      getMerchantResp.Merchant.VerifiedStatus,
			MerchantStatus:      getMerchantResp.Merchant.MerchantStatus,
			BusinessDescription: getMerchantResp.Merchant.BusinessDescription,
			CreatedTime:         getMerchantResp.Merchant.CreatedTime,
			UpdatedTime:         getMerchantResp.Merchant.UpdatedTime,
		},
	}, nil
}
