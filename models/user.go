package models

type Role = int

const (
	RNormal Role = iota + 1
	RAdmin
)

type User struct {
	ID       string `json:"id"`
	UserName string `gorm:"type:varchar(255);not null;unique" json:"user_name" binding:"required"`
	UserPwd  string `gorm:"type:varchar(255);not null" json:"user_pwd" binding:"required"`
	Role     Role   `gorm:"not null" json:"role"`
}

type UserExecutorDatasourceInfo struct {
	ID               int    `gorm:"primary_key" json:"id"`
	UserID           string `gorm:"not null" json:"user_id"`
	HubInfoID        int    `json:"hub_id"`
	ExecutorInfoID   int    `json:"executor_id"`
	DataSourceInfoID int    `json:"db_id"`
}
