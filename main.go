package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/priyawadhwa/tracer-minikube/pkg/minikube"
	"github.com/priyawadhwa/tracer-minikube/pkg/skaffold"
	"github.com/priyawadhwa/tracer-minikube/pkg/tracer"
)

func main() {
	ctx, span := tracer.StartSpan(context.Background(), "container-tools")
	defer span.End()
	if err := execute(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	time.Sleep(15 * time.Second)
}

func execute(ctx context.Context) error {
	if err := minikube.Start(ctx); err != nil {
		return errors.Wrap(err, "starting minikube")
	}
	if err := skaffold.DevLoop(ctx); err != nil {
		return errors.Wrap(err, "skaffold dev loop")
	}
	return nil
}
