package dao

import (
	"context"
	"fmt"
	"time"

	"blogs/lib/infra"
)

type Article struct {
	ID           int64     `json:"id" gorm:"column:id"`
	ArticleID    string    `json:"article_id" gorm:"column:article_id"`
	Title        string    `json:"title" gorm:"column:title"`
	Summary      string    `json:"summary" gorm:"column:summary"`
	Content      string    `json:"content" gorm:"column:content"`
	ViewCount    int64     `json:"view_count" gorm:"column:view_count"`
	LikeCount    int64     `json:"like_count" gorm:"column:like_count"`
	CommentCount int64     `json:"comment_count" gorm:"column:comment_count"`
	UserID       string    `json:"user_id" gorm:"column:user_id"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Article) TableName() string {
	return "articles"
}

func GetArticleByArticleID(ctx context.Context, articleID string) (result *Article, err error) {
	result = &Article{}
	err = infra.MysqlClient.WithContext(ctx).Model(&Article{}).Where("article_id = ?", articleID).Find(result).Error
	return
}

func GetArticleListByLikeCountSort(ctx context.Context, page int, size int) (result []*Article, err error) {
	err = infra.MysqlClient.WithContext(ctx).Model(&Article{}).
		Order("like_count DESC").
		Offset((page - 1) * size).
		Limit(size).Error
	return
}

func CreateArticle(ctx context.Context, article *Article) error {
	article.ArticleID = fmt.Sprintf("nid_%v", time.Now().UnixMilli())
	return infra.MysqlClient.WithContext(ctx).Create(article).Error
}

func UpdateArticleByArticleID(ctx context.Context, articleID string, updater map[string]interface{}) error {
	if len(updater) == 0 {
		return nil
	}
	return infra.MysqlClient.WithContext(ctx).Model(&Article{}).Where("article_id = ?", articleID).
		Updates(updater).Error
}
