package cardholder

import (
	"net/http"
	"pelago-card-backend/app/pb/cardholder"

	"github.com/zeromicro/go-zero/rest/httpx"
	"pelago-card-backend/portal/internal/svc"
	"pelago-card-backend/portal/internal/types"
)

// cardholder register
func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CardholderRegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		//call rpc
		rpcReq := &cardholder.CreateAccountReq{
			// TODO:
		}

		resp, err := svcCtx.CardHodlerRpc.CreateAccount(r.Context(), rpcReq)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
