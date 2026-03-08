# 🎉 欢迎使用技术博客平台!

## 📖 快速导航

### 🆕 新用户必读
1. **项目概览** - 查看 [`INDEX.md`](INDEX.md) 了解项目结构
2. **快速开始** - 查看 [`QUICK_CHECKLIST.md`](QUICK_CHECKLIST.md) 5分钟上手
3. **部署指南** - 查看 [`BLOG_SERVICE_GUIDE.md`](BLOG_SERVICE_GUIDE.md) 完整部署

### 📚 完整文档
- 📌 [`INDEX.md`](INDEX.md) - **文档索引** (推荐从这里开始!)
- ⚡ [`QUICK_CHECKLIST.md`](QUICK_CHECKLIST.md) - 快速检查清单
- 🔧 [`BLOG_SERVICE_GUIDE.md`](BLOG_SERVICE_GUIDE.md) - 完整部署指南
- 📖 [`API_REFERENCE.md`](API_REFERENCE.md) - API速查表
- 📚 [`BLOG_PROJECT.md`](BLOG_PROJECT.md) - 完整项目文档
- 🚀 [`BLOG_README.md`](BLOG_README.md) - 快速开始
- 📝 [`PROJECT_SUMMARY.md`](PROJECT_SUMMARY.md) - 项目总结
- ✅ [`FINISHED.md`](FINISHED.md) - 完成说明
- 📦 [`DELIVERY.md`](DELIVERY.md) - 交付文档

---

## 🎯 项目亮点

✅ **双功能合一**: 抖音短视频 + 技术博客平台
✅ **统一用户系统**: 共享认证和用户管理
✅ **微服务架构**: 高性能、易扩展
✅ **完整文档**: 9个详细文档,覆盖所有方面
✅ **易于部署**: 3步即可启动项目

---

## 🚀 3步快速启动

### 第1步: 生成Proto文件 ⚠️
```bash
cd base-service/proto/blog
protoc --proto_path=. \
  --go_out=paths=source_relative:. \
  --go-grpc_out=paths=source_relative:. \
  blog.proto
```

### 第2步: 初始化数据库 ⚠️
```bash
mysql -u root -p douyin < blog-sql/init_blog.sql
```

### 第3步: 启动服务 ⚡
```bash
./blog_start.sh
# 选择选项4启动所有服务
```

---

## 📊 核心功能

### 博客功能
- ✅ 发布博客
- ✅ 博客列表
- ✅ 博客详情
- ✅ 更新博客
- ✅ 删除博客
- ✅ 分类管理
- ✅ 标签管理

### 抖音功能 (原有)
- ✅ 用户注册/登录
- ✅ 视频发布/浏览
- ✅ 点赞/评论
- ✅ 关注/粉丝
- ✅ 私信聊天

---

## 🧪 测试API

```bash
# 获取博客列表
curl http://localhost:8888/blog/list

# 发布博客 (需要token)
curl -X POST "http://localhost:8888/blog/publish" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d "title=测试博客&content=这是内容&tag=技术"
```

---

## 📞 需要帮助?

1. **部署问题** → 查看 [`QUICK_CHECKLIST.md`](QUICK_CHECKLIST.md)
2. **使用问题** → 查看 [`BLOG_SERVICE_GUIDE.md`](BLOG_SERVICE_GUIDE.md)
3. **开发问题** → 查看 [`BLOG_PROJECT.md`](BLOG_PROJECT.md)
4. **API问题** → 查看 [`API_REFERENCE.md`](API_REFERENCE.md)

---

## 🎓 学习路径

### 第一天: 熟悉项目
- 阅读 [`INDEX.md`](INDEX.md) 了解文档结构
- 阅读 [`BLOG_README.md`](BLOG_README.md) 了解功能
- 查看 [`API_REFERENCE.md`](API_REFERENCE.md) 了解接口

### 第二天: 部署项目
- 按照 [`QUICK_CHECKLIST.md`](QUICK_CHECKLIST.md) 检查环境
- 按照 [`BLOG_SERVICE_GUIDE.md`](BLOG_SERVICE_GUIDE.md) 部署项目
- 运行测试验证功能

### 第三天: 开发扩展
- 阅读 [`BLOG_PROJECT.md`](BLOG_PROJECT.md) 的开发指南
- 参考现有代码开发新功能
- 运行测试确保兼容性

---

## 🌟 技术栈

- **语言**: Go 1.19+
- **Web框架**: Gin
- **RPC框架**: gRPC
- **ORM**: Gorm
- **数据库**: MySQL 8.0
- **缓存**: Redis 6.0
- **序列化**: Protobuf

---

## 📈 项目统计

- **新增文件**: 20+ 个
- **代码行数**: 2000+ 行
- **核心功能**: 7 个
- **API接口**: 10+ 个
- **数据库表**: 5 个
- **文档数量**: 9 个

---

## ⚠️ 重要提醒

使用项目前,必须完成:
1. ✅ 生成Proto文件 (使用protoc)
2. ✅ 初始化数据库 (执行SQL脚本)
3. ✅ 修改配置文件 (MySQL, Redis连接)

详细步骤请查看 [`BLOG_SERVICE_GUIDE.md`](BLOG_SERVICE_GUIDE.md)

---

## 🎉 开始使用!

**推荐阅读顺序**:
1. [`INDEX.md`](INDEX.md) - 找到你需要的文档
2. [`QUICK_CHECKLIST.md`](QUICK_CHECKLIST.md) - 快速部署
3. [`API_REFERENCE.md`](API_REFERENCE.md) - 了解API

**祝使用愉快! 🚀**

---

**最后更新**: 2026-03-08
**项目版本**: v1.0.0
