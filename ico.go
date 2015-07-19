// https://en.wikipedia.org/wiki/3D_projection
// https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// http://members.chello.at/~easyfilter/bresenham.html
// http://inside.mines.edu/fs_home/gmurray/ArbitraryAxisRotation/
// http://blog.andreaskahler.com/2009/06/creating-icosphere-mesh-in-code.html

package main

import "fmt"
import "math"
import "time"

func OrthographicProjection(point *D3Point, offsetX, offsetZ float64) *D2Point {
	aVal := [][]float64{
		{1, 0, 0},
		{0, 0, 1},
	}

	cVal := [][]float64{
		{offsetX},
		{offsetZ},
	}

	c := NewMatrix(cVal)
	a := NewMatrix(aVal)
	x := MultiplyMatrices(a, point.ToMatrix())
	x = AddMatrices(x, c)
	return NewD2PointFromMatrix(x)
}

var BULLET rune = '•'

func round(x float64) int {
	return int(math.Floor(x + 0.5))
}

func ProjectOntoCanvas(model *Model, canvas *Canvas, offsetRow, offsetColumn uint) {
	offsetX := float64(0)
	offsetZ := float64(-10)
	canvasYScale := 1.0
	canvasXScale := 2.2

	for _, point := range model.Points {
		projection := OrthographicProjection(point, offsetX, offsetZ)
		y := round(projection.Y()*canvasYScale) + int(offsetRow)
		x := round(projection.X()*canvasXScale) + int(offsetColumn)
		if x >= 0 && x < int(canvas.Width()) && y >= 0 && y < int(canvas.Height()) {
			(*canvas)[y][x] = BULLET
		}
	}
}

func main() {
	width := uint(80)
	height := uint(40)
	canvas := NewCanvas(height, width)
	cube := MakeCube()
	cube.Scale(4)

	for true {
		cube.RotateAroundXAxis(0.2)
		cube.RotateAroundYAxis(0.2)
		xOffset := width / 2
		yOffset := uint(float64(height) / 1.5)

		ProjectOntoCanvas(cube, canvas, yOffset, xOffset)
		canvas.Print()
		canvas.Clear()
		time.Sleep(250 * time.Millisecond)
		fmt.Println("------")
	}
}
