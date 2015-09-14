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
	triangles []*Triangle
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

func (this *Model) AddTriangle(triangle *Triangle) {
	this.triangles = append(this.triangles, triangle)
	this.valid = false
}

func (this *Model) SetTriangles(triangles []*Triangle) {
	this.triangles = triangles
	this.valid = false
}

func (this *Model) invalidatePointsAndEdges() {
	this.valid = false
}

func edgesContain(edges []*Edge, x *Edge) bool {
	for _, edge := range edges {
		if edge != nil && edge.Equals(x) {
			return true
		}
	}
	return false
}

func (this *Model) facingTowardsCamera(t *Triangle) bool {
	a := SubMatrices(t.B.ToMatrix(), t.A.ToMatrix())
	b := SubMatrices(t.C.ToMatrix(), t.A.ToMatrix())
	c := CrossProductMatrices(a, b)
	//fmt.Print(c)
	return c.Get(1, 0) < 0
}

func (this *Model) CollectPointsAndEdges(masking bool) ([]*D3Point, []*Edge) {
	pointsMap := make(map[*D3Point]bool)
	edges := make([]*Edge, 0)

	for _, triangle := range this.triangles {
		if masking && !this.facingTowardsCamera(triangle) {
			continue
		}

		pointsMap[triangle.A] = true
		pointsMap[triangle.B] = true
		pointsMap[triangle.C] = true

		triEdges := make([]*Edge, 0)
		triEdges = append(triEdges, &Edge{triangle.A, triangle.B})
		triEdges = append(triEdges, &Edge{triangle.B, triangle.C})
		triEdges = append(triEdges, &Edge{triangle.C, triangle.A})

		for _, edge := range triEdges {
			if !edgesContain(edges, edge) {
				edges = append(edges, edge)
			}
		}
	}

	points := make([]*D3Point, 0)
	for k, _ := range pointsMap {
		points = append(points, k)
	}

	return points, edges
}

func (this *Model) collectModelPointsAndEdges() {
	if this.valid {
		return
	}
	this.valid = true

	this.points, this.edges = this.CollectPointsAndEdges(false)
}

func (this *Model) Points() []*D3Point {
	this.collectModelPointsAndEdges()
	return this.points
}

func (this *Model) Edges() []*Edge {
	this.collectModelPointsAndEdges()
	return this.edges
}

func (this *Model) Scale(factor float64) {
	this.collectModelPointsAndEdges()
	for i := 0; i < len(this.points) && this.points[i] != nil; i++ {
		this.points[i].Scale(factor)
	}
}

func (this *Model) ScaleDim(xf, yf, zf float64) {
	this.collectModelPointsAndEdges()
	for i := 0; i < len(this.points) && this.points[i] != nil; i++ {
		this.points[i].ScaleDim(xf, yf, zf)
	}
}

func (this *Model) RotateAroundXAxis(angle float64) {
	this.collectModelPointsAndEdges()
	for i := 0; i < len(this.points) && this.points[i] != nil; i++ {
		this.points[i].RotateAroundXAxis(angle)
	}
}

func (this *Model) RotateAroundYAxis(angle float64) {
	this.collectModelPointsAndEdges()
	for i := 0; i < len(this.points) && this.points[i] != nil; i++ {
		this.points[i].RotateAroundYAxis(angle)
	}
}
