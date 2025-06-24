package merchant

import (
	"net/http"
	"pelago-card-backend/app/pb/merchant"

	"github.com/zeromicro/go-zero/rest/httpx"
	"pelago-card-backend/portal/internal/svc"
	"pelago-card-backend/portal/internal/types"
)

// merchant register
func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MerchantRegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		//call rpc
		rpcReq := &merchant.RegisterMerchantReq{
			// TODO:
		}

		resp, err := svcCtx.MerchantRpc.RegisterMerchant(r.Context(), rpcReq)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
