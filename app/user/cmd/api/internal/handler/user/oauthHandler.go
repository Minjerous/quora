package user

import (
	"net/http"
	"quora/app/user/cmd/api/internal/logic/user"
	"quora/app/user/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func OauthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewOauthLogic(r.Context(), svcCtx)
		resp, err := l.Oauth()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
