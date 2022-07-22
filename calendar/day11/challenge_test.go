package day11

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

const input = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
`

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
	var c Challenge
	err := c.Prepare(bytes.NewReader([]byte(input)))
	assert.NoError(t, err)
	assert.Equal(t, []int{5, 4, 8, 3, 1, 4, 3, 2, 2, 3, 2, 7, 4, 5, 8, 5, 4, 7, 1, 1, 5, 2, 6, 4, 5, 5, 6, 1, 7, 3, 6, 1, 4, 1, 3, 3, 6, 1, 4, 6, 6, 3, 5, 7, 3, 8, 5, 4, 7, 8, 4, 1, 6, 7, 5, 2, 4, 6, 4, 5, 2, 1, 7, 6, 8, 4, 1, 7, 2, 1, 6, 8, 8, 2, 8, 8, 1, 1, 3, 4, 4, 8, 4, 6, 8, 4, 8, 5, 5, 4, 5, 2, 8, 3, 7, 5, 1, 5, 2, 6}, c.input)
}

func BenchmarkChallenge_Prepare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := bytes.NewReader(fullInput)
		var c Challenge
		c.Prepare(buf)
	}
}

func TestChallenge_Part1(t *testing.T) {
	var c Challenge
	err := c.Prepare(bytes.NewReader([]byte(input)))
	assert.NoError(t, err)
	r, err := c.Part1()
	assert.NoError(t, err)
	assert.Equal(t, "1656", r)
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
	var c Challenge
	err := c.Prepare(bytes.NewReader([]byte(input)))
	assert.NoError(t, err)
	r, err := c.Part2()
	assert.NoError(t, err)
	assert.Equal(t, "195", r)
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
