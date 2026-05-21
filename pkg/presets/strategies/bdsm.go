package strategies

import "github.com/sekudva/strategika/internal/domain"

// masochist
func Masochist() *domain.Strategy {
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix: domain.Hold,
		},

		Trigger: &domain.Trigger{
			Act:   domain.Take,
			Count: 1,
			Reaction: domain.RuleValue{
				Fix: domain.Share,
			},
		},
		State: make(map[string]int),
	}
}

// sadist
func Sadist() *domain.Strategy {
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix: domain.Hold,

			Prob: map[domain.Act]float64{
				domain.Take: 0.4,
			},
		},

		Trigger: &domain.Trigger{
			Act:   domain.Share,
			Count: 1,
			Reaction: domain.RuleValue{
				Fix: domain.Take,
			},
		},
		State: make(map[string]int),
	}
}

// pacifist, never Take
func Pacifist() *domain.Strategy {
	def := domain.MirrorDefense
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix:    domain.Share,
			Mirror: &def,
		},
		Trigger: nil,
		State:   make(map[string]int),
	}
}

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
