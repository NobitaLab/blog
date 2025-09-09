package models

import (
	"time"
)

// Blog 博客文章模型
// @Description 博客文章数据结构
// @Param id 唯一标识符
// @Param title 标题
// @Param content 内容（Markdown格式）
// @Param author 作者
// @Param createdAt 创建时间（ISO 8601格式）
// @Param updatedAt 更新时间（ISO 8601格式）
type Blog struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Title     string    `gorm:"size:255;not null" json:"title"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	Author    string    `gorm:"size:100" json:"author"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// TableName 设置表名
func (Blog) TableName() string {
	return "blogs"
}