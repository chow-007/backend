package services

import "backend/models"

func (db *serviceProxy) GetMenuList() []models.Permission {
	var permission []models.Permission
	db.mysql.Where("level=? OR level=?", 0, 1).Find(&permission)
	return permission

}

func (db *serviceProxy) GetRightList() []models.Permission {
	var permission []models.Permission
	db.mysql.Find(&permission)
	return permission
}

func (db *serviceProxy) DeleteRightById(id uint16) error {
	return db.mysql.Where("id=?", id).Delete(&models.Permission{}).Error
}
