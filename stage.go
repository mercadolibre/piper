package piper

// An Operator is function type used to add logic to each pipeline stage.
// Each Operator takes an input and an output channel.
type Operator func(<-chan interface{}, chan<- interface{})

type stage struct {
	in   chan interface{}
	out  chan interface{}
	done chan struct{}

	op   Operator
	next *stage
}

func (s *stage) run() {
	defer close(s.out)

	s.op(s.in, s.out)
	s.done <- struct{}{}
}

func (s *stage) stop() {
	close(s.in)
}

func newStage(op Operator) *stage {
	return &stage{
		done: make(chan struct{}, 1),
		in:   make(chan interface{}),
		out:  make(chan interface{}),
		next: nil,
		op:   op,
	}
}

func newBufferedStage(bufSize int, op Operator) *stage {
	return &stage{
		done: make(chan struct{}, 1),
		in:   make(chan interface{}, bufSize),
		out:  make(chan interface{}, bufSize),
		next: nil,
		op:   op,
	}
}
