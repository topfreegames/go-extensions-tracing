package tracing

import (
	"context"

	"github.com/opentracing/opentracing-go"
)

func Trace(ctx context.Context, operationName string, tags opentracing.Tags, f func()) {
	var parent opentracing.SpanContext
	if span := opentracing.SpanFromContext(ctx); span != nil {
		parent = span.Context()
	}
	reference := opentracing.ChildOf(parent)
	if tags == nil {
		tags = opentracing.Tags{}
	}
	span := opentracing.StartSpan(operationName, reference, tags)
	defer span.Finish()
	defer LogPanic(span)
	f()
}
