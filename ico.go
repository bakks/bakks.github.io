// https://en.wikipedia.org/wiki/3D_projection
// https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// http://members.chello.at/~easyfilter/bresenham.html
// http://inside.mines.edu/fs_home/gmurray/ArbitraryAxisRotation/
// http://blog.andreaskahler.com/2009/06/creating-icosphere-mesh-in-code.html

package main

import "fmt"
import "math"
import "time"
import "sort"
//import "strconv"
import "github.com/nsf/termbox-go"

var POINT rune = '•'
var EDGE rune = '·'
var BLANK rune = ' '
var CANVASXSCALE float64 = 2.2
var CANVASYSCALE float64 = 1.0

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

func round(x float64) int {
	return int(math.Floor(x + 0.5))
}

func equals(a, b, delta float64) bool {
	return a < b+delta && a > b-delta
}

func sign(x int) int {
	if x >= 0.0 {
		return 1
	}
	return -1
}

func contains(point *D3Point, list []*D3Point) bool {
  for _, x := range list {
    if x == point {
      return true
    }
  }
  return false
}

func ProjectPointsOntoCanvas(model *Model, canvas *Canvas, offsetRow, offsetColumn int, occludedPoints []*D3Point) {
	offsetX := float64(0)
	offsetZ := float64(-10)

	for _, point := range model.Points {
    if occludedPoints != nil && contains(point, occludedPoints) {
      continue
    }

		projection := OrthographicProjection(point, offsetX, offsetZ)
		y := round(projection.Y()*CANVASYSCALE) + offsetRow
		x := round(projection.X()*CANVASXSCALE) + offsetColumn
		canvas.Set(y, x, POINT)
	}
}

func BresenhamProjection(canvas *Canvas, x0, x1, y0, y1 int) {
	if x1 < x0 {
		tmpX := x0
		tmpY := y0
		x0 = x1
		y0 = y1
		x1 = tmpX
		y1 = tmpY
	}

	deltaX := x1 - x0
	deltaY := y1 - y0
	err := 0.0
	y := y0
	errSign := sign(y1 - y0)
	var deltaErr float64

	if deltaX == 0 {
		if y0 >= y1 {
			deltaErr = float64(y0 - y1)
		} else {
			deltaErr = float64(y1 - y0)
		}
	} else {
		deltaErr = math.Abs(float64(deltaY) / float64(deltaX))
	}

	for x := x0; x <= x1; x++ {
		canvas.Set(y, x, EDGE)
		err += deltaErr
		for err >= 0.5 && ((y0 > y1 && y > y1) || (y < y1)) {
			canvas.Set(y, x, EDGE)
			y += errSign
			err -= 1.0
		}
	}
}

func ProjectEdgesOntoCanvas(model *Model, canvas *Canvas, offsetRow, offsetColumn int, occludedPoints []*D3Point) {
	offsetX := float64(0)
	offsetZ := float64(-10)

	for _, edge := range model.Edges {
    if occludedPoints != nil && (contains(edge.P0, occludedPoints) || contains(edge.P1, occludedPoints)) {
      continue
    }
		p0 := OrthographicProjection(edge.P0, offsetX, offsetZ)
		p1 := OrthographicProjection(edge.P1, offsetX, offsetZ)
		x0 := round(p0.X()*CANVASXSCALE) + int(offsetColumn)
		x1 := round(p1.X()*CANVASXSCALE) + int(offsetColumn)
		y0 := round(p0.Y()*CANVASYSCALE) + int(offsetRow)
		y1 := round(p1.Y()*CANVASYSCALE) + int(offsetRow)
		BresenhamProjection(canvas, x0, x1, y0, y1)
	}
}

func initTermbox() chan termbox.Event {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	event_queue := make(chan termbox.Event)
	go func() {
		for {
			event_queue <- termbox.PollEvent()
		}
	}()

	return event_queue
}

type ByZ []*D3Point

func (a ByZ) Len() int           { return len(a) }
func (a ByZ) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByZ) Less(i, j int) bool { return a[i].Y() > a[j].Y() }

func cullPointsOnZ(model *Model, n int) []*D3Point {
  list := make([]*D3Point, len(model.Points))
  top := make([]*D3Point, n)
  copy(list, model.Points)

  sort.Sort(ByZ(list))

  for i := 0; i < n && i < len(list); i++ {
    top = append(top, list[i])
  }

  return top
}

func printCanvasToTermbox(canvas *Canvas) {
	w1 := int(canvas.Width())
	h1 := int(canvas.Height())
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for y := 0; y < h1; y++ {
		for x := 0; x < w1; x++ {
			c := canvas.Get(y, x)
			if c != BLANK {
				termbox.SetCell(x, y, c, termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}
	termbox.Flush()
}

func main() {
	fmt.Print("foo")
	event_queue := initTermbox()

	width := uint(120)
	height := uint(60)
	canvas := NewCanvas(height, width)
	model := MakeIcosahedron()
	model.Scale(10)
	xOffset := int(width) / 2
	yOffset := int(float64(height) / 1.5)

loop:
	for {
		select {
		case ev := <-event_queue:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyCtrlC {
				break loop
			}
		default:
			model.RotateAroundXAxis(0.008)
			model.RotateAroundYAxis(0.03)

			ProjectEdgesOntoCanvas(model, canvas, yOffset, xOffset, nil)
			ProjectPointsOntoCanvas(model, canvas, yOffset, xOffset, nil)
			printCanvasToTermbox(canvas)
			canvas.Clear()
			time.Sleep(50 * time.Millisecond)
		}
	}

	termbox.Close()
}
