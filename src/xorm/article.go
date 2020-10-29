package xorm

import (
	"errors"

	"lottery_backend/src/utils"
	"lottery_backend/src/xorm/model"
)

// InsertAndUpdateArticle: insert article or update article for user
func InsertAndUpdateArticle(article *model.Article) (string, error) {
	if article == nil || article.UserId == "" || article.Text == "" {
		return "", errors.New("insert or update article error")
	}

	db := GetDB()

	sess := db.NewSession()
	defer sess.Close()
	err := sess.Begin()
	if err != nil {
		return "", err
	}
	oldArticle := &model.Article{
		UserId: article.UserId,
	}
	exist, err := db.Where("user_id = ?", article.UserId).Get(oldArticle)
	if err != nil {
		return "", err
	}
	if exist {
		article.Id = oldArticle.Id
		_, err = db.AllCols().Where("id=?", article.Id).Update(article)
	} else {
		article.Id = utils.NewId()
		_, err = db.Insert(article)
		if err != nil {
			return "", err
		}
		err = InserPermission(&model.Permission{
			UserId:    article.UserId,
			Permitted: 1,
		})
	}
	err = sess.Commit()
	if err != nil {
		return "", err
	}
	return article.Id, err
}

// ListUserArticle: list all article
// Todo: join user
func ListUserArticle(offset, limit int) ([]model.Article, error) {
	db := GetDB()
	data := make([]model.Article, 0)
	err := db.Limit(limit, offset).Find(&data)
	return data, err
}
