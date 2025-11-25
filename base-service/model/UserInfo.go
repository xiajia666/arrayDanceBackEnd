package model

import "time"

type Video struct {
	VideoId       int32
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int32
	CommentCount  int32
	Title         string
	CreatedTime   time.Time
}
type User struct {
	ID        uint64  `gorm:"primaryKey;autoIncrement;comment:主键"`
	UID       string  `gorm:"type:varchar(20);uniqueIndex;comment:对外UID"`
	Username  string  `gorm:"type:varchar(30);uniqueIndex;comment:登录名"`
	Password  string  `gorm:"type:varchar(95);comment:bcrypt哈希"`
	Email     string  `gorm:"type:varchar(100);uniqueIndex;comment:邮箱"`
	Phone     *string `gorm:"type:varchar(20);uniqueIndex;comment:手机"`
	Nickname  string  `gorm:"type:varchar(50);comment:昵称"`
	Avatar    string  `gorm:"type:varchar(255);default:/static/avatar/default.png;comment:头像"`
	Role      string  `gorm:"type:varchar(20);default:reader;comment:角色"`
	State     uint8   `gorm:"type:tinyint;default:1;index;comment:状态 1正常 2禁用 3封禁"`
	DeletedAt uint8   `gorm:"softDelete:flag;comment:软删除"` // 1=已删 0=未删
	CreatedAt int64   `gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt int64   `gorm:"autoUpdateTime;comment:更新时间"`
	LastIP    string  `gorm:"type:varchar(45);comment:最后登录IP"`
	LastAt    *int64  `gorm:"comment:最后登录时间"`
}

// 注册时候返回的数据
type UserResp struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	Token      string `json:"token"`       // 用户鉴权token
	UserID     int64  `json:"user_id"`     // 用户id
}
