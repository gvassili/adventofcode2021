package day09

import (
	"bufio"
	"io"
	"math"
	"strconv"
)

type Challenge struct {
	input  [][]uint8
	height int
	width  int
}

func (c *Challenge) Day() int {
	return 9
}

func (c *Challenge) Prepare(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	buffer := make([]uint8, 0, 1024)
	c.input = make([][]uint8, 0, 16)
	for scanner.Scan() {
		begIdx := len(buffer)
		for _, b := range scanner.Bytes() {
			buffer = append(buffer, b-'0')
		}
		c.input = append(c.input, buffer[begIdx:])
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	c.height, c.width = len(c.input), len(c.input[0])
	return nil
}

func (c *Challenge) heightAt(y, x int) uint8 {
	if y >= 0 && y < c.height && x >= 0 && x < c.width {
		return c.input[y][x]
	}
	return math.MaxUint8
}

func (c *Challenge) Part1() (string, error) {
	sum := 0
	for y, row := range c.input {
		for x, cell := range row {
			if cell < c.heightAt(y, x+1) &&
				cell < c.heightAt(y, x-1) &&
				cell < c.heightAt(y+1, x) &&
				cell < c.heightAt(y-1, x) {
				sum += int(cell) + 1
			}
		}
	}
	return strconv.Itoa(sum), nil
}

func (c *Challenge) Part2() (string, error) {
	gapMap := make([][]bool, c.height)
	for i := range c.input {
		gapMap[i] = make([]bool, c.width)
	}
	var markGapRec func(y, x int) (size int)
	markGapRec = func(y, x int) int {
		if c.heightAt(y, x) >= 9 || gapMap[y][x] {
			return 0
		}
		gapMap[y][x] = true
		return markGapRec(y, x+1) +
			markGapRec(y, x-1) +
			markGapRec(y+1, x) +
			markGapRec(y-1, x) + 1
	}
	biggestGaps := make([]int, 3)
	for y, row := range c.input {
		for x := range row {
			gapSize := markGapRec(y, x)
			if gapSize > 0 {
				for i, otherGap := range biggestGaps {
					if gapSize > otherGap {
						copy(biggestGaps[i+1:], biggestGaps[i:])
						biggestGaps[i] = gapSize
						break
					}
				}
			}
		}
	}
	return strconv.Itoa(biggestGaps[0] * biggestGaps[1] * biggestGaps[2]), nil
}
