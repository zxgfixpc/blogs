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
		cookie, err := c.Request.Cookie(consts.CookieKeySessionID)
		if err != nil || cookie == nil || cookie.Value == "" {
			ginsugar.FailNotLogin(c, fmt.Errorf("用户未登录"))
			return
		}

		loginInfo, err := dao.GetUserLoginBySessionID(context.TODO(), cookie.Value)
		if err != nil || loginInfo == nil || loginInfo.SessionID == "" {
			ginsugar.FailNotLogin(c, fmt.Errorf("系统错误，请重新登录"))
		}
		if time.Now().Unix() > loginInfo.SessionExpr {
			ginsugar.FailNotLogin(c, fmt.Errorf("登录过期，请重新登录"))
		}

		// TODO 刷新登录cookie

		c.Set(consts.CtxKeyUserID, loginInfo.UserID)

		c.Next()
	}
}
