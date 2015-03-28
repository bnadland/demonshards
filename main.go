package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"strconv"
	"unicode/utf8"
)

type Viking struct {
	Sprite     rune
	Color      termbox.Attribute
	Background termbox.Attribute
	X          int
	Y          int
}

func NewViking(sprite string, x int, y int) *Viking {
	spriteRune, _ := utf8.DecodeRuneInString(sprite)
	return &Viking{
		Sprite:     spriteRune,
		Color:      termbox.ColorGreen,
		Background: termbox.ColorBlack,
		X:          x,
		Y:          y,
	}
}

func render(text string, x int, y int) {
	for i, textRune := range text {
		termbox.SetCell(x+i, y, textRune, termbox.ColorWhite, termbox.ColorBlack)
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	vikings := []*Viking{
		NewViking("1", 5, 7),
		NewViking("2", 7, 7),
		NewViking("3", 9, 7),
	}

	cursor := struct {
		X int
		Y int
	}{
		X: vikings[0].X,
		Y: vikings[0].Y,
	}

	termbox.SetCursor(cursor.X, cursor.Y)

	running := true

	tiles := make(map[string]rune)

	for _, tile := range []string{".", "#"} {
		tileRune, _ := utf8.DecodeRuneInString(tile)
		tiles[tile] = tileRune
	}

	selectedViking := 1

	for running {

		render(fmt.Sprintf("[%s] Viking", strconv.Itoa(selectedViking)), 0, 0)

		for x := 0; x < 80; x++ {
			for y := 1; y < 20; y++ {
				termbox.SetCell(x, y, tiles["."], termbox.ColorWhite, termbox.ColorBlack)
			}
		}

		for _, viking := range vikings {
			termbox.SetCell(viking.X, viking.Y, viking.Sprite, viking.Color, viking.Background)
		}

		termbox.SetCursor(cursor.X, cursor.Y)

		termbox.Flush()

		event := termbox.PollEvent()
		if event.Type == termbox.EventKey {
			if event.Key == termbox.KeyEsc || event.Ch == 'q' {
				running = false
			}
			if event.Key == termbox.KeyArrowUp || event.Ch == 'k' {
				cursor.Y -= 1
			}
			if event.Key == termbox.KeyArrowDown || event.Ch == 'j' {
				cursor.Y += 1
			}
			if event.Key == termbox.KeyArrowLeft || event.Ch == 'h' {
				cursor.X -= 1
			}
			if event.Key == termbox.KeyArrowRight || event.Ch == 'l' {
				cursor.X += 1
			}

			if event.Ch == '1' {
				selectedViking = 1
				cursor.X = vikings[0].X
				cursor.Y = vikings[0].Y
			}
			if event.Ch == '2' {
				selectedViking = 2
				cursor.X = vikings[1].X
				cursor.Y = vikings[1].Y
			}
			if event.Ch == '3' {
				selectedViking = 3
				cursor.X = vikings[2].X
				cursor.Y = vikings[2].Y
			}
		}
	}
}
