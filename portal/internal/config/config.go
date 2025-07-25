package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	CardRpcConf       zrpc.RpcClientConf
	CardHolderRpcConf zrpc.RpcClientConf
	MerchantRpcConf   zrpc.RpcClientConf
	IntegrationRcConf zrpc.RpcClientConf
	// 其他服务的 rpc 配置...
}
