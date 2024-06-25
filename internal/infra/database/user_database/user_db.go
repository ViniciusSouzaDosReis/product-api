package user_database

import (
	"github.com/ViniciusSouzaDosReis/product-api/internal/entity/user"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user user.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(email string) (*user.User, error) {
	var user user.User

	// if err := u.DB.First(&user, "email = ?", email).Error; err != nil{
	// 	return nil, err
	// }
	err := u.DB.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
