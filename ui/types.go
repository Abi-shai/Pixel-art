package ui

import (
	"pixl/apptype"
	"pixl/swatch"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	pixelWindow fyne.Window
	State       *apptype.State
	Swatches    []*swatch.Swatch
}
