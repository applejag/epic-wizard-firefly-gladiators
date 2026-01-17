package shop

import (
	"firefly-jam-2026/assets"

	"github.com/firefly-zero/firefly-go/firefly"
)

type Scene struct {
	Shop Shop
	Frog Frog
}

func (s *Scene) Boot() {
	s.Shop.Boot()
	s.Frog.Boot()
}

func (s *Scene) Update() {
	s.Shop.Update()
	s.Frog.Update()
}

func (s *Scene) Render() {
	assets.ShopBG.Draw(firefly.P(0, 0))

	s.Shop.Render()
	s.Frog.Render()
}
