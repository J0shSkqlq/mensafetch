package entity

import (
	"encoding/json"
)

type Meal struct {
	ID       int                `json:"id"`
	Name     string             `json:"name"`
	Notes    []string           `json:"notes"`
	Prices   map[string]float64 `json:"prices"`
	Category string             `json:"category"`
	Image    string             `json:"image"`
	URL      string             `json:"url"`
}

func NewMealListFromJson(jsonData string) (meals []Meal, err error) {
	err = json.Unmarshal([]byte(jsonData), &meals)
	return
}
