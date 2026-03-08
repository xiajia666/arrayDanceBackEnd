# 技术博客平台 - 文档索引

欢迎使用技术博客平台! 🎉

本索引帮助您快速找到所需的文档和信息。

---

## 🚀 新手快速开始

### 我是第一次使用,想快速启动项目

1. **5分钟快速开始** → 查看 [`QUICK_CHECKLIST.md`](QUICK_CHECKLIST.md)
   - 环境检查
   - 服务启动
   - 基础测试

2. **完整部署指南** → 查看 [`BLOG_SERVICE_GUIDE.md`](BLOG_SERVICE_GUIDE.md)
   - 详细步骤说明
   - 常见问题解答
   - 故障排查

---

## 📖 文档导航

### 项目概述类

| 文档 | 用途 | 阅读时间 |
|------|------|----------|
| [`BLOG_README.md`](BLOG_README.md) | 快速了解项目功能和技术栈 | 5分钟 |
| [`BLOG_PROJECT.md`](BLOG_PROJECT.md) | 完整项目文档,包含架构设计 | 15分钟 |
| [`PROJECT_SUMMARY.md`](PROJECT_SUMMARY.md) | 开发总结和功能清单 | 10分钟 |

### 开发部署类

| 文档 | 用途 | 阅读时间 |
|------|------|----------|
| [`BLOG_SERVICE_GUIDE.md`](BLOG_SERVICE_GUIDE.md) | **部署指南** - 必读 | 20分钟 |
| [`QUICK_CHECKLIST.md`](QUICK_CHECKLIST.md) | **快速检查清单** - 必读 | 5分钟 |
| [`blog_start.sh`](blog_start.sh) | 启动脚本 | - |

### API接口类

| 文档 | 用途 | 阅读时间 |
|------|------|----------|
| [`API_REFERENCE.md`](API_REFERENCE.md) | **API速查表** - 开发必备 | 10分钟 |
| [`base-service/proto/blog/blog.proto`](base-service/proto/blog/blog.proto) | Proto定义文件 | - |

### 代码示例类

| 文档 | 用途 |
|------|------|
| [`base-service/test/blog/blog_test.go`](base-service/test/blog/blog_test.go) | gRPC测试示例 |
| [`base-service/test/blog/blog_full_test.go`](base-service/test/blog/blog_full_test.go) | 完整测试示例 |

---

## 🎯 根据需求查找

### 我想了解项目功能
→ 查看 [`BLOG_README.md`](BLOG_README.md) 和 [`BLOG_PROJECT.md`](BLOG_PROJECT.md)

### 我要部署项目
→ 查看 [`BLOG_SERVICE_GUIDE.md`](BLOG_SERVICE_GUIDE.md) 和 [`QUICK_CHECKLIST.md`](QUICK_CHECKLIST.md)

### 我要调用API
→ 查看 [`API_REFERENCE.md`](API_REFERENCE.md)

### 我要开发新功能
→ 查看 [`BLOG_PROJECT.md`](BLOG_PROJECT.md) 第九部分"开发指南"

### 我遇到问题了
→ 查看 [`BLOG_SERVICE_GUIDE.md`](BLOG_SERVICE_GUIDE.md) 第八部分"常见问题"

### 我想了解数据库设计
→ 查看 [`BLOG_PROJECT.md`](BLOG_PROJECT.md) 第七部分"数据库设计"

### 我想了解技术架构
→ 查看 [`BLOG_PROJECT.md`](BLOG_PROJECT.md) 第三部分"技术架构"

---

## 📁 重要文件位置

### 配置文件
- `base-service/global/blog/global.go` - 博客服务配置 ⚠️ 需要修改
- `base-service/global/global.go` - 全局配置 ⚠️ 需要修改
- `douyin-api/global/global.go` - API网关配置 ⚠️ 需要修改

### 核心代码
- `base-service/model/blog/bloginfo.go` - 数据模型
- `base-service/handler/blog/blog.go` - 业务逻辑
- `douyin-api/api/blog/blog.go` - HTTP接口
- `base-service/proto/blog/blog.proto` - Proto定义

