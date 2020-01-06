package services

import (
	"backend/configs"
	"backend/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//初始化MySQL
func initMysql() {
	option, err := gorm.Open("mysql", configs.Default.MysqlUrl)
	if err != nil {
		panic(err)
	}
	option.LogMode(configs.Default.Debug)
	Service.mysql = option
	Service.mysql.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.PermissionApi{})

	Service.mysql.Where(models.User{UserName:"admin"}).Attrs(&models.User{Password:"admin"}).FirstOrCreate(&models.User{})

	println("mysql services started")
}
