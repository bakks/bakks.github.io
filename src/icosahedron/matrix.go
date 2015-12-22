package icosahedron

import "fmt"

type Matrix struct {
	n        uint
	m        uint
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
		n:        n,
		m:        m,
		elements: make([]float64, n*m),
	}

}

func (this *Matrix) Width() uint {
	return this.n
}

func (this *Matrix) Height() uint {
	return this.m
}

func (this *Matrix) Get(n, m uint) float64 {
	return this.elements[n+this.n*m]
}

func (this *Matrix) Put(n, m uint, value float64) {
	this.elements[n+this.n*m] = value
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

func addSubMatrices(a, b *Matrix, add bool) *Matrix {
	if a.Height() != b.Height() || a.Width() != b.Width() {
		panic("Cannot add matrices of different sizes")
	}

	n := a.Width()
	m := a.Height()
	x := NewEmptyMatrix(n, m)

	if add {
		for j := uint(0); j < m; j++ {
			for i := uint(0); i < n; i++ {
				x.Put(i, j, a.Get(i, j) + b.Get(i, j))
			}
		}
	} else {
		for j := uint(0); j < m; j++ {
			for i := uint(0); i < n; i++ {
				x.Put(i, j, a.Get(i, j) - b.Get(i, j))
			}
		}
	}

	return x
}

func AddMatrices(a, b *Matrix) *Matrix {
	return addSubMatrices(a, b, true)
}

func SubMatrices(a, b *Matrix) *Matrix {
	return addSubMatrices(a, b, false)
}

func CrossProductMatrices(a, b *Matrix) *Matrix {
	if a.Height() != 3 || b.Height() != 3 || a.Width() != 1 || b.Width() != 1 {
		panic("Can only crossproduct 2 3x1 matrices")
	}

	c := NewMatrix([][]float64{
		{ a.Get(1, 0) * b.Get(2, 0) - a.Get(2, 0) * b.Get(1, 0) },
		{ (a.Get(0, 0) * b.Get(2, 0) - a.Get(2, 0) * b.Get(0, 0)) * -1 },
		{ a.Get(0, 0) * b.Get(1, 0) - a.Get(1, 0) * b.Get(0, 0) },
	})

	return c
}

