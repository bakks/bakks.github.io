// https://en.wikipedia.org/wiki/3D_projection
// https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// http://members.chello.at/~easyfilter/bresenham.html
// http://inside.mines.edu/fs_home/gmurray/ArbitraryAxisRotation/
// http://blog.andreaskahler.com/2009/06/creating-icosphere-mesh-in-code.html

package main

import "fmt"
import "math"
import "time"

type Matrix struct {
  n uint
  m uint
  elements []float64
}

func NewMatrix(values [][]float64) *Matrix {
  m := uint(len(values))
  n := uint(len(values[0]))
  x := NewEmptyMatrix(n, m)

  for j, row := range values {
    if uint(len(row)) != n {
      panic("Not enough columns in value")
    }

    for i, value := range row {
      x.Put(uint(i), uint(j), value)
    }
  }

  return x
}

func NewEmptyMatrix(n, m uint) *Matrix {
  return &Matrix{
    n: n,
    m: m,
    elements: make([]float64, n * m),
  }

}

func (this *Matrix) Width() uint {
  return this.n
}

func (this *Matrix) Height() uint {
  return this.m
}

func (this *Matrix) Get(n, m uint) float64 {
  return this.elements[n + this.n * m]
}

func (this *Matrix) Put(n, m uint, value float64) {
  this.elements[n + this.n * m] = value
}

func (this *Matrix) Equals(that *Matrix) bool {
  if this.Width() != that.Width() || this.Height() != that.Height() {
    return false
  }

  for j := uint(0); j < this.Height(); j++ {
    for i := uint(0); i < this.Width(); i++ {
      if this.Get(i, j) != that.Get(i, j) {
        return false
      }
    }
  }

  return true
}

func (this *Matrix) Print() {
  for j := uint(0); j < this.Height(); j++ {
    for i := uint(0); i < this.Width(); i++ {
      fmt.Printf("%.2f ", this.Get(i, j))
    }
    fmt.Printf("\n")
  }
}

