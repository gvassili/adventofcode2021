package day16

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
)

type Challenge struct {
	transmission []uint8
}

type typeId int

const (
	addTypeId typeId = 0
	mulTypeId        = 1
	minTypeId        = 2
	maxTypeId        = 3
	valTypeId        = 4
	gtTypeId         = 5
	ltTypeId         = 6
	eqTypeId         = 7
)

func (Challenge) Day() int {
	return 16
}

func hexByteBoInt(b byte) uint8 {
	if b <= '9' {
		return b - '0'
	}
	return b - 'A' + 10
}

func (c *Challenge) Prepare(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		return errors.New("empty scanner")
	}
	for i := 0; i < len(scanner.Text()); i += 2 {
		c.transmission = append(c.transmission, (hexByteBoInt(scanner.Text()[i])<<4)|hexByteBoInt(scanner.Text()[i+1]))
	}
	return scanner.Err()
}

type header struct {
	Version uint8
}

func (h header) version() uint8 {
	return h.Version
}

type value struct {
	header
	Number int
}

type operator struct {
	header
	Type   typeId
	Params []packet
}

type packet any

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (c Challenge) decodTransmission() packet {
	const (
		bitSize = 16
		bitMask = bitSize - 1
	)
	bitCursor := 0
	var read = func(bits int) int {
		v := 0
		bitsToRead := bits
		for bitsToRead > 0 {
			localBitsToRead := minInt(bitsToRead, bitSize-(bitCursor&bitMask))
			v = (v << localBitsToRead) | (int(c.transmission[bitCursor>>3])>>((bitSize-(localBitsToRead))-(bitCursor&bitMask)))&((1<<localBitsToRead)-1)
			bitCursor += localBitsToRead
			bitsToRead -= localBitsToRead
		}
		return v
	}

	var decodPacket func() packet
	decodPacket = func() packet {
		h := header{
			Version: uint8(read(3)),
		}
		t := uint8(read(3))
		switch t {
		// value packet
		case valTypeId:
			v := value{header: h}
			for {
				chunk := read(5)
				v.Number = (v.Number << 4) | (int(chunk) & 0b1111)
				if (chunk & (1 << 4)) == 0 {
					break
				}
			}
			return v
		// operator packet
		default:
			o := operator{header: h, Type: typeId(t)}
			lengthType := read(1)
			if lengthType == 0 {
				end := bitCursor + read(15)
				for bitCursor < end {
					o.Params = append(o.Params, decodPacket())
				}
			} else {
				length := read(11)
				for i := 0; i < length; i++ {
					o.Params = append(o.Params, decodPacket())
				}
			}
			return o
		}
	}
	return decodPacket()
}

func (c *Challenge) Part1() (string, error) {
	var sumVersion func(p packet) int
	sumVersion = func(p packet) int {
		switch p := p.(type) {
		case value:
			return int(p.Version)
		case operator:
			sum := int(p.Version)
			for _, p := range p.Params {
				sum += sumVersion(p)
			}
			return sum
		}
		panic(fmt.Errorf("invalid packet type %T", p))
	}
	return strconv.Itoa(sumVersion(c.decodTransmission())), nil
}

func (c *Challenge) Part2() (string, error) {
	getOperator := func(typeId typeId) func(a, b int) int {
		switch typeId {
		case addTypeId:
			return func(a, b int) int { return a + b }
		case mulTypeId:
			return func(a, b int) int { return a * b }
		case minTypeId:
			return func(a, b int) int {
				if a < b {
					return a
				}
				return b
			}
		case maxTypeId:
			return func(a, b int) int {
				if a > b {
					return a
				}
				return b
			}
		case gtTypeId:
			return func(a, b int) int {
				if a > b {
					return 1
				}
				return 0
			}
		case ltTypeId:
			return func(a, b int) int {
				if a < b {
					return 1
				}
				return 0
			}
		case eqTypeId:
			return func(a, b int) int {
				if a == b {
					return 1
				}
				return 0
			}
		}
		panic(fmt.Errorf("unknow operator type id %d", typeId))
	}
	var evaluateTransmission func(p packet) int
	evaluateTransmission = func(p packet) int {
		switch p := p.(type) {
		case value:
			return p.Number
		case operator:
			op := getOperator(p.Type)
			v := evaluateTransmission(p.Params[0])
			for i := 1; i < len(p.Params); i++ {
				b := evaluateTransmission(p.Params[i])
				v = op(v, b)
			}
			return v
		}
		panic(fmt.Errorf("invalid packet type %T", p))
	}
	return strconv.Itoa(evaluateTransmission(c.decodTransmission())), nil
}
