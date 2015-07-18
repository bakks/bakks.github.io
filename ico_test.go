package main

import "testing"

func TestMultiplyMatrices(t *testing.T) {
  // test identity matrix multiplied by itself
  aVal := [][]float64 {
    {1, 0},
    {0, 1},
  }
  a := NewMatrix(aVal)

  x := MultiplyMatrices(a, a)
  if !x.Equals(a) {
    x.Print()
    t.Fatal("one")
  }

  // test identity matrix multiplied by a simple matrix
  bVal := [][]float64 {
    {1, 2},
    {3, 4},
  }
  b := NewMatrix(bVal)

  x = MultiplyMatrices(a, b)
  if !x.Equals(b) {
    x.Print()
    t.Fatal("two")
  }

  // test more complex multiplication example
  cVal := [][]float64 {
    {1, 2, 3},
    {4, 5, 6},
  }
  c := NewMatrix(cVal)

  if c.N() != 3 && c.M() != 2 {
    t.Fatal("row/column mixup")
  }

  dVal := [][]float64 {
    {7, 8},
    {9, 10},
    {11, 12},
  }
  d := NewMatrix(dVal)

  eVal := [][]float64 {
    {58, 64},
    {139, 154},
  }
  e := NewMatrix(eVal)

  x = MultiplyMatrices(c, d)
  if !x.Equals(e) {
    x.Print()
    t.Fatal("three")
  }
}

func TestAddMatrices(t *testing.T) {
  aVal := [][]float64 {
    {1, 2, 3},
    {4, 5, 6},
  }

  bVal := [][]float64 {
    {7, 8, 9},
    {10, 11, 12},
  }

  cVal := [][]float64 {
    {8, 10, 12},
    {14, 16, 18},
  }

  a := NewMatrix(aVal)
  b := NewMatrix(bVal)
  c := NewMatrix(cVal)

  x := AddMatrices(a, b)
  if !x.Equals(c) {
    x.Print()
    t.Fatal("Add matrices")
  }
}

