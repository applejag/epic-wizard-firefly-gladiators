package mainmenu

import (
	"firefly-jam-2026/assets"
	"firefly-jam-2026/pkg/util"

	"github.com/firefly-zero/firefly-go/firefly"
)

type Menu struct {
	TitleScreen util.AnimatedSheet
	Transition  Transition
}

func (w *Menu) Boot() {
	w.TitleScreen = assets.TitleScreen.Animated(2)
	w.Transition = NewTransition(assets.TransitionSheet.Animated(10))
}

func (w *Menu) Update() {
	w.TitleScreen.Update()
	w.Transition.Update()

	if firefly.ReadButtons(firefly.Combined).S {
		w.Transition.Play()
	}
}

func (w *Menu) Render() {
	w.TitleScreen.Draw(firefly.P(0, 0))
	w.Transition.Draw()
}
