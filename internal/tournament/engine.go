package tournament

import (
	"fmt"
	"strings"

	"github.com/sekudva/strategika/internal/domain"
)

// RunSimulation выполняет все раунды согласно конфигурации.
// Возвращает итоговые счета агентов и ошибку (пока nil).
func RunSimulation(cfg SimConfig, agents []*domain.Agent) (map[domain.AgID]int, error) {
	for round := 1; round <= cfg.Rounds; round++ {

		// 1. Фаза решений (параллельно)
		decisions := decidePhase(agents, cfg.Pairs, round)

		// 2. Фаза шума
		decisions = noisePhase(decisions, cfg.Noise, cfg.RNG)

		// 3. Фаза применения
		applyPhase(agents, decisions, cfg.Pairs, round)

	}

	scores := make(map[domain.AgID]int, len(agents))

	for _, a := range agents {
		scores[a.ID] = a.Score
	}

	return scores, nil
}

func LogRound(round int, a1, a2 *domain.Agent, act1, act2 domain.Act, s1, s2 int) {
	fmt.Printf(
		"[%3d] %-6s vs %-6s | %-5s (%-+3d) vs %-5s (%-+3d)\n",
		round,
		a1.ID, a2.ID,
		act1.String(), s1,
		act2.String(), s2,
	)
}

func LogResult(a1, a2 *domain.Agent, score1, score2 int) {
	fmt.Println(strings.Repeat("─", 55))
	fmt.Printf(
		"Result: %s %d — %d %s\n",
		a1.ID, score1, score2, a2.ID,
	)
	fmt.Println(strings.Repeat("─", 55))
}
