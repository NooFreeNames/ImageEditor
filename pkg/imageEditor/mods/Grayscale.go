package mods

import (
	"image"
	"image/color"
)

// NewGrayscale creates a new Grayscale object.
func NewGrayscale() *Grayscale {
	return new(Grayscale)
}

// Grayscale is a type representing a modifier that converts an image to
// grayscale.
type Grayscale struct{}

// ModifyPixel converts the image pixel to grayscale.
func (grayscale *Grayscale) ModifyPixel(_ image.Point, col color.RGBA,
	_ image.Image) color.RGBA {
	intensity := (uint16(col.R) + uint16(col.G) + uint16(col.B)) / 3
	col.R = uint8(intensity)
	col.G = uint8(intensity)
	col.B = uint8(intensity)
	return col
}
