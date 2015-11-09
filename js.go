package js


func main() {
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
