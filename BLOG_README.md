# 技术博客平台后端说明文档

## 项目概述

本项目在原有抖音微服务架构基础上,新增技术博客平台功能,采用Go语言开发,基于gin + gRPC + Gorm + Redis技术栈。

## 技术架构

### 微服务架构
- **API层 (douyin-api)**: 负责HTTP请求处理、鉴权、路由分发
- **业务层 (base-service)**: 提供gRPC服务,处理业务逻辑
- **数据层**: MySQL + Redis,负责数据存储和缓存

### 技术栈
- **Web框架**: gin
- **RPC框架**: gRPC
- **ORM框架**: Gorm
- **数据库**: MySQL 8.0+
- **缓存**: Redis
- **认证**: JWT
- **日志**: zap

## 博客模块实现

### 1. 数据库设计

#### 博客表 (blogs)
```sql
- blog_id: 博客ID(主键)
- user_id: 作者ID
- title: 标题
- content: 内容
- excerpt: 摘要
- views: 阅读量
- like_count: 点赞数
- comment_count: 评论数
- tag: 标签
- is_published: 是否发布
- is_deleted: 软删除标志
- is_recommended: 是否推荐
- create_time: 创建时间
- update_time: 更新时间
```

#### 其他表
- **categories**: 博客分类表
- **tags**: 博客标签表
- **blog_tags**: 博客与标签关联表
- **blog_categories**: 博客与分类关联表

### 2. 项目结构

```
arrayDanceBackEnd/
├── base-service/                   # 业务服务层
│   ├── model/blog/                 # 博客数据模型
│   │   └── bloginfo.go             # 博客、分类、标签模型
│   ├── proto/blog/                 # 博客proto定义
│   │   ├── blog.proto              # proto文件
│   │   ├── blog.pb.go              # 生成的pb文件
│   │   └── blog_grpc.pb.go         # 生成的gRPC文件
│   ├── handler/blog/               # 博客业务逻辑
│   │   └── blog.go                 # 博客服务实现
│   ├── cmd/blog/                   # 博客服务启动入口
│   │   └── main.go                 # 主函数
│   └── util/blog/                  # 博客工具类
│       └── redisUtil.go            # Redis工具
├── douyin-api/                     # API网关层
│   ├── api/blog/                   # 博客API
│   │   └── blog.go                 # HTTP接口实现
│   ├── router/blog/                # 博客路由
│   │   └── blog.go                 # 路由配置
│   └── proto/blog/                 # API层proto引用
├── blog-sql/                       # 博客数据库脚本
│   └── init_blog.sql               # 初始化SQL
└── test/                           # 测试文件
```

### 3. 核心功能

#### 博客文章
- ✅ 发布博客
- ✅ 获取博客列表
- ✅ 获取博客详情
- ✅ 更新博客
- ✅ 删除博客(软删除)

#### 分类与标签
- ✅ 分类列表
- ✅ 标签列表

### 4. API接口文档

#### 发布博客
- **路径**: `/blog/publish`
- **方法**: POST
- **需要登录**: 是
- **参数**:
  - title (string): 标题
  - content (string): 内容
  - tag (string): 标签
  - category_id (int64): 分类ID
  - excerpt (string): 摘要

#### 获取博客列表
- **路径**: `/blog/list`
- **方法**: GET
- **需要登录**: 否
- **参数**:
  - user_id (int64): 用户ID(可选)
  - category_id (int64): 分类ID(可选)
  - tag_id (int64): 标签ID(可选)
  - keyword (string): 搜索关键词(可选)
  - page (int32): 页码
  - page_size (int32): 每页数量

#### 获取博客详情
- **路径**: `/blog/detail`
- **方法**: GET
- **需要登录**: 否
- **参数**:
  - blog_id (int64): 博客ID

#### 更新博客
- **路径**: `/blog/update`
- **方法**: POST
- **需要登录**: 是
- **参数**:
  - blog_id (int64): 博客ID
  - title (string): 标题
  - content (string): 内容
  - tag (string): 标签
  - category_id (int64): 分类ID
  - excerpt (string): 摘要

#### 删除博客
- **路径**: `/blog/delete`
- **方法**: POST
- **需要登录**: 是
- **参数**:
  - blog_id (int64): 博客ID

#### 获取分类列表
- **路径**: `/blog/categories`
- **方法**: GET
- **需要登录**: 否

#### 获取标签列表
- **路径**: `/blog/tags`
- **方法**: GET
- **需要登录**: 否

## 快速开始

### 1. 初始化数据库

```bash
# 执行博客初始化SQL
mysql -u root -p douyin < blog-sql/init_blog.sql
```

### 2. 启动博客服务

```bash
# 进入博客服务目录
cd base-service/cmd/blog

# 启动博客gRPC服务(端口8889)
go run main.go
```

### 3. 启动API网关

```bash
# 进入API网关目录
cd douyin-api/cmd

# 启动HTTP服务(端口8888)
go run main.go
```

### 4. 访问API

```bash
# 获取博客列表
curl http://localhost:8888/blog/list

# 发布博客(需要登录token)
curl -X POST http://localhost:8888/blog/publish \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "title=测试博客&content=这是一篇测试博客&tag=技术,Go"
```

## 兼容性说明

本博客平台完全兼容原有抖音功能:
- ✅ 用户系统共享
- ✅ 认证系统共享
- ✅ 视频功能保留
- ✅ 点赞评论系统保留
- ✅ 关注社交功能保留

用户可以同时使用抖音功能和博客功能,数据相互独立但用户体系统一。

## 性能优化

1. **阅读量缓存**: 使用Redis缓存博客阅读量,定时同步到数据库
2. **gRPC连接池**: 复用gRPC连接,减少连接开销
3. **数据库索引**: 为常用查询字段添加索引
4. **软删除**: 避免物理删除,提高性能

## 后续扩展建议

1. **Markdown支持**: 添加Markdown编辑器和渲染
2. **图片上传**: 集成七牛云对象存储
3. **SEO优化**: 添加meta标签,优化搜索引擎收录
4. **全文搜索**: 使用Elasticsearch实现博客搜索
5. **评论系统**: 复用现有评论系统,支持博客评论
6. **点赞系统**: 复用现有点赞系统,支持博客点赞

## 注意事项

1. **proto文件生成**: 实际使用时需要使用protoc命令生成完整的pb.go文件
2. **Redis配置**: 修改`base-service/global/blog/global.go`中的Redis连接信息
3. **MySQL配置**: 修改`base-service/global/blog/global.go`中的MySQL连接信息
4. **端口配置**: 博客服务默认使用8889端口,避免与现有服务冲突

## 技术支持

如有问题,请联系开发团队。
