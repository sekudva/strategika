package mod

import (
	"github.com/sekudva/strategika/internal/domain"
)

// Doubler
func Doubler() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if len(ctx.History) == 0 {
			ctx.ModState[domain.RepeatCounter] = 0
			return core
		}

		if ctx.ModState[domain.RepeatCounter] == 0 {
			ctx.ModState[domain.ActRecorder] = int(ctx.History.OpLastAct())
		}

		ctx.ModState[domain.RepeatCounter]++

		if ctx.ModState[domain.RepeatCounter] <= 2 {
			return domain.Act(ctx.ModState[domain.ActRecorder])
		}

		ctx.ModState[domain.ActRecorder] = int(ctx.History.OpLastAct())
		ctx.ModState[domain.RepeatCounter] = 1
		return domain.Act(ctx.ModState[domain.ActRecorder])
	}
}

// Дважды повторяет только Share и Take
func JournalistMod() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if len(ctx.History) == 0 {
			ctx.ModState[domain.RepeatCounter] = 0
			return core
		}

		if ctx.ModState[domain.RepeatCounter] == 0 {
			ctx.ModState[domain.ActRecorder] = int(ctx.History.OpLastAct())
		}

		ctx.ModState[domain.RepeatCounter]++

		if ctx.ModState[domain.RepeatCounter] <= 2 {
			if domain.Act(ctx.ModState[domain.ActRecorder]) != domain.Hold {
				return domain.Act(ctx.ModState[domain.ActRecorder])
			}
		}

		ctx.ModState[domain.ActRecorder] = int(ctx.History.OpLastAct())

		if domain.Act(ctx.ModState[domain.ActRecorder]) != domain.Hold {
			ctx.ModState[domain.RepeatCounter] = 1
			return domain.Act(ctx.ModState[domain.ActRecorder])
		}

		return core
	}
}
