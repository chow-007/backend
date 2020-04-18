package services

import (
	"backend/configs"
	"backend/serializers"
	"backend/utils"
	"github.com/influxdata/influxdb/client/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/streadway/amqp"
	"log"
)

const (
	queueName = "monitor_queue"
)

type dataHubProxy struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (d *dataHubProxy) Publish(msg []byte) {
	go func() {
		q, err := d.Channel.QueueDeclare(
			queueName, // name
			false,   // durable
			false,   // delete when unused
			false,   // exclusive
			false,   // no-wait
			nil,     // arguments
		)
		failOnError(err, "Failed to declare a queue")

		err = d.Channel.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        msg,
			})
		failOnError(err, "Failed to publish a message")

	}()

}

type MonitorData struct {

	Name string
}
func (d *dataHubProxy) Receive()  {
	q, err := d.Channel.QueueDeclare(
		queueName, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := d.Channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var data serializers.RabbitMqData
			jsoniter.Unmarshal(d.Body, &data)

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
	}()

}

func (d *dataHubProxy) Close() {
	d.Conn.Close()
	d.Channel.Close()
}