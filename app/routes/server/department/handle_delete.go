package department

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jithui555/goatnd/models"
	"github.com/jithui555/goatnd/pkg/db"
)

type deleteInput struct {
	Auth string `header:"Authorization" required:"true"`
	ID   uint   `path:"id" doc:"部门ID"`
}

type deleteOutput struct {
	Body struct {
		Message string `json:"message" example:"OK" doc:"提示信息"`
	}
}

func handleDeleteLogic(ctx context.Context, i *deleteInput) (*deleteOutput, error) {
	database := db.GetDB()

	var dept models.Department
	if err := database.First(&dept, i.ID).Error; err != nil {
		return nil, huma.Error404NotFound("未找到部门")
	}

	if err := database.Delete(&dept).Error; err != nil {
		return nil, huma.Error500InternalServerError(err.Error())
	}

	resp := &deleteOutput{}
	resp.Body.Message = "删除成功"
	return resp, nil
}
