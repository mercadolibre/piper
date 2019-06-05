package piper

/*
import (
	"fmt"
)

type Operator func(chan interface{}, chan interface{})

type Stage struct {
	in  chan interface{}
	out chan interface{}
	op  Operator
}

func NewSyncStage(op Operator) Stage {
	return Stage{
		in:  make(chan interface{}),
		out: make(chan interface{}),
		op:  op,
	}
}

func (s *Stage) Compose(next Stage) {
	s.out = next.in
}

func (s *Stage) Run() {
	go func() {
		s.op(s.in, s.out)
		close(s.out)
	}()
}

type Pipeline struct {
	stages []Stage
	in     chan interface{}
	out    chan interface{}
}

func NewPipeline(stages ...Stage) Pipeline {
	return Pipeline{
		stages: stages,
		in:     make(chan interface{}, 3),
		out:    make(chan interface{}, 3),
	}
}

func (p Pipeline) Run() {
	var (
		first *Stage
		last  *Stage
	)
	first = &p.stages[0]
	last = &p.stages[len(p.stages)-1]

	first.in = p.in
	last.out = p.out

	for i := 0; i < len(p.stages)-1; i++ {
		p.stages[i].Compose(p.stages[i+1])
	}

	for i := 0; i < len(p.stages); i++ {
		p.stages[i].Run()
	}
}

func main() {
	square := NewSyncStage(func(in chan interface{}, out chan interface{}) {
		for _n := range in {
			n := _n.(int)
			out <- n * n
		}
	})

	plus := NewSyncStage(func(in chan interface{}, out chan interface{}) {
		for _n := range in {
			n := _n.(int)
			out <- n + 1
		}
	})

	flow := NewPipeline(plus, square)

	flow.in <- 1
	flow.in <- 2
	flow.in <- 3
	close(flow.in)

	flow.Run()

	for x := range flow.out {
		fmt.Println(x)
	}
}

*/
