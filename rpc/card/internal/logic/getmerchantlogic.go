package logic

import (
	"context"
	"fmt"

	"pelago-card-backend/rpc/card/internal/svc"
	"pelago-card-backend/rpc/pb/card"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type GetMerchantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMerchantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMerchantLogic {
	return &GetMerchantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取商户信息
func (l *GetMerchantLogic) GetMerchant(in *card.GetMerchantReq) (*card.GetMerchantResp, error) {
	// 从数据库获取商户信息
	merchant, err := l.svcCtx.MerchantsModel.FindOneByMerchantId(l.ctx, in.MerchantId)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, fmt.Errorf("商户不存在: %s", in.MerchantId)
		}
		l.Logger.Errorf("Failed to get merchant: %v", err)
		return nil, fmt.Errorf("获取商户信息失败: %v", err)
	}

	// 构建响应数据
	merchantInfo := &card.MerchantInfo{
		MerchantId:      merchant.MerchantId,
		EntityName:      merchant.EntityName,
		BrandName:       merchant.BrandName,
		MerchantCountry: merchant.MerchantCountry,
		ContactName:     merchant.ContactName,
		ContactEmail:    merchant.ContactEmail,
		VerifiedStatus:  merchant.VerifiedStatus,
		MerchantStatus:  merchant.MerchantStatus,
		ApiKey:          merchant.ApiKey,
		CreatedTime:     merchant.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedTime:     merchant.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	// 处理可空字段
	if merchant.WebsiteUrl.Valid {
		merchantInfo.WebsiteUrl = merchant.WebsiteUrl.String
	}
	if merchant.MerchantLogo.Valid {
		merchantInfo.MerchantLogo = merchant.MerchantLogo.String
	}
	if merchant.BusinessDescription.Valid {
		merchantInfo.BusinessDescription = merchant.BusinessDescription.String
	}

	return &card.GetMerchantResp{
		Merchant: merchantInfo,
	}, nil
}
