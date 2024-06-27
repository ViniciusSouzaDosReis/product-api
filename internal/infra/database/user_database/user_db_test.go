package user_database

import (
	"testing"

	"github.com/ViniciusSouzaDosReis/product-api/internal/entity"
	"github.com/ViniciusSouzaDosReis/product-api/internal/infra/database/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
)

var sqliteDialecto = sqlite.Open("file::memory:")

func TestCreateUser(t *testing.T) {
	db, err := utils.CreateDBConnection(sqliteDialecto, &entity.User{})
	assert.NoError(t, err)

	newUser, _ := entity.NewUser("John Doe", "j@j.com", "123456")
	userDB := NewUserDB(db)

	err = userDB.Create(newUser)
	assert.Equal(t, nil, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", newUser.ID).Error
	assert.Equal(t, nil, err)

	assert.Equal(t, userFound.ID, newUser.ID)
	assert.Equal(t, userFound.Email, newUser.Email)
	assert.Equal(t, userFound.Name, newUser.Name)
	assert.Equal(t, userFound.Password, newUser.Password)
}

func TestFindByEmail(t *testing.T) {
	db, err := utils.CreateDBConnection(sqliteDialecto, &entity.User{})
	assert.NoError(t, err)

	newUser, _ := entity.NewUser("John Doe", "j@j.com", "123456")
	userDB := NewUserDB(db)

	userDB.Create(newUser)

	userFound, err := userDB.FindByEmail(newUser.Email)
	assert.Equal(t, nil, err)
	assert.Equal(t, newUser.Email, userFound.Email)
	assert.Equal(t, newUser.ID, userFound.ID)
	assert.Equal(t, newUser.Name, userFound.Name)
	assert.Equal(t, newUser.Password, userFound.Password)
}
