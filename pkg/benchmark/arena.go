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
	agents := presets.ClassicStrategies()
	// agents := presets.NonClassicGroup()
	// ===================================

	f, err := os.Create("arena_log.txt")
	if err != nil {
		return fmt.Errorf("failed to create log file: %w", err)
	}
	defer f.Close()

	cfg, err := presets.ArenaConfig(len(agents), 20, 0.0)
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	cfg.Logger = tournament.NewAggregateLogger(20, cfg.Pairs, agents, f)
	cfg.InfoTo(f)

	scores := cfg.RunSimulation(agents)

	fmt.Fprintf(f, "\n=== FINAL SCORES ===\n")
	for i, a := range agents {
		fmt.Fprintf(f, "%2d. %-20s: %d\n", i+1, a.Name, scores[a.ID])
	}

	tournament.PrintLeaderboard(agents)

	return nil
}
