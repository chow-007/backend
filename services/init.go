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
//func initDataHub() {
	//connExit := make(chan *amqp.Error) // TODO:注释
	//chExit := make(chan *amqp.Error)
	//conn, err := amqp.Dial(configs.Default.DataHubQueue)
	//if err != nil {
	//	panic(err)
	//}
	//conn.NotifyClose(connExit)
	//ch, err := conn.Channel()
	//if err != nil {
	//	panic(err)
	//}
	//ch.NotifyClose(chExit)
	//Service.dataHub = &dataHubProxy{Conn: conn, Channel: ch}
	//
	//println("datahub service started")
//
//	go func() {
//	    time.Sleep(time.Second * 20)
//		for{
//			_t := time.Now()
//			for i:=1; i<=30; i++ {
//
//				rand.Seed(int64(i))
//
//				tmpData := serializers.RabbitMqData{
//					Time: _t,
//					Code: fmt.Sprintf("%d# 温室", i),
//					Co2:            0.04 + utils.Decimal(rand.Float64()),
//					O2:             20.9 + utils.Decimal(rand.Float64()),
//					Temperature:    15 + utils.Decimal(rand.Float64()),
//					AirHumidity:    70 + utils.Decimal(rand.Float64()),
//					GroundHumidity: 60 + utils.Decimal(rand.Float64()),
//					Illumination:   59 + utils.Decimal(rand.Float64()),
//				}
//
//				tmp, err := jsoniter.Marshal(tmpData)
//				if err != nil {
//					panic(err)
//				}
//				Service.dataHub.Publish(tmp)
//			}
//			time.Sleep(time.Second * 10)
//		}
//	}()
//
//	// 查询阈值并缓存到缓存
//	var thresholds []models.Threshold
//	Service.mysql.Find(&thresholds)
//	for _, v := range thresholds{
//		item := map[string]float64{
//			"low": v.Low,
//			"height": v.Height,
//		}
//		configs.Thresholds[v.Name] = item
//	}
//
//	// 接收rabbitQq消息
//	Service.dataHub.Receive()
//
//}

func initDataHub ()  {
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
	go func() {
		time.Sleep(time.Second * 20)
		for{
			_t := time.Now()
			for i:=1; i<=30; i++ {

				rand.Seed(int64(i))

				data := serializers.RabbitMqData{
					Time: _t,
					Code: fmt.Sprintf("%d# 温室", i),
					Co2:            0.04 + utils.Decimal(rand.Float64()),
					O2:             20.9 + utils.Decimal(rand.Float64()),
					Temperature:    15 + utils.Decimal(rand.Float64()),
					AirHumidity:    70 + utils.Decimal(rand.Float64()),
					GroundHumidity: 60 + utils.Decimal(rand.Float64()),
					Illumination:   59 + utils.Decimal(rand.Float64()),
				}

// ==================

				configs.RealtimeData[data.Code] = map[string]interface{}{
					"code": data.Code,
					"co2": utils.Decimal(data.Co2),
					"o2": utils.Decimal(data.O2),
					"temperature": utils.Decimal(data.Temperature),
					"air_humidity": utils.Decimal(data.AirHumidity),
					"ground_humidity": utils.Decimal(data.GroundHumidity),
					"illumination": utils.Decimal(data.Illumination),
					"time": data.Time,
				}
				configs.RealtimeTemperature[data.Code] = utils.Decimal(data.Temperature)
				fmt.Println(configs.RealtimeData)

				// 判断是否报警
				var alarm configs.Alarm
				if data.Co2 < configs.Thresholds["co2"]["low"] || data.Co2 > configs.Thresholds["co2"]["height"] {
					alarm.Co2 = true
				}else {
					alarm.Co2 = false
				}
				if data.O2 < configs.Thresholds["o2"]["low"] || data.O2 > configs.Thresholds["o2"]["height"] {
					alarm.O2 = true
				}else {
					alarm.O2 = false
				}
				if data.Temperature < configs.Thresholds["temperature"]["low"] || data.Temperature > configs.Thresholds["temperature"]["height"] {
					alarm.Temperature = true
				}else {
					alarm.Temperature = false
				}
				if data.AirHumidity < configs.Thresholds["air_humidity"]["low"] || data.AirHumidity > configs.Thresholds["air_humidity"]["height"] {
					alarm.AirHumidity = true
				}else {
					alarm.AirHumidity = false
				}
				if data.GroundHumidity < configs.Thresholds["ground_humidity"]["low"] || data.GroundHumidity > configs.Thresholds["ground_humidity"]["height"] {
					alarm.GroundHumidity = true
				}else {
					alarm.GroundHumidity = false
				}
				alarm.Code = data.Code
				alarm.Time = data.Time
				configs.AlarmCache[data.Code] = alarm

				// 采集的数据存入InfluxDb库
				bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
					Precision: "ms",
					Database:  configs.Default.InfluxDBName,
				})

				point, _ := client.NewPoint(
					"greenhouse",
					map[string]string{"code": "value"},
					map[string]interface{}{
						"code": data.Code,
						"co2": data.Co2,
						"o2": data.O2,
						"temperature": data.Temperature,
						"air_humidity": data.AirHumidity,
						"ground_humidity": data.GroundHumidity,
						"illumination": data.Illumination},
					data.Time)
				bp.AddPoint(point)

				Service.influx.Write(bp)

			}
			time.Sleep(time.Second * 10)
		}
	}()

}