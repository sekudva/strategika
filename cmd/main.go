package main

import (
	"fmt"

	"github.com/sekudva/strategika/internal/domain"
	"github.com/sekudva/strategika/internal/tournament"
	"github.com/sekudva/strategika/pkg/presets"
)

func main() {
	// Генератор ID
	gen := domain.NewIDGenerator()
	id1 := gen.Next()
	id2 := gen.Next()

	// Создание агентов
	a1 := domain.NewAgent(presets.AlwaysShare(), id1)
	a2 := domain.NewAgent(presets.AlwaysTake(), id2)

	fmt.Printf("Duel: %s vs %s\n", a1.ID, a2.ID)
	fmt.Println("Strategy:", "AlwaysShare vs AlwaysTake")
	fmt.Println()

	// Симуляция
	rounds := 10
	//noise := 0.0
	score1, score2 := 0, 0

	for r := 1; r <= rounds; r++ {
		act1 := a1.Decide(a2.ID, r)
		act2 := a2.Decide(a1.ID, r)

		s1, s2 := domain.Payoff(act1, act2)
		score1 += s1
		score2 += s2

		a1.Memory.Record(r, a2.ID, act1, act2)
		a2.Memory.Record(r, a1.ID, act2, act1)

		tournament.LogRound(r, a1, a2, act1, act2, s1, s2)
	}

	tournament.LogResult(a1, a2, score1, score2)
}
