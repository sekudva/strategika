package tournament

import (
	"github.com/sekudva/strategika/internal/domain"
)

// RoundRobin запускает круговой турнир: каждая стратегия играет с каждой (включая саму с собой)
// Агенты передаются уже созданными. Функция сама чистит их память перед каждой дуэлью,
// но сохраняет накопленный счёт.
func (cfg SimConfig) RoundRobin(agents []*domain.Agent) {
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
			cfg.Pairs = DuelPairs()

			// Запускаем дуэль
			cfg.RunSimulation([]*domain.Agent{agents[i], agents[j]})
		}
	}

	cfg.Logger.Finalize(agents)
}
