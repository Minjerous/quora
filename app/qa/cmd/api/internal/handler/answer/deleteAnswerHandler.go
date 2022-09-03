package answer

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"quora/app/qa/cmd/api/internal/logic/answer"
	"quora/app/qa/cmd/api/internal/svc"
	"quora/app/qa/cmd/api/internal/types"
)

func DeleteAnswerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteAnswerReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := answer.NewDeleteAnswerLogic(r.Context(), svcCtx)
		resp, err := l.DeleteAnswer(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
