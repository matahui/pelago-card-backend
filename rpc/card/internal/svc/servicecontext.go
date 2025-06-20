package svc

import (
	"pelago-card-backend/rpc/card/internal/config"
	"pelago-card-backend/rpc/card/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	MerchantsModel model.MerchantsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:         c,
		MerchantsModel: model.NewMerchantsModel(conn),
	}
}
