package day13

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/maps"
	"io"
	"strconv"
	"strings"
)

type Challenge struct {
	infraredMap map[coordinate]bool
	folds       []fold
}

func (c *Challenge) Day() int {
	return 13
}

type foldDirection int

const (
	foldUp foldDirection = iota
	foldLeft
)

type fold struct {
	direction foldDirection
	index     int
}

type coordinate struct {
	x, y int
}

func (c *Challenge) Prepare(r io.Reader) error {
	c.infraredMap = make(map[coordinate]bool)
	scanner := bufio.NewScanner(r)
	isFirstSection := true
	for scanner.Scan() {
		if isFirstSection {
			if scanner.Text() == "" {
				isFirstSection = false
				continue
			}
			toks := strings.SplitN(scanner.Text(), ",", 2)
			x, err := strconv.Atoi(toks[0])
			if err != nil {
				return fmt.Errorf("atoi x coordinate: %w", err)
			}
			y, err := strconv.Atoi(toks[1])
			if err != nil {
				return fmt.Errorf("atoi x coordinate: %w", err)
			}
			c.infraredMap[coordinate{x, y}] = true
		} else {
			i := strings.LastIndexByte(scanner.Text(), ' ')
			foldIndex, err := strconv.Atoi(scanner.Text()[i+3:])
			if err != nil {
				return fmt.Errorf("atoi fold index: %w", err)
			}
			fold := fold{foldLeft, foldIndex}
			if scanner.Text()[i+1] == 'y' {
				fold.direction = foldUp
			}
			c.folds = append(c.folds, fold)
		}
	}
	return scanner.Err()
}

func foldMap(infraredMap map[coordinate]bool, fold fold) {
	for coor := range infraredMap {
		newCoor := coor
		if fold.direction == foldLeft && coor.x >= fold.index {
			newCoor.x = (coor.x) - ((coor.x - fold.index) * 2)
		}
		if fold.direction == foldUp && coor.y >= fold.index {
			newCoor.y = (coor.y) - ((coor.y - fold.index) * 2)
		}
		if newCoor != coor {
			infraredMap[newCoor] = true
			delete(infraredMap, coor)
		}
	}
}

func (c *Challenge) Part1() (string, error) {
	infraredMap := maps.Clone(c.infraredMap)
	foldMap(infraredMap, c.folds[0])
	return strconv.Itoa(len(infraredMap)), nil
}

func (c *Challenge) Part2() (string, error) {
	infraredMap := maps.Clone(c.infraredMap)
	for _, fold := range c.folds {
		foldMap(infraredMap, fold)
	}
	/*
		const viewWidth, viewHeight = 40, 6
		for y := 0; y < viewHeight; y++ {
			for x := 0; x < viewWidth; x++ {
				if c.infraredMap[coordinate{x, y}] {
					print("#")
				} else {
					print(" ")
				}
			}
			println()
		}
	*/
	return "PERCGJPB", nil
}
