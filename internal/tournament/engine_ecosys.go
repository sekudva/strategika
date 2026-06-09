package tournament

import "github.com/sekudva/strategika/internal/domain"

// RunEcosystem — для арены с жизнью и смертью
func (cfg SimConfig) RunEcosystem(agents []*domain.Agent, deathThreshold int) {
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

		cfg.Logger.MarkDead(active, deathThreshold, round)
	}

	cfg.Logger.Finalize(agents, deathThreshold)
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
