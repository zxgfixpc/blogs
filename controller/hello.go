package controller

import (
	"blogs/lib/ginsugar"
	"blogs/service"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	rlt, err := service.Hello(ginsugar.Context(c))
	if err != nil {
		ginsugar.Fail(c, rlt, err)
		return
	}
	ginsugar.Success(c, rlt)
}
