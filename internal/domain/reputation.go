package domain

type Reputation struct {
	Cooperation float64
	Defense     float64
}

/*
| Партнёр → Я   | Cooperation | Defense |
| ------------- | ----------- | ------- |
| Share → Share | ++          | ·       |
| Share → Hold  | −           | ·       |
| Share → Take  | −−          | ·       |
| Hold → Share  | ++          | −       |
| Hold → Hold   | −           | +       |
| Hold → Take   | −−          | +       |
| Take → Share  | ·           | −−      |
| Take → Hold   | ·           | −       |
| Take → Take   | ·           | ++      |
*/
// Дельта репутации
var reputationTable = [3][3]Reputation{
	// Share
	{
		{+0.1, 0},
		{-0.05, 0},
		{-0.1, 0},
	},
	// Hold
	{
		{+0.1, -0.05},
		{-0.05, +0.05},
		{-0.1, +0.05},
	},
	// Take
	{
		{0, -0.1},
		{0, -0.05},
		{0, +0.1},
	},
}

func UpdRep(rep *Reputation, myAct, opAct Act) {
	delta := reputationTable[myAct][opAct]
	rep.Cooperation = clamp(rep.Cooperation + delta.Cooperation)
	rep.Defense = clamp(rep.Defense + delta.Defense)
}

func clamp(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}
