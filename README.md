# Piper
Hello, this is Piper. 

Piper is your friend. 

Piper is also a pipeline library for Go heavily influenced by [whiskybadger's article](https://whiskybadger.io/post/introducing-go-pipeline/) and [Netty](https://netty.io/)

## Operators
Operators are the basic building block for the pipeline. These come in two flavors: regular and sink.
The former takes an input and an output channel and can be used either as an intermediate or ouput stage.
The latter is used to end a pipeline that doesn't have an output.

```go
var myOp = piper.Operator(func(in <-chan interface{}, out chan<- interface{}){
    // cool stuff
})

var mySinkOp = piper.SinkOperator(func(in <-chan interface{}){
    // cool stuff but without an output
})
```

Operators are later wrapped into Stages to be added to the pipeline.

## What is a Pipeline?
A pipeline is a list of Stages with an input channel and optionally an output

### How do i build one?
With a PipelineBuilder of course! The builder provides a way to add stages to the pipeline. 
These can buffered or not. This defines the way the entry channel will work.

```go
// a sample pipeline with an output
pipeline, out := piper.NewBuilder().
    AddLastBuffered(1000, anOperator).
    Output(anotherOperator)

// this one does not output anything
pipeline := piper.NewBuilder().
    AddLast(anOperator).
    Sink(sinkOperator)

```
All building functions come with its buffered flavor.

### Default Sink
Piper also includes the default sink stage. This receives messages and discards them
```go
pipeline := piper.NewBuilder().
    AddLast(anOperator).
    DefaultSink()
```

### Splitting
Splitting is a way to create more complex pipelines. This works as a sink stage that ends the initial pipeline
and forwards the output to a list of pipelnes
```go
pipeline := piper.NewBuilder().
    Split(aPipeline, anotherPipeline, ...)
```
The forwaring is made sequentially, so this could cause a starvation if one of the pipeline was stalled.


## Ok, i have my pipeline. How do i use it?
```go
pipeline.Run() // This starts the pipeline, duh

// You should stop the pipeline when you are done with it
// multiple stops can cause a panic
pipline.Stop()

// if you want to wait for the pipeline to consume all remaining messages:
pipeline.Wait()

// or maybe you want to wait for other things as well
<-pipeline.Done()
```

## A *"fun"* example
```go
package main

import (
	"fmt"
	"time"

	"github.com/sebisujar/piper"
)

func main() {
	plusOperator := piper.Operator(func(in <-chan interface{}, out chan<- interface{}) {
		for _n := range in {
			n := _n.(int)
			out <- n + 1
		}
	})

	squareOperator := piper.Operator(func(in <-chan interface{}, out chan<- interface{}) {
		for _n := range in {
			n := _n.(int)
			out <- n * n
		}
		time.Sleep(2 * time.Second)
	})

	idOperator := piper.Operator(func(in <-chan interface{}, out chan<- interface{}) {
		for _n := range in {
			out <- _n
		}
	})

	idPipeline, idOut := piper.NewBuilder().Output(idOperator) // complex pipeline
	squarePipeline, squareOut := piper.NewBuilder().Output(squareOperator) // much more complex pipeline

	p := piper.NewBuilder().
		AddLast(plusOperator).
		Split(idPipeline, squarePipeline)

	go func() {
		for x := range idOut {
			fmt.Println("id pipeline: ", x)
		}
	}()

	go func() {
		for x := range squareOut {
			fmt.Println("square pipeline: ", x)
		}
	}()

	p.Run()
	in := p.In()
	in <- 1
	in <- 2
	in <- 3

	p.Stop()
    p.Wait()
   	fmt.Println("bye")
}

```
