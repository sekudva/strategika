package agents

import (
	"github.com/sekudva/strategika/internal/domain"
	"github.com/sekudva/strategika/presets/strategies"
)

// ========== 1. TIT-FOR-TAT FAMILY ==========

// just copy
func TitForTat() *domain.Agent {
	return &domain.Agent{
		Name:     "TitForTat",
		ID:       RequestID(10),
		Strategy: strategies.TitForTatNEW(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// triggers only on take
func TitForTatOrig() *domain.Agent {
	return &domain.Agent{
		Name:     "TitForTatOrig",
		ID:       RequestID(11),
		Strategy: strategies.TitForTat(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// triggers on 2 takes AND to 1 take - do hold
func TitFor2Tats() *domain.Agent {
	return &domain.Agent{
		Name:     "TitFor2Tats",
		ID:       RequestID(12),
		Strategy: strategies.TitFor2Tats(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// triggers only on 2 takes
func TitFor2TatsOrig() *domain.Agent {
	return &domain.Agent{
		Name:     "TitFor2TatsOLD",
		ID:       RequestID(13),
		Strategy: strategies.TitFor2TatsOLD(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// not always triggers on take
func ForgivingTFT() *domain.Agent {
	return &domain.Agent{
		Name:     "ForgivingTFT",
		ID:       RequestID(12),
		Strategy: strategies.ForgivingTitForTat(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// start with take
func EvilTFT() *domain.Agent {
	return &domain.Agent{
		Name:     "EvilTFT",
		ID:       RequestID(13),
		Strategy: strategies.TitForTatEVIL(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// take prob
func Joss() *domain.Agent {
	return &domain.Agent{
		Name:     "Joss",
		ID:       RequestID(14),
		Strategy: strategies.Joss(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// take prob, take start
func HardJoss() *domain.Agent {
	return &domain.Agent{
		Name:     "HardJoss",
		ID:       RequestID(15),
		Strategy: strategies.HardJoss(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}
