package dao

import (
	"time"
)

// MyLikeArticles 结构体
type MyLikeArticles struct {
	ID        uint64    `json:"id" gorm:"column:id"`
	UserID    string    `json:"user_id" gorm:"column:user_id"`
	ArticleID string    `json:"article_id" gorm:"column:article_id"`
	Status    int8      `json:"status" gorm:"column:status"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
} 
