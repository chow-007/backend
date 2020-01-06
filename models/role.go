package models

type Role struct {
	ID       uint16 `json:"id"`
	RoleName string `gorm:"type:varchar(100);not null; comment:'角色名称'" json:"role_name" binding:"required"`
	PsIds    string `gorm:"type:varchar(512); default:''; comment:'权限ids 1,2,3,6'" json:"ps_ids"`
	RoleDesc string `gorm:"type:text; comment:'角色描述'" json:"role_desc" binding:"required"`
	Children []Permission `gorm:"-" json:"children"`
}
