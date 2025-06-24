package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"pelago-card-backend/app/pb/card"
	"pelago-card-backend/app/pb/cardholder"
	"pelago-card-backend/app/pb/integration"
	"pelago-card-backend/app/pb/merchant"
	"pelago-card-backend/portal/internal/config"
	"pelago-card-backend/portal/internal/middleware"
)

type ServiceContext struct {
	Config     config.Config
	Logging    rest.Middleware
	Prometheus rest.Middleware

	CardRpc        card.CardServiceClient
	CardHodlerRpc  cardholder.CardHolderServiceClient
	MerchantRpc    merchant.MerchantServiceClient
	IntegrationRpc integration.IntegrationServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		Logging:        middleware.NewLoggingMiddleware().Handle,
		Prometheus:     middleware.NewPrometheusMiddleware().Handle,
		CardRpc:        card.NewCardServiceClient(zrpc.MustNewClient(c.CardRpcConf).Conn()),
		CardHodlerRpc:  cardholder.NewCardHolderServiceClient(zrpc.MustNewClient(c.CardHolderRpcConf).Conn()),
		MerchantRpc:    merchant.NewMerchantServiceClient(zrpc.MustNewClient(c.MerchantRpcConf).Conn()),
		IntegrationRpc: integration.NewIntegrationServiceClient(zrpc.MustNewClient(c.IntegrationRcConf).Conn()),
	}
}
