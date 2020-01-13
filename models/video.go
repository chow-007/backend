package models

type Video struct {
	ID int `json:"id"`
	Code string `gorm:"type:varchar(128); not null; comment:'厂房编号'"`
	Url string `gorm:"type:text; comment:'摄像头地址'" json:"url"`
}
