package dao

import (
	"context"
	"fmt"
	"time"

	"blogs/lib/infra"
)

type Article struct {
	ID           int64  `json:"id" gorm:"column:id"`
	NID          string `json:"nid" gorm:"column:nid"`
	Title        string `json:"title" gorm:"column:title"`
	Summary      string `json:"summary" gorm:"column:summary"`
	Content      string `json:"content" gorm:"column:content"`
	ViewCount    int64  `json:"view_count" gorm:"column:view_count"`
	LikeCount    int64  `json:"like_count" gorm:"column:like_count"`
	CommentCount int64  `json:"comment_count" gorm:"column:comment_count"`
	Username     string `json:"username" gorm:"column:username"`
}

func (Article) TableName() string {
	return "articles"
}

func GetArticleByNID(ctx context.Context, nid string) (result *Article, err error) {
	result = &Article{}
	err = infra.MysqlClient.WithContext(ctx).Model(&Article{}).Where("nid = ?", nid).Find(result).Error
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
	article.NID = fmt.Sprintf("nid_%v", time.Now().UnixMilli())
	return infra.MysqlClient.WithContext(ctx).Create(article).Error
}

func UpdateArticle(ctx context.Context, nid string, updater map[string]interface{}) error {
	if len(updater) == 0 {
		return nil
	}
	return infra.MysqlClient.WithContext(ctx).Model(&Article{}).Where("nid = ?", nid).
		Updates(updater).Error
}
