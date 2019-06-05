package piper

func newSplitterStage(ps ...Pipeline) Stage {
	return Stage{
		in:  make(chan interface{}),
		out: make(chan interface{}),
		op:  makeSplitter(ps...),
	}
}

func makeSplitter(ps ...Pipeline) Operator {
	return Operator(func(in chan interface{}, out chan interface{}) {
		for _, p := range ps {
			p.Run()
		}

		for msg := range in {
			for _, p := range ps {
				p.In() <- msg
			}
		}

		for _, p := range ps {
			p.Stop()
			p.Wait()
		}
	})
}
