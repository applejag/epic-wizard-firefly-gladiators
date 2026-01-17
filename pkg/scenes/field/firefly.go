package field

import (
	"firefly-jam-2026/assets"
	"firefly-jam-2026/pkg/util"

	"github.com/firefly-zero/firefly-go/firefly"
)

const FireflyAnimationFPS = 15

type Firefly struct {
	id         int
	pos        util.Vec2
	sprites    util.AnimatedSheet
	spritesRev util.AnimatedSheet
}

func NewFirefly(id int) Firefly {
	sprites := assets.FireflySheet.Animated(FireflyAnimationFPS)
	spritesRev := assets.FireflySheetRev.Animated(FireflyAnimationFPS)
	frame := util.RandomRange(0, len(assets.FireflySheet))
	sprites.SetFrame(frame)
	spritesRev.SetFrame(frame)
	return Firefly{
		id: id,
		pos: util.V(
			float32(util.RandomRange(40, firefly.Width-40)),
			float32(util.RandomRange(30, firefly.Height-30)),
		),
		sprites:    sprites,
		spritesRev: spritesRev,
	}
}

func (f *Firefly) Update() {
	f.sprites.Update()
	f.spritesRev.Update()
}

func (f *Firefly) Render() {
	f.sprites.Draw(f.pos.Point())
}
