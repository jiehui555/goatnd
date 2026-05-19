package department

import (
	"context"
	"fmt"

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
		return nil, fmt.Errorf("未找到部门")
	}

	if err := database.Delete(&dept).Error; err != nil {
		return nil, err
	}

	resp := &deleteOutput{}
	resp.Body.Message = "删除成功"
	return resp, nil
}
