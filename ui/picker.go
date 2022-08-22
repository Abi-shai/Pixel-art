package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/lusingander/colorpicker"
)

func SetupColorPicker(app *AppInit) *fyne.Container {
	picker := colorpicker.New(200, colorpicker.StyleHue)
	picker.SetOnChanged(func(c color.Color) {
		// Set the color of the brush to the selected color
		// on the color picker
		app.State.BrushColor = c
		app.Swatches[app.State.SwatchSelected].SetColor(c)
	})
	return container.NewVBox(picker)
}
