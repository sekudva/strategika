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

// ПРАВИЛА ДЛЯ МОДИФИКАТОРОВ: И или ИЛИ
// Возвращает последнее
func All(mods ...Modifier) Modifier {
	return func(core Act, ctx ModContext) Act {
		result := core
		for _, m := range mods {
			result = m(core, ctx)
			if result == core {
				return core
			}
		}
		return result
	}
}

// Возвращает первое которое меняет основной ответ
func Any(mods ...Modifier) Modifier {
	return func(core Act, ctx ModContext) Act {
		for _, m := range mods {
			result := m(core, ctx)
			if result != core {
				return result
			}
		}
		return core
	}
}

func Greater(a, b float64) bool { return a > b }
func Less(a, b float64) bool    { return a < b }
