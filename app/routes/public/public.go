package public

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/jithui555/goatnd/app/routes/public/auth"
)

func RegisterRoutes(api huma.API) {
	auth.RegisterRoutes(api)
}
