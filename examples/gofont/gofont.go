package main

import (
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/photostorm/gg"
	"golang.org/x/image/draw"
	"golang.org/x/image/font/gofont/goregular"
)

func main() {
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	face := truetype.NewFace(font, &truetype.Options{Size: 48})

	dc := gg.NewContext(1024, 1024)
	dc.SetFontFace(face, 96)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored("Hello, world!", 512, 512, 0.5, 0.5, draw.BiLinear, nil, false)
	dc.SavePNG("out.png")
}
