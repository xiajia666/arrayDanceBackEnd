# 博客平台API接口速查表

## 📌 基本信息

- **Base URL**: `http://localhost:8888`
- **认证方式**: JWT Bearer Token
- **Content-Type**: `application/x-www-form-urlencoded`

---

## 🔓 无需登录的接口

### 1. 获取博客列表
```
GET /blog/list
```

**Query Parameters**:
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| user_id | int64 | 否 | 筛选特定用户的博客 |
| category_id | int64 | 否 | 筛选特定分类 |
| tag_id | int64 | 否 | 筛选特定标签 |
| keyword | string | 否 | 搜索关键词 |
| page | int32 | 否 | 页码 (默认: 1) |
| page_size | int32 | 否 | 每页数量 (默认: 10) |

**示例**:
```bash
curl "http://localhost:8888/blog/list?page=1&page_size=10"
```

**响应**:
```json
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
      "like_count": 10,
      "comment_count": 5,
      "tag": "技术,Go",
      "create_time": "2026-03-08 10:00:00"
    }
  ]
}
```

---

### 2. 获取博客详情
```
GET /blog/detail
```

**Query Parameters**:
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| blog_id | int64 | 是 | 博客ID |

**示例**:
```bash
curl "http://localhost:8888/blog/detail?blog_id=1"
```

**响应**:
```json
{
  "status_code": 0,
  "status_msg": "success",
  "blog_info": {
    "blog_id": 1,
    "user_id": 1,
    "title": "博客标题",
    "content": "博客完整内容...",
    "excerpt": "博客摘要",
    "views": 100,
    "like_count": 10,
    "comment_count": 5,
    "tag": "技术,Go",
    "author_name": "作者名",
    "author_avatar": "头像URL",
    "create_time": "2026-03-08 10:00:00",
    "update_time": "2026-03-08 10:00:00"
  }
}
```

---

### 3. 获取分类列表
```
GET /blog/categories
```

**示例**:
```bash
curl "http://localhost:8888/blog/categories"
```

**响应**:
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
    },
    {
      "category_id": 2,
      "name": "Go语言",
      "article_count": 5,
      "display_order": 2
    }
  ]
}
```

---

### 4. 获取标签列表
```
GET /blog/tags
```

**示例**:
```bash
curl "http://localhost:8888/blog/tags"
```

**响应**:
```json
{
  "status_code": 0,
  "status_msg": "success",
  "tag_list": [
    {
      "tag_id": 1,
      "name": "Go",
      "article_count": 5
    },
    {
      "tag_id": 2,
      "name": "技术",
      "article_count": 10
    }
  ]
}
```

---

## 🔒 需要登录的接口

### 5. 发布博客
```
POST /blog/publish
```

**Headers**:
```
Authorization: Bearer YOUR_TOKEN
Content-Type: application/x-www-form-urlencoded
```

**Body Parameters**:
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| title | string | 是 | 博客标题 |
| content | string | 是 | 博客内容 |
| tag | string | 否 | 标签,多个用逗号分隔 |
| category_id | int64 | 否 | 分类ID |
| excerpt | string | 否 | 摘要(不填自动生成) |

**示例**:
```bash
curl -X POST "http://localhost:8888/blog/publish" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "title=Go语言最佳实践&content=# Go语言%0A%0A这是一篇关于Go的博客...&tag=Go,编程,技术"
```

**响应**:
```json
{
  "status_code": 0,
  "status_msg": "博客发布成功",
  "blog_id": 1
}
```

---

### 6. 更新博客
```
POST /blog/update
```

**Headers**:
```
Authorization: Bearer YOUR_TOKEN
Content-Type: application/x-www-form-urlencoded
```

**Body Parameters**:
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| blog_id | int64 | 是 | 博客ID |
| title | string | 否 | 新标题 |
| content | string | 否 | 新内容 |
| tag | string | 否 | 新标签 |
| category_id | int64 | 否 | 新分类ID |
| excerpt | string | 否 | 新摘要 |

**示例**:
```bash
curl -X POST "http://localhost:8888/blog/update" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "blog_id=1&title=更新后的标题&content=更新后的内容"
```

**响应**:
```json
{
  "status_code": 0,
  "status_msg": "博客更新成功"
}
```

---

### 7. 删除博客
```
POST /blog/delete
```

**Headers**:
```
Authorization: Bearer YOUR_TOKEN
Content-Type: application/x-www-form-urlencoded
```

**Body Parameters**:
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| blog_id | int64 | 是 | 博客ID |

**示例**:
```bash
curl -X POST "http://localhost:8888/blog/delete" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "blog_id=1"
```

**响应**:
```json
{
  "status_code": 0,
  "status_msg": "博客删除成功"
}
```

---

## 📝 状态码说明

| 状态码 | 说明 |
|--------|------|
| 0 | 成功 |
| 400 | 请求参数错误 |
| 401 | 未登录或token无效 |
| 403 | 无权限操作 |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |

---

## 🔗 相关API (抖音原有功能)

### 用户相关
- `POST /douyin/user/register/` - 用户注册
- `POST /douyin/user/login/` - 用户登录
- `GET /douyin/user/` - 获取用户信息

### 视频相关
- `GET /douyin/feed/` - 视频流
- `POST /douyin/publish/action/` - 发布视频
- `GET /douyin/publish/list/` - 发布列表

### 互动相关
- `POST /douyin/favorite/action/` - 点赞/取消赞
- `GET /douyin/favorite/list/` - 喜欢列表
- `POST /douyin/comment/action/` - 评论
- `GET /douyin/comment/list/` - 评论列表

### 社交相关
- `POST /douyin/relation/action/` - 关注/取消关注
- `GET /douyin/relation/follow/list/` - 关注列表
- `GET /douyin/relation/follower/list/` - 粉丝列表

---

## 💡 使用提示

1. **获取Token**: 先调用登录接口获取token,然后在需要登录的接口中添加 `Authorization: Bearer YOUR_TOKEN`

2. **博客标签**: 多个标签用英文逗号分隔,如: `tag=Go,编程,技术`

3. **博客内容**: 支持Markdown格式 (未来版本)

4. **分页查询**: 使用 `page` 和 `page_size` 参数,默认每页10条

5. **错误处理**: 所有接口都返回 `status_code` 和 `status_msg`,根据状态码判断操作是否成功

---

## 🧪 测试工具推荐

1. **curl**: 命令行测试
2. **Postman**: GUI接口测试工具
3. **Apifox**: 在线API文档和测试平台

---

## 📚 完整文档

- 详细部署指南: `BLOG_SERVICE_GUIDE.md`
- 完整项目文档: `BLOG_PROJECT.md`
- 快速检查清单: `QUICK_CHECKLIST.md`

---

**最后更新**: 2026-03-08
**API版本**: v1.0.0
