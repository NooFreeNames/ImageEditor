package mods

import (
	"image"
	"image/color"
)

// PixelModifier is an interface that defines a method to modify a pixel of an 
// image.
type PixelModifier interface {
	ModifyPixel(image.Point, color.RGBA, image.Image) color.RGBA
}
