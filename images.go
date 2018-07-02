package main

import (
	"image"
	"image/color"
	"golang.org/x/tour/pic"
)

type Image struct{
	width int
	height int
	red uint8
	green uint8
	blue uint8
	alpha uint8
}

func (im Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, im.width, im.height)
}

func (im Image) At(x, y int) color.Color {
	return color.RGBA{im.red, im.green, im.blue, im.alpha}
}

func main() {
	m := Image{40, 60, 30, 120, 200, 100}
	pic.ShowImage(m)
}
