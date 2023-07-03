// Package imageEditor provides functionality for editing images. It allowsyou
// to change pixels, crop images, and save them in various formats.
package imageEditor

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"mime"
	"runtime"
	"sync"

	"github.com/NooFreeNames/ImageEditor/pkg/imageEditor/geom"
	"github.com/NooFreeNames/ImageEditor/pkg/imageEditor/mods"
)

// Supported mime types
const (
	MIMEPNG  = "image/png"
	MIMEJPEG = "image/jpeg"
)

// IsSupportedImageFormat checks whether ImageEditor supports this image format
// or not.
func IsSupportedImageFormat(mimeType string) bool {
	switch mimeType {
	case MIMEPNG, MIMEJPEG:
		return true
	default:
		return false
	}
}

// NewImageEditor creates a new ImageEditor instance by decoding the image from
// the given io.Reader. It returns an error if the decoding fails.
func NewImageEditor(reader io.Reader) (*ImageEditor, error) {
	if reader == nil {
		return nil, errors.New("io.Reader is nil")
	}

	source, err := decode(reader)
	if err != nil {
		return nil, err
	}

	destination := image.NewRGBA(source.Bounds())
	return &ImageEditor{source, destination, false, false}, nil
}

// decode decodes an image from the given io.Reader and returns the image.Image.
// It returns an error if the decoding fails.
func decode(reader io.Reader) (image.Image, error) {
	img, ext, err := image.Decode(reader)
	if err != nil || !IsSupportedImageFormat(mime.TypeByExtension("."+ext)) {
		return nil, image.ErrFormat
	}

	return img, nil
}

// ImageEditor is an image editor that contains various tools for manipulating
// images. You can create a new object using NewImageEditor()
type ImageEditor struct {
	// source is original image that is loaded into the editor.
	source image.Image
	// destination is the final edited image that is produced from the original.
	destination *image.RGBA
	// isModifiedPixels is boolean indicating if the pixels of the image have.
	// been modified.
	isModifiedPixels bool
	// isCropped is boolean indicating if the image has been cropped.
	isCropped bool
}

// IsModifiedImage checks if the image has been modified.
func (editor *ImageEditor) IsModifiedImage() bool {
	return editor.isCropped || editor.isModifiedPixels
}

// EditedImage returns the edited image.
func (editor *ImageEditor) EditedImage() image.Image {
	var editedImage image.Image

	if !editor.IsModifiedImage() {
		// If the image has not been modified, the original source image will be
		// returned.
		editedImage = editor.source
	} else {
		if !editor.isModifiedPixels {
			// If the destination field's pixels are uninitialized, they will be
			// copied from the source field via mods.NewCopy().
			editor.ModifyPixels(mods.NewCopy())
		}
		editedImage = editor.destination
	}

	return editedImage
}

func (editor *ImageEditor) Size() geom.Size {
	return geom.NewSize(
		editor.destination.Rect.Dx(),
		editor.destination.Rect.Dy(),
	)
}

// ModifyPixels gets pixels from the source field, changes pixel values using
// mods.PixelModifier and sets new values in the destination field.
func (editor *ImageEditor) ModifyPixels(pixelModifer mods.PixelModifier) {
	if pixelModifer == nil {
		return
	}

	bounds := editor.destination.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	if editor.isModifiedPixels {
		// If pixels have been changed, we move the original field and create a
		// new destination field to modify pixels again.
		editor.source = editor.destination
		editor.destination = image.NewRGBA(bounds)
	}

	var waitGroup sync.WaitGroup
	groupCount := runtime.NumCPU()
	waitGroup.Add(groupCount)
	pixelCount := width * height
	pixelPerGroup := pixelCount / groupCount

	for currentGroup := 0; currentGroup < groupCount; currentGroup++ {
		go func(group int) {
			defer waitGroup.Done()

			var startPixel int = pixelPerGroup * group
			var endPixel int
			if group == groupCount-1 {
				endPixel = pixelCount
			} else {
				endPixel = startPixel + pixelPerGroup
			}

			for currentPixel := startPixel; currentPixel < endPixel; currentPixel++ {
				x := currentPixel%width + bounds.Min.X
				y := currentPixel/width + bounds.Min.Y
				oldColor := color.RGBAModel.Convert(
					editor.source.At(x, y)).(color.RGBA)
				newColor := pixelModifer.ModifyPixel(
					image.Pt(x, y),
					oldColor,
					editor.source,
				)
				editor.destination.SetRGBA(x, y, newColor)
			}
		}(currentGroup)
	}

	waitGroup.Wait()
	editor.isModifiedPixels = true
}

