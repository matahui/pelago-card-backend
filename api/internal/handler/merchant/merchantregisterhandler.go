package merchant

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"pelago-card-backend/api/internal/logic/merchant"
	"pelago-card-backend/api/internal/svc"
	"pelago-card-backend/api/internal/types"
)

// 商户入驻注册
func MerchantRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MerchantRegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := merchant.NewMerchantRegisterLogic(r.Context(), svcCtx)
		resp, err := l.MerchantRegister(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
