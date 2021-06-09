package handler

import (
	"net/http"

	"book/service/search/api/internal/logic"
	"book/service/search/api/internal/svc"
	"book/service/search/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func searchHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSearchLogic(r.Context(), ctx)
		resp, err := l.Search(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}