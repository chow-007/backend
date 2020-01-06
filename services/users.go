package services

import (
	"backend/models"
)

func (db *serviceProxy) CreateUser(user models.User) error {
	return db.mysql.Create(&user).Error
}

func (db *serviceProxy) GetUserById(userId uint64) (user models.User, err error) {
	err = db.mysql.Where("id=?", userId).First(&user).Error
	return
}

func (db *serviceProxy) GetUserByNameAndPassword(username, password string) (user models.User, err error) {
	err = db.mysql.Where("username=? AND password=?", username, password).First(&user).Error
	return
}

func (db *serviceProxy) DeleteUser(userId string) error {
	return db.mysql.Where("id=?", userId).Delete(&models.User{}).Error
}

func (db *serviceProxy) UpdateUserById(id uint, user map[string]interface{}) (err error) {
	err = db.mysql.Model(&models.User{}).Where("id=?", id).Updates(user).Error
	return
}


func (db *serviceProxy) IsExistUserName(userName string) bool {
	rowAffected := db.mysql.Where("username=?", userName).First(&models.User{}).RowsAffected
	if rowAffected == 0 {
		return false
	}
	return true
}

func (db *serviceProxy) GetUserList(limit, offset int, keyWords string) (users []models.User, total int) {
	ormCount := db.mysql.Model(&models.User{})
	orm := db.mysql.Preload("Role")

	if keyWords != "" {
		ormCount = ormCount.Where("username LIKE '?'", "%"+keyWords+"%")
		orm.Where("username LIKE ?", "%"+keyWords+"%")
	}

	ormCount.Count(&total)
	orm.Limit(limit).Offset(limit * (offset - 1)).Find(&users)
	return
}
