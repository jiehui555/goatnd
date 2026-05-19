package department

import (
	"context"
	"fmt"

	"github.com/jithui555/goatnd/models"
	"github.com/jithui555/goatnd/pkg/db"
)

type createInput struct {
	Body struct {
		Name string `json:"name" minLength:"1" maxLength:"32" doc:"名称"`
	}
}

type createOutput struct {
	Body struct {
		ID      uint   `json:"id" doc:"部门ID"`
		Name    string `json:"name" doc:"部门名称"`
		Message string `json:"message" example:"OK" doc:"提示信息"`
	}
}

func handlerCreateLogic(ctx context.Context, i *createInput) (*createOutput, error) {
	database := db.GetDB()

	var count int64
	database.Model(&models.Department{}).Where("name = ?", i.Body.Name).Count(&count)
	if count > 0 {
		return nil, fmt.Errorf("部门名称已存在")
	}

	dept := models.Department{
		Name: i.Body.Name,
	}

	if err := database.Create(&dept).Error; err != nil {
		return nil, err
	}

	resp := &createOutput{}
	resp.Body.ID = dept.ID
	resp.Body.Name = dept.Name
	resp.Body.Message = "OK"
	return resp, nil
}
