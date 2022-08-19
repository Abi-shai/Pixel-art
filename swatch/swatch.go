package swatch

import (
	"image/color"
	"pixl/apptype"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// Setting up swatch struct params
type Swatch struct {
	widget.BaseWidget
	Selected     bool
	Color        color.Color
	SwatchIndex  int
	clickHandler func(swatch *Swatch)
}

func (swatch *Swatch) SetColor(color color.Color) {
	swatch.Color = color
	// Whenever the swatch changes,
	// it will update on the screen itself
	swatch.Refresh()
}

// Creates a swatch
func NewSwatch(state *apptype.State, color color.Color, swatchIndex int, clickHandler func(s *Swatch)) *Swatch {
	swatch := &Swatch{
		Selected:     false,
		Color:        color,
		clickHandler: clickHandler,
		SwatchIndex:  swatchIndex,
	}
	swatch.ExtendBaseWidget(swatch)
	return swatch
}

// func (swatch *Swatch) CreateRenderer() fyne.WidgetRenderer {
// 	square := canvas.NewRectangle(swatch.color)
// 	objects := []fyne.CanvasObject{square}

// 	return &SwatchRenderer{
// 		square:  *square,
// 		objects: objects,
// 		parent:  swatch,
// 	}
// }

func (swatch *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(swatch.Color)
	objects := []fyne.CanvasObject{square}

	return &SwatchRenderer{
		square:  *square,
		objects: objects,
		parent:  swatch,
	}
}
