package mainmenu

import (
	"firefly-jam-2026/pkg/util"

	"github.com/firefly-zero/firefly-go/firefly"
)

type Transition struct {
	util.AnimatedSheet
	size firefly.Size
}

func NewTransition(sprites util.AnimatedSheet) Transition {
	sprites.AutoPlay = false
	sprites.Stop()
	return Transition{AnimatedSheet: sprites}
}

func (t *Transition) Draw() {
	const size = 8
	// tile the sprite
	for x := 0; x < firefly.Width; x += size {
		for y := 0; y < firefly.Height; y += size {
			t.AnimatedSheet.Draw(firefly.P(x, y))
		}
	}
}
