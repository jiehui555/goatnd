package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jithui555/goatnd/models"
	"github.com/jithui555/goatnd/pkg/token"
	"github.com/jithui555/goatnd/pkgs/db"
	"golang.org/x/crypto/bcrypt"
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
		Token   string `json:"token" doc:"JWT Token"`
	}
}

type RefreshInput struct {
	Auth string `header:"Authorization" required:"true"`
}

type RefreshOutput struct {
	Body struct {
		Message string `json:"message" example:"OK" doc:"提示信息"`
		Token   string `json:"token" doc:"JWT Token"`
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
		var user models.User
		database := db.GetDB()
		if err := database.Where("email = ?", i.Body.Email).First(&user).Error; err != nil {
			return nil, huma.Error401Unauthorized("用户不存在或密码错误")
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(i.Body.Password)); err != nil {
			return nil, huma.Error401Unauthorized("用户不存在或密码错误")
		}

		t, err := token.GenerateToken(user.ID)
		if err != nil {
			return nil, huma.Error500InternalServerError("生成 token 失败")
		}

		resp := &LoginOutput{}
		resp.Body.Message = "登录成功"
		resp.Body.Token = t
		return resp, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "auth-refresh",
		Method:      http.MethodPost,
		Path:        "/auth/refresh",
		Summary:     "刷新 token",
		Tags:        []string{"公共"},
	}, func(ctx context.Context, i *RefreshInput) (*RefreshOutput, error) {
		authHeader := i.Auth
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return nil, huma.Error400BadRequest("无效的授权格式")
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := token.ValidateToken(tokenStr)
		if err != nil {
			return nil, huma.Error401Unauthorized("无效的 token")
		}

		// Verify user still exists
		database := db.GetDB()
		var user models.User
		if err := database.First(&user, claims.UserID).Error; err != nil {
			return nil, huma.Error401Unauthorized("用户不存在")
		}

		newToken, err := token.GenerateToken(user.ID)
		if err != nil {
			return nil, huma.Error500InternalServerError("刷新 token 失败")
		}

		resp := &RefreshOutput{}
		resp.Body.Message = "刷新成功"
		resp.Body.Token = newToken
		return resp, nil
	})
}
