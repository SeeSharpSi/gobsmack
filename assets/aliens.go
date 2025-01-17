package assets

import "math/rand/v2"

type Alien interface {
	Init()
	Damage()
	Heal()
}

// GenericAlien alien type
type AlienDef struct {
	Type     string
	HP       int
	Strength int
}

type GenericAlien AlienDef
type GaurdAlien AlienDef

// Inits a GenericAlien alien. HP is 3-5
func (a *GenericAlien) Init() {
	a.Type = "generic"
	a.HP = 3 + rand.IntN(3)
	a.Strength = 2 + rand.IntN(4)
}
func (a *GenericAlien) Damange() {
}
func (a *GenericAlien) Heal() {
}

// Inits a GaurdAlien alien. HP is 5-7
func (a *GaurdAlien) Init() {
	a.Type = "gaurd"
	a.HP = 5 + rand.IntN(3)
	a.Strength = 3 + rand.IntN(5)
}
func (a *GaurdAlien) Damange() {
}
func (a *GaurdAlien) Heal() {
}
