package models

type Permission struct {
	ID       uint16       `json:"id"`
	AuthName string       `gorm:"type:varchar(128); not null; comment:'权限名称'" json:"auth_name"`
	PsPid    uint16       `gorm:"not null; comment:'父id'" json:"ps_pid"`
	Path     string       `gorm:"type:varchar(256); not null; comment:'操作方法'" json:"path"`
	Level    uint8        `gorm:"not null; comment:'权限等级'" json:"level"`
	Order    uint8        `gorm:"default:0; not null; comment:'权限排序'" json:"Order"`
	Children PermissionSlice `gorm:"-" json:"children"`
}


type PermissionSlice []Permission

func (p PermissionSlice) Len() int      { return len(p) }
func (p PermissionSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PermissionSlice) Less(i, j int) bool { return p[i].Order < p[j].Order }
