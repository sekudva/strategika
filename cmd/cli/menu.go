package cli

import (
	"fmt"
)

type menuItem struct {
	id   int
	name string
	desc string
}

var modes = []menuItem{
	{1, "Duel", "1 × 1"},
	{2, "Round Robin", "1 × 1 many times, TOURNAMENT"},
	{3, "Trial", "1 × group"},
	{4, "Circulaire", "1 × group many times"},
	{5, "Arena", "all × all"},
	{6, "Ecosystem Arena", "all × all WITH dying threshold"},
	{7, "INFO", "some info about cli menu"},
}

func MainMenu() int {
	fmt.Println("\n=== STRATEGIKA ===")
	fmt.Println("\nChoose mode:")
	for _, m := range modes {
		fmt.Printf("%-1d) %-18s |%-30s\n", m.id, m.name, m.desc)
	}
	fmt.Println("\nAny other option is ESCAPE.")
	fmt.Println("Any other next options besides the ones provided returns DEFAULT values.")
	return readInt("> ", 0, len(modes), 0)
}
