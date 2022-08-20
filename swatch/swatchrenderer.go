package swatch

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type SwatchRenderer struct {
	square  canvas.Rectangle
	objects []fyne.CanvasObject
	parent  *Swatch
}

// Returns the the SwatchRenderer
func (renderer *SwatchRenderer) MinSize() fyne.Size {
	return renderer.square.MinSize()
}

// Determine where in the layout the swatch will be placed
func (renderer *SwatchRenderer) Layout(size fyne.Size) {

	// Will resize the existing square to the maximun size available
	renderer.objects[0].Resize(size)
}

func (renderer *SwatchRenderer) Refresh() {
	renderer.Layout(fyne.NewSize(20, 20))
	renderer.square.FillColor = renderer.parent.Color

	// When the parent is selected change some strokes values
	if renderer.parent.Selected {
		renderer.square.StrokeWidth = 3
		renderer.square.StrokeColor = color.NRGBA{102, 102, 102, 102}
		renderer.objects[0] = &renderer.square
	} else {
		renderer.square.StrokeWidth = 0
		renderer.objects[0] = &renderer.square
	}

	canvas.Refresh(renderer.parent)
}

func (renderer *SwatchRenderer) Objects() []fyne.CanvasObject {
	return renderer.objects
}

func (renderer *SwatchRenderer) Destroy() {}
