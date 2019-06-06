package piper

type Operator func(chan interface{}, chan interface{})

type Stage struct {
	in   chan interface{}
	out  chan interface{}
	done chan struct{}

	op   Operator
	next *Stage
}

func (s *Stage) compose(next *Stage) {
	if s.next != nil {
		s.next.compose(next)
	} else {
		s.out = next.in
		s.next = next
	}
}

func (s *Stage) run() {
	go func() {
		defer close(s.out)

		if s.next != nil {
			s.next.run()
		}
		s.op(s.in, s.out)
		s.done <- struct{}{}
	}()
}

func (s *Stage) stop() {
	close(s.in)
}

func (s *Stage) wait() {
	<-s.done

	if s.next != nil {
		s.next.wait()
	}
}

func (s *Stage) Out() chan interface{} {
	if s.next != nil {
		return s.next.Out()
	} else {
		return s.out
	}
}

func NewSyncStage(op Operator) *Stage {
	return &Stage{
		done: make(chan struct{}, 1),
		in:   make(chan interface{}),
		out:  make(chan interface{}),
		next: nil,
		op:   op,
	}
}

func NewBufferedStage(bufSize int, op Operator) *Stage {
	return &Stage{
		done: make(chan struct{}, 1),
		in:   make(chan interface{}, bufSize),
		out:  make(chan interface{}, bufSize),
		next: nil,
		op:   op,
	}
}
