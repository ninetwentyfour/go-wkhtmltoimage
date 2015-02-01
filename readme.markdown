# go-wkhtmltoimage

[![GoDoc](https://godoc.org/github.com/ninetwentyfour/go-wkhtmltoimage?status.png)](https://godoc.org/github.com/ninetwentyfour/go-wkhtmltoimage) [![Build Status](https://travis-ci.org/ninetwentyfour/go-wkhtmltoimage.svg)](https://travis-ci.org/ninetwentyfour/go-wkhtmltoimage)

go wrapper around [wkhtmltoimage](http://wkhtmltopdf.org/).

Not production ready.

## Install

### go-wkhtmltoimage

    go get github.com/ninetwentyfour/go-wkhtmltoimage

### wkhtmltoimage
Install by hand: [http://wkhtmltopdf.org/downloads.html](http://wkhtmltopdf.org/downloads.html)

## Examples
Get url and save image yourself.

    package main

    import (
      "fmt"
      "github.com/ninetwentyfour/go-wkhtmltoimage"
      "os"
    )

    func main() {
      c := wkhtmltoimage.ImageOptions{Input: "http://example.com", Format: "png"}
      out, _ := wkhtmltoimage.GenerateImage(&c)
      f, _ := os.Create("/tmp/example.png")
      defer f.Close()
      n2, _ := f.Write(out)
      fmt.Printf("wrote %d bytes\n", n2)
    }

Load a file and auto save the result

    package main

    import (
      "github.com/ninetwentyfour/go-wkhtmltoimage"
    )

    func main() {
      c := wkhtmltoimage.ImageOptions{Input: "/tmp/example.html", Format: "png", Output: "/tmp/example.png"}
      wkhtmltoimage.GenerateImage(&c)
    }

Save html string to image

    package main

    import (
      "github.com/ninetwentyfour/go-wkhtmltoimage"
    )

    func main() {
      html := "<html><head></head><body><p style='color:red;'>example</p></body></html>"
      c := wkhtmltoimage.ImageOptions{Input: "-", Format: "png", Output: "/tmp/example.png", Html: html}
      wkhtmltoimage.GenerateImage(&c)
    }

## Contributing:
1. Fork it
1. Create your feature branch (`git checkout -b my-new-feature`)
1. Commit your changes (`git commit -am 'Add some feature'`)
1. Push to the branch (`git push origin my-new-feature`)
1. Create new Pull Request