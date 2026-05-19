package department

import (
	"context"

	"github.com/jithui555/goatnd/models"
	"github.com/jithui555/goatnd/pkg/db"
)

type searchInput struct {
	Auth string `header:"Authorization" required:"true"`
	Name string `query:"name" doc:"搜索名称"`
}

type searchOutput struct {
	Body struct {
		Items []models.Department `json:"items" doc:"部门列表"`
	}
}

func handleSearchLogic(ctx context.Context, i *searchInput) (*searchOutput, error) {
	database := db.GetDB()

	var depts []models.Department
	query := database.Model(&models.Department{})

	if i.Name != "" {
		query = query.Where("name LIKE ?", "%"+i.Name+"%")
	}

	if err := query.Find(&depts).Error; err != nil {
		return nil, err
	}

	resp := &searchOutput{}
	resp.Body.Items = depts
	return resp, nil
}
