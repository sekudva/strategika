package tournament

import (
	"github.com/sekudva/strategika/internal/domain"
)

type Tournament interface {
	Run(agents []*domain.Agent, rounds int, noise float64) Result
}

type Result interface {
	Summary() string
}
