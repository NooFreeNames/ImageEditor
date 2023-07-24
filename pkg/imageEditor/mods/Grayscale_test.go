package mods

import (
	"image"
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewGrayscale(t *testing.T) {
	got := NewGrayscale()
	want := new(Grayscale)
	assert.Equal(t, want, got, "NewGrayscale() = %#v, want %#v", got, want)
}

func TestGrayscale_ModifyPixel(t *testing.T) {
	tests := []struct {
		name  string
		color color.RGBA
		want  color.RGBA
	}{
		{
			name:  "normal values",
			color: color.RGBA{1, 2, 3, 4},
			want:  color.RGBA{2, 2, 2, 4},
		},
		{
			name:  "maximum channel values",
			color: color.RGBA{255, 255, 255, 255},
			want:  color.RGBA{255, 255, 255, 255},
		},
		{
			name:  "minimum channel values",
			color: color.RGBA{0, 0, 0, 0},
			want:  color.RGBA{0, 0, 0, 0},
		},
		{
			name:  "multiple zero channels",
			color: color.RGBA{0, 3, 0, 81},
			want:  color.RGBA{1, 1, 1, 81},
		},
	}

	grayscale := NewGrayscale()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := grayscale.ModifyPixel(image.Point{}, test.color, nil)
			assert.Equal(t, test.want, got, 
				"Grayscale.ModifyPixel(image.Point{}, %#v, nil) = %#v, want %#v",
				test.color, got, test.want)
		})
	}
}
