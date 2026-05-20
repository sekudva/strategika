package presets

import "github.com/sekudva/strategika/internal/domain"

// original logic tit-for-tat: with only share/take
func Tit_for_Tat() *domain.Strategy {
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

// tit-for-two-tats NEW logic
func Tit_for_Two_Tats() *domain.Strategy {
	direct := domain.MirrorDirect
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix:    domain.Share,
			Mirror: &direct,
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
func Tit_for_Tat_NEW() *domain.Strategy {
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
func Tit_for_Tat_EVIL() *domain.Strategy {
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
