package tournament

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/sekudva/strategika/internal/domain"
)

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
	Agents   []*domain.Agent // для получения имён/ID и счета

	// Внутренние счётчики для каждой пары
	counters map[Pair][3]int // [ShareCount, HoldCount, TakeCount]
	round    int

	Writer io.Writer // чтобы писать в файлики а не в терминал
}

func NewAggregateLogger(interval int, pairs []Pair, agents []*domain.Agent, w io.Writer) *AggregateLogger {
	if w == nil {
		w = os.Stdout // по умолчанию в консоль
	}

	return &AggregateLogger{
		Interval: interval,
		Pairs:    pairs,
		Agents:   agents,
		counters: make(map[Pair][3]int),
		Writer:   w,
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
	l.counters[p] = c

	// Аналогично для Act2 (второго агента в паре)
	p2 := Pair{1, 0}
	c2 := l.counters[p2]
	switch entry.Act2 {
	case domain.Share:
		c2[0]++
	case domain.Hold:
		c2[1]++
	case domain.Take:
		c2[2]++
	}
	l.counters[p2] = c2

	l.round = entry.Round
}

func (l *AggregateLogger) Flush() []RoundLog {
	fmt.Fprint(l.Writer, l.Stats())
	return nil // не хранит историю
}

// PrintStats вызывается из Flush или отдельно
func (l *AggregateLogger) Stats() string {
	var sb strings.Builder
	for _, p := range l.Pairs {
		i, j := p[0], p[1]

		// Статистика агента i против j
		c1 := l.counters[Pair{i, j}]
		total1 := c1[0] + c1[1] + c1[2]

		// Статистика агента j против i
		c2 := l.counters[Pair{j, i}]
		total2 := c2[0] + c2[1] + c2[2]

		if total1 == 0 && total2 == 0 {
			continue
		}

		fmt.Fprintf(&sb, "%-15s → %-15s | %3d / %3d / %3d	| Score %-15s : %d\n",
			l.Agents[i].Name, l.Agents[j].Name,
			pct(c1[0], total1), pct(c1[1], total1), pct(c1[2], total1), l.Agents[i].Name, l.Agents[i].Score)

		fmt.Fprintf(&sb, "%-15s ← %-15s | %3d / %3d / %3d	| Score %-15s : %d\n\n",
			l.Agents[i].Name, l.Agents[j].Name,
			pct(c2[0], total2), pct(c2[1], total2), pct(c2[2], total2), l.Agents[j].Name, l.Agents[j].Score)
	}

	return sb.String()
}

func pct(count, total int) int {
	if total == 0 {
		return 0
	}
	return count * 100 / total
}

// AllLogger сохраняет все раунды в срез.
type AllLogger struct {
	Logs   []RoundLog
	Writer io.Writer
}

func NewAllLogger(w io.Writer) *AllLogger {
	if w == nil {
		w = os.Stdout
	}
	return &AllLogger{Writer: w}
}

func (l *AllLogger) Log(entry RoundLog) {
	l.Logs = append(l.Logs, entry)
}

func (l *AllLogger) Flush() []RoundLog {
	for _, entry := range l.Logs {
		fmt.Fprintf(l.Writer, "[%3d] %d → %d | %s vs %s\n",
			entry.Round,
			entry.Agent1, entry.Agent2,
			entry.Act1.String(), entry.Act2.String(),
		)
	}

	return l.Logs
}

// SilentLogger не сохраняет и не выводит ничего.
type SilentLogger struct{}

func (l *SilentLogger) Log(entry RoundLog) {}
func (l *SilentLogger) Flush() []RoundLog  { return nil }

func PrintLeaderboard(agents []*domain.Agent) {
	sort.Slice(agents, func(i, j int) bool {
		return agents[i].Score > agents[j].Score
	})

	fmt.Println()
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("FINAL LEADERBOARD")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("%-20s | %8s\n", "Strategy", "Score")
	fmt.Println(strings.Repeat("-", 35))

	for _, a := range agents {
		fmt.Printf("%-20s | %8d\n", a.Name, a.Score)
	}
	fmt.Println(strings.Repeat("=", 60))
}
