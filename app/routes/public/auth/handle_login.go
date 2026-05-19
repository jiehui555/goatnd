package auth

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jithui555/goatnd/models"
	"github.com/jithui555/goatnd/pkg/db"
	"github.com/jithui555/goatnd/pkg/token"
	"golang.org/x/crypto/bcrypt"
)

// LoginInput 登录请求输入
type LoginInput struct {
	Body struct {
		Email    string `json:"email" minLength:"1" doc:"邮箱"`
		Password string `json:"password" minLength:"8" maxLength:"32" doc:"密码"`
	}
}

// LoginOutput 登录响应输出
type LoginOutput struct {
	Body struct {
		Message string `json:"message" example:"OK" doc:"提示信息"`
		Token   string `json:"token" doc:"JWT Token"`
	}
}

// handlerLoginLogic 处理登录的逻辑
func handlerLoginLogic(ctx context.Context, i *LoginInput) (*LoginOutput, error) {
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
}
