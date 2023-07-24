package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAlignment(t *testing.T) {
	type args struct {
		vertical   string
		horizontal string
	}
	tests := []struct {
		name string
		args args
		want Alignment
	}{
		{
			name: "correct values",
			args: args{LEFT, TOP},
			want: Alignment{LEFT, TOP},
		},
		{
			name: "correct values",
			args: args{RIGHT, BOTTOM},
			want: Alignment{RIGHT, BOTTOM},
		},
		{
			name: "correct values",
			args: args{CENTER, CENTER},
			want: Alignment{CENTER, CENTER},
		},
		{
			name: "default values",
			args: args{DEFAULT_VERTICAL, DEFAULT_HORIZONTL},
			want: DefaultAlignment,
		},
		{
			name: "incorrect vertical",
			args: args{"middle", DEFAULT_HORIZONTL},
			want: DefaultAlignment,
		},
		{
			name: "incorrect horizontal",
			args: args{DEFAULT_VERTICAL, "tor"},
			want: DefaultAlignment,
		},
		{
			name: "incorrect vertical and horizontal",
			args: args{"middle", "middle"},
			want: DefaultAlignment,
		},
		{
			name: "empty vertical",
			args: args{"", DEFAULT_HORIZONTL},
			want: DefaultAlignment,
		},
		{
			name: "empty horizontal",
			args: args{DEFAULT_VERTICAL, ""},
			want: DefaultAlignment,
		},
		{
			name: "empty vertical and horizontal",
			args: args{"", ""},
			want: DefaultAlignment,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewAlignment(test.args.vertical, test.args.horizontal)
			assert.Equal(t, test.want, got,
				"NewAlignment(%q, %q) = %#v, want %#v",
				test.args.vertical, test.args.horizontal, got, test.want,
			)
		})
	}
}

func TestValidateVertical(t *testing.T) {
	tests := []struct {
		name     string
		vertical string
		want     bool
	}{
		{
			name:     "correct vertical",
			vertical: LEFT,
			want:     true,
		},
		{
			name:     "correct vertical",
			vertical: RIGHT,
			want:     true,
		},
		{
			name:     "correct vertical",
			vertical: CENTER,
			want:     true,
		},
		{
			name:     "incorrect vertical",
			vertical: TOP,
			want:     false,
		},
		{
			name:     "incorrect vertical",
			vertical: BOTTOM,
			want:     false,
		},
		{
			name:     "incorrect vertical",
			vertical: "beleberda",
			want:     false,
		},
		{
			name:     "empty vertical",
			vertical: "",
			want:     false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ValidateVertical(test.vertical)
			assert.Equal(t, test.want, got,
				"ValidateVertical(%q) = %t, want %t",
				test.vertical, got, test.want,
			)
		})
	}
}

func TestValidateHorizontal(t *testing.T) {
	tests := []struct {
		name       string
		horizontal string
		want       bool
	}{
		{
			name:       "correct horizontal",
			horizontal: TOP,
			want:       true,
		},
		{
			name:       "correct horizontal",
			horizontal: BOTTOM,
			want:       true,
		},
		{
			name:       "correct horizontal",
			horizontal: CENTER,
			want:       true,
		},
		{
			name:       "incorrect horizontal",
			horizontal: LEFT,
			want:       false,
		},
		{
			name:       "incorrect horizontal",
			horizontal: RIGHT,
			want:       false,
		},
		{
			name:       "incorrect horizontal",
			horizontal: "beleberda",
			want:       false,
		},
		{
			name:       "empty horizontal",
			horizontal: "",
			want:       false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ValidateHorizontal(test.horizontal)
			assert.Equal(t, test.want, got,
				"ValidateVertical(%q) = %t, want %t",
				test.horizontal, got, test.want,
			)
		})
	}
}

func TestAlignment_Vertical(t *testing.T) {
	testValues := []string{CENTER, LEFT, RIGHT}

	for _, testValue := range testValues {
		alignment := NewAlignment(testValue, DEFAULT_HORIZONTL)
		got := alignment.Vertical()
		assert.Equal(t, testValue, got,
			"%#v.Vertical() = %q, want %q",
			alignment, got, testValue,
		)
	}
}

func TestAlignment_SetVertical(t *testing.T) {
	testTabel := []struct {
		name     string
		vertical string
		want     string
	}{
		{
			name:     "correct vertical",
			vertical: LEFT,
			want:     LEFT,
		},
		{
			name:     "correct vertical",
			vertical: RIGHT,
			want:     RIGHT,
		},
		{
			name:     "correct vertical",
			vertical: CENTER,
			want:     CENTER,
		},
		{
			name:     "incorrect vertical",
			vertical: TOP,
			want:     DEFAULT_VERTICAL,
		},
		{
			name:     "incorrect vertical",
			vertical: BOTTOM,
			want:     DEFAULT_VERTICAL,
		},
		{
			name:     "incorrect vertical",
			vertical: "beleberda",
			want:     DEFAULT_VERTICAL,
		},
		{
			name:     "empty vertical",
			vertical: "",
			want:     DEFAULT_VERTICAL,
		},
	}

	alignment := DefaultAlignment
	for _, test := range testTabel {
		t.Run(test.name, func(t *testing.T) {
			alignment.SetVertical(test.vertical)
			assert.Equal(t, test.want, alignment.vertical,
				"Alignment.SetVertical(%q) set %q, want %q",
				test.vertical, alignment.vertical, test.want,
			)
		})
	}
}

func TestAlignment_Horizontal(t *testing.T) {
	testValues := []string{TOP, BOTTOM, CENTER}

	for _, testValue := range testValues {
		alignment := NewAlignment(DEFAULT_VERTICAL, testValue)
		got := alignment.Horizontal()
		assert.Equal(t, testValue, got,
			"%#v.Horizontal() = %q, want %q",
			alignment, got, testValue,
		)
	}
}

func TestAlignment_SetHorizontal(t *testing.T) {
	testTabel := []struct {
		name       string
		horizontal string
		want       string
	}{
		{
			name:       "correct horizontal",
			horizontal: TOP,
			want:       TOP,
		},
		{
			name:       "correct horizontal",
			horizontal: BOTTOM,
			want:       BOTTOM,
		},
		{
			name:       "correct horizontal",
			horizontal: CENTER,
			want:       CENTER,
		},
		{
			name:       "incorrect horizontal",
			horizontal: LEFT,
			want:       DEFAULT_HORIZONTL,
		},
		{
			name:       "incorrect horizontal",
			horizontal: RIGHT,
			want:       DEFAULT_HORIZONTL,
		},
		{
			name:       "incorrect horizontal",
			horizontal: "beleberda",
			want:       DEFAULT_HORIZONTL,
		},
		{
			name:       "empty horizontal",
			horizontal: "",
			want:       DEFAULT_HORIZONTL,
		},
	}

	alignment := DefaultAlignment
	for _, test := range testTabel {
		t.Run(test.name, func(t *testing.T) {
			alignment.SetHorizontal(test.horizontal)
			assert.Equal(t, test.want, alignment.horizontal,
				"Alignment.SetVertical(%q) set %q, want %q",
				test.horizontal, alignment.horizontal, test.want,
			)
		})
	}
}
