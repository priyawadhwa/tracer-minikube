package minikube

import (
	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/pkg/errors"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"go.opencensus.io/trace"
)

// Measures for the stats quickstart.
var (
	// The latency in seconds
	mLatencyS = stats.Float64("repl/startTime", "The latency in start time", stats.UnitSeconds)
)

// TagKeys for minikube start.
var (
	osKey = tag.MustNewKey("minikube.sigs.k8s.io/keys/os")
)

// Trace traces minikube start
func Trace() error {
	// Create and register a OpenCensus Stackdriver Trace exporter.
	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID:               "priya-wadhwa",
		DefaultMonitoringLabels: &stackdriver.Labels{},
	})
	if err != nil {
		return errors.Wrap(err, "getting exporter")
	}
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	trace.RegisterExporter(exporter)

	// Export to Stackdriver Monitoring.
	if err = exporter.StartMetricsExporter(); err != nil {
		return errors.Wrap(err, "starting metrics exporter")
	}

	// Subscribe views to see stats in Stackdriver Monitoring.
	if err := enableViews(); err != nil {
		return errors.Wrap(err, "enabling views")
	}

	return start()
}

func enableViews() error {
	startTimeView := &view.View{
		Name:        "minikube/startTime",
		Measure:     mLatencyS,
		Description: "minikube start over time",
		Aggregation: view.Distribution(1, 500),
		TagKeys:     []tag.Key{osKey},
	}

	return view.Register(startTimeView)
}
