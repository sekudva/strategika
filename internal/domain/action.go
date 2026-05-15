package domain

type Act int

const (
	Share Act = iota
	Hold
	Take
)

// для первых ходов
const NoAct Act = -1

func (a Act) String() string {
	switch a {
	case Share:
		return "Share"
	case Hold:
		return "Hold"
	case Take:
		return "Take"
	}
	return "?"
}
