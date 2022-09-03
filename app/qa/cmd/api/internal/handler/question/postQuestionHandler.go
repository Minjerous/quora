package question

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"quora/app/qa/cmd/api/internal/logic/question"
	"quora/app/qa/cmd/api/internal/svc"
	"quora/app/qa/cmd/api/internal/types"
)

func PostQuestionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostQuestionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := question.NewPostQuestionLogic(r.Context(), svcCtx)
		resp, err := l.PostQuestion(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
