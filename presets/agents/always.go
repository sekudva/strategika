package agents

import (
	"github.com/sekudva/strategika/internal/domain"
	"github.com/sekudva/strategika/presets/strategies"
)

// ========== 2. ALWAYS FAMILY ==========

func AlwaysShare() *domain.Agent {
	return &domain.Agent{
		Name:     "AlwaysShare",
		ID:       RequestID(20),
		Strategy: strategies.AlwaysShare(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

func AlwaysHold() *domain.Agent {
	return &domain.Agent{
		Name:     "AlwaysHold",
		ID:       RequestID(21),
		Strategy: strategies.AlwaysHold(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

func AlwaysTake() *domain.Agent {
	return &domain.Agent{
		Name:     "AlwaysTake",
		ID:       RequestID(22),
		Strategy: strategies.AlwaysTake(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

func Random() *domain.Agent {
	return &domain.Agent{
		Name:     "Random",
		ID:       RequestID(23),
		Strategy: strategies.Random(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

func RandomOLD() *domain.Agent {
	return &domain.Agent{
		Name:     "RandomOLD",
		ID:       RequestID(24),
		Strategy: strategies.Random_OLD(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// after 2 shares forever share, starts with take
func AntiGrudger() *domain.Agent {
	return &domain.Agent{
		Name:     "AntiGrudger",
		ID:       RequestID(25),
		Strategy: strategies.AntiGrudger(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}
