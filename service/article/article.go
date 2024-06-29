package article

import (
	"context"

	"blogs/dao"
)

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

	usernames := make([]string, 0, len(articles))
	for _, v := range articles {
		usernames = append(usernames, v.Username)
	}

	// 获取作者的头像+昵称
	users, err := dao.FindUserInfoByUsernames(ctx, usernames)
	if err != nil {
		return nil, err
	}

	userMap := make(map[string]dao.User, len(users))
	for _, v := range users {
		userMap[v.Username] = *v
	}

	ret.List = make([]ArtiInfo, 0, len(articles))
	for _, v := range articles {
		item := ArtiInfo{
			ID:           v.ID,
			NID:          v.NID,
			Title:        v.Title,
			Summary:      v.Summary,
			ViewCount:    v.ViewCount,
			LikeCount:    v.LikeCount,
			CommentCount: v.CommentCount,
			Nick:         userMap[v.Username].Nick,
			Avatar:       userMap[v.Username].Avatar,
		}
		ret.List = append(ret.List, item)
	}

	return ret, nil
}
