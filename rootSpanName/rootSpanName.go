package rootspanName

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor/processorhelper"
)

type rootspanName struct {
	config Config
}

func newTracesProcessor(ctx context.Context, set component.ProcessorCreateSettings, cfg *Config, nextConsumer consumer.Traces) (component.TracesProcessor, error) {
	sn := &rootspanName{}

	return processorhelper.NewTracesProcessor(
		ctx,
		set,
		cfg,
		nextConsumer,
		sn.processTraces,
		processorhelper.WithCapabilities(consumer.Capabilities{MutatesData: true}))
}

func (sn *rootspanName) processTraces(_ context.Context, td ptrace.Traces) (ptrace.Traces, error) {

	rootSpanName := ""
	spanCount := 0

	rss := td.ResourceSpans()
	for i := 0; i < rss.Len(); i++ {
		rs := rss.At(i)
		ilss := rs.ScopeSpans() // ilss = scopeSpans
		for j := 0; j < ilss.Len(); j++ {
			ils := ilss.At(j)
			spans := ils.Spans()
			//library := ils.Scope()
			for k := 0; k < spans.Len(); k++ {
				spanCount = spanCount + 1
				s := spans.At(k)

				spanID := s.SpanID().HexString()
				traceID := s.TraceID().HexString()
				spanName := s.Name()
				parentspanID := s.ParentSpanID().HexString()
				//start_nano := s.StartTimestamp().AsTime().UnixNano()
				//end_time := s.EndTimestamp().AsTime().UnixNano()

				start_time := s.StartTimestamp().String()
				end_time := s.EndTimestamp().String()

				fmt.Printf("SpanID = " + spanID + "\n" + "traceID = " + traceID + "\n")
				fmt.Printf("start_time:" + start_time + "\n" + "end_time:" + end_time + "\n")

				if parentspanID == "" {
					rootSpanName = spanName
					fmt.Printf("Set rootSpanName = " + rootSpanName)
				}
			}

		}

	}
	return td, nil

}
