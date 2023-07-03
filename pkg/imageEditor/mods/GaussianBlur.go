package mods

import (
	"image"
	"image/color"
	"math"
)

// NewGaussianBlur creates a new GaussianBlur object with given sigma value.
// sigma is the degree of blur.
func NewGaussianBlur(sigma float64) *GaussianBlur {
	size := int(3 * sigma + 0.5)
	if size%2 == 0 {
		size++
	}

	halfSize := size / 2

	kernel := make([]float64, size)
	var sum float64
	for i := 0; i < size; i++ {
		x := float64(i - halfSize)
		kernel[i] = math.Exp(-x*x/(2*sigma*sigma)) / (sigma * math.Sqrt(2*math.Pi))
		sum += kernel[i]
	}
	for i := 0; i < size; i++ {
		kernel[i] /= sum
	}

	return &GaussianBlur{halfSize, kernel}
}

// GaussianBlur represents a type that applies a gaussian blur effect to an
// image.
type GaussianBlur struct {
	halfSize int
	kernel   []float64
}

// ModifyPixel applies a gaussian blur to an image pixel.
func (blur *GaussianBlur) ModifyPixel(position image.Point, c color.RGBA,
	src image.Image) color.RGBA {

	var resultR, resultG, resultB, resultA, weight float64
	imageBounds := src.Bounds()
	blurBounds := image.Rect(
		position.X-blur.halfSize,
		position.Y-blur.halfSize,
		position.X+blur.halfSize,
		position.Y+blur.halfSize,
	)

	for x := blurBounds.Min.X; x <= blurBounds.Max.X; x++ {
		if x < imageBounds.Min.X || x >= imageBounds.Max.X {
			continue
		}
		for y := blurBounds.Min.Y; y <= blurBounds.Max.Y; y++ {
			if y < imageBounds.Min.Y || y >= imageBounds.Max.Y {
				continue
			}
			r, g, b, a := src.At(x, y).RGBA()
			mul := blur.kernel[x-blurBounds.Min.X] * blur.kernel[y-blurBounds.Min.Y]

			resultR += float64(r) * mul
			resultG += float64(g) * mul
			resultB += float64(b) * mul
			resultA += float64(a) * mul
			weight += mul
		}
	}

	resultR /= weight
	resultG /= weight
	resultB /= weight
	resultA /= weight

	return color.RGBA{
		uint8(resultR / 256),
		uint8(resultG / 256),
		uint8(resultB / 256),
		uint8(resultA / 256),
	}
}
