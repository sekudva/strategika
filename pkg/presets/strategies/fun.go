package strategies

import "github.com/sekudva/strategika/internal/domain"

// crazy and irrational
func Crazy() *domain.Strategy {
	m := domain.MirrorOpp
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix:    domain.Share,
			Mirror: &m,
		},
		Trigger: nil,
	}
}

// never coop
func Unfriendly() *domain.Strategy {
	m := domain.MirrorCold
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix:    domain.Share,
			Mirror: &m,
		},
		Trigger: nil,
	}
}

// zero tolerance to Hold
func Irreconcilable() *domain.Strategy {
	m := domain.MirrorBad
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix:    domain.Share,
			Mirror: &m,
		},
		Trigger: nil,
	}
}

// StrictTeacher
// start Share, for two Takes - one Take and forgives
func StrictTeacher() *domain.Strategy {
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix: domain.Share,
		},

		Trigger: &domain.Trigger{
			Act:   domain.Take,
			Count: 2,
			Mode:  domain.TriggerExactly,
			Reaction: domain.RuleValue{
				Fix: domain.Take,
			},
		},
	}
}

// При первом Take → с вероятностью 2/7 ответить Take, 5/7 Share
func GroFman() *domain.Strategy {
	return &domain.Strategy{
		Neutral: domain.RuleValue{Fix: domain.Share},
		Trigger: &domain.Trigger{
			Act:   domain.Take,
			Count: 1,
			Mode:  domain.TriggerStreakAfter,
			Reaction: domain.RuleValue{
				Fix: domain.Share,
				Prob: map[domain.Act]float64{
					domain.Share: 0.7142857, // 5/7 ≈ 71.4%
					domain.Take:  0.2857143, // 2/7 ≈ 28.6%
				},
			},
		},
	}
}

// Only after 3 Takes do Take
func Patient() *domain.Strategy {
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix: domain.Share,
		},
		Trigger: &domain.Trigger{
			Act:   domain.Take,
			Count: 3,
			Mode:  domain.TriggerStreakAfter,
			Reaction: domain.RuleValue{
				Fix: domain.Take,
			},
		},
	}
}
