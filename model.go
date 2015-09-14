package main

type Edge struct {
	P0 *D3Point
	P1 *D3Point
}

func (this *Edge) Equals(that *Edge) bool {
	return (this.P0 == that.P0 && this.P1 == that.P1) || (this.P0 == that.P1 && this.P1 == that.P0)
}

type Triangle struct {
	A, B, C *D3Point
}

func NewTriangle(a, b, c *D3Point) *Triangle {
	return &Triangle{a, b, c}
}

type Model struct {
	points    []*D3Point
	edges     []*Edge
	Triangles []*Triangle
	valid     bool
}

func NewModel() *Model {
	return &Model{
		nil,
		nil,
		make([]*Triangle, 0),
		false,
	}
}

func (this *Model) invalidatePointsAndEdges() {
	this.valid = false
}

func (this *Model) edgesContain(x *Edge) bool {
	for _, edge := range this.edges {
		if edge != nil && edge.Equals(x) {
			return true
		}
	}
	return false
}

func (this *Model) collectPointsAndEdges() {
	if this.valid {
		return
	}
	this.valid = true

	points := make(map[*D3Point]bool)
	this.edges = make([]*Edge, 0)

	for _, triangle := range this.Triangles {
		points[triangle.A] = true
		points[triangle.B] = true
		points[triangle.C] = true

		triEdges := make([]*Edge, 0)
		triEdges = append(triEdges, &Edge{triangle.A, triangle.B})
		triEdges = append(triEdges, &Edge{triangle.B, triangle.C})
		triEdges = append(triEdges, &Edge{triangle.C, triangle.A})

		for _, edge := range triEdges {
			if !this.edgesContain(edge) {
				this.edges = append(this.edges, edge)
			}
		}
	}

	this.points = make([]*D3Point, 0)
	for k, _ := range points {
		this.points = append(this.points, k)
	}
}

func (this *Model) Points() []*D3Point {
	this.collectPointsAndEdges()
	return this.points
}

func (this *Model) Edges() []*Edge {
	this.collectPointsAndEdges()
	return this.edges
}

func (this *Model) Scale(factor float64) {
	this.collectPointsAndEdges()
	for i := 0; i < len(this.points) && this.points[i] != nil; i++ {
		this.points[i].Scale(factor)
	}
}

func (this *Model) ScaleDim(xf, yf, zf float64) {
	this.collectPointsAndEdges()
	for i := 0; i < len(this.points) && this.points[i] != nil; i++ {
		this.points[i].ScaleDim(xf, yf, zf)
	}
}

func (this *Model) RotateAroundXAxis(angle float64) {
	this.collectPointsAndEdges()
	for i := 0; i < len(this.points) && this.points[i] != nil; i++ {
		this.points[i].RotateAroundXAxis(angle)
	}
}

func (this *Model) RotateAroundYAxis(angle float64) {
	this.collectPointsAndEdges()
	for i := 0; i < len(this.points) && this.points[i] != nil; i++ {
		this.points[i].RotateAroundYAxis(angle)
	}
}
