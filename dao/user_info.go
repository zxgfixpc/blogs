package dao

import (
	"context"
)

type UserInfo struct {
	Base
	UserID string `json:"user_id" gorm:"column:user_id"`
	Nick   string `json:"nick" gorm:"column:nick"`
	Avatar string `json:"avatar" gorm:"column:avatar"`
}

func (UserInfo) TableName() string {
	return "user_info"
}

func GetUserInfoByUserID(ctx context.Context, userID string) (result *UserInfo, err error) {
	result = &UserInfo{}
	err = defaultDB(ctx).Model(&UserInfo{}).Where("user_id = ?", userID).Find(result).Error
	return
}

func FindUserInfoByUserIDs(ctx context.Context, userIDs []string) (result []*UserInfo, err error) {
	if len(userIDs) == 0 {
		return nil, nil
	}
	err = defaultDB(ctx).Model(&UserInfo{}).Where("user_id in (?)", userIDs).Find(&result).Error
	return
}

func CreateUserInfo(ctx context.Context, userInfo *UserInfo) error {
	return defaultDB(ctx).Create(userInfo).Error
}

func UpdateUserInfo(ctx context.Context, userID string, updater map[string]interface{}) error {
	if len(updater) == 0 {
		return nil
	}
	return defaultDB(ctx).Model(&UserInfo{}).Where("user_id = ?", userID).
		Updates(updater).Error
}
