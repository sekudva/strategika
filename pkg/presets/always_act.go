package presets

import (
	"github.com/sekudva/strategika/internal/domain"
)

// always share
func AlwaysShare() *domain.Strategy {
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix: domain.Share,
		},
		Trigger: nil,
		State:   make(map[string]int),
	}
}

// always hold
func AlwaysHold() *domain.Strategy {
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix: domain.Hold,
		},
		Trigger: nil,
		State:   make(map[string]int),
	}
}

// always take
func AlwaysTake() *domain.Strategy {
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix: domain.Take,
		},
		Trigger: nil,
		State:   make(map[string]int),
	}
}
