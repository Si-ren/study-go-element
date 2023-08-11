package main

import (
	"context"
	"flag"
	"log"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
)

var exporterHttpEndpoint = flag.String("exporter-http-endpoint", "localhost:4318", "")
var exporterGRPCEndpoint = flag.String("exporter-grpc-endpoint", "localhost:4317", "")

func main() {
	// tracer 有 tracer provider 初始化
	// 由tracer开启span
	initProvider()
}

func initProvider() {
	// 1.初始化添加到provider资源信息
	res, err := resource.New(context.Background(),
		resource.WithAttributes(
			semconv.ServiceName("traces-basic"),
			semconv.ServiceVersion("1.0.0"),
			attribute.String("env", "dev"),
		),
		resource.WithOS(),
		resource.WithHost())
	if err != nil {
		log.Fatal(err)
		return
	}
	var traceExporter *otlpTrace.Exporter

}
