package piper

// The SinkOperator is a special type of operator used to end a pipeline
// without producing and output.
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

func defaultSinkStage() *stage {
	return newStage(defaultSink)
}

func newSinkStage(op SinkOperator) *stage {
	return newStage(makeSink(op))
}

func newBufferedSinkStage(bufSize int, op SinkOperator) *stage {
	return newBufferedStage(bufSize, makeSink(op))
}
