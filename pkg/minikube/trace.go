package minikube

import (
	"os"
	"runtime"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/pkg/errors"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
)

// Measures for the stats quickstart.
var (
	// The latency in seconds
	mLatencyS = stats.Float64("repl/latency", "The latency in seconds per REPL loop", stats.UnitSeconds)
)

// Trace traces minikube start
func Trace() error {
	// Create and register a OpenCensus Stackdriver Trace exporter.
	labels := &stackdriver.Labels{}
	labels.Set("os", runtime.GOOS, "operating system")
	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID:               os.Getenv("priya-wadhwa"),
		DefaultMonitoringLabels: labels,
	})
	if err != nil {
		return errors.Wrap(err, "getting exporter")
	}
	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	// Export to Stackdriver Monitoring.
	if err = exporter.StartMetricsExporter(); err != nil {
		return errors.Wrap(err, "starting metrics exporter")
	}

	// Subscribe views to see stats in Stackdriver Monitoring.
	if err := enableViews(); err != nil {
		return errors.Wrap(err, "enabling views")
	}

	return nil
}

func enableViews() error {
	startTimeView := &view.View{
		Name:        "minikube/startTime",
		Measure:     mLatencyS,
		Description: "minikube start times",
		Aggregation: view.LastValue(),
	}

	return view.Register(startTimeView)
}
