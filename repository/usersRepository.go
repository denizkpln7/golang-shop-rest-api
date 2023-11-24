package repository

import (
	"denizkpln7/rest-api/database"
	"denizkpln7/rest-api/models"
	"errors"
	"gorm.io/gorm"
)

type IUsersRepository interface {
	CreateUser(user *models.User) error
	GetAllUser() []models.User
	GetById(id int) models.User
	Login(m *models.User) (models.User, error)
}

type usersRepository struct {
	db *gorm.DB
}

func UsersRespository(db *gorm.DB) IUsersRepository {
	return &usersRepository{
		db: db,
	}
}

func (u *usersRepository) CreateUser(user *models.User) error {
	err := database.DB.Model(models.User{}).Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *usersRepository) Login(m *models.User) (models.User, error) {
	var user models.User
	database.DB.Model(models.User{}).Where("email=? AND name=?", m.Email, m.Name).First(&user)

	if user.ID != 0 {
		println("böyle bir kullanıcı yok")
		return user, errors.New("böyle bir kullanıcı yok")
	}
	return user, nil

}

func (u *usersRepository) GetAllUser() []models.User {
	var users []models.User
	database.DB.Model(models.User{}).Find(&users)
	return users
}

func (u *usersRepository) GetById(id int) models.User {
	var user models.User
	database.DB.Model(models.User{}).Where("id=?", id).Preload("CreditCards").First(&user)
	return user
}
