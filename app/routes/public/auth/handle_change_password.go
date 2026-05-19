package auth

import (
	"context"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jithui555/goatnd/models"
	"github.com/jithui555/goatnd/pkg/db"
	"github.com/jithui555/goatnd/pkg/token"
	"golang.org/x/crypto/bcrypt"
)

// changePasswordInput 修改密码请求输入
type changePasswordInput struct {
	Auth string `header:"Authorization" required:"true"`
	Body struct {
		OldPassword string `json:"old_password" minLength:"1" doc:"旧密码"`
		NewPassword string `json:"new_password" minLength:"8" maxLength:"32" doc:"新密码"`
	}
}

// changePasswordOutput 修改密码响应输出
type changePasswordOutput struct {
	Body struct {
		Message string `json:"message" example:"OK" doc:"提示信息"`
	}
}

// handleChangePasswordLogic 处理更改密码的逻辑
func handleChangePasswordLogic(ctx context.Context, i *changePasswordInput) (*changePasswordOutput, error) {
	authHeader := i.Auth
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, huma.Error400BadRequest("无效的授权格式")
	}
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := token.ValidateToken(tokenStr)
	if err != nil {
		return nil, huma.Error401Unauthorized("无效的 token")
	}

	database := db.GetDB()
	var user models.User
	if err := database.First(&user, claims.UserID).Error; err != nil {
		return nil, huma.Error401Unauthorized("用户不存在")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(i.Body.OldPassword)); err != nil {
		return nil, huma.Error401Unauthorized("旧密码错误")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(i.Body.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, huma.Error500InternalServerError("生成新密码哈希失败")
	}

	if err := database.Model(&user).Update("password", string(hashedPassword)).Error; err != nil {
		return nil, huma.Error500InternalServerError("更新密码失败")
	}

	resp := &changePasswordOutput{}
	resp.Body.Message = "修改密码成功"
	return resp, nil
}
