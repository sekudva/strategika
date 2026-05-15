package domain

type Agent struct {
	ID        AgID
	Memory    *Memory
	Strategy  *Strategy
	Modifiers []Modifier
	Score     int
}

func NewAgent(strat *Strategy, id AgID) *Agent {
	return &Agent{
		ID:       id,
		Strategy: strat,
		Memory:   NewMemory(),
		Score:    0,
	}
}
