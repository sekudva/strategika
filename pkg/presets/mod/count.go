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

func Exploiter() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		n := len(ctx.History)
		if n == 0 {
			ctx.ModState[domain.PunishCounter] = 0 // сколько раз мой Take наказали
			ctx.ModState[domain.SafeCounter] = 0   // сколько раз мой Take прошёл
			return core
		}

		if n >= 1 {
			myPrev := ctx.History.My2LastAct()
			op := ctx.History.OpLastAct()

			if myPrev != domain.Share {
				if op == domain.Take {
					ctx.ModState[domain.PunishCounter]++
				} else {
					ctx.ModState[domain.SafeCounter]++
				}
			}

		}

		if n <= 10 {
			return core
		}

		myPrev := ctx.History.My2LastAct()
		my := ctx.History.MyLastAct()
		op := ctx.History.OpLastAct()

		punish, safe := ctx.ModState[domain.PunishCounter], ctx.ModState[domain.SafeCounter]

		switch myPrev {
		case domain.Take:
			if op != domain.Take || safe > punish {
				return domain.Take
			}
			return domain.Hold

		case domain.Hold:
			if op != domain.Take {
				return domain.Share
			}
			return domain.Hold

		case domain.Share:
			if op != domain.Take && my != domain.Take {
				return domain.Take
			} else if punish > (safe*4) || my == domain.Take {
				return domain.Hold
			}

			return domain.Share

		default:
			return core
		}

	}
}
