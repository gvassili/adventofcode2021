package day12

import (
	"bufio"
	"io"
	"strconv"
	"strings"
	"unicode"
)

type Challenge struct {
	startNode      *Node
	smallNodeCount int
}

func (c *Challenge) Day() int {
	return 12
}

type NodeType int

const (
	visitedSmallRoomType NodeType = iota
	smallCaveNodeType
	bigCaveNodeType
	startNodeType
	endNodeType
)

type Node struct {
	name      string
	neighbors []*Node
	nodeFlag  NodeType
}

func (c *Challenge) Prepare(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	nodes := make(map[string]*Node)
	getNode := func(name string) *Node {
		node, ok := nodes[name]
		if !ok {
			node = &Node{
				name: name,
			}
			nodes[name] = node
			switch {
			case name == "start":
				node.nodeFlag = startNodeType
				c.startNode = node
			case name == "end":
				node.nodeFlag = endNodeType
			case unicode.IsUpper(rune(name[0])):
				node.nodeFlag = bigCaveNodeType
			case unicode.IsLower(rune(name[0])):
				node.nodeFlag = smallCaveNodeType
				c.smallNodeCount++
			default:
				node.nodeFlag = visitedSmallRoomType
			}
		}
		return node
	}
	for scanner.Scan() {
		toks := strings.SplitN(scanner.Text(), "-", 2)
		node, neighbourNode := getNode(toks[0]), getNode(toks[1])
		node.neighbors = append(node.neighbors, neighbourNode)
		neighbourNode.neighbors = append(neighbourNode.neighbors, node)
	}
	return scanner.Err()
}

func (c *Challenge) Part1() (string, error) {
	var countPathRec func(*Node) int
	countPathRec = func(node *Node) int {
		if node.nodeFlag == smallCaveNodeType {
			node.nodeFlag = visitedSmallRoomType
			defer func() {
				node.nodeFlag = smallCaveNodeType
			}()
		}
		pathCount := 0
		for _, neighbour := range node.neighbors {
			if neighbour.nodeFlag == bigCaveNodeType || neighbour.nodeFlag == smallCaveNodeType {
				pathCount += countPathRec(neighbour)
			} else if neighbour.nodeFlag == endNodeType {
				pathCount++
			}
		}
		if node.nodeFlag&visitedSmallRoomType != 0 {
			node.nodeFlag = smallCaveNodeType
		}
		return pathCount
	}
	return strconv.Itoa(countPathRec(c.startNode)), nil
}

func (c *Challenge) Part2() (string, error) {
	var countPathRec func(*Node, bool) int
	countPathRec = func(node *Node, visitedSmallRoomTwice bool) int {
		if node.nodeFlag == smallCaveNodeType {
			node.nodeFlag = visitedSmallRoomType
			defer func() {
				node.nodeFlag = smallCaveNodeType
			}()
		}
		pathCount := 0
		for _, neighbour := range node.neighbors {
			if neighbour.nodeFlag == bigCaveNodeType || neighbour.nodeFlag == smallCaveNodeType {
				pathCount += countPathRec(neighbour, visitedSmallRoomTwice)
			} else if neighbour.nodeFlag == visitedSmallRoomType && !visitedSmallRoomTwice {
				pathCount += countPathRec(neighbour, true)
			} else if neighbour.nodeFlag == endNodeType {
				pathCount++
			}
		}
		if node.nodeFlag&visitedSmallRoomType != 0 {
			node.nodeFlag = smallCaveNodeType
		}
		return pathCount
	}
	return strconv.Itoa(countPathRec(c.startNode, false)), nil
}
