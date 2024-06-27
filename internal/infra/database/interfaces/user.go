package interfaces

import "github.com/ViniciusSouzaDosReis/product-api/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
