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
		fmt.Printf("Ошибка создания файла: %v\n", err)
		return
	}
	defer f.Close()

	cfg, err := presets.DuelConfig(rounds, noise)
	if err != nil {
		fmt.Printf("Ошибка конфигурации: %v\n", err)
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

	f, _ := os.Create("roundrobin_log.txt")
	defer f.Close()

	cfg, _ := presets.DuelConfig(rounds, noise)
	cfg.Logger = tournament.NewAggregateLogger(rounds, tournament.DuelPairs(), agents, f)
	cfg.InfoTo(f)
	cfg.RoundRobin(agents, noSelf)

	fmt.Printf("\nResults in generated file: %s\n", f.Name())
}

func handleTrial(rounds int, noise float64) {
	fmt.Println("\n=== TRIAL ===")
	fmt.Println("Choose one lider and group:")

	fmt.Println("\nВыберите лидера:")
	leader := selectOne(selectGroup())

	fmt.Println("Выберите группу для испытания:")
	group := selectEach(selectGroup())

	fmt.Printf("\nTrial: %s против %d агентов\n", leader.Name, len(group))

	all := append([]*domain.Agent{leader}, group...)
	cfg, _ := presets.TrialConfig(0, len(all), rounds, noise)

	f, _ := os.Create("trial_log.txt")
	defer f.Close()

	cfg.Logger = tournament.NewAggregateLogger(rounds, cfg.Pairs, all, f)
	cfg.InfoTo(f)
	cfg.RunTrial(leader, group)
	cfg.Logger.Finalize(all)

	fmt.Printf("\nResults in generated file: %s\n", f.Name())
}

func handleCirculaire(rounds int, noise float64) {
	fmt.Println("\n=== CIRCULAIRE ===")
	fmt.Println("Choose many strategies, many LIDERS and one GROUP:")

	fmt.Println("\nВыберите лидеров:")
	leaders := selectEach(selectGroup())
	fmt.Println("Выберите группу для испытаний:")
	group := selectEach(selectGroup())

	fmt.Printf("\nCirculaire: %d лидеров против %d агентов\n", len(leaders), len(group))

	f, _ := os.Create("circulaire_log.txt")
	defer f.Close()

	all := append(leaders, group...)
	cfg, _ := presets.TrialConfig(0, len(all), rounds, noise)
	cfg.Logger = tournament.NewAggregateLogger(rounds, tournament.TrialPairs(0, len(all)), all, f)
	cfg.InfoTo(f)
	cfg.Circulaire(leaders, group)
	fmt.Printf("\nResults in generated file: %s\n", f.Name())
}

func handleArena(rounds int, noise float64) {
	fmt.Println("\n=== ARENA ===")
	fmt.Println("Choose many strategies:")

	agents := selectEach(selectGroup())

	fmt.Printf("\nАрена: %d агентов\n", len(agents))

	f, _ := os.Create("arena_log.txt")
	defer f.Close()

	cfg, _ := presets.ArenaConfig(len(agents), rounds, noise)
	cfg.Logger = tournament.NewAggregateLogger(rounds, cfg.Pairs, agents, f)
	cfg.InfoTo(f)
	cfg.RunSimulation(agents)

	fmt.Printf("\nResults in generated file: %s\n", f.Name())
}

func handleEcosystem(rounds int, noise float64) {
	fmt.Println("\n=== ECOSYSTEM ===")
	fmt.Println("Choose many strategies:")

	agents := selectEach(selectGroup())

	threshold := readInt("Порог смерти (очки) [0]: ", -100000, 100000, 0)
	fmt.Printf("\nЭкосистема: %d агентов, порог смерти: %d\n", len(agents), threshold)

	f, _ := os.Create("ecosystem_log.txt")
	defer f.Close()

	cfg, _ := presets.ArenaConfig(len(agents), rounds, noise)
	cfg.Logger = tournament.NewAggregateLogger(rounds, cfg.Pairs, agents, f)
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
