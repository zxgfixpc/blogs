package dao

import (
	"time"
)

// ArticleComments 结构体
type ArticleComments struct {
	ID        int64     `json:"id" gorm:"column:id"`
	ArticleID string    `json:"article_id" gorm:"column:article_id"`
	UserID    string    `json:"user_id" gorm:"column:user_id"`
	ParentID  int64     `json:"parent_id" gorm:"column:parent_id"`
	Content   string    `json:"content" gorm:"column:content"`
	Status    int8      `json:"status" gorm:"column:status"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}
