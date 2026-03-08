# 技术博客平台部署指南

## 快速开始

### 1. 环境要求
- Go 1.21+
- MySQL 8.0+
- Redis 5.0+
- ffmpeg (可选,用于视频处理)

### 2. 数据库初始化

```bash
# 创建数据库
mysql -u root -p
CREATE DATABASE douyin CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
EXIT;

# 初始化博客相关表
mysql -u root -p douyin < blog-sql/init_blog.sql
```

### 3. 配置修改

#### MySQL配置
编辑 `base-service/global/blog/global.go`:
```go
dsn := "root:你的密码@tcp(localhost:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
```

#### Redis配置
编辑 `base-service/global/blog/global.go`:
```go
RS = redis.NewClient(&redis.Options{
    Addr:     "你的Redis地址:6379",
    Password: "你的Redis密码", // 如果有密码
    DB:       0,
})
```

### 4. 启动服务

#### 启动博客服务
```bash
cd base-service/cmd/blog
go run main.go -ip 127.0.0.1 -port 8889
```

#### 启动API网关
```bash
cd douyin-api/cmd
go run main.go
```

### 5. 验证部署

```bash
# 测试博客列表接口
curl http://localhost:8080/blog/list

# 测试发布博客(需要登录)
curl -X POST http://localhost:8080/blog/publish \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d "title=测试博客&content=这是测试内容"
```

## 服务端口说明

| 服务 | 端口 | 说明 |
|------|------|------|
| 博客服务 | 8889 | 博客gRPC服务 |
| 基础服务 | 8885 | 用户、视频gRPC服务 |
| 互动服务 | 8887 | 评论、点赞gRPC服务 |
| 社交服务 | 8886 | 关注、消息gRPC服务 |
| API网关 | 8080 | HTTP API接口 |

## 常见问题

### 1. 数据库连接失败
检查MySQL是否启动,配置是否正确

### 2. gRPC连接失败
确保博客服务已启动并监听正确端口

### 3. Redis连接失败
检查Redis服务是否启动,配置是否正确

## 生产环境建议

1. 使用Docker容器化部署
2. 配置负载均衡
3. 使用连接池优化gRPC连接
4. 配置监控和日志收集
5. 定期备份数据库
6. 使用HTTPS加密通信
