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
	MysqlUrl:        getDefaultEnvString("MYSQL_URL", "root:123456@tcp(39.104.150.48:9106)/demo?charset=utf8mb4&parseTime=True"),
	InfluxUrl:       getDefaultEnvString("INFLUX_URL", "influx://:@39.104.150.48:8086"),
	InfluxDBName:    getDefaultEnvString("INFLUX_DATABASE", "monitor"),
	BurrowUrl:       getDefaultEnvString("BURROW_URL", "127.0.0.1"),
	BurrowPort:      getDefaultEnvInt("BURROW_PORT", 8000),
	LoadUrl:         getDefaultEnvString("LOAD_URL", "http://39.104.150.48:9029"),
	ExecutorAddress: getDefaultEnvSlice("EXECUTOR_ADDRESS", []string{"39.104.150.48"}),
	//DataSourceAddress: getDefaultEnvSlice("DATA_SOURCE_ADDRESS", []string{"192.168.31.126", "39.104.150.48"}),
	DataSourceAddress: getDefaultEnvSlice("DATA_SOURCE_ADDRESS", []string{"39.104.150.48"}),
	DockerLoaderPort:  getDefaultEnvInt("DOCKER_LOADER_PORT", 9020),
	DockerMonitorPort: getDefaultEnvInt("DOCKER_MONITOR_PORT", 9010),
	DockerHostVolumes: getDefaultEnvString("DOCKER_HOST_VOLUMES", "/var/lib/container-data/"),
}

var ServerHost map[string]int

type Config struct {
	Debug             bool
	SecretKey         string
	MysqlUrl          string
	InfluxUrl         string
	InfluxDBName      string
	BurrowUrl         string
	BurrowPort        int
	LoadUrl           string
	ExecutorAddress   []string
	DataSourceAddress []string
	DockerLoaderPort  int
	DockerMonitorPort int
	TokenExpireTime   int
	DockerHostVolumes string
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
