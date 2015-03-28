package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"math/rand"
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
	return &Viking{
		Sprite:     tile(sprite),
		Color:      termbox.ColorBlack,
		Background: termbox.ColorGreen,
		X:          x,
		Y:          y,
	}
}

func render(text string, x int, y int) {
	for i, textRune := range text {
		termbox.SetCell(x+i, y, textRune, termbox.ColorWhite, termbox.ColorBlack)
	}
}

func tile(ch string) rune {
	tileRune, _ := utf8.DecodeRuneInString(ch)
	return tileRune
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	rows := 24
	cols := 80

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

	running := true

	enemies := []string{"z", "m", "s"}

	selectedViking := 1

	for running {

		render(fmt.Sprintf("[%s] Viking", strconv.Itoa(selectedViking)), 0, 0)

		// floor
		func() {
			for x := 0; x < cols; x++ {
				for y := 1; y < rows; y++ {
					termbox.SetCell(x, y, tile("."), termbox.ColorWhite, termbox.ColorBlack)
				}
			}
		}()

		// walls
		func() {
			for i := 0; i < rand.Intn(rows*cols); i++ {
				termbox.SetCell(
					rand.Intn(cols),
					rand.Intn(rows-1)+1,
					tile("#"),
					termbox.ColorWhite,
					termbox.ColorBlack)
			}

		}()

		// walls
		func() {
			for i := 0; i < rand.Intn(rows*cols); i++ {
				termbox.SetCell(
					rand.Intn(cols),
					rand.Intn(rows-1)+1,
					tile("t"),
					termbox.ColorWhite,
					termbox.ColorBlack)
			}

		}()

		// enemies
		func() {
			for i := 0; i < rand.Intn(9); i++ {
				termbox.SetCell(
					rand.Intn(cols),
					rand.Intn(rows-1)+1,
					tile(enemies[rand.Intn(len(enemies))]),
					termbox.ColorWhite,
					termbox.ColorRed)
			}
		}()

		// vikings
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
