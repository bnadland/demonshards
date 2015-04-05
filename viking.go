package main

import (
	"github.com/nsf/termbox-go"
)

type Viking struct {
	Sprite   rune
	Position *Point
}

func (v *Viking) Draw() {
	termbox.SetCell(v.Position.X, v.Position.Y, v.Sprite, termbox.ColorDefault, termbox.ColorDefault)
}

func NewViking(sprite rune, x int, y int) *Viking {
	return &Viking{
		Sprite:   sprite,
		Position: &Point{X: x, Y: y},
	}
}
