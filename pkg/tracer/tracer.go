package tracer

import (
	"context"
	"fmt"
	"os"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/trace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var (
	projectID = "priya-wadhwa"
	t         trace.Tracer
)

func init() {
	fmt.Println("Initializing tracer...")
	exporter, err := texporter.NewExporter(texporter.WithProjectID(projectID))

	if err != nil {
		exit(errors.Wrap(err, "getting exporter"))
	}
	tp, err := sdktrace.NewProvider(sdktrace.WithSyncer(exporter))
	if err != nil {
		exit(errors.Wrap(err, "new provider"))
	}
	tp.ApplyConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()})
	global.SetTraceProvider(tp)
	t = global.TraceProvider().Tracer("container-tools")
}

func exit(err error) {
	fmt.Println("failed to initialize: ", err)
	os.Exit(1)
}

// StartSpan starts a new span
func StartSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	return t.Start(ctx, name)
}
