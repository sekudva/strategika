package strategies

import "github.com/sekudva/strategika/internal/domain"

// Grudger (after Take only Take) orig logic
func Grudger() *domain.Strategy {
	m := domain.MirrorSelf
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix:    domain.Share,
			Mirror: &m,
		},
		Trigger: &domain.Trigger{
			Act:   domain.Take,
			Count: 1,
			Mode:  domain.TriggerExactly,
			Reaction: domain.RuleValue{
				Fix: domain.Take,
			},
		},
	}
}
