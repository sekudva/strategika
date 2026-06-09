package tournament

import (
	"math"

	"github.com/sekudva/strategika/internal/domain"
)

// RoundRobin запускает круговой турнир: каждая стратегия играет с каждой (включая саму с собой)
// Агенты передаются уже созданными. Функция сама чистит их память перед каждой дуэлью,
// но сохраняет накопленный счёт.
func RoundRobin(cfg SimConfig, agents []*domain.Agent) error {
	n := len(agents)

	for _, a := range agents {
		a.Score = 0
	}

	for i := range n {
		for j := i; j < n; j++ { // j = i включает само-игры
			// Очищаем память агентов перед дуэлью (счёт сохраняется)
			agents[i].ResetMemory()
			agents[j].ResetMemory()

			// Настраиваем пары для этой дуэли
			duelCfg := cfg
			duelCfg.Pairs = []Pair{{0, 1}}

			// log
			duelCfg.Logger = NewAggregateLogger(
				cfg.Logger.(*AggregateLogger).Interval,
				duelCfg.Pairs,
				[]*domain.Agent{agents[i], agents[j]},
				cfg.Logger.(*AggregateLogger).Writer,
			)

			// Запускаем дуэль
			duelCfg.RunSimulation([]*domain.Agent{agents[i], agents[j]})
		}
	}

	cfg.Logger.Finalize(agents, math.MinInt)
	return nil
}
