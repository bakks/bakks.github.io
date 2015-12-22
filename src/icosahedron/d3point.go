package icosahedron

import "math"

type D3Point struct {
	x, y, z float64
}

func NewD3PointFromMatrix(matrix *Matrix) *D3Point {
	return &D3Point{matrix.Get(0, 0), matrix.Get(0, 1), matrix.Get(0, 2)}
}

func (this *D3Point) X() float64 {
	return this.x
}

func (this *D3Point) SetX(x float64) {
	this.x = x
}

func (this *D3Point) Y() float64 {
	return this.y
}

func (this *D3Point) SetY(y float64) {
	this.y = y
}

func (this *D3Point) Z() float64 {
	return this.z
}

func (this *D3Point) SetZ(z float64) {
	this.z = z
}

func (this *D3Point) ToMatrix() *Matrix {
	val := [][]float64{
		{this.x},
		{this.y},
		{this.z},
	}
	return NewMatrix(val)
}

func (this *D3Point) Scale(factor float64) {
	this.x *= factor
	this.y *= factor
	this.z *= factor
}

func (this *D3Point) ScaleDim(xf, yf, zf float64) {
	this.x *= xf
	this.y *= yf
	this.z *= zf
}

func (this *D3Point) RotateAroundXAxis(angle float64) {
	y := this.Y()*math.Cos(angle) - this.Z()*math.Sin(angle)
	z := this.Y()*math.Sin(angle) + this.Z()*math.Cos(angle)

	this.SetY(y)
	this.SetZ(z)
}

func (this *D3Point) RotateAroundYAxis(angle float64) {
	x := this.Z()*math.Sin(angle) + this.X()*math.Cos(angle)
	z := this.Z()*math.Cos(angle) - this.X()*math.Sin(angle)

	this.SetX(x)
	this.SetZ(z)
}

func (this *D3Point) RotateAroundZAxis(angle float64) {
	x := this.X()*math.Cos(angle) - this.Y()*math.Sin(angle)
	y := this.X()*math.Sin(angle) + this.Y()*math.Cos(angle)

	this.SetX(x)
	this.SetY(y)
}

func (this *D3Point) GetMiddlePoint(that *D3Point) *D3Point {
	return &D3Point{
		0.5 * (that.x + this.x),
		0.5 * (that.y + this.y),
		0.5 * (that.z + this.z),
	}
}

func (this *D3Point) NormalizeDistanceToOrigin() {
	dist := math.Sqrt(this.x * this.x + this.y * this.y + this.z * this.z)
	this.x /= dist
	this.y /= dist
	this.z /= dist
}

