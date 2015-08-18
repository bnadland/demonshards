package main

import (
	"github.com/nickdavies/go-astar/astar"
	"github.com/nsf/termbox-go"
	"math/rand"
	"os"
	"time"
)

func printString(col, row int, text string) {
	for i, ch := range text {
		termbox.SetCell(col+i, row, ch, termbox.ColorDefault, termbox.ColorDefault)
	}
}

func showIntro() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	printString(10, 5, "After a long demon hunt, you finally arrive back home at the Isle of Hunters")
	printString(10, 6, "to rest for the winter...")
	termbox.Flush()
	termbox.PollEvent()
}

func showGrid() {
	cols, rows := termbox.Size()
	a := astar.NewAStar(rows, cols)
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	source := astar.Point{Col: r.Intn(cols), Row: r.Intn(rows)}
	target := astar.Point{Col: r.Intn(cols), Row: r.Intn(rows)}

	for i := 0; i < 400; i++ {
		wall := astar.Point{Col: r.Intn(cols), Row: r.Intn(rows)}
		a.FillTile(wall, -1)
		termbox.SetCell(wall.Col, wall.Row, '#', termbox.ColorDefault, termbox.ColorDefault)
	}

	path := a.FindPath(astar.NewPointToPoint(), []astar.Point{source}, []astar.Point{target})

	for path != nil {
		path = path.Parent
		termbox.SetCell(path.Col, path.Row, '.', termbox.ColorDefault, termbox.ColorDefault)
		path = nil //path.Parent
	}

	termbox.SetCell(source.Col, source.Row, 's', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(target.Col, target.Row, 't', termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
	termbox.PollEvent()
}

func main() {
	err = termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	showIntro()
	showGrid()
}
