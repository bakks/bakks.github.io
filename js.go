package main 

import "github.com/gopherjs/gopherjs/js"

func main() {
	js.Global.Set("ico", map[string]interface{}{
		"New": New,
	})
}

func New(width, height int) *js.Object {
	ico := Ico{
		Width: uint(width),
		Height: uint(height),
		Canvas: NewCanvas(uint(height), uint(width)),
		Model: MakeIcosahedron(1),
		XOffset: width / 2,
		YOffset: int(float64(height) / 1.5),
	}

	ico.Model.Scale(15)

	return js.MakeWrapper(&ico)
}

func (this *Ico) RotateX(f float64) {
	this.Model.RotateAroundZAxis(f)
}

func (this *Ico) RotateY(f float64) {
	this.Model.RotateAroundXAxis(f)
}

func (this *Ico) Render() string {
	_, edges := this.Model.CollectPointsAndEdges(true)
	ProjectEdgesOntoCanvas(edges, this.Canvas, this.YOffset, this.XOffset, nil, 0.4)

	str := this.Canvas.ToString()
	this.Canvas.Clear()
	return str
}

