package presets

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

// pacifist
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

// crazy and irrational
func Crazy() *domain.Strategy {
	m := domain.MirrorOpp
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix:    domain.Share,
			Mirror: &m,
		},
		Trigger: nil,
		State:   make(map[string]int),
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
		State:   make(map[string]int),
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
		State:   make(map[string]int),
	}
}
