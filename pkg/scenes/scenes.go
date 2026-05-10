package scenes

var SwitchScene func(scene Scene)

const (
	LongestSceneName = 15
)

var AllScenes = []Scene{
	MainMenu,
	Field,
	Insectarium,
	Shop,
	RacingBattle,
	RacingTraining,

	RacingBattle2,
}

type Scene byte

const (
	MainMenu Scene = iota
	Field
	Insectarium
	Shop
	RacingBattle
	RacingTraining

	// New racebattle design
	// Will replace the old racebattle eventually
	RacingBattle2
)

func (s Scene) String() string {
	switch s {
	case Insectarium:
		return "insectarium"
	case Field:
		return "field"
	case MainMenu:
		return "main menu"
	case RacingBattle:
		return "racing battle"
	case RacingTraining:
		return "racing training"
	case Shop:
		return "shop"
	case RacingBattle2:
		return "racing battle (new)"
	default:
		panic("unexpected Scene")
	}
}
