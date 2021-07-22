package main

import (
	"fmt"
	"strings"
)

type Game struct {
	Quiz string
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
	for _, v := range numbers {
		if numberOutOfRange(v) {
			return nil, fmt.Errorf("invalid numbers: %s", numbers)
		}
	}
	return &Game{s}, nil
}

func invalidLength(numbers []string) bool {
	return len(numbers) != NumberLength
}

func numberOutOfRange(v string) bool {
	return v < MinNumber || v > MaxNumber
}
