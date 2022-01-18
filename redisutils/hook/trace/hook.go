package trace

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/hashicorp/go-multierror"
	"github.com/soyacen/bytebufferpool"
	"github.com/soyacen/goutils/stacktraceutils"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

type Hook struct {
	tracer trace.Tracer
	db     int
}

func New(tracer trace.Tracer, db int) redis.Hook {
	return &Hook{
		tracer: tracer,
		db:     db,
	}
}

func (hook *Hook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	ctx, span := hook.tracer.Start(ctx, cmd.FullName(), trace.WithSpanKind(trace.SpanKindInternal))
	lines := strings.Split(stacktraceutils.CallersFrames(4), "\n")
	caller := lines[0] + "@" + strings.TrimSpace(lines[1])
	span.SetAttributes(
		semconv.DBSystemRedis,
		semconv.DBRedisDBIndexKey.Int(hook.db),
		semconv.DBRedisDBIndexKey.Int(hook.db),
		semconv.DBStatementKey.String(fmt.Sprintf("%v", cmd.Args())),
		attribute.Key("caller").String(caller),
	)
	return ctx, nil
}

func (hook *Hook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	span := trace.SpanFromContext(ctx)
	defer span.End()
	if err := cmd.Err(); err != nil {
		span.RecordError(err, trace.WithTimestamp(time.Now()))
	}
	return nil
}

func (hook *Hook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	ctx, span := hook.tracer.Start(ctx, "pipeline")
	lines := strings.Split(stacktraceutils.CallersFrames(4), "\n")
	caller := lines[0] + "@" + strings.TrimSpace(lines[1])
	span.SetAttributes(
		semconv.DBSystemRedis,
		semconv.DBRedisDBIndexKey.Int(hook.db),
		semconv.DBRedisDBIndexKey.Int(hook.db),
		semconv.DBStatementKey.String(getPipelineStatement(cmds)),
		attribute.Key("caller").String(caller),
		attribute.Key("db.cmd.count").Int(len(cmds)),
	)
	return ctx, nil
}

func getPipelineStatement(cmds []redis.Cmder) string {
	byteBuffer := bytebufferpool.Get()
	defer byteBuffer.Free()
	for _, cmd := range cmds {
		byteBuffer.WriteString(fmt.Sprintf("%v ", cmd.Args()))
	}
	return byteBuffer.String()
}

func (hook *Hook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	span := trace.SpanFromContext(ctx)
	defer span.End()
	var err error
	for _, cmd := range cmds {
		if e := cmd.Err(); e != nil {
			e = multierror.Append(err, e)
		}
	}
	if err != nil {
		span.RecordError(err, trace.WithTimestamp(time.Now()))
	}
	return nil
}
