/*-- Handling mouse events --*/
package swatch

import "fyne.io/fyne/v2/driver/desktop"

func (swatch *Swatch) MouseDown(event *desktop.MouseEvent) {
	swatch.ClickHandler(swatch)
	swatch.Selected = true
	swatch.Refresh()
}

func (swatch *Swatch) MouseUp(event *desktop.MouseEvent) {}
