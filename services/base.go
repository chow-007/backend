package services

import (
	"github.com/influxdata/influxdb/client/v2"
	"github.com/jinzhu/gorm"
)

type serviceProxy struct {
	mysql  *gorm.DB
	influx client.Client
	dataHub *dataHubProxy
}

var Service *serviceProxy

func init() {
	Service = new(serviceProxy)
	initMysql()
	intiInfluxDB()
	initDataHub()
}
func (db *serviceProxy) Close() {
	db.mysql.Close()
	db.influx.Close()
	db.dataHub.Close()
}

