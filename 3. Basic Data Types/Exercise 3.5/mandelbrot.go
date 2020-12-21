// Mandelbrot emits a PNG image of the mandelbrot fractal.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {

		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z. img.Set(px, py, mandelbrot(z))
			img.Set(px, py, mandelbrot(z))
		}

		if py%100 == 0 || py == height-1 {
			fmt.Fprintf(os.Stderr, "rendered row %d\n", py+1)
		}
	}

	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			// Nice purple color
			return color.RGBA{(255 - contrast*n) / 2, 0, 255 - contrast*n, 255}
		}
	}
	return color.Black
}
