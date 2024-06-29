package article

type ArtiInfo struct {
	ID           int64  `json:"id"`
	NID          string `json:"nid"`
	Title        string `json:"title"`
	Summary      string `json:"summary"`
	ViewCount    int64  `json:"view_count"`
	LikeCount    int64  `json:"like_count"`
	CommentCount int64  `json:"comment_count"`
	Nick         string `json:"nick"`
	Avatar       string `json:"avatar"`
}

type GetRecommendArticleRsp struct {
	HasMore bool       `json:"has_more"`
	List    []ArtiInfo `json:"list"`
}
