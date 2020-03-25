package configs

import (
	"os"
	"strconv"
	"strings"
	"time"
)

var Default = &Config{
	Debug:           getDefaultEnvBool("DEBUG", true),
	SecretKey:       getDefaultEnvString("SECRET_KEY", "test"),
	TokenExpireTime: getDefaultEnvInt("TokenExpireTime", 3600*1000000),
	MysqlUrl:        getDefaultEnvString("MYSQL_URL", "root:root@tcp(localhost:3306)/beiyou?charset=utf8mb4&parseTime=True"),
	InfluxUrl:       getDefaultEnvString("INFLUX_URL", "influx://root:root@39.104.150.48:8086"),
	//InfluxDBName:    getDefaultEnvString("INFLUX_DATABASE", "monitor"),
	InfluxDBName:    getDefaultEnvString("INFLUX_DATABASE", "monitor"),
	DataHubQueue:  getDefaultEnvString("DATA_HUB_QUEUE", "amqp://root:root@127.0.0.1:5672/beiyou"),
}

var ServerHost map[string]int

type Config struct {
	Debug             bool
	SecretKey         string
	MysqlUrl          string
	InfluxUrl         string
	InfluxDBName      string
	TokenExpireTime   int
	DataHubQueue      string
}

func getDefaultEnvString(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}

func getDefaultEnvInt(key string, defaultValue int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	bValue, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return bValue
}

func getDefaultEnvBool(key string, defaultValue bool) bool {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	bValue, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}
	return bValue
}

func getDefaultEnvSlice(key string, defaultValue []string) []string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	tmp := strings.Split(value, ",")
	res := make([]string, 0)
	for _, v := range tmp {
		if len(v) == 0 {
			continue
		}
		res = append(res, v)
	}
	return res
}

var Thresholds = make(map[string]map[string]float64)


type Alarm struct {
	Time time.Time `json:"time"`
	Code string `json:"code"`
	Co2 bool `json:"co2"`
	O2 bool `json:"o2"`
	Temperature bool `json:"temperature"`
	AirHumidity bool `json:"air_humidity"`
	GroundHumidity bool `json:"ground_humidity"`
	Illumination bool `json:"illumination"`
}
var AlarmCache = make(map[string]Alarm)
