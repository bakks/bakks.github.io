// https://en.wikipedia.org/wiki/3D_projection
// https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// http://members.chello.at/~easyfilter/bresenham.html
// http://inside.mines.edu/fs_home/gmurray/ArbitraryAxisRotation/
// http://blog.andreaskahler.com/2009/06/creating-icosphere-mesh-in-code.html

package main

//import "fmt"
import "math"
import "time"
import "log"
import "os"
//import "strconv"
import "github.com/nsf/termbox-go"

var POINT rune = '•'
var EDGE rune = '·'
var BLANK rune = ' '
var CANVASXSCALE float64 = 2.2
var CANVASYSCALE float64 = 1.0
var logger *log.Logger

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

func ProjectPointsOntoCanvas(points []*D3Point, canvas *Canvas, offsetRow, offsetColumn int, occludedPoints []*D3Point) {
	offsetX := float64(0)
	offsetZ := float64(-10)

	for _, point := range points {
    if occludedPoints != nil && contains(point, occludedPoints) {
      continue
    }

		projection := OrthographicProjection(point, offsetX, offsetZ)
		y := round(projection.Y()*CANVASYSCALE) + offsetRow
		x := round(projection.X()*CANVASXSCALE) + offsetColumn
		canvas.Set(y, x, POINT)
	}
}

//FILLS := map[int]int[]{
//	0:	 ['|', '|', '|'],
//	45:	 ['/', '/', '/'],
//	90:  ['_', '-', '-'],
//	135: [',', '\\', '`'],
//]

func pointLineDist(x0, y0, deltaX, deltaY, adj, length float64) float64 {
	return math.Abs(deltaY * x0 - deltaX * y0 + adj) / length
}

func pointDist(x0, y0, x1, y1 float64) float64 {
	return math.Sqrt((x1 - x0) * (x1 - x0) + (y1 - y0) * (y1 - y0))
}

func WideLineProjection(canvas *Canvas, x2, x1, y2, y1, width float64) {
	if x2 < x1 {
		tmpX := x2
		tmpY := y2
		x2 = x1
		y2 = y1
		x1 = tmpX
		y1 = tmpY
	}

	maxX := int(canvas.Width())
	maxY := int(canvas.Height())
	deltaX := x2 - x1
	deltaY := y2 - y1

	var m float64
	if math.Abs(x2 - x1) < 0.000001 {
		m = math.NaN()
	} else {
		m = (y2 - y1) / (x2 - x1)
	}
	b := y1 - m * x1

	adj := x2 * y1 - y2 * x1
	length := math.Sqrt((y2 - y1) * (y2 - y1) + (x2 - x1) * (x2 - x1))

	var y0 float64

	for x0 := round(x1 - width - 1); x0 <= round(x2 + width + 1) && x0 <= maxX; x0++ {
		//logger.Printf("(%f,%f) (%f,%f) : %d, %f\n", x1, y1, x2, y2, x0, y0)
		if !math.IsNaN(m) {
			y0 = m * float64(x0) + b
		} else {
			y0 = y1
		}

		for yi := round(y0); yi >= 0 && yi <= maxY; yi++ {  // exit condition
			lineDist := pointLineDist(float64(x0), float64(yi), deltaX, deltaY, adj, length)
			pointDist1 := pointDist(float64(x0), float64(yi), x1, y1)
			pointDist2 := pointDist(float64(x0), float64(yi), x2, y2)

			if lineDist > width || pointDist1 + pointDist2 > length + 2 {
				break
			}

			canvas.Set(yi, x0, EDGE)
		}

		for yi := round(y0) - 1; yi >= 0 && yi <= maxY; yi-- {  // exit condition
			lineDist := pointLineDist(float64(x0), float64(yi), deltaX, deltaY, adj, length)
			pointDist1 := pointDist(float64(x0), float64(yi), x1, y1)
			pointDist2 := pointDist(float64(x0), float64(yi), x2, y2)

			if lineDist > width || (pointDist1 + pointDist2) > length + 2 {
				break
			}

			canvas.Set(yi, x0, EDGE)
		}
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

func ProjectEdgesOntoCanvas(edges []*Edge, canvas *Canvas, offsetRow, offsetColumn int, occludedPoints []*D3Point) {
	offsetX := float64(0)
	offsetZ := float64(-10)

	for _, edge := range edges {
    if occludedPoints != nil && (contains(edge.P0, occludedPoints) || contains(edge.P1, occludedPoints)) {
      continue
    }
		p0 := OrthographicProjection(edge.P0, offsetX, offsetZ)
		p1 := OrthographicProjection(edge.P1, offsetX, offsetZ)
		x0 := p0.X()*CANVASXSCALE + float64(offsetColumn)
		x1 := p1.X()*CANVASXSCALE + float64(offsetColumn)
		y0 := p0.Y()*CANVASYSCALE + float64(offsetRow)
		y1 := p1.Y()*CANVASYSCALE + float64(offsetRow)
		WideLineProjection(canvas, x0, x1, y0, y1, 1.5)
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
	f, err := os.OpenFile("ico.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	logger = log.New(f, "ico: ", log.Lshortfile)
	logger.Print("Starting ico run")

	event_queue := initTermbox()

	width := uint(120)
	height := uint(60)
	canvas := NewCanvas(height, width)
	model := MakeIcosahedron(1)
	model.Scale(30)
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
			model.RotateAroundXAxis(0.02)
			//model.RotateAroundYAxis(0.08)

			points, edges := model.CollectPointsAndEdges(true)
			ProjectEdgesOntoCanvas(edges, canvas, yOffset, xOffset, nil)
			ProjectPointsOntoCanvas(points, canvas, yOffset, xOffset, nil)
			printCanvasToTermbox(canvas)
			canvas.Clear()
			time.Sleep(50 * time.Millisecond)
		}
	}

	termbox.Close()
}
