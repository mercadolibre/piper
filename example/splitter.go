package main

import (
	"fmt"
	"time"

	"github.com/sebisujar/piper"
)

func main() {
	plus := piper.Operator(func(in chan interface{}, out chan interface{}) {
		for _n := range in {
			n := _n.(int)
			out <- n + 1
		}
	})

	square := piper.Operator(func(in chan interface{}, out chan interface{}) {
		for _n := range in {
			n := _n.(int)
			out <- n * n
		}
		time.Sleep(2 * time.Second)
	})

	id := piper.Operator(func(in chan interface{}, out chan interface{}) {
		for _n := range in {
			out <- _n
		}
	})

	p1, out1 := piper.NewBuilder().Output(id)
	p2, out2 := piper.NewBuilder().Output(square)

	p := piper.NewBuilder().
		AddLastBuffered(10, plus).
		Split(p1, p2)

	go func() {
		for x := range out1 {
			fmt.Println("id: ", x)
		}
	}()

	go func() {
		for x := range out2 {
			fmt.Println("square: ", x)
		}
	}()

	p.Run()
	in := p.In()
	in <- 1
	in <- 2
	in <- 3

	p.Stop()
	fmt.Println("will wait")
	<-p.Done()
	fmt.Println("done ")
}
