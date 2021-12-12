package day08

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type line struct {
	in  []string
	out []string
}

type Challenge struct {
	input []line
}

func (c *Challenge) Day() int {
	return 8
}

func (c *Challenge) Prepare(r io.Reader) error {
	c.input = make([]line, 0, 1024)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		m := strings.LastIndexByte(scanner.Text(), '|')
		i := line{
			in:  strings.Split(scanner.Text()[:m-1], " "),
			out: strings.Split(scanner.Text()[m+2:], " "),
		}
		c.input = append(c.input, i)
	}
	return scanner.Err()
}

func (c *Challenge) Part1() (string, error) {
	unqCount := 0
	for _, i := range c.input {
		for _, o := range i.out {
			if len(o) == 2 || len(o) == 3 || len(o) == 4 || len(o) == 7 {
				unqCount++
			}
		}
	}
	return strconv.Itoa(unqCount), nil
}

func isSet(n int, mask int) bool {
	return n&mask == mask
}

func makeMask(str string) (mask int) {
	for _, r := range str {
		mask |= 1 << (r - 'a')
	}
	return
}

func isPartiallySet(n int, mask int) bool {
	bits := n & mask
	return (bits != 0) && ((bits & (bits - 1)) == 0)
}

func (c *Challenge) Part2() (string, error) {
	sum := 0
	for _, l := range c.input {
		string2, string3, string4, string7 := "", "", "", ""
		for _, i := range l.in {
			switch len(i) {
			case 2:
				string2 = i
			case 3:
				string3 = i
			case 4:
				string4 = i
			case 7:
				string7 = i
			}
		}
		maskCf := makeMask(string2)
		maskBd := makeMask(string4) & ^maskCf
		maskEg := makeMask(string7) & ^(maskBd | makeMask(string3))
		n := 0
		for i, o := range l.out {
			if i != 0 {
				n *= 10
			}
			switch {
			case len(o) == 2:
				n += 1
			case len(o) == 4:
				n += 4
			case len(o) == 3:
				n += 7
			case len(o) == 7:
				n += 8
			default:
				binN := makeMask(o)
				switch {
				case isPartiallySet(binN, maskCf) && isPartiallySet(binN, maskBd) && isSet(binN, maskEg):
					n += 2
				case isSet(binN, maskCf) && isPartiallySet(binN, maskBd) && isPartiallySet(binN, maskEg):
					n += 3
				case isPartiallySet(binN, maskCf) && isSet(binN, maskBd) && isPartiallySet(binN, maskEg):
					n += 5
				case isPartiallySet(binN, maskCf) && isSet(binN, maskBd) && isSet(binN, maskEg):
					n += 6
				case isSet(binN, maskCf) && isSet(binN, maskBd) && isPartiallySet(binN, maskEg):
					n += 9
				}
			}
		}
		sum += n
	}
	strings.NewReplacer()
	return strconv.Itoa(sum), nil
}
