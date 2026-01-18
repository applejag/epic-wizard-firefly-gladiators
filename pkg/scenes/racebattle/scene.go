package racebattle

import (
	"cmp"
	"slices"

	"github.com/applejag/firefly-jam-2026/assets"
	"github.com/applejag/firefly-jam-2026/pkg/state"
	"github.com/applejag/firefly-jam-2026/pkg/util"

	"github.com/firefly-zero/firefly-go/firefly"
)

type Scene struct {
	AnimatedClouds util.AnimatedSheet
	Players        []Firefly
	Standing       []Standing
	Camera         Camera
}

type Standing struct {
	Progress float32
	*Firefly
}

func (s *Scene) Boot() {
	s.AnimatedClouds = assets.RacingMapClouds.Animated(2)
}

func (s *Scene) Update() {
	for i := range s.Players {
		s.Players[i].Update()
	}
	s.nudgeFirefliesAwayFromEachOther()
	s.updateStanding()
	// Sort by Y-axis so that they're drawn in the right order
	slices.SortFunc(s.Players, func(a, b Firefly) int {
		return cmp.Compare(a.Pos.Y, b.Pos.Y)
	})

	s.Camera.Update(s)
	s.AnimatedClouds.Update()
}

func (s *Scene) nudgeFirefliesAwayFromEachOther() {
	for i := 0; i < len(s.Players); i++ {
		for j := i + 1; j < len(s.Players); j++ {
			s.Players[i].MoveAwayFrom(&s.Players[j])
		}
	}
}

func (s *Scene) updateStanding() {
	clear(s.Standing)
	if len(s.Standing) < len(s.Players) {
		s.Standing = slices.Grow(s.Standing, len(s.Players)-len(s.Standing))
	}
	s.Standing = s.Standing[:len(s.Players)]

	for i, player := range s.Players {
		s.Standing[i] = Standing{
			Progress: player.PathTracker.Progress(player.Pos),
			Firefly:  &s.Players[i],
		}
	}

	// sort so the highest progress is on index 0
	slices.SortFunc(s.Standing, func(a, b Standing) int {
		return cmp.Compare(b.Progress, a.Progress)
	})
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
		if player.IsPlayer && player.Peer == state.Input.Me {
			me = &s.Players[i]
		} else {
			player.Render(s)
		}
	}
	// Draw my player last
	if me != nil {
		me.Render(s)
	}
	// Draw tree tops layer on top
	assets.RacingMapTreetops.Draw(mapPos)
	s.AnimatedClouds.Draw(mapPos)
}

func (s *Scene) OnSceneEnter() {
	clear(s.Players)
	s.Players = s.Players[:0]
	for peer := range state.Game.InRaceBattle {
		s.Players = append(s.Players, NewFireflyPlayer(peer, util.V(41, 390).Add(offsetForPlayer(len(s.Players))), firefly.Degrees(271)))
	}
	if len(s.Players) < 2 {
		s.Players = append(s.Players, NewFireflyAI(util.V(41, 390).Add(offsetForPlayer(len(s.Players))), firefly.Degrees(271)))
	}
	s.Camera.Update(s)
}

func offsetForPlayer(index int) util.Vec2 {
	if index == 0 {
		return util.V(0, 0)
	}
	angle := firefly.Degrees(60 * float32(index-1))
	return util.AngleToVec2(angle).Scale(12)
}
