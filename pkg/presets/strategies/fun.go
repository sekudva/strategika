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
