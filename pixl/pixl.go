package main

import (
	"image/color"
	"pixl/apptype"
	"pixl/swatch"
	"pixl/ui"

	"fyne.io/fyne/v2"
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

	pixelCanvasConfig := apptype.PixelCanvasConfig{
		// On screen size of the entire drawing area
		DrawingArea:  fyne.NewSize(600, 600),
		CanvasOffset: fyne.NewPos(0, 0),
		PixelRows:    10,
		PixelColumns: 10,
		PixelSize:    30,
	}

	pixelCanvas := pixelCanvas.NewPixelCanvas()

	appInit := ui.AppInit{
		PixelWindow: pixelWindow,
		State:       &state,
		Swatches:    make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixelWindow.ShowAndRun()
}
