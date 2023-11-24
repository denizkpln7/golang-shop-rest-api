package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string       `json:"name"`
	Surname     string       `json:"surname"`
	Email       string       `json:"email"`
	CreditCards []CreditCard `gorm:"foreignKey:UserRefer"`
}

type CreditCard struct {
	gorm.Model
	Number    string `json:"number"`
	UserRefer uint   `json:"userRefer"`
}
