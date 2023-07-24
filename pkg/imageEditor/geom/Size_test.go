package geom

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSize(t *testing.T) {
	type args struct {
		width  int
		height int
	}
	tests := []struct {
		name string
		args args
		want Size
	}{
		{
			name: "positive width and height",
			args: args{100, 100},
			want: Size{100, 100},
		},
		{
			name: "zero width and height",
			args: args{0, 0},
			want: Size{0, 0},
		},
		{
			name: "negative height",
			args: args{100, -100},
			want: Size{100, 0},
		},
		{
			name: "negative width",
			args: args{-100, 100},
			want: Size{0, 100},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewSize(test.args.width, test.args.height)
			assert.Equal(t, got, test.want, "NewSize(%q, %q) = %#v, wnat %#v",
			test.args.width, test.args.height, got, test.want)
		})
	}
}

func TestSize_Width(t *testing.T) {
	testValues := []int{0, 500, math.MaxInt, 9999}

	for _, testValue := range testValues {
		size := Size{
			width:  testValue,
			height: 0,
		}
		got := size.Width()
		assert.Equal(t, testValue, got, "%#v.Width() = %q, wnat %q",
		size, got, testValue)
	}
}

func TestSize_SetWidth(t *testing.T) {
	tests := []struct {
		name  string
		width int
		want  int
	}{
		{
			name:  "positive value",
			width: 500,
			want:  500,
		},
		{
			name:  "zero value",
			width: 0,
			want:  0,
		},
		{
			name:  "negative value",
			width: -500,
			want:  0,
		},
		{
			name:  "MaxInt value",
			width: math.MaxInt,
			want:  math.MaxInt,
		},
		{
			name:  "MinInt value",
			width: math.MinInt,
			want:  0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			size := &Size{}
			size.SetWidth(test.width)
			assert.Equal(t, test.want, size.width, 
				"Size.SetWidth(%q) set %q, want %q",
				test.width, size.width, test.want)
		})
	}
}

func TestSize_Height(t *testing.T) {
	testValues := []int{0, 500, math.MaxInt}

	for _, testValue := range testValues {
		size := Size{
			width:  0,
			height: testValue,
		}
		got := size.Height()
		assert.Equal(t, testValue, got, "%#v.Height() = %q, wnat %q",
		size, got, testValue)
	}
}

func TestSize_SetHeight(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  int
	}{
		{
			name:  "positive value",
			value: 500,
			want:  500,
		},
		{
			name:  "zero value",
			value: 0,
			want:  0,
		},
		{
			name:  "negative value",
			value: -500,
			want:  0,
		},
		{
			name:  "MaxInt value",
			value: math.MaxInt,
			want:  math.MaxInt,
		},
		{
			name:  "MinInt value",
			value: math.MinInt,
			want:  0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			size := Size{}
			size.SetHeight(test.value)
			assert.Equal(t, test.want, size.height, 
				"Size.SetHeight(%q) set %q, want %q",
				test.value, size.height, test.want)
		})
	}
}

func TestSize_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		size Size
		want bool
	}{
		{
			name: "zero width",
			size: Size{0, 100},
			want: true,
		},
		{
			name: "zero height",
			size: Size{100, 0},
			want: true,
		},
		{
			name: "zero width and height",
			size: Size{0, 0},
			want: true,
		},
		{
			name: "negative width",
			size: Size{-100, 100},
			want: true,
		},
		{
			name: "negative height",
			size: Size{100, -100},
			want: true,
		},
		{
			name: "negative width and height",
			size: Size{-100, -100},
			want: true,
		},
		{
			name: "positive width and height",
			size: Size{100, 100},
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.size.IsEmpty()
			assert.Equal(t, test.want, got, "%#v.IsEmpty() = %t, want %t",
				test.size, got, test.want)
		})
	}
}
