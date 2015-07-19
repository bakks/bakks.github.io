package main

import "fmt"

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
