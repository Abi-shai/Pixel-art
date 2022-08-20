package ui

import (
	"pixl/apptype"
	"pixl/swatch"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixelWindow fyne.Window
	State       *apptype.State
	Swatches    []*swatch.Swatch
}
