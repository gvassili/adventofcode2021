package day01

import (
	"bufio"
	"io"
	"strconv"
)

type Challenge struct {
	input []int
}

func (c *Challenge) Day() int {
	return 1
}

func (c *Challenge) Prepare(r io.Reader) error {
	c.input = make([]int, 0, 1024)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		c.input = append(c.input, n)
	}
	return scanner.Err()
}

func (c *Challenge) Part1() (string, error) {
	count := 0
	for i, n := range c.input {
		if i > 0 && n > c.input[i-1] {
			count++
		}
	}
	return strconv.Itoa(count), nil
}

func (c *Challenge) Part2() (string, error) {
	count := 0
	for i, n := range c.input {
		if i > 2 && c.input[i-3] < n {
			count++
		}
	}
	return strconv.Itoa(count), nil
}
