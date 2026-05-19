package department

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jithui555/goatnd/models"
	"github.com/jithui555/goatnd/pkg/db"
)

type updateInput struct {
	Auth string `header:"Authorization" required:"true"`
	ID   uint   `path:"id" doc:"部门ID"`
	Body struct {
		Name string `json:"name" minLength:"1" maxLength:"32" doc:"名称"`
	}
}

type updateOutput struct {
	Body struct {
		Message string `json:"message" example:"OK" doc:"提示信息"`
	}
}

func handlerUpdateLogic(ctx context.Context, i *updateInput) (*updateOutput, error) {
	database := db.GetDB()

	var dept models.Department
	if err := database.First(&dept, i.ID).Error; err != nil {
		return nil, huma.Error404NotFound("未找到部门")
	}

	// 检查名称唯一性 (如果名称改变了)
	if dept.Name != i.Body.Name {
		var count int64
		database.Model(&models.Department{}).Where("name = ? AND id != ?", i.Body.Name, i.ID).Count(&count)
		if count > 0 {
			return nil, huma.Error400BadRequest("部门名称已存在")
		}
		dept.Name = i.Body.Name
	}

	if err := database.Save(&dept).Error; err != nil {
		return nil, huma.Error500InternalServerError(err.Error())
	}

	resp := &updateOutput{}
	resp.Body.Message = "修改成功"
	return resp, nil
}
