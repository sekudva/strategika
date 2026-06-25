package tournament

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/sekudva/strategika/internal/domain"
)

/*
	SILENT LOGGER
*/

// SilentLogger не сохраняет и не выводит ничего, кроме итогов
type SilentLogger struct {
	Writer io.Writer
}

func NewSilentLogger(w io.Writer) *SilentLogger {
	if w == nil {
		w = os.Stdout
	}
	return &SilentLogger{
		Writer: w,
	}
}

func (l *SilentLogger) ForDuel(pairs []Pair, agents []*domain.Agent) RoundLogger {
	return l // silent — не важно, можно тот же
}

func (l *SilentLogger) Log(entry RoundLog) {}

func (l *SilentLogger) Flush() []RoundLog { return nil }

func (l *SilentLogger) MarkDead(agents []*domain.Agent, threshold int, round int) {
	for _, a := range agents {
		if a.Score <= threshold && !a.Dead {
			a.Dead = true
		}
	}
}

func (l *SilentLogger) Finalize(agents []*domain.Agent) {
	fmt.Fprint(l.Writer, Leaderboard(agents, nil))
}

/*
	ALL LOGGER
*/

// RoundLog — одна запись о взаимодействии в раунде
type RoundLog struct {
	Round          int
	Agent1, Agent2 domain.AgID
	Act1, Act2     domain.Act
	Score1, Score2 int // очки, начисленные в этом раунде
}

// AllLogger сохраняет все раунды в срез.
type AllLogger struct {
	Logs   []RoundLog
	Agents []*domain.Agent
	Writer io.Writer
}

func NewAllLogger(agents []*domain.Agent, w io.Writer) *AllLogger {
	if w == nil {
		w = os.Stdout
	}
	return &AllLogger{
		Agents: agents,
		Writer: w,
		Logs:   make([]RoundLog, 0),
	}
}

func (l *AllLogger) ForDuel(pairs []Pair, agents []*domain.Agent) RoundLogger {
	return &AllLogger{
		Agents: agents,
		Writer: l.Writer,
		Logs:   make([]RoundLog, 0),
	}
}

func (l *AllLogger) Log(entry RoundLog) {
	l.Logs = append(l.Logs, entry)
}

func (l *AllLogger) Flush() []RoundLog {
	names := make(map[domain.AgID]string)
	for _, a := range l.Agents {
		names[a.ID] = a.Name
	}

	for _, entry := range l.Logs {
		name1 := names[entry.Agent1]
		name2 := names[entry.Agent2]
		if name1 == "" {
			name1 = fmt.Sprintf("Agent-%d", entry.Agent1)
		}
		if name2 == "" {
			name2 = fmt.Sprintf("Agent-%d", entry.Agent2)
		}

		fmt.Fprintf(l.Writer, "[%3d] %-20s vs %-20s | %-5s vs %-5s | %+3d / %+3d\n",
			entry.Round,
			name1, name2,
			entry.Act1.String(), entry.Act2.String(),
			entry.Score1, entry.Score2,
		)
	}

	result := l.Logs
	l.Logs = nil
	return result
}

func (l *AllLogger) MarkDead(agents []*domain.Agent, threshold int, round int) {
	for _, a := range agents {
		if a.Score <= threshold && !a.Dead {
			a.Dead = true

			fmt.Fprintf(l.Writer, strings.Repeat("=", 50)+"\n")
			fmt.Fprintf(l.Writer, "AGENT %s IS DEAD!	Score: %d	Round: %d\n",
				a.Name, a.Score, round,
			)

			fmt.Fprintf(l.Writer, strings.Repeat("=", 50)+"\n")
		}
	}
}
func (l *AllLogger) Finalize(agents []*domain.Agent) {
	fmt.Fprint(l.Writer, Leaderboard(agents, nil))
}

/*
	AGGREGATE LOGGER
*/

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

	// дата смерти
	deathLog map[domain.AgID]int

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
		deathLog: make(map[domain.AgID]int),
		counters: make(map[Pair][3]int),
		Writer:   w,
	}
}

func (l *AggregateLogger) ForDuel(pairs []Pair, agents []*domain.Agent) RoundLogger {
	return &AggregateLogger{
		Interval: l.Interval,
		Pairs:    pairs,
		Agents:   agents,
		counters: make(map[Pair][3]int),
		deathLog: make(map[domain.AgID]int),
		Writer:   l.Writer,
	}
}

func (l *AggregateLogger) Log(entry RoundLog) {
	p1, ok1 := l.findPair(entry.Agent1, entry.Agent2)
	p2, ok2 := l.findPair(entry.Agent2, entry.Agent1)

	if ok1 {
		c := l.counters[p1]
		l.increment(&c, entry.Act1)
		l.counters[p1] = c
	}

	if ok2 {
		c := l.counters[p2]
		l.increment(&c, entry.Act2)
		l.counters[p2] = c
	}

	l.round = entry.Round
}

