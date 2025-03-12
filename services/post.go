package services

import (
	"blog/libs"
	"blog/models"
)

func CreatePost(title string, content string, author models.User) (*models.Post, error) {
	post := &models.Post{
		Title:   title,
		Content: content,
		Author:  author,
	}

	err := models.CreatePost(libs.DB, post)
	return post, err
}

func GetPost(id int) (*models.Post, error) {
	post, err := models.GetPost(libs.DB, id)
	return &post, err
}

func GetAllPosts() ([]models.Post, error) {
	posts, err := models.GetAllPosts(libs.DB)
	return posts, err
}

func UpdatePost(id int, title string, content string) (*models.Post, error) {
	post, err := models.GetPost(libs.DB, id)
	if err != nil {
		return nil, err
	}

	post.Title = title
	post.Content = content

	err = models.UpdatePost(libs.DB, &post)
	return &post, err
}

func DeletePost(id int) error {
	post, err := models.GetPost(libs.DB, id)
	if err != nil {
		return err
	}

	err = models.DeletePost(libs.DB, &post)
	return err
}
