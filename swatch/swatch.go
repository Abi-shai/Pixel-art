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
	selected     bool
	color        color.Color
	swatchIndex  int
	clickHandler func(swatch *Swatch)
}

func (swatch *Swatch) SetColor(color color.Color) {
	swatch.color = color
	// Whenever the swatch changes,
	// it will update on the screen itself
	swatch.Refresh()
}

// Creates a swatch
func NewSwatch(state *apptype.State, color color.Color, swatchIndex int, clickHandler func(swatch *Swatch)) {
	// Set a pointer to the Swatch
	swatch := &Swatch{
		selected:     false,
		color:        color,
		swatchIndex:  swatchIndex,
		clickHandler: clickHandler,
	}

	swatch.ExtendBaseWidget(swatch)
	return swatch
}

func (swatch *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(swatch.color)
	objects := []fyne.CanvasObject{square}

	return &SwatchRenderer{
		square:  *square,
		objects: objects,
		parent:  swatch,
	}
}
