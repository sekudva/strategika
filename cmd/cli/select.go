package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sekudva/strategika/internal/domain"
	"github.com/sekudva/strategika/pkg/presets"
)

var groups = []struct {
	name string
	fn   func() []*domain.Agent
}{
	{"All agents", presets.AllStrategies},
	{"Classic", presets.ClassicStrategies},
	{"Non-Classic", presets.NonClassicGroup},
}

func selectGroup() []*domain.Agent {
	fmt.Println("\nChoose group:")
	for i, g := range groups {
		fmt.Printf("%d. %s\n", i+1, g.name)
	}

	choice := readInt("> ", 0, len(groups), 1)
	if choice == 0 {
		return nil
	}

	selected := groups[choice-1].fn()
	fmt.Printf("Chosen group: «%s», numbers of agents: %d\n", groups[choice-1].name, len(selected))
	return selected
}

func selectEach(agents []*domain.Agent) []*domain.Agent {
	fmt.Println("\nAvialable agents:")
	for i, a := range agents {
		fmt.Printf("%2d. %s\n", i+1, a.Name)
	}

	fmt.Print("\nIf nothing chosen, the whole group will be returned!")
	fmt.Print("\nEnter agent numbers separated by space:\n> ")

	var input string
	fmt.Scanf("%s", &input)

	var selected []*domain.Agent
	for _, s := range strings.Split(input, " ") {
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err == nil && n >= 1 && n <= len(agents) {
			selected = append(selected, agents[n-1])
		}
	}

	if len(selected) == 0 {
		fmt.Println("Nothing chosen - entire group returned.")
		return agents
	}

	return selected
}

func selectOne(agents []*domain.Agent) *domain.Agent {
	fmt.Print("\nIf nothing chosen, first agent will be returned!")
	fmt.Println("\nChoose agent:")
	for i, a := range agents {
		fmt.Printf("%2d. %s\n", i+1, a.Name)
	}

	choice := readInt("> ", 0, len(agents), 0)

	if choice == 0 {
		fmt.Println("Nothing chosen - first agent returned.")
		return agents[0]
	}

	return agents[choice-1]
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
