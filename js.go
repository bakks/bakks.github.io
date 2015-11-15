package main 

import "github.com/gopherjs/gopherjs/js"

func main() {
	js.Global.Set("ico", render)
}

func render() string {
	width := uint(120)
	height := uint(60)
	canvas := NewCanvas(height, width)
	model := MakeIcosahedron(1)
	model.Scale(30)
	xOffset := int(width) / 2
	yOffset := int(float64(height) / 1.5)

	model.RotateAroundXAxis(0.02)
	//model.RotateAroundYAxis(0.08)

	_, edges := model.CollectPointsAndEdges(true)
	ProjectEdgesOntoCanvas(edges, canvas, yOffset, xOffset, nil)

	str := canvas.ToString()
	canvas.Clear()
	return str
}
