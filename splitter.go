package piper

func newSplitterStage(stages ...*Stage) *Stage {
	return &Stage{
		in:  make(chan interface{}),
		out: make(chan interface{}),
		op:  makeSplitter(stages...),
	}
}

func makeSplitter(stages ...*Stage) Operator {
	return Operator(func(in chan interface{}, out chan interface{}) {
		for _, p := range stages {
			p.run()
		}

		for msg := range in {
			for _, s := range stages {
				s.in <- msg
			}
		}

		for _, p := range stages {
			p.stop()
			p.wait()
		}
	})
}
