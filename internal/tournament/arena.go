package tournament

import (
	"fmt"

	"github.com/sekudva/strategika/internal/domain"
)

///// !!!!!!!

type Arena struct {
	Agents []*domain.Agent
}

type ArenaResult struct {
	Scores map[domain.AgID]int
	Rounds int
}

func (a *Arena) Run(rounds int, noise float64) *ArenaResult {
	//
	return nil
}

func (r ArenaResult) Summary() string {
	return fmt.Sprintf("Arena: %d agents, %d rounds", len(r.Scores), r.Rounds)
}
