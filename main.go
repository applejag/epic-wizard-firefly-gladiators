package main

import (
	"github.com/applejag/epic-wizard-firefly-gladiators/assets"
	"github.com/applejag/epic-wizard-firefly-gladiators/pkg/game"
	"github.com/applejag/epic-wizard-firefly-gladiators/pkg/scenes"
	"github.com/applejag/epic-wizard-firefly-gladiators/pkg/state"

	"github.com/firefly-zero/firefly-go/firefly"
)

func init() {
	firefly.Boot = boot
	firefly.Update = update
	firefly.Render = render
}

var scenemanager = game.SceneManager{}

func boot() {
	assets.Load()

	state.Input.Boot()
	scenemanager.Boot()

	// TODO: switching to this scene from start while I'm developing the scene code
	scenemanager.SwitchSceneNoTransition(scenes.RacingBattle2)
}

func update() {
	state.Input.Update()
	scenemanager.Update()
}

func render() {
	scenemanager.Render()
}
