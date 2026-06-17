package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/sekudva/strategika/internal/domain"
	"github.com/sekudva/strategika/internal/tournament"
	"github.com/sekudva/strategika/pkg/presets"
)

func handleDuel(rounds int, noise float64) {
	fmt.Println("\n=== DUEL ===")
	fmt.Println("Choose 2 strategies:")

	a1 := selectOne(selectGroup())
	a2 := selectOne(selectGroup())

	fmt.Printf("\nДуэль: %s vs %s\n", a1.Name, a2.Name)

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

	cfg.Logger = tournament.NewAllLogger([]*domain.Agent{a1, a2}, f)
	cfg.InfoTo(f)

	cfg.RunSimulation([]*domain.Agent{a1, a2})

	cfg.Logger.Finalize([]*domain.Agent{a1, a2})
	fmt.Printf("\nResults in generated file: %s\n", f.Name())
}

func handleRoundRobin(rounds int, noise float64) {
	fmt.Println("\n=== ROUND-ROBIN ===")
	fmt.Println("Choose many strategies:")

	agents := selectEach(selectGroup())

	noSelf := readInt("Исключить само-игры? 1 - да: ", 0, 1, 0) == 1
	fmt.Printf("\nRound Robin: %d агентов, само-игры: %v\n", len(agents), !noSelf)

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

	cfg.Logger = tournament.NewAggregateLogger(rounds/2, tournament.DuelPairs(), agents, f)
	cfg.InfoTo(f)

	cfg.RoundRobin(agents, noSelf)

	fmt.Printf("\nResults in generated file: %s\n", f.Name())
}

func handleTrial(rounds int, noise float64) {
	fmt.Println("\n=== TRIAL ===")
	fmt.Println("Choose one lider and group:")

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

	cfg.Logger = tournament.NewAggregateLogger(rounds, cfg.Pairs, all, f)
	cfg.InfoTo(f)

	cfg.RunTrial(leader, group)

	cfg.Logger.Finalize(all)
	fmt.Printf("\nResults in generated file: %s\n", f.Name())
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

	cfg.Logger = tournament.NewAggregateLogger(rounds, tournament.TrialPairs(0, len(all)), all, f)
	cfg.InfoTo(f)

	cfg.Circulaire(leaders, group)

	fmt.Printf("\nResults in generated file: %s\n", f.Name())
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

	cfg.Logger = tournament.NewAggregateLogger(rounds/2, cfg.Pairs, agents, f)
	cfg.InfoTo(f)

	cfg.RunSimulation(agents)

	fmt.Printf("\nResults in generated file: %s\n", f.Name())
}

func handleEcosystem(rounds int, noise float64) {
	fmt.Println("\n=== ECOSYSTEM ===")
	fmt.Println("Choose many strategies:")

	agents := selectEach(selectGroup())

	threshold := readInt("Death threshold (scores) [0]: ", -100000, 100000, 0)
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

	cfg.Logger = tournament.NewAggregateLogger(rounds/2, cfg.Pairs, agents, f)
	cfg.InfoTo(f)

	cfg.RunEcosystem(agents, threshold)

	fmt.Printf("\nResults in generated file: %s\n", f.Name())
}

func handleInfo() {
	data, err := os.ReadFile("info.md")
	if err != nil {
		fmt.Println("Info file not found.")
		return
	}
	fmt.Println(strings.TrimSpace(string(data)))
}
