package day14

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"
)

/* Naive algorithm
type Challenge struct {
	polymerTemplate    string
	pairInsertionRules map[string]byte
}

func (c *Challenge) Prepare(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	c.polymerTemplate = ""
	c.pairInsertionRules = make(map[string]byte)
	for scanner.Scan() {
		if c.polymerTemplate == "" {
			c.polymerTemplate = scanner.Text()
		} else if scanner.Text() != "" {
			toks := strings.SplitN(scanner.Text(), " -> ", 2)
			c.pairInsertionRules[toks[0]] = toks[1][0]
		}
	}
	return scanner.Err()
}

func (c *Challenge) Part1() (string, error) {
	polymer := c.polymerTemplate
	for n := 0; n < 40; n++ {
		nextPolymer := strings.Builder{}
		nextPolymer.Grow((len(c.polymerTemplate) * 2) - 1)
		for i := 0; i < len(polymer)-1; i++ {
			nextPolymer.WriteByte(polymer[i])
			nextPolymer.WriteByte(c.pairInsertionRules[polymer[i:i+2]])
		}
		nextPolymer.WriteByte(polymer[len(polymer)-1])
		polymer = nextPolymer.String()
		elementSums := make(map[byte]int, 4)
		for _, element := range polymer {
			elementSums[byte(element)]++
		}
		orderedElementSums := maps.Values(elementSums)
		slices.Sort(orderedElementSums)
	}
	elementSums := make(map[byte]int, 4)
	for _, element := range polymer {
		elementSums[byte(element)]++
	}
	orderedElementSums := maps.Values(elementSums)
	slices.Sort(orderedElementSums)
	return strconv.Itoa(orderedElementSums[len(orderedElementSums)-1] - orderedElementSums[0]), nil
}
*/

type element uint8
type pair uint8

type rule struct {
	e1       element
	e2       element
	insertP1 pair
	insertP2 pair
}

type Challenge struct {
	polymerTemplate    []int  // index is pair
	pairInsertionRules []rule // index is pair

	maxElement   int
	firstElement element
	lastElement  element
}

func (c *Challenge) Day() int {
	return 14
}

func (c *Challenge) Prepare(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	elementMap := make(map[byte]element)
	pairMap := make(map[[2]byte]pair)

	getElement := func(b byte) (e element) {
		e, ok := elementMap[b]
		if !ok {
			e = element(len(elementMap))
			elementMap[b] = e
		}
		return
	}
	getPair := func(b [2]byte) (p pair) {
		p, ok := pairMap[b]
		if !ok {
			p = pair(len(pairMap))
			pairMap[b] = p
		}
		return
	}

	var polymerTemplateString string
	for scanner.Scan() {
		if polymerTemplateString == "" {
			polymerTemplateString = scanner.Text()
		} else if scanner.Text() != "" {
			toks := strings.SplitN(scanner.Text(), " -> ", 2)
			p := getPair(*(*[2]byte)([]byte(toks[0])))
			e1, e2 := getElement(toks[0][0]), getElement(toks[0][1])
			if int(p) >= len(c.pairInsertionRules) {
				tmpPairInsertionRules := make([]rule, p+1)
				copy(tmpPairInsertionRules, c.pairInsertionRules)
				c.pairInsertionRules = tmpPairInsertionRules
			}
			c.pairInsertionRules[p] = rule{
				e1:       e1,
				e2:       e2,
				insertP1: getPair([2]byte{toks[0][0], toks[1][0]}),
				insertP2: getPair([2]byte{toks[1][0], toks[0][1]}),
			}
		}
	}
	c.polymerTemplate = make([]int, len(c.pairInsertionRules))
	for i := 0; i < len(polymerTemplateString)-1; i++ {
		p := getPair(*(*[2]byte)([]byte(polymerTemplateString[i : i+2])))
		c.polymerTemplate[p]++
	}
	c.firstElement = getElement(polymerTemplateString[0])
	c.lastElement = getElement(polymerTemplateString[len(polymerTemplateString)-1])
	c.maxElement = len(elementMap)
	return scanner.Err()
}

func (c *Challenge) polymerization(step int) int {
	polymer, nextPolymer := make([]int, len(c.polymerTemplate)), make([]int, len(c.polymerTemplate))
	copy(polymer, c.polymerTemplate)
	for i := 0; i < step; i++ {
		copy(nextPolymer, polymer)
		for pi, count := range polymer {
			if count > 0 {
				rule := c.pairInsertionRules[pi]
				nextPolymer[pi] -= count
				nextPolymer[rule.insertP1] += count
				nextPolymer[rule.insertP2] += count
			}
		}
		polymer, nextPolymer = nextPolymer, polymer
	}

	elementCount := make([]int, c.maxElement)
	elementCount[c.firstElement]++
	elementCount[c.lastElement]++
	for pi, count := range polymer {
		rule := c.pairInsertionRules[pi]
		elementCount[rule.e1] += count
		elementCount[rule.e2] += count
	}
	minCount, maxCount := math.MaxInt, 0
	for _, count := range elementCount {
		if minCount > count {
			minCount = count
		}
		if maxCount < count {
			maxCount = count
		}
	}
	return (maxCount >> 1) - (minCount >> 1)
}

func (c *Challenge) Part1() (string, error) {
	return strconv.Itoa(c.polymerization(10)), nil
}

func (c *Challenge) Part2() (string, error) {
	return strconv.Itoa(c.polymerization(40)), nil
}
