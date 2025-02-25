package assets

import "math/rand/v2"

type Player struct {
	Name string
	// North South East West
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
	// still need a screenrender function. not sure where to put it
	// it could go here? or could exist on the game? WAIT it should exist on the game
	// and iterate over each player... or then it could exist on each player and
	// still be iterated over
	// actually, it should exist on game because game can access everything (aliens, ship, etc)
	CurrentScreen string
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
