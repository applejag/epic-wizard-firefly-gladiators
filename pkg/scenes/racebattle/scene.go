package racebattle

import (
	"cmp"
	"firefly-jam-2026/assets"
	"firefly-jam-2026/pkg/state"
	"firefly-jam-2026/pkg/util"
	"slices"

	"github.com/firefly-zero/firefly-go/firefly"
)

var path = []util.Vec2{
	util.V(41, 406),
	util.V(56, 451),
	util.V(125, 504),
	util.V(232, 576),
	util.V(420, 556),
	util.V(431, 516),
	util.V(418, 440),
	util.V(550, 327),
	util.V(542, 257),
	util.V(512, 212),
	util.V(470, 195),
	util.V(426, 112),
	util.V(361, 123),
	util.V(242, 252),
	util.V(213, 233),
	util.V(235, 135),
	util.V(226, 82),
	util.V(154, 84),
	util.V(98, 119),
	util.V(76, 172),
	util.V(78, 250),
	util.V(72, 288),
	util.V(79, 319),
	util.V(47, 372),
}

type Scene struct {
	AnimatedClouds util.AnimatedSheet
	Players        []Firefly
	Camera         Camera
}

func (s *Scene) Boot() {
	s.AnimatedClouds = assets.RacingMapClouds.Animated(2)
	// Sort by Y-axis so that they're drawn in the right order
	slices.SortFunc(s.Players, func(a, b Firefly) int {
		return cmp.Compare(a.Pos.Y, b.Pos.Y)
	})
}

func (s *Scene) Update() {
	for i := range s.Players {
		s.Players[i].Update()
	}
	s.Camera.Update(s)
	s.AnimatedClouds.Update()
}

func (s *Scene) Render() {
	// Background
	firefly.ClearScreen(firefly.ColorDarkGray)
	mapPos := s.Camera.WorldPointToCameraSpace(firefly.P(0, 0))
	assets.RacingMap.Draw(mapPos)
	assets.RacingMapTrees.Draw(mapPos)
	// Players
	var me *Firefly
	for i, player := range s.Players {
		if player.Peer == state.Input.Me {
			me = &s.Players[i]
		} else {
			player.Draw(s)
		}
	}
	// Draw my player last
	if me != nil {
		me.Draw(s)
	}
	// Draw tree tops layer on top
	assets.RacingMapTreetops.Draw(mapPos)
	s.AnimatedClouds.Draw(mapPos)
}

func (s *Scene) OnSceneEnter() {
	clear(s.Players)
	s.Players = s.Players[:0]
	for peer := range state.Game.InRaceBattle {
		// TODO: randomize the position a little
		s.Players = append(s.Players, NewFireflyPlayer(peer, util.V(41, 390), firefly.Degrees(270)))
	}
	// TODO: add AI players
	s.Camera.Update(s)
}
