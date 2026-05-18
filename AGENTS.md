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

## 代码提交规范

请遵循 `类型: 具体内容` 的格式并使用中文进行提交：

- 功能: 新功能
- 修复: 修补 bug
- 文档: 文档修改
- 格式: 代码格式修改 (不影响代码运行的变动)
- 重构: 代码重构
- 性能: 性能优化
- 测试: 测试相关
- 构建: 影响构建系统或外部依赖的更改
- CI: CI 配置文件和脚本的更改
- 杂务: 其他日常事务 (不修改源代码或测试文件)
