Preset strategies can have default values, as is the case for those created through the UI.

```
func STRATEGY() *Strategy {
	return &Strategy{
		Neutral: RuleValue{

			Fix: DEFAULT Share
            Prob: DEFAULT nil (disabled)
            Mirror: DEFAULT nil (disabled)

		},

		Trigger: &Trigger{      DEFAULT nil (disabled)

			Act:   DEFAULT Share
			Count: DEFAULT 1
			Mode:  DEFAULT TriggerStreakAfter,

			Reaction: RuleValue{
				Fix: DEFAULT Share
			},
		},
		State: DEFAULT nil
	}
}
```

AGENT STRUCTURE
An agent is a unit that carries a strategy.
```
type Agent struct {
	ID   AgID // ID is assigned either by request or sequentially
	Name string // Set during creation for display in the logger

	Memory   *Memory // Stores reputation and action history
	Strategy *Strategy // The strategy itself

	Modifiers []Modifier // Strategy extensions, applied in order

	Score int // OVERALL! score (can be set for non-equilibrium simulations)
	Dead  bool // Marks a dead agent if an evolutionary parameter is enabled
}
```

FOLDERS
benchmark — for advanced testing via code modifications
presets/mod — contains modifiers
presets/strategies — contains the agent-side logic of strategies
presets — contains agents and groups built from them

___ 

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

СТРУКТУРА АГЕНТА ВЫГЛЯДИТ ТАК
Агент - это юнит, который переносит стратегию
```
type Agent struct {
	ID   AgID // Айди выдается либо по реквесту, либо по порядку
	Name string // Задается при создании для отображения в логгере

	Memory   *Memory // Хранится репутация и история ходов
	Strategy *Strategy // Сама стратегия

	Modifiers []Modifier // Надстройки стратегии по порядку

	Score int // ОБЩИЙ! счет (можно назначить сразу для неравновесной симуляции)
	Dead  bool // Обозначение для умершего агента, если есть эволюционный параметр
}
```

ПАПКИ
benchmark - для продвинутого тестирования через изменения кода
presets/mod - здесь лежат моды
presets/strategies - здесь лежит агентная часть стратегий
presets - здесь лежат агенты и собранные из них группы