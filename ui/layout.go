package ui

import "fyne.io/fyne/v2/container"

/*-- Function that build the ui elements and swatches --*/
func Setup(app *AppInit) {
	swatchesContainer := BuildSwatches(app)
	colorPicker := SetupColorPicker(app)

	appLayout := container.NewBorder(nil, swatchesContainer, nil, colorPicker)

	app.PixelWindow.SetContent(appLayout)
}
