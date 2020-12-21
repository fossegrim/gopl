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
		width, height          = 2560, 2560
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {

		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin

			// supersampling
			// subpixel resolution
			subpxPerRow := 2
			subpxPerColumn := 2
			subpxTotal := subpxPerRow * subpxPerColumn

			redAvg := float64(0)
			greenAvg := float64(0)
			blueAvg := float64(0)

			// I don't really understand
			noe := (float64(2)/height*(xmax-xmin) + xmin) - (float64(1)/height*(xmax-xmin) + xmin)
			// fmt.Fprintf(os.Stderr, "%g\n", noe)

			for subpxYIndex := 0; subpxYIndex < subpxPerRow; subpxYIndex++ {
				for subpxXIndex := 0; subpxXIndex < subpxPerColumn; subpxXIndex++ {

					xOffset := noe * (float64(subpxXIndex) / float64(subpxPerColumn))
					yOffset := noe * (float64(subpxYIndex) / float64(subpxPerRow))
					z := complex(x+xOffset, y+yOffset)

					// the colors are premultiplide and as such must be divided manutally
					red, green, blue, alpha := mandelbrot(z).RGBA()
					// fmt.Fprintf(os.Stderr, "red: %d\n", red)
					// fmt.Fprintf(os.Stderr, "alpha: %d\n", alpha)
					red /= alpha / 255
					green /= alpha / 255
					blue /= alpha / 255
					redAvg += float64(red) / float64(subpxTotal)
					greenAvg += float64(green) / float64(subpxTotal)
					blueAvg += float64(blue) / float64(subpxTotal)
					// fmt.Fprintf(os.Stderr, "red: %d\n", red)
					// fmt.Fprintf(os.Stderr, "redAvg: %g\n", redAvg)
				}
			}
			// fmt.Fprintf(os.Stderr, "---===---\n")
			img.Set(px, py, color.RGBA{uint8(redAvg), uint8(greenAvg), uint8(blueAvg), 255})

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
