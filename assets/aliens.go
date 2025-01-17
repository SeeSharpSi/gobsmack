package assets

import "math/rand/v2"

// Generic alien type
type Generic struct {
	HP       int
	Strength int
}

// Returns a new Generic alien. HP is 3-6
func (a Generic) New() Generic {
	buff := rand.IntN(4)
	return Generic{
		HP:       buff + 3,
		Strength: 2,
	}
}
