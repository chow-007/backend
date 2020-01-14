package services

import (
	"backend/configs"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/influxdata/influxdb/models"
	"log"
)

func (db *serviceProxy) Query(sql string) ([]models.Row, error) {
	log.Println("influxdb sql:", sql)
	resp, err := db.influx.Query(client.NewQuery(sql, configs.Default.InfluxDBName, "ms"))
	if err != nil {
		return []models.Row{}, err
	}
	if resp.Error() != nil{
		return []models.Row{}, resp.Error()
	}

	return resp.Results[0].Series, nil
}
