package racebattle

import (
	"firefly-jam-2026/assets"
	"firefly-jam-2026/pkg/util"

	"github.com/firefly-zero/firefly-go/firefly"
)

var path = []util.Vec2{
	util.V(68, 331),
	util.V(93, 368),
	util.V(107, 429),
	util.V(207, 448),
	util.V(240, 437),
	util.V(284, 351),
	util.V(309, 337),
	util.V(384, 343),
	util.V(442, 278),
	util.V(429, 254),
	util.V(374, 216),
	util.V(310, 175),
}

type World struct {
	Players []Firefly
	Camera  Camera
	Me      firefly.Peer
	// path Path
}

func (w *World) Update() {
	for i := range w.Players {
		w.Players[i].Update()
	}
	w.Camera.Update(w)
}

func (w *World) Render() {
	// Background
	firefly.ClearScreen(firefly.ColorDarkGray)
	assets.RacingMap.Draw(w.Camera.WorldPointToCameraSpace(firefly.P(0, 0)))
	// Players
	var me *Firefly
	for i, player := range w.Players {
		if player.Peer == w.Me {
			me = &w.Players[i]
		} else {
			player.Draw(w)
		}
	}
	// Draw my player last (on top)
	if me != nil {
		me.Draw(w)
	}
}
