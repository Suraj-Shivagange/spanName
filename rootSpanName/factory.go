package rootspanName

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
)

const (
	// typeStr is the value of "type" Span processor in the configuration.
	typeStr = "rootspanName_"
	// The stability level of the processor.
	stability = component.StabilityLevelAlpha
)

func NewFactory() component.ProcessorFactory {
	return component.NewProcessorFactory(
		typeStr,
		createDefaultConfig,
		component.WithTracesProcessor(createTracesProcessor, stability))
}

func createDefaultConfig() config.Processor {
	return &Config{
		ProcessorSettings: config.NewProcessorSettings(config.NewComponentID(typeStr)),
	}
}

func createTracesProcessor(
	ctx context.Context,
	set component.ProcessorCreateSettings,
	cfg config.Processor,
	nextConsumer consumer.Traces) (component.TracesProcessor, error) {
	return newTracesProcessor(ctx, set, cfg.(*Config), nextConsumer) // return the fuction name from the rootSpanName.go file
}
