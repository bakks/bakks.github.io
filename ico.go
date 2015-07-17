// https://en.wikipedia.org/wiki/3D_projection
// https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// http://members.chello.at/~easyfilter/bresenham.html
// http://inside.mines.edu/fs_home/gmurray/ArbitraryAxisRotation/
// http://blog.andreaskahler.com/2009/06/creating-icosphere-mesh-in-code.html

package main

import "fmt"

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

type Point struct {
  Matrix
}

func (this *Point) X() float64 {
  return this.elements[0]
}

func (this *Point) Y() float64 {
  return this.elements[1]
}

func (this *Point) Z() float64 {
  return this.elements[2]
}

func NewPoint(x, y, z float64) *Point {
  arr := make([]float64, 3)
  arr[0] = x
  arr[1] = y
  arr[2] = z

  return &Point{
    Matrix{
      n: 1,
      m: 3,
      elements: []float64{x, y, z,},
    },
  }
}

func MultiplyMatrices(a, b *Matrix) *Matrix {
  if a.M() != b.N() {
    panic("Could not multiply non-matching matrices")
  }

  n := a.N()
  m := a.M()
  //p := b.M()
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

func main() {
  fmt.Println("hi\n")
}

