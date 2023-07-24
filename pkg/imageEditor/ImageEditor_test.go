// Package imageEditor provides functionality for editing images. It allowsyou
// to change pixels, crop images, and save them in various formats.
package imageEditor

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"io"
	"os"
	"testing"

	"github.com/NooFreeNames/ImageEditor/pkg/imageEditor/geom"
	"github.com/NooFreeNames/ImageEditor/pkg/imageEditor/mods"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockImage struct {
	mock.Mock
	image.Image
}

func (mock *MockImage) ColorModel() color.Model {
	args := mock.Called()
	return args.Get(0).(color.Model)
}

func (mock *MockImage) Bounds() image.Rectangle {
	args := mock.Called()
	return args.Get(0).(image.Rectangle)
}

func (mock *MockImage) At(x, y int) color.Color {
	args := mock.Called(x, y)
	return args.Get(0).(color.Color)
}

func TestIsSupportedImageFormat(t *testing.T) {
	type args struct {
		mimeType string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Correct image format: " + MIMEJPEG,
			args: args{MIMEJPEG},
			want: true,
		},
		{
			name: "Correct image format: " + MIMEPNG,
			args: args{MIMEPNG},
			want: true,
		},
		{
			name: "Correct image format: " + MIMEPNG,
			args: args{MIMEPNG},
			want: true,
		},
		{
			name: "Unsupported format: image/gif",
			args: args{"image/gif"},
			want: false,
		},
		{
			name: "Incorrect image format",
			args: args{"beleberda"},
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := IsSupportedImageFormat(test.args.mimeType)
			assert.Equal(t, test.want, got, "image format: "+test.args.mimeType)
		})
	}
}

func TestNewImageEditor(t *testing.T) {

	tests := []struct {
		name     string
		fileName string
		wantErr  bool
	}{
		{
			name:     "jpeg extension",
			fileName: "test_image.jpeg",
			wantErr:  false,
		},
		{
			name:     "jpg extension",
			fileName: "test_image.jpg",
			wantErr:  false,
		},
		{
			name:     "png extension",
			fileName: "test_image.png",
			wantErr:  false,
		},
		{
			name:     "gif extension",
			fileName: "test_image.gif",
			wantErr:  true,
		},
		{
			name:     "broken jepg file",
			fileName: "broken_image.jpeg",
			wantErr:  true,
		},
		{
			name:     "broken png file",
			fileName: "broken_image.png",
			wantErr:  true,
		},
		{
			name:     "no image file",
			fileName: "no_image.txt",
			wantErr:  true,
		},
	}
	t.Run("file is nil", func(t *testing.T) {
		_, err := NewImageEditor(nil)
		assert.Error(t, err)
	})

	path := "../../test/images/"
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filePath := path + test.fileName
			file, err := os.Open(filePath)
			if err != nil {
				assert.FailNow(t, err.Error(), "Error opening the file: "+filePath)
				return
			}
			defer file.Close()
			editor, err := NewImageEditor(file)

			if test.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			file.Seek(0, io.SeekStart)
			img, _, err := image.Decode(file)

			if err != nil {
				assert.FailNow(t, err.Error(), "Error decoding the image: "+filePath)
				return
			}

			if !assert.ObjectsAreEqual(img, editor.source) {
				t.Error("The source field is incorrectly initialized")
			}

			if assert.NotNil(t, editor.destination, "The destination field must be initialized") {
				assert.Equal(t,
					editor.source.Bounds(),
					editor.destination.Bounds(),
					"The size of the destination image must be equal to the size of the source image",
				)
			}
			assert.False(t, editor.isModifiedPixels, "The isModifiedPixels field must have the value false")
			assert.False(t, editor.isCropped, "The isCropped field must have the value false")
		})
	}
}

func Test_decode(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		wantErr  bool
	}{
		{
			name:     "decoding test_image.jpeg",
			fileName: "test_image.jpeg",
			wantErr:  false,
		},
		{
			name:     "decoding test_image.jpg ",
			fileName: "test_image.jpg",
			wantErr:  false,
		},
		{
			name:     "decoding test_image.png",
			fileName: "test_image.png",
			wantErr:  false,
		},
		{
			name:     "decoding test_image.gif",
			fileName: "test_image.gif",
			wantErr:  true,
		},
		{
			name:     "decoding broken_image.jpeg",
			fileName: "broken_image.jpeg",
			wantErr:  true,
		},
		{
			name:     "decoding broken_image.png",
			fileName: "broken_image.png",
			wantErr:  true,
		},
		{
			name:     "decoding no_image.txt",
			fileName: "no_image.txt",
			wantErr:  true,
		},
	}
	path := "../../test/images/"
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filePath := path + test.fileName
			file, err := os.Open(filePath)
			if err != nil {
				assert.FailNow(t, err.Error(), "Error opening the file: "+filePath)
				return
			}
			defer file.Close()
			img, err := decode(file)

			if test.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			file.Seek(0, io.SeekStart)
			wantImg, _, err := image.Decode(file)

			if err != nil {
				assert.FailNow(t, err.Error(), "Error decoding the image: "+filePath)
				return
			}

			if !assert.ObjectsAreEqual(wantImg, img) {
				t.Error("The source field is incorrectly initialized")
			}
		})
	}
}

