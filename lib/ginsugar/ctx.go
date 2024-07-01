package ginsugar

import (
	"context"

	"github.com/gin-gonic/gin"

	"blogs/lib/consts"
)

func Context(c *gin.Context) context.Context {
	ctx := c.Request.Context()
	userID := GetUserID
	ctx = context.WithValue(ctx, consts.CtxKeyUserID, userID)
	return ctx
}

func GetUserID(c *gin.Context) string {
	userID, _ := c.Get(consts.CtxKeyUserID)
	userIDStr, _ := userID.(string)
	return userIDStr
}
