// https://en.wikipedia.org/wiki/3D_projection
// https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// http://members.chello.at/~easyfilter/bresenham.html
// http://inside.mines.edu/fs_home/gmurray/ArbitraryAxisRotation/
// http://blog.andreaskahler.com/2009/06/creating-icosphere-mesh-in-code.html

package main

import "fmt"
import "math"

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

func (this *Matrix) N() uint {
  return this.n
}

func (this *Matrix) M() uint {
  return this.m
}

func (this *Matrix) Get(n, m uint) float64 {
  return this.elements[n + this.n * m]
}

func (this *Matrix) Put(n, m uint, value float64) {
  this.elements[n + this.n * m] = value
}

func (this *Matrix) Equals(that *Matrix) bool {
  if this.n != that.N() || this.m != that.M() {
    return false
  }

  for j := uint(0); j < this.m; j++ {
    for i := uint(0); i < this.n; i++ {
      if this.Get(i, j) != that.Get(i, j) {
        return false
      }
    }
  }

  return true
}

func (this *Matrix) Print() {
  for j := uint(0); j < this.m; j++ {
    for i := uint(0); i < this.n; i++ {
      fmt.Printf("%.2f ", this.Get(i, j))
    }
    fmt.Printf("\n")
  }
}

type D3Point struct {
  Matrix
}

func (this *D3Point) X() float64 {
  return this.elements[0]
}

func (this *D3Point) SetX(x float64) {
  this.elements[0] = x
}

func (this *D3Point) Y() float64 {
  return this.elements[1]
}

func (this *D3Point) SetY(y float64) {
  this.elements[0] = y
}

func (this *D3Point) Z() float64 {
  return this.elements[2]
}

func (this *D3Point) SetZ(z float64) {
  this.elements[0] = z
}

func NewD3Point(x, y, z float64) *D3Point {
  return &D3Point{
    Matrix{
      n: 1,
      m: 3,
      elements: []float64{x, y, z,},
    },
  }
}

type D2Point struct {
  Matrix
}

func (this *D2Point) X() float64 {
  return this.elements[0]
}

func (this *D2Point) Y() float64 {
  return this.elements[1]
}

func NewD2Point(x, y float64) *D2Point {
  return &D2Point{
    Matrix{
      n: 1,
      m: 2,
      elements: []float64{x, y},
    },
  }
}

func MultiplyMatrices(a, b *Matrix) *Matrix {
  if a.M() != b.N() {
    panic("Could not multiply non-matching matrices")
  }

  n := a.N()
  m := a.M()
  x := NewEmptyMatrix(m, m)

  for i := uint(0); i < m; i++ {
    for j := uint(0); j < m; j++ {
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
  if a.M() != b.M() || a.N() != b.N() {
  }

  n := a.N()
  m := a.M()
  x := NewEmptyMatrix(a.N(), a.M())

  for j := uint(0); j < m; j++ {
    for i := uint(0); i < n; i++ {
      x.Put(i, j, a.Get(i, j) + b.Get(i, j))
    }
  }

  return x
}

type Edge struct {
  start *D3Point
  end *D3Point
}

type Model struct {
  points []*D3Point
  edges []Edge
}

func RotatePointAroundXAxis(point *D3Point, angle float64) {
  y := point.Y() * math.Cos(angle) - point.Z() * math.Sin(angle)
  z := point.Y() * math.Sin(angle) + point.Z() * math.Cos(angle)

  point.SetY(y)
  point.SetZ(z)
}

func RotatePointAroundYAxis(point *D3Point, angle float64) {
  x := point.Z() * math.Sin(angle) + point.X() * math.Cos(angle)
  z := point.Z() * math.Cos(angle) + point.X() * math.Sin(angle)

  point.SetX(x)
  point.SetZ(z)
}

func main() {
  fmt.Println("hi\n")
}

