package handler

import (
	"net/http"

	"foxit-otel-go/api/internal/logic"
	"foxit-otel-go/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DemoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDemoLogic(r.Context(), svcCtx)
		resp, err := l.Demo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
