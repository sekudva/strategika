// cmd/strategika/main.go
package main

import (
	"fmt"
	"os"

	"github.com/sekudva/strategika/cmd/cli"
)

func main() {
	if err := cli.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// func main() {
// 	// ========== ВЫБЕРИ ТЕСТ ==========
// 	// Раскомментируй нужную строку:

// 	// 1. Круговой турнир всех стратегий
// 	// if err := benchmark.RunRoundRobinTournament(); err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// 2. Дуэль двух агентов (выбери агентов внутри benchmarks.RunDuel)
// 	if err := benchmark.RunDuel(); err != nil {
// 		log.Fatal(err)
// 	}

// 	// 3. Арена с группой (выбери группу внутри benchmarks.RunArena)
// 	// if err := benchmark.RunArena(); err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	fmt.Println("Done!")
// }
