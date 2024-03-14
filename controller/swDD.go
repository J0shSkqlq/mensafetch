package controller

import (
	"fmt"
	"io"
	"mensafetch/entity"
	"net/http"
	"time"
)

const (
	baseUrl string = "https://api.studentenwerk-dresden.de/openmensa/v2"
)

func GetMeals(mensaId int, relativeDay int) (meals []entity.Meal, err error) {
	today := time.Now().Add(time.Duration(relativeDay*24) * time.Hour).Format("2006-01-02")
	endpoint := fmt.Sprintf("%s/canteens/%d/days/%s/meals", baseUrl, mensaId, today)
	resp, _ := http.Get(endpoint)
	bodyBytes, _ := io.ReadAll(resp.Body)
	meals, _ = entity.NewMealListFromJson(string(bodyBytes))
	return
}
