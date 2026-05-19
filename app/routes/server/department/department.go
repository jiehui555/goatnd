package department

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jithui555/goatnd/app/routes"
)

func RegisterRoutes(api huma.API) {
	// 创建部门接口
	huma.Register(api, huma.Operation{
		OperationID: "server-department-create",
		Method:      http.MethodPost,
		Path:        "/server/department/create",
		Summary:     "部门-创建",
		Tags:        []string{"服务端"},
		Middlewares: huma.Middlewares{routes.AuthMiddleware(api), routes.AdminMiddleware(api)},
	}, handlerCreateLogic)

	// 更新部门接口
	huma.Register(api, huma.Operation{
		OperationID: "server-department-update",
		Method:      http.MethodPut,
		Path:        "/server/department/update/{id}",
		Summary:     "部门-更新",
		Tags:        []string{"服务端"},
		Middlewares: huma.Middlewares{routes.AuthMiddleware(api), routes.AdminMiddleware(api)},
	}, handlerUpdateLogic)

	// 删除部门接口
	huma.Register(api, huma.Operation{
		OperationID: "server-department-delete",
		Method:      http.MethodDelete,
		Path:        "/server/department/delete/{id}",
		Summary:     "部门-删除",
		Tags:        []string{"服务端"},
		Middlewares: huma.Middlewares{routes.AuthMiddleware(api), routes.AdminMiddleware(api)},
	}, handleDeleteLogic)

	// 搜索部门接口
	huma.Register(api, huma.Operation{
		OperationID: "server-department-search",
		Method:      http.MethodGet,
		Path:        "/server/department/search",
		Summary:     "部门-搜索",
		Tags:        []string{"服务端"},
		Middlewares: huma.Middlewares{routes.AuthMiddleware(api), routes.AdminMiddleware(api)},
	}, handleSearchLogic)
}
