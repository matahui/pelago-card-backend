package merchant

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"pelago-card-backend/api/internal/logic/merchant"
	"pelago-card-backend/api/internal/svc"
	"pelago-card-backend/api/internal/types"
)

// 商户列表查询
func ListMerchantsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListMerchantsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := merchant.NewListMerchantsLogic(r.Context(), svcCtx)
		resp, err := l.ListMerchants(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
