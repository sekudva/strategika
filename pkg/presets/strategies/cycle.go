package strategies

import (
	"github.com/sekudva/strategika/internal/domain"
)

func EvilCyclist() *domain.Strategy {
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix: domain.Take,
		},
		Trigger: &domain.Trigger{
			Act:   domain.Share,
			Count: 2,
			Mode:  domain.TriggerEvery,
			Reaction: domain.RuleValue{
				Fix: domain.Share,
			},
		},
	}
}

func GoodCyclist() *domain.Strategy {
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix: domain.Share,
		},
		Trigger: &domain.Trigger{
			Act:   domain.Take,
			Count: 2,
			Mode:  domain.TriggerEvery,
			Reaction: domain.RuleValue{
				Fix: domain.Take,
			},
		},
	}
}

// scared to friendship
func ColdCyclist() *domain.Strategy {
	m := domain.MirrorCold
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix:    domain.Hold,
			Mirror: &m,
		},
		Trigger: &domain.Trigger{
			Act:   domain.Take,
			Count: 2,
			Mode:  domain.TriggerEvery,
			Reaction: domain.RuleValue{
				Fix: domain.Take,
			},
		},
	}
}

// like to share, every 2 Hold - Hold
func WarmCyclist() *domain.Strategy {
	def := domain.MirrorNice
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix:    domain.Share,
			Mirror: &def,
		},
		Trigger: &domain.Trigger{
			Act:   domain.Hold,
			Count: 2,
			Mode:  domain.TriggerEvery,
			Reaction: domain.RuleValue{
				Fix: domain.Hold,
			},
		},
	}
}

// share if only 2 shares
func ColdGhost() *domain.Strategy {
	cold := domain.MirrorCold
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix:    domain.Hold,
			Mirror: &cold,
		},
		Trigger: &domain.Trigger{
			Act:   domain.Share,
			Count: 2,
			Mode:  domain.TriggerEvery,
			Reaction: domain.RuleValue{
				Fix: domain.Share,
			},
		},
	}
}
