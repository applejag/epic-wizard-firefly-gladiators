package field

import (
	"cmp"
	"firefly-jam-2026/assets"
	"firefly-jam-2026/pkg/state"
	"slices"

	"github.com/firefly-zero/firefly-go/firefly"
)

type Scene struct {
	fireflies []Firefly
}

func (s *Scene) Boot() {
}

func (s *Scene) Update() {
	for i := range s.fireflies {
		s.fireflies[i].Update()
	}
	// Sort by Y-axis so that they're drawn in the right order
	slices.SortFunc(s.fireflies, func(a, b Firefly) int {
		return cmp.Compare(a.pos.Y, b.pos.Y)
	})
}

func (s *Scene) Render() {
	firefly.ClearScreen(firefly.ColorBlack)
	assets.Field.Draw(firefly.P(0, 0))

	for i := range s.fireflies {
		s.fireflies[i].Render()
	}
}

func (s *Scene) FindFireflyByID(id int) int {
	for idx := range s.fireflies {
		if s.fireflies[idx].id == id {
			return idx
		}
	}
	return -1
}

func (s *Scene) OnSceneSwitch() {
	for _, f := range state.Game.Fireflies {
		idx := s.FindFireflyByID(f.ID)
		if idx == -1 {
			s.fireflies = append(s.fireflies, NewFirefly(f.ID))
		}
	}
	s.fireflies = slices.DeleteFunc(s.fireflies, func(f Firefly) bool {
		return state.Game.FindFireflyByID(f.id) == -1
	})
}
