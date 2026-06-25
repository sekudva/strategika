package tournament

import (
	"github.com/sekudva/strategika/internal/domain"
)

func (cfg SimConfig) RunTrial(leader *domain.Agent, group []*domain.Agent) {
	agents := append([]*domain.Agent{leader}, group...)

	cfg.Pairs = TrialPairs(0, len(agents))

	cfg.RunSimulation(agents)
}

// Circulaire запускает турнир испытаний: каждый лидер проходит испытание группой.
// Группа не меняется между лидерами. Для каждого лидера вызывается RunTrial.
func (cfg SimConfig) Circulaire(leaders []*domain.Agent, group []*domain.Agent) {
	for _, leader := range leaders {
		for _, g := range group {
			g.ResetMemory()
		}
		// Создаём новый логгер для этого лидера
		all := append([]*domain.Agent{leader}, group...)
		pairs := TrialPairs(0, len(all))
		cfg.Logger = cfg.Logger.ForDuel(pairs, all)

		cfg.RunTrial(leader, group)
	}

	cfg.Logger.Finalize(leaders)
}
