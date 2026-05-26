package presets

import (
	"fmt"
	"math/rand/v2"

	"github.com/sekudva/strategika/internal/tournament"
)

// DuelConfig возвращает конфигурацию для дуэли двух агентов
func DuelConfig(rounds int, noise float64) (tournament.SimConfig, error) {
	cfg := tournament.SimConfig{
		Rounds: rounds,
		Noise:  noise,
		Pairs:  DuelPairs(),
		RNG:    DefaultRNG(),
		Logger: &tournament.AllLogger{},
	}

	if err := cfg.Validate(); err != nil {
		return tournament.SimConfig{}, err
	}

	return cfg.DefVars(), nil
}

// ArenaConfig возвращает конфигурацию для арены (все против всех)
func ArenaConfig(agentCount int, rounds int, noise float64) (tournament.SimConfig, error) {
	if agentCount < 2 {
		return tournament.SimConfig{}, fmt.Errorf("arena requires at least 2 agents, got %d", agentCount)
	}

	cfg := tournament.SimConfig{
		Rounds: rounds,
		Noise:  noise,
		Pairs:  AllPairs(agentCount),
		RNG:    DefaultRNG(),
		Logger: &tournament.AggregateLogger{Interval: 10},
	}

	if err := cfg.Validate(); err != nil {
		return tournament.SimConfig{}, err
	}

	return cfg.DefVars(), nil
}

// TrialConfig возвращает конфигурацию для испытания (один против группы)
func TrialConfig(leaderIndex, agentCount int, rounds int, noise float64) (tournament.SimConfig, error) {
	if agentCount < 2 {
		return tournament.SimConfig{}, fmt.Errorf("trial requires at least 2 agents, got %d", agentCount)
	}

	if leaderIndex < 0 || leaderIndex >= agentCount {
		return tournament.SimConfig{}, fmt.Errorf("leader index %d out of range [0,%d]", leaderIndex, agentCount-1)
	}

	cfg := tournament.SimConfig{
		Rounds: rounds,
		Noise:  noise,
		Pairs:  TrialPairs(leaderIndex, agentCount),
		RNG:    DefaultRNG(),
		Logger: &tournament.AllLogger{},
	}

	if err := cfg.Validate(); err != nil {
		return tournament.SimConfig{}, err
	}

	return cfg.DefVars(), nil

}

// ------------------
func AllPairs(n int) []tournament.Pair {
	pairs := make([]tournament.Pair, 0, n*(n-1)/2)
	for i := range n {
		for j := i + 1; j < n; j++ {
			pairs = append(pairs, tournament.Pair{i, j})
		}
	}
	return pairs
}

// DuelPairs возвращает одну пару для дуэли
func DuelPairs() []tournament.Pair {
	return []tournament.Pair{{0, 1}}
}

// TrialPairs возвращает пары лидера со всеми остальными
func TrialPairs(leaderIndex, agentCount int) []tournament.Pair {
	pairs := make([]tournament.Pair, 0, agentCount-1)
	for i := range agentCount {
		if i != leaderIndex {
			pairs = append(pairs, tournament.Pair{leaderIndex, i})
		}
	}
	return pairs
}

func DefaultRNG() *rand.Rand {
	return rand.New(rand.NewPCG(0, 0))
}
