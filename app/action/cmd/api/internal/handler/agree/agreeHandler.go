package agree

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"quora/app/action/cmd/api/internal/logic/agree"
	"quora/app/action/cmd/api/internal/svc"
	"quora/app/action/cmd/api/internal/types"
)

func AgreeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgreeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := agree.NewAgreeLogic(r.Context(), svcCtx)
		resp, err := l.Agree(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
