package controller

import (
	"github.com/gin-gonic/gin"

	"blogs/lib/consts"
	"blogs/lib/ginsugar"
	"blogs/service/user"
)

func Login(c *gin.Context) {
	req := &user.LoginReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		ginsugar.InputError(c, err)
		return
	}

	sessionID, err := user.Login(ginsugar.Context(c), req.UserID, req.Password)
	if err != nil {
		ginsugar.Fail(c, nil, err)
		return
	}
	c.SetCookie(consts.CookieKeySessionID, sessionID, consts.CookieKeySessionExpr*3600, "", "localhost", false, true)

	ginsugar.Success(c, nil)
}

func Exit(c *gin.Context) {
	userID := ginsugar.GetUserID(c)
	if userID == "" {
		ginsugar.Success(c, nil)
		return
	}
	err := user.Exit(ginsugar.Context(c), userID)
	if err != nil {
		ginsugar.Fail(c, nil, err)
		return
	}
	ginsugar.Success(c, nil)
}
