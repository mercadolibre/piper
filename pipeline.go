package piper

type Pipeline struct {
	head *Stage
}

func NewPipeline(head *Stage) Pipeline {
	return Pipeline{
		head: head,
	}
}

func (p Pipeline) AddStage(s *Stage) Pipeline {
	p.head.compose(s)
	return p
}

func (p Pipeline) Run() {
	p.head.run()
}

func (p Pipeline) Stop() {
	p.head.stop()
}

func (p Pipeline) Wait() {
	p.head.wait()
}

func (p Pipeline) In() chan interface{} {
	return p.head.in
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
