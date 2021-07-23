package number

import (
	"fmt"
	"strconv"
)

const (
	Separator = ""
	MinNumber = "1"
	MaxNumber = "9"
)

type Number struct {
	value    int
	position int
}

func numberOutOfRange(v string) bool {
	return v < MinNumber || v > MaxNumber
}

func NewNumber(s string, position int) (Number, error) {
	if numberOutOfRange(s) {
		return Number{}, fmt.Errorf("invalid number: %s", s)
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return Number{}, err
	}
	return Number{v, position}, nil
}

func (n Number) compare(target Number) Decision {
	if isStrike(target, n) {
		return strike
	} else if isBall(target, n) {
		return ball
	}
	return nothing
}

func isBallOrStrike(result Decision) bool {
	return result != nothing
}

func isBall(target Number, n Number) bool {
	return n.value == target.value && n.position != target.position
}

func isStrike(target Number, n Number) bool {
	return n.value == target.value && n.position == target.position
}
