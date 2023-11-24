package repository

import (
	"denizkpln7/rest-api/database"
	"denizkpln7/rest-api/models"
	"gorm.io/gorm"
)

type ICreditCardsRepository interface {
	CreateCard(card *models.CreditCard) error
	GetCreateCardById(id int) models.CreditCard
	//GetAllUser() []models.User
	//GetById(id int) models.User
}

type creditcardsRepository struct {
	db *gorm.DB
}

func CreditCardsRespository(db *gorm.DB) ICreditCardsRepository {
	return &creditcardsRepository{
		db: db,
	}
}

func (c creditcardsRepository) CreateCard(card *models.CreditCard) error {
	return database.DB.Model(models.CreditCard{}).Create(card).Error
}

func (c creditcardsRepository) GetCreateCardById(id int) models.CreditCard {
	var creditcard models.CreditCard
	database.DB.Model(models.CreditCard{}).Where("id=?", id).First(&creditcard)
	return creditcard
}
