// Package handlers provides request handlers for the server package.
package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/NooFreeNames/ImageEditor/internal/server/utils"
	"github.com/NooFreeNames/ImageEditor/pkg/imageEditor"
	"github.com/NooFreeNames/ImageEditor/pkg/imageEditor/geom"
	"github.com/NooFreeNames/ImageEditor/pkg/imageEditor/mods"
)

// pingHandler is the handler function for the "/image" URL.
// Edits the image and writes the result to the request body, and in case of an
// error writes the error text to the request body. All errors are logged.
//
// Read values of the POST request:
//   - image - image file to edit
//   - width - crop width
//   - height - crop height
//   - vertical - crop vertical position
//   - horizontal - crop horizontal position
//   - filter - image filter
//   - blure_sigma - degree of blur
func ImageHandler(response http.ResponseWriter, request *http.Request) {
	file, meta, err := request.FormFile("image")
	if err != nil {
		http.Error(response, "Could not read the file", http.StatusBadRequest)
		log.Println("Failed to parse 'image' parameter: ", err)
		return
	}
	defer file.Close()
	contentType := meta.Header.Get("Content-Type")

	editor, err := imageEditor.NewImageEditor(file)
	if err != nil {
		http.Error(response, "Unsupported file format",
			http.StatusUnsupportedMediaType)
		log.Println("Failed to create ImageEditor: ", err)
		return
	}

	width, err := utils.ParsePositiveInt(request.FormValue("width"))
	if err != nil {
		utils.LogAndWriteError(response,
			"The width must be a positive integer",
			http.StatusBadRequest)
		return
	}

	height, err := utils.ParsePositiveInt(request.FormValue("height"))
	if err != nil {
		utils.LogAndWriteError(response,
			"The height must be a positive integer",
			http.StatusBadRequest)
		return
	}

	size := geom.NewSize(width, height)
	if !size.IsEmpty() {
		imageSize := editor.Size()

		if size.Width() > imageSize.Width() {
			utils.LogAndWriteError(response,
				"The width should not be greater than the width of the image",
				http.StatusBadRequest)
			return
		}

		if size.Height() > imageSize.Height() {
			utils.LogAndWriteError(response,
				"The height should not be greater than the height of the image",
				http.StatusBadRequest)
			return
		}

		vertical := request.FormValue("vertical")
		if !geom.ValidateVertical(vertical) {
			utils.LogAndWriteError(response,
				"Incorrect vertical value",
				http.StatusBadRequest)
			return
		}
		horizontal := request.FormValue("horizontal")
		if !geom.ValidateHorizontal(horizontal) {
			utils.LogAndWriteError(response,
				"Incorrect horizontal value",
				http.StatusBadRequest)
			return
		}

		alignment := geom.NewAlignment(vertical, horizontal)
		editor.CropBySizeAndAlignment(size, alignment)
	}

	filter := request.FormValue("filter")
	switch filter {
	case "grayscale":
		editor.ModifyPixels(mods.NewGrayscale())
	case "negative":
		editor.ModifyPixels(mods.NewNegative())
	case "blure":
		sigma, _ := strconv.ParseFloat(request.FormValue("blure_sigma"), 64)
		if sigma <= 0 {
			sigma = 2
		}
		editor.ModifyPixels(mods.NewGaussianBlur(sigma))
	default:
		if filter != "" {
			utils.LogAndWriteError(response,
				"Invalid filter value",
				http.StatusBadRequest)
			return
		}
	}

	buff, err := editor.BytesBuffer(contentType)
	if err != nil {
		http.Error(response, "File decoding error",
			http.StatusInternalServerError)
		log.Println("Failed to get image bytes: ", err)
		return
	}

	response.Header().Set("Content-Type", contentType)
	response.Header().Set("Content-Length", strconv.Itoa(buff.Len()))
	response.Write(buff.Bytes())
}
