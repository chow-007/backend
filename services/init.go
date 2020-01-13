package services

import (
	"backend/configs"
	"backend/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/jinzhu/gorm"
	"net/url"
	"time"
)

//初始化MySQL
func initMysql() {
	option, err := gorm.Open("mysql", configs.Default.MysqlUrl)
	if err != nil {
		panic(err)
	}
	option.LogMode(configs.Default.Debug)
	Service.mysql = option
	Service.mysql.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.PermissionApi{}, &models.Video{})

	Service.mysql.Where(models.User{UserName:"admin"}).Attrs(&models.User{Password:"admin"}).FirstOrCreate(&models.User{})

	println("mysql services started")
}

//初始化InfluxDB
func intiInfluxDB() {
	opt, err := url.Parse(configs.Default.InfluxUrl)
	if err != nil {
		panic(err)
	}
	pwd, _ := opt.User.Password()
	Service.influx, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     fmt.Sprintf("http://%s", opt.Host),
		Username: opt.User.Username(),
		Password: pwd,
		Timeout:  time.Second * 5,
	})
	if err != nil {
		panic(err)
	}


	println("influxDB services started")
}
