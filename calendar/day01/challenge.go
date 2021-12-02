package day01

import (
	"bufio"
	"io"
	"strconv"
)

type Challenge struct {
	input    []int
}

func (d *Challenge) Day() int {
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

func (d *Challenge) Part1() (string, error) {
	count := 0
	for i, n := range d.input {
		if i > 0 && n > d.input[i - 1] {
			count++
		}
	}
	return strconv.Itoa(count), nil
}

func (d *Challenge) Part2() (string, error) {
	count, slidingSum, prevSlidingSum := 0, 0, 0
	for i, n := range d.input {
	    slidingSum += n
	    if i > 2 {
	    	slidingSum -= d.input[i - 3]
	    	if slidingSum > prevSlidingSum {
	    		count++
	    	}
	    }
		prevSlidingSum = slidingSum
	}
	return strconv.Itoa(count), nil
}
