package field

import (
	"firefly-jam-2026/assets"

	"github.com/firefly-zero/firefly-go/firefly"
)

type Scene struct{}

func (s *Scene) Boot() {
}

func (s *Scene) Update() {
}

func (s *Scene) Render() {
	firefly.ClearScreen(firefly.ColorBlack)

	assets.Field.Draw(firefly.P(0, 0))
}
