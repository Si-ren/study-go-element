package bus

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"strconv"
	"time"
)

func main() {

}

type Bus interface {
	Sum(ctx context.Context, a, b int) int
	Product(ctx context.Context, a, b int) int
}

type bus struct {
	tracer trace.Tracer
}

func NewBus(tracer trace.Tracer) Bus {
	return &bus{
		tracer: tracer,
	}
}
func (bus bus) Sum(ctx context.Context, a, b int) int {
	c := a + b
	ctx, span := bus.tracer.Start(ctx, "sum", trace.WithAttributes(attribute.String("c", strconv.Itoa(c))))
	defer span.End()
	<-time.After(time.Second)
	return c
}

func (bus bus) Product(ctx context.Context, a, b int) int {
	c := a * b
	ctx, span := bus.tracer.Start(ctx, "product", trace.WithAttributes(attribute.String("c", strconv.Itoa(c))))
	defer span.End()
	<-time.After(2 * time.Second)
	return c
}