type D3Point struct {
  x, y, z float64
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

func NewD3PointFromMatrix(matrix *Matrix) *D3Point {
  return &D3Point{matrix.Get(0, 0), matrix.Get(0, 1), matrix.Get(0, 2)}
}

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

func MultiplyMatrices(a, b *Matrix) *Matrix {
  if a.Width() != b.Height() {
    panic("Could not multiply non-matching matrices")
  }

  n := a.Width()
  m := a.Height()
  x := NewEmptyMatrix(b.Width(), m)

  for i := uint(0); i < m; i++ {
    for j := uint(0); j < b.Width(); j++ {
      var sum float64 = 0

      for k := uint(0); k < n; k++ {
        sum += a.Get(k, i) * b.Get(j, k)
      }

      x.Put(j, i, sum)
    }
  }

  return x
}

func AddMatrices(a, b *Matrix) *Matrix {
  if a.Height() != b.Height() || a.Width() != b.Width() {
    panic("Cannot add matrices of different sizes")
  }

  n := a.Width()
  m := a.Height()
  x := NewEmptyMatrix(n, m)

  for j := uint(0); j < m; j++ {
    for i := uint(0); i < n; i++ {
      x.Put(i, j, a.Get(i, j) + b.Get(i, j))
    }
  }

  return x
}

type Edge struct {
  Start *D3Point
  End *D3Point
}

type Model struct {
  Points []*D3Point
  Edges []Edge
}

func NewModel() *Model {
  return &Model{
    make([]*D3Point, 0),
    make([]Edge, 0),
  }
}

func (this *Model) Scale(factor float64) {
  for i := 0; i < len(this.Points); i++ {
    this.Points[i].Scale(factor)
  }
}

func (this *Model) ScaleDim(xf, yf, zf float64) {
  for i := 0; i < len(this.Points); i++ {
    this.Points[i].ScaleDim(xf, yf, zf)
  }
}

func (this *D3Point) RotateAroundXAxis(angle float64) {
  y := this.Y() * math.Cos(angle) - this.Z() * math.Sin(angle)
  z := this.Y() * math.Sin(angle) + this.Z() * math.Cos(angle)

  this.SetY(y)
  this.SetZ(z)
}

func (this *D3Point) RotateAroundYAxis(angle float64) {
  x := this.Z() * math.Sin(angle) - this.X() * math.Cos(angle)
  z := this.Z() * math.Cos(angle) + this.X() * math.Sin(angle)

  this.SetX(x)
  this.SetZ(z)
}

func (this *Model) RotateAroundXAxis(angle float64) {
  for i := 0; i < len(this.Points); i++ {
    this.Points[i].RotateAroundXAxis(angle)
  }
}

func (this *Model) RotateAroundYAxis(angle float64) {
  for i := 0; i < len(this.Points); i++ {
    this.Points[i].RotateAroundYAxis(angle)
  }
}

func MakeCube() *Model {
  x := NewModel()

  x.Points = append(x.Points, &D3Point{-1, 1, -1})
  x.Points = append(x.Points, &D3Point{1, 1, -1})
  x.Points = append(x.Points, &D3Point{-1, -1, -1})
  x.Points = append(x.Points, &D3Point{1, -1, -1})
  x.Points = append(x.Points, &D3Point{-1, 1, 1})
  x.Points = append(x.Points, &D3Point{1, 1, 1})
  x.Points = append(x.Points, &D3Point{-1, -1, 1})
  x.Points = append(x.Points, &D3Point{1, -1, 1})

  x.Edges = append(x.Edges, Edge{x.Points[0], x.Points[1]})
  x.Edges = append(x.Edges, Edge{x.Points[1], x.Points[3]})
  x.Edges = append(x.Edges, Edge{x.Points[3], x.Points[2]})
  x.Edges = append(x.Edges, Edge{x.Points[2], x.Points[0]})
  x.Edges = append(x.Edges, Edge{x.Points[4], x.Points[5]})
  x.Edges = append(x.Edges, Edge{x.Points[5], x.Points[7]})
  x.Edges = append(x.Edges, Edge{x.Points[7], x.Points[6]})
  x.Edges = append(x.Edges, Edge{x.Points[6], x.Points[4]})
  x.Edges = append(x.Edges, Edge{x.Points[0], x.Points[4]})
  x.Edges = append(x.Edges, Edge{x.Points[1], x.Points[5]})
  x.Edges = append(x.Edges, Edge{x.Points[3], x.Points[7]})
  x.Edges = append(x.Edges, Edge{x.Points[2], x.Points[6]})

  return x
}

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

type Canvas [][]rune

func NewCanvas(rows, columns uint) *Canvas {
  var x Canvas = make([][]rune, rows)
  for i := range x {
    x[i] = make([]rune, columns)
    for j := range x[i] {
      x[i][j] = ' '
    }
  }

  return &x
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
    y := round(projection.Y() * canvasYScale) + int(offsetRow)
    x := round(projection.X() * canvasXScale) + int(offsetColumn)
    if x >= 0 && x < int(canvas.Width()) && y >= 0 && y < int(canvas.Height()) {
      (*canvas)[y][x] = BULLET
    }
  }
}

func (this *Canvas) Width() uint {
  return uint(len((*this)[0]))
}

func (this *Canvas) Height() uint {
  return uint(len(*this))
}

func (this *Canvas) Clear() {
  for j := 0; j < int(this.Height()); j++ {
    for i := 0; i < int(this.Width()); i++ {
      (*this)[j][i] = ' '
    }
  }
}

func (this *Canvas) Print() {
  fmt.Print(this.ToString())
}

func (this *Canvas) ToString() string {
  str := ""

  for y := uint(0); y < this.Height(); y++ {
    for x := uint(0); x < this.Width(); x++ {
      str += string((*this)[y][x])
    }
    str += "\n"
  }

  return str
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

