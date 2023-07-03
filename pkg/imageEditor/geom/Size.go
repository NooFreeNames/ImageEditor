package geom

// NewSize creates a new Size object based on the width and height parameters.
func NewSize(width, height int) Size {
	size := Size{}
	size.SetHeight(height)
	size.SetWidth(width)
	return size
}

// Size represents the size of an object in two dimensions.
type Size struct {
	width, height int
}

// Width returns the value of the width field.
func (size Size) Width() int {
	return size.width
}

// SetWidth sets the value of the width field of a Size object. If the provided
// value is negative, it is set to 0.
func (size *Size) SetWidth(width int) {
	if width < 0 {
		width = 0
	}
	size.width = width
}

// Height returns the value of the height field.
func (size Size) Height() int {
	return size.height
}

// SetHeight sets the value of the height field of a Size object. If the
// provided value is negative, it is set to 0.
func (size *Size) SetHeight(height int) {
	if height < 0 {
		height = 0
	}
	size.height = height
}

// IsEmpty returns true if either the width or height field is 0. Otherwise, it
// returns false.
func (size Size) IsEmpty() bool {
	return size.width <= 0 || size.height <= 0
}
