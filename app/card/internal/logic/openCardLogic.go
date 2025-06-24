package logic

import (
	"context"

	"pelago-card-backend/app/card/internal/svc"
	"pelago-card-backend/app/pb/card"

	"github.com/zeromicro/go-zero/core/logx"
)

type OpenCardLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOpenCardLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenCardLogic {
	return &OpenCardLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OpenCardLogic) OpenCard(in *card.OpenCardReq) (*card.OpenCardResp, error) {
	// todo: add your logic here and delete this line

	return &card.OpenCardResp{}, nil
}
