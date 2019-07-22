package piper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultSinkStage(t *testing.T) {
	sink := defaultSinkStage()
	go sink.run()
	defer sink.stop()

	sink.in <- 1
	sink.in <- 2
	sink.in <- 3

	assert.Empty(t, sink.out)
}

func TestNewSink(t *testing.T) {
	x := 0
	sinkOp := SinkOperator(func(in <-chan interface{}) {
		for range in {
			x++
		}
	})

	sink := newSinkStage(sinkOp)
	go sink.run()
	defer sink.stop()

	n := 5
	for i := 0; i < n; i++ {
		sink.in <- struct{}{}
	}

	assert.Equal(t, x, n)
}

func TestNewBufferedSink(t *testing.T) {
	x := 0
	sinkOp := SinkOperator(func(in <-chan interface{}) {
		for range in {
			x++
		}
	})

	bufSize := 1
	sink := newBufferedSinkStage(bufSize, sinkOp)
	assert.Equal(t, bufSize, cap(sink.in))
}
