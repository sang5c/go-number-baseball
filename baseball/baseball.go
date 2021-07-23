package baseball

import (
	number2 "baseball/baseball/number"
	score2 "baseball/baseball/score"
	"strings"
)

type Game struct {
	numbers number2.Numbers
}

func (g *Game) compare(s string) score2.Score {
	sc := score2.Score{}
	targetSlice := strings.Split(s, number2.Separator)

	for index, value := range targetSlice {
		result := g.numbers.Contains(value, index)
		sc.Add(result)
	}
	return sc
}

func NewGame(s string) (*Game, error) {
	numbers, err := number2.NewNumbers(s)
	if err != nil {
		return nil, err
	}
	return &Game{numbers}, nil
}
