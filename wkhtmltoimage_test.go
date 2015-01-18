package wkhtmltoimage

import (
	"os"
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

	// sets the input
	if v[1] != "http://example.com" {
		t.Error("Expected http://example.com, got ", v[1])
	}

	// sets the output
	if v[2] != "-" {
		t.Error("Expected -, got ", v[2])
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

	// sets the format
	if v[1] != "--format" {
		t.Error("Expected --format, got ", v[1])
	}
	if v[2] != "svg" {
		t.Error("Expected svg, got ", v[2])
	}

	// sets the height
	if v[3] != "--height" {
		t.Error("Expected --height, got ", v[3])
	}
	if v[4] != "600" {
		t.Error("Expected 600, got ", v[4])
	}

	// sets the width
	if v[5] != "--width" {
		t.Error("Expected --width, got ", v[5])
	}
	if v[6] != "800" {
		t.Error("Expected 800, got ", v[5])
	}

	// sets the quality
	if v[7] != "--quality" {
		t.Error("Expected --quality, got ", v[7])
	}
	if v[8] != "80" {
		t.Error("Expected 80, got ", v[8])
	}

	// sets the input
	if v[9] != "http://example.com" {
		t.Error("Expected http://example.com, got ", v[9])
	}

	// sets the output
	if v[10] != "test.svg" {
		t.Error("Expected test.svg, got ", v[10])
	}
}

func TestGetImageReturnsImage(t *testing.T) {
	c := ImageOptions{Input: "http://example.com"}
	output, err := GenerateImage(&c)

	if err != nil {
		t.Error("Expected err to be nil, got ", err.Error())
	}

	if output == nil {
		t.Error("Expected output to not be nil, got nil")
	}
}

// this test has to be last cause it kills the env var - pretty hacky
func TestGetImageReturnsErrorIfNoBinaryPath(t *testing.T) {
	os.Setenv("WKHTMLTOIMAGE_PATH", "")
	c := ImageOptions{Input: "http://example.com"}
	_, err := GenerateImage(&c)

	if err == nil {
		t.Error("Expected err to not be nil, got nil")
	}
	if err.Error() != "WKHTMLTOIMAGE_PATH env var not set" {
		t.Error("Expected WKHTMLTOIMAGE_PATH env var not set, got ", err.Error())
	}
}
