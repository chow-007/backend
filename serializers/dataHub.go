package serializers

import "time"

type RabbitMqData struct {
	Time time.Time `json:"time"`
	Code string `json:"code"`
	Co2 float64 `json:"co2"`
	O2 float64 `json:"o2"`
	Temperature float64 `json:"temperature"`
	AirHumidity float64 `json:"air_humidity"`
	GroundHumidity float64 `json:"ground_humidity"`
	Illumination float64 `json:"illumination"`
}
