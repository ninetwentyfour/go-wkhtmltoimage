// Package wkhtmltoimage provides a simple wrapper around wkhtmltoimage (http://wkhtmltopdf.org/) binary.
package wkhtmltoimage

import (
	"errors"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// ImageOptions represent the options to generate the image.
type ImageOptions struct {
	// Input is the content to turn into an image. REQUIRED
	//
	// Can be a url (http://example.com), a local file (/tmp/example.html), or html as a string (send "-" and set the Html value)
	Input string
	// Format is the type of image to generate
	//
	// jpg, png, svg, bmp supported. Defaults to local wkhtmltoimage default
	Format string
	// Height is the height of the screen used to render in pixels.
	//
	// Default is calculated from page content. Default 0 (renders entire page top to bottom)
	Height int
	// Width is the width of the screen used to render in pixels.
	//
	// Note that this is used only as a guide line. Default 1024
	Width int
	// Quality determines the final image quality.
	//
	// Values supported between 1 and 100. Default is 94
	Quality int
	// Html is a string of html to render into and image.
	//
	// Only needed to be set if Input is set to "-"
	Html string
	// Output controls how to save or return the image.
	//
	// Leave nil to return a []byte of the image. Set to a path (/tmp/example.png) to save as a file.
	Output string
}

// GenerateImage creates an image from an input.
// It returns the image ([]byte) and any error encountered.
func GenerateImage(options *ImageOptions) ([]byte, error) {
	arr, err := buildParams(options)
	if err != nil {
		return []byte{}, err
	}

	execPath := os.Getenv("WKHTMLTOIMAGE_PATH")
	if execPath == "" {
		return []byte{}, errors.New("WKHTMLTOIMAGE_PATH env var not set")
	}

	cmd := exec.Command(execPath, arr...)

	if options.Html != "" {
		cmd.Stdin = strings.NewReader(options.Html)
	}

	output, _ := cmd.CombinedOutput()

	return output, nil
}

// buildParams takes the image options set by the user and turns them into command flags for wkhtmltoimage
// It returns an array of command flags.
func buildParams(options *ImageOptions) ([]string, error) {
	a := []string{}

	if options.Input == "" {
		return []string{}, errors.New("Must provide input")
	}

	// silence extra wkhtmltoimage output
	a = append(a, "-q")

	if options.Format != "" {
		a = append(a, "--format")
		a = append(a, options.Format)
	}

	if options.Height != 0 {
		a = append(a, "--height")
		a = append(a, strconv.Itoa(options.Height))
	}

	if options.Width != 0 {
		a = append(a, "--width")
		a = append(a, strconv.Itoa(options.Width))
	}

	if options.Quality != 0 {
		a = append(a, "--quality")
		a = append(a, strconv.Itoa(options.Quality))
	}

	// url and output come last
	if options.Input != "-" {
		// make sure we dont pass stdin if we aren't expecting it
		options.Html = ""
	}

	a = append(a, options.Input)

	if options.Output == "" {
		a = append(a, "-")
	} else {
		a = append(a, options.Output)
	}

	return a, nil
}
