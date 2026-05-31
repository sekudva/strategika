package main

import (
	"fmt"
	"os"

	"github.com/sekudva/strategika/internal/tournament"
	"github.com/sekudva/strategika/pkg/presets"
)

func main() {
	agents := presets.EvilDominates()

	//agents := append(presets.ClassicStrategies(), presets.NonClassicGroup()...)

	f, _ := os.Create("log.txt")
	//s := os.Stdout
	defer f.Close()

	// Конфигурация с AggregateLogger для статистики
	cfg, err := presets.DuelConfig(10, 0.0)
	if err != nil {
		fmt.Printf("config error: %v\n", err)
		return
	}
	// Переопределяем логгер на AggregateLogger с выводом каждые 10 раундов
	cfg.Logger = tournament.NewAggregateLogger(10, cfg.Pairs, agents, f)

	cfg.InfoTo(f)

	// Запускаем RoundRobin
	err = tournament.RoundRobin(cfg, agents)
	if err != nil {
		fmt.Printf("tournament error: %v\n", err)
		return
	}

	// Итоговая таблица
	tournament.PrintLeaderboard(agents)
}
