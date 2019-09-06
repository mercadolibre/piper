package piper

// The Pipeline is the main component of piper.
// It is in charge of linking running and stopping all stages added to it.
type Pipeline struct {
	head *stage
	tail *stage
}

func newPipeline() *Pipeline {
	return &Pipeline{
		head: nil,
		tail: nil,
	}
}

func (p *Pipeline) addLast(s *stage) *Pipeline {
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

// Run starts all the stages.
// It creates a goroutine for each stage.
func (p Pipeline) Run() {
	for s := p.head; s != nil; s = s.next {
		go s.run()
	}
}

// Stop closes the pipeline.
// The stop is propagated through each stage.
func (p Pipeline) Stop() {
	p.head.stop()
}

// Wait blocks until all messages are consumed.
func (p Pipeline) Wait() {
	<-p.Done()
}

// Done returns a channel that is closed once all messages are consumed.
func (p Pipeline) Done() <-chan struct{} {
	return p.tail.done
}

// In return the input channel for the pipeline.
func (p Pipeline) In() chan<- interface{} {
	return p.head.in
}

// Drop removes consumes the oldest element in the input channel.
// This serves as a way to make room in the empty channel when blocking is not an option.
func (p Pipeline) Drop() {
	select {
	case <-p.head.in:
	default:
	}
}
