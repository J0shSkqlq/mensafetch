package usecase

import (
	"fmt"
	"mensafetch/config"
	"mensafetch/controller"
	"mensafetch/entity"
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
	fmt.Print(entity.PrettyStringJson(f.meals))
}
