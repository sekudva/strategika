package mod

import "github.com/sekudva/strategika/internal/domain"

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
