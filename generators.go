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

func MakeIcosahedron() *Model {
  x := NewModel()
	points := make([]*D3Point, 0)

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

	x.Triangles = append(x.Triangles, NewTriangle(points[ 0], points[11], points[ 5]))
	x.Triangles = append(x.Triangles, NewTriangle(points[ 0], points[ 5], points[ 1]))
	x.Triangles = append(x.Triangles, NewTriangle(points[ 0], points[ 1], points[ 7]))
	x.Triangles = append(x.Triangles, NewTriangle(points[ 0], points[ 7], points[10]))
	x.Triangles = append(x.Triangles, NewTriangle(points[ 0], points[10], points[11]))
	x.Triangles = append(x.Triangles, NewTriangle(points[ 1], points[ 5], points[ 9]))
	x.Triangles = append(x.Triangles, NewTriangle(points[ 5], points[11], points[ 4]))
	x.Triangles = append(x.Triangles, NewTriangle(points[11], points[10], points[ 2]))
	x.Triangles = append(x.Triangles, NewTriangle(points[10], points[ 7], points[ 6]))
	x.Triangles = append(x.Triangles, NewTriangle(points[ 7], points[ 1], points[ 8]))
	x.Triangles = append(x.Triangles, NewTriangle(points[ 3], points[ 9], points[ 4]))
	x.Triangles = append(x.Triangles, NewTriangle(points[ 3], points[ 4], points[ 2]))
	x.Triangles = append(x.Triangles, NewTriangle(points[ 3], points[ 2], points[ 6]))
	x.Triangles = append(x.Triangles, NewTriangle(points[ 3], points[ 6], points[ 8]))
	x.Triangles = append(x.Triangles, NewTriangle(points[ 3], points[ 8], points[ 9]))
	x.Triangles = append(x.Triangles, NewTriangle(points[ 4], points[ 9], points[ 5]))
	x.Triangles = append(x.Triangles, NewTriangle(points[ 2], points[ 4], points[11]))
	x.Triangles = append(x.Triangles, NewTriangle(points[ 6], points[ 2], points[10]))
	x.Triangles = append(x.Triangles, NewTriangle(points[ 8], points[ 6], points[ 7]))
	x.Triangles = append(x.Triangles, NewTriangle(points[ 9], points[ 8], points[ 1]))

  return x
}

