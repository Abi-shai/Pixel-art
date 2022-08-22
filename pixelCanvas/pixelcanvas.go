package pixelcanvas

import (
	"image"
	"image/color"
	"pixl/apptype"

	"fyne.io/fyne"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/widget"
)

type PixelCanvasMouseState struct {
	previousCoords *fyne.PointEvent
}

type PixelCanvas struct {
	widget.BaseWidget
	apptype.PixelCanvasConfig
	renderer    *PixelCanvasRenderer
	PixelData   image.Image
	mouseState  PixelCanvasMouseState
	appState    *apptype.State
	reloadImage bool
}

// Bounds func to calculate the position of the pixel canvas in the ui
func (pixelCanvas *PixelCanvas) Bounds() image.Rectangle {
	x0 := int(pixelCanvas.CanvasOffset.X)
	y0 := int(pixelCanvas.CanvasOffset.Y)
	// To calculate the exact position of the canvas for the X and Y axes
	// Take the numbers of cols * by the size of the pixel
	// And add the offset
	x1 := int(pixelCanvas.PixelColumns*pixelCanvas.PixelSize + int(pixelCanvas.CanvasOffset.X))
	y1 := int(pixelCanvas.PixelColumns*pixelCanvas.PixelSize + int(pixelCanvas.CanvasOffset.Y))
	return image.Rect(x0, y0, x1, y1)
}

// To determine if the mouse is in the bounding box of the canvas
func InBounds(pos fyne.Position, bounds image.Rectangle) bool {
	if pos.X >= int(bounds.Min.X) &&
		pos.X < int(bounds.Max.X) &&
		pos.Y >= int(bounds.Min.Y) &&
		pos.Y < int(bounds.Max.Y) {
		return true
	}
	return false
}

func NewBlankImage(columns, rows int, color color.Color) image.Image {
	image := image.NewNRGBA(image.Rect(0, 0, columns, rows))

	// Set image color for rows and columns
	for y := 0; y < rows; y++ {
		for x := 0; x < columns; x++ {
			image.Set(x, y, color)
		}
	}
	return image
}

func NewPixelCanvas(state *apptype.State, config apptype.PixelCanvasConfig) *PixelCanvas {
	pixelCanvas := &PixelCanvas{
		PixelCanvasConfig: config,
		appState:          state,
	}

	pixelCanvas.PixelData = NewBlankImage(pixelCanvas.PixelColumns, pixelCanvas.PixelRows, color.NRGBA{128, 128, 128, 255})
	pixelCanvas.ExtendBaseWidget(pixelCanvas)

	return pixelCanvas
}

// Creating the renderer widget
func (pixelCanvas *PixelCanvas) CreateRenderer() fyne.WidgetRenderer {
	canvasImage := canvas.NewImageFromImage(pixelCanvas.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvasImage.FillMode = canvas.ImageFillContain

	canvasBorder := make([]canvas.Line, 4)
	for i := 0; i < len(canvasBorder); i++ {
		canvasBorder[i].StrokeColor = color.NRGBA{100, 100, 100, 255}
		canvasBorder[i].StrokeWidth = 2
	}

	renderer := &PixelCanvasRenderer{
		pixelCanvas:  pixelCanvas,
		canvasImage:  canvasImage,
		canvasBorder: canvasBorder,
	}

	pixelCanvas.renderer = renderer

	return renderer
}
