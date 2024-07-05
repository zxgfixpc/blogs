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
	sessionID := getSessionID(userID)
	sessionExpr := getSessionExpr()
	err = dao.UpdateUserLogin(ctx, userID, map[string]interface{}{
		"session_id":   sessionID,
		"session_expr": sessionExpr,
	})
	return sessionID, err
}

func Exit(ctx context.Context, userID string) error {
	return dao.UpdateUserLogin(ctx, userID, map[string]interface{}{
		"session_id":   "",
		"session_expr": 0,
	})
}

func Register(ctx context.Context, userID, password string) (string, error) {
	var sessionID string
	err := dao.Trans(ctx, func(transCtx context.Context) error {
		userInfo, err := dao.GetUserLoginByUserID(transCtx, userID)
		if err != nil {
			return err
		}
		if userInfo != nil && userInfo.UserID != "" {
			return fmt.Errorf("用户已注册")
		}

		sessionID = getSessionID(userID)
		err = dao.CreateUserLogin(transCtx, &dao.UserLogin{
			UserID:      userID,
			Password:    password,
			SessionID:   sessionID,
			SessionExpr: getSessionExpr(),
		})
		if err != nil {
			return err
		}

		err = dao.CreateUserInfo(transCtx, &dao.UserInfo{
			UserID: userID,
			Nick:   fmt.Sprintf("Note_%v", time.Now().Unix()),
		})
		if err != nil {
			return err
		}
		return nil
	})

	return sessionID, err
}

func getSessionID(userID string) string {
	sessionID := time.Now().Unix()
	return fmt.Sprintf("%v%v", userID, sessionID)
}

func getSessionExpr() int64 {
	return time.Now().Add(2 * time.Hour).Unix()
}
