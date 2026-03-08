# 技术博客平台与抖音集成项目

## 一、项目介绍

本项目是一个集成了技术博客平台和短视频平台(抖音)的微服务后端系统。项目基于Go语言开发,采用高性能的微服务架构,支持高并发访问。

### 核心特点

- ✅ **双功能合一**: 同时支持技术博客和短视频两大核心功能
- ✅ **统一用户系统**: 博客和抖音共享同一套用户认证体系
- ✅ **高性能架构**: 基于gRPC + Gin的微服务架构
- ✅ **数据持久化**: MySQL + Redis双层存储,保证数据安全和访问性能
- ✅ **模块化设计**: 服务分层清晰,易于扩展和维护

## 二、功能模块

### 2.1 抖音功能模块(原有)
- 用户注册/登录
- 视频发布/浏览
- 点赞/取消赞
- 评论功能
- 关注/粉丝系统
- 私信聊天

### 2.2 博客功能模块(新增)
- ✅ 博客文章发布
- ✅ 博客列表浏览
- ✅ 博客详情查看
- ✅ 博客更新/删除
- ✅ 博客分类管理
- ✅ 博客标签管理
- ⏳ 博客评论系统(复用抖音评论)
- ⏳ 博客点赞系统(复用抖音点赞)
- ⏳ 博客搜索功能
- ⏳ 博客推荐算法

## 三、技术架构

### 3.1 架构图

```
┌─────────────────────────────────────────────────┐
│              客户端 (Web/移动端)                 │
└────────────────────┬────────────────────────────┘
                     │ HTTP/HTTPS
┌────────────────────▼────────────────────────────┐
│         API网关层 (douyin-api)                   │
│   - 路由分发                                     │
│   - 用户认证 (JWT)                               │
│   - 权限验证                                     │
│   - 请求限流                                     │
└────────────────────┬────────────────────────────┘
         │           │           │
         │ HTTP     │ gRPC     │ gRPC
         │          │          │
┌────────▼───┐ ┌───▼────┐ ┌───▼────┐
│  用户服务   │ │视频服务 │ │博客服务 │
│(8887端口)  │ │(8888端口)││(8889端口)│
└────────────┘ └─────────┘ └─────────┘
         │           │           │
         └───────────┴───────────┘
                     │
         ┌───────────▼───────────┐
         │   数据存储层           │
         │   - MySQL (主数据)     │
         │   - Redis (缓存)       │
         └───────────────────────┘
```

### 3.2 技术选型

| 技术栈 | 用途 | 版本 |
|--------|------|------|
| Go | 开发语言 | 1.19+ |
| Gin | Web框架 | 1.8+ |
| gRPC | RPC框架 | 1.40+ |
| Gorm | ORM框架 | 1.23+ |
| MySQL | 关系型数据库 | 8.0+ |
| Redis | 缓存数据库 | 6.0+ |
| JWT | 认证授权 | - |
| Zap | 日志框架 | 1.20+ |
| Protobuf | 序列化 | 3.19+ |

## 四、项目结构

```
arrayDanceBackEnd/
├── base-service/               # 业务服务层
│   ├── model/                  # 数据模型
│   │   ├── UserInfo.go         # 用户模型
│   │   ├── video/              # 视频模型
│   │   └── blog/               # 博客模型 ⭐
│   │       └── bloginfo.go     # 博客数据结构
│   ├── proto/                  # Proto定义
│   │   ├── user.proto          # 用户proto
│   │   ├── video.proto         # 视频proto
│   │   └── blog/               # 博客proto ⭐
│   │       ├── blog.proto      # 博客proto定义
│   │       ├── blog.pb.go      # 生成的pb文件
│   │       └── blog_grpc.pb.go # gRPC服务文件
│   ├── handler/                # 业务逻辑
│   │   ├── user.go             # 用户服务
│   │   ├── video.go            # 视频服务
│   │   └── blog/               # 博客服务 ⭐
│   │       └── blog.go         # 博客业务逻辑
│   ├── cmd/                    # 服务启动入口
│   │   ├── main.go             # 主服务
│   │   └── blog/               # 博客服务 ⭐
│   │       └── main.go         # 博客服务启动
│   ├── global/                 # 全局配置
│   │   ├── global.go           # 全局变量
│   │   └── blog/               # 博客配置 ⭐
│   │       └── global.go       # 博客全局配置
│   └── util/                   # 工具类
│       ├── jwtUtil.go          # JWT工具
│       └── blog/               # 博客工具 ⭐
│           └── redisUtil.go    # Redis工具
│
├── douyin-api/                 # API网关层
│   ├── api/                    # API实现
│   │   ├── user.go             # 用户API
│   │   ├── video.go            # 视频API
│   │   ├── comment.go          # 评论API
│   │   ├── favorite.go         # 点赞API
│   │   └── blog/               # 博客API ⭐
│   │       └── blog.go         # 博客HTTP接口
│   ├── router/                 # 路由配置
│   │   ├── user.go             # 用户路由
│   │   ├── comment.go          # 评论路由
│   │   ├── favorite.go         # 点赞路由
│   │   └── blog/               # 博客路由 ⭐
│   │       └── blog.go         # 博客路由配置
│   ├── global/                 # 全局配置
│   │   └── global.go           # 全局变量
│   ├── util/                   # 工具类
│   │   ├── authMiddleware.go   # 认证中间件
│   │   └── grpcConnUtils.go    # gRPC连接工具
│   └── cmd/                    # 启动入口
│       └── main.go             # API服务启动
│
├── social-service/             # 社交服务
│   ├── handler/                # 社交业务逻辑
│   └── proto/                  # 社交proto
│
├── interaction-service/        # 互动服务
│   ├── handler/                # 互动业务逻辑
│   └── dao/                    # 数据访问层
│
├── blog-sql/                   # 博客数据库脚本 ⭐
│   └── init_blog.sql           # 博客表初始化
│
├── BLOG_README.md              # 博客功能说明 ⭐
├── README.md                   # 项目总说明
└── run.sh                      # 启动脚本
```

