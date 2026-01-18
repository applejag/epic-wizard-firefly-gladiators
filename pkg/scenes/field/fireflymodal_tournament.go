package field

import (
	"github.com/applejag/firefly-jam-2026/assets"
	"github.com/applejag/firefly-jam-2026/pkg/util"
	"github.com/firefly-zero/firefly-go/firefly"
)

type TournamentPage struct {
	tournamentAnim util.AnimatedSheet
}

func (p *TournamentPage) Boot() {
	p.tournamentAnim = assets.TournamentButton.Animated(6)
}

func (p *TournamentPage) Update() {
	p.tournamentAnim.Update()
}

func (p *TournamentPage) Render(innerScrollPoint firefly.Point) {
}