func TestImageEditor_IsModifiedImage(t *testing.T) {
	type fields struct {
		isModifiedPixels bool
		isCropped        bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "the image was not edited",
			fields: fields{false, false},
			want:   false,
		},
		{
			name:   "the image has been cropped",
			fields: fields{false, true},
			want:   true,
		},
		{
			name:   "the pixels of the image were modified",
			fields: fields{true, false},
			want:   true,
		},
		{
			name:   "the image has been edited in several ways",
			fields: fields{true, false},
			want:   true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			editor := &ImageEditor{
				source:           new(MockImage),
				destination:      new(image.RGBA),
				isModifiedPixels: test.fields.isModifiedPixels,
				isCropped:        test.fields.isCropped,
			}

			got := editor.IsModifiedImage()
			assert.Equal(t, test.want, got,
				"isModifiedPixels = %d; isCropped = %d;",
				editor.isModifiedPixels,
				editor.isCropped,
			)
		})
	}
}

func TestImageEditor_EditedImage(t *testing.T) {
	type fields struct {
		isModifiedPixels bool
		isCropped        bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "the image was not edited",
			fields: fields{false, false},
		},
		{
			name:   "the image has been cropped",
			fields: fields{false, true},
		},
		{
			name:   "the pixels of the image were modified",
			fields: fields{true, false},
		},
		{
			name:   "the image has been edited in several ways",
			fields: fields{true, false},
		},
	}
	bounds := image.Rect(0, 0, 10, 10)
	source := new(MockImage)
	source.On("Bounds").Return(bounds)
	source.On("At", mock.Anything, mock.Anything).Return(color.RGBA{100, 100, 100, 255})
	source.On("ColorModel").Return(color.RGBAModel)
	destination := image.NewRGBA(bounds)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			editor := &ImageEditor{
				source:           source,
				destination:      destination,
				isModifiedPixels: tt.fields.isModifiedPixels,
				isCropped:        tt.fields.isCropped,
			}
			got := editor.EditedImage()

			if !editor.IsModifiedImage() {
				if !assert.ObjectsAreEqual(editor.source, got) {
					t.Error("If the image has not been modified, then the source image should be returned.")
				}
			} else if !editor.isModifiedPixels {
				t.Error("If the ModifyPixels function has never been called, it must be called in the EditedImage function.")
			} else if !assert.ObjectsAreEqual(editor.destination, got) {
				t.Error("If the image has been modified it must be taken from the destination")
			}

		})
	}
}

