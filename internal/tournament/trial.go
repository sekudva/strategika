package tournament

import (
	"math"

	"github.com/sekudva/strategika/internal/domain"
)

func RunTrial(cfg SimConfig, leader *domain.Agent, group []*domain.Agent) {
	agents := append([]*domain.Agent{leader}, group...)

	cfg.Pairs = TrialPairs(0, len(agents))

	cfg.RunSimulation(agents)
}

// Circulaire запускает турнир испытаний: каждый лидер проходит испытание группой.
// Группа не меняется между лидерами. Для каждого лидера вызывается RunTrial.
// Результаты всех испытаний выводятся через cfg.Logger.
func Circulaire(cfg SimConfig, leaders []*domain.Agent, group []*domain.Agent) error {
	for _, leader := range leaders {
		for _, g := range group {
			g.ResetMemory()
		}

		agents := append([]*domain.Agent{leader}, group...)

		pairs := make([]Pair, len(group))
		for i := range group {
			pairs[i] = Pair{0, i + 1}
		}

		trialCfg := cfg
		trialCfg.Pairs = pairs
		trialCfg.Logger = NewAggregateLogger(
			cfg.Logger.(*AggregateLogger).Interval,
			pairs,
			agents,
			cfg.Logger.(*AggregateLogger).Writer,
		)

		trialCfg.RunSimulation(agents)
	}

	cfg.Logger.Finalize(leaders, math.MinInt)
	return nil
}
