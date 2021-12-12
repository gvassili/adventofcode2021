package calendar

import (
	"fmt"
	"github.com/gvassili/adventofcode2021/calendar/day01"
	"github.com/gvassili/adventofcode2021/calendar/day02"
	"github.com/gvassili/adventofcode2021/calendar/day03"
	"github.com/gvassili/adventofcode2021/calendar/day04"
	"github.com/gvassili/adventofcode2021/calendar/day05"
	"github.com/gvassili/adventofcode2021/calendar/day06"
	"github.com/gvassili/adventofcode2021/calendar/day07"
	"github.com/gvassili/adventofcode2021/calendar/day08"
	"github.com/gvassili/adventofcode2021/calendar/day09"
	"io"
	"sort"
)

var challengeList = []func() Challenge{
	func() Challenge { return new(day01.Challenge) },
	func() Challenge { return new(day02.Challenge) },
	func() Challenge { return new(day03.Challenge) },
	func() Challenge { return new(day04.Challenge) },
	func() Challenge { return new(day05.Challenge) },
	func() Challenge { return new(day06.Challenge) },
	func() Challenge { return new(day07.Challenge) },
	func() Challenge { return new(day08.Challenge) },
	func() Challenge { return new(day09.Challenge) },
}

var challengeMap = func() map[int]func() Challenge {
	m := make(map[int]func() Challenge, len(challengeList))
	for _, c := range challengeList {
		m[c().Day()] = c
	}
	return m
}()

func Load(day int) (Challenge, error) {
	loader, ok := challengeMap[day]
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
	challengeNames := make([]int, 0, len(challengeMap))
	for name := range challengeMap {
		challengeNames = append(challengeNames, name)
	}
	sort.Ints(challengeNames)
	result := make([]Challenge, 0, len(challengeMap))
	for _, day := range challengeNames {
		result = append(result, challengeMap[day]())
	}
	return result
}
