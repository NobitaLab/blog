package models

import (
	"time"
)

// User 用户模型
// @Description 用户数据结构
// @Param id 唯一标识符
// @Param username 用户名
// @Param password 密码（哈希后的）
// @Param email 邮箱
// @Param role 用户角色（admin/user）
// @Param createdAt 创建时间
// @Param updatedAt 更新时间
type User struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Username  string    `gorm:"size:100;not null;unique" json:"username"`
	Password  string    `gorm:"size:255;not null" json:"password,omitempty"` // 不返回密码字段
	Email     *string   `gorm:"size:255;uniqueIndex:idx_unique_email_null" json:"email"`
	Role      string    `gorm:"size:20;not null;default:'user'" json:"role"` // admin 或 user
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// TableName 设置表名
func (User) TableName() string {
	return "users"
}