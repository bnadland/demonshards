package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"io/ioutil"
	"strings"
)

type Tile struct {
	Symbol     rune
	IsPassable bool
}

type Point struct {
	X int
	Y int
}

type MissionMap struct {
	Spawn *Point
	Tiles [][]*Tile
}

func (m *MissionMap) Draw() {
	for y, row := range m.Tiles {
		for x, tile := range row {
			termbox.SetCell(x, y, tile.Symbol, termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}

func (m *MissionMap) IsPassable(position *Point) bool {
	return m.Tiles[position.Y][position.X].IsPassable
}

func NewMap(mapName string) *MissionMap {
	file, err := ioutil.ReadFile(fmt.Sprintf("%s.map", mapName))
	if err != nil {
		panic(err)
	}

	missionMap := &MissionMap{}

	for y, line := range strings.Split(string(file), "\n") {
		var row []*Tile
		for x, symbol := range line {

			tile := &Tile{
				Symbol: symbol,
			}

			switch symbol {
			case Rune("."):
				tile.IsPassable = true
				break
			case Rune("@"):
				tile.Symbol = '.'
				tile.IsPassable = true
				missionMap.Spawn = &Point{X: x, Y: y}
			default:
				tile.IsPassable = false
			}

			row = append(row, tile)
		}
		missionMap.Tiles = append(missionMap.Tiles, row)
	}

	return missionMap
}
