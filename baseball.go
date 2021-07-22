package main

import (
	"fmt"
	"strings"
)

type Game struct {
	numbers []string
}

type Score struct {
	strike int
	ball   int
}

func (g *Game) compare(s string) Score {
	score := Score{}
	targetSlice := strings.Split(s, NumberSeparator)

	for i, source := range g.numbers {
		for k, target := range targetSlice {
			if source == target {
				if i == k {
					score.strike++
					continue
				}
				score.ball++
			}
		}
	}
	return score
}

const (
	NumberSeparator = ""
	MinNumber       = "1"
	MaxNumber       = "9"
	NumberLength    = 3
)

func NewGame(s string) (*Game, error) {
	numbers := strings.Split(s, NumberSeparator)
	if invalidLength(numbers) {
		return nil, fmt.Errorf("invalid length: %s", numbers)
	}

	used := make(map[string]struct{})
	for _, v := range numbers {
		if numberOutOfRange(v) {
			return nil, fmt.Errorf("invalid numbers: %s", numbers)
		}
		if _, ok := used[v]; ok {
			return nil, fmt.Errorf("duplicated numbers: %s", numbers)
		}
		used[v] = struct{}{}
	}
	return &Game{numbers}, nil
}

func invalidLength(numbers []string) bool {
	return len(numbers) != NumberLength
}

func numberOutOfRange(v string) bool {
	return v < MinNumber || v > MaxNumber
}
