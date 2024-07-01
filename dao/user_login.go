package dao

import (
	"context"
)

type UserLogin struct {
	Base
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
	err = defaultDB(ctx).Model(&UserLogin{}).Where("user_id = ?", userID).Find(result).Error
	return
}

func GetUserLoginBySessionID(ctx context.Context, sessionID string) (result *UserLogin, err error) {
	result = &UserLogin{}
	err = defaultDB(ctx).Model(&UserLogin{}).Where("session_id = ?", sessionID).Find(result).Error
	return
}

func CreateUserLogin(ctx context.Context, userLogin *UserLogin) error {
	return defaultDB(ctx).Create(userLogin).Error
}

func UpdateUserLogin(ctx context.Context, userID string, updater map[string]interface{}) error {
	if len(updater) == 0 {
		return nil
	}
	return defaultDB(ctx).Model(&UserLogin{}).Where("user_id = ?", userID).
		Updates(updater).Error
}
