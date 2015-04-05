package main

import (
	"container/list"
	"fmt"
	"github.com/nsf/termbox-go"
	"io/ioutil"
	"strings"
	"time"
)

type Tile struct {
	Symbol     rune
	IsPassable bool
	IsBlueMove bool
	IsGoldMove bool
}

type Point struct {
	X int
	Y int
}

type MissionMap struct {
	Spawn *Point
	Tiles [][]*Tile
	MaxX  int
	MaxY  int
}

func (m *MissionMap) Draw() {
	for y, row := range m.Tiles {
		for x, tile := range row {
			termbox.SetCell(x, y, tile.Symbol, termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}

func (m *MissionMap) Neighbours(p Point) []Point {
	var neighbours []Point
	for y := p.Y - 1; y < p.Y+2; y++ {
		for x := p.X - 1; x < p.X+2; x++ {
			if x > -1 && y > -1 {
				neighbours = append(neighbours, Point{X: x, Y: y})
			}
		}
	}
	return neighbours
}

func (m *MissionMap) FindPath(start *Point, target *Point) bool {
	termbox.HideCursor()

	frontier := list.New()
	frontier.PushBack(Point{X: start.X, Y: start.Y})

	visited := make(map[int]map[int]*Point)
	for y := 0; y < m.MaxY; y++ {
		visited[y] = make(map[int]*Point)
	}
	visited[start.Y][start.X] = nil

	for frontier.Len() > 0 {
		head := frontier.Front()
		frontier.Remove(head)
		current := head.Value.(Point)

		if current.X == target.X && current.Y == target.Y {
			path := []Point{Point{X: target.X, Y: target.Y}}
			m.Draw()
			for _, p := range path {
				termbox.SetCell(p.X, p.Y, '.', termbox.ColorBlue, termbox.ColorBlue)
			}
			termbox.Flush()
			time.Sleep(500 * time.Millisecond)
			// get path
			return true
		}

		m.Draw()

		for _, point := range m.Neighbours(current) {
			if m.Tiles[point.Y][point.X].IsPassable {
				if visited[point.Y][point.X] == nil {
					frontier.PushBack(point)
					visited[point.Y][point.X] = &current
				}
			}
		}

		for y := 0; y < m.MaxY; y++ {
			for x := 0; x < m.MaxX; x++ {
				if visited[y][x] != nil {
					termbox.SetCell(x, y, '.', termbox.ColorYellow|termbox.AttrBold, termbox.ColorBlack)
				}
			}
		}

		termbox.SetCell(current.X, current.Y, '.', termbox.ColorBlue|termbox.AttrBold, termbox.ColorBlack)
		termbox.Flush()
		time.Sleep(300 * time.Millisecond)
	}
	return false
}

func NewMap(mapName string) *MissionMap {
	file, err := ioutil.ReadFile(fmt.Sprintf("%s.map", mapName))
	if err != nil {
		panic(err)
	}

	missionMap := &MissionMap{}

	maxX := 0
	maxY := 0

	for y, line := range strings.Split(string(file), "\n") {
		maxY = y
		var row []*Tile
		for x, symbol := range line {
			maxX = x

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

	missionMap.MaxX = maxX
	missionMap.MaxY = maxY

	return missionMap
}
