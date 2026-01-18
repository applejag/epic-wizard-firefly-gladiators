package field

import (
	"firefly-jam-2026/assets"
	"firefly-jam-2026/pkg/state"
	"firefly-jam-2026/pkg/util"

	"github.com/firefly-zero/firefly-go/firefly"
)

type FireflyModal struct {
	isOpen          bool
	scrollOpenAnim  util.AnimatedSheet
	scrollCloseAnim util.AnimatedSheet
	scrollSprite    firefly.SubImage
	firefly         *Firefly
}

func (m *FireflyModal) IsOpen() bool {
	return m.isOpen || !m.scrollCloseAnim.IsPaused()
}

func (m *FireflyModal) IsClosing() bool {
	return !m.isOpen && !m.scrollCloseAnim.IsPaused()
}

func (m *FireflyModal) Open(firefly *Firefly) {
	m.scrollOpenAnim.Play()
	m.isOpen = true
	m.firefly = firefly
}

func (m *FireflyModal) Close() {
	if m.IsClosing() {
		return
	}

	m.scrollCloseAnim.Play()
	m.firefly = nil
	m.isOpen = false
}

func (m *FireflyModal) Boot() {
	m.scrollOpenAnim = assets.ScrollOpen.Animated(12)
	m.scrollOpenAnim.AutoPlay = false
	m.scrollOpenAnim.Stop()
	m.scrollCloseAnim = assets.ScrollClose.Animated(12)
	m.scrollCloseAnim.AutoPlay = false
	m.scrollCloseAnim.Stop()
	m.scrollSprite = assets.ScrollClose[0]
}

func (m *FireflyModal) Update() {
	m.scrollOpenAnim.Update()
	m.scrollCloseAnim.Update()

	if m.IsClosing() {
		return
	}

	buttons := firefly.ReadButtons(firefly.GetMe())
	if buttons.E {
		m.Close()
	}
}

func (m *FireflyModal) Render() {
	const scrollWidth = 111
	point := firefly.P(firefly.Width/2-scrollWidth/2, 24)
	m.scrollOpenAnim.Draw(point)
	m.scrollCloseAnim.Draw(point)

	if m.isOpen && m.scrollCloseAnim.IsPaused() && m.scrollOpenAnim.IsPaused() {
		m.renderScroll(point)
	}
}

func (m *FireflyModal) renderScroll(point firefly.Point) {
	m.scrollSprite.Draw(point)

	dataIndex := state.Game.FindFireflyByID(m.firefly.id)
	if dataIndex == -1 {
		panic("should never be -1 here")
	}
	data := state.Game.Fireflies[dataIndex]

	const scrollInnerWidth = 77

	innerScrollPoint := point.Add(firefly.P(18, 18))

	text := util.WordWrap(
		data.Name.String(),
		scrollInnerWidth,
		assets.FontEG_6x9.CharWidth(),
	)

	assets.FontEG_6x9.Draw(text, innerScrollPoint.Add(firefly.P(0, 10)), firefly.ColorDarkGray)
}
