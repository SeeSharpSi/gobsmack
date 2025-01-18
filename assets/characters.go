package assets

import "math/rand/v2"

type Player struct {
	// North South East West
	Name       string
	Facing     string
	Alive      bool
	TurnOver   bool
	PosX, PosY int
	Type       string
	HP         int
	Strength   int
	Perks      []Perk
	Luck       int
	// x, y T if explored, F if not
	Explored [][]bool
	Items    []Item
}

func (p *Player) Init(username string) {
	p.Name = username
	p.Alive = true
	p.TurnOver = false
}

// moves player spaces in x, y dir
func (p *Player) Move(dir string) bool {
	switch dir {
	case "forwards":
		switch p.Facing {
		case "north":
			p.PosY++
		case "south":
			p.PosY--
		case "east":
			p.PosX++
		case "west":
			p.PosX--
		default:
			return false
		}
	case "backwards":
		switch p.Facing {
		case "north":
			p.PosY--
		case "south":
			p.PosY++
		case "east":
			p.PosX--
		case "west":
			p.PosX++
		default:
			return false
		}
	case "left":
		switch p.Facing {
		case "north":
			p.PosX--
		case "south":
			p.PosX++
		case "east":
			p.PosY++
		case "west":
			p.PosY--
		default:
			return false
		}
	case "right":
		switch p.Facing {
		case "north":
			p.PosX++
		case "south":
			p.PosX--
		case "east":
			p.PosY--
		case "west":
			p.PosY++
		default:
			return false
		}
	default:
		return false
	}
	return false
}

func (p *Player) Pass() bool {
	return true
}

type Class interface {
	Init()
}

type Soldier Player
type Medic Player

func (s *Soldier) Init() {
	s.Type = "soldier"
	s.HP = 5 + rand.IntN(3)
	s.Strength = 5 + rand.IntN(3)
	s.Luck = 2 + rand.IntN(1)
}

func (s *Medic) Init() {
	s.Type = "medic"
	s.HP = 7 + rand.IntN(4)
	s.Strength = 2 + rand.IntN(3)
	s.Luck = 3 + rand.IntN(2)
}

type Perk interface {
	BeforeRolls()
	AfterRolls()
}
