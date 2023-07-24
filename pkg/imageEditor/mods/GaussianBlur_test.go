package mods

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGaussianBlur(t *testing.T) {
	tests := []struct {
		name  string
		sigma float64
		want  *GaussianBlur
	}{
		{
			name:  "sigma=1.0",
			sigma: 1.0,
			want: &GaussianBlur{
				halfSize: 1,
				kernel: []float64{
					0.274068619061197,
					0.45186276187760605,
					0.274068619061197,
				},
			},
		},
		{
			name:  "sigma=1.2",
			sigma: 1.2,
			want: &GaussianBlur{
				halfSize: 2,
				kernel: []float64{
					0.08562916395501294,
					0.24266759672960792,
					0.3434064786307582,
					0.24266759672960792,
					0.08562916395501294,
				},
			},
		},
		{
			name:  "sigma=2.0",
			sigma: 2.0,
			want: &GaussianBlur{
				halfSize: 3,
				kernel: []float64{
					0.07015932695902607,
					0.13107487896736597,
					0.19071282356963737,
					0.21610594100794114,
					0.19071282356963737,
					0.13107487896736597,
					0.07015932695902607,
				},
			},
		},
		{
			name:  "sigma=2.4",
			sigma: 2.4,
			want: &GaussianBlur{
				halfSize: 3,
				kernel: []float64{
					0.08868143961720941,
					0.13687640922660566,
					0.17759311499808242,
					0.193698072316205,
					0.17759311499808242,
					0.13687640922660566,
					0.08868143961720941,
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewGaussianBlur(test.sigma)
			assert.Equal(t, test.want, got,
				"NewGaussianBlur(%g) = %#v, vant %#v",
				test.sigma, got,  test.want,
			)
		})
	}
}
func TestGaussianBlur_ModifyPixel(t *testing.T) {
	tests := []struct {
		name         string
		sigma        float64
		bounds       image.Rectangle
		sourceMatrix []uint8
		wantMatrix   []uint8
	}{
		{
			name:   "sigma=2.0",
			sigma:  2.0,
			bounds: image.Rect(0, 0, 2, 2),
			sourceMatrix: []uint8{
				1, 45, 3, 4,
				33, 77, 4, 231,
				10, 4, 122, 4,
				4, 5, 10, 23,
			},
			wantMatrix: []uint8{
				11, 34, 34, 64,
				15, 37, 22, 103,
				13, 27, 46, 73,
				10, 24, 27, 64,
			},
		},
		{
			name:   "sigma=1.0",
			sigma:  1.0,
			bounds: image.Rect(0, 0, 2, 2),
			sourceMatrix: []uint8{
				0, 2, 3, 4,
				33, 22, 23, 3,
				20, 4, 1, 4,
				4, 5, 0, 0,
			},
			wantMatrix: []uint8{
				13, 7, 6, 3,
				19, 12, 11, 2,
				16, 7, 5, 2,
				11, 7, 4, 1,
			},
		},
		{
			name:   "sigma=2.4",
			sigma:  2.4,
			bounds: image.Rect(2, 2, 4, 4),
			sourceMatrix: []uint8{
				0, 2, 3, 4,
				33, 22, 23, 3,
				20, 4, 1, 4,
				4, 5, 0, 0,
			},
			wantMatrix: []uint8{
				14, 8, 6, 2,
				17, 11, 9, 1,
				17, 9, 7, 2,
				12, 8, 5, 1,
			},
		},
		{
			name:         "empty matrix",
			sigma:        1.0,
			bounds:       image.Rect(0, 0, 0, 2),
			sourceMatrix: []uint8{},
			wantMatrix:   []uint8{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			blur := NewGaussianBlur(test.sigma)
			img := image.NewRGBA(test.bounds)
			img.Pix = test.sourceMatrix

			for x := test.bounds.Min.X; x <= test.bounds.Max.X; x++ {
				for y := test.bounds.Min.Y; y <= test.bounds.Max.Y; y++ {
					img.SetRGBA(x, y, blur.ModifyPixel(
						image.Point{x, y},
						img.RGBAAt(x, y),
						img,
					))
				}
			}

			assert.Equal(t, img.Pix, test.wantMatrix, 
				"Incorrect pixel modification")
		})
	}
}
