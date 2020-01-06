package services

import "backend/models"

func (db *serviceProxy) CreateRole(role models.Role) error {
	return db.mysql.Create(&role).Error
}

func (db *serviceProxy) DeleteRole(id uint16) error {
	return db.mysql.Where("id=?", id).Delete(models.Role{}).Error
}

func (db *serviceProxy) UpdateRoleById(id uint16, columns map[string]interface{}) error {
	return db.mysql.Model(&models.Role{}).Where("id=?", id).Updates(columns).Error
}

func (db *serviceProxy) DeleteRoleRightById(role models.Role) error {
	return db.mysql.Model(&models.Role{}).Where("id=?", role.ID).Updates(map[string]string{
		"role_name":role.RoleName,
		"role_desc": role.RoleDesc,
	}).Error
}

func (db *serviceProxy) GetRoleById(id uint16) (role models.Role, err error) {
	err = db.mysql.Where("id=?", id).First(&role).Error
	return
}

func (db *serviceProxy) GetRoleList() (roles []models.Role, err error) {
	//err = db.mysql.Select([]string{"id", "role_name", "role_desc"}).Find(&roles).Error
	err = db.mysql.Find(&roles).Error
	//db.mysql.Exec("SELECT r.id,r.role_name,r.role_desc, p.id as pid, p.auth_name,p.ps_pid,p.path,p.level FROM roles as r LEFT JOIN permissions as p ON p.id=r.ps_ids GROUP BY r.id")
	return
}
