package services

import (
	"blog/libs"
	"blog/models"
)

func CreateComment(content string, articleId int, user models.User) (*models.Comment, error) {
	var article, articleErr = GetArticle(articleId)
	if articleErr != nil {
		return nil, articleErr
	}

	comment := &models.Comment{
		User:    user,
		Content: content,
		Article: *article,
	}

	err := models.CreateComment(libs.DB, comment)
	return comment, err
}
