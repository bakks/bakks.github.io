package main

import "github.com/nsf/termbox-go"
import "os"
import "log"
import "time"

var logger *log.Logger

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

	ico := NewIco(uint(120), uint(60), 120 / 2, int(60.0 / 1.5), 30.0)

loop:
	for {
		select {
		case ev := <-event_queue:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyCtrlC {
				break loop
			}
		default:
			ico.Model.RotateAroundXAxis(0.02)
			//model.RotateAroundYAxis(0.08)

			_, edges := ico.Model.CollectPointsAndEdges(true)
			ProjectEdgesOntoCanvas(edges, ico.Canvas, ico.YOffset, ico.XOffset, nil)
			//ProjectPointsOntoCanvas(points, ico.Canvas, ico.YOffset, ico.XOffset, nil)
			printCanvasToTermbox(ico.Canvas)
			ico.Canvas.Clear()
			time.Sleep(50 * time.Millisecond)
		}
	}

	termbox.Close()
}
