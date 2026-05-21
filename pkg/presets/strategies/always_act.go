package strategies

import (
	"github.com/sekudva/strategika/internal/domain"
)

// always share
func AlwaysShare() *domain.Strategy {
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix: domain.Share,
		},
		Trigger: nil,
	}
}

// always hold
func AlwaysHold() *domain.Strategy {
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix: domain.Hold,
		},
		Trigger: nil,
	}
}

// always take
func AlwaysTake() *domain.Strategy {
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix: domain.Take,
		},
		Trigger: nil,
	}
}

// true random
func Random() *domain.Strategy {
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix: domain.Hold,
			Prob: map[domain.Act]float64{
				domain.Share: 0.33,
				domain.Hold:  0.34,
				domain.Take:  0.33,
			},
		},
		Trigger: nil,
	}
}

// true random without Hold
func Random_OLD() *domain.Strategy {
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix: domain.Hold,
			Prob: map[domain.Act]float64{
				domain.Share: 0.5,
				domain.Take:  0.5,
			},
		},
		Trigger: nil,
	}
}
