package wkhtmltoimage

import (
	"testing"
)

func TestBuildParamsReturnsErrorIfNoInput(t *testing.T) {
	params := new(ImageOptions)

	_, err := buildParams(params)
	if err == nil {
		t.Error("Expected err to not be nil, got nil")
	}
	if err.Error() != "Must provide input" {
		t.Error("Expected Must provide input, got ", err.Error())
	}
}

func TestBuildParamsSetsDefaultParams(t *testing.T) {
	params := ImageOptions{Input: "http://example.com"}

	v, err := buildParams(&params)
	if err != nil {
		t.Error("Expected err to be nil, got ", err)
	}

	// sets quiet
	if v[0] != "-q" {
		t.Error("Expected -q, got ", v[0])
	}

	// sets disable-plugins
	if v[1] != "--disable-plugins" {
		t.Error("Expected --disable-plugins, got ", v[1])
	}

	// sets format
	if v[2] != "--format" {
		t.Error("Expected --format, got ", v[2])
	}
	if v[3] != "png" {
		t.Error("Expected png, got ", v[3])
	}

	// sets the input
	if v[4] != "http://example.com" {
		t.Error("Expected http://example.com, got ", v[4])
	}

	// sets the output
	if v[5] != "-" {
		t.Error("Expected -, got ", v[5])
	}
}

func TestBuildParamsSetsAllParams(t *testing.T) {
	params := ImageOptions{Input: "http://example.com", Format: "svg", Height: 600, Width: 800, Quality: 80, Output: "test.svg"}

	v, err := buildParams(&params)
	if err != nil {
		t.Error("Expected err to be nil, got ", err)
	}

	// sets quiet
	if v[0] != "-q" {
		t.Error("Expected -q, got ", v[0])
	}

	// sets disable-plugins
	if v[1] != "--disable-plugins" {
		t.Error("Expected --disable-plugins, got ", v[1])
	}

	// sets the format
	if v[2] != "--format" {
		t.Error("Expected --format, got ", v[2])
	}
	if v[3] != "svg" {
		t.Error("Expected svg, got ", v[3])
	}

	// sets the height
	if v[4] != "--height" {
		t.Error("Expected --height, got ", v[4])
	}
	if v[5] != "600" {
		t.Error("Expected 600, got ", v[5])
	}

	// sets the width
	if v[6] != "--width" {
		t.Error("Expected --width, got ", v[6])
	}
	if v[7] != "800" {
		t.Error("Expected 800, got ", v[7])
	}

	// sets the quality
	if v[8] != "--quality" {
		t.Error("Expected --quality, got ", v[8])
	}
	if v[9] != "80" {
		t.Error("Expected 80, got ", v[9])
	}

	// sets the input
	if v[10] != "http://example.com" {
		t.Error("Expected http://example.com, got ", v[10])
	}

	// sets the output
	if v[11] != "test.svg" {
		t.Error("Expected test.svg, got ", v[11])
	}
}

// TODO: Figure out how to get this to pass on travic ci
// func TestGetImageReturnsImage(t *testing.T) {
// 	c := ImageOptions{Input: "http://example.com"}
// 	output, err := GenerateImage(&c)

// 	if err != nil {
// 		t.Error("Expected err to be nil, got ", err.Error())
// 	}

// 	if output == nil {
// 		t.Error("Expected output to not be nil, got nil")
// 	}
// }

// this test has to be last cause it kills the env var - pretty hacky
func TestGetImageReturnsErrorIfNoBinaryPath(t *testing.T) {
	c := ImageOptions{Input: "http://example.com"}
	_, err := GenerateImage(&c)

	if err == nil {
		t.Error("Expected err to not be nil, got nil")
	}
	if err.Error() != "BinaryPath not set" {
		t.Error("Expected BinaryPath not set, got ", err.Error())
	}
}
