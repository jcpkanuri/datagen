package types

import (
	"errors"
)

type SequenceMap map[string]int

type SequenceConfig struct {
	Start int
	Step  int
	Max   int
	Cycle bool
}

func (s SequenceMap) IncrementAndGet(key string, seqCon SequenceConfig) (int, error) {

	if val, exists := s[key]; exists {
		val += seqCon.Step
		if val == seqCon.Max && seqCon.Cycle {
			val = seqCon.Start
		}

		s[key] = val
		return val, nil
	}

	return 0, errors.New("key does not exist")
}
