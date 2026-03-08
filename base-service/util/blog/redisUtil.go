package util

import (
	"demotest/base-service/global"
	"time"
)

// 增加阅读量计数
func IncreaseBlogViews(blogId int64) {
	viewsKey := fmt.Sprintf("blog:views:%d", blogId)
	global.RS.Incr(viewsKey)

	// 设置过期时间,防止内存泄漏
	// 每天定时同步到数据库
}

// 获取博客阅读量
func GetBlogViews(blogId int64) int64 {
	viewsKey := fmt.Sprintf("blog:views:%d", blogId)
	count, err := global.RS.Get(viewsKey).Int64()
	if err != nil {
		return 0
	}
	return count
}

// 批量同步阅读量到数据库
func SyncBlogViews() {
	// 获取所有博客阅读量缓存
	keys := global.RS.Keys("blog:views:*")

	for _, key := range keys.Val() {
		views, err := global.RS.Get(key).Int64()
		if err != nil {
			continue
		}

		// 解析博客ID
		// TODO: 从key中提取blog_id并更新数据库

		// 删除缓存
		// global.RS.Del(key)
	}
}

// 定时同步任务
func StartSyncTask() {
	// 每天凌晨2点同步
	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for {
			select {
			case <-ticker.C:
				SyncBlogViews()
			}
		}
	}()
}
