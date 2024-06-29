package dao

import (
	"context"

	"blogs/lib/infra"
)

type User struct {
	ID       int64  `json:"id" gorm:"column:id"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
	Nick     string `json:"nick" gorm:"column:nick"`
	Avatar   string `json:"avatar" gorm:"column:avatar"`
}

func (User) TableName() string {
	return "users"
}

func GetUserInfoByUserName(ctx context.Context, username string) (result *User, err error) {
	result = &User{}
	err = infra.MysqlClient.WithContext(ctx).Model(&User{}).Where("username = ?", username).Find(result).Error
	return
}

func CreateUser(ctx context.Context, userInfo *User) error {
	return infra.MysqlClient.WithContext(ctx).Create(userInfo).Error
}

func UpdateUserInfo(ctx context.Context, username string, updater map[string]interface{}) error {
	if len(updater) == 0 {
		return nil
	}
	return infra.MysqlClient.WithContext(ctx).Model(&User{}).Where("username = ?", username).
		Updates(updater).Error
}
