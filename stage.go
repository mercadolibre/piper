package piper

import (
	"sync"
)

type Operator func(chan interface{}, chan interface{})

type Stage struct {
	in  chan interface{}
	out chan interface{}
	op  Operator
}

func (s *Stage) Compose(next Stage) {
	s.out = next.in
}

func (s *Stage) run(wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		defer close(s.out)

		s.op(s.in, s.out)
	}()
}

func NewSyncStage(op Operator) Stage {
	return Stage{
		in:  make(chan interface{}),
		out: make(chan interface{}),
		op:  op,
	}
}

func NewBufferedStage(bufSize int, op Operator) Stage {
	return Stage{
		in:  make(chan interface{}, bufSize),
		out: make(chan interface{}, bufSize),
		op:  op,
	}
}
