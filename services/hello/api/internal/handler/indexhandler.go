package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"test/services/hello/api/internal/logic"
	"test/services/hello/api/internal/svc"
	"test/services/hello/api/internal/types"
)

func indexHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetIndexReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewIndexLogic(r.Context(), svcCtx)
		resp, err := l.Index(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
