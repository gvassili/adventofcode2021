package day03

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Challenge struct {
	input []uint64
	bits  int
}

func (c *Challenge) Day() int {
	return 3
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (c *Challenge) Prepare(r io.Reader) error {
	c.input = make([]uint64, 0, 1024)
	scanner := bufio.NewScanner(r)
	c.bits = 64
	for scanner.Scan() {
		n, err := strconv.ParseUint(scanner.Text(), 2, 64)
		if err != nil {
			return fmt.Errorf("parseUint line: %w", err)
		}
		c.bits = min(c.bits, len(scanner.Bytes()))
		c.input = append(c.input, n)
	}
	return scanner.Err()
}

func (c *Challenge) Part1() (string, error) {
	counts := make([]int64, c.bits)
	for _, n := range c.input {
		for i := 0; i < c.bits; i++ {
			counts[i] += int64((((n >> i) & 1) << 1) - 1)
		}
	}
	var gammaRate uint64
	for i, c := range counts {
		if c > 0 {
			gammaRate |= 1 << i
		}
	}
	epsilonRate := (^gammaRate) & (^(^uint64(0) << c.bits))
	return strconv.Itoa(int(gammaRate * epsilonRate)), nil
}

func (c *Challenge) Part2() (string, error) {
	split := func(arr []uint64, bit int) ([]uint64, []uint64) {
		zeroCount, oneCount := 0, 0
		for zeroCount+oneCount < len(arr) {
			leftIsZero, rightIsZero := arr[zeroCount]&(1<<bit) == 0, arr[len(arr)-oneCount-1]&(1<<bit) == 0
			if !leftIsZero && rightIsZero {
				arr[zeroCount], arr[len(arr)-oneCount-1] = arr[len(arr)-oneCount-1], arr[zeroCount]
				zeroCount++
				oneCount++
			}
			if leftIsZero {
				zeroCount++
			} else if !rightIsZero {
				oneCount++
			}
		}
		return arr[:zeroCount], arr[zeroCount:]
	}
	var findRatingRec func(arr []uint64, bit int, keepHigher bool) uint64
	findRatingRec = func(arr []uint64, bit int, keepHigher bool) uint64 {
		if len(arr) == 1 {
			return arr[0]
		}
		zeroArr, oneArr := split(arr, bit)
		if (len(oneArr) >= len(zeroArr)) == keepHigher {
			return findRatingRec(oneArr, bit-1, keepHigher)
		} else {
			return findRatingRec(zeroArr, bit-1, keepHigher)
		}
	}
	oneArr, zeroArr := split(c.input, c.bits-1)
	if len(oneArr) >= len(zeroArr) {
		zeroArr, oneArr = oneArr, zeroArr
	}
	oxygenRating := findRatingRec(zeroArr, c.bits-2, true)
	co2ScrubberRating := findRatingRec(oneArr, c.bits-2, false)
	return strconv.Itoa(int(oxygenRating * co2ScrubberRating)), nil
}
