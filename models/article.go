package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Title   string `json:"title"`
	Content string `json:"content"`

	Likes    []User    `json:"likes" gorm:"many2many:article_likes;"`
	Comments []Comment `json:"comments" gorm:"foreignKey:ArticleID"`
	Author   User      `json:"author" gorm:"foreignKey:UserID"`

	UserID uint `json:"-"`
}

type CreateOrUpdateArticleInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// =====================
// CRUD functions
// =====================

func GetAllArticles(db *gorm.DB) ([]Article, error) {
	var articles []Article
	result := db.Preload("Likes").Preload("Comments").Find(&articles)
	return articles, result.Error
}

func GetArticle(db *gorm.DB, id int) (Article, error) {
	var article Article
	result := db.Preload("Likes").Preload("Comments").First(&article, id)
	return article, result.Error
}

func CreateArticle(db *gorm.DB, article *Article) error {
	result := db.Create(article)
	return result.Error
}

func UpdateArticle(db *gorm.DB, article *Article) error {
	result := db.Save(article)
	return result.Error
}

func DeleteArticle(db *gorm.DB, article *Article) error {
	result := db.Delete(article)
	return result.Error
}

func LikeArticle(db *gorm.DB, article *Article, user *User) error {
	err := db.Model(article).Association("Likes").Append(user)
	return err
}

func UnlikeArticle(db *gorm.DB, article *Article, user *User) error {
	err := db.Model(article).Association("Likes").Delete(user)
	return err
}
