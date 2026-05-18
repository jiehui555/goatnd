package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/jithui555/goatnd/models"
	"github.com/jithui555/goatnd/pkgs/db"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

// Options CLI 配置选项。
type Options struct {
	Port   int    `help:"监听端口" short:"p" default:"8888"`
	DBPath string `help:"数据库文件路径" short:"d" default:"data.db"`
}

// GreetingOutput 表示问候操作的响应。
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"问候信息"`
	}
}

// ReviewInput 表示评论操作的请求。
type ReviewInput struct {
	Auth string `header:"Authorization" required:"true"`
	Body struct {
		Author  string `json:"author" maxLength:"10" doc:"评论者"`
		Rating  int    `json:"rating" minimum:"1" maximum:"5" doc:"评分 (1-5)"`
		Message string `json:"message,omitempty" maxLength:"100" doc:"评论内容"`
	}
}

func Run() {
	// 创建一个支持端口选项的 CLI 应用。
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		// 创建新的路由器和 API
		router := http.NewServeMux()
		api := humago.New(router, huma.DefaultConfig("我的 API", "1.0.0"))

		// 注册 GET /greeting/{name}
		huma.Register(api, huma.Operation{
			OperationID: "get-greeting",
			Method:      http.MethodGet,
			Path:        "/greeting/{name}",
			Summary:     "获取问候语",
			Description: "根据姓名获取问候语。",
			Tags:        []string{"问候"},
		}, func(ctx context.Context, input *struct {
			Name string `path:"name" maxLength:"30" example:"world" doc:"要问候的名字"`
		}) (*GreetingOutput, error) {
			resp := &GreetingOutput{}
			resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
			return resp, nil
		})

		// 注册 POST /reviews
		huma.Register(api, huma.Operation{
			OperationID:   "post-review",
			Method:        http.MethodPost,
			Path:          "/reviews",
			Summary:       "提交评论",
			Tags:          []string{"评论"},
			DefaultStatus: http.StatusCreated,
			Middlewares:   huma.Middlewares{AuthMiddleware(api)},
		}, func(ctx context.Context, i *ReviewInput) (*struct{}, error) {
			// TODO: 将评论保存到数据存储中。
			return nil, nil
		})

		// 初始化数据库
		database, err := db.InitDB(options.DBPath)
		if err != nil {
			panic(fmt.Sprintf("无法初始化数据库: %v", err))
		}

		// 执行数据库迁移
		err = database.AutoMigrate(&models.User{})
		if err != nil {
			panic(fmt.Sprintf("数据库迁移失败: %v", err))
		}

		// 自动创建默认管理员账号
		var adminCount int64
		database.Model(&models.User{}).Where("email = ?", "admin@example.com").Count(&adminCount)
		if adminCount == 0 {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
			if err != nil {
				panic(fmt.Sprintf("生成管理员密码失败: %v", err))
			}
			adminUser := &models.User{
				Username: "admin",
				Email:    "admin@example.com",
				Password: string(hashedPassword),
				IsAdmin:  true,
			}
			if err := database.Create(adminUser).Error; err != nil {
				panic(fmt.Sprintf("创建管理员账号失败: %v", err))
			}
			fmt.Println("已自动创建默认管理员账号: admin@example.com")
		}

		// 告诉 CLI 如何启动服务器。
		hooks.OnStart(func() {
			fmt.Printf("正在端口 %d 上启动服务器...\n", options.Port)
			fmt.Printf("数据库文件: %s\n", options.DBPath)
			fmt.Printf("文档地址：http://127.0.0.1:%d/docs\n", options.Port)
			http.ListenAndServe(fmt.Sprintf(":%d", options.Port), router)
		})
	})

	// 运行 CLI。如果没有传入命令，则启动服务器。
	cli.Run()
}
