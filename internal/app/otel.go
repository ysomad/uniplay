package app

import (
	"context"

	"github.com/ysomad/uniplay/internal/config"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func newJaegerExporter(conf config.Jaeger) (*jaeger.Exporter, error) {
	return jaeger.New(
		jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(conf.Endpoint)),
	)
}

// newTracerProvider returns and registers cnfigured tracer provider.
//
// docs: https://opentelemetry.io/docs/instrumentation/go/exporting_data/
func newTracerProvider(conf config.App, exp sdktrace.SpanExporter) func(context.Context) error {
	r := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(conf.Name),
		semconv.ServiceVersionKey.String(conf.Ver),
		attribute.String("environment", conf.Environment),
	)

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		// sdktrace.WithBatcher(exp),
		sdktrace.WithSyncer(exp),
		sdktrace.WithResource(r),
	)

	otel.SetTracerProvider(tp)

	return tp.Shutdown
}
