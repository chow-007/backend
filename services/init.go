package services

import (
	"backend/configs"
	"backend/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//初始化MySQL
func initMysql() {
	option, err := gorm.Open("mysql", configs.Default.MysqlUrl)
	//option, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",configs.Default.MysqlUserName, configs.Default.MysqlPwd, configs.Default.MysqlIp, configs.Default.MysqlPort, configs.Default.MysqlDatabase))
	if err != nil {
		panic(err)
	}
	option.LogMode(configs.Default.Debug)
	Service.mysql = option
	Service.mysql.AutoMigrate(&models.User{})

	Service.mysql.Where(models.User{UserName:"admin"}).Attrs(&models.User{ID: uuid.NewV4().String(), UserPwd:"admin"}).FirstOrCreate(&models.User{})

	println("mysql services started")
}
