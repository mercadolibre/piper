package piper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitter(t *testing.T) {
	inc, incOut := NewBuilder().Output(incOp)
	square, squareOut := NewBuilder().Output(squareOp)

	splitter := newSplitterStage(inc, square)
	go splitter.run()
	defer splitter.stop()

	go func() {
		for _, x := range inputs {
			splitter.in <- x
		}
	}()

	for _, x := range inputs {
		assert.Equal(t, x+1, <-incOut)
		assert.Equal(t, x*x, <-squareOut)
	}

}
