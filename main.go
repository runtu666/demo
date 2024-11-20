/*
 * @Author: tao_yang1 tao_yang1@foxitsoftware.com
 * @Date: 2023-05-06 09:10:27
 * @LastEditors: tao_yang1 tao_yang1@foxitsoftware.com
 * @LastEditTime: 2023-07-25 15:36:20
 * @FilePath: /foxit-otel-go/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func main() {
	//
	os.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "https://otel-collector-azk8s.connectedpdf.com")
	l := log.New(os.Stdout, "", 0)
	client := otlptracehttp.NewClient()
	exporter, err := otlptrace.New(context.TODO(), client)
	if err != nil {
		l.Fatal(err)
	}

	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithBatcher(exporter, trace.WithBatchTimeout(time.Second*5)),
		trace.WithResource(newResource()),
	)
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			l.Fatal(err)
		}
	}()
	otel.SetTracerProvider(tp)

	app := NewApp(os.Stdin, l)
	app.Run(context.Background())

}

func newExporter(w io.Writer) (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human-readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
}

// newResource returns a resource describing this application.
func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("foxit-go-test"),
			semconv.ServiceVersion("v0.1.0"),
			attribute.String("foxit.category", "otel_traces"),
		),
	)
	return r
}
