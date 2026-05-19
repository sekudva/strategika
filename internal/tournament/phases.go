package tournament

import (
	"math/rand/v2"
	"sync"

	"github.com/sekudva/strategika/internal/domain"
)

// ФАЗЫ ТУРНИРА

// Ядро принятия решения, логика взаимодействия агентов реализуется здесь
func decidePhase(agents []*domain.Agent, pairs []Pair, round int) map[DirectedPair]domain.Act {

	n := len(pairs) * 2 // для направленных решений

	results := make(map[DirectedPair]domain.Act, n)

	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, p := range pairs {
		wg.Add(2) // по одной горутине на каждое направление
		go func(from, to int) {
			defer wg.Done()
			act := agents[from].Decide(agents[to].ID, round)
			mu.Lock()
			results[DirectedPair{from, to}] = act
			mu.Unlock()
		}(p[0], p[1])

		go func(from, to int) {
			defer wg.Done()
			act := agents[from].Decide(agents[to].ID, round)
			mu.Lock()
			results[DirectedPair{from, to}] = act
			mu.Unlock()
		}(p[1], p[0])
	}
	wg.Wait()

	return results
}

// Фаза наложения шума
func noisePhase(decisions map[DirectedPair]domain.Act, noise float64, rng *rand.Rand) map[DirectedPair]domain.Act {
	if noise <= 0 {
		return decisions
	}

	noisy := make(map[DirectedPair]domain.Act, len(decisions))

	for d, act := range decisions {
		if rng.Float64() < noise {
			noisy[d] = randomAct(rng)
		} else {
			noisy[d] = act
		}
	}

	return noisy
}
func randomAct(rng *rand.Rand) domain.Act {
	return domain.Act(rng.IntN(3))
}

// Бухгалтерская фаза, считает выплаты и записывает в память агентов историю
func applyPhase(agents []*domain.Agent, decisions map[DirectedPair]domain.Act, pairs []Pair, round int) {
	for _, p := range pairs {
		i, j := p[0], p[1]
		actItoJ := decisions[DirectedPair{i, j}]
		actJtoI := decisions[DirectedPair{j, i}]

		payoffI, payoffJ := domain.Payoff(actItoJ, actJtoI)

		agents[i].Score += payoffI
		agents[j].Score += payoffJ

		agents[i].Memory.Record(round, agents[j].ID, actItoJ, actJtoI)
		agents[j].Memory.Record(round, agents[i].ID, actJtoI, actItoJ)
	}
}
