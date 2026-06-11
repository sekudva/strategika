package benchmark

import (
	"fmt"
	"os"

	"github.com/sekudva/strategika/internal/tournament"
	"github.com/sekudva/strategika/pkg/presets"
)

// RunRoundRobinTournament запускает круговой турнир всех стратегий
// Результат записывается в log.txt в текущей директории
func RunRoundRobinTournament() error {
	// ========== ВЫБЕРИ ГРУППУ ==========
	//agents := presets.EvilDominates()
	//agents := presets.ClassicStrategies()
	//agents := presets.NonClassicGroup()
	// ===================================

	agents := append(presets.ClassicStrategies(), presets.NonClassicGroup()...)

	f, err := os.Create("log.txt")
	if err != nil {
		return fmt.Errorf("failed to create log file: %w", err)
	}
	defer f.Close()

	cfg, err := presets.DuelConfig(100, 0.0)
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	cfg.Logger = tournament.NewAggregateLogger(100, cfg.Pairs, agents, f)
	cfg.InfoTo(f)

	cfg.RoundRobin(agents, true)

	return nil
}
