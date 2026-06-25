package agents

import (
	"github.com/sekudva/strategika/internal/domain"
)

// ClassicStrategies возвращает набор классических стратегий(с новой логикой) из турнира Аксельрода
func Classic() []*domain.Agent {
	return []*domain.Agent{
		TitForTat(),
		TitFor2Tats(),
		ForgivingTFT(),
		EvilTFT(),
		Joss(),
		HardJoss(),

		AlwaysShare(),
		AlwaysTake(),
		Random(),

		GroFman(),

		GoByMajority(),
		Grudger(),
		WMAdams(),
		Eatherley(),
		Cave(),
		Pavlov(),
		Tullock(),
		Champion(),
		Leyvraz(),
	}
}

// Неклассические стратегии, использующие необычную логику из конструкторов и частое действие Hold
func NonClassic() []*domain.Agent {
	return []*domain.Agent{
		AlwaysHold(),

		Masochist(),
		Sadist(),
		Pacifist(),
		Bully(),
		Melting(),

		EvilCyclist(),
		GoodCyclist(),
		ColdCyclist(),
		WarmCyclist(),
		ColdGhost(),

		Crazy(),
		Unfriendly(),
		ZeroHold(),
		StrictTeacher(),
		Patient(),

		AntiGrudger(),

		Journalist(),
		Boev(),
		Abuser(),
		EvilPavlov(),
		Sycheva(),

		Paladin(),
		Barbarian(),
		Warlock(),
		Sorcerer(),
		Monk(),
	}
}

// Evil Group
func Evil() []*domain.Agent {
	return []*domain.Agent{
		EvilTFT(),
		Joss(),
		HardJoss(),

		AlwaysHold(),
		AlwaysTake(),
		Random(),
		RandomOLD(),

		Sadist(),
		Bully(),
		Melting(),

		EvilCyclist(),
		ColdCyclist(),
		ColdGhost(),

		Unfriendly(),

		AntiGrudger(),

		Boev(),
		Abuser(),
		EvilPavlov(),

		Barbarian(),
		Warlock(),
		Sorcerer(),
		Rogue(),

		Cave(),
	}
}

// Nice Group
func Nice() []*domain.Agent {
	return []*domain.Agent{
		TitFor2Tats(),
		ForgivingTFT(),

		AlwaysShare(),

		Pacifist(),

		GoodCyclist(),
		WarmCyclist(),

		StrictTeacher(),
		Patient(),

		GroFman(),

		Paladin(),
		Monk(),

		GoByMajority(),
		WMAdams(),
		Eatherley(),
		Pavlov(),
		Tullock(),
		Champion(),
		Leyvraz(),
	}
}

// Adaptable — стратегии, которые адаптируются к оппоненту.
func Adaptable() []*domain.Agent {
	return []*domain.Agent{
		Boev(),
		Abuser(),
		EvilPavlov(),
		Sycheva(),

		Rogue(),

		GoByMajority(),
		WMAdams(),
		Eatherley(),
		Cave(),
		Pavlov(),
		Tullock(),
		Champion(),
		Leyvraz(),
	}
}
