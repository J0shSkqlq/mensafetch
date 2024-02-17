package usecase

import (
	"fmt"
	"mensafetch/config"
	"mensafetch/controller"
	"mensafetch/entity"
	"reflect"
)

type Fetcher struct {
	Config *config.Configuration
	Flags  *config.FlagSet
	meals  []entity.Meal
}

func NewFetcher(flags *config.FlagSet, config *config.Configuration) (*Fetcher, error) {
	meals, err := controller.GetMeals(flags.MensaId, flags.DayOffSet)

	if err != nil {
		return nil, err
	}

	return &Fetcher{
		config,
		flags,
		meals,
	}, nil
}

func (f *Fetcher) PrintMeals() {
	for _, meal := range f.meals {
		fmt.Print(PrettyStringJson(meal))
		fmt.Print("-----------\n") // Separator between meals

	}
}
func PrettyStringJson(meal entity.Meal) string {
	output := ""
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
	return output
}
