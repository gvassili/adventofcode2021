package day06

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
)

type Challenge struct {
	input []int
}

func (c *Challenge) Day() int {
	return 6
}

func (c *Challenge) Prepare(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.IndexByte(data, ','); i >= 0 {
			return i + 1, data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	})

	c.input = make([]int, 0, 1024)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		c.input = append(c.input, n)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func simulatePopulation(basePopulation []int, duration int) int {
	var population [9]int
	for _, fish := range basePopulation {
		population[fish]++
	}
	for day := 0; day <= duration; day++ {
		population[(day+6)%len(population)] += population[(day+8)%len(population)]
	}
	sum := 0
	for _, pop := range population {
		sum += pop
	}
	return sum
}

func (c *Challenge) Part1() (string, error) {
	return strconv.Itoa(simulatePopulation(c.input, 80)), nil
}

func (c *Challenge) Part2() (string, error) {
	return strconv.Itoa(simulatePopulation(c.input, 256)), nil
}
