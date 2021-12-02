package day02

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type command struct {
	direction string
	amount    int
}

type Challenge struct {
	input []command
}

func (c *Challenge) Day() int {
	return 2
}

func (c *Challenge) Prepare(r io.Reader) error {
	c.input = make([]command, 0, 1024)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		command := command{}
		if _, err := fmt.Sscanf(scanner.Text(), "%s %d", &command.direction, &command.amount); err != nil {
			return fmt.Errorf("scanf line: %w", err)
		}
		c.input = append(c.input, command)
	}
	return scanner.Err()
}

func (c *Challenge) Part1() (string, error) {
	depth, position := 0, 0
	for _, command := range c.input {
		switch command.direction {
		case "forward":
			position += command.amount
		case "up":
			depth -= command.amount
		case "down":
			depth += command.amount
		}
	}
	return strconv.Itoa(depth * position), nil
}

func (c *Challenge) Part2() (string, error) {
	pitch, position, depth := 0, 0, 0
	for _, command := range c.input {
		switch command.direction {
		case "forward":
			position += command.amount
			depth += command.amount * pitch
		case "up":
			pitch -= command.amount
		case "down":
			pitch += command.amount
		}
	}
	return strconv.Itoa(depth * position), nil
}