func TestImageEditor_ModifyPixels(t *testing.T) {

	tests := []struct {
		name         string
		filter       mods.PixelModifier
		bounds       image.Rectangle
		sourceMatrix []uint8
		wantMatrix   []uint8
	}{
		{
			name:   "nil modifier",
			filter: nil,
			bounds: image.Rect(0, 0, 2, 2),
			sourceMatrix: []uint8{
				1, 45, 3, 4,
				33, 77, 4, 231,
				10, 4, 122, 4,
				4, 5, 10, 23,
			},
			wantMatrix: []uint8{
				1, 45, 3, 4,
				33, 77, 4, 231,
				10, 4, 122, 4,
				4, 5, 10, 23,
			},
		},
		{
			name:   "copy",
			filter: mods.NewCopy(),
			bounds: image.Rect(0, 0, 2, 2),
			sourceMatrix: []uint8{
				1, 45, 3, 4,
				33, 77, 4, 231,
				10, 4, 122, 4,
				4, 5, 10, 23,
			},
			wantMatrix: []uint8{
				1, 45, 3, 4,
				33, 77, 4, 231,
				10, 4, 122, 4,
				4, 5, 10, 23,
			},
		},
		{
			name:   "negative",
			filter: mods.NewNegative(),
			bounds: image.Rect(0, 0, 2, 2),
			sourceMatrix: []uint8{
				0, 2, 3, 4,
				33, 22, 23, 3,
				255, 255, 255, 255,
				0, 0, 0, 0,
			},
			wantMatrix: []uint8{
				255, 253, 252, 4,
				222, 233, 232, 3,
				0, 0, 0, 255,
				255, 255, 255, 0,
			},
		},
		{
			name:   "grayscale",
			filter: mods.NewGrayscale(),
			bounds: image.Rect(2, 2, 4, 4),
			sourceMatrix: []uint8{
				1, 2, 3, 4,
				33, 22, 11, 44,
				255, 255, 255, 255,
				0, 0, 0, 0,
			},
			wantMatrix: []uint8{
				2, 2, 2, 4,
				22, 22, 22, 44,
				255, 255, 255, 255,
				0, 0, 0, 0,
			},
		},
	}
	t.Run("ImageEditor.isModifiedPixels is true", func(t *testing.T) {
		source := image.NewRGBA(image.Rect(0, 0, 10, 10))
		destination := image.NewRGBA(image.Rect(0, 0, 5, 5))

		editor := ImageEditor{
			source:           source,
			destination:      destination,
			isModifiedPixels: true,
		}

		editor.ModifyPixels(mods.NewCopy())
		if !assert.ObjectsAreEqual(destination, editor.source) {
			t.Error("Destination must become the source")
		}
	})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			img := image.NewRGBA(test.bounds)
			img.Pix = test.sourceMatrix
			editor := ImageEditor{
				source:      img,
				destination: image.NewRGBA(img.Bounds()),
			}

			editor.ModifyPixels(test.filter)

			editedImage := editor.EditedImage().(*image.RGBA).Pix
			assert.Equal(t, editedImage, test.wantMatrix)
		})
	}
}

func TestImageEditor_Encode(t *testing.T) {
	type args struct {
		mimeType string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "encoding jpeg type",
			args:    args{MIMEJPEG},
			wantErr: false,
		},
		{
			name:    "encoding png type",
			args:    args{MIMEPNG},
			wantErr: false,
		},
		{
			name:    "encoding gif type",
			args:    args{"image/gif"},
			wantErr: true,
		},
		{
			name:    "encoding invalid type",
			args:    args{"beleberda"},
			wantErr: true,
		},
	}
	bounds := image.Rect(0, 0, 3, 3)
	source := new(MockImage)
	source.On("Bounds").Return(bounds)
	source.On("At", mock.Anything, mock.Anything).Return(color.RGBA{100, 100, 100, 255})
	source.On("ColorModel").Return(color.RGBAModel)
	destination := image.NewRGBA(bounds)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			editor := &ImageEditor{
				source:      source,
				destination: destination,
			}

			writer := new(bytes.Buffer)
			err := editor.Encode(writer, test.args.mimeType)

			message := "mimeType = " + test.args.mimeType
			if test.wantErr {
				assert.Error(t, err, message)
			} else {
				assert.NoError(t, err, message)
			}

		})
	}
}

func TestImageEditor_CropByRectangle(t *testing.T) {

	tests := []struct {
		name          string
		imageBounds   image.Rectangle
		cropBounds    image.Rectangle
		wantIsCropped bool
	}{
		{
			name:          "empty cropBounds",
			imageBounds:   image.Rect(0, 0, 2, 2),
			cropBounds:    image.Rect(0, 0, 0, 0),
			wantIsCropped: false,
		},
		{
			name:          "cropBounds equals imageBounds",
			imageBounds:   image.Rect(2, -2, 10, 10),
			cropBounds:    image.Rect(2, -2, 10, 10),
			wantIsCropped: false,
		},
		{
			name:          "cropBounds not in imageBounds",
			imageBounds:   image.Rect(0, 0, 10, 10),
			cropBounds:    image.Rect(-100, -100, 100, 100),
			wantIsCropped: false,
		},
		{
			name:          "cropBounds not in imageBounds",
			imageBounds:   image.Rect(-1, -2, 10, 10),
			cropBounds:    image.Rect(0, 0, -100, -100),
			wantIsCropped: false,
		},
		{
			name:          "cropBounds not in imageBounds",
			imageBounds:   image.Rect(2, 2, 100, 100),
			cropBounds:    image.Rect(-100, -100, -2, -2),
			wantIsCropped: false,
		},
		{
			name:          "cropBounds in imageBounds",
			imageBounds:   image.Rect(-3, -3, -10, -10),
			cropBounds:    image.Rect(-1, -2, -5, -5),
			wantIsCropped: false,
		},
		{
			name:          "cropBounds in imageBounds",
			imageBounds:   image.Rect(0, 0, 10, 10),
			cropBounds:    image.Rect(1, 2, 8, 8),
			wantIsCropped: true,
		},
		{
			name:          "cropBounds in imageBounds",
			imageBounds:   image.Rect(0, -10, 10, 10),
			cropBounds:    image.Rect(0, -10, 8, 8),
			wantIsCropped: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			editor := &ImageEditor{
				source:      image.NewRGBA(test.imageBounds),
				destination: image.NewRGBA(test.imageBounds),
				isCropped:   false,
			}
			editor.CropByRectangle(test.cropBounds)

			wantBounds := test.imageBounds
			if test.wantIsCropped {
				wantBounds = test.cropBounds
			}

			assert.Equal(t, wantBounds, editor.destination.Rect,
				"incorrect value in ImageEditor.destination.Rect.\nimageBounds %v, cropBounds %v",
				test.imageBounds, test.cropBounds)
			assert.Equal(t, test.wantIsCropped, editor.isCropped,
				"incorrect value in ImageEditor.isCropped.\nimageBounds %v, cropBounds %v",
				test.imageBounds, test.cropBounds)
		})
	}
}

