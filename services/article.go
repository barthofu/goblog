package services

import (
	"blog/libs"
	"blog/models"
)

func CreateArticle(title string, content string, author models.User) (*models.Article, error) {
	article := &models.Article{
		Title:   title,
		Content: content,
		Author:  author,
	}

	err := models.CreateArticle(libs.DB, article)
	return article, err
}

func GetArticle(id int) (*models.Article, error) {
	article, err := models.GetArticle(libs.DB, id)
	return &article, err
}

func GetAllArticles() ([]models.Article, error) {
	articles, err := models.GetAllArticles(libs.DB)
	return articles, err
}

func UpdateArticle(id int, title string, content string) (*models.Article, error) {
	article, err := models.GetArticle(libs.DB, id)
	if err != nil {
		return nil, err
	}

	article.Title = title
	article.Content = content

	err = models.UpdateArticle(libs.DB, &article)
	return &article, err
}

func DeleteArticle(id int) error {
	article, err := models.GetArticle(libs.DB, id)
	if err != nil {
		return err
	}

	err = models.DeleteArticle(libs.DB, &article)
	return err
}

func LikeArticle(id int, user models.User) error {
	article, err := models.GetArticle(libs.DB, id)
	if err != nil {
		return err
	}

	err = models.LikeArticle(libs.DB, &article, &user)
	return err
}

func UnlikeArticle(id int, user models.User) error {
	article, err := models.GetArticle(libs.DB, id)
	if err != nil {
		return err
	}

	err = models.UnlikeArticle(libs.DB, &article, &user)
	return err
}
