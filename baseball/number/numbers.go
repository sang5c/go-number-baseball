package number

import (
	"fmt"
	"strings"
)

const (
	Length = 3
)

type Numbers []Number

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

func invalidLength(numbers []string) bool {
	return len(numbers) != Length
}
