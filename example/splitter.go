package main

import (
	"fmt"
	"time"

	"github.com/sebisujar/piper"
)

func main() {
	plus := piper.NewBufferedStage(10, func(in chan interface{}, out chan interface{}) {
		for _n := range in {
			n := _n.(int)
			out <- n + 1
		}
	})

	square := piper.NewSyncStage(func(in chan interface{}, out chan interface{}) {
		for _n := range in {
			n := _n.(int)
			out <- n * n
		}
		time.Sleep(2 * time.Second)
	})

	id := piper.NewSyncStage(func(in chan interface{}, out chan interface{}) {
		for _n := range in {
			out <- _n
		}
	})

	p1 := piper.NewPipeline(id)
	p2 := piper.NewPipeline(square)

	p := piper.NewPipeline(plus).
		Split(p1, p2)

	go func() {
		for x := range p1.Out() {
			fmt.Println("id: ", x)
		}
	}()

	go func() {
		for x := range p2.Out() {
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
	p.Wait()
	fmt.Println("done ")
}
