// BLOG_SERVICE_GUIDE.md - 博客服务开发和部署指南

# 技术博客平台后端 - 完整部署指南

## 第一部分: 项目概述

本项目成功在原有抖音微服务架构基础上,新增了完整的技术博客平台功能。项目保持了原有的高性能架构,同时扩展了博客相关的业务能力。

## 第二部分: 已完成的工作

### 2.1 数据库层 ✅
- 创建了博客相关的数据模型 (`base-service/model/blog/bloginfo.go`)
- 设计了完整的数据库表结构 (blogs, categories, tags等)
- 编写了数据库初始化SQL脚本 (`blog-sql/init_blog.sql`)

### 2.2 业务服务层 ✅
- 创建了博客gRPC服务 (`base-service/proto/blog/blog.proto`)
- 实现了博客业务逻辑 (`base-service/handler/blog/blog.go`)
- 配置了博客服务启动入口 (`base-service/cmd/blog/main.go`)
- 添加了Redis缓存工具 (`base-service/util/blog/redisUtil.go`)

### 2.3 API网关层 ✅
- 实现了博客HTTP接口 (`douyin-api/api/blog/blog.go`)
- 配置了博客路由 (`douyin-api/router/blog/blog.go`)
- 更新了主路由配置 (`douyin-api/globalinit/router.go`)

### 2.4 测试和文档 ✅
- 编写了单元测试 (`base-service/test/blog/blog_test.go`)
- 创建了完整的项目文档 (`BLOG_README.md`, `BLOG_PROJECT.md`)

## 第三部分: 需要手动完成的步骤

### 3.1 Proto文件生成 (重要!)

当前的proto相关文件是手动创建的简化版本。在实际运行前,需要使用protoc工具生成完整的代码。

#### 步骤:

```bash
# 1. 安装protoc编译器和Go插件
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 2. 进入proto目录
cd /Users/xiajia/Desktop/arrayDanceBackEnd/base-service/proto/blog

# 3. 备份现有的proto文件(如果需要)
cp blog.proto blog.proto.backup

# 4. 使用protoc生成完整的Go代码
protoc --proto_path=. \
  --go_out=paths=source_relative:. \
  --go-grpc_out=paths=source_relative:. \
  blog.proto

# 5. 生成的文件会覆盖:
#    - blog.pb.go
#    - blog_grpc.pb.go
```

### 3.2 数据库初始化

```bash
# 1. 登录MySQL
mysql -u root -p

# 2. 创建数据库(如果不存在)
CREATE DATABASE IF NOT EXISTS douyin CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

# 3. 执行博客表初始化脚本
USE douyin;
SOURCE /Users/xiajia/Desktop/arrayDanceBackEnd/blog-sql/init_blog.sql;

# 4. 验证表是否创建成功
SHOW TABLES LIKE 'blog%';
```

### 3.3 配置文件修改

#### 修改数据库连接:

文件: `base-service/global/blog/global.go`

```go
// 修改MySQL连接字符串
dsn := "root:你的密码@tcp(localhost:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"

// 修改Redis连接信息
RS = redis.NewClient(&redis.Options{
    Addr:     "你的Redis地址:端口",
    Password: "你的Redis密码",
    DB:       0,
})
```

#### 注意:
- 原有的 `base-service/global/global.go` 和 `douyin-api/global/global.go` 也需要检查配置
- 确保所有服务的数据库和Redis配置一致

### 3.4 端口配置

默认端口分配:
- **API网关**: 8888
- **主业务服务**: 8887
- **博客服务**: 8889 (新增)

如果端口冲突,可以在对应的main.go文件中修改:
```go
Port := flag.Int("port", 8889, "端口号")  // 修改这里的端口号
```

## 第四部分: 启动服务

### 4.1 方式一: 使用启动脚本

```bash
cd /Users/xiajia/Desktop/arrayDanceBackEnd
./run.sh
```

### 4.2 方式二: 手动启动(推荐开发环境)

打开三个终端:

#### 终端1: 启动博客gRPC服务
```bash
cd /Users/xiajia/Desktop/arrayDanceBackEnd/base-service/cmd/blog
go run main.go
# 输出: ip: 127.0.0.1  port: 8889  Blog Service is running
```

#### 终端2: 启动主业务服务
```bash
cd /Users/xiajia/Desktop/arrayDanceBackEnd/base-service/cmd
go run main.go
# 输出: ip: 127.0.0.1  port: 8887  Service is running
```

#### 终端3: 启动API网关
```bash
cd /Users/xiajia/Desktop/arrayDanceBackEnd/douyin-api/cmd
go run main.go
# 输出: API服务启动成功
```

## 第五部分: API测试

### 5.1 使用curl测试

#### 1. 用户注册
```bash
curl -X POST "http://localhost:8888/douyin/user/register/" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "username=testuser&password=123456"
```

#### 2. 用户登录
```bash
curl -X POST "http://localhost:8888/douyin/user/login/" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "username=testuser&password=123456"
```

响应示例:
```json
{
  "status_code": 0,
  "status_msg": "success",
  "user_id": 1,
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### 3. 发布博客
```bash
curl -X POST "http://localhost:8888/blog/publish" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -H "Authorization: Bearer 你的token" \
  -d "title=我的第一篇Go博客&content=# Go语言入门指南%0A%0A这是一篇关于Go语言的博客...&tag=Go,编程,技术"
