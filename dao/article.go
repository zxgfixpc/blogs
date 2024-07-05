package dao

import (
	"context"
)

type Article struct {
	Base
	ArticleID    string  `json:"article_id" gorm:"column:article_id"`
	Title        string  `json:"title" gorm:"column:title"`
	Summary      string  `json:"summary" gorm:"column:summary"`
	CoverImage   string  `json:"cover_image" gorm:"column:cover_image"`
	Tags         Strings `json:"tags" gorm:"column:tags"`
	Content      string  `json:"content" gorm:"column:content"`
	ViewCount    int64   `json:"view_count" gorm:"column:view_count"`
	LikeCount    int64   `json:"like_count" gorm:"column:like_count"`
	CommentCount int64   `json:"comment_count" gorm:"column:comment_count"`
	UserID       string  `json:"user_id" gorm:"column:user_id"`
	Status       int8    `json:"status" gorm:"column:status"`
}

func (Article) TableName() string {
	return "articles"
}

func GetArticleByArticleID(ctx context.Context, articleID string) (result *Article, err error) {
	result = &Article{}
	err = defaultDB(ctx).Model(&Article{}).Where("article_id = ?", articleID).Find(result).Error
	return
}

func GetArticleListByLikeCountSort(ctx context.Context, page int, size int) (result []*Article, err error) {
	err = defaultDB(ctx).Model(&Article{}).
		Order("like_count DESC").
		Offset((page - 1) * size).
		Limit(size).
		Find(&result).
		Error
	return
}

func CreateArticle(ctx context.Context, article *Article) error {
	return defaultDB(ctx).Create(article).Error
}

func UpdateArticleByArticleID(ctx context.Context, articleID string, updater map[string]interface{}) error {
	if len(updater) == 0 {
		return nil
	}
	return defaultDB(ctx).Model(&Article{}).Where("article_id = ?", articleID).
		Updates(updater).Error
}
