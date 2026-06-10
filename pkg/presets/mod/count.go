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

func GoByMajorityMod() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if len(ctx.History) == 0 {
			return core
		}

		var shareCount, holdCount, takeCount int
		for _, round := range ctx.History {
			switch round.OpAct {
			case domain.Share:
				shareCount++
			case domain.Hold:
				holdCount++
			case domain.Take:
				takeCount++
			}
		}

		// Большинство Share → Share
		if shareCount > holdCount && shareCount > takeCount {
			return domain.Share
		}
		// Большинство Take → Take
		if takeCount > shareCount && takeCount > holdCount {
			return domain.Take
		}
		// Большинство Hold → Hold
		if holdCount >= shareCount && holdCount >= takeCount {
			return domain.Hold
		}

		return core
	}
}
