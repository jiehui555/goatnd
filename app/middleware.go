package app

import (
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
)

func AuthMiddleware(api huma.API) func(ctx huma.Context, next func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		r, _ := humago.Unwrap(ctx)

		// 获取请求头中的 Token
		token := r.Header.Get("Authorization")
		if token == "" {
			huma.WriteErr(api, ctx, http.StatusInternalServerError,
				"Some friendly message", fmt.Errorf("error detail"),
			)
			return
		}

		next(ctx)
	}
}
