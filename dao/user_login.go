package dao

import (
	"context"

	"blogs/lib/infra"
)

type UserLogin struct {
	ID          int64  `json:"id" gorm:"column:id"`
	UserID      string `json:"user_id" gorm:"column:user_id"`
	Password    string `json:"password" gorm:"column:password"`
	SessionID   string `json:"session_id" gorm:"column:session_id"`
	SessionExpr int64  `json:"session_expr" gorm:"column:session_expr"`
}

func (UserLogin) TableName() string {
	return "user_login"
}

func GetUserLoginByUserID(ctx context.Context, userID string) (result *UserLogin, err error) {
	result = &UserLogin{}
	err = infra.MysqlClient.WithContext(ctx).Model(&UserLogin{}).Where("user_id = ?", userID).Find(result).Error
	return
}

func CreateUserLogin(ctx context.Context, userLogin *UserLogin) error {
	return infra.MysqlClient.WithContext(ctx).Create(userLogin).Error
}

func UpdateUserLogin(ctx context.Context, userID string, updater map[string]interface{}) error {
	if len(updater) == 0 {
		return nil
	}
	return infra.MysqlClient.WithContext(ctx).Model(&UserLogin{}).Where("user_id = ?", userID).
		Updates(updater).Error
}
