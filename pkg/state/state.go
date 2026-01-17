package state

import (
	"firefly-jam-2026/pkg/util"
)

var Game GameState

type GameState struct {
	Fireflies []Firefly
}

func (g *GameState) AddFirefly() {
	name := util.RandomName()
	g.Fireflies = append(g.Fireflies, Firefly{
		Name: name,
	})
}

type Firefly struct {
	Name util.Name
}
