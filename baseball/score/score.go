package score

import (
	number2 "baseball/baseball/number"
)

type Score struct {
	strike int
	ball   int
}

func (s *Score) Add(d number2.Decision) {
	if d.IsStrike() {
		s.strike++
	} else if d.IsBall() {
		s.ball++
	}
}

func New(strike, ball int) Score {
	return Score{strike, ball}
}