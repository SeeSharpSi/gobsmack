package assets

import "math/rand/v2"

type Player struct {
	PosX, PosY int
	Class
}

type Soldier struct {
}

type Class struct {
	Type     string
	HP       int
	Strength int
	Perks    []Perk
	Luck     int
}

// Returns a class of type soldier
func (c Class) NewSoldier() Class {
	HPmod := rand.IntN(3)
	Strmod := rand.IntN(3)
	Luckmod := rand.IntN(1)
	return Class{
		Type:     "soldier",
		HP:       5 + HPmod,
		Strength: 5 + Strmod,
		Luck:     2 + Luckmod,
	}
}

type Perk struct {
}
