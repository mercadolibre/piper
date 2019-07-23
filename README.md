# Piper
Hello, this is Piper. 

Piper is your friend. 

Piper is also a pipeline lib for Go heavily influenced by [whiskybadger's article](https://whiskybadger.io/post/introducing-go-pipeline/) and [Netty](https://netty.io/)

## TL;DR
```go
```

## Operators
Operators are the basic building block for the pipeline. These come in two flavors: regular and sink.
The former takes an input and an output channel and can be used either as an intermediate or ouput stage.
The latter is used to end a pipeline that doesn't have an output.

```go
var myOp = piper.Operator()
var mySinkOp = piper.SinkOperator()
```

## What is a Pipeline?
A pipeline is a collection of Stage

### How do i build one?
