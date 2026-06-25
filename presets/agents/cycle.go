package agents

import (
	"github.com/sekudva/strategika/internal/domain"
	"github.com/sekudva/strategika/presets/strategies"
)

// ========== 4. CYCLE FAMILY ==========

// always take, but if two shares - do share
func EvilCyclist() *domain.Agent {
	return &domain.Agent{
		Name:     "EvilCyclist",
		ID:       RequestID(40),
		Strategy: strategies.EvilCyclist(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// answer take only to 2 takes
func GoodCyclist() *domain.Agent {
	return &domain.Agent{
		Name:     "GoodCyclist",
		ID:       RequestID(41),
		Strategy: strategies.GoodCyclist(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// never share, react only if 2 takes
func ColdCyclist() *domain.Agent {
	return &domain.Agent{
		Name:     "ColdCyclist",
		ID:       RequestID(42),
		Strategy: strategies.ColdCyclist(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// likes to share, but every 2 hold do hold, on take do take
func WarmCyclist() *domain.Agent {
	return &domain.Agent{
		Name:     "WarmCyclist",
		ID:       RequestID(43),
		Strategy: strategies.WarmCyclist(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// share if only 2 shares, on take do take
func ColdGhost() *domain.Agent {
	return &domain.Agent{
		Name:     "ColdGhost",
		ID:       RequestID(44),
		Strategy: strategies.ColdGhost(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}
