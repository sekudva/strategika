package tournament

import "github.com/sekudva/strategika/internal/domain"

// RoundLog — одна запись о взаимодействии в раунде
type RoundLog struct {
	Round          int
	Agent1, Agent2 domain.AgID
	Act1, Act2     domain.Act
	Score1, Score2 int // очки, начисленные в этом раунде
}

// AggregateLogger выводит агрегированную статистику каждые N раундов.
// Хранит только последний снапшот для каждой пары.
// Используется в турнире дуэлей и на арене.
type AggregateLogger struct {
	Interval int             // каждые сколько раундов выводить
	Pairs    []Pair          // для знания, какие пары существуют
	Agents   []*domain.Agent // для получения имён/ID

	// Внутренние счётчики для каждой пары
	counters map[Pair][3]int // [ShareCount, HoldCount, TakeCount]
	round    int
}

func NewAggregateLogger(interval int, pairs []Pair, agents []*domain.Agent) *AggregateLogger {
	return &AggregateLogger{
		Interval: interval,
		Pairs:    pairs,
		Agents:   agents,
		counters: make(map[Pair][3]int),
	}
}

func (l *AggregateLogger) Log(entry RoundLog) {
	p := Pair{0, 1} // для дуэли
	// В общем случае нужно найти пару по Agent1/Agent2
	// Но для дуэли это тривиально, для арены — ищем в l.Pairs
	c := l.counters[p]
	switch entry.Act1 {
	case domain.Share:
		c[0]++
	case domain.Hold:
		c[1]++
	case domain.Take:
		c[2]++
	}
	// Аналогично для Act2 (второго агента в паре)
	// Здесь упрощённо — для дуэли
	l.counters[p] = c

	l.round = entry.Round
}

func (l *AggregateLogger) Flush() []RoundLog {
	return nil // не хранит историю
}

// PrintStats вызывается из Flush или отдельно
func (l *AggregateLogger) Stats() string {
	// Формат: "TFT: 80/20/0 Joss: 10/0/90"
	// для каждой пары
	return "" // реализация ниже
}

// AllLogger сохраняет все раунды в срез.
type AllLogger struct {
	Logs []RoundLog
}

func (l *AllLogger) Log(entry RoundLog) {
	l.Logs = append(l.Logs, entry)
}
func (l *AllLogger) Flush() []RoundLog {
	return l.Logs
}

// SilentLogger не сохраняет и не выводит ничего.
type SilentLogger struct{}

func (l *SilentLogger) Log(entry RoundLog) {}
func (l *SilentLogger) Flush() []RoundLog  { return nil }
