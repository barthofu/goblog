package services

import (
	"blog/libs"
	"blog/models"
)

// ==========================================
// CRUD functions
// ==========================================

func CreateUser(email string, firstName string, lastName string) (*models.User, error) {
	user := &models.User{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}

	err := models.CreateUser(libs.DB, user)
	return user, err
}

func GetUser(id int) (*models.User, error) {
	user, err := models.GetUser(libs.DB, id)
	return &user, err
}

func GetUserByEmail(email string) (*models.User, error) {
	user, err := models.GetUserByEmail(libs.DB, email)
	return &user, err
}

func GetAllUsers() ([]models.User, error) {
	users, err := models.GetAllUsers(libs.DB)
	return users, err
}

func UpdateUser(id int, email string, firstName string, lastName string) (*models.User, error) {
	user, err := models.GetUser(libs.DB, id)
	if err != nil {
		return nil, err
	}

	user.Email = email
	user.FirstName = firstName
	user.LastName = lastName

	err = models.UpdateUser(libs.DB, &user)
	return &user, err
}

func DeleteUser(id int) error {
	user, err := models.GetUser(libs.DB, id)
	if err != nil {
		return err
	}

	err = models.DeleteUser(libs.DB, &user)
	return err
}

// ==========================================
// Follow functions
// ==========================================

func FollowUser(followerID int, followingID int) error {
	follower, err := models.GetUser(libs.DB, followerID)
	if err != nil {
		return err
	}

	following, err := models.GetUser(libs.DB, followingID)
	if err != nil {
		return err
	}

	err = models.FollowUser(libs.DB, &follower, &following)
	return err
}

func UnfollowUser(followerID int, followingID int) error {
	follower, err := models.GetUser(libs.DB, followerID)
	if err != nil {
		return err
	}

	following, err := models.GetUser(libs.DB, followingID)
	if err != nil {
		return err
	}

	err = models.UnfollowUser(libs.DB, &follower, &following)
	return err
}
