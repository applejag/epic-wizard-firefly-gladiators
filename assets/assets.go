package assets

import (
	"firefly-jam-2026/pkg/util"

	"github.com/firefly-zero/firefly-go/firefly"
)

var (
	Field           firefly.Image
	RacingMap       firefly.Image
	FireflySheet    util.SpriteSheet
	FireflySheetRev util.SpriteSheet
)

func Load() {
	Field = firefly.LoadImage("field", nil)
	RacingMap = firefly.LoadImage("racing-map", nil)
	FireflySheet = util.SplitImageByCount(firefly.LoadImage("firefly", nil), firefly.S(7, 1))
	FireflySheetRev = util.SplitImageByCount(firefly.LoadImage("firefly-rev", nil), firefly.S(7, 1))
}
