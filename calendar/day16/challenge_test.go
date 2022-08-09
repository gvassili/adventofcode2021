package day16

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

type testCase struct {
	input string
	part1 string
	part2 string
}

var testCases = []testCase{
	{"D2FE28", "6", ""},
	{"8A004A801A8002F478", "16", ""},
	{"620080001611562C8802118E34", "12", ""},
	{"C0015000016115A2E0802F182340", "23", ""},
	{"A0016C880162017C3686B18A3D4780", "31", ""},
	{"C200B40A82", "", "3"},
	{"04005AC33890", "", "54"},
	{"880086C3E88112", "", "7"},
	{"CE00C43D881120", "", "9"},
	{"D8005AC2A8F0", "", "1"},
	{"F600BC2D8F", "", "0"},
	{"9C005AC2F8F0", "", "0"},
	{"9C0141080250320F1802104A08", "", "1"},
}

var fullInput = func() []byte {
	r, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return b
}()

func TestChallenge_Prepare(t *testing.T) {
	for _, testCase := range testCases {
		var c Challenge
		err := c.Prepare(bytes.NewReader([]byte(testCase.input)))
		assert.NoError(t, err)
	}
}

func BenchmarkChallenge_Prepare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := bytes.NewReader(fullInput)
		var c Challenge
		c.Prepare(buf)
	}
}

func TestChallenge_Part1(t *testing.T) {
	for _, testCase := range testCases {
		if testCase.part1 == "" {
			continue
		}
		var c Challenge
		err := c.Prepare(bytes.NewReader([]byte(testCase.input)))
		assert.NoError(t, err)
		r, err := c.Part1()
		assert.NoError(t, err)
		assert.Equal(t, testCase.part1, r)
	}
}

func BenchmarkChallenge_Part1(b *testing.B) {
	buf := bytes.NewReader(fullInput)
	var c Challenge
	c.Prepare(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Part1()
	}
}

func TestChallenge_Part2(t *testing.T) {
	for _, testCase := range testCases {
		if testCase.part2 == "" {
			continue
		}
		var c Challenge
		err := c.Prepare(bytes.NewReader([]byte(testCase.input)))
		assert.NoError(t, err)
		r, err := c.Part2()
		assert.NoError(t, err)
		assert.Equal(t, testCase.part2, r)
	}
}

func BenchmarkChallenge_Part2(b *testing.B) {
	buf := bytes.NewReader(fullInput)
	var c Challenge
	c.Prepare(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Part2()
	}
}
