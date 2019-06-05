package piper

import (
	"sync"
)

type Pipeline struct {
	stages []Stage
	wg     *sync.WaitGroup
}

func NewPipeline() *Pipeline {
	return &Pipeline{
		stages: make([]Stage, 0),
		wg:     &sync.WaitGroup{},
	}
}

func (p *Pipeline) AddStage(s Stage) *Pipeline {
	n := len(p.stages)

	if n > 0 {
		p.stages[n-1].Compose(s)
	}

	p.stages = append(p.stages, s)
	return p
}

func (p Pipeline) Run() {
	p.wg.Add(len(p.stages))

	for i := 0; i < len(p.stages); i++ {
		p.stages[i].run(p.wg)
	}
}

func (p Pipeline) Stop() {
	close(p.stages[0].in)
}

func (p Pipeline) Wait() {
	p.wg.Wait()
}

func (p Pipeline) In() chan interface{} {
	return p.stages[0].in
}

func (p Pipeline) Out() chan interface{} {
	return p.stages[len(p.stages)-1].out
}

func (p *Pipeline) Split(ps ...Pipeline) *Pipeline {
	return p.AddStage(newSplitterStage(ps...))
}

func (p *Pipeline) Sink() *Pipeline {
	return p.AddStage(newSinkStage())
}
