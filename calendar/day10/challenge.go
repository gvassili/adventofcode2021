package day10

import (
	"bufio"
	"io"
	"sort"
	"strconv"
)

type Challenge struct {
	input [][]uint8
}

func (c *Challenge) Day() int {
	return 10
}

func (c *Challenge) Prepare(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	buffer := make([]uint8, 0, 1024)
	c.input = make([][]uint8, 0, 16)
	for scanner.Scan() {
		begIdx := len(buffer)
		buffer = append(buffer, scanner.Bytes()...)
		c.input = append(c.input, buffer[begIdx:])
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func errorValueFromByte(b byte) int {
	switch b {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	}
	return 0
}

func (c *Challenge) Part1() (string, error) {
	openStack := make([]byte, 0, 32)
	score := 0
chunkLoop:
	for _, chunk := range c.input {
		openStack = openStack[0:]
		for _, b := range chunk {
			switch b {
			case '(', '[', '{', '<':
				if b == '(' {
					openStack = append(openStack, ')')
				} else {
					openStack = append(openStack, b+2)
				}

			default:
				if b != openStack[len(openStack)-1] {
					score += errorValueFromByte(b)
					continue chunkLoop
				} else {
					openStack = openStack[:len(openStack)-1]
				}
			}
		}
	}
	return strconv.Itoa(score), nil
}

func (c *Challenge) Part2() (string, error) {
	openStack := make([]byte, 0, 32)
	scores := make([]int, 0, 128)

chunkLoop:
	for _, chunk := range c.input {
		openStack = openStack[:0]
		for _, b := range chunk {
			switch b {
			case '(', '[', '{', '<':
				if b == '(' {
					openStack = append(openStack, ')')
				} else {
					openStack = append(openStack, b+2)
				}

			default:
				if b != openStack[len(openStack)-1] {
					continue chunkLoop
				} else {
					openStack = openStack[:len(openStack)-1]
				}
			}
		}
		chunkScore := 0
		for i := len(openStack) - 1; i >= 0; i-- {
			chunkScore *= 5
			switch openStack[i] {
			case ')':
				chunkScore += 1
			case ']':
				chunkScore += 2
			case '}':
				chunkScore += 3
			case '>':
				chunkScore += 4
			}
		}
		scores = append(scores, chunkScore)
	}
	sort.Ints(scores)
	return strconv.Itoa(scores[(len(scores) >> 1)]), nil
}
