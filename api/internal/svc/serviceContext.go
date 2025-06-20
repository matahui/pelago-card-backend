package svc

import (
	"pelago-card-backend/api/internal/config"
	"pelago-card-backend/rpc/pb/card"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	CardRpcClient card.CardServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		CardRpcClient: card.NewCardServiceClient(zrpc.MustNewClient(c.CardRpc).Conn()),
	}
}