// Encode encodes the image and saves it using the specified writer. It returns
// an error if there is a problem with encoding the image.
func (editor *ImageEditor) Encode(writer io.Writer, mimeType string) error {
	switch mimeType {
	case MIMEPNG:
		return png.Encode(writer, editor.EditedImage())
	case MIMEJPEG:
		return jpeg.Encode(writer, editor.EditedImage(), nil)
	default:
		return image.ErrFormat
	}
}

// CropByRectangle crops the image to the given rectangle.
func (editor *ImageEditor) CropByRectangle(bounds image.Rectangle) {
	if bounds.Empty() {
		return
	}

	if !bounds.In(editor.destination.Rect) || bounds.Eq(editor.destination.Rect) {
		return
	}

	editor.destination = editor.destination.SubImage(bounds).(*image.RGBA)
	editor.isCropped = true
}

// CropBySizeAndAlignment crops the image according to the given geometry.Size
// and geometry.Alignment.
func (editor *ImageEditor) CropBySizeAndAlignment(size geom.Size,
	alignment geom.Alignment) {
	if size.IsEmpty() {
		return
	}

	rect := editor.sizeAndAlignmentToRectangle(size, alignment)
	editor.CropByRectangle(rect)
}

// BytesBuffer returns a byte  representation of the Edited Image, encoded using 
// the specified mimeType. It returns an error if the encoding fails.
func (editor *ImageEditor) BytesBuffer(mimeType string) (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)
	err := editor.Encode(buffer, mimeType)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}

// sizeAndAlignmentToRectangle converts the geometry.Size and geometry.Alignment
// into an image.Rectangle for the edited image.
func (editor *ImageEditor) sizeAndAlignmentToRectangle(size geom.Size,
	alignment geom.Alignment) image.Rectangle {
	bounds := editor.destination.Bounds()
	var rect image.Rectangle

	if size.Width() >= bounds.Dx() || size.Height() >= bounds.Dy() {
		return bounds
	}

	switch alignment.Vertical() {
	case geom.LEFT:
		rect.Min.X = bounds.Min.X
		rect.Max.X = bounds.Min.X + size.Width()
	case geom.RIGHT:
		rect.Min.X = bounds.Max.X - size.Width()
		rect.Max.X = bounds.Max.X
	default:
		indent := bounds.Dx() - size.Width()
		if indent%2 != 0 {
			rect.Min.X++
		}
		indent /= 2
		rect.Min.X += bounds.Min.X + indent
		rect.Max.X = bounds.Max.X - indent
	}

	switch alignment.Horizontal() {
	case geom.TOP:
		rect.Min.Y = bounds.Min.Y
		rect.Max.Y = bounds.Min.Y + size.Height()
	case geom.BOTTOM:
		rect.Min.Y = bounds.Max.Y - size.Height()
		rect.Max.Y = bounds.Max.Y
	default:
		indent := bounds.Dy() - size.Height()
		if indent%2 != 0 {
			rect.Min.Y++
		}
		indent /= 2
		rect.Min.Y += bounds.Min.Y + indent
		rect.Max.Y = bounds.Max.Y - indent
	}

	return rect
}
