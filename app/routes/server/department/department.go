package department

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterRoutes(api huma.API) {
	// 创建部门接口
	huma.Register(api, huma.Operation{
		OperationID: "server-department-create",
		Method:      http.MethodPost,
		Path:        "/server/department/create",
		Summary:     "部门-创建",
		Tags:        []string{"部门"},
	}, handlerCreateLogic)
}
