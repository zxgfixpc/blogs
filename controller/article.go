package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"blogs/dao"
	"blogs/lib/ginsugar"
	"blogs/service/article"
)

func CreateOrUpdateArticle(c *gin.Context) {
	req := &dao.Article{}
	if err := c.ShouldBindJSON(req); err != nil {
		ginsugar.InputError(c, err)
		return
	}
	req.UserID = ginsugar.GetUserID(c)

	id, err := article.CreateOrUpdateArticle(ginsugar.Context(c), req)
	if err != nil {
		ginsugar.Fail(c, nil, err)
		return
	}

	ginsugar.Success(c, map[string]interface{}{
		"article_id": id,
	})
}

func GetRecommendArticles(c *gin.Context) {
	type Req struct {
		Page int `form:"page"`
		Size int `form:"size"`
	}
	var req Req
	if err := c.ShouldBind(&req); err != nil {
		ginsugar.InputError(c, err)
		return
	}

	list, err := article.GetRecommendArticle(ginsugar.Context(c), req.Page, req.Size)
	if err != nil {
		ginsugar.Fail(c, nil, err)
		return
	}

	ginsugar.Success(c, list)
}

func GetArticleByID(c *gin.Context) {
	articleID, ok := c.GetQuery("article_id")
	if !ok {
		ginsugar.InputError(c, fmt.Errorf("not article_id"))
	}

	articleInfo, err := article.GetArticleByID(ginsugar.Context(c), articleID)
	if err != nil {
		ginsugar.Fail(c, nil, err)
		return
	}

	ginsugar.Success(c, articleInfo)
}
