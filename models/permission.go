package models

type Permission struct {
	ID     uint16 `json:"id"`
	AuthName string `gorm:"type:varchar(20); not null; comment:'权限名称'" json:"auth_name"`
	PsPid  uint16 `gorm:"not null; comment:'父id'" json:"ps_pid"`
	Path    string `gorm:"type:varchar(32); not null; comment:'操作方法'" json:"path"`
	Level  uint8  `gorm:"not null; comment:'权限等级'" json:"level"`
	Order  uint8  `gorm:"default:0; not null; comment:'权限等级'" json:"Order"`
	Children []Permission `gorm:"-" json:"children"`
}

type PermissionApi struct {
	ID           uint   `json:"id"`
	PsId         uint   `gorm:"index" json:"ps_id"`
	PsApiService string `grom:"type:varchar(255)" json:"ps_api_service"`
	PsApiAction  string `grom:"type:varchar(255)" json:"ps_api_action"`
	PsApiPath    string `grom:"type:varchar(255)" json:"ps_api_path"`
	PsApiOrder   uint8  `json:"ps_api_order"`
}
