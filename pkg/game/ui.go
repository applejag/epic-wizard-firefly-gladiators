package game

import (
	"bytes"

	"github.com/applejag/epic-wizard-firefly-gladiators/assets"
	"github.com/applejag/epic-wizard-firefly-gladiators/pkg/state"

	"github.com/firefly-zero/firefly-go/firefly"
)

var (
	moneyTextBuf      [4]byte
	firefliesCountBuf [2]byte
)

type UI struct {
	cachedMoneyValue     int
	cachedFirefliesValue int
}

func (u *UI) Render() {
	if len(state.Game.Fireflies) == 0 {
		// wait until player has at least 1 firefly
		return
	}
	assets.CashBanner.Draw(firefly.P(2, 2))

	money := min(state.Game.Money, 9999)
	if moneyTextBuf[0] == 0 || money != u.cachedMoneyValue {
		formatPaddedIntInto(moneyTextBuf[:], money, 4)
		u.cachedMoneyValue = money
	}
	fireflies := min(len(state.Game.Fireflies), 99)
	if firefliesCountBuf[0] == 0 || fireflies != u.cachedFirefliesValue {
		formatPaddedIntInto(firefliesCountBuf[:], fireflies, 2)
		u.cachedFirefliesValue = fireflies
	}

	drawRightAlignedWithColoredZeros(
		assets.FontEG_6x9,
		moneyTextBuf[:],
		firefly.P(29, 12),
		firefly.ColorLightGray,
		firefly.ColorDarkGray)
	drawRightAlignedWithColoredZeros(
		assets.FontEG_6x9,
		firefliesCountBuf[:],
		firefly.P(59, 12),
		firefly.ColorLightGray,
		firefly.ColorDarkGray)
}

func drawRightAlignedWithColoredZeros(font firefly.Font, text []byte, right firefly.Point, zeroColor, textColor firefly.Color) {
	width := font.CharWidth() * len(text)
	left := right.Add(firefly.P(-width, 0))
	withoutZeros := bytes.TrimLeft(text, "0")
	zeros := text[:len(text)-len(withoutZeros)]
	font.DrawBytes(zeros, left, zeroColor)
	font.DrawBytes(withoutZeros, left.Add(firefly.P(font.CharWidth()*len(zeros), 0)), textColor)
}

func formatPaddedIntInto(buf []byte, value, width int) {
	for i := range width {
		buf[i] = '0'
	}
	index := len(buf) - 1
	for value > 0 && index >= 0 {
		buf[index] = '0' + byte(value%10)
		value /= 10
		index -= 1
	}
}
