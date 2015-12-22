package icosahedron

import "testing"

func TestMultiplyMatrices(t *testing.T) {
	// test identity matrix multiplied by itself
	aVal := [][]float64{
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
	bVal := [][]float64{
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
	cVal := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}
	c := NewMatrix(cVal)

	if c.Width() != 3 && c.Height() != 2 {
		t.Fatal("row/column mixup")
	}

	dVal := [][]float64{
		{7, 8},
		{9, 10},
		{11, 12},
	}
	d := NewMatrix(dVal)

	eVal := [][]float64{
		{58, 64},
		{139, 154},
	}
	e := NewMatrix(eVal)

	x = MultiplyMatrices(c, d)
	if !x.Equals(e) {
		x.Print()
		t.Fatal("three")
	}

	fVal := [][]float64{
		{1},
		{1},
	}
	f := NewMatrix(fVal)

	gVal := [][]float64{
		{122},
		{293},
	}
	g := NewMatrix(gVal)

	x = MultiplyMatrices(e, f)
	if !x.Equals(g) {
		x.Print()
		t.Fatal("four")
	}
}

func TestAddMatrices(t *testing.T) {
	aVal := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}

	bVal := [][]float64{
		{7, 8, 9},
		{10, 11, 12},
	}

	cVal := [][]float64{
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

func TestCrossProduct(t *testing.T) {
	aVal := [][]float64{
		//{3, -3, 1},
		{3},
		{-3},
		{1},
	}

	bVal := [][]float64{
		//{4, 9, 2},
		{4},
		{9},
		{2},
	}

	cVal := [][]float64{
		//{-15, -2, 39},
		{-15},
		{-2},
		{39},
	}

	a := NewMatrix(aVal)
	b := NewMatrix(bVal)
	c := NewMatrix(cVal)

	x := CrossProductMatrices(a, b)
	if !x.Equals(c) {
		x.Print()
		t.Fatal("Cross product")
	}
}

//func TestProjection(t *testing.T) {
//	width := uint(80)
//	height := uint(40)
//	canvas := NewCanvas(height, width)
//	cube := MakeCube()
//	ProjectOntoCanvas(cube, canvas, height/2, width/2)
//}

