package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{
	w, h int
	cA, cB uint8
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i *Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), i.cA, i.cB}
}

func exerciseImages() {
	m := &Image{100,100, 250,250}
	pic.ShowImage(m)
}