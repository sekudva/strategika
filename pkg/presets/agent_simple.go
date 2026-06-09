package presets

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

// just copy
func TitForTat() *domain.Agent {
	return &domain.Agent{
		Name:     "TitForTat",
		ID:       RequestID(10),
		Strategy: strategies.TitForTatNEW(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// triggers only on take
func TitForTatOrig() *domain.Agent {
	return &domain.Agent{
		Name:     "TitForTatOrig",
		ID:       RequestID(11),
		Strategy: strategies.TitForTat(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// triggers on 2 takes AND to 1 take - do hold
func TitFor2Tats() *domain.Agent {
	return &domain.Agent{
		Name:     "TitFor2Tats",
		ID:       RequestID(12),
		Strategy: strategies.TitFor2Tats(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// triggers only on 2 takes
func TitFor2TatsOrig() *domain.Agent {
	return &domain.Agent{
		Name:     "TitFor2TatsOLD",
		ID:       RequestID(13),
		Strategy: strategies.TitFor2TatsOLD(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// not always triggers on take
func ForgivingTFT() *domain.Agent {
	return &domain.Agent{
		Name:     "ForgivingTFT",
		ID:       RequestID(12),
		Strategy: strategies.ForgivingTitForTat(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// start with take
func EvilTFT() *domain.Agent {
	return &domain.Agent{
		Name:     "EvilTFT",
		ID:       RequestID(13),
		Strategy: strategies.TitForTatEVIL(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// take prob
func Joss() *domain.Agent {
	return &domain.Agent{
		Name:     "Joss",
		ID:       RequestID(14),
		Strategy: strategies.Joss(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// take prob, take start
func HardJoss() *domain.Agent {
	return &domain.Agent{
		Name:     "HardJoss",
		ID:       RequestID(15),
		Strategy: strategies.HardJoss(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// ========== 2. ALWAYS FAMILY ==========

func AlwaysShare() *domain.Agent {
	return &domain.Agent{
		Name:     "AlwaysShare",
		ID:       RequestID(20),
		Strategy: strategies.AlwaysShare(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

func AlwaysHold() *domain.Agent {
	return &domain.Agent{
		Name:     "AlwaysHold",
		ID:       RequestID(21),
		Strategy: strategies.AlwaysHold(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

func AlwaysTake() *domain.Agent {
	return &domain.Agent{
		Name:     "AlwaysTake",
		ID:       RequestID(22),
		Strategy: strategies.AlwaysTake(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

func Random() *domain.Agent {
	return &domain.Agent{
		Name:     "Random",
		ID:       RequestID(23),
		Strategy: strategies.Random(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

func RandomOLD() *domain.Agent {
	return &domain.Agent{
		Name:     "RandomOLD",
		ID:       RequestID(24),
		Strategy: strategies.Random_OLD(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// ========== 3. BDSM FAMILY  ==========

// Likes to be forced taken
func Masochist() *domain.Agent {
	return &domain.Agent{
		Name:     "Masochist",
		ID:       RequestID(30),
		Strategy: strategies.Masochist(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// Sometimes do take and never share
func Sadist() *domain.Agent {
	return &domain.Agent{
		Name:     "Sadist",
		ID:       RequestID(31),
		Strategy: strategies.Sadist(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// never do take, only hold
func Pacifist() *domain.Agent {
	return &domain.Agent{
		Name:     "Pacifist",
		ID:       RequestID(32),
		Strategy: strategies.Pacifist(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// never share, only take, response to take is hold
func Bully() *domain.Agent {
	return &domain.Agent{
		Name:     "Bully",
		ID:       RequestID(33),
		Strategy: strategies.Bully(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// always take, but if share - answer share
func Melting() *domain.Agent {
	return &domain.Agent{
		Name:     "Melting",
		ID:       RequestID(34),
		Strategy: strategies.Melting(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// ========== 4. CYCLE FAMILY ==========

// always take, but if two shares - do share
func EvilCyclist() *domain.Agent {
	return &domain.Agent{
		Name:     "EvilCyclist",
		ID:       RequestID(40),
		Strategy: strategies.EvilCyclist(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// answer take only to 2 takes
func GoodCyclist() *domain.Agent {
	return &domain.Agent{
		Name:     "GoodCyclist",
		ID:       RequestID(41),
		Strategy: strategies.GoodCyclist(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// never share, react only if 2 takes
func ColdCyclist() *domain.Agent {
	return &domain.Agent{
		Name:     "ColdCyclist",
		ID:       RequestID(42),
		Strategy: strategies.ColdCyclist(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// likes to share, but every 2 hold do hold, on take do take
func WarmCyclist() *domain.Agent {
	return &domain.Agent{
		Name:     "WarmCyclist",
		ID:       RequestID(43),
		Strategy: strategies.WarmCyclist(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// share if only 2 shares, on take do take
func ColdGhost() *domain.Agent {
	return &domain.Agent{
		Name:     "ColdGhost",
		ID:       RequestID(44),
		Strategy: strategies.ColdGhost(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// ========== 5. FUN FAMILY ==========

// on share do take, on take do share
func Crazy() *domain.Agent {
	return &domain.Agent{
		Name:     "Crazy",
		ID:       RequestID(50),
		Strategy: strategies.Crazy(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// never coop
func Unfriendly() *domain.Agent {
	return &domain.Agent{
		Name:     "Unfriendly",
		ID:       RequestID(51),
		Strategy: strategies.Unfriendly(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// on hold do take
func Irreconcilable() *domain.Agent {
	return &domain.Agent{
		Name:     "Irreconcilable",
		ID:       RequestID(52),
		Strategy: strategies.Irreconcilable(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// for two Takes - one Take and forgives
func StrictTeacher() *domain.Agent {
	return &domain.Agent{
		Name:     "StrictTeacher",
		ID:       RequestID(53),
		Strategy: strategies.StrictTeacher(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// Only after 3 Takes do Take, TF3T
func Patient() *domain.Agent {
	return &domain.Agent{
		Name:     "Patient",
		ID:       RequestID(54),
		Strategy: strategies.Patient(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// ========== 6. ORIGINAL SIMPLE LOGIC FAMILY ==========

// after 2 shares forever share, starts with take
func AntiGrudger() *domain.Agent {
	return &domain.Agent{
		Name:     "AntiGrudger",
		ID:       RequestID(60),
		Strategy: strategies.AntiGrudger(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

func GroFman() *domain.Agent {
	return &domain.Agent{
		Name:     "GroFman",
		ID:       RequestID(61),
		Strategy: strategies.GroFman(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}
