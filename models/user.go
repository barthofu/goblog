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

	Following []*User   `json:"following" gorm:"many2many:user_following;"`
	Posts     []Post    `json:"posts" gorm:"foreignKey:UserID"`
	Likes     []Post    `json:"likes" gorm:"many2many:user_likes;"`
	Comments  []Comment `json:"comments" gorm:"foreignKey:UserID"`
}

type CreateOrUpdateUserInput struct {
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// =====================
// Repository functions
// =====================

func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	result := db.Preload("Following").Preload("Posts").Preload("Likes").Preload("Comments").Find(&users)
	return users, result.Error
}

func GetUser(db *gorm.DB, id int) (User, error) {
	var user User
	result := db.Preload("Following").Preload("Posts").Preload("Likes").Preload("Comments").First(&user, id)
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

func GetUserByEmail(db *gorm.DB, email string) (User, error) {
	var user User
	result := db.Where("email = ?", email).First(&user)
	return user, result.Error
}

func FollowUser(db *gorm.DB, follower *User, following *User) error {
	err := db.Model(follower).Association("Following").Append(following)
	return err
}

func UnfollowUser(db *gorm.DB, follower *User, following *User) error {
	err := db.Model(follower).Association("Following").Delete(following)
	return err
}
