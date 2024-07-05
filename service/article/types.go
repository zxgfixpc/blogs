package article

import (
	"blogs/dao"
)

type ArtiInfo struct {
	ID           int64       `json:"id"`
	ArticleID    string      `json:"article_id"`
	Title        string      `json:"title"`
	Summary      string      `json:"summary"`
	CoverImage   string      `json:"cover_image"`
	Tags         dao.Strings `json:"tags"`
	ViewCount    int64       `json:"view_count"`
	LikeCount    int64       `json:"like_count"`
	CommentCount int64       `json:"comment_count"`
	Nick         string      `json:"nick"`
	Avatar       string      `json:"avatar"`
	CreatedAt    string      `json:"created_at"`
}

type GetRecommendArticleRsp struct {
	HasMore bool       `json:"has_more"`
	List    []ArtiInfo `json:"list"`
}
