package db

import "github.com/jinzhu/gorm"

type DBUser struct {
	gorm.Model
	Username string
}

func AddDBUser(username string) error {
	db, err := DbConnect()
	if err != nil {
		return err
	}

	defer db.Close()

	db.AutoMigrate(&DBUser{})

	err = db.Create(&DBUser{
		Username: username,
	}).Error

	if err != nil {
		return err
	}
	return nil
}
