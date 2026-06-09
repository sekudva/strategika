package domain

type ModContext struct {
	History  Rounds
	Round    int
	ModState map[Counter]int
	OpRep    Reputation
}

// храниться в агенте
type Modifier func(core Act, ctx ModContext) Act

// храниться в памяти, используется агентом
type Counter int

const (
	RepeatCounter        Counter = iota // сколько раз повторили действие
	GoodStreakCounter                   // сколько Share подряд от партнёра
	EscalationCounter                   // шаг эскалации
	PhaseStepCounter                    // шаг фазы оценки
	ActRecorder                         // Так как действия тоже INT их можно хранить в счетчиках для удобства
	CurrThresholdCounter                // Текущее пороговое значение
)
