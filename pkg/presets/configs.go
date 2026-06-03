package presets

import (
	"fmt"
	"math"
	"math/rand/v2"

	"github.com/sekudva/strategika/internal/tournament"
)

// DuelConfig возвращает конфигурацию для дуэли двух агентов
func DuelConfig(rounds int, noise float64) (tournament.SimConfig, error) {
	cfg := tournament.SimConfig{
		Rounds:         rounds,
		Noise:          noise,
		Pairs:          tournament.DuelPairs(),
		RNG:            DefaultRNG(),
		Logger:         &tournament.AllLogger{},
		DeathThreshold: math.MinInt,
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
		Rounds:         rounds,
		Noise:          noise,
		Pairs:          tournament.AllPairs(agentCount),
		RNG:            DefaultRNG(),
		Logger:         &tournament.AggregateLogger{Interval: 10},
		DeathThreshold: math.MinInt,
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
		Rounds:         rounds,
		Noise:          noise,
		Pairs:          tournament.TrialPairs(leaderIndex, agentCount),
		RNG:            DefaultRNG(),
		Logger:         &tournament.AllLogger{},
		DeathThreshold: math.MinInt,
	}

	if err := cfg.Validate(); err != nil {
		return tournament.SimConfig{}, err
	}

	return cfg.DefVars(), nil

}

// ------------------

func DefaultRNG() *rand.Rand {
	return rand.New(rand.NewPCG(0, 0))
}
