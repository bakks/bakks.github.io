package main

type D2Point struct {
	x, y float64
}

func (this *D2Point) X() float64 {
	return this.x
}

func (this *D2Point) Y() float64 {
	return this.y
}

func NewD2PointFromMatrix(matrix *Matrix) *D2Point {
	return &D2Point{matrix.Get(0, 0), matrix.Get(0, 1)}
}
