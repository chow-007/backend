package services

import (
	"github.com/influxdata/influxdb/client/v2"
	"github.com/jinzhu/gorm"
)

type serviceProxy struct {
	mysql  *gorm.DB
	influx client.Client
}

var Service *serviceProxy

func init() {
	Service = new(serviceProxy)
	initMysql()
	intiInfluxDB()
}
func (db *serviceProxy) Close() {
	db.mysql.Close()
	db.influx.Close()
}

