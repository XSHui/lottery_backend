package xorm

import (
	"errors"
	"lottery_backend/src/xorm/model"
)

// InsertArticle: insert article for user
func InsertArticle(article *model.Article) error {
	if article == nil || article.Id == "" || article.UserId == "" || article.Text == "" {
		return errors.New("insert article error")
	}

	db := GetDB()
	_, err := db.Insert(article)
	return err
}

// ListUserArticle: list all article
// Todo: join user
func ListUserArticle(offset, limit int)([]model.Article, error) {
	db := GetDB()
	data := make([]model.Article, 0)
	err := db.Limit(limit, offset).Find(&data)
	return data, err
}

