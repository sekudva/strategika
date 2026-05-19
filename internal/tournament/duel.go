package tournament

// !!!!!!!!!!!

import (
	"fmt"
	"math/rand/v2"

	"github.com/sekudva/strategika/internal/domain"
)

// Duel реализует Tournament для игры 1 на 1.
type Duel struct {
	Rounds int
	Noise  float64
	RNG    *rand.Rand // если nil, будет создан новый
}

// DuelResult содержит итоги дуэли.
type DuelResult struct {
	Agent1, Agent2     domain.AgID
	Score1, Score2     int
	History1, History2 []domain.Round // опционально для Summary
}

func (r DuelResult) Scores() map[domain.AgID]int {
	return map[domain.AgID]int{
		r.Agent1: r.Score1,
		r.Agent2: r.Score2,
	}
}

func (r DuelResult) Summary() string {
	return fmt.Sprintf("%d vs %d: %d - %d", r.Agent1, r.Agent2, r.Score1, r.Score2)
}

// Run запускает дуэль между двумя агентами.
// Агенты должны быть ровно два.
func (d Duel) Run(agents []*domain.Agent) (Result, error) {
	if len(agents) != 2 {
		return nil, fmt.Errorf("duel requires exactly 2 agents, got %d", len(agents))
	}

	rng := d.RNG
	if rng == nil {
		rng = rand.New(rand.NewPCG(0, 0)) // детерминированный seed по умолчанию
	}

	cfg := SimConfig{
		Rounds: d.Rounds,
		Noise:  d.Noise,
		Pairs:  []Pair{{0, 1}}, // единственная пара
		RNG:    rng,
	}

	// Внимание: RunSimulation изменяет Score и Memory агентов,
	// поэтому для чистоты можно предварительно сбросить их до начального состояния.
	// Но это ответственность вызывающего (создать свежих агентов через presets).

	scores, err := RunSimulation(cfg, agents)
	if err != nil {
		return nil, err
	}

	return DuelResult{
		Agent1: agents[0].ID,
		Agent2: agents[1].ID,
		Score1: scores[agents[0].ID],
		Score2: scores[agents[1].ID],
	}, nil
}