func TestImageEditor_CropBySizeAndAlignment(t *testing.T) {
	type args struct {
		size      geom.Size
		alignment geom.Alignment
	}
	tests := []struct {
		name          string
		imageBounds   image.Rectangle
		args          args
		wantIsCropped bool
	}{
		{
			name:          "empty cropSize",
			imageBounds:   image.Rect(0, 0, 2, 2),
			args:          args{geom.NewSize(0, 0), geom.DefaultAlignment},
			wantIsCropped: false,
		},
		{
			name:          "imageSize is less than cropSize",
			imageBounds:   image.Rect(0, 0, 10, 10),
			args:          args{geom.NewSize(20, 20), geom.DefaultAlignment},
			wantIsCropped: false,
		},
		{
			name:          "imageSize is less than cropSize",
			imageBounds:   image.Rect(-1, -2, -100, -10),
			args:          args{geom.NewSize(20, 20), geom.DefaultAlignment},
			wantIsCropped: false,
		},
		{
			name:          "cropSize is less than imageSize",
			imageBounds:   image.Rect(2, 3, 15, 15),
			args:          args{geom.NewSize(10, 10), geom.DefaultAlignment},
			wantIsCropped: true,
		},
		{
			name:          "cropSize is less than imageSize",
			imageBounds:   image.Rect(-16, -14, -2, -1),
			args:          args{geom.NewSize(10, 10), geom.DefaultAlignment},
			wantIsCropped: true,
		},
		{
			name:          fmt.Sprintf("Alignment(%s, %s)", geom.LEFT, geom.TOP),
			imageBounds:   image.Rect(0, 0, 10, 10),
			args:          args{geom.NewSize(7, 6), geom.NewAlignment(geom.LEFT, geom.TOP)},
			wantIsCropped: true,
		},
		{
			name:          fmt.Sprintf("Alignment(%s, %s)", geom.RIGHT, geom.BOTTOM),
			imageBounds:   image.Rect(1, 1, 10, 10),
			args:          args{geom.NewSize(7, 3), geom.NewAlignment(geom.RIGHT, geom.BOTTOM)},
			wantIsCropped: true,
		},
		{
			name:          fmt.Sprintf("Alignment(%s, %s)", geom.CENTER, geom.CENTER),
			imageBounds:   image.Rect(0, 0, 10, 10),
			args:          args{geom.NewSize(7, 3), geom.NewAlignment(geom.CENTER, geom.CENTER)},
			wantIsCropped: true,
		},
		{
			name:          "incorrect image size",
			imageBounds:   image.Rect(2, 2, 10, 10),
			args:          args{geom.NewSize(7, 3), geom.NewAlignment(geom.CENTER, geom.CENTER)},
			wantIsCropped: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			editor := &ImageEditor{
				source:      image.NewRGBA(test.imageBounds),
				destination: image.NewRGBA(test.imageBounds),
				isCropped:   false,
			}
			editor.CropBySizeAndAlignment(test.args.size, test.args.alignment)

			wantBounds := test.imageBounds
			if test.wantIsCropped {
				wantBounds = editor.sizeAndAlignmentToRectangle(test.args.size, test.args.alignment)
			}

			assert.Equal(t, wantBounds, editor.destination.Rect,
				"incorrect value in ImageEditor.destination.Rect.\nimageBounds %v, size %v, alignment %v",
				test.imageBounds, test.args.size, test.args.alignment)
			assert.Equal(t, test.wantIsCropped, editor.isCropped,
				"incorrect value in ImageEditor.isCropped.\nimageBounds %v, size %v, alignment %v",
				test.imageBounds, test.args.size, test.args.alignment)
		})
	}
}

