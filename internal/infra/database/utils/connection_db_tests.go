package utils

import (
	"gorm.io/gorm"
)

func CreateDBConnection(dialector gorm.Dialector, models ...interface{}) (*gorm.DB, error) {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(models...)
	if err != nil {
		return nil, err
	}
	return db, nil
}
