package main

import (
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

func (p1 *Point) equal(p2 *Point) bool {
	if p1.X == p2.X && p1.Y == p2.Y {
		return true
	}
	return false
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
	frontier := make(chan Point, m.MaxY*m.MaxX)
	frontier <- Point{X: start.X, Y: start.Y}

	visited := make(map[int]map[int]bool)
	for y := 0; y < m.MaxY; y++ {
		visited[y] = make(map[int]bool)
		for x := 0; x < m.MaxX; x++ {
			visited[y][x] = false
		}
	}
	visited[start.Y][start.X] = true

	for {
		current := <-frontier

		m.Draw()

		for _, point := range m.Neighbours(current) {
			if m.Tiles[point.Y][point.X].IsPassable {
				if visited[point.Y][point.X] == false {
					frontier <- point
					visited[point.Y][point.X] = true
				}
			}
		}

		for y := 0; y < m.MaxY; y++ {
			for x := 0; x < m.MaxX; x++ {
				if visited[y][x] == true {
					termbox.SetCell(x, y, '.', termbox.ColorYellow, termbox.ColorYellow)
				}
			}
		}

		termbox.SetCell(current.X, current.Y, '.', termbox.ColorBlue, termbox.ColorBlue)
		termbox.Flush()
		time.Sleep(500 * time.Millisecond)

		if current.X == target.X && current.Y == target.Y {
			return true
		}
	}
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
