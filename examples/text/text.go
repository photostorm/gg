package main

import (
	"github.com/photostorm/gg"
	"golang.org/x/image/draw"
)

func main() {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("/Library/Fonts/Arial.ttf", 96, 96); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("Hello, world!", S/2, S/2, 0.5, 0.5, draw.BiLinear, nil, false)
	dc.SavePNG("out.png")
}
