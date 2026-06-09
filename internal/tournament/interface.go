package tournament

import "github.com/sekudva/strategika/internal/domain"

type RoundLogger interface {
	// Log вызывается для каждой пары агентов после applyPhase.
	Log(entry RoundLog)

	// Flush вызывается в конце симуляции (для вывода накопленного).
	// Возвращает историю раундов (может быть nil, если логгер её не хранит).
	Flush() []RoundLog

	// Finalize prints the final scores and any summary.
	Finalize(agents []*domain.Agent)

	// MarkDead used in ecosystem to check dead agents.
	MarkDead(agents []*domain.Agent, threshold int, round int)
}
