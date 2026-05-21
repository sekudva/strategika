У стратегий в пресетах могут быть значения по умолчанию, также у стратегий, создающихйся через UI

```
func СТРАТЕГИЯ() *Strategy {
	return &Strategy{
		Neutral: RuleValue{

			Fix: ПО УМОЛЧАНИЮ Share
            Prob: ПО УМОЛЧАНИЮ nil (не включено)
            Mirror: ПО УМОЛЧАНИЮ nil (не включено)

		},

		Trigger: &Trigger{      ПО УМОЛЧАНИЮ nil (не включено)

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
```