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
		State:   make(map[string]int),
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
			Reaction: domain.RuleValue{
				Fix: domain.Take,
			},
		},
		State: make(map[string]int),
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
			Reaction: domain.RuleValue{
				Fix: domain.Take,
			},
		},
		State: make(map[string]int),
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
		State:   make(map[string]int),
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
		State:   make(map[string]int),
	}
}

// Probalistic tit-for-tat, orig logic, Joss
func Joss() *domain.Strategy {
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix: domain.Share,

			Prob: map[domain.Act]float64{
				domain.Take: 0.1,
			},
		},
		Trigger: &domain.Trigger{
			Act:   domain.Take,
			Count: 1,
			Reaction: domain.RuleValue{
				Fix: domain.Take,
			},
		},
		State: make(map[string]int),
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
			Reaction: domain.RuleValue{
				Fix: domain.Take,
				Prob: map[domain.Act]float64{
					domain.Share: 0.2,
					domain.Hold:  0.1,
				},
			},
		},
		State: make(map[string]int),
	}
}
