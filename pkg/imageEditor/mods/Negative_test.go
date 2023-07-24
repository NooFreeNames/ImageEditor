package mods

import (
	"image"
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewNegative(t *testing.T) {
	invertValues := [256]uint8{}
	for i := uint8(0); i < 255; i++ {
		invertValues[i] = 255 - i
	}

	got := NewNegative()
	want := &Negative{invertValues}
	assert.Equal(t, want, got, "NewNegative() = %#v, want %#v", got, want)
}

func TestNegative_ModifyPixel(t *testing.T) {
	tests := []struct {
		name  string
		color color.RGBA
		want  color.RGBA
	}{
		{
			name:  "normal values",
			color: color.RGBA{1, 2, 3, 4},
			want:  color.RGBA{254, 253, 252, 4},
		},
		{
			name:  "maximum channel values",
			color: color.RGBA{255, 255, 255, 255},
			want:  color.RGBA{0, 0, 0, 255},
		},
		{
			name:  "minimum channel values",
			color: color.RGBA{0, 0, 0, 0},
			want:  color.RGBA{255, 255, 255, 0},
		},
		{
			name:  "multiple zero channels",
			color: color.RGBA{0, 3, 0, 81},
			want:  color.RGBA{255, 252, 255, 81},
		},
	}

	negative := NewNegative()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := negative.ModifyPixel(image.Point{}, test.color, nil)
			assert.Equal(t, test.want, got, 
				"Negative.ModifyPixel(image.Point{}, %#v, nil) = %#v, want %#v",
				test.color, got, test.want)
		})
	}
}
