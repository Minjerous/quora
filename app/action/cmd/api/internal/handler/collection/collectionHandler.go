package collection

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"quora/app/action/cmd/api/internal/logic/collection"
	"quora/app/action/cmd/api/internal/svc"
	"quora/app/action/cmd/api/internal/types"
)

func CollectionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CollectionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := collection.NewCollectionLogic(r.Context(), svcCtx)
		resp, err := l.Collection(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
