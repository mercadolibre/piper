package piper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	inputs = []int{1, 2, 3, 4, 5, 6}

	incOp = Operator(func(in <-chan interface{}, out chan<- interface{}) {
		for x := range in {
			out <- (x.(int)) + 1
		}
	})

	squareOp = Operator(func(in <-chan interface{}, out chan<- interface{}) {
		for _x := range in {
			x := _x.(int)
			out <- x * x
		}
	})

	printOp = SinkOperator(func(in <-chan interface{}) {
		for x := range in {
			fmt.Println(x)
		}
	})
)

func IncSquare(x int) int {
	return (x + 1) * (x + 1)
}

func testIncSquare(t *testing.T, p *Pipeline, out <-chan interface{}) {
	for _, x := range inputs {
		p.In() <- x
		assert.Equal(t, IncSquare(x), <-out)
	}
}
