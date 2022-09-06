package like

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"quora/app/action/cmd/api/internal/logic/like"
	"quora/app/action/cmd/api/internal/svc"
	"quora/app/action/cmd/api/internal/types"
)

func LikeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LikeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := like.NewLikeLogic(r.Context(), svcCtx)
		resp, err := l.Like(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
