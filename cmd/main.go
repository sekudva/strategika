package main

import (
	"fmt"
	"log"

	"github.com/sekudva/strategika/pkg/benchmark"
)

func main() {
	// ========== ВЫБЕРИ ТЕСТ ==========
	// Раскомментируй нужную строку:

	// 1. Круговой турнир всех стратегий
	// if err := benchmark.RunRoundRobinTournament(); err != nil {
	// 	log.Fatal(err)
	// }

	// 2. Дуэль двух агентов (выбери агентов внутри benchmarks.RunDuel)
	if err := benchmark.RunDuel(); err != nil {
		log.Fatal(err)
	}

	// 3. Арена с группой (выбери группу внутри benchmarks.RunArena)
	// if err := benchmark.RunArena(); err != nil {
	//     log.Fatal(err)
	// }

	fmt.Println("Done!")
}
