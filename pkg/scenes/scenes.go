package scenes

var SwitchScene func(scene Scene)

type Scene byte

const (
	MainMenu Scene = iota
	Field
	Insectarium
	Shop
	RaceBattle
)

func (s Scene) String() string {
	switch s {
	case Insectarium:
		return "insectarium"
	case Field:
		return "field"
	case MainMenu:
		return "main menu"
	case RaceBattle:
		return "race battle"
	case Shop:
		return "shop"
	default:
		panic("unexpected Scene")
	}
}
