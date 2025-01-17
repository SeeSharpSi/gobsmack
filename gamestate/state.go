package gamestate

import "seesharpsi/gobsmack/assets"

type GameState struct {
	Players map[User]assets.Player
	Aliens  []assets.Alien
	Ship    assets.Ship
}

type User string

type PlayerStats struct {
}
