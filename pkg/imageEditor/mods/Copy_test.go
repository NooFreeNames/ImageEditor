package mods

import (
	"image"
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCopy(t *testing.T) {
	copy := NewCopy()
	want := new(Copy)
	assert.Equal(t, want, copy, "NewCopy() = %#v, wnat %#v", copy, want)
}

func TestCopy_ModifyPixel(t *testing.T) {
	copy := NewCopy()

	testValues := []color.RGBA{
		{1, 2, 3, 4},
		{255, 255, 255, 255},
		{0, 0, 0, 0},
	}

	for _, testValue := range testValues {
		point := image.Point{}
		got := copy.ModifyPixel(point, testValue, nil)
		assert.Equal(t, testValue, got, 
			"Copy.ModifyPixel(%#v, %#v, nil) = %#v, want %#v",
			point, testValue, got, testValue)
	}
}
