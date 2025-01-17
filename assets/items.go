package assets

import "math/rand/v2"

type Item interface {
	Init()
	Use() bool
	Pickup()
}

type Gun struct {
	Power int
	Ammo  int
}

func (g *Gun) Init() {
	g.Power = rand.IntN(3)
	g.Ammo = rand.IntN(7)
}

func (g *Gun) Use() bool {
	if g.Ammo > 0 {
		g.Ammo--
		return true
	}
	return false
}

func (g *Gun) Pickup() {
	g.Init()
}

type Flashlight struct {
	Battery int
	On      bool
}

func (f *Flashlight) Init() {
	f.Battery = 20 + rand.IntN(80)
	f.On = true
}

func (f *Flashlight) Use() {
	if !f.On && f.Battery > 0 {
		f.On = true
	} else {
		f.On = false
	}
}

func (f *Flashlight) Pickup() {
	f.Init()
}