func TestImageEditor_BytesBuffer(t *testing.T) {
	type args struct {
		mimeType string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "encoding jpeg type",
			args:    args{MIMEJPEG},
			wantErr: false,
		},
		{
			name:    "encoding png type",
			args:    args{MIMEPNG},
			wantErr: false,
		},
		{
			name:    "encoding gif type",
			args:    args{"image/gif"},
			wantErr: true,
		},
		{
			name:    "encoding invalid type",
			args:    args{"beleberda"},
			wantErr: true,
		},
	}
	bounds := image.Rect(0, 0, 3, 3)
	source := new(MockImage)
	source.On("Bounds").Return(bounds)
	source.On("At", mock.Anything, mock.Anything).Return(color.RGBA{100, 100, 100, 255})
	source.On("ColorModel").Return(color.RGBAModel)
	destination := image.NewRGBA(bounds)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			editor := &ImageEditor{
				source:      source,
				destination: destination,
			}

			got, err := editor.BytesBuffer(test.args.mimeType)

			message := "mimeType = " + test.args.mimeType
			if test.wantErr {
				assert.Error(t, err, message)
			} else {
				assert.NoError(t, err, message)
				assert.NotEmpty(t, got.Len(),
					"ImageEditor.BytesBuffer(%v) should not return an empty value",
					test.args.mimeType,
				)
			}
		})
	}
}

func TestImageEditor_sizeAndAlignmentToRectangle(t *testing.T) {
	type args struct {
		size      geom.Size
		alignment geom.Alignment
	}
	tests := []struct {
		name        string
		imageBounds image.Rectangle
		args        args
		wantBounds  image.Rectangle
	}{
		{
			name:        "empty cropSize",
			imageBounds: image.Rect(0, 0, 10, 10),
			args:        args{geom.NewSize(0, 0), geom.NewAlignment(geom.LEFT, geom.TOP)},
			wantBounds:  image.Rect(0, 0, 0, 0),
		},
		{
			name:        "empty cropSize",
			imageBounds: image.Rect(-2, -1, 2, 2),
			args:        args{geom.NewSize(0, 0), geom.NewAlignment(geom.LEFT, geom.TOP)},
			wantBounds:  image.Rect(-2, -1, -2, -1),
		},
		{
			name:        "empty cropSize",
			imageBounds: image.Rect(-2, -1, 2, 2),
			args:        args{geom.NewSize(0, 0), geom.NewAlignment(geom.RIGHT, geom.BOTTOM)},
			wantBounds:  image.Rect(2, 2, 2, 2),
		},
		{
			name:        "empty cropSize",
			imageBounds: image.Rect(-4, -5, 2, 2),
			args:        args{geom.NewSize(0, 0), geom.NewAlignment(geom.CENTER, geom.CENTER)},
			wantBounds:  image.Rect(-1, -1, -1, -1),
		},
		{
			name:        "notmal cropSize",
			imageBounds: image.Rect(0, 0, 10, 10),
			args:        args{geom.NewSize(5, 5), geom.NewAlignment(geom.LEFT, geom.TOP)},
			wantBounds:  image.Rect(0, 0, 5, 5),
		},
		{
			name:        "notmal cropSize",
			imageBounds: image.Rect(-10, -10, 0, 0),
			args:        args{geom.NewSize(5, 7), geom.NewAlignment(geom.RIGHT, geom.BOTTOM)},
			wantBounds:  image.Rect(-5, -7, 0, 0),
		},
		{
			name:        "notmal cropSize",
			imageBounds: image.Rect(-7, -13, 3, 0),
			args:        args{geom.NewSize(5, 2), geom.NewAlignment(geom.CENTER, geom.CENTER)},
			wantBounds:  image.Rect(-4, -7, 1, -5),
		},
		{
			name:        "imageSize is less then cropimage",
			imageBounds: image.Rect(-3, -2, 12, 10),
			args:        args{geom.NewSize(100, 100), geom.NewAlignment(geom.CENTER, geom.CENTER)},
			wantBounds:  image.Rect(-3, -2, 12, 10),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			editor := &ImageEditor{
				source:      image.NewRGBA(test.imageBounds),
				destination: image.NewRGBA(test.imageBounds),
			}

			got := editor.sizeAndAlignmentToRectangle(test.args.size, test.args.alignment)

			assert.Equal(t, test.wantBounds, got,
				"incorrect value.\nImage bounds %v, size %v, alignment %v.",
				test.imageBounds, test.args.size, test.args.alignment,
			)
		})
	}
}
