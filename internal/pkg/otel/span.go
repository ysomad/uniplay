package otel

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// StartTrace starting trace with global provider.
func StartTrace(ctx context.Context, tracerName, spanName string) (context.Context, trace.Span) {
	return otel.GetTracerProvider().Tracer(tracerName).Start(ctx, spanName)
}
