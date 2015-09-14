package main

import "math"

//func MakeCube() *Model {
//	x := NewModel()
//
//	x.Points = append(x.Points, &D3Point{-1,  1, -1})
//	x.Points = append(x.Points, &D3Point{ 1,  1, -1})
//	x.Points = append(x.Points, &D3Point{-1, -1, -1})
//	x.Points = append(x.Points, &D3Point{ 1, -1, -1})
//	x.Points = append(x.Points, &D3Point{-1,  1,  1})
//	x.Points = append(x.Points, &D3Point{ 1,  1,  1})
//	x.Points = append(x.Points, &D3Point{-1, -1,  1})
//	x.Points = append(x.Points, &D3Point{ 1, -1,  1})
//
//	x.Edges = append(x.Edges, Edge{x.Points[0], x.Points[1]})
//	x.Edges = append(x.Edges, Edge{x.Points[1], x.Points[3]})
//	x.Edges = append(x.Edges, Edge{x.Points[3], x.Points[2]})
//	x.Edges = append(x.Edges, Edge{x.Points[2], x.Points[0]})
//	x.Edges = append(x.Edges, Edge{x.Points[4], x.Points[5]})
//	x.Edges = append(x.Edges, Edge{x.Points[5], x.Points[7]})
//	x.Edges = append(x.Edges, Edge{x.Points[7], x.Points[6]})
//	x.Edges = append(x.Edges, Edge{x.Points[6], x.Points[4]})
//	x.Edges = append(x.Edges, Edge{x.Points[0], x.Points[4]})
//	x.Edges = append(x.Edges, Edge{x.Points[1], x.Points[5]})
//	x.Edges = append(x.Edges, Edge{x.Points[3], x.Points[7]})
//	x.Edges = append(x.Edges, Edge{x.Points[2], x.Points[6]})
//
//	return x
//}

func MakeIcosahedron(refines int) *Model {
  x := NewModel()
	points := make([]*D3Point, 0)
	triangles := make([]*Triangle, 0)

  t := (1.0 + math.Sqrt(5.0)) / 2.0;

  points = append(points, &D3Point{-1,  t,  0})
  points = append(points, &D3Point{ 1,  t,  0})
  points = append(points, &D3Point{-1, -t,  0})
  points = append(points, &D3Point{ 1, -t,  0})
  points = append(points, &D3Point{ 0, -1,  t})
  points = append(points, &D3Point{ 0,  1,  t})
  points = append(points, &D3Point{ 0, -1, -t})
  points = append(points, &D3Point{ 0,  1, -t})
  points = append(points, &D3Point{ t,  0, -1})
  points = append(points, &D3Point{ t,  0,  1})
  points = append(points, &D3Point{-t,  0, -1})
  points = append(points, &D3Point{-t,  0,  1})

	triangles = append(triangles, NewTriangle(points[ 0], points[11], points[ 5]))
	triangles = append(triangles, NewTriangle(points[ 0], points[ 5], points[ 1]))
	triangles = append(triangles, NewTriangle(points[ 0], points[ 1], points[ 7]))
	triangles = append(triangles, NewTriangle(points[ 0], points[ 7], points[10]))
	triangles = append(triangles, NewTriangle(points[ 0], points[10], points[11]))
	triangles = append(triangles, NewTriangle(points[ 1], points[ 5], points[ 9]))
	triangles = append(triangles, NewTriangle(points[ 5], points[11], points[ 4]))
	triangles = append(triangles, NewTriangle(points[11], points[10], points[ 2]))
	triangles = append(triangles, NewTriangle(points[10], points[ 7], points[ 6]))
	triangles = append(triangles, NewTriangle(points[ 7], points[ 1], points[ 8]))
	triangles = append(triangles, NewTriangle(points[ 3], points[ 9], points[ 4]))
	triangles = append(triangles, NewTriangle(points[ 3], points[ 4], points[ 2]))
	triangles = append(triangles, NewTriangle(points[ 3], points[ 2], points[ 6]))
	triangles = append(triangles, NewTriangle(points[ 3], points[ 6], points[ 8]))
	triangles = append(triangles, NewTriangle(points[ 3], points[ 8], points[ 9]))
	triangles = append(triangles, NewTriangle(points[ 4], points[ 9], points[ 5]))
	triangles = append(triangles, NewTriangle(points[ 2], points[ 4], points[11]))
	triangles = append(triangles, NewTriangle(points[ 6], points[ 2], points[10]))
	triangles = append(triangles, NewTriangle(points[ 8], points[ 6], points[ 7]))
	triangles = append(triangles, NewTriangle(points[ 9], points[ 8], points[ 1]))

	for _, point := range points {
		point.NormalizeDistanceToOrigin()
	}

	for i := 0; i < refines; i++ {
		newTriangles := make([]*Triangle, 0)

		for _, t := range triangles {
			a := t.A.GetMiddlePoint(t.B)
			b := t.B.GetMiddlePoint(t.C)
			c := t.C.GetMiddlePoint(t.A)
			a.NormalizeDistanceToOrigin()
			b.NormalizeDistanceToOrigin()
			c.NormalizeDistanceToOrigin()

			newTriangles = append(newTriangles, &Triangle{t.A, a, c})
			newTriangles = append(newTriangles, &Triangle{t.B, b, a})
			newTriangles = append(newTriangles, &Triangle{t.C, c, b})
			newTriangles = append(newTriangles, &Triangle{a, b, c})
		}

		triangles = newTriangles
	}

	x.SetTriangles(triangles)

  return x
}

