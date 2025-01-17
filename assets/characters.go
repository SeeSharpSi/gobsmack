package assets

import "math/rand/v2"

type Player struct {
	PosX, PosY int
	Type       string
	HP         int
	Strength   int
	Perks      []Perk
	Luck       int
	// x, y T if explored, F if not
	Explored [][]bool
}

func (p *Player) Init(s Ship) {
}

// moves player spaces in x, y dir
func (p *Player) Move(xmod, ymod int) {
	p.PosX += xmod
	p.PosY += ymod
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
	Before()
	During()
	After()
}
