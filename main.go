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

	viking := NewViking("@", missionMap.Spawn.X, missionMap.Spawn.Y)

	for running {

		missionMap.Draw()
		viking.Draw()

		termbox.Flush()

		event := termbox.PollEvent()
		if event.Type == termbox.EventKey {
			if event.Key == termbox.KeyEsc || event.Ch == 'q' {
				running = false
			}

			position := &Point{
				X: viking.Position.X,
				Y: viking.Position.Y,
			}
			if event.Ch == 'j' {
				position.Y += 1
			}
			if event.Ch == 'k' {
				position.Y -= 1
			}
			if event.Ch == 'h' {
				position.X -= 1
			}
			if event.Ch == 'l' {
				position.X += 1
			}
			if missionMap.IsPassable(position) == true {
				viking.Position = position
			}
		}
	}
}
