package main

import (
	"fmt"
	"strings"
)

type Game struct {
	Quiz string
}

func NewGame(s string) (*Game, error) {
	numbers := strings.Split(s, "")
	if len(numbers) != 3 {
		return nil, fmt.Errorf("invalid length: %s", numbers)
	}
	for _, v := range numbers {
		if v < "1" || v > "9" {
			return nil, fmt.Errorf("invalid numbers: %s", numbers)
		}
	}
	return &Game{s}, nil
}
