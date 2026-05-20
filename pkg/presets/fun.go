package presets

import "github.com/sekudva/strategika/internal/domain"

// Fear then opponent do Take
func Bully() *domain.Strategy {
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix: domain.Take,
		},
		Trigger: &domain.Trigger{
			Act:   domain.Take,
			Count: 1,
			Reaction: domain.RuleValue{
				Fix: domain.Hold,
			},
		},
		State: make(map[string]int),
	}
}
