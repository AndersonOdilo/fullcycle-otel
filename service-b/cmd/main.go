package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/AndersonOdilo/otel/service-b/configs"
	"github.com/AndersonOdilo/otel/service-b/internal/api"
	"github.com/AndersonOdilo/otel/service-b/internal/infra/web"
	"github.com/AndersonOdilo/otel/service-b/internal/infra/web/webserver"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)


func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	shutdown, err := InitProvider(configs.OtelServiceName, configs.OtelExporterOtlpEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal("failed to shutdown TracerProvider: %w", err)
		}
	}()

	
	tracer := otel.Tracer("temp-system-api")

	fmt.Println("Starting web server on port", configs.WebServerPort)
	locationRepository := api.NewLocationRepository();
	tempRepository := api.NewTempRepository(*configs);

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webTempHandler := web.NewWebTempHandler(locationRepository, tempRepository, tracer);
	webserver.AddHandler("GET /temp/{cep}", webTempHandler.Get)
	go webserver.Start()

	select {
	case <-sigCh:
		log.Println("Shutting down gracefully, CTRL+c pressed...")
	case <-ctx.Done():
		log.Println("Shutting down due other reason...")
	}

	_, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
}


func InitProvider(serviceName, collectorURL string) (func(context.Context) error, error) {
	ctx := context.Background()

	res, err := resource.New(
		ctx,
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}
	
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure(),otlptracegrpc.WithTimeout(20*time.Second), otlptracegrpc.WithEndpoint(collectorURL));
	
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)

	otel.SetTextMapPropagator(propagation.TraceContext{})

	return tracerProvider.Shutdown, nil
}