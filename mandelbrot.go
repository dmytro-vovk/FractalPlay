package main

import (
	"fmt"
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/dmytro-vovk/FractalPlay/fractal"
)

func main() {
	currentW, currentH := 1000, 1000
	minX, maxX := -2.0, 1.0
	minY, maxY := -1.5, 1.5

	w := app.New().NewWindow("Mandelbrot")
	w.Resize(fyne.NewSize(float32(currentW), float32(currentH)))
	w.SetPadded(false)
	m := fractal.NewMandelbrot()
	img := m.Image(minX, maxX, minY, maxY, currentW, currentH)
	w.SetContent(canvas.NewRaster(func(w, h int) image.Image {
		if w != currentW || h != currentH {
			currentW, currentH = w, h
			img = m.Image(
				minX, maxX,
				minY, maxY,
				currentW, currentH,
			)
			fmt.Printf("Scale is %g\n", maxX-minX)
		}
		return img
	}))
	h := float32(1.0)
	w.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		shiftY := (maxY - minY) * 0.1
		shiftX := (maxX - minX) * 0.1
		switch event.Name {
		case fyne.KeyUp:
			minY += shiftY
			maxY += shiftY
		case fyne.KeyDown:
			minY -= shiftY
			maxY -= shiftY
		case fyne.KeyLeft:
			minX += shiftX
			maxX += shiftX
		case fyne.KeyRight:
			minX -= shiftX
			maxX -= shiftX
		case fyne.KeyPlus:
			minX += shiftX
			maxX -= shiftX
			minY += shiftY
			maxY -= shiftY
		case fyne.KeyMinus:
			minX -= shiftX
			maxX += shiftX
			minY -= shiftY
			maxY += shiftY
		default:
			return
		}
		// hack to force repaint the canvas
		size := w.Canvas().Size()
		size.Height += h
		h = -h
		w.Resize(size)
	})
	w.ShowAndRun()
}
