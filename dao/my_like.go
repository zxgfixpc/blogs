package dao

import (
	"context"

	"gorm.io/gorm"
)

// MyLikeArticles 结构体
type MyLikeArticles struct {
	Base
	UserID    string `json:"user_id" gorm:"column:user_id"`
	ArticleID string `json:"article_id" gorm:"column:article_id"`
	Status    int8   `json:"status" gorm:"column:status"`
}

func (MyLikeArticles) TableName() string {
	return "my_like_articles"
}

// GetTheArticleLikeUsers 文章喜欢的人
func GetTheArticleLikeUsers(ctx context.Context, articleID string) (userIDs []string, err error) {
	err = defaultDB(ctx).Model(&Follows{}).
		Select("user_id").
		Where("article_id = ?", articleID).
		Order("created_at ASC").
		Find(&userIDs).Error
	return
}

// GetMyLikeArticleList 我喜欢的文章列表
func GetMyLikeArticleList(ctx context.Context, userID string) (articleIDs []string, err error) {
	err = defaultDB(ctx).Model(&Follows{}).
		Select("article_id").
		Where("user_id = ?", userID).
		Order("created_at ASC").
		Find(&articleIDs).Error
	return
}

func CreateMyLikeArticles(ctx context.Context, userID, articleID string) error {
	return defaultDB(ctx).Create(&MyLikeArticles{
		UserID:    userID,
		ArticleID: articleID,
	}).Error
}

func CancelMyLikeArticles(ctx context.Context, userID, articleID string) error {
	return defaultDB(ctx).Model(&MyLikeArticles{}).
		Where("user_id = ?", userID).
		Where("article_id = ?", articleID).
		Updates(map[string]interface{}{"deleted_id": gorm.Expr(`id`)}).
		Error
}
