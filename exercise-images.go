package main

import ("golang.org/x/tour/pic"
	    "image/color"
		"image"
	   )

type Image struct{
	model color.Model
	w, h int
}

func (img Image) ColorModel() color.Model{
  return img.model
}

func (img Image) Bounds() image.Rectangle{
	return image.Rect(0, 0, img.w, img.h)
}

func (img Image) At(x, y int) color.Color{
	v := uint8(x^y - x*y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{color.RGBAModel, 255, 255}
	pic.ShowImage(m)
}
