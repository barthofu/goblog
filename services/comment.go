package services

import (
	"blog/libs"
	"blog/models"
)

func CreateComment(userID uint, content string) (*models.Comment, error) {
	comment := &models.Comment{
		UserID:  userID,
		Content: content,
	}

	err := models.CreateComment(libs.DB, comment)
	return comment, err
}

func GetComment(id uint) (*models.Comment, error) {
	comment, err := models.GetComment(libs.DB, id)
	return &comment, err
}

func GetAllComments() ([]models.Comment, error) {
	comments, err := models.GetAllComments(libs.DB)
	return comments, err
}

func UpdateComment(id uint, content string) (*models.Comment, error) {
	comment, err := models.GetComment(libs.DB, id)
	if err != nil {
		return nil, err
	}

	comment.Content = content
	err = models.UpdateComment(libs.DB, &comment)
	return &comment, err
}

func DeleteComment(id uint) error {
	comment, err := models.GetComment(libs.DB, id)
	if err != nil {
		return err
	}

	err = models.DeleteComment(libs.DB, &comment)
	return err
}
