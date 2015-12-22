package main

import "icosahedron"
import "github.com/gopherjs/gopherjs/js"

func main() {
	js.Global.Set("ico", map[string]interface{}{
		"New": New,
	})
}

func New(width, height, xOffset, yOffset, joints int, scale float64) *js.Object {
	ico := icosahedron.Ico{
		Width: uint(width),
		Height: uint(height),
		Canvas: icosahedron.NewCanvas(uint(height), uint(width)),
		Model: icosahedron.MakeIcosahedron(1),
		XOffset: xOffset,
		YOffset: yOffset,
	}

	ico.Model.Scale(scale)

	return js.MakeWrapper(&ico)
}

