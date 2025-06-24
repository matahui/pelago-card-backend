package card

import (
	"net/http"

	"pelago-card-backend/app/pb/card"
	"pelago-card-backend/portal/internal/svc"
	"pelago-card-backend/portal/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// open card
func OpenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OpenCardReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		//call rpc
		rpcReq := &card.OpenCardReq{
			// TODO:
		}

		resp, err := svcCtx.CardRpc.OpenCard(r.Context(), rpcReq)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
