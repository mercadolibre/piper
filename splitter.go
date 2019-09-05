package piper

func newSplitterStage(ps ...*Pipeline) *stage {
	return newSinkStage(makeSplitter(ps...))
}

func makeSplitter(ps ...*Pipeline) SinkOperator {
	return SinkOperator(func(in <-chan interface{}) {
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
