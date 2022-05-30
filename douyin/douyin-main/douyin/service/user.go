package service

import (
	"douyin/model"
	"log"

	"github.com/jinzhu/gorm"
)

func GetUser(u model.User, db *gorm.DB) (model.User, error) {
	user, err := u.Get(db)
	if err != nil {
		log.Println(err)
		return user, err
	}
	return user, nil
}

func CreateUser(u model.User, db *gorm.DB) error {
	err := u.Create(db)
	if err != nil {
		log.Println(err)
	}
	return nil
}
func UpdateUser(u model.User, db *gorm.DB) error {
	err := u.Update(db)
	if err != nil {
		log.Println(err)
	}
	return nil
}
