package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Title   string `json:"title"`
	Content string `json:"content"`

	Likes    []User    `json:"likes" gorm:"many2many:post_likes;"`
	Comments []Comment `json:"comments" gorm:"foreignKey:PostID"`
	Author   User      `json:"author" gorm:"foreignKey:UserID"`

	UserID uint `json:"-"`
}

type CreateOrUpdatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// =====================
// CRUD functions
// =====================

func GetAllPosts(db *gorm.DB) ([]Post, error) {
	var posts []Post
	result := db.Preload("Likes").Preload("Comments").Find(&posts)
	return posts, result.Error
}

func GetPost(db *gorm.DB, id int) (Post, error) {
	var post Post
	result := db.Preload("Likes").Preload("Comments").First(&post, id)
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

func LikePost(db *gorm.DB, post *Post, user *User) error {
	err := db.Model(post).Association("Likes").Append(user)
	return err
}

func UnlikePost(db *gorm.DB, post *Post, user *User) error {
	err := db.Model(post).Association("Likes").Delete(user)
	return err
}
