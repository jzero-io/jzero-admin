package route

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"server/internal/logic/route"
	"server/internal/svc"
	types "server/internal/types/route"
)

func IsRouteExist(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IsRouteExistRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := route.NewIsRouteExist(r.Context(), svcCtx, r)
		resp, err := l.IsRouteExist(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
