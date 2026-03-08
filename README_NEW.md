# 抖音基础版(数组跳动后端代码) + 技术博客平台

## 一、项目介绍

本项目包含两个主要模块:
1. **抖音基础版**: 基于grpc通讯协议开发的高性能微服务
2. **技术博客平台**: 兼容抖音功能的技术博客后端系统

### 技术博客平台特点
- 使用Go语言开发,沿用原有抖音微服务架构
- 支持博客文章发布、编辑、删除、查看详情
- 支持分类和标签管理
- 兼容抖音用户系统、互动系统、消息系统
- 使用Redis缓存优化阅读量统计
- 支持软删除和推荐功能

## 二、技术选型

### 共用技术栈
- **Gin**: Web框架,提供HTTP服务
- **gRPC**: 微服务间通信协议
- **Protocol Buffers**: 数据序列化协议
- **JWT**: Token生成和权限校验
- **Gorm**: MySQL ORM操作
- **Go-redis**: Redis操作,缓存频繁访问数据
- **Zap**: 高性能日志打印
- **pprof**: 性能测试工具

### 抖音模块特有
- **ffmpeg**: 视频取帧,生成视频封面
- **七牛云**: 对象存储,存储视频、图片等静态资源

## 三、架构设计

### 微服务架构
项目分为三层:
1. **API层**: 负责鉴权和分发请求,调用远程服务返回数据
2. **业务服务层**: 负责与数据库交互和逻辑处理
3. **数据层**: MySQL数据库和Redis缓存

### 服务划分
- **douyin-api**: API网关层
- **base-service**: 基础服务层(用户、视频、博客)
- **interaction-service**: 互动服务层(评论、点赞)
- **social-service**: 社交服务层(关注、消息)

## 四、项目结构

```
arrayDanceBackEnd/
├── base-service/              # 基础服务层
│   ├── cmd/
│   │   ├── main.go           # 主服务启动入口
│   │   └── blog/             # 博客服务启动入口
│   │       └── main.go
│   ├── global/               # 全局配置
│   │   ├── global.go         # 全局变量
│   │   └── blog/             # 博客服务全局变量
│   │       └── global.go
│   ├── handler/              # 业务逻辑层
│   │   ├── user.go          # 用户处理
│   │   ├── video.go         # 视频处理
│   │   └── blog/            # 博客处理
│   │       └── blog.go
│   ├── model/                # 数据模型
│   │   ├── UserInfo.go      # 用户模型
│   │   ├── video/           # 视频模型
│   │   └── blog/            # 博客模型
│   │       └── bloginfo.go
│   ├── proto/                # Protocol Buffers定义
│   │   ├── user.proto
│   │   ├── video.proto
│   │   └── blog/            # 博客相关protobuf
│   │       ├── blog.proto
│   │       ├── blog.pb.go
│   │       ├── blog_grpc.pb.go
│   │       └── messages.pb.go
│   ├── test/                 # 测试代码
│   │   └── blog/            # 博客测试
│   │       └── blog_test.go
│   └── util/                 # 工具类
│       ├── jwtUtil.go
│       ├── redisUtil.go
│       └── blog/            # 博客工具
│           └── redisUtil.go
├── douyin-api/               # API网关层
│   ├── api/                  # API接口
│   │   ├── user.go
│   │   ├── video.go
│   │   ├── comment.go
│   │   ├── favorite.go
│   │   └── blog/            # 博客API
│   │       └── blog.go
│   ├── global/               # 全局配置
│   ├── globalinit/           # 初始化配置
│   ├── proto/                # API层protobuf
│   ├── router/               # 路由配置
│   │   ├── user.go
│   │   ├── comment.go
│   │   ├── favorite.go
│   │   └── blog/            # 博客路由
│   │       └── blog.go
│   └── util/                 # 工具类
│       ├── authMiddleware.go
│       ├── grpcConnUtils.go
│       └── blogGrpcUtil.go
├── interaction-service/      # 互动服务(评论、点赞)
├── social-service/           # 社交服务(关注、消息)
├── blog-sql/                 # 博客数据库初始化脚本
│   └── init_blog.sql
├── BLOG_README.md           # 博客平台详细文档
└── README.md                # 项目说明
```

## 五、数据库设计

### 博客相关表

