package calendar

import (
	"fmt"
	"github.com/gvassili/adventofcode2021/calendar/day01"
	"github.com/gvassili/adventofcode2021/calendar/day02"
	"github.com/gvassili/adventofcode2021/calendar/day03"
	"io"
	"sort"
)

var challenges = map[int]func() Challenge{
	1: func() Challenge { return new(day01.Challenge) },
	2: func() Challenge { return new(day02.Challenge) },
	3: func() Challenge { return new(day03.Challenge) },
}

func Load(day int) (Challenge, error) {
	loader, ok := challenges[day]
	if !ok {
		return nil, fmt.Errorf("could not find challenge %d", day)
	}
	return loader(), nil
}

type Challenge interface {
	Day() int
	Prepare(r io.Reader) error
	Part1() (string, error)
	Part2() (string, error)
}

func LoadAllChallenges() []Challenge {
	challengeNames := make([]int, 0, len(challenges))
	for name := range challenges {
		challengeNames = append(challengeNames, name)
	}
	sort.Ints(challengeNames)
	result := make([]Challenge, 0, len(challenges))
	for _, day := range challengeNames {
		result = append(result, challenges[day]())
	}
	return result
}