## 五、快速开始

### 5.1 环境准备

```bash
# 1. 安装Go (1.19+)
# 2. 安装MySQL (8.0+)
# 3. 安装Redis (6.0+)
# 4. 安装protoc (3.19+)
```

### 5.2 数据库初始化

```bash
# 创建数据库
mysql -u root -p
CREATE DATABASE douyin CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

# 初始化抖音表结构 (使用原有SQL)
# 初始化博客表结构
mysql -u root -p douyin < blog-sql/init_blog.sql
```

### 5.3 配置修改

```bash
# 修改数据库连接信息
vim base-service/global/global.go
# 修改: dsn := "root:密码@tcp(localhost:3306)/douyin..."

vim base-service/global/blog/global.go
# 修改: dsn := "root:密码@tcp(localhost:3306)/douyin..."

# 修改Redis连接信息
vim base-service/global/global.go
# 修改: Addr, Password

vim base-service/global/blog/global.go
# 修改: Addr, Password
```

### 5.4 启动服务

```bash
# 方式一: 使用启动脚本
./run.sh

# 方式二: 手动启动

# 1. 启动博客gRPC服务 (端口8889)
cd base-service/cmd/blog
go run main.go

# 2. 启动主服务 (端口8887)
cd base-service/cmd
go run main.go

# 3. 启动API网关 (端口8888)
cd douyin-api/cmd
go run main.go
```

### 5.5 测试API

```bash
# 获取博客列表
curl http://localhost:8888/blog/list

# 注册用户
curl -X POST http://localhost:8888/douyin/user/register/ \
  -d "username=testuser&password=123456"

# 登录获取token
curl -X POST http://localhost:8888/douyin/user/login/ \
  -d "username=testuser&password=123456"

# 发布博客 (需要token)
curl -X POST http://localhost:8888/blog/publish \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "title=我的第一篇博客&content=这是一篇测试博客&tag=技术,Go"
```

## 六、API接口文档

### 6.1 博客接口

#### 发布博客
```
POST /blog/publish
Header: Authorization: Bearer <token>
Body:
  - title: 博客标题
  - content: 博客内容
  - tag: 标签(可选)
  - category_id: 分类ID(可选)
  - excerpt: 摘要(可选,自动生成)
Response:
{
  "status_code": 0,
  "status_msg": "博客发布成功",
  "blog_id": 1
}
```

#### 获取博客列表
```
GET /blog/list
Query:
  - user_id: 用户ID(可选)
  - category_id: 分类ID(可选)
  - tag_id: 标签ID(可选)
  - keyword: 搜索关键词(可选)
  - page: 页码(默认1)
  - page_size: 每页数量(默认10)
Response:
{
  "status_code": 0,
  "status_msg": "success",
  "blog_list": [
    {
      "blog_id": 1,
      "title": "博客标题",
      "excerpt": "博客摘要",
      "author_name": "作者名",
      "views": 100,
      ...
    }
  ]
}
```

#### 获取博客详情
```
GET /blog/detail
Query:
  - blog_id: 博客ID
Response:
{
  "status_code": 0,
  "status_msg": "success",
  "blog_info": {
    "blog_id": 1,
    "title": "博客标题",
    "content": "博客内容",
    "author_name": "作者名",
    "views": 100,
    ...
  }
}
```