#### 1. blogs - 博客文章表
```sql
CREATE TABLE `blogs` (
    `blog_id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `user_id` BIGINT NOT NULL,
    `title` VARCHAR(255) NOT NULL,
    `content` TEXT NOT NULL,
    `excerpt` VARCHAR(500),
    `views` BIGINT DEFAULT 0,
    `like_count` BIGINT DEFAULT 0,
    `comment_count` BIGINT DEFAULT 0,
    `tag` VARCHAR(100),
    `is_published` BOOLEAN DEFAULT TRUE,
    `is_deleted` BOOLEAN DEFAULT FALSE,
    `is_recommended` BOOLEAN DEFAULT FALSE,
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `update_time` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

#### 2. categories - 分类表
```sql
CREATE TABLE `categories` (
    `category_id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(50) NOT NULL UNIQUE,
    `display_order` BIGINT DEFAULT 0,
    `article_count` BIGINT DEFAULT 0,
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX `idx_display_order` (`display_order`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

#### 3. tags - 标签表
```sql
CREATE TABLE `tags` (
    `tag_id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(50) NOT NULL UNIQUE,
    `article_count` BIGINT DEFAULT 0,
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

#### 4. blog_tags - 博客标签关联表
```sql
CREATE TABLE `blog_tags` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `blog_id` BIGINT NOT NULL,
    `tag_id` BIGINT NOT NULL,
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX `idx_blog_id` (`blog_id`),
    INDEX `idx_tag_id` (`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

#### 5. blog_categories - 博客分类关联表
```sql
CREATE TABLE `blog_categories` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `blog_id` BIGINT NOT NULL,
    `category_id` BIGINT NOT NULL,
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX `idx_blog_id` (`blog_id`),
    INDEX `idx_category_id` (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

### 初始化数据库

```bash
# 初始化博客相关表
mysql -u root -p < blog-sql/init_blog.sql
```

## 六、API接口文档

### 博客相关接口

#### 1. 发布博客
```http
POST /blog/publish
Content-Type: application/x-www-form-urlencoded

title=我的第一篇博客&content=这里是内容&tag=技术,Go&category_id=1
```

**响应:**
```json
{
  "status_code": 0,
  "status_msg": "博客发布成功",
  "blog_id": 123
}
```

#### 2. 获取博客列表
```http
GET /blog/list?page=1&page_size=10&category_id=1
```

**响应:**
```json
{
  "status_code": 0,
  "status_msg": "success",
  "blog_list": [
    {
      "blog_id": 1,
      "user_id": 1,
      "title": "标题",
      "excerpt": "摘要",
      "views": 100,
      "like_count": 10,
      "comment_count": 5,
      "tag": "技术,Go",
      "author_name": "作者",
      "author_avatar": "头像URL",
      "create_time": "2026-03-08 12:00:00",
      "update_time": "2026-03-08 12:00:00"
    }
  ]
}
```

#### 3. 获取博客详情
```http
GET /blog/detail?blog_id=1
```

**响应:**
```json
{
  "status_code": 0,
  "status_msg": "success",
  "blog_info": {
    "blog_id": 1,
    "user_id": 1,
    "title": "标题",
    "content": "完整内容",
    "excerpt": "摘要",
    "views": 100,
    "like_count": 10,
    "comment_count": 5,
    "tag": "技术,Go",
    "author_name": "作者",
    "author_avatar": "头像URL",
    "create_time": "2026-03-08 12:00:00",
    "update_time": "2026-03-08 12:00:00"
  }
}
```

#### 4. 更新博客
```http
POST /blog/update
Content-Type: application/x-www-form-urlencoded

blog_id=1&title=更新后的标题&content=更新后的内容&tag=技术,Go,后端
```

**响应:**
```json
{
  "status_code": 0,
  "status_msg": "博客更新成功"
}
```

#### 5. 删除博客
```http
POST /blog/delete
Content-Type: application/x-www-form-urlencoded

blog_id=1
```

**响应:**
```json
{
  "status_code": 0,
  "status_msg": "博客删除成功"
}
```

#### 6. 获取分类列表
```http
GET /blog/categories
```

**响应:**
```json
{
  "status_code": 0,
  "status_msg": "success",
  "category_list": [
    {
      "category_id": 1,
      "name": "技术分享",
      "article_count": 10,
      "display_order": 1
    }
  ]
}
```

#### 7. 获取标签列表
```http
GET /blog/tags
```

**响应:**
```json
{
  "status_code": 0,
  "status_msg": "success",
  "tag_list": [
    {
      "tag_id": 1,
      "name": "技术",
      "article_count": 20
    }
  ]
}
```

## 七、服务启动

### 1. 博客服务启动
```bash
cd base-service/cmd/blog
go run main.go -ip 127.0.0.1 -port 8889
```

### 2. API网关启动
```bash
cd douyin-api/cmd
go run main.go
```

### 3. 其他服务启动(抖音功能)
```bash
# 基础服务
cd base-service/cmd
go run main.go -ip 127.0.0.1 -port 8885

# 互动服务
cd interaction-service
go run main.go

# 社交服务
cd social-service/cmd
go run main.go
```

## 八、环境配置

### MySQL配置
修改 `base-service/global/blog/global.go` 中的DSN配置:
```go
dsn := "root:password@tcp(localhost:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
```

### Redis配置
修改 `base-service/global/blog/global.go` 中的Redis配置:
```go
RS = redis.NewClient(&redis.Options{
    Addr:     "127.0.0.1:6379",
    Password: "", // Redis密码
    DB:       0,  // Redis数据库
})
```

### gRPC端口配置
- 博客服务: 默认8889端口
- 基础服务: 默认8885端口
- 互动服务: 默认8887端口
- 社交服务: 默认8886端口

## 九、性能优化

1. **阅读量统计**: 使用Redis缓存,定时同步到数据库
2. **gRPC连接池**: 使用连接池管理gRPC连接,减少创建销毁开销
3. **数据库索引**: 关键字段添加索引优化查询性能
4. **缓存策略**: 热门文章、分类、标签使用Redis缓存

## 十、测试

### 单元测试
```bash
# 博客服务测试
cd base-service/test/blog
go test -v

# 抖音服务测试
cd base-service/test
go test -v
```

### 性能测试
使用pprof进行性能分析:
```bash
go tool pprof http://localhost:6060/debug/pprof/profile
```

## 十一、项目亮点

### 抖音基础版
- 使用gRPC通讯协议开发高性能微服务
- 使用连接池技术构建连接工厂,复用gRPC连接
- 使用预编译SQL防止SQL注入
- 使用Redis缓存频繁更改的数据
- 使用ffmpeg进行视频取帧生成封面
- 使用七牛云对象存储

### 技术博客平台
- 完全兼容抖音用户系统和互动系统
- 支持博客文章的完整生命周期管理
- 使用Redis优化阅读量统计
- 支持分类和标签管理
- 支持软删除和推荐功能
- 代码结构清晰,易于扩展

## 十二、待完成功能

1. [ ] 博客评论功能集成
2. [ ] 博客点赞功能集成
3. [ ] 博客搜索功能(关键词搜索)
4. [ ] 博客推荐算法
5. [ ] 博客置顶功能
6. [ ] 博客草稿功能
7. [ ] 博客访问统计(按日/月)
8. [ ] 博客导出功能(Markdown/HTML)
9. [ ] 博客Markdown编辑器支持
10. [ ] 博客图片上传和管理

## 十三、注意事项

1. 当前gRPC连接使用临时方案,生产环境需要优化为连接池
2. 需要在 `douyin-api/global/global.go` 中初始化博客服务的gRPC连接
3. 数据库表需要先执行初始化SQL创建
4. 用户认证复用现有的JWT Token机制
5. 与现有抖音功能完全兼容,可共用用户系统、互动系统

## 十四、项目总结

### 已实现功能
- [x] 抖音用户系统(注册、登录、信息查询)
- [x] 抖音视频功能(上传、播放、推荐)
- [x] 抖音互动功能(评论、点赞)
- [x] 抖音社交功能(关注、消息)
- [x] 博客文章管理(发布、查看、更新、删除)
- [x] 博客分类和标签管理
- [x] 博客阅读量统计(使用Redis)

### 技术收获
- 掌握了gRPC微服务架构设计
- 熟悉了Go语言在高并发场景下的应用
- 学习了连接池技术和性能优化方法
- 了解了Redis缓存策略和应用场景
- 实践了数据库索引优化

## 十五、部署

1. 安装Go环境(1.21+)
2. 安装MySQL 8.0
3. 安装Redis
4. 安装ffmpeg环境(抖音视频功能需要)
5. 修改各模块中global的MySQL连接和Redis连接
6. 修改base-service下的video中的七牛云密钥和仓库名称
7. 初始化数据库表(执行blog-sql/init_blog.sql)
8. 启动各服务

## 十六、联系方式

如有问题,请联系开发团队。
