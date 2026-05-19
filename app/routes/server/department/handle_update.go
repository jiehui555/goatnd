package department

import (
	"context"
	"fmt"

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
		return nil, fmt.Errorf("未找到部门")
	}

	// 检查名称唯一性 (如果名称改变了)
	if dept.Name != i.Body.Name {
		var count int64
		database.Model(&models.Department{}).Where("name = ? AND id != ?", i.Body.Name, i.ID).Count(&count)
		if count > 0 {
			return nil, fmt.Errorf("部门名称已存在")
		}
		dept.Name = i.Body.Name
	}

	if err := database.Save(&dept).Error; err != nil {
		return nil, err
	}

	resp := &updateOutput{}
	resp.Body.Message = "修改成功"
	return resp, nil
}
