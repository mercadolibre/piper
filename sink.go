package piper

var sinkOperator = Operator(func(in chan interface{}, out chan interface{}) {
	for range in {
	}
})

func newSinkStage() *Stage {
	return &Stage{
		in:  make(chan interface{}),
		out: make(chan interface{}),
		op:  sinkOperator,
	}
}
