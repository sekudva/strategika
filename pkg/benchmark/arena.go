package benchmark

import (
	"fmt"
	"os"

	"github.com/sekudva/strategika/internal/tournament"
	"github.com/sekudva/strategika/pkg/presets"
)

// RunArena запускает арену с предустановленной группой стратегий
// Выбери группу, раскомментировав нужную строку
func RunArena() error {
	// ========== ВЫБЕРИ ГРУППУ ==========
	//agents := presets.EvilDominates()
	//agents := presets.ClassicStrategies()
	//agents := presets.NonClassicGroup()
	// ===================================

	agents := append(presets.ClassicStrategies(), presets.NonClassicGroup()...)

	f, err := os.Create("arena_log.txt")
	if err != nil {
		return fmt.Errorf("failed to create log file: %w", err)
	}
	defer f.Close()

	cfg, err := presets.ArenaConfig(len(agents), 600, 0.0)
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	cfg.Logger = tournament.NewAggregateLogger(200, cfg.Pairs, agents, f)
	cfg.InfoTo(f)

	//cfg.RunEcosystem(agents, math.MinInt)
	cfg.RunEcosystem(agents, 0)
	return nil
}
