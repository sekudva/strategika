package tournament

import (
	"fmt"
	"math/rand/v2"
)

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

// валидация значений
func (cfg SimConfig) Validate() error {
	if cfg.Rounds <= 0 {
		return fmt.Errorf("rounds must be > 0, got %d", cfg.Rounds)
	}

	if cfg.Noise < 0 || cfg.Noise > 1 {
		return fmt.Errorf("noise must be between 0 and 1, got %f", cfg.Noise)
	}

	if len(cfg.Pairs) == 0 {
		return fmt.Errorf("no pairs defined")
	}

	return nil
}

// заполнение незаполненных полей
func (cfg SimConfig) DefVars() SimConfig {
	if cfg.RNG == nil {
		cfg.RNG = rand.New(rand.NewPCG(0, 0))
	}
	if cfg.Logger == nil {
		cfg.Logger = &SilentLogger{}
	}
	return cfg
}
