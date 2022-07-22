package day11

import (
	"bufio"
	"errors"
	"io"
	"strconv"
)

type Challenge struct {
	input []int
}

func (c *Challenge) Day() int {
	return 11
}

const size = 10

func (c *Challenge) Prepare(r io.Reader) error {
	c.input = make([]int, 0, size*size)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		for _, v := range scanner.Text() {
			c.input = append(c.input, int(v-'0'))
		}
	}
	return scanner.Err()
}

func coordToIndex(x, y int) int {
	return x + y*size
}

func incRec(flashed []bool, input []int, x, y int) int {
	flashCount := 0
	index := coordToIndex(x, y)
	if x < 0 || x >= size || y < 0 || y >= size || flashed[index] {
		return 0
	}
	input[index]++
	if input[index] > 9 {
		flashed[index] = true
		flashCount++
		input[index] = 0
		flashCount += incRec(flashed, input, x, y+1)
		flashCount += incRec(flashed, input, x+1, y+1)
		flashCount += incRec(flashed, input, x+1, y)
		flashCount += incRec(flashed, input, x+1, y-1)
		flashCount += incRec(flashed, input, x, y-1)
		flashCount += incRec(flashed, input, x-1, y-1)
		flashCount += incRec(flashed, input, x-1, y)
		flashCount += incRec(flashed, input, x-1, y+1)
	}
	return flashCount
}

func (c *Challenge) Part1() (string, error) {
	flashCount := 0
	input := make([]int, len(c.input))
	var flashed []bool
	copy(input, c.input)
	for step := 0; step < 100; step++ {
		flashed = make([]bool, len(input))
		for i := range input {
			flashCount += incRec(flashed, input, i%size, i/size)
		}
	}
	return strconv.Itoa(flashCount), nil
}

func (c *Challenge) Part2() (string, error) {
	input := make([]int, len(c.input))
	var flashed []bool
	copy(input, c.input)
	for step := 0; step < 1000000; step++ {
		flashCount := 0
		flashed = make([]bool, len(input))
		for i := range input {
			flashCount += incRec(flashed, input, i%size, i/size)
		}
		if flashCount >= size*size {
			return strconv.Itoa(step + 1), nil
		}
	}
	return "", errors.New("answer not found")
}
