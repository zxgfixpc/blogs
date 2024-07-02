package user

import (
	"context"
	"fmt"
	"time"

	"blogs/dao"
)

func Login(ctx context.Context, userID, password string) (string, error) {
	loginInfo, err := dao.GetUserLoginByUserID(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("get user login err:%v", err)
	}

	if loginInfo == nil || loginInfo.UserID == "" {
		return "", fmt.Errorf("账户不存在")
	}

	if loginInfo.Password != password {
		return "", fmt.Errorf("密码错误")
	}

	// sessionID
	sessionID := time.Now().Unix()
	sessionIDStr := fmt.Sprintf("%v%v", userID, sessionID)
	sessionExpr := time.Now().Add(1 * time.Hour).Unix()
	err = dao.UpdateUserLogin(ctx, userID, map[string]interface{}{
		"session_id":   sessionIDStr,
		"session_expr": sessionExpr,
	})
	return sessionIDStr, err
}

func Exit(ctx context.Context, userID string) error {
	return dao.UpdateUserLogin(ctx, userID, map[string]interface{}{
		"session_id":   "",
		"session_expr": 0,
	})
}

func Register(ctx context.Context, userID, password string) error {
	userInfo, err := dao.GetUserLoginByUserID(ctx, userID)
	if err != nil {
		return err
	}
	if userInfo != nil && userInfo.UserID != "" {
		return fmt.Errorf("用户已注册")
	}

	err = dao.CreateUserLogin(ctx, &dao.UserLogin{
		UserID:   userID,
		Password: password,
	})
	if err != nil {
		return err
	}

	err = dao.CreateUserInfo(ctx, &dao.UserInfo{
		UserID: userID,
		Nick:   fmt.Sprintf("Note_%v", time.Now().Unix()),
	})
	if err != nil {
		return err
	}

	return nil
}