### 数据库
- `blog-sql/init_blog.sql` - 数据库初始化脚本 ⚠️ 需要执行

### 启动入口
- `base-service/cmd/blog/main.go` - 博客服务
- `douyin-api/cmd/main.go` - API网关

---

## ⚡ 快速命令

### 1. 初始化数据库
```bash
mysql -u root -p douyin < blog-sql/init_blog.sql
```

### 2. 生成Proto文件
```bash
cd base-service/proto/blog
protoc --proto_path=. \
  --go_out=paths=source_relative:. \
  --go-grpc_out=paths=source_relative:. \
  blog.proto
```

### 3. 启动所有服务
```bash
./blog_start.sh
# 选择选项4
```

### 4. 测试API
```bash
curl http://localhost:8888/blog/list
```

### 5. 运行测试
```bash
cd base-service/test/blog
go test -v .
```

---

## 🎓 学习路径

### 第一天: 熟悉项目
1. 阅读 [`BLOG_README.md`](BLOG_README.md) 了解功能
2. 阅读 [`BLOG_PROJECT.md`](BLOG_PROJECT.md) 了解架构
3. 查看 [`API_REFERENCE.md`](API_REFERENCE.md) 了解接口

### 第二天: 部署项目
1. 按照 [`QUICK_CHECKLIST.md`](QUICK_CHECKLIST.md) 检查环境
2. 按照 [`BLOG_SERVICE_GUIDE.md`](BLOG_SERVICE_GUIDE.md) 部署项目
3. 运行测试验证功能

### 第三天: 开发扩展
1. 阅读 [`BLOG_PROJECT.md`](BLOG_PROJECT.md) 的开发指南
2. 参考现有代码开发新功能
3. 运行测试确保兼容性

---

## ❓ 常见问题快速链接

| 问题 | 解决方案位置 |
|------|-------------|
| Proto编译失败 | [`BLOG_SERVICE_GUIDE.md`](BLOG_SERVICE_GUIDE.md) 第八部分 Q1 |
| 数据库连接失败 | [`BLOG_SERVICE_GUIDE.md`](BLOG_SERVICE_GUIDE.md) 第八部分 Q2 |
| Redis连接失败 | [`BLOG_SERVICE_GUIDE.md`](BLOG_SERVICE_GUIDE.md) 第八部分 Q3 |
| 端口被占用 | [`BLOG_SERVICE_GUIDE.md`](BLOG_SERVICE_GUIDE.md) 第八部分 Q4 |
| gRPC连接失败 | [`BLOG_SERVICE_GUIDE.md`](BLOG_SERVICE_GUIDE.md) 第八部分 Q5 |

---

## 📞 获取帮助

1. **查看文档**: 所有文档都在本目录
2. **查看代码注释**: 关键代码都有详细注释
3. **运行测试**: 测试用例是最好的示例
4. **检查日志**: 服务运行时会输出详细日志

---

## 🎯 下一步

根据您的角色选择:

### 如果您是运维人员
→ 重点阅读: [`BLOG_SERVICE_GUIDE.md`](BLOG_SERVICE_GUIDE.md), [`QUICK_CHECKLIST.md`](QUICK_CHECKLIST.md)

### 如果您是开发人员
→ 重点阅读: [`BLOG_PROJECT.md`](BLOG_PROJECT.md), [`API_REFERENCE.md`](API_REFERENCE.md)

### 如果您是产品经理
→ 重点阅读: [`BLOG_README.md`](BLOG_README.md), [`PROJECT_SUMMARY.md`](PROJECT_SUMMARY.md)

### 如果您是测试人员
→ 重点阅读: [`API_REFERENCE.md`](API_REFERENCE.md), [`base-service/test/blog/`](base-service/test/blog/) 目录

---

## 📊 项目统计

- **总代码量**: ~2000+ 行
- **新增文件**: 20+ 个
- **核心功能**: 7 个
- **API接口**: 10+ 个
- **文档数量**: 6 个

---

**祝您使用愉快!** 🚀

如有问题,请按上述索引查找相关文档。

最后更新: 2026-03-08
