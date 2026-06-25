package agents

import (
	"github.com/sekudva/strategika/internal/domain"
	"github.com/sekudva/strategika/presets/strategies"
)

var gen = domain.NewIDGenerator()

func NextID() domain.AgID {
	return gen.Next()
}

func RequestID(id domain.AgID) domain.AgID {
	return gen.Request(id)
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
func ZeroHold() *domain.Agent {
	return &domain.Agent{
		Name:     "ZeroHold",
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
