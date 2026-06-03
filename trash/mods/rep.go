package mods

import "github.com/sekudva/strategika/internal/domain"

// STRATEGY MODIFIERS FOR ARENA AND TRIAL LIST (1/Group and Arena)
// and REPUTATION-BASED STRATEGY MODIFIERS

// SeekWeak - ищет слабых и эксплуатирует
func SeekWeak() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if ctx.OpRep.Coop > 0.65 && ctx.OpRep.Def < 0.3 {
			return domain.Take
		}
		return core
	}
}

// SeekFreak - ищет недружелюбных защищающихся и не взаимодействует
func SeekFreak() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if ctx.OpRep.Coop < 0.3 && ctx.OpRep.Def > 0.65 {
			return domain.Hold
		}
		return core
	}
}

// SeekSneak - ищет пассивных
func SeekSneak() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if ctx.OpRep.Coop < 0.3 && ctx.OpRep.Def < 0.3 {
			return domain.Take
		}
		return core
	}
}

// Doubler
func Doubler() domain.Modifier {
	repeatCount := 0
	var lastAct domain.Act

	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if len(ctx.History) == 0 {
			return core
		}

		currentLast := ctx.History.OpLastAct()

		if currentLast != lastAct {
			lastAct = currentLast
			repeatCount = 0
		}

		repeatCount++
		if repeatCount <= 2 {
			return lastAct
		}

		return core
	}
}
