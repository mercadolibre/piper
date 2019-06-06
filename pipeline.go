package piper

type Pipeline struct {
	head *Stage
	tail *Stage
}

func NewPipeline(s *Stage) Pipeline {
	return Pipeline{
		head: s,
		tail: s,
	}
}

func (p Pipeline) AddStage(s *Stage) Pipeline {
	p.tail.out = s.in
	p.tail.next = s
	p.tail = s

	return p
}

func (p Pipeline) Run() {
	p.head.run()
}

func (p Pipeline) Stop() {
	p.head.stop()
}

func (p Pipeline) Done() chan struct{} {
	return p.tail.done
}

func (p Pipeline) In() chan interface{} {
	return p.head.in
}

func (p Pipeline) Out() chan interface{} {
	return p.tail.out
}

func (p Pipeline) Split(ps ...Pipeline) Pipeline {
	heads := make([]*Stage, len(ps))
	for i, p := range ps {
		heads[i] = p.head
	}

	return p.AddStage(newSplitterStage(heads...))
}

func (p Pipeline) Sink() Pipeline {
	return p.AddStage(newSinkStage())
}
