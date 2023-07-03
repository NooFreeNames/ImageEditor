// Package mods provides image modifiers that can be used to alter the
// appearance of images.
package mods

import (
	"image"
	"image/color"
)

// NewCopy creates a new Copy object.
func NewCopy() *Copy {
	return new(Copy)
}

// Copy is used to copy pixels from one image to another.
type Copy struct{}

// Returns the received pixel of the image.
func (copy *Copy) ModifyPixel(_ image.Point, c color.RGBA,
	_ image.Image) color.RGBA {
	return c
}
