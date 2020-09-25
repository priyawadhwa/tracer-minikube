package main

import (
	"fmt"
	"os"

	"github.com/priyawadhwa/tracer-minikube/pkg/minikube"
)

func main() {
	if err := minikube.Trace(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
