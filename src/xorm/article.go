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
