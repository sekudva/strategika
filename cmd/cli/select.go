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
	{"Все", presets.AllStrategies},
	{"Классические", presets.ClassicStrategies},
	{"Неклассические", presets.NonClassicGroup},
}

func selectGroup() []*domain.Agent {
	fmt.Println("\nВыберите группу стратегий:")
	for i, g := range groups {
		fmt.Printf("%d. %s\n", i+1, g.name)
	}

	choice := readInt("> ", 0, len(groups), 1)
	if choice == 0 {
		return nil
	}

	return groups[choice-1].fn()
}

func selectEach(agents []*domain.Agent) []*domain.Agent {
	fmt.Println("\nДоступные агенты:")
	for i, a := range agents {
		fmt.Printf("%2d. %s\n", i+1, a.Name)
	}

	fmt.Print("\nЕсли ничего не выбрать, то группа будет в полном составе!")
	fmt.Print("\nВведите номера через запятую:\n> ")

	var input string
	fmt.Scanf("%s", &input)

	var selected []*domain.Agent
	for _, s := range strings.Split(input, ",") {
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err == nil && n >= 1 && n <= len(agents) {
			selected = append(selected, agents[n-1])
		}
	}

	if len(selected) == 0 {
		fmt.Println("Ничего не выбрано — беру всех.")
		return agents
	}

	return selected
}

func selectOne(agents []*domain.Agent) *domain.Agent {
	fmt.Print("\nЕсли ничего не выбрать, то будет выбран первый в группе!")
	fmt.Println("\nВыберите агента:")
	for i, a := range agents {
		fmt.Printf("%2d. %s\n", i+1, a.Name)
	}

	choice := readInt("> ", 0, len(agents), 0)

	if choice == 0 {
		return agents[0]
	}

	return agents[choice-1]
}
