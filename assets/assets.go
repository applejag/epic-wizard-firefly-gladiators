package assets

import (
	"slices"

	"github.com/applejag/epic-wizard-firefly-gladiators/pkg/util"

	"github.com/firefly-zero/firefly-go/firefly"
)

var (
	buf [1366914]byte

	Field            firefly.Image
	FireflyHighlight util.SpriteSheet
	ScrollClose      util.SpriteSheet
	ScrollOpen       util.SpriteSheet
	TournamentButton util.SpriteSheet
	TrainButton      util.SpriteSheet
	ShopButton       util.SpriteSheet

	RacingMap         firefly.Image
	RacingMapTrees    firefly.Image
	RacingMapTreetops firefly.Image
	RacingMapClouds   util.SpriteSheet
	RacingMapMask     firefly.Image
	RacingPlace       util.SpriteSheet
	VictorySplash     util.SpriteSheet
	DefeatSplash      util.SpriteSheet
	FireflySheet      util.SpriteSheet
	FireflySheetRev   util.SpriteSheet

	FireflyHats    util.SpriteSheet
	FireflyHatsRev util.SpriteSheet

	TitleScreen          util.SpriteSheet
	TitleButtonHighlight util.SpriteSheet
	TitleNoContinue      firefly.Image

	ShopBG      firefly.Image
	ShopFrog    util.SpriteSheet
	ShopProps   util.SpriteSheet
	ShopChatbox firefly.Image
	ShopItem    util.SpriteSheet

	Exit firefly.Image

	TransitionSheet util.SpriteSheet

	CashBanner firefly.Image

	FontEG_6x9    firefly.Font
	FontPico8_4x6 firefly.Font
)

func Load() {
	loader := Loader{buf: buf[:]}

	Field = loader.LoadImage("field")
	FireflyHighlight = util.SplitImageBySize(loader.LoadImage("firefly-hi"), firefly.S(32, 32))
	allHats := util.SplitImageBySize(loader.LoadImage("firefly-hats"), firefly.S(10, 8))
	FireflyHats = allHats[:22]
	FireflyHatsRev = allHats[22:]
	ScrollClose = util.SplitImageByCount(loader.LoadImage("scroll"), firefly.S(4, 1))
	ScrollOpen = slices.Clone(ScrollClose)
	slices.Reverse(ScrollOpen)
	TournamentButton = util.SplitImageByCount(loader.LoadImage("tournament-btn"), firefly.S(7, 1))
	TrainButton = util.SplitImageBySize(loader.LoadImage("train-btn"), firefly.S(61, 14))
	ShopButton = util.SplitImageBySize(loader.LoadImage("shop-btn"), firefly.S(42, 14))
	RacingMap = loader.LoadImage("racing-map")
	RacingMapTrees = loader.LoadImage("racing-map-trees")
	RacingMapTreetops = loader.LoadImage("racing-map-treetops")
	RacingMapClouds = util.SplitImageByCount(loader.LoadImage("racing-map-clouds"), firefly.S(2, 1))
	RacingMapMask = loader.LoadImage("racing-map-mask")
	RacingPlace = util.SplitImageBySize(loader.LoadImage("racing-place"), firefly.S(28, 33))
	VictorySplash = util.SplitImageByCount(loader.LoadImage("victory-splash"), firefly.S(3, 3))
	DefeatSplash = util.SplitImageByCount(loader.LoadImage("defeat-splash"), firefly.S(3, 3))
	fireflyCombinedSheet := util.SplitImageBySize(loader.LoadImage("firefly"), firefly.S(9, 10))
	FireflySheet = fireflyCombinedSheet[0:2]
	FireflySheetRev = fireflyCombinedSheet[2:4]
	TitleScreen = util.SplitImageByCount(loader.LoadImage("title-screen"), firefly.S(2, 1))
	TitleButtonHighlight = util.SplitImageByCount(loader.LoadImage("title-button-hi"), firefly.S(2, 1))
	TitleNoContinue = loader.LoadImage("title-no-continue")
	ShopBG = loader.LoadImage("shop-bg")
	ShopFrog = util.SplitImageByCount(loader.LoadImage("shop-frog"), firefly.S(3, 2))
	ShopProps = util.SplitImageByCount(loader.LoadImage("shop-props"), firefly.S(3, 2))
	ShopChatbox = loader.LoadImage("shop-chatbox")
	ShopItem = util.SplitImageByCount(loader.LoadImage("shop-item"), firefly.S(4, 3))
	TransitionSheet = util.SplitImageByCount(loader.LoadImage("transition"), firefly.S(4, 4))
	Exit = loader.LoadImage("exit")
	CashBanner = loader.LoadImage("cash-banner")
	FontEG_6x9 = loader.LoadFont("eg_6x9")
	FontPico8_4x6 = loader.LoadFont("pico8_4x6")

	loader.AssertBufferFullyUtilized()
}

type Loader struct {
	buf  []byte
	used int
}

func (loader *Loader) LoadFile(path string) firefly.File {
	// Have to get the file size separately, even though LoadFile returns the size
	// https://github.com/firefly-zero/firefly-runtime/issues/7
	fileSize := firefly.GetFileSize(path)
	file := firefly.LoadFile(path, loader.buf[:fileSize])
	written := len(file.Bytes())
	loader.used += written
	loader.buf = loader.buf[written:]
	return file
}

func (loader *Loader) LoadImage(path string) firefly.Image {
	return loader.LoadFile(path).Image()
}

func (loader *Loader) LoadFont(path string) firefly.Font {
	return loader.LoadFile(path).Font()
}

func (loader *Loader) AssertBufferFullyUtilized() {
	var buf [100]byte
	n := copy(buf[:], "All assets loaded. Total size: ")
	n += util.FormatIntInto(buf[n:], loader.used)
	firefly.LogDebug(string(buf[:n]))

	if len(loader.buf) > 0 {
		var buf [100]byte
		n := copy(buf[:], "Assets buffer should be smaller. Want: [")
		n += util.FormatIntInto(buf[n:], loader.used)
		n += copy(buf[n:], "]byte, but got [")
		n += util.FormatIntInto(buf[n:], loader.used+len(loader.buf))
		n += copy(buf[n:], "]byte")
		err := string(buf[:n])
		firefly.LogError(err)
		panic("assets.Loader: buffer was not fully utilized")
	}
}
