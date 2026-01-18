package state

import "github.com/firefly-zero/firefly-go/firefly"

var Input = InputState{}

type InputState struct {
	Me         firefly.Peer
	dpad4Old   firefly.DPad4
	dpad4      firefly.DPad4
	padOld     firefly.Pad
	hasPadOld  bool
	pad        firefly.Pad
	hasPad     bool
	buttonsOld firefly.Buttons
	buttons    firefly.Buttons
}

func (i *InputState) Boot() {
	i.Me = firefly.GetMe()
}

func (i *InputState) Update() {
	i.padOld = i.pad
	i.hasPadOld = i.hasPad
	pad, hasPad := firefly.ReadPad(i.Me)
	i.pad = pad
	i.hasPad = hasPad
	if hasPad {
		i.dpad4Old = i.dpad4
		i.dpad4 = pad.DPad4()
	} else {
		i.dpad4Old = i.dpad4
		i.dpad4 = firefly.DPad4None
	}

	i.buttonsOld = i.buttons
	i.buttons = firefly.ReadButtons(i.Me)
}

func (i *InputState) JustPressedButtons() firefly.Buttons {
	return i.buttons.JustPressed(i.buttonsOld)
}

func (i *InputState) JustPressedDPad4() firefly.DPad4 {
	return i.dpad4.JustPressed(i.dpad4Old)
}
