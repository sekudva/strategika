package presets

import (
	"github.com/sekudva/strategika/internal/domain"
	"github.com/sekudva/strategika/pkg/presets/mod"
	"github.com/sekudva/strategika/pkg/presets/strategies"
)

// ========== 13. ORIGINAL TOURNAMENT ==========

// GoByMajority - K31R by PAULA GAIL GRISELL.
func GoByMajority() *domain.Agent {
	return &domain.Agent{
		Name:      "GoByMajority",
		ID:        RequestID(130),
		Strategy:  strategies.AlwaysShare(),
		Memory:    domain.NewMemory(),
		Score:     0,
		Modifiers: []domain.Modifier{mod.GoByMajorityMod()},
	}
}

// Grudger - K34R by JAMES W. FRIEDMAN.
func Grudger() *domain.Agent {
	return &domain.Agent{
		Name:     "Grudger",
		ID:       RequestID(131),
		Strategy: strategies.Grudger(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// WMAdams - K44R by WM. ADAMS.
func WMAdams() *domain.Agent {
	return &domain.Agent{
		Name:      "WMAdams",
		ID:        RequestID(132),
		Strategy:  strategies.AlwaysShare(),
		Memory:    domain.NewMemory(),
		Score:     0,
		Modifiers: []domain.Modifier{mod.AdamsMod()},
	}
}

// Eatherley - K46R by GRAHAM J. EATHERLEY
func Eatherley() *domain.Agent {
	return &domain.Agent{
		Name:      "Eatherley",
		ID:        RequestID(133),
		Strategy:  strategies.AlwaysShare(),
		Memory:    domain.NewMemory(),
		Score:     0,
		Modifiers: []domain.Modifier{mod.EatherleyMod()},
	}
}

// Cave - K49R by ROB CAVE.
func Cave() *domain.Agent {
	return &domain.Agent{
		Name:      "Cave",
		ID:        RequestID(133),
		Strategy:  strategies.AlwaysShare(),
		Memory:    domain.NewMemory(),
		Score:     0,
		Modifiers: []domain.Modifier{mod.CaveMod()},
	}
}
