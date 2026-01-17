package state

var state GameState

type GameState struct {
	Fireflies []Firefly
}

type Firefly struct {
	Name Name
}

type Name [2]byte
