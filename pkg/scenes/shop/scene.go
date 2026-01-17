package shop

import (
	"firefly-jam-2026/assets"

	"github.com/firefly-zero/firefly-go/firefly"
)

type Scene struct {
	Frog Frog
}

func (s *Scene) Boot() {
	s.Frog.Boot()
}

func (s *Scene) Update() {
	s.Frog.Update()
}

func (s *Scene) Render() {
	assets.Shop.Draw(firefly.P(0, 0))

	s.Frog.Render()
}
