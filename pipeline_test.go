package piper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeIncSquarePipeline() (*Pipeline, chan interface{}) {
	p := newPipeline()
	p.addLast(newStage(incOp))

	lastStage := newStage(squareOp)
	p.addLast(lastStage)
	return p, lastStage.out
}

func TestPipelineEmpty(t *testing.T) {
	p := newPipeline()

	assert.Nil(t, p.head)
	assert.Nil(t, p.tail)
}

func TestPipelineAddLast(t *testing.T) {
	p := newPipeline()

	incStage := newStage(incOp)
	p.addLast(incStage)
	assert.Equal(t, incStage, p.head)
	assert.Equal(t, incStage, p.tail)

	squareStage := newStage(squareOp)
	p.addLast(squareStage)

	assert.Equal(t, incStage, p.head)
	assert.Equal(t, squareStage, p.tail)
	assert.Equal(t, incStage.out, squareStage.in)
	assert.Equal(t, squareStage, incStage.next)
}

func TestPipelineRunAndStop(t *testing.T) {
	p, out := makeIncSquarePipeline()
	p.Run()
	assert.Empty(t, p.Done())

	testIncSquare(t, p, out)

	p.Stop()
	assert.Zero(t, len(p.Done()))
	p.Wait()
}

func TestPipelineDrop(t *testing.T) {
	p := newPipeline()
	p.addLast(newBufferedStage(1, incOp))
	p.In() <- 1

	assert.NotZero(t, len(p.In()))
	p.Drop()
	assert.Zero(t, len(p.In()))
	p.Drop()
	assert.Zero(t, len(p.In()))
}
