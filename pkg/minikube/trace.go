package minikube

import (
	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/api/global"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var (
	projectID = "priya-wadhwa"
)

// Trace traces minikube start
func Trace() error {
	// Create and register a OpenCensus Stackdriver Trace exporter.
	exporter, err := texporter.NewExporter(texporter.WithProjectID(projectID))

	if err != nil {
		return errors.Wrap(err, "getting exporter")
	}
	tp, err := sdktrace.NewProvider(sdktrace.WithSyncer(exporter))
	if err != nil {
		return errors.Wrap(err, "new provider")
	}
	tp.ApplyConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()})
	global.SetTraceProvider(tp)

	return start()
}
