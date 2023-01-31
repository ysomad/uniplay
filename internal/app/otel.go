package app

import (
	"context"
	"time"

	"github.com/ysomad/uniplay/internal/config"

	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func newResource(conf config.App) *resource.Resource {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(conf.Name),
		semconv.ServiceVersionKey.String(conf.Ver),
		attribute.String("environment", conf.Environment),
	)
}

// newTracerProvider returns and registers cnfigured tracer provider.
//
// docs: https://opentelemetry.io/docs/instrumentation/go/exporting_data/
func newTracerProvider(r *resource.Resource, exp sdktrace.SpanExporter) func(context.Context) error {
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)

	otel.SetTracerProvider(tp)

	return tp.Shutdown
}

func newJaegerExporter(conf config.Jaeger) (*jaeger.Exporter, error) {
	return jaeger.New(
		jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(conf.Endpoint)),
	)
}

func newMeterProvider(res *resource.Resource, exp *prometheus.Exporter) (func(context.Context) error, error) {
	mp := metric.NewMeterProvider(
		metric.WithReader(exp),
		metric.WithResource(res),
	)

	global.SetMeterProvider(mp)

	if err := runtime.Start(runtime.WithMinimumReadMemStatsInterval(time.Second)); err != nil {
		return nil, err
	}

	return mp.Shutdown, nil
}
