package tracing

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSdk "go.opentelemetry.io/otel/sdk/trace"
	semConv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/xiaohubai/go-grpc-layout/configs/conf"
)

func RegisterTracer(endpoint string, g *conf.Global) error {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint)))
	if err != nil {
		return err
	}
	tp := traceSdk.NewTracerProvider(
		traceSdk.WithSampler(traceSdk.ParentBased(traceSdk.TraceIDRatioBased(1.0))),
		traceSdk.WithBatcher(exp),
		traceSdk.WithResource(resource.NewSchemaless(
			semConv.ServiceNameKey.String(g.AppName),
			semConv.ServiceVersionKey.String(g.Version),
			semConv.ServiceInstanceIDKey.String(g.Id),
			attribute.String("env", g.Env),
		)),
	)

	otel.SetTracerProvider(tp)
	return nil
}

// NewSpan creates a span using default tracer.
func NewSpan(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return otel.Tracer("").Start(ctx, spanName, opts...)
}

func TraceID(ctx context.Context) string {
	if span := trace.SpanContextFromContext(ctx); span.HasTraceID() {
		return span.TraceID().String()
	}
	return ""
}

func SpanID(ctx context.Context) string {
	if span := trace.SpanContextFromContext(ctx); span.HasSpanID() {
		return span.SpanID().String()
	}
	return ""
}
