package auth

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

// RegisterRoutes 注册认证相关的公共路由
func RegisterRoutes(api huma.API) {
	// 登录接口
	huma.Register(api, huma.Operation{
		OperationID: "auth-login",
		Method:      http.MethodPost,
		Path:        "/auth/login",
		Summary:     "获取 token",
		Tags:        []string{"公共"},
	}, handlerLoginLogic)

	// 刷新 Token 接口
	huma.Register(api, huma.Operation{
		OperationID: "auth-refresh",
		Method:      http.MethodPost,
		Path:        "/auth/refresh",
		Summary:     "刷新 token",
		Tags:        []string{"公共"},
	}, handleRefreshTokenLogic)

	// 修改密码接口
	huma.Register(api, huma.Operation{
		OperationID: "auth-change-password",
		Method:      http.MethodPost,
		Path:        "/auth/change-password",
		Summary:     "修改密码",
		Tags:        []string{"公共"},
	}, handleChangePasswordLogic)
}
