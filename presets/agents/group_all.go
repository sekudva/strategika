package agents

import (
	"github.com/sekudva/strategika/internal/domain"
)

func AllStrategies() []*domain.Agent {
	return []*domain.Agent{
		TitForTat(),
		TitForTatOrig(),
		TitFor2Tats(),
		TitFor2TatsOrig(),
		ForgivingTFT(),
		EvilTFT(),
		Joss(),
		HardJoss(),

		AlwaysShare(),
		AlwaysHold(),
		AlwaysTake(),
		Random(),
		RandomOLD(),

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
		GroFman(),

		Journalist(),
		Boev(),
		Abuser(),
		EvilPavlov(),
		Sycheva(),

		Paladin(),
		Barbarian(),
		Warlock(),
		Sorcerer(),
		Rogue(),
		Monk(),

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
