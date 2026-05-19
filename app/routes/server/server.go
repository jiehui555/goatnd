package server

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/jithui555/goatnd/app/routes/server/department"
)

func RegisterRoutes(api huma.API) {
	department.RegisterRoutes(api)
}
