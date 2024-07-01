package dao

import (
	"time"
)

type Follows struct {
	ID             int64     `json:"id" gorm:"column:id"`
	UserID         string    `json:"user_id" gorm:"column:user_id"`
	FollowerUserID string    `json:"follower_user_id" gorm:"column:follower_user_id"`
	Status         int8      `json:"status" gorm:"column:status"`
	CreatedAt      time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"column:updated_at"`
}
