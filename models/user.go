package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Email     string `json:"email" gorm:"not null;uniqueIndex"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	Posts    []Post    `json:"-"`
	Comments []Comment `json:"-"`
	Likes    []Post    `json:"-"`
}

type CreateOrUpdateUserInput struct {
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// =====================
// CRUD functions
// =====================

func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	result := db.Find(&users)
	return users, result.Error
}

func GetUser(db *gorm.DB, id int) (User, error) {
	var user User
	result := db.First(&user, id)
	return user, result.Error
}

func CreateUser(db *gorm.DB, user *User) error {
	result := db.Create(user)
	return result.Error
}

func UpdateUser(db *gorm.DB, user *User) error {
	result := db.Save(user)
	return result.Error
}

func DeleteUser(db *gorm.DB, user *User) error {
	result := db.Delete(user)
	return result.Error
}