#### 更新博客
```
POST /blog/update
Header: Authorization: Bearer <token>
Body:
  - blog_id: 博客ID
  - title: 标题(可选)
  - content: 内容(可选)
  - tag: 标签(可选)
Response:
{
  "status_code": 0,
  "status_msg": "博客更新成功"
}
```

#### 删除博客
```
POST /blog/delete
Header: Authorization: Bearer <token>
Body:
  - blog_id: 博客ID
Response:
{
  "status_code": 0,
  "status_msg": "博客删除成功"
}
```

### 6.2 抖音接口(原有)

参考原有文档: [接口文档地址](https://apifox.com/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50717106)

## 七、数据库设计

### 7.1 博客相关表

#### blogs (博客表)
| 字段 | 类型 | 说明 |
|------|------|------|
| blog_id | BIGINT | 博客ID(主键) |
| user_id | BIGINT | 作者ID |
| title | VARCHAR(255) | 标题 |
| content | TEXT | 内容 |
| excerpt | VARCHAR(500) | 摘要 |
| views | BIGINT | 阅读量 |
| like_count | BIGINT | 点赞数 |
| comment_count | BIGINT | 评论数 |
| tag | VARCHAR(100) | 标签 |
| is_published | BOOLEAN | 是否发布 |
| is_deleted | BOOLEAN | 软删除标志 |
| is_recommended | BOOLEAN | 是否推荐 |
| create_time | DATETIME | 创建时间 |
| update_time | DATETIME | 更新时间 |

#### categories (分类表)
| 字段 | 类型 | 说明 |
|------|------|------|
| category_id | BIGINT | 分类ID(主键) |
| name | VARCHAR(50) | 分类名称 |
| display_order | BIGINT | 显示顺序 |
| article_count | BIGINT | 文章数量 |

#### tags (标签表)
| 字段 | 类型 | 说明 |
|------|------|------|
| tag_id | BIGINT | 标签ID(主键) |
| name | VARCHAR(50) | 标签名称 |
| article_count | BIGINT | 文章数量 |

#### blog_tags (博客标签关联表)
| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT | ID(主键) |
| blog_id | BIGINT | 博客ID |
| tag_id | BIGINT | 标签ID |

#### blog_categories (博客分类关联表)
| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT | ID(主键) |
| blog_id | BIGINT | 博客ID |
| category_id | BIGINT | 分类ID |

## 八、性能优化

### 8.1 已实现的优化

1. **gRPC连接池**: 复用gRPC连接,减少创建销毁开销
2. **Redis缓存**:
   - 博客阅读量缓存
   - 用户信息缓存
   - 点赞数据缓存
3. **数据库索引**: 为常用查询字段添加索引
4. **软删除**: 避免物理删除,提高性能
5. **Keep-Alive**: gRPC连接保活

### 8.2 建议的优化

1. **CDN加速**: 静态资源使用CDN
2. **数据库读写分离**: 主从复制
3. **分库分表**: 数据量大时考虑
4. **消息队列**: 异步处理耗时操作
5. **全文搜索**: 使用Elasticsearch

## 九、开发指南

### 9.1 添加新功能

1. **定义数据模型** (`base-service/model/`)
2. **编写Proto文件** (`base-service/proto/`)
3. **生成gRPC代码**: `protoc --go_out=. --go-grpc_out=. xxx.proto`
4. **实现业务逻辑** (`base-service/handler/`)
5. **编写API接口** (`douyin-api/api/`)
6. **配置路由** (`douyin-api/router/`)
7. **测试验证**

### 9.2 Proto文件生成命令

```bash
# 进入proto目录
cd base-service/proto/blog

# 生成Go代码
protoc --go_out=. --go-grpc_out=. blog.proto

# 或使用完整路径
protoc --proto_path=. \
  --go_out=paths=source_relative:. \
  --go-grpc_out=paths=source_relative:. \
  blog.proto
```

## 十、常见问题

### Q1: Proto文件生成失败?
A: 确保已安装protoc和protoc-gen-go插件:
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Q2: 数据库连接失败?
A: 检查MySQL服务是否启动,以及连接字符串是否正确

### Q3: Redis连接失败?
A: 检查Redis服务是否启动,以及密码和地址配置

### Q4: 端口被占用?
A: 修改对应服务的端口配置,博客服务默认8889,避免冲突

## 十一、项目贡献

欢迎提交Issue和Pull Request!

## 十二、许可证

MIT License

## 十三、联系方式

- 项目地址: https://github.com/xiajia666/arrayDanceBackEnd
- 问题反馈: 提交Issue

---

**注意**: 本项目博客功能为新增模块,原有抖音功能完全保留且不受影响。博客模块的proto文件需要使用protoc工具重新生成完整的pb.go文件后才能编译运行。
