package racebattle2

import (
	"github.com/applejag/epic-wizard-firefly-gladiators/assets"
	"github.com/applejag/epic-wizard-firefly-gladiators/pkg/util"
	"github.com/firefly-zero/firefly-go/firefly"
)

const FireflyAnimationFPS = 10.0

type Scene struct {
	fireflySprite util.AnimatedSheet
}

func (s *Scene) Boot() {
	s.fireflySprite = util.NewAnimatedSheet(assets.FireflySheet, FireflyAnimationFPS)
}

func (s *Scene) Update() {
	s.fireflySprite.Update()
}

func (s *Scene) Render() {
	firefly.ClearScreen(firefly.ColorGreen)
	// sky
	firefly.DrawRect(firefly.P(0, 0), firefly.S(firefly.Width, 60), firefly.Solid(firefly.ColorCyan))

	s.fireflySprite.Draw(firefly.P(50, 80))
}

func (s *Scene) OnSceneEnter() {
}
