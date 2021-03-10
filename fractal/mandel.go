package fractal

import (
	"image"
	"image/color"
	"sync"
)

type Mandelbrot struct{}

const (
	maxIterations = 100
	blockSize     = 500
)

func NewMandelbrot() *Mandelbrot {
	return &Mandelbrot{}
}

func (f *Mandelbrot) Image(xMin, xMax, yMin, yMax float64, width, height int) image.Image {
	dx, dy := (xMax-xMin)/float64(width), (yMax-yMin)/float64(height)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	var wg sync.WaitGroup
	for sx := 0; sx < width; sx += blockSize {
		for sy := 0; sy < height; sy += blockSize {
			wg.Add(1)
			go func(sx, sy int) {
				for y := 0; y < sy+blockSize; y++ {
					cy := yMin + float64(y)*dy
					for x := sx; x < sx+blockSize; x++ {
						cx := xMin + float64(x)*dx
						img.SetRGBA(x, y, f.getColor(cx, cy))
					}
				}
				wg.Done()
			}(sx, sy)
		}
	}
	wg.Wait()
	return img
}

func (f *Mandelbrot) getColor(x, y float64) color.RGBA {
	var x0 float64
	var y0 float64
	for i := 0; i < maxIterations; i++ {
		x0, y0 = x0*x0-y0*y0+x, 2*x0*y0+y
		if x0*x0+y0*y0 > 4 {
			return color.RGBA{
				R: byte(i * 8),
				G: byte(i * 16),
				B: byte(i * 24),
				A: 255,
			}
		}
	}
	return color.RGBA{}
}
