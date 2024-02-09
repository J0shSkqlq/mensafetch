package entity

import (
	"encoding/json"
)

type Mensa struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Address     string    `json:"address"`
	Coordinates []float64 `json:"coordinates"`
	URL         string    `json:"url"`
	MenuURL     string    `json:"menu"`
}

func NewMensaFromJson(jsonData string) (mensa Mensa, err error) {
	err = json.Unmarshal([]byte(jsonData), &mensa)
	return
}
