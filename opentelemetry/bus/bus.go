package bus

import (
	"context"
	"errors"
	"strconv"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
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
	_, span := bus.tracer.Start(ctx, "sum", trace.WithAttributes(attribute.String("c", strconv.Itoa(c))))
	defer span.End()
	<-time.After(time.Second)
	return c
}

func (bus bus) Product(ctx context.Context, a, b int) int {
	c := a * b
	_, span := bus.tracer.Start(ctx, "product", trace.WithAttributes(attribute.String("c", strconv.Itoa(c))))
	span.SetStatus(codes.Error, "this is product error")
	span.RecordError(errors.New("record error in product"))
	defer span.End()
	<-time.After(2 * time.Second)
	return c
}
