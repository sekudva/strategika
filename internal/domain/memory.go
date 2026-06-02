package domain

type Round struct {
	N     int
	MyAct Act
	OpAct Act
}
type Rounds []Round

type Memory struct {
	History map[AgID]Rounds
	Rep     Reputation
}

func NewMemory() *Memory {
	return &Memory{
		History: make(map[AgID]Rounds),
		Rep: Reputation{
			Coop: 0.5,
			Def:  0.5,
		},
	}
}

// Запись хода в память агента
func (m *Memory) Record(round int, partnerID AgID, my, their Act) {
	m.History[partnerID] = append(m.History[partnerID], Round{
		N:     round,
		MyAct: my,
		OpAct: their,
	})
}

// Возвращает последний ход ПРОТИВНИКА
// Если это первый ход в истории то возвращает -1
func (r Rounds) OpLastAct() Act {
	if len(r) <= 0 {
		return NoAct
	}
	return r[len(r)-1].OpAct
}

// Возвращает МОЙ последний ход
// Если это первый ход в истории то возвращает -1
func (r Rounds) MyLastAct() Act {
	if len(r) <= 0 {
		return NoAct
	}
	return r[len(r)-1].MyAct
}

// Счетчик триггера TriggerStreakAfter  для МОДИФИКАТОРОЫ
// Возможно не нужен?
func (r Rounds) CountTrigger(trigger Act) int {
	if len(r) <= 0 {
		return 0
	}

	count := 0

	for i := len(r) - 1; i >= 0; i-- {
		if r[i].OpAct == trigger {
			count++
		} else {
			break
		}
	}

	return count
}

// Для триггера внутри стратегии
func (r Rounds) streak(act Act) int {
	s := 0
	for i := len(r) - 1; i >= 0; i-- {
		if r[i].OpAct == act {
			s++
		} else {
			break
		}
	}
	return s
}
