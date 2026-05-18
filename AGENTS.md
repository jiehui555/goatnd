# Agent 指南

基于 [Huma v2](https://github.com/danielgtaylor/huma) 和 [Chi v5](https://github.com/go-chi/chi) 的 Go API。

## 运行
- `go run main.go` (默认端口 8888)
- `go run main.go -p <port>`

## 路由
- `GET /greeting/{name}`
- `POST /reviews` (占位符)

## 注意事项
- CLI 由 `humacli` 提供。
- 目前无测试，新增功能请添加单元测试。
