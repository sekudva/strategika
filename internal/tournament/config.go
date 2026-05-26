package tournament

import "math/rand/v2"

// Типы для хранения индексов в срезе []*domain.Agent
// DirectedPair — направленная пара агентов
type DirectedPair struct{ From, To int }

// Pair — неупорядоченная пара агентов
type Pair [2]int

// Конфигурация турнира
type SimConfig struct {
	Rounds int
	Noise  float64
	Pairs  []Pair     // список взаимодействий на каждый раунд
	RNG    *rand.Rand // воспроизводимость
	Logger RoundLogger
}

func DuelConfig(rounds int, noise float64) SimConfig {
	return SimConfig{
		Rounds: rounds,
		Noise:  noise,
		Pairs:  []Pair{{0, 1}},
		RNG:    rand.New(rand.NewPCG(0, 0)),
		Logger: &AllLogger{},
	}
}

func ArenaConfig(agentCount int, rounds int, noise float64) SimConfig {
	pairs := make([]Pair, 0, agentCount*(agentCount-1)/2)
	for i := range agentCount {
		for j := i + 1; j < agentCount; j++ {
			pairs = append(pairs, Pair{i, j})
		}
	}
	return SimConfig{
		Rounds: rounds,
		Noise:  noise,
		Pairs:  pairs,
		RNG:    rand.New(rand.NewPCG(0, 0)),
		Logger: &AggregateLogger{Interval: 10},
	}
}
