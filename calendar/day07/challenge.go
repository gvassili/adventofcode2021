package day07

import (
	"bufio"
	"bytes"
	"io"
	"sort"
	"strconv"
)

type Challenge struct {
	input []int
	max   int
}

func (c *Challenge) Day() int {
	return 7
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
		c.max = max(c.max, n)
		c.input = append(c.input, n)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func sumInts(elems []int) (sum int) {
	for _, e := range elems {
		sum += e
	}
	return
}

func (c *Challenge) Part1() (string, error) {
	sort.Ints(c.input)
	consumption := sumInts(c.input)
	prev := c.input[0]
	for i, n := range c.input {
		if n == prev {
			continue
		}
		diff := n - prev
		nextConsumption := consumption + (i*diff - (len(c.input)-i)*diff)
		if nextConsumption > consumption {
			return strconv.Itoa(consumption), nil
		}
		consumption = nextConsumption
		prev = n
	}
	return strconv.Itoa(consumption), nil
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (c *Challenge) Part2() (string, error) {
	computeConsumption := func(p int) int {
		consumption := 0
		for _, n := range c.input {
			delta := abs(n - p)
			consumption += (delta * (delta + 1)) >> 1
			//		fmt.Printf("%d => %d (%d) == %d\n", n, idx, delta, (delta*(delta+1))>>1)
		}
		return consumption
	}
	return strconv.Itoa(computeConsumption(sort.Search(c.max, func(p int) bool {
		left, right := computeConsumption(p), computeConsumption(p+1)
		return left <= right
	}))), nil
}
