package day05

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type vector struct {
	x int
	y int
}

type vent struct {
	p1 vector
	p2 vector
}

type Challenge struct {
	input  []vent
	width  int
	height int
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

func (c *Challenge) Day() int {
	return 5
}

func (c *Challenge) Prepare(r io.Reader) error {
	c.input = make([]vent, 0, 1024)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		v := vent{}
		_, err := fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &v.p1.x, &v.p1.y, &v.p2.x, &v.p2.y)
		if err != nil {
			return fmt.Errorf("scan line: %w", err)
		}
		c.width = max(c.width, max(v.p1.x, v.p2.x))
		c.height = max(c.height, max(v.p1.y, v.p2.y))
		c.input = append(c.input, v)
	}
	c.width++
	c.height++
	return scanner.Err()
}

func (c *Challenge) Part1() (string, error) {
	oceanFloor := make([]int, c.width*c.height)
	dangerousAreas := 0
	for _, v := range c.input {
		if v.p1.x == v.p2.x {
			y1, y2 := v.p1.y, v.p2.y
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for ; y1 <= y2; y1++ {
				oceanFloor[v.p1.x+(y1*c.width)]++
				if oceanFloor[v.p1.x+(y1*c.width)] == 2 {
					dangerousAreas++
				}
			}
		} else if v.p1.y == v.p2.y {
			x1, x2 := v.p1.x, v.p2.x
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			for ; x1 <= x2; x1++ {
				oceanFloor[x1+(v.p1.y*c.width)]++
				if oceanFloor[x1+(v.p1.y*c.width)] == 2 {
					dangerousAreas++
				}
			}
		}
	}
	return strconv.Itoa(dangerousAreas), nil
}

func (c *Challenge) Part2() (string, error) {
	oceanFloor := make([]int, c.width*c.height)
	dangerousAreas := 0

	for _, v := range c.input {
		p1 := v.p1
		dirX, dirY := 1, 1
		if v.p2.x < v.p1.x {
			dirX = -1
		}
		if v.p2.y < v.p1.y {
			dirY = -1
		}

		for {
			c := p1.x + p1.y*c.width
			oceanFloor[c]++
			if oceanFloor[c] == 2 {
				dangerousAreas++
			}
			if p1.x == v.p2.x && p1.y == v.p2.y {
				break
			}
			if p1.x != v.p2.x {
				p1.x += dirX
			}
			if p1.y != v.p2.y {
				p1.y += dirY
			}
		}
	}
	return strconv.Itoa(dangerousAreas), nil
}
