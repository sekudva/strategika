package domain

// Правило внутри стратегии, которое предполагает повторение ходов противника
type MirrorMode int

const (
	MirrorDirect MirrorMode = iota // полное копирование

	MirrorDefense // Take → Hold
	MirrorBad     // Hold → Take

	MirrorCold // Share → Hold
	MirrorNice // Hold → Share

	MirrorOpp // Share ↔ Take

	MirrorSelf // копирование своего хода

	// Можно расширить опционал, но чисто теоретически это бессмысленно
)

var mirrorTable = map[MirrorMode]map[Act]Act{
	MirrorDirect: {
		Share: Share,
		Hold:  Hold,
		Take:  Take,
	},

	MirrorDefense: {
		Share: Share,
		Hold:  Hold,
		Take:  Hold, // Take → Hold RULE
	},

	MirrorBad: {
		Share: Share,
		Hold:  Take, // Hold → Take RULE
		Take:  Take,
	},

	MirrorCold: {
		Share: Hold, // Share → Hold RULE
		Hold:  Hold,
		Take:  Take,
	},

	MirrorNice: {
		Share: Share,
		Hold:  Share, // Hold → Share RULE
		Take:  Take,
	},

	MirrorOpp: {
		Share: Take, // Share ↔ Take RULE
		Hold:  Hold,
		Take:  Share, // Take ↔ Share RULE
	},
}

type TriggerMode int

const (
	TriggerStreakAfter TriggerMode = iota // После N раз (пока не закончится серия)
	TriggerEvery                          // каждый N-й подряд в серии (n=1 = на каждый)
	TriggerExactly                        // ровно на N-й, потом пропускать ответ на триггер

)

type TriggerTest func(history []Round, act Act, n int) bool

var triggerTable = map[TriggerMode]TriggerTest{

	TriggerStreakAfter: func(h []Round, act Act, n int) bool {
		streak := 0
		for i := len(h) - 1; i >= 0; i-- {
			if h[i].OpAct == act {
				streak++
			} else {
				break
			}
		}
		return streak >= n
	},

	TriggerEvery: func(h []Round, act Act, n int) bool {
		streak := 0
		for i := len(h) - 1; i >= 0; i-- {
			if h[i].OpAct == act {
				streak++
			} else {
				break
			}
		}
		return streak >= n && streak%n == 0
	},

	TriggerExactly: func(h []Round, act Act, n int) bool {
		streak := 0
		for i := len(h) - 1; i >= 0; i-- {
			if h[i].OpAct == act {
				streak++
			} else {
				break
			}
		}
		return streak == n
	},
}

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
	State   map[string]int // счетчик стратегии, используется редко
}

func NewStrategy(neutral RuleValue, trigger *Trigger) *Strategy {
	return &Strategy{
		Neutral: neutral,
		Trigger: trigger,
		State:   make(map[string]int),
	}
}
