package models

type Threshold struct {
	ID      uint16 `json:"id"`
	Name string `json:"name"`
	Low float64 `json:"low"`
	Height float64 `json:"height"`
}

