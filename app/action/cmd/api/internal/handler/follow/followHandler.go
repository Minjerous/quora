package follow

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"quora/app/action/cmd/api/internal/logic/follow"
	"quora/app/action/cmd/api/internal/svc"
	"quora/app/action/cmd/api/internal/types"
)

func FollowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FollowReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := follow.NewFollowLogic(r.Context(), svcCtx)
		resp, err := l.Follow(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