// findPair ищет направленную пару по ID агентов.
// Возвращает Pair{fromIndex, toIndex} и true, если пара найдена.
func (l *AggregateLogger) findPair(a1, a2 domain.AgID) (Pair, bool) {
	for _, p := range l.Pairs {
		if l.Agents[p[0]].ID == a1 && l.Agents[p[1]].ID == a2 {
			return Pair{p[0], p[1]}, true
		}
		if l.Agents[p[0]].ID == a2 && l.Agents[p[1]].ID == a1 {
			return Pair{p[1], p[0]}, true
		}
	}
	return Pair{}, false
}

// increment увеличивает соответствующий счётчик действия.
func (l *AggregateLogger) increment(c *[3]int, act domain.Act) {
	switch act {
	case domain.Share:
		c[0]++
	case domain.Hold:
		c[1]++
	case domain.Take:
		c[2]++
	}
}

// Flush выводит финальную статистику и не хранит историю.
func (l *AggregateLogger) Flush() []RoundLog {
	if l.round%l.Interval == 0 {
		fmt.Fprintf(l.Writer, "\n--- Round %d ---\n%s", l.round, l.Stats())
		l.counters = make(map[Pair][3]int)
	}

	return nil
}

// PrintStats вызывается из Flush или отдельно
func (l *AggregateLogger) Stats() string {
	var sb strings.Builder
	for _, p := range l.Pairs {
		i, j := p[0], p[1]

		first := l.writeStatsLine(&sb, i, j, "→")
		second := l.writeStatsLine(&sb, j, i, "←")
		if first || second {
			fmt.Fprintf(&sb, "\n")
		}
	}

	return sb.String()
}

func (l *AggregateLogger) writeStatsLine(sb *strings.Builder, from, to int, arrow string) bool {
	c := l.counters[Pair{from, to}]
	total := c[0] + c[1] + c[2]
	if total == 0 {
		return false
	}
	fmt.Fprintf(sb, "%-15s %s %-15s | %3d / %3d / %3d | Score %-15s : %5d		:: %5d\n",
		l.Agents[from].Name, arrow, l.Agents[to].Name,
		pct(c[0], total), pct(c[1], total), pct(c[2], total),
		l.Agents[from].Name, l.Agents[from].Memory.DuelScore,
		l.Agents[from].Score,
	)
	return true
}

func pct(count, total int) int {
	if total == 0 {
		return 0
	}
	return count * 100 / total
}

func (l *AggregateLogger) MarkDead(agents []*domain.Agent, threshold int, round int) {
	for _, a := range agents {
		if a.Score <= threshold && !a.Dead {
			a.Dead = true
			l.deathLog[a.ID] = round

			fmt.Fprintf(l.Writer, strings.Repeat("=", 50)+"\n")
			fmt.Fprintf(l.Writer, "AGENT %s IS DEAD!\n", a.Name)
			fmt.Fprintf(l.Writer, strings.Repeat("=", 50)+"\n")
		}
	}
}

func (l *AggregateLogger) Finalize(agents []*domain.Agent) {
	fmt.Fprint(l.Writer, Leaderboard(agents, l.deathLog))
}

/*
FINALIZE
*/

func Leaderboard(agents []*domain.Agent, deathLog map[domain.AgID]int) string {
	sort.Slice(agents, func(i, j int) bool {
		return agents[i].Score > agents[j].Score
	})

	deathPRINT := false

	var sb strings.Builder
	fmt.Fprintf(&sb, "\n")
	fmt.Fprintf(&sb, strings.Repeat("=", 50)+"\n")
	fmt.Fprintf(&sb, "FINAL LEADERBOARD\n")
	fmt.Fprintf(&sb, strings.Repeat("=", 50)+"\n")
	fmt.Fprintf(&sb, "%s\t| %-23s\t| %8s\n", "TOP", "Strategy", "Score")
	fmt.Fprintf(&sb, strings.Repeat("-", 50)+"\n")

	for i, a := range agents {
		if !deathPRINT && a.Dead {
			fmt.Fprintf(&sb, strings.Repeat("-", 50)+"\n")
			fmt.Fprintf(&sb, "☠ DEAD AGENTS ☠\n")
			fmt.Fprintf(&sb, strings.Repeat("-", 50)+"\n")
			deathPRINT = true
		}

		deathInfo := ""
		if a.Dead {
			if r, ok := deathLog[a.ID]; ok {
				deathInfo = fmt.Sprintf(" ☠ r%d", r)
			} else {
				deathInfo = " ☠"
			}
		}

		fmt.Fprintf(&sb, "№%-3d\t%-3d) %-20s\t|%8d	%s\n", i+1, a.ID, a.Name, a.Score, deathInfo)
	}

	fmt.Fprintf(&sb, strings.Repeat("=", 50)+"\n")
	return sb.String()
}
