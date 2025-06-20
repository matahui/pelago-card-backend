package merchant

import (
	"context"

	"pelago-card-backend/api/internal/svc"
	"pelago-card-backend/api/internal/types"
	"pelago-card-backend/rpc/pb/card"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMerchantsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 商户列表查询
func NewListMerchantsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMerchantsLogic {
	return &ListMerchantsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListMerchantsLogic) ListMerchants(req *types.ListMerchantsReq) (resp *types.ListMerchantsResp, err error) {
	// 调用RPC服务查询商户列表
	listResp, err := l.svcCtx.CardRpcClient.ListMerchants(l.ctx, &card.ListMerchantsReq{
		Page:           int32(req.Page),
		PageSize:       int32(req.PageSize),
		VerifiedStatus: req.VerifiedStatus,
		MerchantStatus: req.MerchantStatus,
	})
	if err != nil {
		logx.Errorf("查询商户列表失败: %v", err)
		return nil, err
	}

	// 转换响应格式
	var merchants []types.MerchantInfo
	for _, merchant := range listResp.Merchants {
		merchants = append(merchants, types.MerchantInfo{
			MerchantId:          merchant.MerchantId,
			EntityName:          merchant.EntityName,
			BrandName:           merchant.BrandName,
			WebsiteUrl:          merchant.WebsiteUrl,
			MerchantLogo:        merchant.MerchantLogo,
			MerchantCountry:     merchant.MerchantCountry,
			ContactName:         merchant.ContactName,
			ContactEmail:        merchant.ContactEmail,
			ApiKey:              merchant.ApiKey,
			VerifiedStatus:      merchant.VerifiedStatus,
			MerchantStatus:      merchant.MerchantStatus,
			BusinessDescription: merchant.BusinessDescription,
			CreatedTime:         merchant.CreatedTime,
			UpdatedTime:         merchant.UpdatedTime,
		})
	}

	return &types.ListMerchantsResp{
		Merchants: merchants,
		Total:     int(listResp.Total),
		Page:      int(listResp.Page),
		PageSize:  int(listResp.PageSize),
	}, nil
}
