package swatch

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type SwatchRenderer struct {
	square  canvas.Rectangle
	objects fyne.CanvasObject
	parent  *Swatch
}

// Returns the the SwatchRenderer
func (renderer *SwatchRenderer) MinSize() fyne.Size {
	return renderer.square.MinSize()
}

// Determine where in the layout the swatch will be placed
func (renderer *SwatchRenderer) LayoutSize(size fyne.Size) {
}
