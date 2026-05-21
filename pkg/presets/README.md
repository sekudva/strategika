У стратегий в пресетах могут быть значения по умолчанию, также у стратегий, создающихйся через UI


func СТРАТЕГИЯ() *Strategy {
	return &Strategy{
		Neutral: RuleValue{
			Fix: ПО УМОЛЧАНИЮ Share
		},

		Trigger: &Trigger{
			Act:   ПО УМОЛЧАНИЮ Share
			Count: ПО УМОЛЧАНИЮ 1
			Mode:  ПО УМОЛЧАНИЮ TriggerStreakAfter,
			Reaction: RuleValue{
				Fix: ПО УМОЛЧАНИЮ Share
			},
		},
		State: ПО УМОЛЧАНИЮ nil
	}
}