module github.com/priyawadhwa/tracer-minikube

go 1.14

require go.opentelemetry.io/otel v0.11.0

require (
	cloud.google.com/go v0.66.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace v0.11.0
	github.com/cloudevents/sdk-go/v2 v2.2.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/stackdriverexporter v0.11.0
	github.com/pkg/errors v0.9.1
	go.opentelemetry.io/otel/sdk v0.11.0
	golang.org/x/net v0.0.0-20200925080053-05aa5d4ee321 // indirect
	golang.org/x/sys v0.0.0-20200923182605-d9f96fdee20d // indirect
	google.golang.org/genproto v0.0.0-20200925023002-c2d885f95484 // indirect
)
