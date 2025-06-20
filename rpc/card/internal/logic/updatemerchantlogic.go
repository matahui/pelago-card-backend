package logic

import (
	"context"
	"database/sql"
	"time"

	"pelago-card-backend/rpc/card/internal/svc"
	"pelago-card-backend/rpc/pb/card"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMerchantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMerchantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMerchantLogic {
	return &UpdateMerchantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新商户信息
func (l *UpdateMerchantLogic) UpdateMerchant(in *card.UpdateMerchantReq) (*card.UpdateMerchantResp, error) {
	// 验证商户ID是否存在
	merchant, err := l.svcCtx.MerchantsModel.FindOneByMerchantId(l.ctx, in.MerchantId)
	if err != nil {
		if err == sql.ErrNoRows {
			logx.Errorf("商户不存在: %s", in.MerchantId)
			return nil, err
		}
		logx.Errorf("查询商户失败: %v", err)
		return nil, err
	}

	// 更新商户信息
	merchant.EntityName = in.EntityName
	merchant.BrandName = in.BrandName
	merchant.WebsiteUrl = sql.NullString{String: in.WebsiteUrl, Valid: in.WebsiteUrl != ""}
	merchant.MerchantLogo = sql.NullString{String: in.MerchantLogo, Valid: in.MerchantLogo != ""}
	merchant.MerchantCountry = in.MerchantCountry
	merchant.ContactName = in.ContactName
	merchant.ContactEmail = in.ContactEmail
	merchant.BusinessDescription = sql.NullString{String: in.BusinessDescription, Valid: in.BusinessDescription != ""}
	merchant.UpdatedAt = time.Now()

	err = l.svcCtx.MerchantsModel.Update(l.ctx, merchant)
	if err != nil {
		logx.Errorf("更新商户失败: %v", err)
		return nil, err
	}

	// 返回更新后的商户信息
	return &card.UpdateMerchantResp{
		Merchant: &card.MerchantInfo{
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
		},
	}, nil
}
