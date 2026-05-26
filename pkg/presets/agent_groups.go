// pkg/presets/agent_groups.go

package presets

import (
	"github.com/sekudva/strategika/internal/domain"
)

// ClassicStrategies возвращает набор классических стратегий(с новой логикой) из турнира Аксельрода
// которые не используют модификаторы
func ClassicStrategies() []*domain.Agent {
	return []*domain.Agent{
		TitForTat(),
		AlwaysShare(),
		AlwaysTake(),
		Joss(),
		TitFor2Tats(),
		Grudger(),
		EvilTFT(),
		Random(),
	}
}

// AggressiveGroup возвращает стратегии, склонные к атаке
func NonClassicGroup() []*domain.Agent {
	return []*domain.Agent{}
}
