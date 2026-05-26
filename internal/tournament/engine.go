package tournament

import (
	"fmt"
	"math/rand/v2"

	"github.com/sekudva/strategika/internal/domain"
)

// Run запускает симуляцию и проверяет параметры конфигурации
func (cfg SimConfig) Run(agents []*domain.Agent) (map[domain.AgID]int, error) {

	if err := cfg.validate(len(agents)); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	cfg = cfg.defVars()

	scores := cfg.runSimulation(agents)

	return scores, nil
}

// RunSimulation выполняет все раунды согласно конфигурации.
// Возвращает итоговые счета агентов и ошибку (пока nil).
func (cfg SimConfig) runSimulation(agents []*domain.Agent) map[domain.AgID]int {
	for round := 1; round <= cfg.Rounds; round++ {

		// 1. Фаза решений (параллельно)
		decisions := decidePhase(agents, cfg.Pairs, round)

		// 2. Фаза шума
		decisions = noisePhase(decisions, cfg.Noise, cfg.RNG)

		// 3. Фаза применения
		applyPhase(agents, decisions, cfg.Pairs, round, cfg.Logger)

	}

	scores := make(map[domain.AgID]int, len(agents))

	for _, a := range agents {
		scores[a.ID] = a.Score
	}

	return scores
}

// валидация значений
func (cfg SimConfig) validate(agentCount int) error {
	if agentCount <= 0 {
		return fmt.Errorf("no agents provided")
	}

	if cfg.Rounds <= 0 {
		return fmt.Errorf("rounds must be > 0, got %d", cfg.Rounds)
	}

	if cfg.Noise < 0 || cfg.Noise > 1 {
		return fmt.Errorf("noise must be between 0 and 1, got %f", cfg.Noise)
	}

	if len(cfg.Pairs) == 0 {
		return fmt.Errorf("no pairs defined")
	}

	for _, p := range cfg.Pairs {

		if p[0] >= agentCount || p[1] >= agentCount {
			return fmt.Errorf("pair (%d,%d) references non-existent agent (max index %d)", p[0], p[1], agentCount-1)
		}

		if p[0] == p[1] {
			return fmt.Errorf("pair (%d,%d) is self-play, not allowed", p[0], p[1])
		}
	}

	return nil
}

// заполнение незаполненных полей
func (cfg SimConfig) defVars() SimConfig {
	if cfg.RNG == nil {
		cfg.RNG = rand.New(rand.NewPCG(0, 0))
	}
	if cfg.Logger == nil {
		cfg.Logger = &SilentLogger{}
	}
	return cfg
}
