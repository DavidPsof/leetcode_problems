package postgre

import "gorm.io/gorm"

type UserDB struct {
	Name         string
	SecondName   string
	email        string
	PasswordHash string
	gorm.Model
}
