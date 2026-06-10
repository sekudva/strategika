package mod

import "github.com/sekudva/strategika/internal/domain"

// STRATEGY MODIFIERS FOR ARENA AND TRIAL LIST (1/Group and Arena)
// and REPUTATION-BASED STRATEGY MODIFIERS

// SeekWeak - ищет слабых и эксплуатирует
func SeekWeak() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if ctx.OpRep.Coop > 0.6 && ctx.OpRep.Def < 0.35 {
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
		if ctx.OpRep.Coop < 0.2 && ctx.OpRep.Def < 0.2 {
			return domain.Take
		}
		return core
	}
}

// Воин - ищет агрессивных и воюет с ними
func Warrior() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if ctx.OpRep.Coop < 0.15 && ctx.OpRep.Def > 0.75 {
			return domain.Take
		}
		return core
	}
}

// Кооперирует со слабыми
func Assistant() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if ctx.OpRep.Coop > 0.6 && ctx.OpRep.Def < 0.35 {
			return domain.Share
		}
		return core
	}
}

// Кооперирует даже если случайно вошел в цикл зла
func Mature() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if ctx.OpRep.Coop > 0.7 {
			return domain.Share
		}
		return core
	}
}

// Драчун, дерется с сильными
func Brawler() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if ctx.OpRep.Def > 0.9 {
			return domain.Take
		}
		return core
	}
}

// ищет слабых и тянет их на дно
func Quicksand() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if ctx.OpRep.Def < 0.35 {
			return domain.Take
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
		//opPrev := ctx.History.OpLastAct()
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
