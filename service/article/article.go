package article

import (
	"context"

	"blogs/dao"
)

func CreateOrUpdateArticle(ctx context.Context, article *dao.Article) error {
	if article.ArticleID == "" {
		return dao.CreateArticle(ctx, article)
	}

	updater := map[string]interface{}{
		"title":       article.Title,
		"summary":     article.Summary,
		"cover_image": article.CoverImage,
		"tags":        article.Tags,
		"content":     article.Content,
	}
	return dao.UpdateArticleByArticleID(ctx, article.ArticleID, updater)
}

func GetRecommendArticle(ctx context.Context, page, size int) (*GetRecommendArticleRsp, error) {
	articles, err := dao.GetArticleListByLikeCountSort(ctx, page, size)
	if err != nil {
		return nil, err
	}

	ret := &GetRecommendArticleRsp{}
	if len(articles) == 0 {
		return ret, nil
	}
	ret.HasMore = len(articles) == size

	userIds := make([]string, 0, len(articles))
	for _, v := range articles {
		userIds = append(userIds, v.UserID)
	}

	// 获取作者的头像+昵称
	users, err := dao.FindUserInfoByUserIDs(ctx, userIds)
	if err != nil {
		return nil, err
	}

	userMap := make(map[string]dao.UserInfo, len(users))
	for _, v := range users {
		userMap[v.UserID] = *v
	}

	ret.List = make([]ArtiInfo, 0, len(articles))
	for _, v := range articles {
		item := ArtiInfo{
			ID:           v.ID,
			ArticleID:    v.ArticleID,
			Title:        v.Title,
			Summary:      v.Summary,
			CoverImage:   v.CoverImage,
			Tags:         v.Tags,
			ViewCount:    v.ViewCount,
			LikeCount:    v.LikeCount,
			CommentCount: v.CommentCount,
			Nick:         userMap[v.UserID].Nick,
			Avatar:       userMap[v.UserID].Avatar,
		}
		ret.List = append(ret.List, item)
	}

	return ret, nil
}
