package cli

import (
	"fmt"
)

// Run — основной хендлер, запускает меню и диспатчит режимы.
func Run() error {
	for {
		mode := MainMenu()
		if mode == 0 {
			return nil
		}
		var rounds int
		var noise float64

		if mode != 7 {
			rounds = readInt("Количество раундов [200]: ", 1, 10000, 200)
			noise = readFloat("Шум канала (0.0-1.0) [0.0]: ", 0.0, 1.0, 0.0)
		}

		switch mode {
		case 1:
			handleDuel(rounds, noise)
		case 2:
			handleRoundRobin(rounds, noise)
		case 3:
			handleTrial(rounds, noise)
		case 4:
			handleCirculaire(rounds, noise)
		case 5:
			handleArena(rounds, noise)
		case 6:
			handleEcosystem(rounds, noise)
		case 7:
			handleInfo()
			continue
		}
	}
}

func readInt(prompt string, min, max int, defaultVal int) int {
	fmt.Print(prompt)
	var n int
	_, err := fmt.Scanf("%d", &n)

	if err != nil || n < min || n > max {
		fmt.Scanln()
		return defaultVal
	}
	return n
}

func readFloat(prompt string, min, max float64, defaultVal float64) float64 {
	fmt.Print(prompt)
	var f float64
	_, err := fmt.Scanf("%f", &f)

	if err != nil || f < min || f > max {
		fmt.Scanln()
		return defaultVal
	}
	return f
}
