package logic

import (
	"context"

	"pelago-card-backend/rpc/card/internal/svc"
	"pelago-card-backend/rpc/pb/card"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMerchantsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListMerchantsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMerchantsLogic {
	return &ListMerchantsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 商户列表查询
func (l *ListMerchantsLogic) ListMerchants(in *card.ListMerchantsReq) (*card.ListMerchantsResp, error) {
	// 设置默认值
	page := in.Page
	if page <= 0 {
		page = 1
	}
	pageSize := in.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100 // 限制最大页面大小
	}

	// 查询总数
	total, err := l.svcCtx.MerchantsModel.Count(l.ctx, in.VerifiedStatus, in.MerchantStatus)
	if err != nil {
		logx.Errorf("查询商户总数失败: %v", err)
		return nil, err
	}

	// 查询列表数据
	merchantList, err := l.svcCtx.MerchantsModel.FindList(l.ctx, page, pageSize, in.VerifiedStatus, in.MerchantStatus)
	if err != nil {
		logx.Errorf("查询商户列表失败: %v", err)
		return nil, err
	}

	// 转换为响应格式
	var merchants []*card.MerchantInfo
	for _, merchant := range merchantList {
		merchants = append(merchants, &card.MerchantInfo{
			MerchantId:          merchant.MerchantId,
			EntityName:          merchant.EntityName,
			BrandName:           merchant.BrandName,
			WebsiteUrl:          merchant.WebsiteUrl.String,
			MerchantLogo:        merchant.MerchantLogo.String,
			MerchantCountry:     merchant.MerchantCountry,
			ContactName:         merchant.ContactName,
			ContactEmail:        merchant.ContactEmail,
			ApiKey:              merchant.ApiKey,
			VerifiedStatus:      merchant.VerifiedStatus,
			MerchantStatus:      merchant.MerchantStatus,
			BusinessDescription: merchant.BusinessDescription.String,
			CreatedTime:         merchant.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedTime:         merchant.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &card.ListMerchantsResp{
		Merchants: merchants,
		Total:     int32(total),
		Page:      page,
		PageSize:  pageSize,
	}, nil
}
