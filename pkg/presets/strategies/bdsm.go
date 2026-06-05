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
			Mode:  domain.TriggerStreakAfter,
			Reaction: domain.RuleValue{
				Fix: domain.Share,
			},
		},
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
			Mode:  domain.TriggerStreakAfter,
			Reaction: domain.RuleValue{
				Fix: domain.Take,
			},
		},
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
	}
}

// dont mind hold, on take do take
func Saint() *domain.Strategy {
	saint := domain.MirrorSaint
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix:    domain.Share,
			Mirror: &saint,
		},
		Trigger: nil,
	}
}

// Fear then opponent do Take
func Bully() *domain.Strategy {
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix: domain.Take,
		},
		Trigger: &domain.Trigger{
			Act: domain.Take,
			Reaction: domain.RuleValue{
				Fix: domain.Hold,
			},
		},
	}
}

// Always Take, but if Share - Share
func Melting() *domain.Strategy {
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix: domain.Take,
		},
		Trigger: &domain.Trigger{
			Act: domain.Share,
			Reaction: domain.RuleValue{
				Fix: domain.Share,
			},
		},
	}
}
