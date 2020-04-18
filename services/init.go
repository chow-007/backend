package services

import (
	"backend/configs"
	"backend/models"
	"backend/serializers"
	"backend/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/jinzhu/gorm"
	jsoniter "github.com/json-iterator/go"
	"github.com/streadway/amqp"
	"math/rand"
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
	Service.mysql.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Video{}, &models.Threshold{})

	Service.mysql.Where(models.User{UserName: "admin"}).Attrs(&models.User{Password: "admin"}).FirstOrCreate(&models.User{})
	Service.mysql.Where(models.Threshold{Name:"co2"}).Attrs(&models.Threshold{Low:0.1, Height:0.5}).FirstOrCreate(&models.Threshold{})
	Service.mysql.Where(models.Threshold{Name:"o2"}).Attrs(&models.Threshold{Low:18, Height:23}).FirstOrCreate(&models.Threshold{})
	Service.mysql.Where(models.Threshold{Name:"temperature"}).Attrs(&models.Threshold{Low:25, Height:40}).FirstOrCreate(&models.Threshold{})
	Service.mysql.Where(models.Threshold{Name:"air_humidity"}).Attrs(&models.Threshold{Low:30, Height:100}).FirstOrCreate(&models.Threshold{})
	Service.mysql.Where(models.Threshold{Name:"ground_humidity"}).Attrs(&models.Threshold{Low:30, Height:100}).FirstOrCreate(&models.Threshold{})
	Service.mysql.Where(models.Threshold{Name:"illumination"}).Attrs(&models.Threshold{Low:1, Height:2}).FirstOrCreate(&models.Threshold{})

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

//初始化DataHub
func initDataHub() {
	connExit := make(chan *amqp.Error)
	chExit := make(chan *amqp.Error)
	conn, err := amqp.Dial(configs.Default.DataHubQueue)
	if err != nil {
		panic(err)
	}
	conn.NotifyClose(connExit)
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	ch.NotifyClose(chExit)
	Service.dataHub = &dataHubProxy{Conn: conn, Channel: ch}

	println("datahub service started")

	go func() {
		for{
			_t := time.Now()
			for i:=1; i<=30; i++ {

				rand.Seed(int64(i))

				tmpData := serializers.RabbitMqData{
					Time: _t,
					Code: fmt.Sprintf("%d# 温室", i),
					Co2:            0.04 + utils.Decimal(rand.Float64()),
					O2:             20.9 + utils.Decimal(rand.Float64()),
					Temperature:    15 + utils.Decimal(rand.Float64()),
					AirHumidity:    70 + utils.Decimal(rand.Float64()),
					GroundHumidity: 60 + utils.Decimal(rand.Float64()),
					Illumination:   59 + utils.Decimal(rand.Float64()),
				}

				tmp, err := jsoniter.Marshal(tmpData)
				if err != nil {
					panic(err)
				}
				Service.dataHub.Publish(tmp)
			}
			time.Sleep(time.Second * 10)
		}
	}()

	// 查询阈值并缓存到缓存
	var thresholds []models.Threshold
	Service.mysql.Find(&thresholds)
	for _, v := range thresholds{
		item := map[string]float64{
			"low": v.Low,
			"height": v.Height,
		}
		configs.Thresholds[v.Name] = item
	}

	// 接收rabbitQq消息
	Service.dataHub.Receive()

}
