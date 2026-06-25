package strategies

import "github.com/sekudva/strategika/internal/domain"

// Grudger (after Take only Take) orig logic
func Grudger() *domain.Strategy {
	m := domain.MirrorSelf
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix:    domain.Share,
			Mirror: &m,
		},
		Trigger: &domain.Trigger{
			Act:   domain.Take,
			Count: 1,
			Mode:  domain.TriggerExactly,
			Reaction: domain.RuleValue{
				Fix: domain.Take,
			},
		},
	}
}

// Starts with Take, but if 2 Shares - Share forever
// 2 shares bcs dont count first
func AntiGrudger() *domain.Strategy {
	m := domain.MirrorSelf
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix:    domain.Take,
			Mirror: &m,
		},
		Trigger: &domain.Trigger{
			Act:   domain.Share,
			Count: 2,
			Mode:  domain.TriggerExactly,
			Reaction: domain.RuleValue{
				Fix: domain.Share,
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
