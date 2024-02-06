package entity

import (
	"fmt"
	"reflect"
)

type Mealplan struct {
	mensa Mensa
	meals []Meal
}

func (mp *Mealplan) PrettyPrintJson() string {
	output := ""
	for _, meal := range mp.meals {
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
