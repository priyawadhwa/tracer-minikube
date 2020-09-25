package main

import (
	"fmt"
	"os"
	"time"

	"github.com/priyawadhwa/tracer-minikube/pkg/minikube"
)

func main() {
	if err := minikube.Trace(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("sleeping 15 seconds")
	time.Sleep(15 * time.Second)
}
