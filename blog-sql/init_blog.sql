-- 博客平台数据库初始化SQL

-- 创建博客表
CREATE TABLE IF NOT EXISTS `blogs` (
    `blog_id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '博客ID',
    `user_id` BIGINT NOT NULL COMMENT '作者ID',
    `title` VARCHAR(255) NOT NULL COMMENT '标题',
    `content` TEXT NOT NULL COMMENT '内容',
    `excerpt` VARCHAR(500) COMMENT '摘要',
    `views` BIGINT DEFAULT 0 COMMENT '阅读量',
    `like_count` BIGINT DEFAULT 0 COMMENT '点赞数',
    `comment_count` BIGINT DEFAULT 0 COMMENT '评论数',
    `tag` VARCHAR(100) COMMENT '标签,多个标签用逗号分隔',
    `is_published` BOOLEAN DEFAULT TRUE COMMENT '是否发布',
    `is_deleted` BOOLEAN DEFAULT FALSE COMMENT '软删除',
    `is_recommended` BOOLEAN DEFAULT FALSE COMMENT '是否推荐',
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='博客文章表';

-- 创建分类表
CREATE TABLE IF NOT EXISTS `categories` (
    `category_id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '分类ID',
    `name` VARCHAR(50) NOT NULL UNIQUE COMMENT '分类名称',
    `display_order` BIGINT DEFAULT 0 COMMENT '显示顺序',
    `article_count` BIGINT DEFAULT 0 COMMENT '文章数量',
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    INDEX `idx_display_order` (`display_order`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='博客分类表';

-- 创建标签表
CREATE TABLE IF NOT EXISTS `tags` (
    `tag_id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '标签ID',
    `name` VARCHAR(50) NOT NULL UNIQUE COMMENT '标签名称',
    `article_count` BIGINT DEFAULT 0 COMMENT '文章数量',
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='博客标签表';

-- 创建博客标签关联表
CREATE TABLE IF NOT EXISTS `blog_tags` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
    `blog_id` BIGINT NOT NULL COMMENT '博客ID',
    `tag_id` BIGINT NOT NULL COMMENT '标签ID',
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    INDEX `idx_blog_id` (`blog_id`),
    INDEX `idx_tag_id` (`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='博客标签关联表';

-- 创建博客分类关联表
CREATE TABLE IF NOT EXISTS `blog_categories` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
    `blog_id` BIGINT NOT NULL COMMENT '博客ID',
    `category_id` BIGINT NOT NULL COMMENT '分类ID',
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    INDEX `idx_blog_id` (`blog_id`),
    INDEX `idx_category_id` (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='博客分类关联表';

-- 插入默认分类
INSERT INTO `categories` (`name`, `display_order`) VALUES
('技术分享', 1),
('Go语言', 2),
('微服务', 3),
('数据库', 4),
('前端技术', 5),
('运维部署', 6),
('随笔', 7);

-- 创建索引优化查询
CREATE INDEX `idx_blog_published` ON `blogs` (`is_published`, `create_time`);
