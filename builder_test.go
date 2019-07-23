package piper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBuilder(t *testing.T) {
	builder := NewBuilder()

	assert.Equal(t, newPipeline(), builder.pipeline)
}

func TestAddLast(t *testing.T) {
	bufSize := 10
	builder := NewBuilder().
		AddLast(incOp).
		AddLastBuffered(bufSize, squareOp)

	assert.NotEqual(t, builder.pipeline.head, builder.pipeline.tail)
	assert.Equal(t, bufSize, cap(builder.pipeline.tail.in))

	go builder.pipeline.Run()
	defer builder.pipeline.Stop()
	testIncSquare(t, builder.pipeline, builder.pipeline.tail.out)
}

func TestDefaultSink(t *testing.T) {
	pipeline := NewBuilder().DefaultSink()
	pipeline.Run()
	defer pipeline.Stop()

	inputs := []int{1, 2, 3, 4, 5, 6}
	for _, x := range inputs {
		pipeline.In() <- x
		assert.Empty(t, pipeline.tail.out)
	}
}

func TestSink(t *testing.T) {
	count := 0
	sink := SinkOperator(func(in <-chan interface{}) {
		for range in {
			count++
		}
	})

	pipeline := NewBuilder().Sink(sink)
	pipeline.Run()

	for _, x := range inputs {
		pipeline.In() <- x
	}
	pipeline.Stop()
	pipeline.Wait()
	assert.Equal(t, len(inputs), count)
}

func TestBufferedSink(t *testing.T) {
	bufSize := len(inputs)
	sink := SinkOperator(func(in <-chan interface{}) {
	})

	pipeline := NewBuilder().BufferedSink(bufSize, sink)

	for _, x := range inputs {
		pipeline.In() <- x
	}
	assert.Equal(t, len(inputs), len(pipeline.In()))
}

func TestOutput(t *testing.T) {
	pipeline, out := NewBuilder().
		AddLast(incOp).
		Output(squareOp)
	pipeline.Run()
	defer pipeline.Stop()

	testIncSquare(t, pipeline, out)
}

func TestSplit(t *testing.T) {
	incPipeline, incOut := NewBuilder().BufferedOutput(1, incOp)
	squarePipeline, squareOut := NewBuilder().BufferedOutput(1, squareOp)

	pipeline := NewBuilder().Split(incPipeline, squarePipeline)
	pipeline.Run()
	defer pipeline.Stop()

	for _, x := range inputs {
		pipeline.In() <- x

		assert.Equal(t, x+1, <-incOut)
		assert.Equal(t, x*x, <-squareOut)
	}
}
