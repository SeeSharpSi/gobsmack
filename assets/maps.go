package assets

import (
	"fmt"
	"math/rand/v2"
)

type Ship struct {
	Width  int
	Height int

	// Rooms axis is X/Y (width/height)
	Rooms [][]Room
}

func (s *Ship) NewShip() {
	gen := rand.IntN(3)
	w := 5 + gen
	h := 5 - gen
	s.Width = w
	s.Height = h
	roomMatrix := make([][]Room, h)
	for i := range h {
		roomMatrix[i] = make([]Room, w)
		for j := range w {
			edges := make(map[Direction]bool)
			if i == 0 {
				edges["north"] = true
			}
			if i == h-1 {
				edges["south"] = true
			}
			if j == 0 {
				edges["east"] = true
			}
			if j == w-1 {
				edges["west"] = true
			}
			r := Room{
				AxisX: j,
				AxisY: i,
			}
			// fmt.Printf("\ni: %d, j: %d  ", i, j)
			// fmt.Println(edges)
			r.NewRoom(edges)
			roomMatrix[i][j] = r
		}
	}
	s.Rooms = roomMatrix
}

func (s Ship) RenderHTML() string {
	var str string
	str = fmt.Sprintf("<div style='display: grid; grid-template-columns: %d auto; grid-template-rows: %d auto;'>", s.Width, s.Height)
	for i := range s.Height {
		for j := range s.Width {
			currentRoom := s.Rooms[i][j]
			str += fmt.Sprintf("<div style='grid-row: %d; grid-column: %d; border-top: 2px %s; border-bottom: 2px %s; border-left: 2px %s; border-right: 2px %s;'>", i+1, j+1, currentRoom.Walls["north"].Render, currentRoom.Walls["south"].Render, currentRoom.Walls["east"].Render, currentRoom.Walls["west"].Render)
			str += currentRoom.Render
			str += "</div>"
		}
	}
	str += "</div>"
	return str[:len(str)]
}

type Room struct {
	AxisX, AxisY int
	Type         string
	Render       string
	Contains     []Item
	Walls        map[Direction]Wall
}

// Direction of either north, south, east, or west
type Direction string

type Wall struct {
	Type     string
	Health   int
	ShipEdge bool
	Render   string
}

func (w *Wall) NewWall() {
	keys := make([]string, 0, len(WallTypes)-1)
	for k := range WallTypes {
		if k != "edge" {
			keys = append(keys, k)
		}
	}
	randnum := rand.IntN(len(keys))
	*w = WallTypes[keys[randnum]]
}

var WallTypes = map[string]Wall{
	"edge": {
		Type:   "edge",
		Health: 100,
		Render: "solid red",
	},
	"empty": {
		Type:   "empty",
		Health: 0,
		Render: "solid black",
	},
	"solid": {
		Type:   "solid",
		Health: 10,
		Render: "solid blue",
	},
}

// whether or not a direction is the edge of the ship
func (r *Room) NewRoom(edges map[Direction]bool) {
	keys := make([]string, 0, len(RoomTypes))
	for k := range RoomTypes {
		keys = append(keys, k)
	}
	randnum := rand.IntN(len(keys))
	*r = RoomTypes[keys[randnum]]

	r.Walls = make(map[Direction]Wall)
	wall := Wall{}
	wall.NewWall()
	r.Walls["north"] = wall
	wall.NewWall()
	r.Walls["south"] = wall
	wall.NewWall()
	r.Walls["east"] = wall
	wall.NewWall()
	r.Walls["west"] = wall
	for k, v := range edges {
		if v {
			r.Walls[Direction(k)] = WallTypes["edge"]
		}
	}
}

var RoomTypes = map[string]Room{
	"empty": {
		Type:     "empty",
		Render:   "E",
		Contains: nil,
	},
	"gun": {
		Type:     "gun",
		Render:   "X",
		Contains: []Item{&Gun{}},
	},
}
