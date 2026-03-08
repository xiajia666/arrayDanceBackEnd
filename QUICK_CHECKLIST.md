# 技术博客平台 - 快速检查清单

## 🚀 启动前必做 (5分钟)

### 1. 环境检查
- [ ] 已安装 Go 1.19+
- [ ] 已安装 MySQL 8.0+
- [ ] 已安装 Redis 6.0+
- [ ] 已安装 protoc 3.19+

### 2. Proto文件生成 (必须!)
```bash
cd base-service/proto/blog
protoc --proto_path=. \
  --go_out=paths=source_relative:. \
  --go-grpc_out=paths=source_relative:. \
  blog.proto
```
- [ ] blog.pb.go 已生成
- [ ] blog_grpc.pb.go 已生成

### 3. 数据库初始化 (必须!)
```bash
mysql -u root -p douyin < blog-sql/init_blog.sql
```
- [ ] MySQL服务已启动
- [ ] 数据库 `douyin` 已创建
- [ ] 博客表已创建

### 4. 配置修改 (必须!)
- [ ] 修改 `base-service/global/blog/global.go` 的MySQL密码
- [ ] 修改 `base-service/global/blog/global.go` 的Redis地址和密码
- [ ] 修改 `base-service/global/global.go` 的配置
- [ ] 修改 `douyin-api/global/global.go` 的配置

### 5. 依赖安装
```bash
cd /Users/xiajia/Desktop/arrayDanceBackEnd
go mod tidy
```
- [ ] 依赖已安装

---

## 🎯 启动服务 (3分钟)

### 方式一: 使用启动脚本
```bash
chmod +x blog_start.sh
./blog_start.sh
```
- [ ] 选择选项4启动所有服务

### 方式二: 手动启动 (推荐)

**终端1 - 博客服务:**
```bash
cd base-service/cmd/blog
go run main.go
```
- [ ] 博客服务运行在 8889 端口

**终端2 - 主服务:**
```bash
cd base-service/cmd
go run main.go
```
- [ ] 主服务运行在 8887 端口

**终端3 - API网关:**
```bash
cd douyin-api/cmd
go run main.go
```
- [ ] API网关运行在 8888 端口

---

## ✅ 功能测试 (10分钟)

### 1. 测试博客列表 (无需登录)
```bash
curl http://localhost:8888/blog/list
```
- [ ] 返回成功 (status_code: 0)
- [ ] 返回空列表 (刚开始没有博客)

### 2. 注册用户
```bash
curl -X POST "http://localhost:8888/douyin/user/register/" \
  -d "username=testuser&password=123456"
```
- [ ] 返回 user_id 和 token

### 3. 登录获取token
```bash
curl -X POST "http://localhost:8888/douyin/user/login/" \
  -d "username=testuser&password=123456"
```
- [ ] 返回 token
- [ ] 保存 token 用于后续测试

### 4. 发布博客
```bash
curl -X POST "http://localhost:8888/blog/publish" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d "title=我的第一篇博客&content=这是博客内容&tag=技术,Go"
```
- [ ] 返回 blog_id

### 5. 获取博客列表
```bash
curl "http://localhost:8888/blog/list"
```
- [ ] 看到刚发布的博客

### 6. 获取博客详情
```bash
curl "http://localhost:8888/blog/detail?blog_id=1"
```
- [ ] 看到博客完整内容

### 7. 更新博客
```bash
curl -X POST "http://localhost:8888/blog/update" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d "blog_id=1&title=更新后的标题&content=更新后的内容"
```
- [ ] 返回成功

### 8. 删除博客
```bash
curl -X POST "http://localhost:8888/blog/delete" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d "blog_id=1"
```
- [ ] 返回成功

### 9. 获取分类列表
```bash
curl "http://localhost:8888/blog/categories"
```
- [ ] 看到默认分类

### 10. 获取标签列表
```bash
curl "http://localhost:8888/blog/tags"
```
- [ ] 看到标签列表

---

## 📊 数据库验证

### 检查表是否存在
```sql
USE douyin;
SHOW TABLES LIKE 'blog%';
```
应该看到:
- [ ] blogs
- [ ] categories
- [ ] tags
- [ ] blog_tags
- [ ] blog_categories

### 检查默认数据
```sql
SELECT * FROM categories;
```
应该看到7个默认分类:
- [ ] 技术分享
- [ ] Go语言
- [ ] 微服务
- [ ] 数据库
- [ ] 前端技术
- [ ] 运维部署
- [ ] 随笔

---

## 🐛 常见问题排查

### 问题1: Proto编译失败
**症状**: `cannot find package`
**解决**:
```bash
go mod tidy
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### 问题2: 数据库连接失败
**症状**: `connection refused`
**解决**:
- 检查MySQL是否启动: `mysql.server status`
- 检查密码是否正确
- 检查端口是否为3306

### 问题3: Redis连接失败
**症状**: `connection refused` 或 `WRONGPASS`
**解决**:
- 检查Redis是否启动: `redis-cli ping`
- 检查配置文件中的地址和密码
- 检查防火墙设置

### 问题4: 端口被占用
**症状**: `bind: address already in use`
**解决**:
```bash
# 查找占用端口的进程
lsof -i :8889

# 杀死进程
kill -9 PID

# 或修改端口号
```

### 问题5: gRPC连接失败
**症状**: `connection refused`
**解决**:
- 检查对应的gRPC服务是否启动
- 检查API网关的gRPC连接配置
- 检查防火墙设置

---

## 📚 文档索引

- **快速开始**: `BLOG_README.md`
- **完整文档**: `BLOG_PROJECT.md`
- **部署指南**: `BLOG_SERVICE_GUIDE.md`
- **项目总结**: `PROJECT_SUMMARY.md`
- **检查清单**: `QUICK_CHECKLIST.md` (本文件)

---

## 🎯 下一步

### 立即进行
1. [ ] 完成上述所有检查项
2. [ ] 运行测试用例: `go test ./base-service/test/blog/... -v`
3. [ ] 浏览API文档了解所有接口

### 短期优化
1. [ ] 复用现有的点赞系统
2. [ ] 复用现有的评论系统
3. [ ] 添加参数验证中间件
4. [ ] 完善错误处理

### 中期规划
1. [ ] 添加Markdown支持
2. [ ] 集成七牛云图片上传
3. [ ] 实现博客搜索功能
4. [ ] 添加访问统计

---

## ✨ 完成标准

当以下所有条件满足时,说明部署成功:

- [ ] 所有服务正常启动,无错误日志
- [ ] 所有API接口测试通过
- [ ] 数据库表结构正确
- [ ] 博客CRUD功能正常
- [ ] 分类和标签功能正常
- [ ] 与原有抖音功能兼容

---

**预估完成时间**: 30-60分钟

**遇到问题?** 查看 `BLOG_SERVICE_GUIDE.md` 的常见问题部分

**祝部署顺利! 🚀**