```

#### 4. 获取博客列表
```bash
curl "http://localhost:8888/blog/list?page=1&page_size=10"
```

#### 5. 获取博客详情
```bash
curl "http://localhost:8888/blog/detail?blog_id=1"
```

#### 6. 获取分类列表
```bash
curl "http://localhost:8888/blog/categories"
```

### 5.2 使用Postman测试

1. 导入API文档
2. 创建环境变量: `baseUrl = http://localhost:8888`
3. 测试各个接口

### 5.3 使用Go测试

```bash
# 运行博客服务测试
cd /Users/xiajia/Desktop/arrayDanceBackEnd/base-service/test/blog
go test -v .

# 运行所有测试
cd /Users/xiajia/Desktop/arrayDanceBackEnd
go test ./... -v
```

## 第六部分: 集成现有抖音功能

### 6.1 用户系统集成

博客平台完全复用抖音的用户系统:
- ✅ 用户注册/登录接口不变
- ✅ JWT认证机制不变
- ✅ 用户信息查询接口不变

### 6.2 互动功能集成(建议)

可以复用现有的点赞和评论服务:

#### 复用点赞功能:
```go
// 在博客服务中调用interaction-service的点赞服务
// proto文件可以引用: interaction-service/proto/favorite/favorite.proto
```

#### 复用评论功能:
```go
// 在博客服务中调用interaction-service的评论服务
// proto文件可以引用: interaction-service/proto/comment/comment.proto
```

### 6.3 API路由整合

所有API统一通过 `douyin-api` 网关访问:
- 抖音功能: `/douyin/*`
- 博客功能: `/blog/*`

## 第七部分: 后续开发建议

### 7.1 短期优化

1. **Proto文件生成**: 使用protoc生成完整的pb.go文件
2. **错误处理**: 完善各层的错误处理和日志记录
3. **参数验证**: 添加请求参数验证中间件
4. **单元测试**: 补充完整的单元测试

### 7.2 中期功能

1. **博客评论**: 复用现有评论系统
2. **博客点赞**: 复用现有点赞系统
3. **Markdown支持**: 添加Markdown编辑器和渲染
4. **图片上传**: 集成七牛云对象存储
5. **博客搜索**: 实现基于关键词的搜索功能

### 7.3 长期规划

1. **全文搜索**: 集成Elasticsearch
2. **推荐算法**: 基于用户行为的推荐
3. **SEO优化**: 优化搜索引擎收录
4. **统计分析**: 添加访问统计功能
5. **权限管理**: 实现更细粒度的权限控制

## 第八部分: 常见问题排查

### Q1: Proto编译失败
```
错误: "cannot find package ..."
解决: 确保go.mod文件存在且包含所有依赖
go mod tidy
```

### Q2: 数据库连接失败
```
错误: "dial tcp 127.0.0.1:3306: connect: connection refused"
解决:
1. 检查MySQL服务是否启动: mysql.server status
2. 检查端口是否正确: netstat -an | grep 3306
3. 检查密码是否正确
```

### Q3: Redis连接失败
```
错误: "dial tcp: lookup redis: no such host"
解决:
1. 检查Redis服务: redis-cli ping
2. 检查配置文件中的地址和密码
3. 检查防火墙设置
```

### Q4: 端口被占用
```
错误: "bind: address already in use"
解决:
1. 查找占用端口的进程: lsof -i :8889
2. 杀死进程: kill -9 PID
3. 或修改端口号
```

### Q5: gRPC连接失败
```
错误: "connection refused"
解决:
1. 检查对应的gRPC服务是否启动
2. 检查端口号是否正确
3. 检查防火墙设置
```

## 第九部分: 项目文件清单

### 核心文件

```
✅ 数据库层:
   base-service/model/blog/bloginfo.go
   blog-sql/init_blog.sql

✅ Proto层:
   base-service/proto/blog/blog.proto
   base-service/proto/blog/blog.pb.go (需重新生成)
   base-service/proto/blog/blog_grpc.pb.go (需重新生成)

✅ 业务层:
   base-service/handler/blog/blog.go
   base-service/cmd/blog/main.go
   base-service/global/blog/global.go

✅ API层:
   douyin-api/api/blog/blog.go
   douyin-api/router/blog/blog.go
   douyin-api/globalinit/router.go (已更新)

✅ 测试:
   base-service/test/blog/blog_test.go

✅ 文档:
   BLOG_README.md
   BLOG_PROJECT.md
   BLOG_SERVICE_GUIDE.md (本文件)
```

### 需要检查的配置文件

```
⚠️  需要修改:
   base-service/global/global.go (MySQL, Redis配置)
   base-service/global/blog/global.go (MySQL, Redis配置)
   douyin-api/global/global.go (gRPC连接配置)
```

## 第十部分: 总结

### 10.1 项目亮点

1. **架构清晰**: 保持了原有的微服务架构,层次分明
2. **代码复用**: 充分复用了现有的用户系统、认证机制
3. **易于扩展**: 模块化设计,便于后续功能扩展
4. **文档完善**: 提供了详细的开发和部署文档

### 10.2 技术难点

1. **Proto文件生成**: 需要正确配置protoc工具
2. **服务间调用**: 需要正确配置gRPC连接
3. **数据库设计**: 博客数据模型需要合理设计

### 10.3 下一步行动

1. ✅ 使用protoc生成完整的proto代码
2. ✅ 初始化数据库表结构
3. ✅ 修改配置文件
4. ✅ 启动服务并测试
5. ⏳ 复用点赞和评论功能
6. ⏳ 添加更多博客功能

---

**祝开发顺利!** 🚀

如有问题,请查看:
- 完整项目文档: `BLOG_PROJECT.md`
- 快速开始指南: `BLOG_README.md`
- API接口文档: 项目代码中的注释
