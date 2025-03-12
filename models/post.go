package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Title    string `json:"title"`
	Content  string `json:"content"`
	MediaUrl string `json:"media_url"`
	User     User   `json:"user" gorm:"foreignKey:UserID"`

	Views    []User    `json:"views" gorm:"many2many:post_views;"`
	Likes    []User    `json:"likes" gorm:"many2many:post_likes;"`
	Comments []Comment `json:"comments" gorm:"foreignKey:PostID"`

	UserID uint `json:"user_id"`
}

// =====================
// CRUD functions
// =====================

func GetAllPosts(db *gorm.DB) ([]Post, error) {
	var posts []Post
	result := db.Find(&posts)
	return posts, result.Error
}

func GetPost(db *gorm.DB, id uint) (Post, error) {
	var post Post
	result := db.First(&post, id)
	return post, result.Error
}

func CreatePost(db *gorm.DB, post *Post) error {
	result := db.Create(post)
	return result.Error
}

func UpdatePost(db *gorm.DB, post *Post) error {
	result := db.Save(post)
	return result.Error
}

func DeletePost(db *gorm.DB, post *Post) error {
	result := db.Delete(post)
	return result.Error
}
