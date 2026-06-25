package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/sekudva/strategika/internal/domain"
	"github.com/sekudva/strategika/internal/tournament"
	"github.com/sekudva/strategika/presets"
)

func handleDuel(rounds int, noise float64) {
	fmt.Println("\n=== DUEL ===")
	fmt.Println("Choose 2 strategies:")

	agent1 := selectOne(selectGroup())
	agent2 := selectOne(selectGroup())

	fmt.Printf("\nDuel: %s VS %s\n", agent1.Name, agent2.Name)

	f, err := os.Create("duel_log.txt")
	if err != nil {
		fmt.Printf("File creating error: %v\n", err)
		return
	}
	defer f.Close()

	cfg, err := presets.DuelConfig(rounds, noise)
	if err != nil {
		fmt.Printf("Strategika's configuration error: %v\n", err)
		return
	}

	if Silent {
		cfg.Logger = tournament.NewSilentLogger(f)
	} else {
		cfg.Logger = tournament.NewAllLogger([]*domain.Agent{agent1, agent2}, f)
	}
	cfg.InfoTo(f)

	cfg.RunSimulation([]*domain.Agent{agent1, agent2})

	cfg.Logger.Finalize([]*domain.Agent{agent1, agent2})
	Results(f.Name())
}

func handleRoundRobin(rounds int, noise float64) {
	fmt.Println("\n=== ROUND-ROBIN ===")
	fmt.Println("Choose many strategies:")

	agents := selectEach(selectGroup())

	noSelf := true
	if !Quick {
		noSelf = readInt("Allow self-games? 0 - NO, 1 - YES [NO]: ", 0, 1, 0) == 0
	}

	fmt.Printf("\nRound Robin: %d agents, self-games: %v\n", len(agents), !noSelf)

	f, err := os.Create("roundrobin_log.txt")
	if err != nil {
		fmt.Printf("File creating error: %v\n", err)
		return
	}
	defer f.Close()

	cfg, err := presets.DuelConfig(rounds, noise)
	if err != nil {
		fmt.Printf("Strategika's configuration error: %v\n", err)
		return
	}

	if Silent {
		cfg.Logger = tournament.NewSilentLogger(f)
	} else {
		cfg.Logger = tournament.NewAggregateLogger(rounds/2, tournament.DuelPairs(), agents, f)
	}
	cfg.InfoTo(f)

	cfg.RoundRobin(agents, noSelf)

	Results(f.Name())
}

func handleTrial(rounds int, noise float64) {
	fmt.Println("\n=== TRIAL ===")
	fmt.Println("Choose one leader and group:")

	fmt.Println("\nChoose leader (or victim):")
	leader := selectOne(selectGroup())

	fmt.Println("Choose group:")
	group := selectEach(selectGroup())

	fmt.Printf("\nTrial: %s VS %d agents\n", leader.Name, len(group))
	all := append([]*domain.Agent{leader}, group...)

	f, err := os.Create("trial_log.txt")
	if err != nil {
		fmt.Printf("File creating error: %v\n", err)
		return
	}
	defer f.Close()

	cfg, err := presets.TrialConfig(0, len(all), rounds, noise)
	if err != nil {
		fmt.Printf("Strategika's configuration error: %v\n", err)
		return
	}

	if Silent {
		cfg.Logger = tournament.NewSilentLogger(f)
	} else {
		cfg.Logger = tournament.NewAggregateLogger(rounds, cfg.Pairs, all, f)
	}
	cfg.InfoTo(f)

	cfg.RunTrial(leader, group)

	cfg.Logger.Finalize([]*domain.Agent{leader})
	Results(f.Name())
}

func handleCirculaire(rounds int, noise float64) {
	fmt.Println("\n=== CIRCULAIRE ===")
	fmt.Println("Choose many strategies, many LIDERS and one GROUP:")

	fmt.Println("\nChoose leaders:")
	leaders := selectEach(selectGroup())

	fmt.Println("Choose group:")
	group := selectEach(selectGroup())

	fmt.Printf("\nCirculaire: %d leaders VS %d agents\n", len(leaders), len(group))
	all := append(leaders, group...)

	f, err := os.Create("circulaire_log.txt")
	if err != nil {
		fmt.Printf("File creating error: %v\n", err)
		return
	}
	defer f.Close()

	cfg, err := presets.TrialConfig(0, len(all), rounds, noise)
	if err != nil {
		fmt.Printf("Strategika's configuration error: %v\n", err)
		return
	}

	if Silent {
		cfg.Logger = tournament.NewSilentLogger(f)
	} else {
		cfg.Logger = tournament.NewAggregateLogger(rounds, tournament.TrialPairs(0, len(all)), all, f)
	}
	cfg.InfoTo(f)

	cfg.Circulaire(leaders, group)

	Results(f.Name())
}

func handleArena(rounds int, noise float64) {
	fmt.Println("\n=== ARENA ===")
	fmt.Println("Choose many strategies:")

	agents := selectEach(selectGroup())

	fmt.Printf("\nArena: %d agents\n", len(agents))

	f, err := os.Create("arena_log.txt")
	if err != nil {
		fmt.Printf("File creating error: %v\n", err)
		return
	}
	defer f.Close()

	cfg, err := presets.ArenaConfig(len(agents), rounds, noise)
	if err != nil {
		fmt.Printf("Strategika's configuration error: %v\n", err)
		return
	}

	if Silent {
		cfg.Logger = tournament.NewSilentLogger(f)
	} else {
		cfg.Logger = tournament.NewAggregateLogger(rounds/2, cfg.Pairs, agents, f)
	}
	cfg.InfoTo(f)

	cfg.RunSimulation(agents)

	cfg.Logger.Finalize(agents)
	Results(f.Name())
}

func handleEcosystem(rounds int, noise float64) {
	fmt.Println("\n=== ECOSYSTEM ===")
	fmt.Println("Choose many strategies:")

	agents := selectEach(selectGroup())

	threshold := 0
	if !Quick {
		threshold = readInt("Death threshold (scores) [0]: ", -100000, 100000, 0)
	}

	fmt.Printf("\nEcosystem: %d agents, death threshold: %d\n", len(agents), threshold)

	f, err := os.Create("ecosystem_log.txt")
	if err != nil {
		fmt.Printf("File creating error: %v\n", err)
		return
	}
	defer f.Close()

	cfg, err := presets.ArenaConfig(len(agents), rounds, noise)
	if err != nil {
		fmt.Printf("Strategika's configuration error: %v\n", err)
		return
	}

	if Silent {
		cfg.Logger = tournament.NewSilentLogger(f)
	} else {
		cfg.Logger = tournament.NewAggregateLogger(rounds/2, cfg.Pairs, agents, f)
	}
	cfg.InfoTo(f)

	cfg.RunEcosystem(agents, threshold)

	Results(f.Name())
}

func handleInfo() {
	data, err := os.ReadFile("info.txt")
	if err != nil {
		fmt.Println("Info file not found.")
		return
	}
	fmt.Println(strings.TrimSpace(string(data)))
}

func Results(name string) {
	fmt.Printf("\nResults in generated file: %s\n", name)
	fmt.Print("WARNING: This file will override previous version\nif you run simulation with same Simulation Mode!\n")
	fmt.Print("If you want to save result, rename previous file.\n")
}
