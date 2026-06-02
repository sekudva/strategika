package domain

// STRATEGY MODIFIERS LIST (1/1 duels)
// Offended (Обижающаяся) - имеет увеличивающийся счетчик (сколько нужно сделать Share чтобы снова начать делать Share)

type ModContext struct {
	History  Rounds
	Round    int
	Strategy *Strategy
	OpRep    Reputation
}

type Modifier func(core Act, ctx ModContext) Act
