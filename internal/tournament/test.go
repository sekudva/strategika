package tournament

// // !!!!!!!!!!!

// import (
// 	"testing"

// 	"github.com/sekudva/strategika/internal/domain"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// func TestDuel_AlwaysShareVsAlwaysTake(t *testing.T) {
// 	gen := domain.NewIDGenerator()
// 	a1 := domain.NewAgent(presets.AlwaysShare(), gen.Next())
// 	a2 := domain.NewAgent(presets.AlwaysTake(), gen.Next())

// 	duel := Duel{Rounds: 10, Noise: 0}
// 	res, err := duel.Run([]*domain.Agent{a1, a2})
// 	require.NoError(t, err)

// 	scores := res.Scores()
// 	// Share vs Take: Share получает -3, Take получает +7 каждый раунд
// 	assert.Equal(t, -30, scores[a1.ID])
// 	assert.Equal(t, 70, scores[a2.ID])
// }
