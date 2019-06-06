package main

import (
	"fmt"
	"time"

	"github.com/sebisujar/piper"
)

func main() {
	square := piper.NewSyncStage(func(in chan interface{}, out chan interface{}) {
		for _n := range in {
			n := _n.(int)
			out <- n * n
		}
	})

	plus := piper.NewSyncStage(func(in chan interface{}, out chan interface{}) {
		for _n := range in {
			n := _n.(int)
			out <- n + 1
		}
		time.Sleep(time.Second)
	})

	output := piper.NewSyncStage(func(in chan interface{}, out chan interface{}) {
		for _n := range in {
			fmt.Println(_n)
		}
	})

	p := piper.NewPipeline(plus).
		AddStage(square).
		AddStage(output)

	p.Run()
	in := p.In()
	in <- 1
	in <- 2
	in <- 3
	p.Stop()

	p.Wait()
	fmt.Println("done")
}
