package comment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"quora/app/comment/cmd/api/internal/logic/comment"
	"quora/app/comment/cmd/api/internal/svc"
	"quora/app/comment/cmd/api/internal/types"
)

func GetChildCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetChildCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := comment.NewGetChildCommentLogic(r.Context(), svcCtx)
		resp, err := l.GetChildComment(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
