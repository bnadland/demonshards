package main

import (
	"github.com/nsf/termbox-go"
	"unicode/utf8"
)

type Viking struct {
	Sprite rune
	X      int
	Y      int
}

func NewViking(sprite string, x int, y int) *Viking {
	spriteRune, _ := utf8.DecodeRuneInString(sprite)
	return &Viking{
		Sprite: spriteRune,
		X:      x,
		Y:      y,
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

	for running {
		for _, viking := range vikings {
			termbox.SetCell(viking.X, viking.Y, viking.Sprite, termbox.ColorGreen, termbox.ColorBlack)
		}
		termbox.SetCursor(cursor.X, cursor.Y)
		termbox.Flush()

		event := termbox.PollEvent()
		if event.Type == termbox.EventKey {
			switch event.Key {
			case termbox.KeyEsc:
				running = false
				break
			case termbox.KeyArrowUp:
				cursor.Y -= 1
				break
			case termbox.KeyArrowDown:
				cursor.Y += 1
				break
			case termbox.KeyArrowLeft:
				cursor.X -= 1
				break
			case termbox.KeyArrowRight:
				cursor.X += 1
				break
			}
		}
	}
}
