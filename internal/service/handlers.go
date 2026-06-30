package service

import (
	"fmt"
	"os"

	"github.com/sekudva/strategika/internal/domain"
	"github.com/sekudva/strategika/internal/tournament"
	"github.com/sekudva/strategika/presets"
)

var (
	Silent bool
	Quick  bool
)

func RunDuel(a1, a2 *domain.Agent, rounds int, noise float64, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("File creating error: %v\n", err)
	}
	defer f.Close()

	cfg, err := presets.DuelConfig(rounds, noise)
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	cfg.Logger = newLogger(cfg, []*domain.Agent{a1, a2}, f, false)

	cfg.InfoTo(f)

	cfg.RunSimulation([]*domain.Agent{a1, a2})

	cfg.Logger.Finalize([]*domain.Agent{a1, a2})
	return nil
}

func RunRoundRobin(agents []*domain.Agent, rounds int, noise float64, noSelf bool, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("file error: %w", err)
	}
	defer f.Close()

	cfg, err := presets.DuelConfig(rounds, noise)
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	cfg.Logger = newLogger(cfg, agents, f, true)

	cfg.InfoTo(f)

	cfg.RoundRobin(agents, noSelf)

	return nil
}

func RunTrial(leader *domain.Agent, group []*domain.Agent, rounds int, noise float64, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("file error: %w", err)
	}
	defer f.Close()

	all := append([]*domain.Agent{leader}, group...)
	cfg, err := presets.TrialConfig(0, len(all), rounds, noise)
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}
	cfg.Logger = newLogger(cfg, all, f, true)
	cfg.InfoTo(f)
	cfg.RunTrial(leader, group)
	cfg.Logger.Finalize([]*domain.Agent{leader})
	return nil
}

func RunCirculaire(leaders []*domain.Agent, group []*domain.Agent, rounds int, noise float64, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("file error: %w", err)
	}
	defer f.Close()

	all := append(leaders, group...)
	cfg, err := presets.TrialConfig(0, len(all), rounds, noise)
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}
	cfg.Logger = newLogger(cfg, all, f, true)
	cfg.InfoTo(f)
	cfg.Circulaire(leaders, group)
	return nil
}

func RunArena(agents []*domain.Agent, rounds int, noise float64, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("file error: %w", err)
	}
	defer f.Close()

	cfg, err := presets.ArenaConfig(len(agents), rounds, noise)
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}
	cfg.Logger = newLogger(cfg, agents, f, true)
	cfg.InfoTo(f)
	cfg.RunSimulation(agents)
	cfg.Logger.Finalize(agents)
	return nil
}

func RunEcosystem(agents []*domain.Agent, rounds int, noise float64, deathThreshold int, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("file error: %w", err)
	}
	defer f.Close()

	cfg, err := presets.ArenaConfig(len(agents), rounds, noise)
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}
	cfg.Logger = newLogger(cfg, agents, f, true)
	cfg.InfoTo(f)
	cfg.RunEcosystem(agents, deathThreshold)
	return nil
}

func newLogger(cfg tournament.SimConfig, agents []*domain.Agent, w *os.File, all bool) tournament.RoundLogger {
	/// flag
	if Silent {
		return tournament.NewSilentLogger(w)
	}

	if all {
		return tournament.NewAllLogger(agents, w)
	}

	return tournament.NewAggregateLogger(cfg.Rounds/2, cfg.Pairs, agents, w)
}
