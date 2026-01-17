package state

import (
	"firefly-jam-2026/pkg/util"
)

var (
	nextID int
	Game   GameState
)

type GameState struct {
	Fireflies []Firefly
}

func (g *GameState) AddFirefly() {
	nextID++
	name := util.RandomName()
	g.Fireflies = append(g.Fireflies, Firefly{
		ID:   nextID,
		Name: name,
	})
}

func (g *GameState) FindFireflyByID(id int) int {
	for idx := range g.Fireflies {
		if g.Fireflies[idx].ID == id {
			return idx
		}
	}
	return -1
}

type Firefly struct {
	ID   int
	Name util.Name
}
