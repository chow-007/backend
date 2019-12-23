package services

import (
	"github.com/jinzhu/gorm"
)

type serviceProxy struct {
	mysql  *gorm.DB
}

var Service *serviceProxy

func init() {
	Service = new(serviceProxy)
	initMysql()
}
func (db *serviceProxy) Close() {
	db.mysql.Close()
}

