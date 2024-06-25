package interfaces

import (
	"github.com/ViniciusSouzaDosReis/product-api/internal/entity/user"
)

type UserInterface interface {
	Create(user *user.User) error
	FindByEmail(email string) (*user.User, error)
}
