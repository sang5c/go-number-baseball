package score

import (
	"baseball/number"
)

type Score struct {
	strike int
	ball   int
}

func (s *Score) Add(d number.Decision) {
	result := d.GetResult()
	if result == "strike" {
		s.strike++
	} else if result == "ball" {
		s.ball++
	}
}

func New(strike, ball int) Score {
	return Score{strike, ball}
}
