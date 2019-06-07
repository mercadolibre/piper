package piper

type Pipeline struct {
	head *Stage
	tail *Stage
}

func newPipeline() *Pipeline {
	return &Pipeline{
		head: nil,
		tail: nil,
	}
}

func (p *Pipeline) addLast(s *Stage) *Pipeline {
	if p.head == nil {
		p.head = s
		p.tail = s
	} else {
		p.tail.out = s.in
		p.tail.next = s
		p.tail = s
	}

	return p
}

func (p Pipeline) Run() {
	p.head.run()
}

func (p Pipeline) Stop() {
	p.head.stop()
}

func (p Pipeline) Wait() {
	<-p.Done()
}

func (p Pipeline) Done() <-chan struct{} {
	return p.tail.done
}

func (p Pipeline) In() chan<- interface{} {
	return p.head.in
}
