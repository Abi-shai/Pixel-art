package pixelcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type PixelCanvasRenderer struct {
	pixelCanvas  *PixelCanvas
	canvasImage  *canvas.Image
	canvasBorder []canvas.Line
}

// Widget renderer interface implementation
func (renderer *PixelCanvasRenderer) MinSize() fyne.Size {
	return renderer.pixelCanvas.DrawingArea
}

// Widget renderer interface implementation
func (renderer *PixelCanvasRenderer) Objects() []fyne.CanvasObject {
	objects := make([]fyne.CanvasObject, 0, 5)
	for i := 0; i < len(renderer.canvasBorder); i++ {
		objects = append(objects, &renderer.canvasBorder[i])
	}
	objects = append(objects, renderer.canvasImage)
	return objects
}

func (renderer *PixelCanvasRenderer) Destroy() {

}

func (renderer *PixelCanvasRenderer) Layout(size fyne.Size) {
	renderer.LayoutCanvas(size)
	renderer.LayoutBorder(size)

}

func (renderer *PixelCanvasRenderer) Refresh() {
	if renderer.pixelCanvas.reloadImage {
		renderer.canvasImage = canvas.NewImageFromImage(renderer.pixelCanvas.PixelData)
		renderer.canvasImage.ScaleMode = canvas.ImageScalePixels
		renderer.canvasImage.FillMode = canvas.ImageFillContain
		renderer.pixelCanvas.reloadImage = false
	}
	renderer.Layout(renderer.pixelCanvas.Size())
	canvas.Refresh(renderer.canvasImage)
}

// Setting up the layout canvas correct size and position
func (renderer *PixelCanvasRenderer) LayoutCanvas(size fyne.Size) {
	imagePixelWidth := renderer.pixelCanvas.PixelColumns
	imagePixelHeight := renderer.pixelCanvas.PixelRows
	pixelSize := renderer.pixelCanvas.PixelSize

	//Set the canvas image to the correct location
	renderer.canvasImage.Move(fyne.NewPos(renderer.pixelCanvas.CanvasOffset.X, renderer.pixelCanvas.CanvasOffset.Y))
	// Resize the canvas image (to draw on)
	// to the correct size of a general pixel art image
	renderer.canvasImage.Resize(fyne.NewSize(float32(imagePixelWidth*pixelSize), float32(imagePixelHeight*pixelSize)))
}

// Setting up the layout border of the image canvas
func (renderer *PixelCanvasRenderer) LayoutBorder(size fyne.Size) {
	offset := renderer.pixelCanvas.CanvasOffset
	imageWidth := renderer.canvasImage.Size().Width
	imageHeight := renderer.canvasImage.Size().Height

	leftBorder := &renderer.canvasBorder[0]
	leftBorder.Position1 = fyne.NewPos(offset.X, offset.Y)
	// Adding the imageHeight for on the Y, for the border to wrap the entire image
	leftBorder.Position2 = fyne.NewPos(offset.X, offset.Y+imageHeight)

	topBorder := &renderer.canvasBorder[1]
	topBorder.Position1 = fyne.NewPos(offset.X, offset.Y)
	// Adding the imageWidth for on the X, for the border to wrap the entire image
	topBorder.Position2 = fyne.NewPos(offset.X+imageWidth, offset.Y)

	rightBorder := &renderer.canvasBorder[2]
	rightBorder.Position1 = fyne.NewPos(offset.X+imageWidth, offset.Y)
	rightBorder.Position2 = fyne.NewPos(offset.X+imageWidth, offset.Y+imageHeight)

	bottomBorder := &renderer.canvasBorder[2]
	bottomBorder.Position1 = fyne.NewPos(offset.X, offset.Y+imageHeight)
	bottomBorder.Position2 = fyne.NewPos(offset.X+imageWidth, offset.Y+imageHeight)

}
