package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	UserName  string    `gorm:"type:varchar(128); column:username; not null;unique; comment:'用户名'" json:"username" binding:"required"`
	Password  string    `gorm:"type:char(64);not null; comment:'密码'" json:"password" binding:"required"`
	Email     string    `gorm:"type:varchar(64); comment:'邮箱'" json:"email"`
	Mobile    string    `gorm:"type:varchar(32); comment:'手机'" json:"mobile"`
	State     bool      `gorm:"default:false; not null; comment:'状态'" json:"state"`
	CreatedAt time.Time `gorm:"comment:'账号创建时间'" json:"created_at"`
	UpdatedAt time.Time `gorm:"comment:'账号更新时间'" json:"updated_at"`
	Introduce string    `gorm:"type:text; comment:'简介'" json:"introduce"`
	Role      Role
	RoleID    uint16 `gorm:"comment:'用户角色Id'" json:"role_id"`
}
