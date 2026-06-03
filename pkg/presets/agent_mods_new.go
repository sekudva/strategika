package presets

import (
	"github.com/sekudva/strategika/internal/domain"
	"github.com/sekudva/strategika/pkg/presets/mod"
)

// ========== 10. PREDICTABLE BEHAVIOR ==========

func Journalist() *domain.Agent {
	return &domain.Agent{
		Name: "Journalist",
		ID:   RequestID(100),

		Strategy: &domain.Strategy{
			Neutral: domain.RuleValue{
				Fix:  domain.Hold,
				Prob: map[domain.Act]float64{domain.Take: 0.1},
			},
		},

		Memory:    domain.NewMemory(),
		Score:     0,
		Modifiers: []domain.Modifier{mod.JournalistMod()},
	}
}
