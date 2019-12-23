package services

import (
	"backend/models"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func (db *serviceProxy) GetUserById(userId string) (user models.User, err error) {
	err = db.mysql.Where("id=?", userId).First(&user).Error
	return
}

func (db *serviceProxy) UpdateUserById(user models.User) (err error) {
	err = db.mysql.Model(&models.User{}).Where("id=?", user.ID).UpdateColumns(user).Error
	return
}

func (db *serviceProxy) CreateUser(user models.User) bool {
	user.ID = uuid.NewV4().String()
	db.mysql.Create(&user)
	return db.mysql.NewRecord(user)
}

func (db *serviceProxy) IsExistUser(userName string) bool {
	err := db.mysql.Where("user_name=?", userName).Find(&models.User{}).Error
	if err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (db *serviceProxy) DeleteUser(userId string) error {
	return db.mysql.Where("id=?", userId).Delete(&models.User{}).Error
}

func (db *serviceProxy) GetUser(username, password string) (user models.User, err error) {
	err = db.mysql.Where("user_name=? AND user_pwd=?", username, password).First(&user).Error
	return
}

func (db *serviceProxy) GetUserList(limit, offset int) (users []models.User) {
	db.mysql.Limit(limit).Offset(limit * (offset - 1)).Find(&users)
	return
}