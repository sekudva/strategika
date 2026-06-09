package tournament

import "github.com/sekudva/strategika/internal/domain"

// RunEcosystem — для арены с жизнью и смертью
func (cfg SimConfig) RunEcosystem(agents []*domain.Agent) map[domain.AgID]int {
	active := make([]*domain.Agent, len(agents))
	copy(active, agents)

	for round := 1; round <= cfg.Rounds; round++ {
		active = filterAlive(active)
		if len(active) < 2 {
			break
		}

		pairs := allPairs(active)
		decisions := decidePhase(active, pairs, round)
		decisions = noisePhase(decisions, cfg.Noise, cfg.RNG)
		applyPhase(active, decisions, pairs, round, cfg.Logger)

		cfg.Logger.MarkDead(active, cfg.DeathThreshold, round)
	}

	cfg.Logger.Finalize(agents)

	return collectScores(agents)
}

// Вспомогательные (приватные)
func collectScores(agents []*domain.Agent) map[domain.AgID]int {
	scores := make(map[domain.AgID]int, len(agents))
	for _, a := range agents {
		scores[a.ID] = a.Score
	}
	return scores
}

func filterAlive(agents []*domain.Agent) []*domain.Agent {
	alive := make([]*domain.Agent, 0, len(agents))
	for _, a := range agents {
		if !a.Dead {
			alive = append(alive, a)
		}
	}
	return alive
}

func allPairs(agents []*domain.Agent) []Pair {
	pairs := make([]Pair, 0, len(agents)*(len(agents)-1)/2)
	for i := range agents {
		for j := i + 1; j < len(agents); j++ {
			pairs = append(pairs, Pair{i, j})
		}
	}
	return pairs
}

func markDead(agents []*domain.Agent, threshold int) {
	for _, a := range agents {
		if a.Score <= threshold {
			a.Dead = true
		}
	}
}
