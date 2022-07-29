package day15

import (
	"bufio"
	"container/heap"
	"io"
	"math"
	"strconv"
)

type Challenge struct {
	dangerMap [][]uint8
}

func (Challenge) Day() int {
	return 15
}

func (c *Challenge) Prepare(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	var buffer []uint8
	for scanner.Scan() {
		for _, c := range scanner.Text() {
			buffer = append(buffer, byte(c)-'0')
		}
		c.dangerMap = append(c.dangerMap, buffer[len(buffer)-len(scanner.Text()):])
	}
	return scanner.Err()
}

type node struct {
	y, x int
}

type nodes []node

func (n nodes) Len() int {
	return len(n)
}

func (n nodes) Less(i, j int) bool {
	return n[i].x+n[i].y < n[j].x+n[j].y
}

func (n nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n *nodes) Push(x any) {
	*n = append(*n, x.(node))
}

func (n *nodes) Pop() any {
	old := *n
	i := len(old)
	x := old[i-1]
	*n = old[0 : i-1]
	return x
}

func (c *Challenge) Part1() (string, error) {
	var nodes nodes
	distanceMap := make([][]int, len(c.dangerMap))
	for i, row := range c.dangerMap {
		distanceMap[i] = make([]int, len(row))
		for j := range distanceMap[i] {
			distanceMap[i][j] = math.MaxInt
		}
	}
	finalNode := node{len(distanceMap) - 1, len(distanceMap[len(distanceMap)-1]) - 1}
	distanceMap[0][0] = int(c.dangerMap[0][0])
	heap.Push(&nodes, node{0, 0})
	for nodes.Len() > 0 {
		n := heap.Pop(&nodes).(node)
		adjacentNodes := [4]node{
			{n.y + 1, n.x},
			{n.y, n.x + 1},
			{n.y - 1, n.x},
			{n.y, n.x - 1},
		}
		for _, adjacentNode := range adjacentNodes {
			if adjacentNode.y < 0 || adjacentNode.y >= len(distanceMap) ||
				adjacentNode.x < 0 || adjacentNode.x >= len(distanceMap[adjacentNode.y]) {
				continue
			}
			distance := distanceMap[n.y][n.x] + int(c.dangerMap[adjacentNode.y][adjacentNode.x])
			if distanceMap[adjacentNode.y][adjacentNode.x] <= distance {
				continue
			}
			distanceMap[adjacentNode.y][adjacentNode.x] = distance
			heap.Push(&nodes, adjacentNode)
		}
	}
	return strconv.Itoa(distanceMap[finalNode.y][finalNode.x] - int(c.dangerMap[0][0])), nil
}

func (c *Challenge) Part2() (string, error) {
	var nodes nodes
	distanceMap := make([][]int, len(c.dangerMap)*5)
	for y := 0; y < len(c.dangerMap)*5; y++ {
		distanceMap[y] = make([]int, len(c.dangerMap[y%len(c.dangerMap)])*5)
		for x := range distanceMap[y] {
			distanceMap[y][x] = math.MaxInt
		}
	}
	finalNode := node{len(distanceMap) - 1, len(distanceMap[len(distanceMap)-1]) - 1}
	distanceMap[0][0] = 0
	heap.Push(&nodes, node{0, 0})
	for nodes.Len() > 0 {
		n := heap.Pop(&nodes).(node)
		adjacentNodes := [4]node{
			{n.y + 1, n.x},
			{n.y, n.x + 1},
			{n.y - 1, n.x},
			{n.y, n.x - 1},
		}
		for _, adjacentNode := range adjacentNodes {
			if adjacentNode.y < 0 || adjacentNode.y >= len(distanceMap) ||
				adjacentNode.x < 0 || adjacentNode.x >= len(distanceMap[adjacentNode.y]) {
				continue
			}
			height := len(c.dangerMap)
			width := len(c.dangerMap[adjacentNode.y%height])
			dangerLevel := (((int(c.dangerMap[adjacentNode.y%width][adjacentNode.x%height]) + adjacentNode.y/height + adjacentNode.x/width) - 1) % 9) + 1
			distance := distanceMap[n.y][n.x] + dangerLevel
			if distanceMap[adjacentNode.y][adjacentNode.x] <= distance {
				continue
			}
			distanceMap[adjacentNode.y][adjacentNode.x] = distance
			heap.Push(&nodes, adjacentNode)
		}
	}
	return strconv.Itoa(distanceMap[finalNode.y][finalNode.x]), nil
}
