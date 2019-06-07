package piper

type PipelineBuilder struct {
	pipeline *Pipeline
}

func NewBuilder() PipelineBuilder {
	return PipelineBuilder{
		pipeline: newPipeline(),
	}
}

func (builder PipelineBuilder) AddLast(op Operator) PipelineBuilder {
	stage := newStage(op)
	builder.pipeline.addLast(stage)
	return builder
}

func (builder PipelineBuilder) AddLastBuffered(bufSize int, op Operator) PipelineBuilder {
	stage := newBufferedStage(bufSize, op)
	builder.pipeline.addLast(stage)
	return builder
}

func (builder PipelineBuilder) DefaultSink() *Pipeline {
	return builder.pipeline.addLast(defaultSinkStage())
}

func (builder PipelineBuilder) Sink(op SinkOperator) *Pipeline {
	stage := newSinkStage(op)
	return builder.pipeline.addLast(stage)
}

func (builder PipelineBuilder) BufferedSink(bufSize int, op SinkOperator) *Pipeline {
	stage := newBufferedSinkStage(bufSize, op)
	return builder.pipeline.addLast(stage)
}

func (builder PipelineBuilder) Output(op Operator) (*Pipeline, <-chan interface{}) {
	pipeline := builder.AddLast(op).pipeline
	return pipeline, pipeline.tail.out
}

func (builder PipelineBuilder) BufferedOutput(bufSize int, op Operator) (*Pipeline, <-chan interface{}) {
	pipeline := builder.AddLastBuffered(bufSize, op).pipeline
	return pipeline, pipeline.tail.out
}

func (builder PipelineBuilder) Split(ps ...*Pipeline) *Pipeline {
	heads := make([]*Stage, len(ps))
	for i, p := range ps {
		heads[i] = p.head
	}
	stage := newSplitterStage(ps...)

	return builder.pipeline.addLast(stage)
}
