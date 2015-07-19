package main

func MakeCube() *Model {
	x := NewModel()

	x.Points = append(x.Points, &D3Point{-1, 1, -1})
	x.Points = append(x.Points, &D3Point{1, 1, -1})
	x.Points = append(x.Points, &D3Point{-1, -1, -1})
	x.Points = append(x.Points, &D3Point{1, -1, -1})
	x.Points = append(x.Points, &D3Point{-1, 1, 1})
	x.Points = append(x.Points, &D3Point{1, 1, 1})
	x.Points = append(x.Points, &D3Point{-1, -1, 1})
	x.Points = append(x.Points, &D3Point{1, -1, 1})

	x.Edges = append(x.Edges, Edge{x.Points[0], x.Points[1]})
	x.Edges = append(x.Edges, Edge{x.Points[1], x.Points[3]})
	x.Edges = append(x.Edges, Edge{x.Points[3], x.Points[2]})
	x.Edges = append(x.Edges, Edge{x.Points[2], x.Points[0]})
	x.Edges = append(x.Edges, Edge{x.Points[4], x.Points[5]})
	x.Edges = append(x.Edges, Edge{x.Points[5], x.Points[7]})
	x.Edges = append(x.Edges, Edge{x.Points[7], x.Points[6]})
	x.Edges = append(x.Edges, Edge{x.Points[6], x.Points[4]})
	x.Edges = append(x.Edges, Edge{x.Points[0], x.Points[4]})
	x.Edges = append(x.Edges, Edge{x.Points[1], x.Points[5]})
	x.Edges = append(x.Edges, Edge{x.Points[3], x.Points[7]})
	x.Edges = append(x.Edges, Edge{x.Points[2], x.Points[6]})

	return x
}
