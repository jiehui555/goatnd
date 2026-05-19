package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/jithui555/goatnd/pkg/token"
)

func AuthMiddleware(api huma.API) func(ctx huma.Context, next func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		r, _ := humago.Unwrap(ctx)

		// 获取请求头中的 Token
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Missing authorization header", nil)
			return
		}

		// 检查 Bearer 前缀
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Invalid authorization header format", nil)
			return
		}

		claims, err := token.ValidateToken(tokenString)
		if err != nil {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Invalid token", err)
			return
		}

		// 将用户信息存入 Header 中以便后续 Middleware 或 Handler 使用
		r.Header.Add("X-User-ID", fmt.Sprintf("%d", claims.UserID))
		r.Header.Add("X-Is-Admin", fmt.Sprintf("%t", claims.IsAdmin))

		next(ctx)
	}
}

func AdminMiddleware(api huma.API) func(ctx huma.Context, next func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		if ctx.Header("X-Is-Admin") != "true" {
			huma.WriteErr(api, ctx, http.StatusForbidden, "Admin access required", nil)
			return
		}
		next(ctx)
	}
}
