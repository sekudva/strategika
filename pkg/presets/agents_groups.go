package presets

import (
	"github.com/sekudva/strategika/internal/domain"
)

func AllStrategies() []*domain.Agent {
	return []*domain.Agent{
		TitForTat(),
		ForgivingTFT(),
		AlwaysShare(),
		AlwaysTake(),
		Joss(),
		HardJoss(),
		TitFor2Tats(),
		Grudger(),
		EvilTFT(),
		Random(),
		GroFman(),
		GoByMajority(),
		WMAdams(),
		Eatherley(),
		Cave(),
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
		Journalist(),
		AntiGrudger(),
		Paladin(),
		Barbarian(),
		Warlock(),
		Sorcerer(),
		Monk(),
		Boev(),
	}
}
