package dao

import (
	"context"
)

// ArticleComments 结构体
type ArticleComments struct {
	Base
	ArticleID string `json:"article_id" gorm:"column:article_id"`
	UserID    string `json:"user_id" gorm:"column:user_id"`
	ParentID  int64  `json:"parent_id" gorm:"column:parent_id"`
	Content   string `json:"content" gorm:"column:content"`
	Status    int8   `json:"status" gorm:"column:status"`
}

func (ArticleComments) TableName() string {
	return "article_comments"
}

func GetArticleCommentsByArticleID(ctx context.Context, articleID string) (result []*ArticleComments, err error) {
	err = defaultDB(ctx).WithContext(ctx).Model(&Article{}).
		Where("article_id = ?", articleID).
		Order("created_at ASC").
		Find(&result).Error
	return
}

func CreateArticleComment(ctx context.Context, articleID string, parentID int64, comment string) error {
	return defaultDB(ctx).WithContext(ctx).Create(&ArticleComments{
		ArticleID: articleID,
		ParentID:  parentID,
		Content:   comment,
	}).Error
}
