package agents

import (
	"github.com/sekudva/strategika/internal/domain"
	"github.com/sekudva/strategika/pkg/presets/strategies"
)

var gen = domain.NewIDGenerator()

func NextID() domain.AgID {
	return gen.Next()
}

func RequestID(id domain.AgID) domain.AgID {
	return gen.Request(id)
}

// ========== 1. TIT-FOR-TAT FAMILY ==========

func TitForTat() *domain.Agent {
	return &domain.Agent{
		Name:     "TitForTat",
		ID:       RequestID(1),
		Strategy: strategies.TitForTatNEW(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

func TitFor2Tats() *domain.Agent {
	return &domain.Agent{
		Name:     "TitFor2Tats",
		ID:       RequestID(2),
		Strategy: strategies.TitFor2Tats(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

func Joss() *domain.Agent {
	return &domain.Agent{
		Name:     "Joss",
		ID:       RequestID(3),
		Strategy: strategies.Joss(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

func ForgivingTitForTat() *domain.Agent {
	return &domain.Agent{
		Name:     "ForgivingTFT",
		ID:       RequestID(4),
		Strategy: strategies.ForgivingTitForTat(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

func Tullock() *domain.Agent {
	return &domain.Agent{
		Name:     "EvilTFT",
		ID:       RequestID(5),
		Strategy: strategies.TitForTatEVIL(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// ========== 2. ALWAYS FAMILY ==========

func AlwaysShare() *domain.Agent {
	return &domain.Agent{
		Name:     "AlwaysShare",
		ID:       RequestID(10),
		Strategy: strategies.AlwaysShare(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

func AlwaysHold() *domain.Agent {
	return &domain.Agent{
		Name:     "AlwaysHold",
		ID:       RequestID(11),
		Strategy: strategies.AlwaysHold(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

func AlwaysTake() *domain.Agent {
	return &domain.Agent{
		Name:     "AlwaysTake",
		ID:       RequestID(12),
		Strategy: strategies.AlwaysTake(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

func Random() *domain.Agent {
	return &domain.Agent{
		Name:     "Random",
		ID:       RequestID(13),
		Strategy: strategies.Random(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// ========== 3. BDSM FAMILY  ==========

func Masochist() *domain.Agent {
	return &domain.Agent{
		Name:     "Masochist",
		ID:       RequestID(20),
		Strategy: strategies.Masochist(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

func Sadist() *domain.Agent {
	return &domain.Agent{
		Name:     "Sadist",
		ID:       RequestID(21),
		Strategy: strategies.Sadist(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

func Pacifist() *domain.Agent {
	return &domain.Agent{
		Name:     "Pacifist",
		ID:       RequestID(22),
		Strategy: strategies.Pacifist(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

func Bully() *domain.Agent {
	return &domain.Agent{
		Name:     "Bully",
		ID:       RequestID(30),
		Strategy: strategies.Bully(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// ========== 4. FUN FAMILY ==========

func Crazy() *domain.Agent {
	return &domain.Agent{
		Name:     "Crazy",
		ID:       RequestID(23),
		Strategy: strategies.Crazy(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

func Unfriendly() *domain.Agent {
	return &domain.Agent{
		Name:     "Unfriendly",
		ID:       RequestID(31),
		Strategy: strategies.Unfriendly(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

func Irreconcilable() *domain.Agent {
	return &domain.Agent{
		Name:     "Irreconcilable",
		ID:       RequestID(32),
		Strategy: strategies.Irreconcilable(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}
