package main

import (
	"baseball/number"
	"baseball/score"
	"strings"
)

type Game struct {
	numbers number.Numbers
}

func (g *Game) compare(s string) score.Score {
	sc := score.Score{}
	targetSlice := strings.Split(s, number.Separator)

	for index, value := range targetSlice {
		result := g.numbers.Contains(value, index)
		sc.Add(result)
	}
	return sc
}

func NewGame(s string) (*Game, error) {
	numbers, err := number.NewNumbers(s)
	if err != nil {
		return nil, err
	}
	return &Game{numbers}, nil
}
