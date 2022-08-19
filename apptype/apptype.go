/* -- Setting up high level required types for the pixel application -- */
package apptype

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type BrushType = int

type PixelCanvasConfig struct {
	drawingArea             fyne.Size
	canvasOffset            fyne.Position
	pixelRows, pixelColumns int
	pixelSize               int
}

// Struct that stors infos about the state of the program
type State struct {
	BrushColor     color.Color
	brushType      int
	SwatchSelected int
	filePath       string
}

func (state *State) seFilePath(path string) {
	state.filePath = path
}
