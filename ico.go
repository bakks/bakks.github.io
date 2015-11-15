// https://en.wikipedia.org/wiki/3D_projection
// https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// http://members.chello.at/~easyfilter/bresenham.html
// http://inside.mines.edu/fs_home/gmurray/ArbitraryAxisRotation/
// http://blog.andreaskahler.com/2009/06/creating-icosphere-mesh-in-code.html
// http://www.utf8-chartable.de/unicode-utf8-table.pl?start=9472

package main

import "math"

var POINT rune = '░'
var EDGE rune = '⣿'
var BLANK rune = ' '
var CANVASXSCALE float64 = 2.4
var CANVASYSCALE float64 = 1.0

type Ico struct {
	Width uint
	Height uint
	XOffset int
	YOffset int
	Canvas *Canvas
	Model *Model
}

func NewIco(width, height uint, xOffset, yOffset int, scale float64) *Ico {
	ico := &Ico{
		width, height, xOffset, yOffset,
		NewCanvas(height, width),
		MakeIcosahedron(1),
	}
	ico.Model.Scale(scale)
	return ico
}

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

var FILLS map[int]rune = map[int]rune{
	0:	 '⣿',
	45:	 '⣼',
	90:  '⠶',
	135: '⠦',
}

func pointLineDist(x0, y0, deltaX, deltaY, adj, length float64) float64 {
	return math.Abs(deltaY * x0 - deltaX * y0 + adj) / length
}

func pointLineDistWithin(x0, y0, deltaX, deltaY, adj, length, dist float64) int {
	if pointLineDist(x0, y0, deltaX, deltaY, adj, length) < dist {
		return 1
	}
	return 0
}

func pointDist(x0, y0, x1, y1 float64) float64 {
	return math.Sqrt((x1 - x0) * (x1 - x0) + (y1 - y0) * (y1 - y0))
}

var BRAILLE rune = '⠀'

func fillPoints(x0, yStart, inc int, x1, y1, x2, y2, deltaX, deltaY, adj, length float64, canvas *Canvas, maxY int, width float64) {
	x := float64(x0)
	for yi := yStart; yi >= 0 && yi <= maxY; yi += inc {
		y := float64(yi)
		//y := float64(yStart) + float64(yi - yStart) * 2.2
		//lineDist := pointLineDist(float64(x0), float64(yi), deltaX, deltaY, adj, length)
		pointDist1 := pointDist(x, float64(yi), x1, y1)
		pointDist2 := pointDist(x, float64(yi), x2, y2)

		filledPoint := canvas.Get(yi, x0)
		runePoint := 0
		if int(filledPoint) > int(BRAILLE) && int(filledPoint) <= int('⣿') {
			//logger.Printf("::: %d\n", filledPoint)
			runePoint = int(filledPoint) - int(BRAILLE)
		}
		runePoint |= (pointLineDistWithin(x - 0.1, y - 0.2, deltaX, deltaY, adj, length, width) << 0)
		runePoint |= (pointLineDistWithin(x - 0.1, y - 0.1, deltaX, deltaY, adj, length, width) << 1)
		runePoint |= (pointLineDistWithin(x - 0.1, y + 0.1, deltaX, deltaY, adj, length, width) << 2)
		runePoint |= (pointLineDistWithin(x + 0.1, y - 0.2, deltaX, deltaY, adj, length, width) << 3)
		runePoint |= (pointLineDistWithin(x + 0.1, y - 0.1, deltaX, deltaY, adj, length, width) << 4)
		runePoint |= (pointLineDistWithin(x + 0.1, y + 0.1, deltaX, deltaY, adj, length, width) << 5)
		runePoint |= (pointLineDistWithin(x - 0.1, y + 0.2, deltaX, deltaY, adj, length, width) << 6)
		runePoint |= (pointLineDistWithin(x + 0.1, y + 0.2, deltaX, deltaY, adj, length, width) << 7)

		if runePoint == 0 || pointDist1 + pointDist2 > length + 2 {
			return
		}

		edge := rune(int(BRAILLE) + runePoint)
		canvas.Set(yi, x0, edge)
	}
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

		fillPoints(x0, round(y0), 1, x1, y1, x2, y2, deltaX, deltaY, adj, length, canvas, maxY, width)
		fillPoints(x0, round(y0), -1, x1, y1, x2, y2, deltaX, deltaY, adj, length, canvas, maxY, width)
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
		WideLineProjection(canvas, x0, x1, y0, y1, 0.5)
	}
}

type ByZ []*D3Point

func (a ByZ) Len() int           { return len(a) }
func (a ByZ) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByZ) Less(i, j int) bool { return a[i].Y() > a[j].Y() }

