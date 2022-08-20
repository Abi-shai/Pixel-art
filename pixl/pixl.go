package main

import (
	"image/color"
	"pixl/apptype"
	"pixl/swatch"
	"pixl/ui"

	"fyne.io/fyne/v2/app"
)

// Instanciating the pixel app
func main() {
	pixelApp := app.New()
	pixelWindow := pixelApp.NewWindow("pixl")

	state := apptype.State{
		color.NRGBA{255, 255, 255, 255},
		1,
		0,
		"lalala",
	}

	appInit := ui.AppInit{
		PixelWindow: pixelWindow,
		State:       &state,
		Swatches:    make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixelWindow.ShowAndRun()
}
