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
		SwatchSelected: 0,
	}

	appInit := ui.AppInit{
		pixelWindow: pixelWindow,
		state:       &state,
		make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.pixelWindow.ShowAndRun()
}
