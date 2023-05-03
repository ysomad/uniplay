package app

import (
	"context"
	"time"

	"github.com/exaring/otelpgx"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/ysomad/uniplay/internal/config"
)

type openTelemetry struct {
	AppTracer trace.Tracer
	PgxTracer *otelpgx.Tracer

	// functions which have to called before or after app is down
	CleanupFuncs [2]func(context.Context) error
}

func newOpenTelemetry(conf *config.Config) (openTelemetry, error) {
	ol := openTelemetry{}
	res := newResource(conf.App)

	jaegerExp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(conf.Jaeger.Endpoint)))
	if err != nil {
		return openTelemetry{}, err
	}

	tracerProvider := newTracerProvider(res, jaegerExp)

	prometheusExp, err := prometheus.New()
	if err != nil {
		return openTelemetry{}, err
	}

	meterProvider, err := newMeterProvider(res, prometheusExp)
	if err != nil {
		return openTelemetry{}, err
	}

	ol.AppTracer = otel.GetTracerProvider().Tracer("uniplay")
	ol.PgxTracer = otelpgx.NewTracer(otelpgx.WithTrimSQLInSpanName())
	ol.CleanupFuncs = [2]func(context.Context) error{
		tracerProvider.Shutdown,
		meterProvider.Shutdown,
	}

	return ol, nil
}

func newResource(conf config.App) *resource.Resource {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(conf.Name),
		semconv.ServiceVersionKey.String(conf.Ver),
		attribute.String("environment", conf.Environment),
	)
}

// newTracerProvider returns and registers configured tracer provider.
//
// docs: https://opentelemetry.io/docs/instrumentation/go/exporting_data/
func newTracerProvider(r *resource.Resource, exp sdktrace.SpanExporter) *sdktrace.TracerProvider {
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)

	otel.SetTracerProvider(tp)

	return tp
}

// newMeterProvider sets global meter provider and returns it.
func newMeterProvider(res *resource.Resource, exp *prometheus.Exporter) (*metric.MeterProvider, error) {
	mp := metric.NewMeterProvider(
		metric.WithReader(exp),
		metric.WithResource(res),
	)

	global.SetMeterProvider(mp)

	if err := runtime.Start(runtime.WithMinimumReadMemStatsInterval(time.Second)); err != nil {
		return nil, err
	}

	return mp, nil
}
