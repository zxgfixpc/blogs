package controller

import (
	"github.com/gin-gonic/gin"

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

	ginsugar.Success(c, map[string]interface{}{
		"session_id": sessionID,
	})
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

func Register(c *gin.Context) {
	req := &user.LoginReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		ginsugar.InputError(c, err)
		return
	}

	sessionID, err := user.Register(ginsugar.Context(c), req.UserID, req.Password)
	if err != nil {
		ginsugar.Fail(c, nil, err)
		return
	}

	ginsugar.Success(c, map[string]interface{}{
		"session_id": sessionID,
	})
}
