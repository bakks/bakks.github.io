package main

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
