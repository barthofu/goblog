package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Content string `json:"content"`
	Post    Post   `json:"post" gorm:"foreignKey:PostID"`
	User    User   `json:"user" gorm:"foreignKey:UserID"`

	PostID uint `json:"post_id"`
	UserID uint `json:"user_id"`
}

// =====================
// Repository functions
// =====================

func GetAllComments(db *gorm.DB) ([]Comment, error) {
	var comments []Comment
	result := db.Find(&comments)
	return comments, result.Error
}

func GetComment(db *gorm.DB, id uint) (Comment, error) {
	var comment Comment
	result := db.First(&comment, id)
	return comment, result.Error
}

func CreateComment(db *gorm.DB, comment *Comment) error {
	result := db.Create(comment)
	return result.Error
}

func UpdateComment(db *gorm.DB, comment *Comment) error {
	result := db.Save(comment)
	return result.Error
}

func DeleteComment(db *gorm.DB, comment *Comment) error {
	result := db.Delete(comment)
	return result.Error
}
