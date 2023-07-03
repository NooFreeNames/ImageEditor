package mods

import (
	"image"
	"image/color"
)

// NewNegative creates a new Negative object.
func NewNegative() *Negative {
	invertValues := [256]uint8{}

	for i := uint8(0); i < 255; i++ {
		invertValues[i] = 255 - i
	}

	return &Negative{invertValues}
}

// Negative is type represents a modifier that inverts the colors of an image.
type Negative struct {
	// invertValues stores the inverted values for each pixel.
	invertValues [256]uint8
}

// ModifyPixel inverts an image pixel.
func (negative *Negative) ModifyPixel(_ image.Point, col color.RGBA,
	_ image.Image) color.RGBA {
	col.R = negative.invertValues[col.R]
	col.G = negative.invertValues[col.G]
	col.B = negative.invertValues[col.B]
	return col
}
