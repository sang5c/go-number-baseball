package number

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Separator = ""
	MinNumber = "1"
	MaxNumber = "9"
	Length    = 3
)

type Numbers []Number

type Number struct {
	value    int
	position int
}

func NewNumbers(str string) (Numbers, error) {
	split := strings.Split(str, Separator)
	if invalidLength(split) {
		return Numbers{}, fmt.Errorf("invalid length: %s", str)
	}

	numbers := Numbers{}
	used := make(map[string]struct{})
	for index, value := range split {
		number, err := NewNumber(value, index)
		if err != nil {
			return Numbers{}, err
		}
		if _, ok := used[value]; ok {
			return nil, fmt.Errorf("duplicated numbers: %v", numbers)
		}
		used[value] = struct{}{}
		numbers = append(numbers, number)
	}
	return numbers, nil
}

func (numbers Numbers) Contains(target string, position int) Decision {
	for _, number := range numbers {
		targetNumber, _ := NewNumber(target, position)
		result := number.compare(targetNumber)
		if isBallOrStrike(result) {
			return result
		}
	}
	return nothing
}

func invalidLength(numbers []string) bool {
	return len(numbers) != Length
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

func isBallOrStrike(result Decision) bool {
	return result != nothing
}

func (n Number) compare(target Number) Decision {
	if isStrike(target, n) {
		return strike
	} else if isBall(target, n) {
		return ball
	}
	return nothing
}

func isBall(target Number, n Number) bool {
	return n.value == target.value && n.position != target.position
}

func isStrike(target Number, n Number) bool {
	return n.value == target.value && n.position == target.position
}
