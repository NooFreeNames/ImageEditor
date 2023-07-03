// Package geometry provides utilities for dealing with geometric concepts.
package geom

const (
	LEFT   = "left"
	RIGHT  = "right"
	TOP    = "top"
	BOTTOM = "bottom"
	CENTER = "center"
)

const (
	DEFAULT_VERTICAL  = CENTER
	DEFAULT_HORIZONTL = CENTER
)

var DefaultAlignment = Alignment{DEFAULT_VERTICAL, DEFAULT_HORIZONTL}

// NewAlignment creates a new Alignment object based on the vertical and
// horizontal parameters. Valid vertical are "left", "right", and "center".
// Valid horizontal are "top", "bottom", and "center".
func NewAlignment(vertical, horizontal string) Alignment {
	alignment := Alignment{}
	alignment.SetVertical(vertical)
	alignment.SetHorizontal(horizontal)
	return alignment
}

// ValidateVertical checks whether a string value is a valid vertical
// alignment. Valid values are "left", "right", and "center".
func ValidateVertical(vertical string) bool {
	switch vertical {
	case LEFT, RIGHT, CENTER:
		return true
	default:
		return false
	}
}

// ValidateHorizontal checks whether a string value is a valid horizontal
// alignment. Valid values are "top", "bottom", and "center".
func ValidateHorizontal(horizontal string) bool {
	switch horizontal {
	case TOP, BOTTOM, CENTER:
		return true
	default:
		return false
	}
}

// Alignment is structure defines the vertical and horizontal alignment of an
// object.
type Alignment struct {
	// vertical is vertical alignment. Valid values are "left", "right", and
	// "center"
	vertical string
	// horizontal is horizontal alignment. Valid values are "top", "bottom", and
	// "center"
	horizontal string
}

// Vertical returns the value of the vertical field.
func (alignment Alignment) Vertical() string {
	return alignment.vertical
}

// SetVertical sets the value of the vertical field. If the value is not valid,
// the default value is set to CENTER.
func (alignment *Alignment) SetVertical(vertical string) {
	if !ValidateVertical(vertical) {
		alignment.vertical = DEFAULT_VERTICAL
	} else {
		alignment.vertical = vertical
	}
}

// Horizontal returns the value of the vertical field.
func (alignment Alignment) Horizontal() string {
	return alignment.horizontal
}

// SetHorizontal sets the value of the vertical field. If the value is not
// valid, the default value is set to CENTER.
func (alignment *Alignment) SetHorizontal(horizontal string) {
	if !ValidateHorizontal(horizontal) {
		alignment.horizontal = DEFAULT_HORIZONTL
	} else {
		alignment.horizontal = horizontal
	}
}
