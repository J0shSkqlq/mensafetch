package entity

import (
	"encoding/json"
	"fmt"
	"reflect"
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

func PrettyStringJson(meals []Meal) string {
	output := ""
	for _, meal := range meals {
		output += "-----------\n" // Separator between meals
		v := reflect.ValueOf(meal)
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i)
			value := v.Field(i)
			key := field.Name
			if field.Type.Kind() == reflect.Map { // Check if the value is a nested dictionary (map in Go)
				output += fmt.Sprintf("\033[1m%s:\033[0m\n", key) // Bold key
				for _, k := range value.MapKeys() {
					nestedValue := value.MapIndex(k)
					output += fmt.Sprintf("  %s: %v\n", k, nestedValue)
				}
			} else {
				format := "\033[1m%s:\033[0m \033[33m%v\033[0m\n"
				if key == "Prices" {
					for k, v := range meal.Prices {
						output += fmt.Sprintf(format, k, fmt.Sprintf("%.2f€", v))
					}
				} else {
					output += fmt.Sprintf(format, key, value)
				}
			}
		}
	}
	return output
}
