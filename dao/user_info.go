package dao

import (
	"context"

	"blogs/lib/infra"
)

type UserInfo struct {
	ID     int64  `json:"id" gorm:"column:id"`
	UserID string `json:"user_id" gorm:"column:user_id"`
	Nick   string `json:"nick" gorm:"column:nick"`
	Avatar string `json:"avatar" gorm:"column:avatar"`
}

func (UserInfo) TableName() string {
	return "user_info"
}

func GetUserInfoByUserID(ctx context.Context, userID string) (result *UserInfo, err error) {
	result = &UserInfo{}
	err = infra.MysqlClient.WithContext(ctx).Model(&UserInfo{}).Where("user_id = ?", userID).Find(result).Error
	return
}

func FindUserInfoByUserIDs(ctx context.Context, userIDs []string) (result []*UserInfo, err error) {
	if len(userIDs) == 0 {
		return nil, nil
	}
	err = infra.MysqlClient.WithContext(ctx).Model(&UserInfo{}).Where("user_id in (?)", userIDs).Find(&result).Error
	return
}

func CreateUserInfo(ctx context.Context, userInfo *UserInfo) error {
	return infra.MysqlClient.WithContext(ctx).Create(userInfo).Error
}

func UpdateUserInfo(ctx context.Context, userID string, updater map[string]interface{}) error {
	if len(updater) == 0 {
		return nil
	}
	return infra.MysqlClient.WithContext(ctx).Model(&UserInfo{}).Where("user_id = ?", userID).
		Updates(updater).Error
}
