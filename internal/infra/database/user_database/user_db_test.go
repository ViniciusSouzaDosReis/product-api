package user_database

import (
	"testing"

	"github.com/ViniciusSouzaDosReis/product-api/internal/entity/user"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Equal(t, nil, err)
	db.AutoMigrate(&user.User{})

	newUser, _ := user.NewUser("John Doe", "j@j.com", "123456")
	userDB := NewUser(db)

	err = userDB.Create(*newUser)
	assert.Equal(t, nil, err)

	var userFound user.User
	err = db.First(&userFound, "id = ?", newUser.ID).Error
	assert.Equal(t, nil, err)

	assert.Equal(t, userFound.ID, newUser.ID)
	assert.Equal(t, userFound.Email, newUser.Email)
	assert.Equal(t, userFound.Name, newUser.Name)
	assert.Equal(t, userFound.Password, newUser.Password)
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Equal(t, nil, err)
	db.AutoMigrate(&user.User{})

	newUser, _ := user.NewUser("John Doe", "j@j.com", "123456")
	userDB := NewUser(db)

	userDB.Create(*newUser)

	userFound, err := userDB.FindByEmail(newUser.Email)
	assert.Equal(t, nil, err)
	assert.Equal(t, newUser.Email, userFound.Email)
	assert.Equal(t, newUser.ID, userFound.ID)
	assert.Equal(t, newUser.Name, userFound.Name)
	assert.Equal(t, newUser.Password, userFound.Password)
}
