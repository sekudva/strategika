package mod

import (
	"math/rand/v2"

	"github.com/sekudva/strategika/internal/domain"
)

// Модификатор для стратегии Туллока, накопительная стратегия
// Снижение кооперации по накоплению предательств
func TullockMod() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if ctx.Round <= 11 {
			return domain.Share
		}

		// Смотрим последние 10 ходов оппонента
		start := max(0, len(ctx.History)-10)
		coopCount := 0
		for i := start; i < len(ctx.History); i++ {
			if ctx.History[i].OpAct == domain.Share {
				coopCount++
			}
		}

		coopRate := float64(coopCount) / 10.0

		// Если 9 из 10 коопераций — щедрый шаг
		if coopCount >= 9 {
			if rand.Float64() < 0.8 {
				return domain.Share
			}
			return domain.Take
		}

		// Иначе: кооперируем на (coopRate - 0.1)%
		prob := max(0, coopRate-0.1)
		if rand.Float64() < prob {
			return domain.Share
		}
		return domain.Take
	}
}
