package main

import (
	"github.com/nsf/termbox-go"
	"unicode/utf8"
)

func Render(text string, x int, y int) {
	for i, textRune := range text {
		termbox.SetCell(x+i, y, textRune, termbox.ColorWhite, termbox.ColorBlack)
	}
}

func Rune(ch string) rune {
	tileRune, _ := utf8.DecodeRuneInString(ch)
	return tileRune
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	running := true

	missionMap := NewMap("001")

	viking := NewViking('@', missionMap.Spawn.X, missionMap.Spawn.Y)
	cursor := &Point{X: viking.Position.X, Y: viking.Position.Y}

	for running {

		missionMap.Draw()
		viking.Draw()
		termbox.SetCursor(cursor.X, cursor.Y)

		termbox.Flush()

		event := termbox.PollEvent()
		if event.Type == termbox.EventKey {
			if event.Key == termbox.KeyEsc || event.Ch == 'q' {
				running = false
			}

			if event.Ch == 'j' {
				cursor.Y += 1
			}
			if event.Ch == 'k' {
				cursor.Y -= 1
			}
			if event.Ch == 'h' {
				cursor.X -= 1
			}
			if event.Ch == 'l' {
				cursor.X += 1
			}

			if event.Ch == 'x' {
				if missionMap.FindPath(viking.Position, cursor) {
					viking.Position.X = cursor.X
					viking.Position.Y = cursor.Y
				}
			}
		}
	}
}
