package field

import (
	"github.com/applejag/epic-wizard-firefly-gladiators/assets"
	"github.com/applejag/epic-wizard-firefly-gladiators/pkg/util"
	"github.com/applejag/firefly-go-math/ffmath"
	"github.com/applejag/firefly-go-math/ffrand"

	"github.com/firefly-zero/firefly-go/firefly"
)

const (
	FPS                       = 60
	FireflyAnimationFPS       = 3
	FireflySpeedPixelsPerTick = 0.2
)

type Firefly struct {
	id            int
	sprites       util.AnimatedSheet
	spritesRev    util.AnimatedSheet
	spritesHat    util.AnimatedSheet
	spritesHatRev util.AnimatedSheet

	pos        ffmath.Vec
	nextPos    ffmath.Vec
	sleepTicks int
}

func NewFirefly(id int) Firefly {
	sprites := assets.FireflySheet.Animated(FireflyAnimationFPS)
	spritesRev := assets.FireflySheetRev.Animated(FireflyAnimationFPS)

	hatIdx := ffrand.Intn(11) * 2
	spritesHat := assets.FireflyHats[hatIdx : hatIdx+2].Animated(FireflyAnimationFPS)
	spritesHatRev := assets.FireflyHatsRev[hatIdx : hatIdx+2].Animated(FireflyAnimationFPS)
	frame := ffrand.Intn(len(assets.FireflySheet))
	sprites.SetFrame(frame)
	spritesRev.SetFrame(frame)
	spritesHat.SetFrame(frame)
	spritesHatRev.SetFrame(frame)
	return Firefly{
		id: id,
		pos: ffmath.V(
			float32(ffrand.IntRange(40, firefly.Width-40)),
			float32(ffrand.IntRange(40, firefly.Height-30)),
		),
		sleepTicks:    ffrand.IntRange(FPS, FPS*3),
		sprites:       sprites,
		spritesRev:    spritesRev,
		spritesHat:    spritesHat,
		spritesHatRev: spritesHatRev,
	}
}

func (f *Firefly) Update() {
	f.sprites.Update()
	f.spritesRev.Update()
	f.spritesHat.Update()
	f.spritesHatRev.Update()

	if f.sleepTicks > 0 {
		// stay idle for a bit
		f.sleepTicks--
		if f.sleepTicks <= 0 {
			f.nextPos = f.pos
			f.nextPos = ffmath.V(
				float32(ffrand.IntRange(40, firefly.Width-40)),
				float32(ffrand.IntRange(30, firefly.Height-30)),
			)
		}
	} else {
		// move to nextPos
		f.pos = f.pos.MoveTowards(f.nextPos, FireflySpeedPixelsPerTick)

		if f.nextPos.Sub(f.pos).RadiusSquared() < 10 {
			f.sleepTicks = ffrand.IntRange(FPS, FPS*3)
		}
	}
}

func (f *Firefly) Render() {
	point := f.pos.Round().Point()
	firefly.DrawCircle(point.Add(firefly.P(-2, 2)), 5, firefly.Solid(firefly.ColorDarkGray))

	if f.nextPos.X > f.pos.X {
		f.sprites.Draw(point.Sub(firefly.P(4, 5)))
		// TODO: hat rendering works, it's just the other stuff that's unimplemented
		// like buying hats, UI for chaning hats, including hats in save state
		// f.spritesHat.Draw(point.Sub(firefly.P(3, 6)))
	} else {
		f.spritesRev.Draw(point.Sub(firefly.P(4, 5)))
		// f.spritesHatRev.Draw(point.Sub(firefly.P(6, 6)))
	}
}
