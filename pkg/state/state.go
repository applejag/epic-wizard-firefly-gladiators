package state

import (
	"github.com/applejag/epic-wizard-firefly-gladiators/pkg/util"
	"github.com/applejag/firefly-go-math/ffrand"

	"github.com/firefly-zero/firefly-go/firefly"
)

var (
	nextID int
	Game   = GameState{
		InRaceBattle: map[firefly.Peer]Firefly{},
	}
)

type Firefly struct {
	ID            int
	Name          util.Name
	Speed         int
	Nimbleness    int
	BattlesPlayed int
	BattlesWon    int
	Hat           int
}

type GameState struct {
	Fireflies          []Firefly
	BattlesPlayedTotal int
	BattlesWonTotal    int
	Money              int

	// Not saved, it's only ephemeral data
	InRaceBattle map[firefly.Peer]Firefly
}

func (g *GameState) AddFirefly() int {
	nextID++
	name := util.RandomName()
	randomness := ffrand.IntRange(8, 14)
	g.Fireflies = append(g.Fireflies, Firefly{
		ID:         nextID,
		Name:       name,
		Speed:      randomness,
		Nimbleness: 8 + (14 - randomness),
	})
	g.Save()
	return nextID
}

func (g *GameState) FindFireflyByID(id int) int {
	for idx := range g.Fireflies {
		if g.Fireflies[idx].ID == id {
			return idx
		}
	}
	return -1
}

func (g *GameState) AddMyFireflyToRaceBattle(id int) {
	dataIndex := g.FindFireflyByID(id)
	if dataIndex == -1 {
		panic("should never be -1 here")
	}
	g.InRaceBattle[Input.Me] = g.Fireflies[dataIndex]
}

func (g *GameState) RemoveMyFireflyFromRaceBattle() {
	delete(g.InRaceBattle, Input.Me)
}

func (g *GameState) Save() {
	var saveBuf [100]byte
	save, err := g.AppendBinary(saveBuf[:0])
	if err != nil {
		firefly.LogError("failed to save: " + err.Error())
		return
	}
	firefly.DumpFile("save", save)

	var buf [32]byte
	n := copy(buf[0:], "saved game, size: ")
	n += util.FormatIntInto(buf[n:], len(save))
	n += copy(buf[n:], " B")
	util.LogDebugBytes(buf[:n])
}

func (g *GameState) HasSave() bool {
	return firefly.FileExists("save")
}

func (g *GameState) LoadSave() bool {
	file := firefly.LoadFile("save", nil)
	if !file.Exists() {
		return false
	}
	g.Reset()
	if err := g.UnmarshalBinary(file); err != nil {
		firefly.LogError("failed to load save: " + err.Error())
		return false
	}

	var buf [42]byte
	n := copy(buf[0:], "loaded saved game, size: ")
	n += util.FormatIntInto(buf[n:], len(file))
	n += copy(buf[n:], " B")
	firefly.LogDebugBytes(buf[:n])
	return true
}

func (g *GameState) Reset() {
	*g = GameState{InRaceBattle: map[firefly.Peer]Firefly{}}
}
