package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"pelago-card-backend/portal/internal/svc"
)

func healthcheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpx.OkJsonCtx(r.Context(), w, "ok")
	}
}
