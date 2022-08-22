package ui

import (
	"pixl/apptype"
	pixelcanvas "pixl/pixelCanvas"
	"pixl/swatch"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixelCanvas *pixelcanvas.PixelCanvas
	PixelWindow fyne.Window
	State       *apptype.State
	Swatches    []*swatch.Swatch
}
