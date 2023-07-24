// Package utils provides various utility functions for the server package.
package utils

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// LogAndWriteError logs an error message and writes it as an HTTP response with
// a specific code.
func LogAndWriteError(response http.ResponseWriter, message string, code int) {
	http.Error(response, message, code)
	log.Println("Error:", message)
}

// ParsePositiveInt converts a string to a positive integer. If the input string
// is empty, it return 0. Returns an error if it is not possible to convert a
// string to a positive integer
func ParsePositiveInt(str string) (int, error) {
	if str == "" {
		return 0, nil
	}

	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("%q is not a number", str)
	}

	if num < 0 {
		return 0, fmt.Errorf("%q is less than zero", str)
	}
	return num, err
}
