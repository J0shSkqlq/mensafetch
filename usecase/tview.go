package usecase

import (
	"fmt"
	"image/jpeg"
	"net/http"

	"github.com/rivo/tview"
)

func loadImageFromURI(uri string) (*tview.Image, error) {
	response, err := http.Get(uri)
	if err != nil {
		return nil, fmt.Errorf("error while requesting image: %s", err)
	}
	defer response.Body.Close()

	photo, _ := jpeg.Decode(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading payload: %s", err)
	}

	tviewImage := tview.NewImage().SetImage(photo)

	return tviewImage, nil
}

func newPrimitive(text string) tview.Primitive {
	return tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(text)
}

func (f *Fetcher) GetUIView() {
	grid := tview.NewGrid().
		SetBorders(true).
		SetColumns(0, -2).
		SetRows(1, 0, 0, 0, 0).
		AddItem(newPrimitive("mensafetch"), 0, 0, 1, 2, 0, -1, false)
	for i, meal := range f.meals {
		image, err := loadImageFromURI("https:" + meal.Image)
		if err != nil {
			fmt.Printf("error while reading payload: %s", err)
		}
		textView := tview.NewTextView().
			SetDynamicColors(true).
			SetRegions(true).
			SetText(PrettyStringJson(meal))
		grid.AddItem(image, i+1, 0, 1, 1, 0, 0, false)
		grid.AddItem(textView, i+1, 1, 1, 1, 0, 0, false)

	}
	if err := tview.NewApplication().SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}
}
