package comment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"quora/app/comment/cmd/api/internal/logic/comment"
	"quora/app/comment/cmd/api/internal/svc"
	"quora/app/comment/cmd/api/internal/types"
)

func DeleteCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteCommmentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := comment.NewDeleteCommentLogic(r.Context(), svcCtx)
		resp, err := l.DeleteComment(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
