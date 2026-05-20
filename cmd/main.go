package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/sekudva/strategika/internal/domain"
	"github.com/sekudva/strategika/internal/tournament"
	"github.com/sekudva/strategika/pkg/presets"
)

func main() {
	// 1. Определяем пул стратегий для турнира
	type NamedStrategy struct {
		Name string
		Fn   func() *domain.Strategy
	}

	strategies := []NamedStrategy{
		{"TitForTat", presets.TitForTat},
		{"TitFor2Tats", presets.TitFor2Tats},
		{"Joss", presets.Joss},
		{"Masochist", presets.Masochist},
		{"Sadist", presets.Sadist},
		{"Bully", presets.Bully},
		{"Pacifist", presets.Pacifist},
		{"Crazy", presets.Crazy},
		{"AlwaysShare", presets.AlwaysShare},
		{"AlwaysTake", presets.AlwaysTake},
	}

	const rounds = 10
	const noise = 0.0 // Без шума для чистоты эксперимента

	fmt.Println(strings.Repeat("=", 70))
	fmt.Printf("🎲 TOURNAMENT: %d strategies × %d rounds\n", len(strategies), rounds)
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println()

	// 2. Структуры для хранения результатов
	type DuelSummary struct {
		Name1, Name2   string
		Score1, Score2 int
	}

	var allResults []DuelSummary
	totalScores := make(map[string]int)
	wins := make(map[string]int)
	losses := make(map[string]int)
	draws := make(map[string]int)

	// 3. Запуск турнира: каждый с каждым
	for i := 0; i < len(strategies); i++ {
		for j := i + 1; j < len(strategies); j++ {
			strategy1, strategy2 := strategies[i], strategies[j]

			// Создаём свежих агентов для каждой дуэли
			gen := domain.NewIDGenerator()
			agent1 := domain.NewAgent(strategy1.Fn(), gen.Next())
			agent2 := domain.NewAgent(strategy2.Fn(), gen.Next())

			fmt.Printf("📌 %s vs %s\n", strategy1.Name, strategy2.Name)
			fmt.Println(strings.Repeat("─", 50))

			// Запуск дуэли через турнирный движок
			duel := tournament.Duel{Rounds: rounds, Noise: noise}
			result, err := duel.Run([]*domain.Agent{agent1, agent2})
			if err != nil {
				fmt.Printf("Error running duel: %v\n", err)
				os.Exit(1)
			}

			duelResult := result.(tournament.DuelResult) // Приведение типа
			score1, score2 := duelResult.Score1, duelResult.Score2

			// Логирование каждого раунда
			for r := 0; r < rounds; r++ {
				act1 := agent1.Memory.History[agent2.ID][r].MyAct
				act2 := agent2.Memory.History[agent1.ID][r].MyAct
				// Выводим накопленный счёт после каждого раунда
				roundScore1, roundScore2 := 0, 0
				for k := 0; k <= r; k++ {
					s1, _ := domain.Payoff(agent1.Memory.History[agent2.ID][k].MyAct, agent2.Memory.History[agent1.ID][k].MyAct)
					roundScore1 += s1
					s2, _ := domain.Payoff(agent2.Memory.History[agent1.ID][k].MyAct, agent1.Memory.History[agent2.ID][k].MyAct)
					roundScore2 += s2
				}
				fmt.Printf(
					"[%2d] %-12s : %-6s (%-+3d)  |  %-12s : %-6s (%-+3d)\n",
					r+1,
					strategy1.Name, act1.String(), roundScore1,
					strategy2.Name, act2.String(), roundScore2,
				)
			}

			// Итог дуэли
			fmt.Println(strings.Repeat("─", 50))
			fmt.Printf("🏆 RESULT: %s %d  —  %d %s\n", strategy1.Name, score1, score2, strategy2.Name)
			fmt.Println()

			// Сохраняем результаты
			allResults = append(allResults, DuelSummary{
				Name1: strategy1.Name, Name2: strategy2.Name,
				Score1: score1, Score2: score2,
			})

			totalScores[strategy1.Name] += score1
			totalScores[strategy2.Name] += score2

			if score1 > score2 {
				wins[strategy1.Name]++
				losses[strategy2.Name]++
			} else if score2 > score1 {
				wins[strategy2.Name]++
				losses[strategy1.Name]++
			} else {
				draws[strategy1.Name]++
				draws[strategy2.Name]++
			}
		}
	}

	// 4. Формируем итоговую таблицу
	type LeaderboardEntry struct {
		Name    string
		Score   int
		Wins    int
		Losses  int
		Draws   int
		WinRate float64
	}

	var leaderboard []LeaderboardEntry
	for _, s := range strategies {
		name := s.Name
		w := wins[name]
		l := losses[name]
		d := draws[name]
		total := w + l + d
		winRate := 0.0
		if total > 0 {
			winRate = float64(w) / float64(total) * 100
		}
		leaderboard = append(leaderboard, LeaderboardEntry{
			Name: name, Score: totalScores[name],
			Wins: w, Losses: l, Draws: d,
			WinRate: winRate,
		})
	}

	// Сортировка по убыванию очков
	sort.Slice(leaderboard, func(i, j int) bool {
		return leaderboard[i].Score > leaderboard[j].Score
	})

	// 5. Вывод таблицы лидеров
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println("📊 FINAL LEADERBOARD")
	fmt.Println(strings.Repeat("=", 70))
	fmt.Printf("\n%-12s | %6s | %4s | %4s | %4s | %6s\n",
		"Strategy", "Score", "W", "L", "D", "Win%")
	fmt.Println(strings.Repeat("─", 50))

	for _, entry := range leaderboard {
		fmt.Printf("%-12s | %6d | %4d | %4d | %4d | %5.1f%%\n",
			entry.Name, entry.Score, entry.Wins, entry.Losses, entry.Draws, entry.WinRate)
	}

	fmt.Println()
	fmt.Println("🏁 Tournament finished.")
}
