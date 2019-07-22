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
}

func TestSink(t *testing.T) {

}

func TestOutput(t *testing.T) {

}

func TestSplit(t *testing.T) {

}
