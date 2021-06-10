package handler

import (
	"net/http"

	"book/service/user/api/internal/logic"
	"book/service/user/api/internal/svc"
	"book/service/user/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func changepwdHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChangePwdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewChangepwdLogic(r.Context(), ctx)
		err := l.Changepwd(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
