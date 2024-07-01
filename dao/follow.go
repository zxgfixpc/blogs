package dao

import (
	"context"

	"gorm.io/gorm"
)

type Follows struct {
	Base
	UserID         string `json:"user_id" gorm:"column:user_id"`
	FollowerUserID string `json:"follower_user_id" gorm:"column:follower_user_id"`
	Status         int8   `json:"status" gorm:"column:status"`
}

func (Follows) TableName() string {
	return "follows"
}

// GetFollowMeUser 关注我的人
func GetFollowMeUser(ctx context.Context, userID string) (followers []string, err error) {
	err = defaultDB(ctx).Model(&Follows{}).
		Select("follower_user_id").
		Where("user_id = ?", userID).
		Order("created_at ASC").
		Find(&followers).Error
	return
}

// GetIFollowUser 我关注的人
func GetIFollowUser(ctx context.Context, follower string) (users []string, err error) {
	err = defaultDB(ctx).Model(&Follows{}).
		Select("user_id").
		Where("follower_user_id = ?", follower).
		Order("created_at ASC").
		Find(&users).Error
	return
}

func CreateFollow(ctx context.Context, userID, followerID string) error {
	return defaultDB(ctx).Create(&Follows{
		UserID:         userID,
		FollowerUserID: followerID,
	}).Error
}

func CancelFollow(ctx context.Context, userID, followerID string) error {
	return defaultDB(ctx).Model(&Follows{}).
		Where("user_id = ?", userID).
		Where("follower_user_id = ?", followerID).
		Updates(map[string]interface{}{"deleted_id": gorm.Expr(`id`)}).
		Error
}
