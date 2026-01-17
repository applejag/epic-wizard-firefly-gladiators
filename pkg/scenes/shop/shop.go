package shop

import "github.com/firefly-zero/firefly-go/firefly"

type Shop struct {
	dpad4Old firefly.DPad4

	Selected int
	Items    []Item
}

func (f *Shop) Boot() {
}

func (f *Shop) Update() {
	if pad, ok := firefly.ReadPad(firefly.GetMe()); ok {
		dpad4 := pad.DPad4()
		justPressed := dpad4.JustPressed(f.dpad4Old)
		if justPressed != firefly.DPad4None {
			f.handleInput(justPressed)
		}
		f.dpad4Old = dpad4
	} else {
		f.dpad4Old = firefly.DPad4None
	}
}

func (f *Shop) handleInput(justPressed firefly.DPad4) {
}

func (f *Shop) Render() {
}

func (f *Shop) AddItemDrug(item Item) {
	f.Items = append(f.Items, item)
}

func (f *Shop) AddItem(item Item) {
	f.Items = append(f.Items, item)
}

type Item struct {
	Price    int
	Icon     firefly.SubImage
	Bg       firefly.SubImage
	Quantity int
}
