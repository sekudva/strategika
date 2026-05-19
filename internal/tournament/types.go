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
}
