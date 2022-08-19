package ui

/*-- Function that build the ui elements and swatches --*/
func Setup(app *AppInit) {
	swatchesContainer := BuildSwatches(app)

	app.pixelWindow.SetContent(swatchesContainer)
}
