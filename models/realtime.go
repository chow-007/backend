package models

import "time"

type RealtimeData struct {
	ID             uint      `json:"id"`
	Code           float64   `gorm: not null;unique; comment:'编号'" json:"code"`
	Co2            float64   `gorm:"not null;unique; comment:'二氧化碳浓度'" json:"co_2"`
	O2             float64   `gorm:"not null;unique; comment:'氧气浓度'" json:"o_2"`
	Temperature    float64   `gorm:"not null;unique; comment:'温度'" json:"temperature"`
	AirHumidity    float64   `gorm:"not null;unique; comment:'空气湿度'" json:"air_humidity"`
	GroundHumidity float64   `gorm:"not null;unique; comment:'土壤湿度'" json:"ground_humidity"`
	Illumination   float64   `gorm:"not null;unique; comment:'光照'" json:"illumination"`
	Time           time.Time `json:"time"`
}
