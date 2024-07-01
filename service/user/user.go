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
