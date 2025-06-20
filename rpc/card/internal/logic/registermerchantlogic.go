package logic

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"time"

	"pelago-card-backend/rpc/card/internal/model"
	"pelago-card-backend/rpc/card/internal/svc"
	"pelago-card-backend/rpc/pb/card"

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

// generateMerchantID 生成商户ID格式: M + yyyyMMddHHmm + 3位随机数
func (l *RegisterMerchantLogic) generateMerchantID() string {
	now := time.Now()
	timeStr := now.Format("200601021504") // yyyyMMddHHmm

	// 生成3位随机数
	randBytes := make([]byte, 2)
	rand.Read(randBytes)
	randNum := int(randBytes[0])<<8 + int(randBytes[1])
	randStr := fmt.Sprintf("%03d", randNum%1000)

	return fmt.Sprintf("M%s%s", timeStr, randStr)
}

// generateAPIKey 生成32位随机API密钥
func (l *RegisterMerchantLogic) generateAPIKey() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// 商户入驻
func (l *RegisterMerchantLogic) RegisterMerchant(in *card.RegisterMerchantReq) (*card.RegisterMerchantResp, error) {
	// 生成商户ID和API密钥
	merchantID := l.generateMerchantID()
	apiKey := l.generateAPIKey()

	// 构建商户数据
	merchant := &model.Merchants{
		MerchantId:      merchantID,
		EntityName:      in.EntityName,
		BrandName:       in.BrandName,
		MerchantCountry: in.MerchantCountry,
		ContactName:     in.ContactName,
		ContactEmail:    in.ContactEmail,
		ApiKey:          apiKey,
		VerifiedStatus:  "Pending for approval", // 默认待审核
		MerchantStatus:  "Inactive",             // 默认未激活
	}

	// 处理可选字段
	if in.WebsiteUrl != "" {
		merchant.WebsiteUrl = sql.NullString{String: in.WebsiteUrl, Valid: true}
	}
	if in.MerchantLogo != "" {
		merchant.MerchantLogo = sql.NullString{String: in.MerchantLogo, Valid: true}
	}
	if in.BusinessDescription != "" {
		merchant.BusinessDescription = sql.NullString{String: in.BusinessDescription, Valid: true}
	}

	// 保存到数据库
	_, err := l.svcCtx.MerchantsModel.Insert(l.ctx, merchant)
	if err != nil {
		l.Logger.Errorf("Failed to insert merchant: %v", err)
		return nil, fmt.Errorf("商户入驻失败: %v", err)
	}

	// 返回响应
	return &card.RegisterMerchantResp{
		MerchantId:     merchantID,
		ApiKey:         apiKey,
		VerifiedStatus: "Pending for approval",
		MerchantStatus: "Inactive",
		CreatedTime:    time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}
