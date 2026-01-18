package field

import (
	"strconv"
	"strings"

	"github.com/applejag/firefly-jam-2026/assets"
	"github.com/applejag/firefly-jam-2026/pkg/state"
	"github.com/applejag/firefly-jam-2026/pkg/util"
	"github.com/firefly-zero/firefly-go/firefly"
	"github.com/orsinium-labs/tinymath"
)

type StatsPage struct {
	changeHatBtn      Button
	giveVitaminsBtn   Button
	playTournamentBtn Button
	focused           ButtonKind
}

func (p *StatsPage) Boot() {
	p.changeHatBtn = NewButton(ButtonChangeHat, "CHANGE HAT")
	p.changeHatBtn.Disabled = true
	p.giveVitaminsBtn = NewButton(ButtonGiveVitamins, "GIVE VITAMINS")
	p.giveVitaminsBtn.Disabled = true
	p.playTournamentBtn = NewButton(ButtonTournament, "RACING")
}

func (p *StatsPage) Update(modal *FireflyModal) {
	p.changeHatBtn.Update()
	p.giveVitaminsBtn.Update()
	p.playTournamentBtn.Update()

	if justPressed := state.Input.JustPressedDPad4(); justPressed != firefly.DPad4None {
		p.handleInputDPad4(justPressed)
	}
	if justPressed := state.Input.JustPressedButtons(); justPressed.Any() {
		p.handleInputButtons(justPressed, modal)
	}
}

func (p *StatsPage) handleInputDPad4(justPressed firefly.DPad4) {
	switch justPressed {
	case firefly.DPad4Up:
		p.focused = p.focused.Up()
	case firefly.DPad4Down:
		p.focused = p.focused.Down()
	}
}

func (p *StatsPage) handleInputButtons(justPressed firefly.Buttons, modal *FireflyModal) {
	switch {
	case justPressed.S:
		switch p.focused {
		case ButtonChangeHat:
			// Shake to signify that the button doesn't work
			p.changeHatBtn.Shake()
			// TODO: allow transition to hats page
			// m.state = ModalHats
		case ButtonGiveVitamins:
			// Shake to signify that the button doesn't work
			p.giveVitaminsBtn.Shake()
		case ButtonTournament:
			modal.state = ModalTournament

			// state.Game.AddMyFireflyToRaceBattle(m.firefly.id)
			// m.CloseWithoutTransition()
			// scenes.SwitchScene(scenes.RacingTraining)
		}
	}
}

func (p *StatsPage) Render(innerScrollPoint firefly.Point, fireflyID int) {
	dataIndex := state.Game.FindFireflyByID(fireflyID)
	if dataIndex == -1 {
		panic("should never be -1 here")
	}
	data := state.Game.Fireflies[dataIndex]

	text := util.WordWrap(
		data.Name.String(),
		scrollInnerWidth,
		assets.FontEG_6x9.CharWidth(),
	)

	charHeight := assets.FontEG_6x9.CharHeight()

	textPos := innerScrollPoint.Add(firefly.P(0, 10))
	assets.FontEG_6x9.Draw(text, textPos, firefly.ColorDarkGray)
	textHeight := charHeight * (strings.Count(text, "\n") + 1)

	speedPoint := textPos.Add(firefly.P(2, textHeight))
	assets.FontEG_6x9.Draw(strconv.Itoa(data.Speed), speedPoint, firefly.ColorBlack)
	assets.FontPico8_4x6.Draw("SPEED", speedPoint.Add(firefly.P(0, charHeight)), firefly.ColorGray)

	nimblenessPoint := speedPoint.Add(firefly.P(32, 0))
	assets.FontEG_6x9.Draw(strconv.Itoa(data.Nimbleness), nimblenessPoint, firefly.ColorBlack)
	assets.FontPico8_4x6.Draw("NIMBLE", nimblenessPoint.Add(firefly.P(0, charHeight)), firefly.ColorGray)

	rectPoint := textPos.Add(firefly.P(64, textHeight+4-charHeight))
	rectSize := firefly.S(22, 22)
	firefly.DrawRoundedRect(rectPoint, rectSize, firefly.S(3, 3), firefly.Outlined(firefly.ColorGray, 1))

	assets.FireflySheet[0].Draw(rectPoint.Add(firefly.P(6, 6)))

	changeHatPoint := innerScrollPoint.Add(firefly.P(0, scrollInnerHeight-26))
	p.changeHatBtn.Render(changeHatPoint, p.focused)

	giveVitaminsPoint := changeHatPoint.Add(firefly.P(0, 8))
	p.giveVitaminsBtn.Render(giveVitaminsPoint, p.focused)

	tournamentPoint := giveVitaminsPoint.Add(firefly.P(0, 8))
	p.playTournamentBtn.Render(tournamentPoint, p.focused)

	// m.tournamentAnim.Draw(tournamentPoint.Add(firefly.P(8, -9)))
	// m.playTournamentBtn.Render(tournamentPoint, m.focused)
	// assets.TrainButton.Draw(tournamentPoint.Add(firefly.P(72, -11)))
}

type ButtonKind byte

const (
	ButtonNone ButtonKind = iota
	ButtonChangeHat
	ButtonGiveVitamins
	ButtonTournament

	buttonCount = 4
)

func (k ButtonKind) Down() ButtonKind {
	switch k {
	case ButtonChangeHat:
		return ButtonGiveVitamins
	case ButtonGiveVitamins:
		return ButtonTournament
	case ButtonNone:
		return ButtonChangeHat
	case ButtonTournament:
		return ButtonChangeHat
	default:
		panic("unexpected field.ButtonKind")
	}
}

func (k ButtonKind) Up() ButtonKind {
	switch k {
	case ButtonChangeHat:
		return ButtonTournament
	case ButtonGiveVitamins:
		return ButtonChangeHat
	case ButtonNone:
		return ButtonTournament
	case ButtonTournament:
		return ButtonGiveVitamins
	default:
		panic("unexpected field.ButtonKind")
	}
}

const ButtonShakeDuration = 45

type Button struct {
	kind     ButtonKind
	text     string
	Disabled bool
	shake    int
}

func NewButton(kind ButtonKind, text string) Button {
	return Button{
		kind: kind,
		text: text,
	}
}

func (b *Button) Update() {
	if b.shake > 0 {
		b.shake--
	}
}

func (b *Button) Render(point firefly.Point, focused ButtonKind) {
	prefix := "- "
	color := firefly.ColorGray
	if focused == b.kind {
		prefix = "> "
		color = firefly.ColorBlack
	}
	if b.Disabled {
		color = firefly.ColorLightGray
	}
	assets.FontPico8_4x6.Draw(prefix, point, color)
	if b.text != "" {
		if b.shake > 0 {
			t := float32(b.shake) / ButtonShakeDuration
			point = point.Add(firefly.P(int(tinymath.Sin(t*45)*t*4), 0))
		}
		textPoint := point.Add(firefly.P(assets.FontPico8_4x6.LineWidth(prefix), 0))
		assets.FontPico8_4x6.Draw(b.text, textPoint, color)
		if b.Disabled {
			// Draw strikethrough
			firefly.DrawLine(textPoint, textPoint.Add(firefly.P(
				assets.FontPico8_4x6.LineWidth(b.text),
				-assets.FontEG_6x9.CharHeight()/2,
			)), firefly.L(firefly.ColorGray, 1))
		}
	}
}

func (b *Button) Shake() {
	b.shake = ButtonShakeDuration
}
