package mainmenu

import (
	"firefly-jam-2026/assets"
	"firefly-jam-2026/pkg/util"

	"github.com/firefly-zero/firefly-go/firefly"
)

type Menu struct {
	TitleScreen util.AnimatedSheet
}

func (w *Menu) Boot() {
	w.TitleScreen = assets.TitleScreen.Animated(2)
}

func (w *Menu) Update() {
	w.TitleScreen.Update()
}

func (w *Menu) Render() {
	w.TitleScreen.Draw(firefly.P(0, 0))
}
