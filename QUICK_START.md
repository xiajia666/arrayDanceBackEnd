# 🚀 快速开始 - 技术博客平台

## 5分钟快速体验

### 1️⃣ 克隆项目
```bash
git clone https://github.com/xiajia666/arrayDanceBackEnd.git
cd arrayDanceBackEnd
```

### 2️⃣ 初始化数据库
```bash
# 连接MySQL
mysql -u root -p

# 创建数据库
CREATE DATABASE douyin CHARACTER SET utf8mb4;

# 退出MySQL
EXIT;

# 初始化博客表
mysql -u root -p douyin < blog-sql/init_blog.sql
```

### 3️⃣ 修改配置
编辑 `base-service/global/blog/global.go`:
```go
// 修改MySQL连接字符串
dsn := "root:你的密码@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"

// 修改Redis配置(如果有密码)
RS = redis.NewClient(&redis.Options{
    Addr:     "127.0.0.1:6379",
    Password: "你的密码",
    DB:       0,
})
```

### 4️⃣ 启动服务

#### 终端1 - 启动博客gRPC服务
```bash
cd base-service/cmd/blog
go run main.go
# 默认监听 127.0.0.1:8889
```

#### 终端2 - 启动API网关
```bash
cd douyin-api/cmd
go run main.go
# 默认监听 :8080
```

### 5️⃣ 测试API

#### 浏览博客列表
```bash
curl http://localhost:8080/blog/list
```

#### 获取博客详情
```bash
curl http://localhost:8080/blog/detail?blog_id=1
```

#### (需要登录)发布博客
```bash
curl -X POST http://localhost:8080/blog/publish \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d "title=我的第一篇博客" \
  -d "content=这是博客内容..." \
  -d "tag=技术,Go"
```

## 📖 接口文档

### 博客相关接口

| 接口 | 路径 | 方法 | 说明 |
|------|------|------|------|
| 发布博客 | `/blog/publish` | POST | 需要登录 |
| 获取列表 | `/blog/list` | GET | 公开访问 |
| 获取详情 | `/blog/detail` | GET | 公开访问 |
| 更新博客 | `/blog/update` | POST | 需要登录 |
| 删除博客 | `/blog/delete` | POST | 需要登录 |
| 分类列表 | `/blog/categories` | GET | 公开访问 |
| 标签列表 | `/blog/tags` | GET | 公开访问 |

## 🎯 核心功能

### ✅ 已实现
- 博客文章的发布、查看、更新、删除
- 博客分类和标签管理
- 阅读量统计(Redis缓存)
- 软删除功能
- 推荐功能
- 与抖音用户系统兼容

### 🔜 待实现
- 博客评论功能
- 博客点赞功能
- 博客搜索
- 博客推荐算法
- 博客草稿功能

## 📚 更多文档

- [完整文档](README_NEW.md) - 详细的项目说明
- [部署指南](DEPLOYMENT_GUIDE.md) - 生产环境部署
- [API文档](INDEX.md) - 完整的API接口说明

## ❓ 常见问题

### Q: 如何注册用户?
A: 使用原有的抖音注册接口 `/user/register`

### Q: 如何获取Token?
A: 使用抖音登录接口 `/user/login` 获得Token,用于博客发布等操作

### Q: 博客数据存在哪里?
A: MySQL数据库的 `blogs` 表中

### Q: 如何查看日志?
A: 启动时会自动打印到控制台,也可配置日志文件

## 📞 需要帮助?

1. 查看 [INDEX.md](INDEX.md) 项目索引
2. 查看 [README_NEW.md](README_NEW.md) 完整文档
3. 提交Issue或联系开发团队

---
**祝你使用愉快!** 🎉
