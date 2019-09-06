package piper

// The PipelineBuilder is the tool for creating pipelines.
// It starts with an empty pipline and ends once a Sink or output stage is added.
type PipelineBuilder struct {
	pipeline *Pipeline
}

// NewBuilder returns a new PipelineBuilder with a fresh Pipeline.
func NewBuilder() PipelineBuilder {
	return PipelineBuilder{
		pipeline: newPipeline(),
	}
}

// AddLast creates a new stage from an Operator and adds it to the end of the Pipeline.
func (builder PipelineBuilder) AddLast(op Operator) PipelineBuilder {
	stage := newStage(op)
	builder.pipeline.addLast(stage)
	return builder
}

// AddLastBuffered does the same as AddLast but makes the input channel from that stage
// have a buffer of the given size.
func (builder PipelineBuilder) AddLastBuffered(bufSize int, op Operator) PipelineBuilder {
	stage := newBufferedStage(bufSize, op)
	builder.pipeline.addLast(stage)
	return builder
}

// DefaultSink adds a stage that consumes every message and return the resulting pipeline.
func (builder PipelineBuilder) DefaultSink() *Pipeline {
	return builder.pipeline.addLast(defaultSinkStage())
}

// Sink adds a SinkStage of the Pipeline from the SinkOperator given,
// and returns the resulting Pipeline.
func (builder PipelineBuilder) Sink(op SinkOperator) *Pipeline {
	stage := newSinkStage(op)
	return builder.pipeline.addLast(stage)
}

// BufferedSink does the same as Sink but makes the input channel from that stage
// have a buffer of the given size.
func (builder PipelineBuilder) BufferedSink(bufSize int, op SinkOperator) *Pipeline {
	stage := newBufferedSinkStage(bufSize, op)
	return builder.pipeline.addLast(stage)
}

// Output works as addLast but it returns the resulting pipeline and
// an output channel to consume the outputs.
func (builder PipelineBuilder) Output(op Operator) (*Pipeline, <-chan interface{}) {
	pipeline := builder.AddLast(op).pipeline
	return pipeline, pipeline.tail.out
}

// BufferedOutput works as Output but makes the input channel from that stage
// have a buffer of the given size
func (builder PipelineBuilder) BufferedOutput(bufSize int, op Operator) (*Pipeline, <-chan interface{}) {
	pipeline := builder.AddLastBuffered(bufSize, op).pipeline
	return pipeline, pipeline.tail.out
}

// Split makes attaches the input of a list of Pipelines to the output of the current one.
// This returns a Pipeline that has no output itself but can produce an output in some
// of the pipelines provided.
//
// Once a Pipeline is attached it should only be run from the main Pipeline.
func (builder PipelineBuilder) Split(ps ...*Pipeline) *Pipeline {
	heads := make([]*stage, len(ps))
	for i, p := range ps {
		heads[i] = p.head
	}
	stage := newSplitterStage(ps...)

	return builder.pipeline.addLast(stage)
}
