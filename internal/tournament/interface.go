package tournament

type RoundLogger interface {
	// Log вызывается для каждой пары агентов после applyPhase.
	Log(entry RoundLog)

	// Flush вызывается в конце симуляции (для вывода накопленного).
	// Возвращает историю раундов (может быть nil, если логгер её не хранит).
	Flush() []RoundLog
}
