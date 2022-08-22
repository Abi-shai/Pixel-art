/* -- Setting up high level required types for the pixel application -- */
package apptype

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type BrushType = int

type PixelCanvasConfig struct {
	DrawingArea             fyne.Size
	CanvasOffset            fyne.Position
	PixelRows, PixelColumns int
	PixelSize               int
}

// Struct that stors infos about the state of the program
type State struct {
	BrushColor     color.Color
	BrushType      int
	SwatchSelected int
	FilePath       string
}

func (state *State) seFilePath(path string) {
	state.FilePath = path
}
