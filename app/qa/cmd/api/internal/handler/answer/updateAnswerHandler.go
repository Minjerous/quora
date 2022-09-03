package answer

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"quora/app/qa/cmd/api/internal/logic/answer"
	"quora/app/qa/cmd/api/internal/svc"
	"quora/app/qa/cmd/api/internal/types"
)

func UpdateAnswerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateAnswerReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := answer.NewUpdateAnswerLogic(r.Context(), svcCtx)
		resp, err := l.UpdateAnswer(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
