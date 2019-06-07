package piper

type SinkOperator func(<-chan interface{})

func makeSink(op SinkOperator) Operator {
	return Operator(func(in <-chan interface{}, out chan<- interface{}) {
		op(in)
	})
}

var defaultSink = makeSink(SinkOperator(func(in <-chan interface{}) {
	for range in {
	}
}))

func defaultSinkStage() *Stage {
	return newStage(defaultSink)
}

func newSinkStage(op SinkOperator) *Stage {
	return newStage(makeSink(op))
}

func newBufferedSinkStage(bufSize int, op SinkOperator) *Stage {
	return newBufferedStage(bufSize, makeSink(op))
}
