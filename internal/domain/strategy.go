package domain

// Фиксированное значение хода
type RuleValue struct {
	Fix    Act
	Prob   map[Act]float64
	Mirror *MirrorMode // nil == no MirrorMode
}

// Действие, на которое реагирует агент
type Trigger struct {
	Act      Act
	Count    int
	Mode     TriggerMode
	Reaction RuleValue // ответ на триггер
}

// Правила стратегии
type Strategy struct {
	Neutral RuleValue // нейтральное состояние
	Trigger *Trigger
}

func NewStrategy(neutral RuleValue, trigger *Trigger) *Strategy {
	return &Strategy{
		Neutral: neutral,
		Trigger: trigger,
	}
}
