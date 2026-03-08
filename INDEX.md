# 项目索引

## 📚 文档目录

### 1. 项目概述
- [`README.md`](README.md) - 项目主文档(抖音基础版)
- [`README_NEW.md`](README_NEW.md) - 新版文档(含博客平台)
- [`DEPLOYMENT_GUIDE.md`](DEPLOYMENT_GUIDE.md) - 部署指南

### 2. 技术博客平台
- **快速开始**: 
  - 查看 [README_NEW.md](README_NEW.md) 了解博客功能
  - 查看 [DEPLOYMENT_GUIDE.md](DEPLOYMENT_GUIDE.md) 进行部署
  
- **核心文件**:
  - [`base-service/cmd/blog/main.go`](base-service/cmd/blog/main.go) - 博客服务入口
  - [`base-service/handler/blog/blog.go`](base-service/handler/blog/blog.go) - 博客业务逻辑
  - [`douyin-api/api/blog/blog.go`](douyin-api/api/blog/blog.go) - 博客API接口
  - [`blog-sql/init_blog.sql`](blog-sql/init_blog.sql) - 数据库初始化脚本

- **数据模型**:
  - [`base-service/model/blog/bloginfo.go`](base-service/model/blog/bloginfo.go) - 博客数据模型
  - [`base-service/proto/blog/blog.proto`](base-service/proto/blog/blog.proto) - Protocol Buffers定义

### 3. 抖音基础功能
- 用户系统: 注册、登录、信息查询
- 视频系统: 上传、播放、推荐
- 互动系统: 评论、点赞
- 社交系统: 关注、消息

### 4. API接口

#### 博客相关接口
| 接口 | 方法 | 路径 | 说明 |
|------|------|------|------|
| 发布博客 | POST | `/blog/publish` | 需要登录 |
| 获取博客列表 | GET | `/blog/list` | 公开访问 |
| 获取博客详情 | GET | `/blog/detail` | 公开访问 |
| 更新博客 | POST | `/blog/update` | 需要登录 |
| 删除博客 | POST | `/blog/delete` | 需要登录 |
| 获取分类列表 | GET | `/blog/categories` | 公开访问 |
| 获取标签列表 | GET | `/blog/tags` | 公开访问 |

#### 抖音相关接口
详见 [README.md](README.md)

### 5. 测试
- 博客单元测试: `base-service/test/blog/blog_test.go`
- 运行测试: `go test -v ./base-service/test/blog/...`

### 6. 项目结构

```
arrayDanceBackEnd/
├── base-service/          # 基础服务层
│   ├── cmd/
│   │   ├── main.go       # 原有抖音服务
│   │   └── blog/         # 博客服务
│   ├── handler/
│   │   ├── user.go       # 用户处理
│   │   ├── video.go      # 视频处理
│   │   └── blog/         # 博客处理
│   ├── model/
│   │   ├── UserInfo.go   # 用户模型
│   │   └── blog/         # 博客模型
│   └── proto/
│       ├── user.proto
│       ├── video.proto
│       └── blog/         # 博客protobuf
├── douyin-api/           # API网关层
│   ├── api/
│   │   ├── user.go
│   │   ├── video.go
│   │   ├── comment.go
│   │   ├── favorite.go
│   │   └── blog/         # 博客API
│   └── router/
│       ├── user.go
│       ├── comment.go
│       ├── favorite.go
│       └── blog/         # 博客路由
├── interaction-service/  # 互动服务(评论、点赞)
├── social-service/       # 社交服务(关注、消息)
└── blog-sql/            # 博客数据库脚本
```

### 7. 快速启动

#### 启动博客服务
```bash
cd base-service/cmd/blog
go run main.go
```

#### 启动API网关
```bash
cd douyin-api/cmd
go run main.go
```

#### 访问博客接口
```bash
# 获取博客列表
curl http://localhost:8080/blog/list

# 获取博客详情
curl http://localhost:8080/blog/detail?blog_id=1
```

### 8. 下一步

1. 阅读 [README_NEW.md](README_NEW.md) 了解完整功能
2. 阅读 [DEPLOYMENT_GUIDE.md](DEPLOYMENT_GUIDE.md) 进行部署
3. 查看 [BLOG_README.md](BLOG_README.md) 了解技术细节

## 📞 联系方式

如有问题,请查看文档或联系开发团队。
