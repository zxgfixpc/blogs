package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"blogs/dao"
	"blogs/lib/consts"
	"blogs/lib/ginsugar"
)

func LoginMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := c.Request.Cookie(consts.CookieKeyUserID)
		if err != nil || userID == nil || userID.Value == "" {
			ginsugar.FailNotLogin(c, fmt.Errorf("用户未登录"))
			return
		}

		sessionID, err := c.Request.Cookie(consts.CookieKeySessionID)
		if err != nil || sessionID == nil || sessionID.Value == "" {
			ginsugar.FailNotLogin(c, fmt.Errorf("用户未登录"))
			return
		}

		loginInfo, err := dao.GetUserLoginByUserID(context.TODO(), userID.Value)
		if err != nil || loginInfo == nil || loginInfo.SessionID == "" || loginInfo.SessionID != sessionID.Value {
			ginsugar.FailNotLogin(c, fmt.Errorf("请重新登录"))
			return
		}
		if time.Now().Unix() > loginInfo.SessionExpr {
			ginsugar.FailNotLogin(c, fmt.Errorf("登录过期，请重新登录"))
			return
		}

		// TODO 刷新登录cookie

		c.Set(consts.CtxKeyUserID, loginInfo.UserID)

		c.Next()
	}
}
