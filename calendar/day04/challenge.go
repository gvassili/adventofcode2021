package day04

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

type Challenge struct {
	answers []int
	cells   []int
}

func (c *Challenge) Day() int {
	return 4
}

func (c *Challenge) Prepare(r io.Reader) error {
	c.cells = make([]int, 0, 1024)
	scanner := bufio.NewScanner(r)
	line := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		if line == 0 {
			strs := strings.Split(scanner.Text(), ",")
			c.answers = make([]int, 0, 32)
			for _, s := range strs {
				n, err := strconv.Atoi(s)
				if err != nil {
					return fmt.Errorf("atoi answer line %d: %w", line, err)
				}
				c.answers = append(c.answers, n)
			}
		} else {
			strs := strings.Split(scanner.Text(), " ")
			for _, s := range strs {
				if s != "" {
					n, err := strconv.Atoi(s)
					if err != nil {
						return fmt.Errorf("atoi board line %d: %w", line, err)
					}
					c.cells = append(c.cells, n)
				}
			}
		}
		line++
	}
	return scanner.Err()
}

const boardWith = 5
const boardHeight = 5
const boardSize = boardWith * boardHeight
const rowMask = 0b11111
const colMask = 0b100001000010000100001

type cell struct {
	value int
	idx   int
}
type cells []cell

func (c cells) Len() int {
	return len(c)
}

func (c cells) Less(i, j int) bool {
	return c[i].value < c[j].value
}

func (c cells) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

type board struct {
	sum    int
	result int
	done   bool
}

func (c *Challenge) Part1() (string, error) {
	boards := make([]board, len(c.cells)/boardSize)
	cells := make(cells, len(c.cells))
	for i, c := range c.cells {
		boards[i/boardSize].sum += c
		cells[i] = cell{
			value: c,
			idx:   i,
		}
	}

	sort.Sort(cells)
	checkCell := func(c cell) bool {
		x, y, b := c.idx%boardWith, (c.idx/boardHeight)%boardHeight, c.idx/boardSize
		cellRowMask, cellColMask := rowMask<<(y*boardWith), colMask<<x
		boards[b].result |= 1 << x << (y * boardWith)
		boards[b].sum -= c.value

		if boards[b].result&cellRowMask == cellRowMask || boards[b].result&cellColMask == cellColMask {
			return true
		}
		return false
	}

	for _, a := range c.answers {
		idx := sort.Search(len(cells), func(i int) bool { return cells[i].value >= a })
		for i := idx - 1; i >= 0 && cells[i].value == a; i-- {
			if checkCell(cells[i]) {
				return strconv.Itoa(boards[cells[i].idx/boardSize].sum * a), nil
			}
		}
		for i := idx; i < len(cells) && cells[i].value == a; i++ {
			if checkCell(cells[i]) {
				return strconv.Itoa(boards[cells[i].idx/boardSize].sum * a), nil
			}
		}
	}
	return "", errors.New("answer not found")
}

func (c *Challenge) Part2() (string, error) {
	boards := make([]board, len(c.cells)/boardSize)
	boardDone := 0
	cells := make(cells, len(c.cells))
	for i, c := range c.cells {
		boards[i/boardSize].sum += c
		cells[i] = cell{
			value: c,
			idx:   i,
		}
	}

	sort.Sort(cells)
	checkCell := func(c cell) bool {
		x, y, b := c.idx%boardWith, (c.idx/boardHeight)%boardHeight, c.idx/boardSize
		cellRowMask, cellColMask := rowMask<<(y*boardWith), colMask<<x
		boards[b].result |= 1 << x << (y * boardWith)
		boards[b].sum -= c.value
		if !boards[b].done && (boards[b].result&cellRowMask == cellRowMask || boards[b].result&cellColMask == cellColMask) {
			boards[b].done = true
			boardDone++
		}
		return boardDone == len(boards)
	}

	for _, a := range c.answers {
		idx := sort.Search(len(cells), func(i int) bool { return cells[i].value >= a })
		for i := idx - 1; i >= 0 && cells[i].value == a; i-- {
			if checkCell(cells[i]) {
				return strconv.Itoa(boards[cells[i].idx/boardSize].sum * a), nil
			}
		}
		for i := idx; i < len(cells) && cells[i].value == a; i++ {
			if checkCell(cells[i]) {
				return strconv.Itoa(boards[cells[i].idx/boardSize].sum * a), nil
			}
		}
	}
	return "", errors.New("answer not found")
}
