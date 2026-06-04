package tournament

// func (cfg SimConfig) RunSimulation(agents []*domain.Agent) map[domain.AgID]int {
// 	active := make([]*domain.Agent, len(agents))
// 	copy(active, agents)

// 	for round := 1; round <= cfg.Rounds; round++ {
// 		// Фильтруем живых (только если есть порог смерти)
// 		if cfg.DeathThreshold > math.MinInt {
// 			var alive []*domain.Agent
// 			for _, a := range active {
// 				if !a.Dead {
// 					alive = append(alive, a)
// 				}
// 			}
// 			active = alive
// 			if len(active) < 2 {
// 				break
// 			}
// 		}

// 		// Строим пары для активных агентов
// 		var activePairs []Pair
// 		for i := range active {
// 			for j := i + 1; j < len(active); j++ {
// 				activePairs = append(activePairs, Pair{i, j})
// 			}
// 		}

// 		decisions := decidePhase(active, activePairs, round)
// 		decisions = noisePhase(decisions, cfg.Noise, cfg.RNG)
// 		applyPhase(active, decisions, activePairs, round, cfg.Logger)

// 		// Проверка смерти
// 		if cfg.DeathThreshold > math.MinInt {
// 			for _, a := range active {
// 				if a.Score <= cfg.DeathThreshold {
// 					a.Dead = true
// 					cfg.Logger.Log(RoundLog{
// 						Round:  round,
// 						Agent1: a.ID,
// 						Act1:   domain.NoAct,
// 						Score1: a.Score,
// 					})
// 				}
// 			}
// 		}
// 	}

// 	if cfg.Logger != nil {
// 		cfg.Logger.Flush()
// 	}

// 	scores := make(map[domain.AgID]int, len(agents))
// 	for _, a := range agents {
// 		scores[a.ID] = a.Score
// 	}
// 	return scores
// }
