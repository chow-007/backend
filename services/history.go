package services

import (
	"backend/configs"
	"backend/serializers"
	"errors"
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/influxdata/influxdb/models"
	"strconv"
)

// 获取历史数据
//func (db *serviceProxy) GetHistoryData(topics []string, startTime, endTime int64, offset, limit int64, period string) (result [][]interface{}, err error) {
//	startTSN := startTime * 1000000
//	endTSN := endTime * 1000000
//	sqlOrFilter := utils.GetSafetySqlOr(topics)
//
//	// 查询总数
//	var baseSQL string
//	var parameter []interface{}
//	if period == "" {
//		baseSQL = "SELECT COUNT(temperature) FROM greenhouse WHERE time > %s AND time < %s AND %s"
//		parameter = []interface{}{strconv.FormatInt(startTSN, 10), strconv.FormatInt(endTSN, 10), sqlOrFilter}
//	} else {
//		baseSQL = "SELECT COUNT(lag) FROM greenhouse WHERE time > %s AND time < %s AND %s group by time(%s)"
//		parameter = []interface{}{strconv.FormatInt(startTSN, 10), strconv.FormatInt(endTSN, 10), sqlOrFilter, period}
//
//	}
//	SQL := fmt.Sprintf(baseSQL, parameter...)
//	fmt.Println(SQL)
//	resp, err := db.influx.Query(client.NewQuery(SQL, configs.Default.InfluxDBName, "ms"))
//	if err != nil {
//		return
//	}
//	if len(resp.Results[0].Series) == 0 { // 无数据
//		return
//	}
//
//	// 查询数据
//	newLimit := limit
//	newOffset := limit * (offset - 1)
//	if period == "" {
//		baseSQL = "SELECT * FROM greenhouse WHERE time > %s AND time < %s AND %s LIMIT %s OFFSET %s"
//		parameter = []interface{}{strconv.FormatInt(startTSN, 10), strconv.FormatInt(endTSN, 10), sqlOrFilter, strconv.FormatInt(newLimit, 10), strconv.FormatInt(newOffset, 10)}
//	} else {
//		baseSQL = "SELECT * FROM greenhouse WHERE time > %s AND time < %s AND %s GROUP BY time(%s) LIMIT %s OFFSET %s"
//		parameter = []interface{}{strconv.FormatInt(startTSN, 10), strconv.FormatInt(endTSN, 10), sqlOrFilter, period, strconv.FormatInt(newLimit, 10), strconv.FormatInt(newOffset, 10)}
//	}
//	SQL = fmt.Sprintf(baseSQL, parameter...)
//	fmt.Println(SQL)
//	resp, err = db.influx.Query(client.NewQuery(SQL, configs.Default.InfluxDBName, "ms"))
//	if err != nil {
//		return
//	}
//	if resp.Results[0].Series == nil {
//		return
//	}
//	result = resp.Results[0].Series[0].Values
//	return
//}

func (db *serviceProxy) GetRealtimeData() (result []models.Row, err error) {
	SQL := "SELECT * FROM greenhouse GROUP BY code ORDER BY time DESC LIMIT 1"
	resp, err := db.influx.Query(client.NewQuery(SQL, configs.Default.InfluxDBName, "ms"))
	if err != nil {
		return
	}
	if resp.Results == nil || len(resp.Results) == 0{
		err = errors.New("Results is null or length is 0;" + resp.Err)
		return
	}
	if resp.Results[0].Series == nil {
		err = errors.New("Series is null" + resp.Results[0].Err)
		return
	}
	return resp.Results[0].Series, nil
}

// 获取历史数据
func (db *serviceProxy) GetHistoryData(filter serializers.HistoryRequest) (*client.Response, error) {
	startTSN := filter.StartTime * 1000000
	endTSN := filter.EndTime * 1000000
	period := filter.GetPeriod()

	var baseSQL string
	var parameter []interface{}

	// 查询数据
	if period == ""{
		baseSQL = "SELECT %s FROM greenhouse WHERE time > %s AND time < %s"
		parameter = []interface{}{filter.GetSelectFields(), strconv.FormatInt(startTSN, 10), strconv.FormatInt(endTSN, 10)}
	}else {
		baseSQL = "SELECT %s FROM greenhouse WHERE time > %s AND time < %s GROUP BY time(%s) fill(50) ORDER BY time"
		parameter = []interface{}{filter.GetSelectFields(), strconv.FormatInt(startTSN, 10), strconv.FormatInt(endTSN, 10), period}
	}
	SQL := fmt.Sprintf(baseSQL, parameter...)
	fmt.Println(SQL)
	resp, err := db.influx.Query(client.NewQuery(SQL, configs.Default.InfluxDBName, "ms"))
	if err != nil {
		return nil, err
	}
	if resp.Results == nil || len(resp.Results) == 0{
		return nil, errors.New(resp.Err)
	}
	if resp.Results[0].Series == nil {
		return nil, errors.New(resp.Results[0].Err)
	}
	//result = resp.Results[0].Series[0]
	return resp, nil
}

func (db *serviceProxy) CreateRealtimeData (data serializers.RabbitMqData) {

}

