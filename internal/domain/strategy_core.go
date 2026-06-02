package domain

import "math/rand/v2"

func (s *Strategy) CoreDecision(history Rounds) Act {
	opLast := history.OpLastAct()
	myLast := history.MyLastAct()

	act := s.evaluate(s.Neutral, opLast, myLast)

	if s.Trigger != nil {

		//default value
		if s.Trigger.Count == 0 {
			s.Trigger.Count = 1
		}

		if test, valid := triggerTable[s.Trigger.Mode]; valid {
			if test(history, s.Trigger.Act, s.Trigger.Count) {
				return s.evaluate(s.Trigger.Reaction, opLast, myLast)
			}
		}
	}

	return act
}

// Решение стратегии
// fixAct → evaluateMirror → evaluateProb → evaluateState → evaluateTrigger
func (s *Strategy) evaluate(rule RuleValue, opLast, myLast Act) Act {
	act := s.evaluateFix(rule)
	act = s.evaluateMirror(act, rule, opLast, myLast)
	act = s.evaluateProb(act, rule)
	return act
}

// Если в реальном столкновении возвращается NoAct - некорректная стратегия
func (s *Strategy) evaluateFix(rule RuleValue) Act {
	// При нормальном поведении NoAct не дойдет до CoreDecision
	return rule.Fix
}
func (s *Strategy) evaluateMirror(act Act, rule RuleValue, opLast, myLast Act) Act {
	if rule.Mirror == nil || opLast == NoAct {
		return act
	}

	mode := *rule.Mirror

	if mode == MirrorSelf {
		return myLast
	}

	if mappedAct, valid := mirrorTable[mode][opLast]; valid {
		return mappedAct
	}

	return act
}
func (s *Strategy) evaluateProb(act Act, rule RuleValue) Act {
	if len(rule.Prob) == 0 {
		return act
	}

	r := rand.Float64()
	cumulative := 0.0
	for a, prob := range rule.Prob {
		cumulative += prob
		if r <= cumulative {
			return a
		}
	}

	return act
}
