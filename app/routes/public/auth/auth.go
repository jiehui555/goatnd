package auth

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type LoginInput struct {
	Body struct {
		Email    string `json:"email" minLength:"1" doc:"邮箱"`
		Password string `json:"password" minLength:"8" maxLength:"32" doc:"密码"`
	}
}

type LoginOutput struct {
	Body struct {
		Message string `json:"message" example:"OK" doc:"提示信息"`
	}
}

type RefreshInput struct {
	Auth string `header:"Authorization" required:"true"`
}

type RefreshOutput struct {
	Body struct {
		Message string `json:"message" example:"OK" doc:"提示信息"`
	}
}

func RegisterAuthRoutes(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "auth-login",
		Method:      http.MethodPost,
		Path:        "/auth/login",
		Summary:     "登录",
		Tags:        []string{"公共"},
	}, func(ctx context.Context, i *LoginInput) (*LoginOutput, error) {
		resp := &LoginOutput{}
		resp.Body.Message = "Hello!"
		return resp, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "auth-refresh",
		Method:      http.MethodPost,
		Path:        "/auth/refresh",
		Summary:     "刷新 token",
		Tags:        []string{"公共"},
	}, func(ctx context.Context, i *RefreshInput) (*RefreshOutput, error) {
		resp := &RefreshOutput{}
		resp.Body.Message = "OK"
		return resp, nil
	})
}
