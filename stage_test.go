package piper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStageCreation(t *testing.T) {
	stage := newStage(incOp)

	assert.Nil(t, stage.next)
	assert.Empty(t, stage.done)
	assert.Empty(t, stage.in)
	assert.Empty(t, stage.out)
}

func TestNewBufferedStageCreation(t *testing.T) {
	bufSize := 10
	stage := newBufferedStage(bufSize, incOp)

	assert.Nil(t, stage.next)
	assert.Empty(t, stage.done)

	assert.Equal(t, bufSize, cap(stage.in))
	assert.Equal(t, bufSize, cap(stage.out))
}

func TestStageRunAndStop(t *testing.T) {
	stage := newBufferedStage(2, incOp)

	go stage.run()
	stage.in <- 1
	stage.in <- 2
	stage.stop()

	assert.Equal(t, 2, <-stage.out)
	assert.Equal(t, 3, <-stage.out)

	_, ok := <-stage.in
	assert.False(t, ok)
	assert.NotEmpty(t, stage.done)
}
