package main

import (
	"context"
	"flag"
	"log"
	"time"
	"traces-basic/bus"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
)

// 链路追踪：记录每个函数的执行时间，并且记录其中的log

var exporterHttpEndpoint = flag.String("exporter-http-endpoint", "localhost:4318", "")
var exporterGRPCEndpoint = flag.String("exporter-grpc-endpoint", "localhost:4317", "")

func main() {
	flag.Parse()
	// tracer 由 tracer provider 初始化
	// 由tracer开启span
	shutdown := initProvider(*exporterHttpEndpoint, *exporterGRPCEndpoint)
	defer func() {
		shutdown(context.Background())
	}()
	tracer := otel.Tracer("main-tracer")
	ctx := context.Background()
	ctx, span := tracer.Start(ctx, "main", trace.WithAttributes(attribute.String("package", "main")))
	defer span.End()
	busTracer := otel.Tracer("bus-tracer")
	busObject := bus.NewBus(busTracer)
	for i := 0; i < 5; i++ {
		busObject.Sum(ctx, i, i+1)
		busObject.Product(ctx, i, i+1)
		<-time.After(300 * time.Millisecond)
	}

}

func initProvider(httpEndpoint, grpcEndpoint string) func(ctx context.Context) error {
	// 初始化traceExpoter
	var traceExporter *otlptrace.Exporter
	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if grpcEndpoint != "" {
			// grpc需要先初始化grpc连接
			conn, err := grpc.DialContext(ctx, grpcEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
			if err != nil {
				log.Fatal(err)
			}
			traceExporter, err = otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure(), otlptracegrpc.WithGRPCConn(conn))
			if err != nil {
				log.Fatal(err)
			}

		} else if httpEndpoint != "" {
			var err error
			traceExporter, err = otlptracehttp.New(ctx, otlptracehttp.WithInsecure(), otlptracehttp.WithEndpoint(httpEndpoint))
			if err != nil {
				log.Fatal(err)
			}

		} else {
			log.Fatal("No traceExporter exists")
		}
		// 1.初始化要添加到provider的resource信息
		// 2.以下代码提供的resource是这样的
		// 	env dev
		// host.name GWSHALTQR03
		// os.description Windows 10 Pro 22H2 (2009) [Version 10.0.19045.3324]
		// os.type windows
		// service.version 1.0.0
		res, err := resource.New(context.Background(),
			resource.WithAttributes(
				semconv.ServiceName("traces-basic"),
				semconv.ServiceVersion("1.0.0"),
				attribute.String("env", "dev"),
			),
			resource.WithOS(),
			resource.WithHost(),
		)

		if err != nil {
			log.Fatal(err)
			return nil
		}
		// 收集span，当收集到一定程度时，再发送到exporter
		bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
		traceProvider := sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithResource(res),
			sdktrace.WithSpanProcessor(bsp))
		otel.SetTracerProvider(traceProvider)
		otel.SetTextMapPropagator(propagation.Baggage{})
		// otel.SetTextMapPropagator(propagation.TraceContext{})
		return traceProvider.Shutdown
	}

}
