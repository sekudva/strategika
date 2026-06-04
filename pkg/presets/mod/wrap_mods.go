package mod

import (
	"math/rand/v2"

	"github.com/sekudva/strategika/internal/domain"
)

// WRAP MODS - вспомогательные надстройки над модификаторами

// Обертка для любого модификатора который смотрит на предыдущее действие оппонента
// Игнорирование применения модификатора если оппонент сделал Hold
func IgnoreHold(next domain.Modifier) domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if len(ctx.History) > 0 && (ctx.History.OpLastAct() == domain.Hold && ctx.History.Op2LastAct() == domain.Hold) {
			return core
		}
		return next(core, ctx)
	}
}

// Применяет модификатор с веротяностью
func WithProbability(prob float64, next domain.Modifier) domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if rand.Float64() < prob {
			return next(core, ctx)
		}
		return core
	}
}
