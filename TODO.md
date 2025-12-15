### 可能优先实现的功能：

1. **依赖管理**（add/update/tidy）- Go 开发最常用
2. **跨平台编译** - Go 的核心优势
3. **测试支持**（test/coverage）- 质量保证必需
4. **代码格式化和检查**（fmt/vet/lint）- 代码规范
5. **watch 模式** - 提升开发体验



## 待添加的功能

### 1. **依赖管理增强**

- `gocar add <package>` - 自动执行 `go get` 并整理 `go.mod`
- `gocar update` - 更新所有依赖（`go get -u xxx`）
- `gocar tidy` - 清理未使用的依赖（`go mod tidy`）
- `gocar vendor` - 创建 vendor 目录

### 2. **测试支持**

- `gocar test` - 运行所有测试（`go test xxx`）
- `gocar test --coverage` - 生成覆盖率报告
- `gocar test --bench` - 运行基准测试
- `gocar test <package>` - 测试指定包

### 3. **代码检查和格式化**

- `gocar fmt` - 格式化代码（`gofmt -w .`）
- `gocar lint` - 代码检查（集成 golangci-lint）
- `gocar vet` - 静态分析（`go vet xxx`）
- `gocar check` - 组合 fmt + vet + test

### 4. **文档生成**

- `gocar doc` - 在浏览器中打开文档（`godoc -http=:6060`）
- `gocar doc <package>` - 显示指定包的文档

### 5. **安装命令**

- `gocar install` - 安装二进制到 $GOPATH/bin
- `gocar install --path` - 安装到指定目录

### 6. **跨平台编译**

- `gocar build --target <os>/<arch>` - 交叉编译
- `gocar build --all` - 为常见平台编译（linux/amd64, darwin/amd64, windows/amd64等）

### 7. **初始化模板**

- `gocar init` - 在现有目录初始化项目（不创建新目录）
- 更多模板：`--template web`（Gin/Echo）、`--template cli`（Cobra）、`--template grpc`

### 8. **工作区支持**

- `gocar workspace` - 管理 Go 1.18+ 的工作区（go.work）
- 多模块项目支持

### 9. **工具管理**

- `gocar tools install <tool>` - 安装开发工具（golangci-lint, mockgen等）
- `gocar tools list` - 列出已安装工具

### 10. **性能分析**

- `gocar profile --cpu` - CPU 性能分析
- `gocar profile --mem` - 内存分析
- `gocar profile --trace` - 跟踪分析

### 11. **发布管理**

- `gocar publish` - 发布到 pkg.go.dev
- `gocar tag` - 创建语义化版本标签

### 12. **配置文件支持**

添加 `gocar.toml` 配置文件（类似 Cargo.toml）：

### 13. **生成代码**

- `gocar generate` - 执行 `go generate xxx.`
- 支持代码生成器集成

### 14. **Watch 模式**

- `gocar watch` - 文件变化时自动重新构建和运行（开发模式）

### 15. **容器化支持**

- `gocar docker` - 生成优化的 Dockerfile（多阶段构建）
- `gocar docker --build` - 构建 Docker 镜像

