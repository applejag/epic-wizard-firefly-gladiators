package shop

import (
	"firefly-jam-2026/assets"

	"github.com/firefly-zero/firefly-go/firefly"
)

const chatWobbleTicks = 40

type Frog struct {
	chatWobbleTime   int
	chatWobbleOffset int
}

func (f *Frog) Boot() {
}

func (f *Frog) Update() {
	f.chatWobbleTime++
	if f.chatWobbleTime >= chatWobbleTicks {
		f.chatWobbleTime -= chatWobbleTicks
		f.chatWobbleOffset = 1 - f.chatWobbleOffset
	}
}

func (f *Frog) Render() {
	assets.ShopChatbox.Draw(firefly.P(4, 22+f.chatWobbleOffset))
	// fits ~17 chars per line, and max 2 lines
	assets.FontEG_6x9.Draw("oy m8, u here 4\nsum foirefloies??", firefly.P(12, 38), firefly.ColorBlack)
}
