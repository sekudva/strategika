package strategies

import "github.com/sekudva/strategika/internal/domain"

// original logic tit-for-tat: with only share/take
func TitForTat() *domain.Strategy {
	nice := domain.MirrorNice
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix:    domain.Share,
			Mirror: &nice,
		},
		Trigger: nil,
	}
}

// tit-for-two-tats NEW logic (If 1 Take - Hold, 2 Take - Take)
func TitFor2Tats() *domain.Strategy {
	m := domain.MirrorDefense
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix:    domain.Share,
			Mirror: &m,
		},
		Trigger: &domain.Trigger{
			Act:   domain.Take,
			Count: 2,
			Mode:  domain.TriggerStreakAfter,
			Reaction: domain.RuleValue{
				Fix: domain.Take,
			},
		},
	}
}

// tit-for-two-tats OLD logic
func TitFor2TatsOLD() *domain.Strategy {
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix: domain.Share,
		},
		Trigger: &domain.Trigger{
			Act:   domain.Take,
			Count: 2,
			Mode:  domain.TriggerStreakAfter,
			Reaction: domain.RuleValue{
				Fix: domain.Take,
			},
		},
	}
}

// new logic tit-for-tat: exact Mirror with Hold
func TitForTatNEW() *domain.Strategy {
	direct := domain.MirrorDirect
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix:    domain.Share,
			Mirror: &direct,
		},
		Trigger: nil,
	}
}

// Evil tit-for-tat: exact Mirror, first round Take
func TitForTatEVIL() *domain.Strategy {
	direct := domain.MirrorDirect
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix:    domain.Take,
			Mirror: &direct,
		},
		Trigger: nil,
	}
}

// Probalistic tit-for-tat, NEW logic, Joss
func Joss() *domain.Strategy {
	direct := domain.MirrorDirect
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix:    domain.Share,
			Mirror: &direct,

			Prob: map[domain.Act]float64{
				domain.Take: 0.1,
			},
		},
	}
}

// Probalistic tit-for-tat, NEW logic, Joss, starts with Take
func HardJoss() *domain.Strategy {
	direct := domain.MirrorDirect
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix:    domain.Take,
			Mirror: &direct,

			Prob: map[domain.Act]float64{
				domain.Take: 0.1,
			},
		},
	}
}

// 30% to not triggered on Take
func ForgivingTitForTat() *domain.Strategy {
	direct := domain.MirrorDirect
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix:    domain.Share,
			Mirror: &direct,
		},
		Trigger: &domain.Trigger{
			Act:   domain.Take,
			Count: 1,
			Mode:  domain.TriggerStreakAfter,
			Reaction: domain.RuleValue{
				Fix: domain.Take,
				Prob: map[domain.Act]float64{
					domain.Share: 0.2,
					domain.Hold:  0.1,
				},
			},
		},
	}
}
