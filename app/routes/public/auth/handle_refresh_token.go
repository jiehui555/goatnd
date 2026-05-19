package auth

import (
	"context"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jithui555/goatnd/models"
	"github.com/jithui555/goatnd/pkg/db"
	"github.com/jithui555/goatnd/pkg/token"
)

// refreshInput 刷新 Token 请求输入
type refreshInput struct {
	Auth string `header:"Authorization" required:"true"`
}

// refreshOutput 刷新 Token 响应输出
type refreshOutput struct {
	Body struct {
		Message string `json:"message" example:"OK" doc:"提示信息"`
		Token   string `json:"token" doc:"JWT Token"`
	}
}

// handleRefreshTokenLogic 处理刷新 Token 的逻辑
func handleRefreshTokenLogic(ctx context.Context, i *refreshInput) (*refreshOutput, error) {
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

	newToken, err := token.GenerateToken(user.ID, user.Username, user.Email, user.IsAdmin)
	if err != nil {
		return nil, huma.Error500InternalServerError("刷新 token 失败")
	}

	resp := &refreshOutput{}
	resp.Body.Message = "刷新成功"
	resp.Body.Token = newToken
	return resp, nil
}
