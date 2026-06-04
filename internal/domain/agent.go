package domain

type Agent struct {
	ID   AgID
	Name string

	Memory   *Memory
	Strategy *Strategy

	Modifiers []Modifier

	Score int
	Dead  bool
}

func NewAgent(strat *Strategy, id AgID) *Agent {
	return &Agent{
		ID:       id,
		Name:     "Default",
		Strategy: strat,
		Memory:   NewMemory(),
		Score:    0,
	}
}

func (a *Agent) Decide(opID AgID, round int, opRep Reputation) Act {
	// инициализация внутренних счетчиков для каждого агента отдельно
	if _, ok := a.Memory.ModState[opID]; !ok {
		a.Memory.ModState[opID] = make(map[Counter]int)
	}

	core := a.Strategy.CoreDecision(a.Memory.History[opID])
	ctx := ModContext{
		Round:    round,
		History:  a.Memory.History[opID],
		ModState: a.Memory.ModState[opID],
		OpRep:    opRep,
	}
	for _, mod := range a.Modifiers {
		core = mod(core, ctx)
	}
	return core
}

func (a *Agent) ResetMemory() {
	a.Memory = NewMemory()
}

func (a *Agent) UpdRep(myAct, opAct Act) {
	UpdRep(&a.Memory.Rep, myAct, opAct)
}
