package minikube

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os/exec"
	"runtime"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/pkg/errors"
	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"go.opencensus.io/trace"
)

func start() error {

	ctx, err := tag.New(context.Background(),
		tag.Insert(osKey, runtime.GOOS),
	)
	if err != nil {
		return errors.Wrap(err, "insert tag")
	}

	cmd := exec.Command("minikube", "start", "--output", "json")
	stdout, _ := cmd.StdoutPipe()

	spanName := "minikube.sigs.k8s.io/StartTime"
	ctx, span := trace.StartSpan(ctx, spanName)
	defer span.End()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
		if err := processStep(ctx, spanName, m); err != nil {
			return err
		}
	}
	return nil
}

func processStep(ctx context.Context, spanName, step string) error {
	ctx, err := tag.New(ctx,
		tag.Insert(osKey, runtime.GOOS),
	)
	if err != nil {
		return errors.Wrap(err, "new key")
	}
	name, err := stepName(step)
	if err != nil {
		return errors.Wrap(err, "step name")
	}
	ctx, span := trace.StartSpan(ctx, fmt.Sprintf("%s/%s", spanName, name))
	defer span.End()

	// Sleep for [1,10] seconds to fake work.
	time.Sleep(time.Duration(rand.Intn(10)+1) * time.Second)

	stats.Record(ctx, mLatencyS.M(25648))
	return nil
}

func stepName(step string) (string, error) {
	event := cloudevents.NewEvent()
	if err := json.Unmarshal([]byte(step), &event); err != nil {
		return "", errors.Wrap(err, "unmarshal cloud event")
	}
	m := map[string]string{}
	data := event.Data()
	if err := json.Unmarshal(data, &m); err != nil {
		return "", errors.Wrap(err, "unmarshal data")
	}
	stepName := m["name"]
	fmt.Println("step name is", stepName)
	return stepName, nil
}
