# Piper
Hello, this is Piper. 

Piper is your friend. 

Piper is also a pipeline library for Go heavily influenced by [whiskybadger's article](https://whiskybadger.io/post/introducing-go-pipeline/) and [Netty](https://netty.io/)

## TL;DR
```go

```

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

Piper also includes the default sink stage. This receives messages and discards them

## What is a Pipeline?
A pipeline is a list of Stages with an input channel and optionally an output

### How do i build one?
With a PipelineBuilder of course! The builder provides a way to add stages to the pipeline. 
These can buffered or not. This defines the way the entry channel will work.

```go
// a sample pipeline with an output
pipeline, out := piper.NewBuilder().
    AddLastBuffered(1000, incOperator).
    Output(squareOperator)

// this one does not output anything
pipeline := piper.NewBuilder().
    AddLast(incOperator).
    Sink(doSomethingOperator)

```

