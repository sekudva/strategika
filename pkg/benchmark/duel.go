package benchmark

import (
	"fmt"
	"os"

	"github.com/sekudva/strategika/internal/domain"
	"github.com/sekudva/strategika/internal/tournament"
	"github.com/sekudva/strategika/pkg/presets"
)

// RunDuel запускает дуэль между двумя агентами с полным логированием
// Выбери агентов, раскомментировав нужные строки
func RunDuel() error {
	// ========== ВЫБЕРИ АГЕНТОВ ==========
	// agent1 := presets.TitForTat()
	// agent2 := presets.AlwaysTake()

	// agent1 := presets.Grudger()
	// agent1 := presets.Joss()
	//agent1 := presets.Paladin()
	//agent1 := presets.AntiGrudger()
	// agent1 := presets.Crazy()

	agent1 := presets.TitFor2Tats()
	//agent2 := presets.ColdGhost()

	//agent1 := presets.Journalist()
	//agent1 := presets.TitForTat()

	//agent1 := presets.AlwaysShare()
	agent2 := presets.Boev()
	// ===================================

	f, err := os.Create("duel_log.txt")
	if err != nil {
		return fmt.Errorf("failed to create log file: %w", err)
	}
	defer f.Close()

	cfg, err := presets.DuelConfig(200, 0.0)
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	cfg.Logger = tournament.NewAllLogger([]*domain.Agent{agent1, agent2}, f)
	cfg.InfoTo(f)

	cfg.RunSimulation([]*domain.Agent{agent1, agent2})

	cfg.Logger.Finalize([]*domain.Agent{agent1, agent2})

	return nil
}
