package configs

import (
	"os"
	"strconv"
	"strings"
)

var Default = &Config{
	Debug:           getDefaultEnvBool("DEBUG", true),
	SecretKey:       getDefaultEnvString("SECRET_KEY", "test"),
	TokenExpireTime: getDefaultEnvInt("TokenExpireTime", 3600*1000000),
	MysqlUrl:        getDefaultEnvString("MYSQL_URL", "root:123456@tcp(39.104.150.48:9106)/beiyou?charset=utf8mb4&parseTime=True"),
	InfluxUrl:       getDefaultEnvString("INFLUX_URL", "influx://:@39.104.150.48:8086"),
	//InfluxDBName:    getDefaultEnvString("INFLUX_DATABASE", "monitor"),
	InfluxDBName:    getDefaultEnvString("INFLUX_DATABASE", "telegraf"),
}

var ServerHost map[string]int

type Config struct {
	Debug             bool
	SecretKey         string
	MysqlUrl          string
	InfluxUrl         string
	InfluxDBName      string
	TokenExpireTime   int
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
