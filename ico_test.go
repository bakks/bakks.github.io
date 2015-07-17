package main

import "testing"
import "fmt"

func TestMultiplyMatrices(t *testing.T) {
  // test identity matrix multiplied by itself
  aVal := [][]float64 {
    []float64{1, 0},
    []float64{0, 1},
  }
  a := NewMatrix(aVal)

  if !MultiplyMatrices(a, a).Equals(a) {
    t.Fatal("one")
  }

  // test identity matrix multiplied by a simple matrix
  bVal := [][]float64 {
    []float64{1, 2},
    []float64{3, 4},
  }
  b := NewMatrix(bVal)

  if !MultiplyMatrices(a, b).Equals(b) {
    MultiplyMatrices(a, b).Print()
    t.Fatal("two")
  }

  // test more complex multiplication example
  cVal := [][]float64 {
    []float64{1, 2, 3},
    []float64{4, 5, 6},
  }
  c := NewMatrix(cVal)

  if c.N() != 3 && c.M() != 2 {
    t.Fatal("row/column mixup")
  }

  dVal := [][]float64 {
    []float64{7, 8,},
    []float64{9, 10,},
    []float64{11, 12,},
  }
  d := NewMatrix(dVal)

  eVal := [][]float64 {
    []float64{58, 64,},
    []float64{139, 154},
  }
  e := NewMatrix(eVal)

  c.Print()
  fmt.Printf("\n")
  d.Print()
  fmt.Printf("\n")
  e.Print()
  fmt.Printf("\n")

  if !MultiplyMatrices(c, d).Equals(e) {
    MultiplyMatrices(c, d).Print()
    t.Fatal("three")
  }
}

