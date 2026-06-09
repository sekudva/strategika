package tournament

import (
	"github.com/sekudva/strategika/internal/domain"
)

// RunSimulation выполняет все раунды согласно конфигурации.
// Возвращает итоговые счета агентов и ошибку (пока nil).
func (cfg SimConfig) RunSimulation(agents []*domain.Agent) {
	for round := 1; round <= cfg.Rounds; round++ {

		// 1. Фаза решений (параллельно)
		decisions := decidePhase(agents, cfg.Pairs, round)

		// 2. Фаза шума
		decisions = noisePhase(decisions, cfg.Noise, cfg.RNG)

		// 3. Фаза применения
		applyPhase(agents, decisions, cfg.Pairs, round, cfg.Logger)

		cfg.Logger.Flush()
	}
}
